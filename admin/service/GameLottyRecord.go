package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SGameLottyRecord struct {
}

func (this *SGameLottyRecord) GetAll(pageInfo request.PageInfo) ([]model.GameLottyRecord, int64) {
	var models []model.GameLottyRecord

	query := global.SHOP_DB.Model(model.GameLottyRecord{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.PlayId != 0 {

		query.Where("play_id=?", pageInfo.PlayId)
	}
	if pageInfo.RewardType != "" {

		query.Where("reward_type=?", pageInfo.RewardType)
	}
	if pageInfo.Username != "" {

		query.Where("user_name=?", pageInfo.Username)
	}
	if pageInfo.Date != "" {

		query.Where("created_at>? and created_at<?", pageInfo.Date, fmt.Sprintf("%s 23:59:59", pageInfo.Date))
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SGameLottyRecord) GetByID(id request.GetById) (*model.GameLottyRecord, error) {
	var models *model.GameLottyRecord
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SGameLottyRecord) Save(models *model.GameLottyRecord) error {
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
func (this *SGameLottyRecord) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.GameLottyRecord{}).Error
}
func (this *SGameLottyRecord) Deletebatch(req request.IdsReq) error {
	var models []model.GameLottyRecord
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
