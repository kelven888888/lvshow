package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitAgreementRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.AgreementController{}
	BaserRouter := Router.Group("agreement")
	{
		BaserRouter.Any("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
	}
	return BaserRouter
}
