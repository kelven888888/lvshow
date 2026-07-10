package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
)

type AgentUser struct {
}

func (this *AgentUser) GetAll(pageInfo request.PageInfo) ([]model.AgentUser, int64) {
	var models []model.AgentUser

	query := global.SHOP_DB.Model(model.AgentUser{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ? or  wallet_path like ?", "%"+pageInfo.Keyword+"%", "%"+pageInfo.Keyword+"%")
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

func (this *AgentUser) GetCode(username string) (agents model.AgentInfo, err error) {
	//var agenrel model.AgentRelatedAccount
	//err = global.SHOP_DB.Table("agent_related_account").Where("username=?", username).Find(&agenrel).Error
	//if err != nil {
	//	return agents, err
	//}
	//agcode := agenrel.AgCode
	var agent model.AgentInfo
	err = global.SHOP_DB.Table("agent_info").Where("agent_user.username=?", username).Select("agent_info.*,agent_user.username ").Joins("left join agent_user on agent_info.ag_code=agent_user.ag_code ").Find(&agent).Error
	if err != nil {
		fmt.Println(err.Error())
		return agent, err
	}
	return agent, nil

}
