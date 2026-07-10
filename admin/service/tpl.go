package service

import (
	"errors"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"strconv"
)

type Tpls struct {
}

func (this *Tpls) GetAll(pageInfo request.PageInfo) ([]model.Tplm, int64) {
	var models []model.Tplm

	query := global.SHOP_DB
	if pageInfo.Account != "" {
		//query.Where("nov_admin.account=?", pageInfo.Account)
	}
	if strconv.Itoa(pageInfo.GroupId) != "" && pageInfo.GroupId != 0 {
		//query.Where("nov_group.id =?", pageInfo.GroupId)
	}
	var count int64 = 0
	query.Model(models).Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *Tpls) GetByID(id request.GetById) (*model.Tplm, error) {
	var models *model.Tplm
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *Tpls) Save(models *model.Tplm) error {
	if models.Id > 0 {
		return global.SHOP_DB.Updates(&models).Error
	} else {
		//models.CreateTime = time.Now()
		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *Tpls) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Tplm{}).Error
}
func (this *Tpls) Deletebatch(req request.IdsReq) error {
	var models []model.Tplm
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
