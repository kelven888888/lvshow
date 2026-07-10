package model

import (
	"github.com/shopspring/decimal"
)

type UserStore struct {
	Model
	UserId     int             `gorm:"column:user_id" comment:"用户id"`
	Username   string          `gorm:"column:username" comment:"用户名"`
	GoodsName  string          `gorm:"column:goods_name" comment:"产品名称"`
	GoodsId    int             `gorm:"column:goods_id" comment:"产品id"`
	Qty        int             `gorm:"column:qty" comment:"数量"`
	Price      decimal.Decimal `gorm:"column:price" comment:"价格"`
	RewardType string          `form:"reward_type"  json:"reward_type" gorm:"size:10;not null;default:0;comment:'商品奖励类型'" comment:"商品奖励类型"`
	GoodsCover string          `gorm:"-"`
}

func (m *UserStore) TableName() string {
	return "user_store"
}
