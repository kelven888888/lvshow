package service

import (
	"errors"
    	"fmt"
    	"ginshop.com/admin/model"
    	"ginshop.com/admin/model/common/request"
    	"ginshop.com/global"
    	"time"
)

type SLevelUpdateLog struct {
}

func (this *SLevelUpdateLog) GetAll(pageInfo request.PageInfo) ([]model.LevelUpdateLog, int64) {
	var models []model.LevelUpdateLog

    	query := global.SHOP_DB.Model(model.LevelUpdateLog{})

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
func (this *SLevelUpdateLog) GetByID(id request.GetById) (*model.LevelUpdateLog, error) {
	var models *model.LevelUpdateLog
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SLevelUpdateLog) Save(models *model.LevelUpdateLog) error {
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
func (this *SLevelUpdateLog) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.LevelUpdateLog{}).Error
}
func (this *SLevelUpdateLog) Deletebatch(req request.IdsReq) error {
	var models []model.LevelUpdateLog
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
