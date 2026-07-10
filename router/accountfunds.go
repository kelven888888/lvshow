package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountFundsRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AccountFunds{}
	BaserRouter := Router.Group("accountfunds")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.Any("/fundsedit", controllers.FundsEdit)

	}
	return BaserRouter
}
