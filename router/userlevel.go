package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitUserLevelRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.UserlevelController{}
	BaserRouter := Router.Group("/userlevel")
	//Router.Use(middleware.OperationRecord())
	{

		BaserRouter.GET("/index", controllers.Index)

		BaserRouter.Any("/add", controllers.Add)
		//
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
