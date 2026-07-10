package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SUserStore struct {
}

func (this *SUserStore) GetAll(pageInfo request.PageInfo) ([]model.UserStore, int64) {
	var models []model.UserStore

	query := global.SHOP_DB.Model(model.UserStore{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.Username != "" {

		query.Where("username=?", pageInfo.Username)
	}
	query.Where("qty>0")

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SUserStore) GetByID(id request.GetById) (*model.UserStore, error) {
	var models *model.UserStore
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SUserStore) Save(models *model.UserStore) error {
	if models.Model.Id > 0 {
		now := model.LocalTime(time.Now())
		models.Model.UpdatedAt = &now
		models.Qty = models.Qty + 1
		return global.SHOP_DB.Updates(&models).Error
	} else {
		now := model.LocalTime(time.Now())
		models.Model.CreatedAt = &now

		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SUserStore) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.UserStore{}).Error
}
func (this *SUserStore) Deletebatch(req request.IdsReq) error {
	var models []model.UserStore
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
