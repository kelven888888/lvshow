package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SGoodsSpecParam struct {
}

func (this *SGoodsSpecParam) GetAll(pageInfo request.PageInfo) ([]model.GoodsSpecParam, int64) {
	var models []model.GoodsSpecParam

	query := global.SHOP_DB.Model(model.GoodsSpecParam{})

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
func (this *SGoodsSpecParam) GetByID(id request.GetById) (*model.GoodsSpecParam, error) {
	var models *model.GoodsSpecParam
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SGoodsSpecParam) Save(models *model.GoodsSpecParam) error {
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
func (this *SGoodsSpecParam) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.GoodsSpecParam{}).Error
}
func (this *SGoodsSpecParam) Deletebatch(req request.IdsReq) error {
	var models []model.GoodsSpecParam
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
