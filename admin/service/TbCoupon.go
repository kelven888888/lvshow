package service

import (
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
	"time"
)

type STbCoupon struct {
}

func (this *STbCoupon) GetAll(pageInfo request.PageInfo) ([]model.TbCoupon, int64) {
	var models []model.TbCoupon

	query := global.SHOP_DB.Model(model.TbCoupon{})

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
func (this *STbCoupon) GetByID(id request.GetById) (*model.TbCoupon, error) {
	var models *model.TbCoupon
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return models, nil
}
func (this *STbCoupon) Coupongen(id request.GetById) error {
	var models *model.TbCoupon
	err := global.SHOP_DB.First(&models, id).Error
	if err != nil {
		return errors.New(err.Error())
	}
	now := model.LocalTime(time.Now())
	len := models.Numbers
	var couponlistarr []model.TbCouponList
	for i := 0; i < len; i++ {

		var couponlist model.TbCouponList
		couponlist.Status = 0

		couponlist.CreatedAt = &now
		couponlist.CouponCode = utils.GenerateRandomString(10)
		couponlist.CouponId = models.Id
		couponlist.CouponTitle = models.Title
		couponlist.Price = models.Price
		extime := model.LocalTime(time.Now().AddDate(0, 0, models.Cycle))
		couponlist.ExpDate = &extime //俩年有效期
		couponlistarr = append(couponlistarr, couponlist)

	}
	err = global.SHOP_DB.CreateInBatches(&couponlistarr, 500).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (this *STbCoupon) Save(models *model.TbCoupon) error {
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
func (this *STbCoupon) Delete(id uint32) error {

	err := global.SHOP_DB.Where("id=?", id).Delete(&model.TbCoupon{}).Error
	if err != nil {
		return err
	}
	err = global.SHOP_DB.Where("coupon_id=?", id).Delete(&model.TbCouponList{}).Error
	if err != nil {
		return err
	}
	return nil

}
func (this *STbCoupon) Deletebatch(req request.IdsReq) error {
	var models []model.TbCoupon
	err := global.SHOP_DB.Find(&models, "id in ?", req.Ids).Delete(&models).Error

	if err != nil {
		return err
	}
	err = global.SHOP_DB.Where("coupon_id in?", req.Ids).Delete(&model.TbCouponList{}).Error
	if err != nil {
		return err
	}

	return nil
}
