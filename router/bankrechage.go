package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitBankRechargeRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.BankRecharge{}
	BaserRouter := Router.Group("bankrecharge")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/add", controllers.Add)

	}
	return BaserRouter
}
