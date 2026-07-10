package service

import (
	"image/color"
	"time"

	"ginshop.com/global"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type Captcha struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (b *Captcha) Captcha(key string) interface{} {
	// 判断验证码是否开启
	openCaptcha := global.SHOP_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.SHOP_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间

	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	//driver := base64Captcha.NewDriverDigit(global.SHOP_CONFIG.Captcha.ImgHeight, global.SHOP_CONFIG.Captcha.ImgWidth, global.SHOP_CONFIG.Captcha.KeyLong, 0.9, 80)
	source := "1234567890abcd"
	rgbaColor := color.RGBA{0, 0, 0, 0}
	// driver
	driver := base64Captcha.NewDriverString(
		global.SHOP_CONFIG.Captcha.ImgHeight, // height int
		global.SHOP_CONFIG.Captcha.ImgWidth,  // width int
		0,                                    // noiseCount int
		0,                                    // showLineOptions int
		4,                                    // length int
		source,                               // source string
		&rgbaColor,                           // bgColor *color.RGBA
		nil,                                  // fontsStorage FontsStorage
		nil,                                  // fonts []string
	)

	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	//fonts := []string{"wqy-microhei.ttc"}
	// 生成driver,g高，宽 背景文字的，画线的调试，背景颜色的指针
	//rgbaColor := color.RGBA{0, 0, 0, 0}
	//driver := base64Captcha.NewDriverMath(80, 240, 0, 0, &rgbaColor, nil, nil)

	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, code, err := cp.Generate()
	if err != nil {
		global.SHOP_LOG.Error("验证码获取失败!", zap.Error(err))
		return err.Error()
	}
	if global.SHOP_CONFIG.System.Env != "debug" {
		code = "0000"
	}
	return map[string]interface{}{
		"CaptchaId":     id,
		"PicPath":       b64s,
		"CaptchaLength": global.SHOP_CONFIG.Captcha.KeyLong,
		"OpenCaptcha":   oc,
		"Captcha":       code,
	}
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
