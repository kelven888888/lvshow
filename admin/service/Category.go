package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"time"
)

type SCategory struct {
}

func (this *SCategory) GetAll(pageInfo request.PageInfo) ([]model.Category, int64) {
	var models []model.Category
	pageInfo.Limit = 100
	query := global.SHOP_DB.Model(model.Category{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s='%s'", pageInfo.SearchField, pageInfo.Keyword))
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" inner_order asc").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SCategory) GetByID(id request.GetById) (*model.Category, error) {
	var models *model.Category
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SCategory) Save(models *model.Category) error {

	var mode model.Category
	global.SHOP_DB.Where("id=?", models.Pid).Find(&mode)
	if mode.Id == 0 {
		models.Level = 1
	} else {
		models.Level = mode.Level + 1
	}
	if models.Level > 3 {
		return errors.New("只支持三级")
	}

	if models.Id > 0 {

		err := global.SHOP_DB.Updates(&models).Error
		this.InnerOrder(0, 0)
		return err
	} else {
		now := model.LocalTime(time.Now())
		models.Model.CreatedAt = &now

		err := global.SHOP_DB.Save(&models).Error
		this.InnerOrder(0, 0)
		return err

	}

}
func (this *SCategory) Delete(id uint32) error {

	var models []model.Category
	global.SHOP_DB.Where("pid=?", id).Find(&models)
	if len(models) > 0 {
		return errors.New("有子类别,请先删除子类别")
	}
	var GOODS []model.Goods
	global.SHOP_DB.Where("goods_cate=?", id).Find(&GOODS)
	if len(GOODS) > 0 {
		return errors.New("有商品属于类别,请先删除商品")
	}

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Category{}).Error
}
func (this *SCategory) Deletebatch(req request.IdsReq) error {
	var models []model.Category
	global.SHOP_DB.Where("pid in ?", req.Ids).Find(&models)
	if len(models) > 0 {
		return errors.New("有子类别,请先删除子类别")
	}
	var GOODS []model.Goods
	global.SHOP_DB.Where("goods_cate in ?", req.Ids).Find(&GOODS)
	if len(GOODS) > 0 {
		return errors.New("有商品属于类别,请先删除商品")
	}
	var modelss []model.Category
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&modelss).Error

	if err != nil {
		return err
	}

	return nil
}
func (this *SCategory) InnerOrder(pid int, innerorder int) int {
	var models []model.Category
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

		var mod model.Category

		global.SHOP_DB.Model(model.Category{}).Where("id=?", v.Id).Find(&mod).Updates(model.Category{
			InnerOrder: innerorder,
		})
		innerorder = innerorder + 1
		innerorder = this.InnerOrder(v.Model.Id, innerorder)

	}
	fmt.Println("--------------------------", innerorder)

	return innerorder
}

// 获取权限分类
func (this *SCategory) GetCategory(isTree ...bool) []*model.Category {
	var Category []*model.Category
	global.SHOP_DB.Find(&Category)

	if len(isTree) > 0 {
		Category = this.getTree(Category)
	}

	return Category
}

// 生成树分类
func (this *SCategory) getTree(menus []*model.Category) []*model.Category {
	tmpMap := make(map[int]*model.Category)
	var tree []*model.Category
	for _, menu := range menus {
		tmpMap[menu.Id] = menu
	}

	for _, menu := range menus {
		if _, ok := tmpMap[menu.Pid]; ok {
			tmpMap[menu.Pid].Child = append(tmpMap[menu.Pid].Child, tmpMap[menu.Id])
		} else {
			tree = append(tree, tmpMap[menu.Id])
		}
	}
	return tree
}
