package service

import (
	"ginshop.com/admin/model"
	"ginshop.com/global"
	"ginshop.com/utils"
	"strings"
)

const (
	system   = iota // 系统消息分组
	event    = 1    // 活动分组
	transfer = 2    // 交易分组
	orders   = 3    // 订单分组
	quant    = 4    // 量化分组
	fund     = 5    //基金分组
	charity  = 6    // 慈善分组

)

type AccountMsgServer struct {
}

func (this *AccountMsgServer) Save(models *model.AccountUserMessage, group int, types int, keys string, parms map[string]string) error {
	if models.Id > 0 {
		return global.SHOP_DB.Updates(&models).Error
	} else {
		//models.CreateTime = time.Now()

		var tpl model.Tplm
		err := global.SHOP_DB.Where("`keys`=?", keys).First(&tpl).Error
		if err != nil {

			global.SHOP_LOG.Log(2, err.Error())

		}
		if tpl.Id <= 0 {

			global.SHOP_LOG.Log(0, "未配置通知key")

		}
		var modelusteauthrity model.User
		err = global.SHOP_DB.Where("username=?", models.Username).First(&modelusteauthrity).Error
		if err != nil {

			global.SHOP_LOG.Log(2, err.Error())

		}
		language := "en"
		if modelusteauthrity.Id > 0 {
			language = modelusteauthrity.Language
		}

		content := utils.Languagebycode(language, tpl.Content)
		for k, v := range parms {
			content = strings.Replace(content, k, v, 1)
		}

		// messages := strings.Replace(tpl.Content, "{wallet_type}", models.PathType, 1)
		// messages = strings.Replace(messages, "{address}", models.WalletPath, 1)
		// messages = strings.Replace(messages, "{amount}", fmt.Sprint(models.Amount), 1)

		models.Title = utils.Languagebycode(language, tpl.Title)
		models.Content = content
		models.Group = group
		models.Read = 0

		models.Type = types
		return global.SHOP_DB.Save(&models).Error
	}

}
