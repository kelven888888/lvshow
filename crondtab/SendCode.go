package crondtab

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"ginshop.com/admin/model"
	"ginshop.com/global"
	"ginshop.com/utils"
)

func SendCode() {
	fmt.Println("Send Sms")
	var sendcodes []model.AccountCheckCode
	global.SHOP_DB.Where("status=0").Limit(2).Find(&sendcodes)
	for _, v := range sendcodes {
		fmt.Println(v.Id)

		var sendcode model.AccountCheckCode
		global.SHOP_DB.Where("id=?", v.Id).Find(&sendcode).Update("status", 2)
		err := errors.New(fmt.Sprintf(""))
		go func(id int64) {
			key := utils.Get_SendCode_Key(v.Type)
			if key == "Invalid_code" {
				global.SHOP_LOG.Log(1, err.Error())
				sendcode.Status = 3
				sendcode.Errmsg = "短信key未配置"
				sendcode.Retry = 3
				global.SHOP_DB.Updates(&sendcode)
				return
			}
			var tpl model.Tplm
			err = global.SHOP_DB.Where("`keys`=?", key).First(&tpl).Error
			if err != nil {

				global.SHOP_LOG.Log(0, err.Error())
				return
			}
			if tpl.Id <= 0 {
				global.SHOP_LOG.Log(1, err.Error())
				sendcode.Status = 3
				sendcode.Errmsg = "短信模板不存在"
				sendcode.Retry = 3
				global.SHOP_DB.Updates(&sendcode)
				return
			}
			message := strings.Replace(utils.Languagebycode(v.Language, tpl.Content), "{code}", v.Captcha, 1)
			to := v.Name
			if utils.ValidateEmail(v.Name) {
				err = utils.SendMailWithRetry(to, message, "")
			} else {

				apiKey := global.SHOP_CONFIG.Sms.Appkey
				secretKey := global.SHOP_CONFIG.Sms.Secretkey
				url := global.SHOP_CONFIG.Sms.Url
				if global.SHOP_CONFIG.System.Version != "NQ" {
					err = utils.SendSms(url, apiKey, secretKey, to, message, 3)
				} else {
					appcode := global.SHOP_CONFIG.Sms.Appcode
					err = utils.SendSmsNq(url, apiKey, secretKey, to, message, appcode, 3)

				}
			}

			if err != nil {
				global.SHOP_LOG.Log(1, err.Error())
				sendcode.Status = 3
				sendcode.Errmsg = err.Error()
				sendcode.Retry = 3
				global.SHOP_DB.Updates(&sendcode)
			} else {
				sendcode.Status = 1

				global.SHOP_DB.Updates(&sendcode)
			}
		}(v.Id)
	}

}

func Initwebtpl() {
	var models []model.Tplm
	var modelslan []model.Language
	global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&modelslan)
	query := global.SHOP_DB.Model(model.Tplm{})
	fmt.Println("init language")
	err := query.Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())

	}

	for _, v := range models {
		var jsontest interface{}
		err := json.Unmarshal([]byte(v.Title), &jsontest)
		if err == nil {

			continue
		}
		var jsontitle, jsond []byte
		var title = make(map[string]interface{})
		var describe = make(map[string]interface{})
		if len(modelslan) > 1 {

			for _, lan := range modelslan {
				_, exists := title[lan.Code]
				if !exists {
					title[lan.Code] = v.Title

					describe[lan.Code] = v.Content
				}

			}

		}
		jsontitle, _ = json.Marshal(&title)
		jsond, _ = json.Marshal(&describe)
		err = global.SHOP_DB.Model(model.Tplm{}).Where("id = ?", v.Id).Updates(model.Tplm{
			Title:   string(jsontitle),
			Content: string(jsond),
		}).Error
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}
func Initbanner() {
	var models []model.WebsiteBanner
	var modelslan []model.Language
	global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&modelslan)
	query := global.SHOP_DB.Model(model.WebsiteBanner{})
	fmt.Println("init language")
	err := query.Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())

	}

	for _, v := range models {
		var jsontest interface{}
		err := json.Unmarshal([]byte(v.Title), &jsontest)
		if err == nil {

			continue
		}
		var jsontitle, jsond, imgjson []byte
		var title = make(map[string]interface{})
		var img = make(map[string]interface{})
		var content = make(map[string]interface{})
		if len(modelslan) > 1 {

			for _, lan := range modelslan {
				_, exists := title[lan.Code]
				if !exists {
					title[lan.Code] = v.Title

					img[lan.Code] = v.Image
					content[lan.Code] = v.Content
				}

			}

		}
		jsontitle, _ = json.Marshal(&title)
		jsond, _ = json.Marshal(&content)
		imgjson, _ = json.Marshal(&img)
		err = global.SHOP_DB.Model(model.WebsiteBanner{}).Where("id = ?", v.Id).Updates(model.WebsiteBanner{
			Title:   string(jsontitle),
			Content: string(jsond),
			Image:   string(imgjson),
		}).Error
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}
func Initstockpre() {
	var models []model.StockPredictionDayList
	var modelslan []model.Language
	global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&modelslan)
	query := global.SHOP_DB.Model(model.StockPredictionDayList{})
	fmt.Println("init language")
	err := query.Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())

	}
	for _, v := range models {
		var jsontest interface{}
		err := json.Unmarshal([]byte(v.Desc), &jsontest)
		if err == nil {

			continue
		}
		var jsond []byte
		var content = make(map[string]interface{})
		if len(modelslan) > 1 {

			for _, lan := range modelslan {
				_, exists := content[lan.Code]
				if !exists {
					content[lan.Code] = v.Desc

				}

			}

		}
		jsond, _ = json.Marshal(&content)
		err = global.SHOP_DB.Model(model.StockPredictionDayList{}).Where("id = ?", v.Id).Updates(model.StockPredictionDayList{
			Desc: string(jsond),
		}).Error
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}
