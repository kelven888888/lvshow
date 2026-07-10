package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SOrderItemsShipping struct {
}

func (this *SOrderItemsShipping) GetAll(pageInfo request.PageInfo) ([]model.OrderItemsShipping, int64) {
	var models []model.OrderItemsShipping

	query := global.SHOP_DB.Model(model.OrderItemsShipping{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SOrderItemsShipping) GetByOrderID(id request.GetById) ([]*model.OrderItemsShipping, error) {
	var models []*model.OrderItemsShipping
	err := global.SHOP_DB.Where("order_id", id.ID).Find(&models).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SOrderItemsShipping) GetByID(id request.GetById) (*model.OrderItemsShipping, error) {
	var models *model.OrderItemsShipping
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SOrderItemsShipping) Save(models *model.OrderItemsShipping) error {
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
func (this *SOrderItemsShipping) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.OrderItemsShipping{}).Error
}
func (this *SOrderItemsShipping) Deletebatch(req request.IdsReq) error {
	var models []model.OrderItemsShipping
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
