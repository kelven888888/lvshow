package service

import (
	"errors"
	"fmt"
	"time"

	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
)

type SbankRecharge struct {
}

func (this *SbankRecharge) GetAll(pageInfo request.PageInfo) ([]model.BankRecharge, int64) {
	var models []model.BankRecharge

	query := global.SHOP_DB.Model(model.BankRecharge{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
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
func (this *SbankRecharge) Save(models *model.BankRecharge) error {
	if models.Id > 0 {
		return global.SHOP_DB.Updates(&models).Error
	} else {
		models.CreateTime = time.Now()
		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SbankRecharge) GetByID(id request.GetById) (*model.BankRecharge, error) {
	var models *model.BankRecharge

	err := global.SHOP_DB.First(&models, id).Error

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
