package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitGoodsSpecParamRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CGoodsSpecParam{}
	BaserRouter := Router.Group("/goodsspecparam")
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
