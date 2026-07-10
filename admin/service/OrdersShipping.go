package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SOrdersShipping struct {
}

func (this *SOrdersShipping) GetAll(pageInfo request.PageInfo) ([]model.OrdersShipping, int64) {
	var models []model.OrdersShipping

	query := global.SHOP_DB.Model(model.OrdersShipping{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.OrderId != 0 {

		query.Where("id=?", pageInfo.OrderId)
	}
	if pageInfo.UserId != 0 {

		query.Where("user_id=?", pageInfo.UserId)
	}
	if pageInfo.Status != nil {

		query.Where("status=?", pageInfo.Status)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SOrdersShipping) GetByID(id request.GetById) (*model.OrdersShipping, error) {
	var models *model.OrdersShipping
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SOrdersShipping) Save(models *model.OrdersShipping) error {
	if models.Id > 0 {
		now := model.LocalTime(time.Now())
		models.Model.UpdatedAt = &now
		if *models.Status == 2 {
			models.ShippingTime = &now
		}
		if *models.Status == 3 {
			models.FinishTime = &now
		}
		return global.SHOP_DB.Updates(&models).Error
	} else {
		now := model.LocalTime(time.Now())
		models.Model.CreatedAt = &now

		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SOrdersShipping) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.OrdersShipping{}).Error
}
func (this *SOrdersShipping) Deletebatch(req request.IdsReq) error {
	var models []model.OrdersShipping
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
func (this *SOrdersShipping) GetByOrderID(id request.GetById) ([]*model.OrderItemsShipping, error) {
	var models []*model.OrderItemsShipping
	err := global.SHOP_DB.Where("order_id", id.ID).Find(&models).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
