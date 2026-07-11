package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// 全局变量
var db *gorm.DB
var rdb *redis.Client

// 初始化数据库和 Redis
func initDB() {
	var err error
	// 假设 MySQL 连接信息，实际项目中应从配置文件读取
	dsn := "root:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// 启动服务
func mainsa() {
	initDB()

	// 自动迁移表结构
	db.AutoMigrate(&SPU{}, &SKU{})

	r := gin.Default()

	// 路由组
	api := r.Group("/api/v1")
	{
		api.POST("/products", createProduct)
		api.GET("/products/:id", getProduct)
		api.POST("/order/place", placeOrder)
	}

	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}

// createProduct 创建商品 (SPU + SKUs)
func createProduct(c *gin.Context) {
	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 事务处理
	err := db.Transaction(func(tx *gorm.DB) error {
		spu := SPU{Name: req.Name, Description: req.Description}
		if err := tx.Create(&spu).Error; err != nil {
			return err
		}

		for _, skuReq := range req.SKUs {
			sku := SKU{
				SPUID:   spu.ID,
				Specs:   skuReq.Specs,
				Price:   skuReq.Price,
				Stock:   skuReq.Stock,
				SkuCode: skuReq.SkuCode,
			}
			if err := tx.Create(&sku).Error; err != nil {
				return err
			}
			// 同步库存到 Redis
			rdb.Set(c.Request.Context(), fmt.Sprintf("stock:sku_%d", sku.ID), sku.Stock, 0)
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

// getProduct 获取商品详情
func getProduct(c *gin.Context) {
	id := c.Param("id")
	var spu SPU
	if err := db.Preload("SKUs").First(&spu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, spu)
}

// placeOrder 下单并扣减库存
func placeOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行 Lua 脚本扣减库存
	luaScript := `
	local key = KEYS[1]
	local quantity = tonumber(ARGV[1])
	local current_stock = tonumber(redis.call('GET', key))
	
	if current_stock == nil then
		return -1
	end
	
	if current_stock < quantity then
		return -2
	end
	
	redis.call('DECRBY', key, quantity)
	return tonumber(redis.call('GET', key))
	`

	key := fmt.Sprintf("stock:sku_%d", req.SKUID)
	result, err := rdb.Eval(c.Request.Context(), luaScript, []string{key}, req.Quantity).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
		return
	}

	switch val := result.(type) {
	case int64:
		if val == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		} else if val == -2 {
			c.JSON(http.StatusConflict, gin.H{"error": "Insufficient stock"})
		} else {
			// 这里应异步创建订单记录到 MySQL
			c.JSON(http.StatusOK, gin.H{"message": "Order placed", "remaining_stock": val})
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
}

// 数据结构定义
type SPU struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SKUs        []SKU  `json:"skus" gorm:"foreignKey:SPUID"`
}

type SKU struct {
	ID      uint64          `json:"id" gorm:"primaryKey"`
	SPUID   uint64          `json:"spu_id"`
	Specs   json.RawMessage `json:"specs" gorm:"type:json"`
	Price   float64         `json:"price"`
	Stock   int             `json:"stock"`
	SkuCode string          `json:"sku_code"`
}

type ProductRequest struct {
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description"`
	SKUs        []SKURequest `json:"skus" binding:"required,dive"`
}

type SKURequest struct {
	Specs   json.RawMessage `json:"specs" binding:"required"`
	Price   float64         `json:"price" binding:"required,gt=0"`
	Stock   int             `json:"stock" binding:"required,gte=0"`
	SkuCode string          `json:"sku_code" binding:"required"`
}

type OrderRequest struct {
	SKUID    uint64 `json:"sku_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,gt=0"`
}
