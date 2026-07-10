package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAdminlogRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.Adminlogcontroll{}
	BaserRouter := Router.Group("/adminlog")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.POST("/delete", controllers.Delete)
		BaserRouter.POST("/deletebatch", controllers.Deletebatch)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
