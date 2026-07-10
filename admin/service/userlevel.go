package service

import (
	"errors"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"strconv"
	"time"
)

type Userlevel struct {
}

func (this *Userlevel) GetAll(pageInfo request.PageInfo) ([]model.MemberLevel, int64) {
	var models []model.MemberLevel

	query := global.SHOP_DB.Model(model.MemberLevel{})
	if pageInfo.Account != "" {
		//query.Where("nov_admin.account=?", pageInfo.Account)
	}
	if pageInfo.IsDisplay == 1 {
		query.Where("is_display=?", pageInfo.IsDisplay)
	}
	if strconv.Itoa(pageInfo.GroupId) != "" && pageInfo.GroupId != 0 {
		//query.Where("nov_group.id =?", pageInfo.GroupId)
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" level asc").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *Userlevel) GetByID(id request.GetById) (*model.MemberLevel, error) {
	var models *model.MemberLevel
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *Userlevel) Save(models *model.MemberLevel) error {
	if models.Id > 0 {
		models.UpdateTime = time.Now()
		return global.SHOP_DB.Updates(&models).Error
	} else {
		models.CreateTime = time.Now()
		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *Userlevel) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.MemberLevel{}).Error
}
func (this *Userlevel) Deletebatch(req request.IdsReq) error {
	var models []model.MemberLevel
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
