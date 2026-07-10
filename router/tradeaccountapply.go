package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitTradeAccountApplyRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.TradeAccountApply{}
	BaserRouter := Router.Group("tradeaccountapply")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.POST("/pass", controllers.Pass)
		BaserRouter.POST("/refuse", controllers.Refuse)

	}
	return BaserRouter
}
