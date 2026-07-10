package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"github.com/shopspring/decimal"
	"strconv"
	"time"

	"ginshop.com/global"
	"ginshop.com/utils"
)

type SUsdtWithdraw struct {
}

func (this *SUsdtWithdraw) GetAll(pageInfo request.PageInfo) ([]model.UsdtWithdrawModel, int64) {
	var models []model.UsdtWithdrawModel

	query := global.SHOP_DB.Model(model.UsdtWithdrawModel{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != 0 {

			query.Where("status =? ", *pageInfo.Status-1)
		}
	}

	if pageInfo.Username != "" {

		query.Where("username =? ", pageInfo.Username)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}

	return models, count

}
func (this *SUsdtWithdraw) Pass(req request.IdsReq) error {
	var models model.UsdtWithdrawModel

	err := global.SHOP_DB.Model(model.UsdtWithdrawModel{}).Where("id in ? and status='0'", req.Ids).Find(&models).Updates(model.UsdtWithdrawModel{Status: 1}).Error

	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return err
	}
	timenow := time.Now()
	var message AccountMsgServer
	var modelmsg model.AccountUserMessage
	modelmsg.Username = models.Username
	modelmsg.CreateTime = timenow
	modelmsg.UpdateTime = timenow
	group := 2
	keys := utils.Get_Code_Key(6)

	parms := make(map[string]string)
	parms["{wallet_type}"] = models.PathType
	parms["{amount}"] = models.Amount.Sub(models.Fee).Round(4).StringFixed(4)
	parms["{address}"] = models.WalletPath
	err = message.Save(&modelmsg, group, group, keys, parms)
	if err != nil {

		global.SHOP_LOG.Log(0, err.Error())

	}

	return nil
}
func (this *SUsdtWithdraw) PassUdun(req request.IdsReq) error {
	var models model.UsdtWithdrawModel

	err := global.SHOP_DB.Find(&models, "id in ? and status='0'", req.Ids).Find(&models).Updates(model.UsdtWithdrawModel{Status: 3}).Error
	if err != nil {
		return err
	}
	timenow := time.Now()
	var message AccountMsgServer
	var modelmsg model.AccountUserMessage
	modelmsg.Username = models.Username
	modelmsg.CreateTime = timenow
	modelmsg.UpdateTime = timenow
	group := 2
	keys := utils.Get_Code_Key(6)

	parms := make(map[string]string)
	parms["{wallet_type}"] = models.PathType
	parms["{amount}"] = fmt.Sprint(models.Amount)
	parms["{address}"] = models.WalletPath
	err = message.Save(&modelmsg, group, group, keys, parms)
	if err != nil {

		global.SHOP_LOG.Log(0, err.Error())

	}
	return nil
	//TODO:
	//请求钱包
	//
	//t_main_type := "0"
	//t_type := "0"
	//if models.PathType == "USDC Ethereum(ERC20)" {
	//	t_main_type = "60"
	//	t_type = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	//}
	//
	//if models.PathType == "USDT Ethereum(ERC20)" {
	//	t_main_type = "60"
	//	t_type = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	//}
	//
	//if models.PathType == "USDT TRX(TRC20)" {
	//	t_main_type = "195"
	//	t_type = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	//}
	//
	//addr := models.WalletPath
	//mainCoinType := models.PathType
	//coinType := models.Type
	//amount := models.Amount
	//businessId := models.Id
	//if addr == "" || t_main_type == "" || t_type == "" || amount == 0 || businessId == 0 {
	//	global.SHOP_LOG.Log(1, "参数错误")
	//	return errors.New("参数错误")
	//}
	//res := UdunWithdraw(mainCoinType, coinType, amount, addr, businessId)
	//if err != nil {
	//	return err
	//}

	return nil
}
func (this *SUsdtWithdraw) Refuse(req request.IdsReq, superusername string) error {
	var models []model.UsdtWithdrawModel

	global.SHOP_DB.Find(&models, "id in ? and status='0'", req.Ids).Find(&models)

	for _, v := range models {
		tran := global.SHOP_DB.Begin()

		username := v.Username
		var accountfunds model.AccountFunds
		err := tran.Where("username=?", username).Find(&accountfunds).Error
		if err != nil {
			tran.Rollback()
			return err
		}
		var mo model.UsdtWithdrawModel
		results := tran.Model(model.UsdtWithdrawModel{}).Where("id=? and status=0", v.Id).Find(&mo).Updates(model.UsdtWithdrawModel{Status: 2})
		if results.Error != nil || results.RowsAffected != 1 {
			tran.Rollback()
			return err
		}
		var conf Config
		rate, err := conf.GetKeyValue("MLECHANGERATE")
		if err != nil {
			tran.Rollback()
			return err
		}
		f, err := strconv.ParseFloat(rate, 64)
		if err != nil {
			tran.Rollback()
			return err
		}
		amounts := decimal.Zero
		remark := ""
		if v.Type == 1 {
			amounts = v.Amount.Mul(utils.Float64ToDecimal(f))
			remark = fmt.Sprintf("%s/%s/+RM%s", superusername, "提现拒绝", amounts)
		}
		if v.Type == 2 {
			amounts = v.Amount
			remark = fmt.Sprintf("%s/%s/+RM%s", superusername, "提现拒绝", amounts)

		}

		err = tran.Exec("update account_funds set ava_funds=ava_funds+? where id=?", amounts, accountfunds.Id).Error
		if err != nil {
			tran.Rollback()
			return err
		}
		timenow := time.Now()
		var saccountfundslog AccountFundsLog

		saccountfundslog.Createlog(accountfunds.Username, amounts, utils.Withdrayapplyrefuetype, remark, 1)
		var message AccountMsgServer
		var modelmsg model.AccountUserMessage
		modelmsg.Username = username
		modelmsg.CreateTime = timenow
		modelmsg.UpdateTime = timenow
		group := 2
		keys := utils.Get_Code_Key(7)
		parms := make(map[string]string)

		err = message.Save(&modelmsg, group, group, keys, parms)
		if err != nil {

			global.SHOP_LOG.Log(0, err.Error())

		}
		tran.Commit()

	}

	return nil
}
