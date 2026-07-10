package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountFeedbackRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AccountFeedbackController{}
	BaserRouter := Router.Group("accountfeedback")
	{
		BaserRouter.GET("/list", controllers.List)
		BaserRouter.POST("/delete", controllers.Delete)
	}
	return BaserRouter
}
