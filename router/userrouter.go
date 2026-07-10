package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.UserController{}
	BaserRouter := Router.Group("user")
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.Any("/changepwd", controllers.Changepwd)
		BaserRouter.POST("/changetradepwd", controllers.Changetradepwd)
		BaserRouter.Any("/setagent", controllers.SetAgent)
		//BaserRouter.POST("/delete", controllers.Delete)

	}
	return BaserRouter
}
