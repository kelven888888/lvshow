package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitVercodeRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.VerCodeController{}
	BaserRouter := Router.Group("vercode")
	{
		BaserRouter.Any("/list", controllers.List)

	}
	return BaserRouter
}
