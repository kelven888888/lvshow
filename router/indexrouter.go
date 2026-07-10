package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitIndexRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.IndexController{}
	Router.GET("/main/index", controllers.Index)
	Router.Any("/home/index", controllers.Console)

	return Router
}
