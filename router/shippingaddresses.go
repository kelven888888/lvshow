package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitShippingAddressesRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CShippingAddresses{}
	BaserRouter := Router.Group("/shippingaddresses")
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
