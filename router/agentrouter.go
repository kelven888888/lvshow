package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAgentRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AgentController{}
	BaserRouter := Router.Group("agent")
	{
		BaserRouter.Any("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/editpassword", controllers.EditPassword)
	}
	return BaserRouter
}
