package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitUserWithdrawRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.UsdtWithdraw{}
	BaserRouter := Router.Group("usdtwithdraw")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.POST("/pass", controllers.Pass)
		BaserRouter.POST("/passudun", controllers.PassUdun)
		BaserRouter.POST("/refuse", controllers.Refuse)

	}
	return BaserRouter
}
