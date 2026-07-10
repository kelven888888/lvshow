package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
	"strconv"
	"strings"
	"time"
)

type SPlaysetting struct {
}

func (this *SPlaysetting) GetAll(pageInfo request.PageInfo) ([]model.Playsetting, int64) {
	var models []model.Playsetting

	query := global.SHOP_DB.Model(model.Playsetting{})

	if pageInfo.Keyword != "" {

		query.Where(fmt.Sprintf("%s like '%s'", pageInfo.SearchField, "%%"+pageInfo.Keyword+"%%"))
	}
	if pageInfo.Id != 0 {

		query.Where("id=?", pageInfo.Id)
	}
	if pageInfo.Status != nil {
		if *pageInfo.Status != 0 {

			query.Where("singel_status=? or double_status=?", pageInfo.Status, pageInfo.Status)
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
func (this *SPlaysetting) GetAllProduct(pageInfo request.PageInfo) ([]model.Goods, int64) {
	var models []model.Goods

	//query := global.SHOP_DB.Model(model.Goods{}).Select("a.id as play_goods_id,a.play_id ,b.*").Joins("")
	//

	query := global.SHOP_DB.Model(model.PlayGoods{}).Joins("JOIN goods ON play_goods.goods_id = goods.id").Select("play_goods.id ," +
		"play_goods.play_id,play_goods.created_at ,goods.goods_name,goods.reward_type,goods.goods_status,goods.unit_price,goods.goods_cover,goods.id as play_goods_id")
	if pageInfo.PlayId != 0 {

		query.Where(fmt.Sprintf("%s = '%d'", "play_id", pageInfo.PlayId))
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id DESC").Find(&models).Error
	if err != nil {
		return nil, 0
	}

	return models, count

}
func (this *SPlaysetting) GetByID(id request.GetById) (*model.Playsetting, error) {
	var models *model.Playsetting
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *SPlaysetting) Save(models *model.Playsetting) error {
	rate_arr := strings.Split(models.RateArr, ",")
	reward_arr := strings.Split(models.RewardArr, ",")
	points_arr := strings.Split(models.PointArr, ",")
	sum := 0.0
	div := 100
	for _, v := range rate_arr {
		vals, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		if vals <= 0.001 {
			return errors.New("只支持小数点后两位")
		}
		sum = sum + vals
		fraction := utils.GetDecimalPlaces(vals)

		fmt.Printf("小数位: %d\n", fraction)

		if fraction == 1 && div < 1000 && fraction > 0 {
			div = 1000
		}
		if fraction == 2 && div < 10000 && fraction > 0 {
			div = 10000
		}
	}
	models.Div = div

	//fmt.Println(models.Type, len(rate_arr), len(points_arr), len(reward_arr), "models.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Typemodels.Type")
	//if models.Type == 6 && len(rate_arr) != 1 && len(points_arr) != 1 && len(reward_arr) != 1 {
	//
	//	return errors.New("盲盒子只能1项")
	//
	//}
	if sum != 100 {
		return errors.New(fmt.Sprintf("设置概率不是100-设置概率%.2f", sum))
	}
	if len(reward_arr) != len(rate_arr) {
		return errors.New("奖励配置与概率配置不对应")
	}
	if len(reward_arr) != len(points_arr) {
		return errors.New("奖励配置与积分配置不对应")
	}
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
func (this *SPlaysetting) Delete(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.Playsetting{}).Error
}
func (this *SPlaysetting) DeleteProduct(id uint32) error {

	return global.SHOP_DB.Where("id=?", id).Delete(&model.PlayGoods{}).Error
}

func (this *SPlaysetting) Deletebatch(req request.IdsReq) error {
	var models []model.Playsetting
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}

	return nil
}
func (this *SPlaysetting) AddProduct(req request.IdsReq) error {
	var models model.Playsetting
	global.SHOP_DB.Find(&models).Where("id  =?", req.PlayId)
	if models.Id == 0 {
		return errors.New("玩法记录没有找到")
	}
	var err error
	for _, v := range req.Ids {
		var mo model.PlayGoods
		global.SHOP_DB.Where("goods_id=? and play_id=?", v, req.PlayId).Find(&mo)
		var goods model.Goods
		global.SHOP_DB.Where("id=? ", v).Find(&goods)
		if mo.Id > 0 {
			now := model.LocalTime(time.Now())
			mo.RewardType = goods.RewardType
			mo.Price = goods.UnitPrice
			mo.Model.UpdatedAt = &now
			err = global.SHOP_DB.Updates(&mo).Error
		} else {
			now := model.LocalTime(time.Now())
			mo.GoodsId = v
			mo.PlayId = req.PlayId
			mo.RewardType = goods.RewardType
			mo.Price = goods.UnitPrice
			mo.Model.CreatedAt = &now

			err = global.SHOP_DB.Save(&mo).Error
		}
	}
	if err != nil {
		return err
	}

	return nil
}
