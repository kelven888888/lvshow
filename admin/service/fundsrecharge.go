package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

type FundsRecharge struct {
}

func (this *FundsRecharge) GetAll(pageInfo request.PageInfo) ([]model.FundRecharge, int64) {
	var models []model.FundRecharge

	query := global.SHOP_DB.Model(model.FundRecharge{})
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
func (this *FundsRecharge) Do(req request.IdsReq, action string) error {
	var models []model.FundRecharge
	global.SHOP_DB.Find(&models, "id in ?", req.Ids).Find(&models)

	var err error
	if action == "1" {

		//var modeltradeaccount model.TradeAccount
		for _, v := range models {
			tran := global.SHOP_DB.Begin()
			username := v.Username
			var accountfunds model.AccountFunds
			err = global.SHOP_DB.Where("username=?", username).First(&accountfunds).Error
			if err != nil {
				global.SHOP_LOG.Error(fmt.Sprintf("更新accountfunds错误%s", err.Error()))
				tran.Rollback()
				return err
			}
			var mFundRecharge model.FundRecharge
			results := tran.Where("id=? and status=0", v.Id).Find(&mFundRecharge).Updates(model.FundRecharge{
				Status: 1,
			})
			if results.Error != nil || results.RowsAffected != 1 {
				tran.Rollback()
				return err
			}
			if mFundRecharge.Id == 0 {

				global.SHOP_LOG.Error(fmt.Sprintf("%d状态错误", v.Id))
				tran.Rollback()
				return errors.New(fmt.Sprintf("%d状态错误", v.Id))
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
			amount := decimal.Zero
			remark := ""
			//exchange_rang:=utils.Float64ToDecimal(f)
			if v.Type == 1 {
				amount = v.Amount.Mul(utils.Float64ToDecimal(f))
				remark = fmt.Sprintf("充值%s/+U%s+RM%s", v.PathType, v.Amount.Round(4).StringFixed(4), amount.Round(4).StringFixed(4))
			}
			if v.Type == 2 {
				amount = v.Amount
				remark = fmt.Sprintf("充值/+RM%s", amount.Round(4).StringFixed(4))

			}

			err = tran.Exec("UPDATE  account_funds set ava_funds=ava_funds+? where id=?", amount, accountfunds.Id).Error
			if err != nil {

				global.SHOP_LOG.Error(fmt.Sprintf("更新accountfunds错误%s", err.Error()))
				tran.Rollback()
				return err
			}
			var saccountfundslog AccountFundsLog
			logidarr := []uint{}
			err, logid := saccountfundslog.Createlog(v.Username, amount, utils.Rechargetype, remark, 1)
			logidarr = append(logidarr, logid)
			if err != nil {
				global.SHOP_LOG.Error(fmt.Sprintf("更新accountfunds错误%s", err.Error()))
				tran.Rollback()
				return err
			}
			var message AccountMsgServer
			var modelmsg model.AccountUserMessage
			modelmsg.Username = username
			modelmsg.CreateTime = time.Now()
			modelmsg.UpdateTime = time.Now()
			group := 2
			keys := utils.Get_Code_Key(5)
			parms := make(map[string]string)
			parms["{wallet_type}"] = v.PathType
			parms["{amount}"] = v.Amount.Round(4).StringFixed(4)

			err = message.Save(&modelmsg, group, group, keys, parms)
			if err != nil {
				var mlog model.AccountFundsLog
				global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
				tran.Rollback()
				global.SHOP_LOG.Log(0, err.Error())
				return err
			}
			tran.Commit()
			//第一单给推荐人加一次盲盒次数
			var mod []model.FundRecharge
			global.SHOP_DB.Where("username=? and status=1", v.Username).Find(&mod)
			if len(mod) == 1 {
				var use model.User
				global.SHOP_DB.Where("username=? ", v.Username).Find(&use)
				if use.Pid != 0 {
					var puse model.User
					global.SHOP_DB.Where("id=?", use.Pid).Find(&puse).Updates(model.User{
						BlindBoxNum: puse.BlindBoxNum + 1,
					})
				}
			}
		}

	}
	if action == "2" {

		//var modeltradeaccount model.TradeAccount
		for _, v := range models {
			tran := global.SHOP_DB.Begin()

			var mFundRecharge model.FundRecharge
			err = global.SHOP_DB.Where("id=? and status=0", v.Id).Find(&mFundRecharge).Updates(model.FundRecharge{
				Status: 2,
			}).Error
			if mFundRecharge.Id == 0 {

				global.SHOP_LOG.Error(fmt.Sprintf("%d状态错误", v.Id))
				tran.Rollback()
				return errors.New(fmt.Sprintf("%d状态错误", v.Id))
			}

			tran.Commit()
		}

	}
	return nil
}
