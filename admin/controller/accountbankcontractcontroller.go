package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
	"time"
)

type AccountBankContractController struct {
	BaseController
}

func (h AccountBankContractController) List(ctx *gin.Context) {
	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Count      int64  `json:"count" form:"count"`
		Username   string `json:"username" form:"username"`
		ContractNo string `json:"contractNo" form:"contractNo"`
		Handler    int    `json:"handler" form:"handler"`
	}

	req := Request{
		PageInfo: request.PageInfo{
			Limit: 20,
			Page:  1,
		},
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.AccountBankContract{})

	if req.Username != "" {
		query = query.Where("username like ?", "%"+req.Username+"%")
	}

	if req.ContractNo != "" {
		query = query.Where("contract_no like ?", "%"+req.ContractNo+"%")
	}

	if req.Handler != 0 {
		if req.Handler == 1 {
			query = query.Where("handler_time is not null")
		}
		if req.Handler == 2 {
			query = query.Where("handler_time is null")
		}
	}

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.AccountBankContract, 0)
	if err = query.
		Where("status = 1").
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}
	webApiUrl := global.SHOP_CONFIG.System.WebApiURL

	ids := make([]int64, 0)
	for _, row := range list {
		ids = append(ids, row.Id)
	}
	if err = db.Model(&model.AccountBankContract{}).Where("id in ?", ids).Update("read", "1").Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "account_bank_contract_list.html", gin.H{
		"status":    "200",
		"list":      list,
		"req":       req,
		"webApiUrl": webApiUrl,
	})
}

func (h AccountBankContractController) Handler(ctx *gin.Context) {

	type Request struct {
		Id int64 `json:"id" form:"id"`
	}

	req := Request{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		h.Error(ctx, err.Error())
		return
	}

	db := global.SHOP_DB

	session := sessions.Default(ctx)
	adminId := session.Get("adminId")
	userId := adminId.(uint)
	admin := model.Admin{}
	global.SHOP_DB.Where("id=?", userId).Find(&admin)
	superusername := admin.Account

	err = db.Transaction(func(tx *gorm.DB) error {
		find := &model.AccountBankContract{}
		if err := db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", req.Id).First(find).Error; err != nil {
			return err
		}

		if find.HandlerTime != nil {
			return errors.New("该订单已被处理")
		}
		fmt.Println("1")

		amount := decimal.NewFromFloat(find.Amount)
		userFunds := model.AccountFunds{}
		if err := tx.Where("username = ?", find.Username).First(&userFunds).Error; err != nil {
			return errors.WithMessage(err, "查询AccountFunds失败")
		}

		if err := tx.Model(&model.AccountFunds{}).Where("id = ?", userFunds.Id).Update("ava_funds", userFunds.AvaFunds.Add(amount)).Error; err != nil {
			return errors.WithMessage(err, "更新余额")
		}
		fmt.Println("2")
		var saccountfundslog service.AccountFundsLog
		remark := fmt.Sprintf("%s/%s/%s/%s/%s", superusername, "处理合同", find.Username, find.ContractNo, amount)
		err := saccountfundslog.CreatelogByTx(tx, find.Username, amount, utils.Rechargetype, remark)
		if err != nil {
			return err
		}
		fmt.Println("3")

		key := utils.Get_Code_Key(13)
		var tpl model.Tplm
		err = tx.Where("`keys`=?", key).Where("status = 1").First(&tpl).Error
		if err != nil {

			return errors.WithMessage(err, "找不到消息模板")
		}

		var modelusteauthrity model.AccountUserAuthority
		err = global.SHOP_DB.Where("username=?", find.Username).First(&modelusteauthrity).Error
		if err != nil {

			global.SHOP_LOG.Log(2, err.Error())

		}
		language := "en"
		if modelusteauthrity.Id > 0 {
			language = modelusteauthrity.Language
		}

		message := utils.Languagebycode(language, tpl.Content)

		message = strings.ReplaceAll(message, "{Amount}", amount.String())
		message = strings.ReplaceAll(message, "{ContractNo}", find.ContractNo)

		// messages := strings.Replace(tpl.Content, "{wallet_type}", models.PathType, 1)
		// messages = strings.Replace(messages, "{address}", models.WalletPath, 1)
		// messages = strings.Replace(messages, "{amount}", fmt.Sprint(models.Amount), 1)

		Title := utils.Languagebycode(language, tpl.Title)

		accountUserMessage := &model.AccountUserMessage{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			Username:   find.Username,
			Title:      Title,
			Content:    message,
			Group:      2,
			Remarks:    "success",
			Type:       2,
			Status:     0,
			Read:       0,
		}

		if err := tx.Create(accountUserMessage).Error; err != nil {
			return errors.WithMessage(err, "添加充值消息")
		}
		fmt.Println("4")

		now := time.Now()
		find.HandlerTime = &now
		find.UpdateTime = &now

		if err := tx.Save(&find).Error; err != nil {
			return err
		}

		return nil
	})

	fmt.Println("err", err)
	if err != nil {
		h.Error(ctx, err.Error())
		return
	}

	h.Success(ctx)
	return
}
