package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitPlaysettingRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CPlaysetting{}
	BaserRouter := Router.Group("/playsetting")
	//Router.Use(middleware.OperationRecord())
	{

		BaserRouter.GET("/index", controllers.Index)

		BaserRouter.Any("/add", controllers.Add)
		//
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
		BaserRouter.Any("/deletebatch", controllers.Deletebatch)
		BaserRouter.POST("/addproduct", controllers.AddProduct)
		BaserRouter.GET("/addproductview", controllers.Productview)
		BaserRouter.POST("/deleteproduct", controllers.DeleteProduct)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
