package router

import (
	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitWebsiteBannerRoute(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.WebsiteBannerController{}
	BaserRouter := Router.Group("websitebanner")
	{
		BaserRouter.GET("/list", controllers.List)
		BaserRouter.Any("/edit", controllers.Edit)
		BaserRouter.POST("/delete", controllers.Delete)
	}
	return BaserRouter
}
