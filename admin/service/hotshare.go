package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SHotShare struct {
}

func (this *SHotShare) GetAll(pageInfo request.PageInfo) ([]model.HotShare, int64) {
	var models []model.HotShare

	query := global.SHOP_DB.Model(model.HotShare{})
	if pageInfo.Keyword != "" {

		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if *pageInfo.Status != 0 {

		query.Where("status =? ", pageInfo.Status)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}

	return models, count

}
func (this *SHotShare) GetByID(id request.GetById) (*model.HotShare, error) {
	var models *model.HotShare
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SHotShare) Save(models *model.HotShare) error {
	if models.Id > 0 {
		return global.SHOP_DB.Updates(&models).Error
	} else {
		models.CreateTime = time.Now()
		models.Market = "US"
		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SHotShare) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.HotShare{}).Error
}
