package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitSettingRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.SettingController{}
	BaserRouter := Router.Group("/setting")
	{
		BaserRouter.Any("/poster", controllers.Poster)
	}
	return BaserRouter
}
