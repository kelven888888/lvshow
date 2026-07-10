package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
)

type Sservercode struct {
}

func (this *Sservercode) GetAll(pageInfo request.PageInfo) ([]model.AccountCheckCode, int64) {
	var models []model.AccountCheckCode

	query := global.SHOP_DB.Model(model.AccountCheckCode{})
	if pageInfo.Keyword != "" {

		query.Where("name LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != 0 {

			query.Where("status =? ", pageInfo.Status)
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
