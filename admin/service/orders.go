package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SOrders struct {
}

func (this *SOrders) GetAll(pageInfo request.PageInfo) ([]model.Orders, int64) {
	var models []model.Orders

	query := global.SHOP_DB.Model(model.Orders{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.OrderId != 0 {

		query.Where("id=?", pageInfo.OrderId)
	}
	if pageInfo.UserId != 0 {
		if pageInfo.Type == 1 {
			query.Where("(user_id_sell=? )", pageInfo.UserId)
		} else if pageInfo.Type == 2 {
			query.Where("(user_id_buy=? )", pageInfo.UserId)
		} else {
			query.Where("(user_id_sell=? or user_id_buy=?)", pageInfo.UserId, pageInfo.UserId)
		}

	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != -1 && *pageInfo.Status != 10 {

			query.Where("(status=?)", pageInfo.Status)
		}
		if *pageInfo.Status == 10 {

			query.Where("(status in(1,2,3))")
		}
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SOrders) GetByID(id request.GetById) (*model.Orders, error) {
	var models *model.Orders
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SOrders) Save(models *model.Orders) error {
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
func (this *SOrders) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Orders{}).Error
}
func (this *SOrders) Deletebatch(req request.IdsReq) error {
	var models []model.Orders
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
