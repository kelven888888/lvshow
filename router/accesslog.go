package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccesslogRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.Acclogctr{}
	BaserRouter := Router.Group("accesslog")
	{

		BaserRouter.GET("/index", controllers.Index)

	}
	return BaserRouter
}
