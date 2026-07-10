package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitWalletRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.WalletPath{}
	BaserRouter := Router.Group("walletpath")
	{

		BaserRouter.GET("/index", controllers.Index)

	}
	return BaserRouter
}
