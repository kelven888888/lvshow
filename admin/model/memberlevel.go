package model

import (
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/shopspring/decimal"
	"time"
)

type MemberLevel struct {
	Id               int    `json:"id"  form:"id"`
	Title            string `json:"title"  form:"title"`
	Thumb            string `json:"thumb"  form:"thumb"`
	Level            *int   `json:"level"  form:"level"`
	OrderBy          int    `json:"order_by"  form:"order_by"`
	IsDisplay        *int   `json:"is_display"  form:"is_display"`
	UpdateTime       time.Time
	CreateTime       time.Time
	InviteMemberNum  int             `json:"invite_member_num"  form:"invite_member_num"`
	InvestMoney      int             `json:"invest_money"  form:"invest_money"`
	DayWithdrawLimit int             `json:"day_withdraw_limit"  form:"day_withdraw_limit"`
	DayBlindNumLimit int             `json:"day_blind_num_limit"  form:"day_blind_num_limit"`
	Exp              decimal.Decimal `json:"exp"  form:"exp"`
	Points           decimal.Decimal `json:"points"  form:"points"`
	CouponId         *int            `json:"coupon_id"  form:"coupon_id"`
	CouponNum        *int            `json:"coupon_num"  form:"coupon_num"`
	CouponName       string          ` form:"-"`
}

func (*MemberLevel) TableName() string {
	return "member_level"
}

func (m MemberLevel) Coupontitle(language string) string {
	var coupon TbCoupon
	if *m.CouponId == 0 {
		return ""
	}
	global.SHOP_DB.Where("id=?", m.CouponId).First(&coupon)
	if coupon.Id == 0 {
		return ""
	}
	return utils.Languagebycode(language, coupon.Title)
}

type LevelUpdateLog struct {
	Model
	Username string          `gorm:"column:username" comment:"用户名"` // 当前积分
	Oldlevel int             `gorm:"column:oldlevel" comment:"旧等级"`
	Newlevel int             `gorm:"column:newlevel" comment:"新等级"`
	Exp      decimal.Decimal `gorm:"column:exp" comment:"当前积分"`  // 当前积分
	Remark   string          `gorm:"column:remark" comment:"备注"` // 当前积分
	Status   int             `gorm:"column:status" comment:"状态"`
}

func (m *LevelUpdateLog) TableName() string {
	return "level_update_log"
}

type ExpRecord struct {
	Model
	Username string          `gorm:"column:username" comment:"用户名"`     // 当前积分
	Oldexp   decimal.Decimal `gorm:"column:oldexp" comment:"老经验"`       // 旧经验
	Exp      decimal.Decimal `gorm:"column:exp" comment:"获取经验"`         // 获得经验
	Newexp   decimal.Decimal `gorm:"column:newexp" comment:"新经验"`       // 新经验
	RecordId int             `gorm:"column:record_id" comment:"抽奖记录id"` // 记录id

}

func (m *ExpRecord) TableName() string {
	return "exp_record"
}
