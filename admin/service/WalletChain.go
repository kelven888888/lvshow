package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SWalletChain struct {
}

func (this *SWalletChain) GetAll(pageInfo request.PageInfo) ([]model.WalletChain, int64) {
	var models []model.WalletChain

	query := global.SHOP_DB.Model(model.WalletChain{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s='%s'", pageInfo.SearchField, pageInfo.Keyword))
	}
	//query.Where("pid = ?", 0).Preload("SubCategory")

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" inner_order asc").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SWalletChain) GetByID(id request.GetById) (*model.WalletChain, error) {
	var models *model.WalletChain
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SWalletChain) Save(models *model.WalletChain) error {
	//models.Label = models.Title
	if models.Id > 0 {

		err := global.SHOP_DB.Updates(&models).Error
		this.InnerOrder(0, 1)
		return err

	} else {
		models.CreateTime = time.Now()

		err := global.SHOP_DB.Save(&models).Error
		this.InnerOrder(0, 1)
		return err
	}

}
func (this *SWalletChain) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.WalletChain{}).Error
}
func (this *SWalletChain) Deletebatch(req request.IdsReq) error {
	var models []model.WalletChain
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
func (this *SWalletChain) InnerOrder(pid int, innerorder int) int {
	var models []model.WalletChain
	global.SHOP_DB.Order("id asc").Find(&models, "pid =?", pid)
	//没有下级是更新内部排序
	if len(models) == 0 {
		//var mod model.WalletChain
		//
		//global.SHOP_DB.Model(model.WalletChain{}).Where("id=?", pid).Find(&mod).Updates(model.WalletChain{
		//	InnerOrder: innerorder,
		//})
		//innerorder = innerorder + 1
		return innerorder
	}
	//modellen := len(models)
	for _, v := range models {

		var mod model.WalletChain

		global.SHOP_DB.Model(model.WalletChain{}).Where("id=?", v.Id).Find(&mod).Updates(model.WalletChain{
			InnerOrder: innerorder,
		})
		innerorder = innerorder + 1
		innerorder = this.InnerOrder(v.Id, innerorder)

	}
	fmt.Println("--------------------------", innerorder)

	return innerorder
}
