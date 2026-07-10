package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/shopspring/decimal"
	"time"
)

type TradeAccountFundsLog struct {
}

func (this *TradeAccountFundsLog) GetAll(pageInfo request.PageInfo) ([]model.TradeAccountFundslog, int64) {
	var models []model.TradeAccountFundslog

	query := global.SHOP_DB.Model(model.TradeAccountFundslog{})
	if pageInfo.Keyword != "" {

		query.Where("td_acc LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if *pageInfo.Status != 0 {

		query.Where("status =? ", pageInfo.Status)
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
func (this *TradeAccountFundsLog) CreateLog(td_acc string, amount decimal.Decimal, types string, remarks string) error {
	var tradeaccountfund model.TradeAccountFunds
	global.SHOP_DB.Where("td_acc=?", td_acc).Find(&tradeaccountfund)

	var log model.TradeAccountFundslog
	log.TdAcc = td_acc
	log.Fee = amount
	log.FType = types
	//amounts := amount.String()
	log.Remarks = remarks
	log.UpdateTime = time.Now()

	log.CreateTime = time.Now()
	log.FundNow = utils.DecimalToFloat(tradeaccountfund.AvaFunds)
	return global.SHOP_DB.Save(&log).Error
}
