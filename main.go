package main

import (
	"embed"
	"fmt"
	"time"

	"os"

	_ "embed"
	"ginshop.com/core"
	"ginshop.com/crondtab"
	"ginshop.com/gen"
	"ginshop.com/global"
	"ginshop.com/initialize"
)

func timeDifferenceInMinutes(t1, t2 time.Time) int {
	diff := t1.Sub(t2)
	return int(diff.Minutes())
}

//go:embed  views/admin/include/*
var Templatess embed.FS

func main() {
	//channle := make(chan uint, 10)
	//for i := 0; i < 1000; i++ {
	//	select {
	//	case channle <- 1:
	//		go func() {
	//
	//			fmt.Println(i)
	//			time.Sleep(time.Second * 1)
	//			<-channle
	//		}()
	//	}
	//}
	const (
		layoutsDir   = "templates/layouts"
		templatesDir = "views/admin/include/"
		extension    = "/*.html"
	)

	//fss, err := Templatess.ReadDir("views/admin/include")
	//{
	//
	//}
	//s, _ := fs.ReadDir(Templatess, "include")
	//fmt.Println(s)
	//
	//tmplFiles, err := fs.ReadDir(Templatess, "views/admin/include")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _, tmpl := range tmplFiles {
	//	if tmpl.IsDir() {
	//		soutpm, _ := fs.ReadDir(Templatess, fmt.Sprintf("views/admin/include/%s", tmpl.Name()))
	//		fmt.Println(fmt.Sprintf("views/admin/include/%s", tmpl))
	//		fmt.Println(soutpm)
	//	}
	//	fmt.Println(tmpl.Name())
	//}
	//time.Sleep(10000 * time.Second)
	core.Viper()                       // 初始化Viper
	global.SHOP_LOG = core.Zap()       // 初始化zap日志库
	global.SHOP_DB = initialize.Gorm() // gorm连接数据库
	initialize.OtherInit()
	initialize.Redis()

	//route := gin.Default()
	//route.GET("/ping", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"data": "",
	//	})
	//
	//})
	//ctx := context.Background()
	////go service.RunPublisher(ctx, "shop_message")
	//go service.RunSubscriber(ctx, "shop_message", "shop_message_queue")
	//time.Local = time.FixedZone("US/Eastern", -4*3600)
	//time.Local = time.FixedZone("Asia/Shanghai", 0)
	//time.Local = time.FixedZone("US/Eastern", -4*3600)
	fmt.Println(time.Now(), "timenoetimenoetimenoetimenoetimenoetimenoetimenoetimenoetimenoetimenoetimenoetimenoe")
	if global.SHOP_DB != nil {

		// 程序结束前关闭数据库链接
		db, _ := global.SHOP_DB.DB()
		defer db.Close()
	}
	args := os.Args[1:]

	if len(args) != 0 {
		if args[0] == "runcrond" {

			go crondtab.Initcrond()

			select {}
		}
	}
	if len(args) != 0 {
		if args[0] == "runother" {

			go crondtab.Levlehook()
			go crondtab.Couponcheck()
			select {}

		}
	}

	if len(args) != 0 {
		if args[0] == "sendmail" {
			go func() {
				for {

					if global.SHOP_CONFIG.System.Env == "pro" {
						crondtab.SendCode()
					}
					time.Sleep(time.Second)
				}
			}()
			select {}
		}

	}

	if len(args) != 0 {

		if args[0] == "gen" {
			gen := gen.Gen{}
			gen.Gener(args[1], args[2], args[3])
			select {}
		}

	}

	core.RunWindowsServer()
	select {}
}
