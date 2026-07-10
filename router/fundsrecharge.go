package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitFundsRechargeRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.FundsRecharge{}
	BaserRouter := Router.Group("fundsrecharge")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.POST("/do", controllers.Do)

	}
	return BaserRouter
}
