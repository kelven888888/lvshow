package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
)

type Banner struct {
}

func (this *Banner) GetAll(pageInfo request.PageInfo) ([]model.Banners, int64) {
	var models []model.Banners

	query := global.SHOP_DB.Model(model.Banners{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s=?", pageInfo.SearchField), pageInfo.Keyword)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *Banner) GetByID(id request.GetById) (*model.Banners, error) {
	var models *model.Banners
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *Banner) Save(models *model.Banners) error {
	if models.Id > 0 {
		return global.SHOP_DB.Updates(&models).Error
	} else {

		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *Banner) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Banners{}).Error
}
func (this *Banner) Deletebatch(req request.IdsReq) error {
	var models []model.Banners
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
