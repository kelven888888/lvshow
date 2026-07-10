package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SPointRedeemRecord struct {
}

func (this *SPointRedeemRecord) GetAll(pageInfo request.PageInfo) ([]model.PointRedeemRecord, int64) {
	var models []model.PointRedeemRecord

	query := global.SHOP_DB.Model(model.PointRedeemRecord{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.Username != "" {

		query.Where("username=?", pageInfo.Username)
	}
	if pageInfo.EndTime != "" {
		query = query.Where("created_at>?", fmt.Sprintf("%s 00:00:00", pageInfo.EndTime))
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SPointRedeemRecord) GetByID(id request.GetById) (*model.PointRedeemRecord, error) {
	var models *model.PointRedeemRecord
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SPointRedeemRecord) Save(models *model.PointRedeemRecord) error {
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
func (this *SPointRedeemRecord) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.PointRedeemRecord{}).Error
}
func (this *SPointRedeemRecord) Deletebatch(req request.IdsReq) error {
	var models []model.PointRedeemRecord
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
