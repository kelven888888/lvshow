package model

import (
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/shopspring/decimal"
)

type GameLottyRecord struct {
	Model
	UserName    string          `gorm:"column:user_name" comment:"用户名称"`  // 用户名称
	GoodsId     int             `gorm:"column:goods_id" comment:"产品id"`   // 产品
	GoodsType   string          `gorm:"column:goods_type" comment:"产品类型"` // 产品类型
	GoodsName   string          `gorm:"column:goods_name" comment:"产品名称"` // 产品类型
	UserId      int             `gorm:"column:user_id" comment:"用户id"`    // 用户id
	PlayId      int             `gorm:"column:play_id" comment:"玩法id"`    // 玩法id
	PlayName    string          `gorm:"column:play_name" comment:"玩法名称"`  // 玩法名称
	Remark      string          `gorm:"column:remark" comment:"备注"`       // 玩法名称
	RewardType  string          `form:"reward_type"  json:"reward_type" gorm:"size:10;not null;default:0;comment:'商品奖励类型'" comment:"商品奖励类型"`
	Image       string          `gorm:"-" comment:"图片"`                             // 玩法名称
	CouponCode  string          `gorm:"column:coupon_code" comment:"优惠券"`           // 玩法名称
	Discount    decimal.Decimal `gorm:"column:discount" comment:"优惠"`               // 玩法名称
	Money       decimal.Decimal `gorm:"column:money" comment:"花费"`                  // 玩法名称
	LotteryType uint64          `gorm:"column:lottery_type" comment:"lottery_type"` // 玩法id
	JoinMember  uint64          `gorm:"-" comment:"lottery_type"`                   // 玩法id

}

func (m *GameLottyRecord) TableName() string {
	return "game_lotty_record"
}
func (m GameLottyRecord) Coupontitle(language string) string {
	var coupon TbCouponList
	if m.CouponCode == "" {
		return ""
	}
	global.SHOP_DB.Where("coupon_code=?", m.CouponCode).First(&coupon)
	if coupon.Id == 0 {
		return ""
	}
	return utils.Languagebycode(language, coupon.CouponTitle)
}
