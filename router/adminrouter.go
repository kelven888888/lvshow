package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAdminRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.Admincontroll{}
	BaserRouter := Router.Group("/admin")
	//Router.Use(middleware.OperationRecord())
	{

		BaserRouter.GET("/index", controllers.Adminlist)
		BaserRouter.GET("/getmenu", controllers.Getmenu)
		BaserRouter.GET("/logout", controllers.Logout)

		BaserRouter.Any("/add", controllers.Add)
		BaserRouter.Any("/getlist", controllers.Getlist)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
		BaserRouter.Any("/deletebatch", controllers.Deletebatch)
		BaserRouter.Any("/changepwd", controllers.Changepwd)
		BaserRouter.Any("/cleancache", controllers.CleanCache)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
