package service

import (
	"fmt"
	"time"

	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
)

type StradeAccountApply struct {
}

func (this *StradeAccountApply) GetAll(pageInfo request.PageInfo) ([]model.MTradeAccountApply, int64) {
	var models []model.MTradeAccountApply

	query := global.SHOP_DB.Model(model.MTradeAccountApply{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != 0 {

			query.Where("status =? ", *pageInfo.Status-1)
		}
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
func (this *StradeAccountApply) Pass(req request.IdsReq, Fee float64, action string) error {
	var models []model.MTradeAccountApply
	tran := global.SHOP_DB.Begin()
	var err error
	if action != "edit" {
		var status int
		status = 1
		err = tran.Find(&models, "id in ?  and status=0", req.Ids).Find(&models).Updates(model.MTradeAccountApply{Status: &status, Fee: Fee}).Error
	} else {
		err = tran.Find(&models, "id in ?", req.Ids).Find(&models).Updates(model.MTradeAccountApply{Fee: Fee}).Error
		return nil

	}
	if err != nil {
		tran.Rollback()
		return err
	}
	//var modeltradeaccount model.TradeAccount
	for _, v := range models {
		username := v.Username

		// 更新用户状态
		{
			user := model.User{}
			err = tran.Where("username = ?", username).First(&user).Error
			if err != nil {
				tran.Rollback()
				return err
			}
			err = tran.Exec("update auth_user set is_auth = 1 where id = ?", user.Id).Error
			if err != nil {
				tran.Rollback()
				return err
			}

			err = tran.Exec("update account_team set status = 1 where user_id = ?", user.Id).Error
			if err != nil {
				tran.Rollback()
				return err
			}
		}

		timenow := time.Now()
		var message AccountMsgServer
		var modelmsg model.AccountUserMessage
		modelmsg.Username = v.Username
		modelmsg.CreateTime = timenow
		modelmsg.UpdateTime = timenow
		group := 0
		keys := utils.Get_Code_Key(10)

		parms := make(map[string]string)

		err = message.Save(&modelmsg, group, group, keys, parms)
		if err != nil {

			global.SHOP_LOG.Log(0, err.Error())
			return err

		}
	}
	tran.Commit()
	return nil
}
func (this *StradeAccountApply) Refust(req request.IdsReq) error {
	var models []model.MTradeAccountApply
	var status int
	status = 2
	err := global.SHOP_DB.Find(&models, "id in ? and status=0", req.Ids).Find(&models).Updates(model.MTradeAccountApply{Status: &status}).Error

	if err != nil {
		return err
	}
	for _, v := range models {
		timenow := time.Now()
		var message AccountMsgServer
		var modelmsg model.AccountUserMessage
		modelmsg.Username = v.Username
		modelmsg.CreateTime = timenow
		modelmsg.UpdateTime = timenow
		group := 0
		keys := utils.Get_Code_Key(11)
		parms := make(map[string]string)

		err = message.Save(&modelmsg, group, group, keys, parms)
		if err != nil {

			global.SHOP_LOG.Log(0, err.Error())
			return err

		}
	}
	return nil
}
