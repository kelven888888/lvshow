package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountFundsLogRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AccountFundsLog{}
	BaserRouter := Router.Group("accountfundslog")
	{

		BaserRouter.GET("/index", controllers.Index)

	}
	return BaserRouter
}
