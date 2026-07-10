package middleware

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/global"
	"ginshop.com/utils"
	"ginshop.com/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"strings"
	"time"
)

// gin自定义日志中间件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {

		//开始时间
		startTime := time.Now()
		language := c.GetHeader("Accept-Language")
		allowlanguage := strings.Split(global.SHOP_CONFIG.System.Language_Array, ",")

		//headacep := c.Request.Header.Get("Content-Type")
		//fmt.Println(c.Request.Header, headacep, "headacepheadacepheadacepheadacepheadacepheadacepheadacepheadacepheadacep")
		if !utils.InArray(language, allowlanguage) {
			language = global.SHOP_CONFIG.System.Language
		}

		c.Set("Language", language)
		fmt.Println(c.Get("Language"))
		tokenString := c.GetHeader("Authorization")
		// 验证token非空
		var username = ""
		userid := 0
		if tokenString != "" {
			token, claims, err := utils.ParseToken(tokenString)
			if err == nil && token.Valid {

				username = claims.UserName
				userid, _ = strconv.Atoi(claims.UserId)
				c.Set("user_id", claims.UserId)
				c.Set("user_name", claims.UserName)

			}
			//如果用户存在 将user信息存入上下文

		}
		c.Next()
		//var err error
		//body, err := io.ReadAll(c.Request.Body)
		//if err != nil {
		//	global.SHOP_LOG.Error("read body from request error:", zap.Error(err))
		//} else {
		//	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		//}
		//var data map[string]interface{}
		//formData := make(map[string][]string)
		//
		//// ShouldBindWith 可以指定绑定引擎为 Form
		//if err := c.ShouldBindWith(&formData, binding.Form); err != nil {
		//	global.SHOP_LOG.Error(err.Error())
		//}
		//var body []byte
		//
		//if c.Request.Method != http.MethodGet {
		//
		//	var err error
		//	body, err = io.ReadAll(c.Request.Body)
		//	if err != nil {
		//		global.SHOP_LOG.Error("read body from request error:", zap.Error(err))
		//	} else {
		//		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		//	}
		//} else {
		//
		//}
		//m := make(map[string]interface{})
		//
		//// 现在 formData 包含了所有表单字段
		//for k, v := range formData {
		//	for _, val := range v {
		//		fmt.Println(k, val)
		//		m[k] = val
		//	}
		//
		//}
		//
		//// 显式指定使用 JSON 绑定引擎
		//if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		//	global.SHOP_LOG.Error(err.Error())
		//}
		//datas, _ := json.Marshal(&m)
		//
		//c.JSON(200, datas)
		//c.Abort()
		//处理请求
		//data := utils.DataMapByRequest(c)
		//datas, _ := json.Marshal(&data)

		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime).Seconds()
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求IP
		ClientIp := c.ClientIP()
		//用户标识
		UserAgent := c.Request.UserAgent()
		if global.SHOP_CONFIG.System.Env == "debug" {
			ClientIp = "107.172.231.100"
		}

		//日志格式
		logger.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    ClientIp,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"user_agent":   UserAgent,
		}).Info()

		go func(username string, userid int) {
			var acclog model.MAccesslog
			var user model.User
			global.SHOP_DB.Where("id=?", userid).Find(&user)
			acclog.Username = user.Username
			acclog.Ip = ClientIp
			acclog.Path = reqUri
			acclog.CreateAt = time.Now()
			acclog.Method = reqMethod
			//if global.SHOP_CONFIG.System.Env == "debug" {
			//	acclog.Body = string(datas)
			//}

			db, err := geoip2.Open("GeoLite2-City.mmdb")
			if err != nil {
				global.SHOP_LOG.Log(2, err.Error())
			}
			defer db.Close()
			ip := net.ParseIP(acclog.Ip)
			record, err := db.City(ip)
			if err != nil {
				global.SHOP_LOG.Log(2, err.Error())
			}
			global.SHOP_DB.Save(&acclog)
			address := fmt.Sprintf("%s_%s", record.Country.Names["zh-CN"], record.City.Names["zh-CN"])
			if len(record.Subdivisions) > 0 {
				address = fmt.Sprintf("%s_%s", address, record.Subdivisions[0].Names["zh-CN"])
			}
			if len(record.Subdivisions) > 0 {
				global.SHOP_DB.Model(model.MAccesslog{}).Where("id=?", acclog.Id).Updates(model.MAccesslog{
					City:        record.City.Names["zh-CN"],
					Country:     record.Country.Names["zh-CN"],
					Subdivision: record.Subdivisions[0].Names["zh-CN"],
					Address:     address,
				})
			} else {
				global.SHOP_DB.Model(model.MAccesslog{}).Where("id=?", acclog.Id).Updates(model.MAccesslog{
					City:    record.City.Names["zh-CN"],
					Country: record.Country.Names["zh-CN"],
					Address: address,
				})
			}

		}(username, userid)

		// token验证是否失效

	}
}
