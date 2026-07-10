package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SGoods struct {
}

func (this *SGoods) GetAll(pageInfo request.PageInfo) ([]model.Goods, int64) {
	var models []model.Goods

	query := global.SHOP_DB.Model(model.Goods{})

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
func (this *SGoods) GetByID(id request.GetById) (*model.Goods, error) {
	var models *model.Goods
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SGoods) Save(models *model.Goods) error {

	global.SHOP_DB.Model(model.PlayGoods{}).Where("goods_id=?", models.Id).Updates(model.PlayGoods{
		RewardType: models.RewardType,
		Price:      models.UnitPrice,
	})
	global.SHOP_DB.Model(model.UserStore{}).Where("goods_id=?", models.Id).Updates(model.UserStore{
		RewardType: models.RewardType,
		Price:      models.UnitPrice,
		GoodsName:  models.GoodsName,
	})
	if models.Id > 0 {
		now := model.LocalTime(time.Now())
		models.Model.UpdatedAt = &now
		return global.SHOP_DB.Updates(&models).Error
	} else {
		now := model.LocalTime(time.Now())
		models.Model.CreatedAt = &now
		//models.CreateTime=time.Now()
		return global.SHOP_DB.Save(&models).Error
	}

}
func (this *SGoods) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Goods{}).Error
}
func (this *SGoods) Deletebatch(req request.IdsReq) error {
	var models []model.Goods
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
