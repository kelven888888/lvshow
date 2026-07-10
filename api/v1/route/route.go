package route

import (
	"ginshop.com/api/v1/controller"
	"ginshop.com/middleware"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	api := r.Group("/api", middleware.LoggerToFile())
	api.Use(middleware.DefaultLimit())
	{
		// v1
		v1 := api.Group("/v1")
		// 登录|注册
		publicGroup := v1.Group("/public")
		{
			publicGroup.POST("/captcha", controller.Captcha)
			publicGroup.POST("/getnotic", controller.Getnotic)
			publicGroup.POST("/sendcode", controller.SendCode)
			publicGroup.POST("/register", controller.Register)
			publicGroup.POST("/getgamebyid", controller.GetGameProductbyid)
			publicGroup.POST("/login", controller.Login)
			publicGroup.POST("/findpwd", controller.Findpwd)
			publicGroup.POST("/getagreement", controller.Agreement)
			publicGroup.POST("/banners", controller.Banners)
			publicGroup.POST("/siteconfig", controller.SiteConfig)
			publicGroup.POST("/lotteryrecord", controller.LotteryRecorddemo)
			publicGroup.POST("/commissionrecord", controller.Commissiondemo)
			publicGroup.Any("/dolotterystr", controller.Dolotterystr)
			publicGroup.POST("/walletcallback", controller.Walletcallback)

		}
		gameGrouppub := v1.Group("/game")
		{

			gameGrouppub.POST("/getgamebyid", controller.GetGameProductbyid)
			gameGrouppub.POST("/getgame", controller.GetGame)

		}
		gameauth := v1.Group("/game").Use(middleware.CheckToken())
		{
			gameauth.POST("/lotteryrecord", controller.LotteryRecord)
			gameauth.POST("/dolottery", controller.Dolottery)
			gameauth.Any("/dolotterystr", controller.Dolotterystr)

		}
		shopauth := v1.Group("/shop").Use(middleware.CheckToken())
		{
			shopauth.POST("/store", controller.Store)
			shopauth.POST("/orderssell", controller.Ordersell)
			shopauth.POST("/ordersbuy", controller.Orderbuy)
			shopauth.POST("/orders", controller.Orders)
			shopauth.POST("/orderscancel", controller.OrdersCancel)
			shopauth.POST("/pointsredeem", controller.PointsRedeem)
			shopauth.POST("/pointsredeemrecord", controller.PointsRedeemrecord)
			shopauth.POST("/doordershipping", controller.DoOrdershipping)
			shopauth.POST("/shippingorders", controller.ShippingOrders)
			shopauth.POST("/shippingorderconfirm", controller.ShippingOrdersConfirm)

		}
		Comauth := v1.Group("/community")
		{

			Comauth.POST("/index", controller.Index)

		}
		authGroup := v1.Group("/account").Use(middleware.CheckToken())
		{
			authGroup.POST("/changepwd", controller.ChangePwd)
			authGroup.POST("/fundsrecord", controller.Fundsrecord)
			authGroup.POST("/changetradepwd", controller.ChangeTradePwd)
			authGroup.POST("/findtradepwd", controller.Findtradepwd)
			authGroup.POST("/uploads", controller.Upload)
			authGroup.POST("/userinfo", controller.UserInfo)
			authGroup.POST("/changenickname", controller.ChangeNickName)
			authGroup.POST("/changeavatar", controller.ChangeAvatar)
			authGroup.POST("/getwallet", controller.GetWallet)
			authGroup.POST("/recharge", controller.Recharge)
			authGroup.POST("/rechargelist", controller.Rechargelist)

			authGroup.POST("/withdraw", controller.Withdraw)
			authGroup.POST("/withdrawlist", controller.Withdrawlist)
			authGroup.POST("/walletaddr", controller.Walletaddr)
			authGroup.POST("/walletaddrlist", controller.Walletaddrlist)
			authGroup.POST("/authsubmit", controller.Authsubmit)
			authGroup.POST("/authlist", controller.Authlist)
			authGroup.POST("/setlanguage", controller.Setlanguage)
			authGroup.POST("/memberlevel", controller.Memberlevel)
			authGroup.POST("/getcoupon", controller.Getcoupon)
			authGroup.POST("/getaddressselect", controller.Getaddressselect)
			authGroup.POST("/doshippingaddr", controller.Doshippingaddr)
			authGroup.POST("/handleshippingaddr", controller.Handleshippingaddr)
			authGroup.POST("/myshippingaddr", controller.Myshippingaddr)
			authGroup.POST("/commissionrecord", controller.Commissionrecord)

		}
		// 分类
		categorynoauthGroup := v1.Group("/category")
		{
			categorynoauthGroup.POST("/list", controller.GetCategoryLists)
			categorynoauthGroup.POST("/goods_list", controller.GetCategoryGoodsLists)
			categorynoauthGroup.POST("/goodsdetails", controller.GoodsDetails)

		}
		authcategorynoauthGroup := v1.Group("/category", middleware.CheckToken())
		{
			authcategorynoauthGroup.POST("/goods_list_enable", controller.GetCategoryGoodsListsenable)
		}
		// 商品
		goodsnoauthGroup := v1.Group("/goods")
		{
			goodsnoauthGroup.POST("/goodsdetails", controller.GoodsDetails)
		}
		// 用户
		userGroup := v1.Group("/user", middleware.CheckToken())
		{
			userGroup.POST("/bind-phone", controller.UserBindPhone)
			userGroup.POST("/un-bind-phone", controller.UnUserBindPhone)
		}

	}
}
