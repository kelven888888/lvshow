package router

import (
	"fmt"
	"ginshop.com/middleware"
	"net/http"

	"ginshop.com/admin/controller"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	controllers := controller.Publiccontroll{}
	//wallcontrollers := controller.WalletCtr{}
	Router.GET("/", func(context *gin.Context) {
		context.Abort()
		context.Redirect(302, "/public/login")

	}).Use(middleware.ChecmAdmin())
	//Router.Any("/wallet/rechargecallbackabcd123", wallcontrollers.RechargeCallBack)
	//Router.Any("/wallet/payoutcalback", wallcontrollers.PayoutBack)

	BaserRouter := Router.Group("public").Use(middleware.ChecmAdmin())
	{
		BaserRouter.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "pong",
			})

		})
		BaserRouter.POST("/message/index", func(context *gin.Context) {
			fmt.Println(1)

		})
		BaserRouter.GET("/register", controllers.Register)
		BaserRouter.GET("/captcha", controllers.Captcha)
		BaserRouter.GET("/login", controllers.Login)
		BaserRouter.POST("/loginsubmit", controllers.Loginsubmit)
		BaserRouter.POST("/upload", controllers.Upload)
		BaserRouter.POST("/uploadeditor", controllers.Uploadeditor)
		BaserRouter.GET("/debugproce", controllers.Debugprice)

	}
	return BaserRouter
}
