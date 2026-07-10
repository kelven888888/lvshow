package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountBankContractRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AccountBankContractController{}
	BaserRouter := Router.Group("accountbankcontract")
	{
		BaserRouter.GET("/list", controllers.List)
		BaserRouter.POST("/handler", controllers.Handler)
	}
	return BaserRouter
}
