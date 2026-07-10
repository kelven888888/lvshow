package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type AccountFundsLog struct {
}

func (this *AccountFundsLog) GetAll(pageInfo request.PageInfo) ([]model.AccountFundsLog, int64) {
	var models []model.AccountFundsLog

	query := global.SHOP_DB.Model(model.AccountFundsLog{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != 0 {

			query.Where("status =? ", pageInfo.Status)
		}
	}
	if pageInfo.LogType != 0 {

		query.Where("log_type =? ", pageInfo.LogType)
	}
	if pageInfo.MoneyType != 0 {

		query.Where("money_type =? ", pageInfo.MoneyType)
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
func (this *AccountFundsLog) Createlog(username string, amount decimal.Decimal, types int, remark string, money_type int) (error, uint) {
	var sagentinfo AgentInfo
	agentinfo, err := sagentinfo.GetCode(username)
	if err != nil {
		return err, 0
	}

	var accountfund model.AccountFunds
	global.SHOP_DB.Where("username=?", username).Find(&accountfund)
	if err != nil {
		return err, 0
	}

	//amounts := amount.String()
	//var saccountfundslog AccountFundsLog
	//saccountfundslog.Createlog(ag_info, username)
	var log model.AccountFundsLog
	log.Username = username
	log.Amount = amount
	log.AgCode = agentinfo.AgCode
	log.Remarks = remark
	log.LevelCode = agentinfo.Level1Code
	log.CreateTime = time.Now()
	log.LogType = types
	log.MoneyType = money_type
	log.Uid = accountfund.Uid
	if money_type == 1 {
		log.FundNow = accountfund.AvaFunds
	} else {
		log.FundNow = accountfund.Points
	}

	return global.SHOP_DB.Save(&log).Error, log.Id
}

func (this *AccountFundsLog) CreatelogByTx(db *gorm.DB, username string, amount decimal.Decimal, types int, remark string) error {
	var sagentinfo AgentInfo
	agentinfo, err := sagentinfo.GetCode(username)
	if err != nil {
		return err
	}
	var accountfund model.AccountFunds
	global.SHOP_DB.Where("username=?", username).Find(&accountfund)
	if err != nil {
		return err
	}
	//amounts := amount.String()
	//var saccountfundslog AccountFundsLog
	//saccountfundslog.Createlog(ag_info, username)
	var log model.AccountFundsLog
	log.Username = username
	log.Amount = amount
	log.AgCode = agentinfo.AgCode
	log.Remarks = remark
	log.LevelCode = agentinfo.Level1Code
	now := time.Now()
	log.CreateTime = now
	log.UpdateTime = &now
	log.LogType = types
	log.FundNow = accountfund.AvaFunds
	return db.Save(&log).Error
}
