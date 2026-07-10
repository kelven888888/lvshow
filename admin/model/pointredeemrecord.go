package model

import "github.com/shopspring/decimal"

type PointRedeemRecord struct {
	Model
	Remarks     string          `gorm:"column:remarks"`
	ProductId   int             `gorm:"column:product_id" comment:"产品id"`   // 产品id
	Username    string          `gorm:"column:username" comment:"用户名称"`     // 用户名称
	UserId      int             `gorm:"column:user_id" comment:"用户id"`      // 用户id
	Points      decimal.Decimal `gorm:"column:points" comment:"消耗积分"`       // 用户名称
	ProductName string          `gorm:"column:product_name" comment:"产品名称"` // 用户名称
	Image       string          `gorm:"-" comment:"图片"`                     // 用户名称
}

func (m *PointRedeemRecord) TableName() string {
	return "point_redeem_record"
}
