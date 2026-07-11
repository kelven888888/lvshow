package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitGoodsSpecGroupRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CGoodsSpecGroup{}
	BaserRouter := Router.Group("/goodsspecgroup")
	//Router.Use(middleware.OperationRecord())
	{

		BaserRouter.GET("/index", controllers.Index)

		BaserRouter.Any("/add", controllers.Add)
		//
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
		BaserRouter.Any("/deletebatch", controllers.Deletebatch)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
