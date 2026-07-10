package model

import (
	"encoding/json"
	"ginshop.com/utils"
	"ginshop.com/utils/Paginate"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 模型结构体
type Goods struct {
	Model

	GoodsCate      int             `form:"goods_cate"  json:"goods_cate" gorm:"size:10;not null;default:0;comment:'商品类别'"  comment:"商品类别"`                        // 商品类别
	GoodsName      string          `form:"goods_name"  json:"goods_name" gorm:"size:255;not null;default:'';index:GoodsNameIndex;comment:'商品名称'"  comment:"商品类别"` // 商品名称
	GoodsProperty  string          `form:"goods_property"  json:"goods_property" gorm:"size:255;not null;default:'';index:GoodsProperty;comment:'商品属性'" edit:"0"` // 商品属性
	GoodsDesc      string          `form:"goods_desc" json:"goods_desc" gorm:"size:255;not null;default:'';index:GoodsDesc;comment:'商品描述'"  comment:"商品描述"`       // 商品简介
	GoodsContent   string          `form:"goods_content"  json:"goods_content" gorm:"longtext;not null;default:'';comment:'商品信息'"  comment:"商品信息"`                // 商品信息
	UnitPrice      decimal.Decimal ` form:"unit_price" json:"unit_price" gorm:"decimal(18,2);not null;default:0;comment:'商品单价'"  comment:"商品单价"`                  // 商品单价
	FavorablePrice float64         `form:"favorable_price" json:"favorable_price" gorm:"decimal(18,2);not null;default:0;comment:'优惠价格'" comment:"优惠价格"`          // 优惠价格
	GoodsStock     uint64          `form:"goods_stock" json:"goods_stock" gorm:"size:10;not null;default:0;comment:'商品库存'" comment:"商品库存"`                        // 商品库存
	GoodsCover     string          `form:"goods_cover" json:"goods_cover" gorm:"size:255;not null;default:'';comment:'商品封面图'" comment:"商品封面图"`                    // 商品封面图
	GoodsSlides    string          `form:"goods_slides" json:"goods_slides" gorm:"size:255;not null;default:'';comment:'商品幻灯片'"`                                  // 商品幻灯片
	GoodsStatus    *int            `form:"goods_status"  json:"goods_status" gorm:"size:10;not null;default:0;comment:'商品状态'" comment:"状态"`                       // 商品状态
	RewardType     string          `form:"reward_type"  json:"reward_type" gorm:"size:10;not null;default:0;comment:'商品奖励类型'" comment:"商品奖励类型"`
	Points         decimal.Decimal ` form:"points" json:"points" gorm:"decimal(18,2);not null;default:0;comment:'积分兑换'"  comment:"积分兑换"` // 商品单价

}

// 获取表名
func (Goods) TableName() string {
	return "goods"
}
func (m *Goods) MarshalJSON() ([]byte, error) {
	type Alias Goods // 创建一个新的类型别名以避免无限递归调用 MarshalJSON() 方法
	aux := struct {
		*Alias          // 将原始User的所有字段嵌入到aux中，但不包括Birthday（除非你已经处理了它）
		PlayGoodsId int ` json:"goods_id" form:"goods_id" ` // 商品

	}{
		Alias:       (*Alias)(m), // 将原始User的所有公共字段赋值给aux的Alias部分（不包括Birthday）
		PlayGoodsId: m.Id,        // 使用自定义的时间格式化逻辑（例如：年-月-日）

	}
	return json.Marshal(aux) // 序列化aux结构体，其中包括格式化后的Birthday字段和所有其他原始字段。
}

// 根据检索条件，获取记录行，并获取总记录条数
func (Goods) FindAll(DB *gorm.DB, params map[string]any) ([]Goods, int64) {
	var GoodResult []Goods
	page := params["page"].(string)
	pageSize := params["pageSize"].(string)
	ParamsFilter := utils.ParamsFilter("page,pageSize", params)
	DB.Scopes(Paginate.Paginate(page, pageSize)).Where(ParamsFilter).Order("created_at desc").Find(&GoodResult)
	GoodCount := DB.Find(&Goods{})
	return GoodResult, GoodCount.RowsAffected
}

// 插入商品操作
func (Goods) AddGoods(DB *gorm.DB, params map[string]any) error {
	GoodsCate, _ := params["GoodsCate"].(int)

	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)

	// 拼装数据
	goods := Goods{
		GoodsCate:     GoodsCate,
		GoodsName:     params["GoodsName"].(string),
		GoodsProperty: params["GoodsProperty"].(string),
		GoodsDesc:     params["GoodsDesc"].(string),
		GoodsContent:  params["GoodsContent"].(string),

		FavorablePrice: FavorablePrice,
		GoodsStock:     GoodsStock,
		GoodsCover:     params["GoodsCover"].(string),
		GoodsSlides:    params["GoodsSlides"].(string),
	}
	// 写入数据库
	result := DB.Create(&goods)
	return result.Error
}

// 编辑商品
func (Goods) EditGoods(DB *gorm.DB, params map[string]any) error {
	int_id, _ := params["id"].(uint)
	GoodsCate, _ := params["GoodsCate"].(int)
	//UnitPrice, _ := params["UnitPrice"].(float64)
	FavorablePrice, _ := params["FavorablePrice"].(float64)
	GoodsStock, _ := params["GoodsStock"].(uint64)

	var goods Goods
	// 查询当前数据
	DB.First(&goods, int_id)
	// 更新数据
	goods.GoodsCate = GoodsCate
	goods.GoodsName = params["GoodsName"].(string)
	goods.GoodsProperty = params["GoodsProperty"].(string)
	goods.GoodsDesc = params["GoodsDesc"].(string)
	goods.GoodsContent = params["GoodsContent"].(string)

	goods.FavorablePrice = FavorablePrice
	goods.GoodsStock = GoodsStock
	goods.GoodsCover = params["GoodsCover"].(string)
	goods.GoodsSlides = params["GoodsSlides"].(string)

	result := DB.Save(&goods)
	return result.Error
}
