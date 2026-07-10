package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SWebNotic struct {
}

func (this *SWebNotic) GetAll(pageInfo request.PageInfo) ([]model.WebNotic, int64) {
	var models []model.WebNotic

	query := global.SHOP_DB.Model(model.WebNotic{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.Status != nil {

		query.Where("status=?", pageInfo.Status)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" orders DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SWebNotic) GetByID(id request.GetById) (*model.WebNotic, error) {
	var models *model.WebNotic
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SWebNotic) Save(models *model.WebNotic) error {
	if models.Id > 0 {
		now := model.LocalTime(time.Now())
		models.Model.UpdatedAt = &now
		return global.SHOP_DB.Updates(&models).Error
	} else {
		now := model.LocalTime(time.Now())
		models.Model.CreatedAt = &now

		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SWebNotic) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.WebNotic{}).Error
}
func (this *SWebNotic) Deletebatch(req request.IdsReq) error {
	var models []model.WebNotic
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
