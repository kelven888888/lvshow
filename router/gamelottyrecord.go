package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitGameLottyRecordRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CGameLottyRecord{}
	BaserRouter := Router.Group("/gamelottyrecord")
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
