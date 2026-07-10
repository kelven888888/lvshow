package controller

import (
	"encoding/json"
	"fmt"
	"ginshop.com/admin/model"
	requests "ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var store = base64Captcha.DefaultMemStore

type Publiccontroll struct {
	BaseController
}

func (this *Publiccontroll) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})

}
func (this *Publiccontroll) Captcha(ctx *gin.Context) {
	var service service.Captcha
	uuids := uuid.New()
	var captcha = service.Captcha(uuids.String())
	this.Success(ctx, "成功", captcha)

}
func (this *Publiccontroll) Login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	adminId := session.Get("adminId")
	//fmt.Println(adminId)
	if adminId != nil {
		ctx.Abort()
		ctx.Redirect(302, "/admin/main/index")

	}
	ctx.HTML(http.StatusOK, "admin_login.html", nil)

}
func (this *Publiccontroll) Debugprice(ctx *gin.Context) {
	codestring := ctx.Query("code")
	var modelss model.LastPrice
	var models model.Tickers

	arr := strings.Split(codestring, ",")
	fmt.Println(arr, codestring)

	jsons := "{\"ticker\":\"AAPL\",\"todaysChangePerc\":-0.20162103310617363,\"todaysChange\":-0.5,\"updated\":1774263120000000000,\"day\":{\"o\":0,\"h\":0,\"l\":0,\"c\":200,\"v\":0,\"vw\":0},\"lastQuote\":{\"P\":247.31,\"S\":800,\"p\":247.15,\"s\":100,\"t\":1774263284445363084},\"lastTrade\":{\"c\":[12],\"i\":\"1481\",\"p\":247.49,\"s\":150,\"t\":1774263073841468955,\"x\":11,\"ds\":\"150.0\"},\"min\":{\"dv\":\"202.0\",\"dav\":\"147767.0\",\"av\":147767,\"t\":1774263060000,\"n\":4,\"o\":247.49,\"h\":247.49,\"l\":247.49,\"c\":247.49,\"v\":202,\"vw\":247.4884},\"prevDay\":{\"o\":247.975,\"h\":249.1999,\"l\":246,\"c\":247.99,\"v\":8.8752566571361e+07,\"vw\":248.0116}}"
	json.Unmarshal([]byte(jsons), &models)
	//for K, _ := range models.Tickers {

	for _, v := range arr {
		var ticker model.Tickers
		ticker = models
		ticker.Ticker = v
		ticker.TodaysChange = float64(rand.Intn(5))
		ticker.TodaysChangePerc = float64(rand.Intn(5))
		ticker.LastTrade.P = models.LastTrade.P * float64(rand.Intn(9)+90) / 100

		modelss.Tickers = append(modelss.Tickers, ticker)
	}
	//}
	jsonecho, _ := json.Marshal(modelss)
	ctx.String(200, string(jsonecho))

}
func (this *Publiccontroll) Loginsubmit(ctx *gin.Context) {
	var loginreq requests.Login
	err := ctx.ShouldBind(&loginreq)

	if err != nil {
		this.Error(ctx, "参数错误", err.Error())
		return
	}

	if !store.Verify(loginreq.CaptchaId, loginreq.Code, true) {
		this.Error(ctx, "验证码错误", nil)
		return
	}
	var adminserver service.AdminServer
	var adminuser *model.Admin
	adminuser, err = adminserver.Login(loginreq)
	//println(global.SHOP_CONFIG.System.SecretKey)

	if err != nil {
		this.Error(ctx, err.Error(), nil)
		return
	}
	session := sessions.Default(ctx)

	// 设置session数据
	session.Set("adminId", adminuser.Id)
	session.Set("groupId", adminuser.GroupId)
	// 保存session数据
	session.Save()
	sessionid := utils.MD5V(session.ID())
	err = global.SHOP_REDIS.Set(ctx, fmt.Sprintf("adminlogin_%d", adminuser.Id), sessionid, 3600*24*7*time.Second).Err()
	if err != nil {
		fmt.Println(sessionid, fmt.Sprintf("adminlogin_%d", adminuser.Id), "---------------------------------------------------------------", err, "---------------------------------------------------------------")
	}
	//var data = make(map[string]string)
	//data["access_token"] = "c262e61cd13ad99fc650e6908c7e5e65b63d2f32185ecfed6b801ee3fbdd5c0a"
	this.Success(ctx, "登录成功", nil)

}

func (this *Publiccontroll) Upload(ctx *gin.Context) {

	//ctx.JSON(http.StatusOK, gin.H{
	//	"error": 0,
	//	"url":   "../../uploads/file/9299961eab1ff4e85e870912f2abb560_20240529172250.jpg",
	//})
	//return
	_, header, err := ctx.Request.FormFile("file")
	contentType := header.Header.Get("Content-Type")
	allowedTypes := []string{"application/pdf", "image/jpeg", "image/png", "image/gif"}
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		this.Error(ctx, "上传文件错误", nil)
		return
	}

	types := ctx.Query("type")
	if err != nil {
		this.Error(ctx, "上传文件错误", nil)
		return
	}
	var service service.FileUploadAndDownloadService
	result, err := service.UploadFile(header, "0", types)
	if err != nil {
		this.Error(ctx, err.Error(), nil)
		return
	}
	this.Success(ctx, "上传成功",
		gin.H{
			"url": result,
		},
	)
}
func (this *Publiccontroll) Uploadeditor(ctx *gin.Context) {

	//ctx.JSON(http.StatusOK, gin.H{
	//	"error": 0,
	//	"url":   "../../uploads/file/9299961eab1ff4e85e870912f2abb560_20240529172250.jpg",
	//})
	//return
	_, header, err := ctx.Request.FormFile("imgFile")
	contentType := header.Header.Get("Content-Type")
	allowedTypes := []string{"application/pdf", "image/jpeg", "image/png", "image/gif"}
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		this.Error(ctx, "上传文件错误", nil)
		return
	}
	types := ctx.Query("type")
	if err != nil {
		this.Error(ctx, "上传文件错误", nil)
		return
	}
	var service service.FileUploadAndDownloadService
	result, err := service.UploadFile(header, "0", types)
	if err != nil {
		this.Error(ctx, err.Error(), nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": 0,
		"url":   "../../" + result,
	})
	return
}
