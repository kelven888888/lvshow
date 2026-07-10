package model

import (
	"github.com/shopspring/decimal"
)

type TbCouponList struct {
	Model
	CouponCode   string          `gorm:"column:coupon_code" comment:"券码" `
	Status       int             `gorm:"column:status" comment:"状态"` // 0未启用1已启用2已激活
	Price        decimal.Decimal `gorm:"column:price" comment:"价值"`
	CouponId     int             `gorm:"column:coupon_id" comment:"优惠券id"` // couponid
	Uid          int             `gorm:"column:uid" comment:"用户id"`        // 用户id
	ActivateTime *LocalTime      `gorm:"column:activate_time" `            // 激活时间
	CouponTitle  string          `gorm:"column:coupon_title" comment:"券名"`
	Username     string          `gorm:"column:username" comment:"用户"`         // 激活用户
	Remarks      string          `gorm:"column:remarks" comment:"备注" json:"-"` // 激活用户

	ExpDate *LocalTime `gorm:"column:exp_date" comment:"失效时间"` // 失效时间

	UsedTime *LocalTime `gorm:"column:used_time" comment:"使用时间"` // 使用时间
}

func (m *TbCouponList) TableName() string {
	return "tb_coupon_list"
}
