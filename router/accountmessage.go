package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAccountMessageRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AccountMessageController{}
	BaserRouter := Router.Group("accountmessage")
	{
		BaserRouter.Any("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
	}
	return BaserRouter
}
