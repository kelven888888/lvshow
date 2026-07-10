package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitTbCouponRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.CTbCoupon{}
	BaserRouter := Router.Group("/tbcoupon")
	//Router.Use(middleware.OperationRecord())
	{

		BaserRouter.GET("/index", controllers.Index)
		BaserRouter.POST("/generate", controllers.Coupongen)
		BaserRouter.Any("/couponlist", controllers.CouponList)

		BaserRouter.Any("/add", controllers.Add)
		//
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.Any("/delete", controllers.Delete)
		BaserRouter.Any("/deletebatch", controllers.Deletebatch)

		//BaserRouter.GET("/main/index", controllers.Console)

	}
	return BaserRouter
}
