package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitVersionRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.VersionController{}
	BaserRouter := Router.Group("version")
	{
		BaserRouter.Any("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
	}
	return BaserRouter
}
