package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitNewsRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.NewsController{}
	BaserRouter := Router.Group("news")
	{
		BaserRouter.Any("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
	}
	return BaserRouter
}
