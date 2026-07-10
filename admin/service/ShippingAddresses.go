package service

import (
	"errors"
    	"fmt"
    	"ginshop.com/admin/model"
    	"ginshop.com/admin/model/common/request"
    	"ginshop.com/global"
    	"time"
)

type SShippingAddresses struct {
}

func (this *SShippingAddresses) GetAll(pageInfo request.PageInfo) ([]model.ShippingAddresses, int64) {
	var models []model.ShippingAddresses

    	query := global.SHOP_DB.Model(model.ShippingAddresses{})

    	if pageInfo.Keyword != "" {

        		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
       }
       if pageInfo.Username != "" {

       		query.Where("username=?", pageInfo.Username)
       	}
       	if pageInfo.Status != nil {
       		if *pageInfo.Status != 0 {

       			query.Where("status =? ", *pageInfo.Status-1)
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
func (this *SShippingAddresses) GetByID(id request.GetById) (*model.ShippingAddresses, error) {
	var models *model.ShippingAddresses
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SShippingAddresses) Save(models *model.ShippingAddresses) error {
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
func (this *SShippingAddresses) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.ShippingAddresses{}).Error
}
func (this *SShippingAddresses) Deletebatch(req request.IdsReq) error {
	var models []model.ShippingAddresses
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
