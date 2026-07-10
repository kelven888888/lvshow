package middleware

import (
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
)

// 全局中间件
func ChecmAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		hostname := c.Request.Host

		//fmt.Println("Hostname:", hostname)
		if hostname != global.SHOP_CONFIG.System.AdminUrl {
			c.Abort()
		}

		c.Next()
	}
}
