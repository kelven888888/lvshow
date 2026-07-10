package model

import "github.com/shopspring/decimal"

type CommissionRecord struct {
	Model

	Remarks      string          `json:"remarks"`       //
	Username     string          `json:"username"`      // 用户名称
	UserId       int             `json:"user_id"`       // 用户id
	Amount       decimal.Decimal `json:"amount"`        // 获得佣金
	FromUserId   int             `json:"from_user_id"`  // 来源用户id
	FromAmount   decimal.Decimal `json:"from_amount"`   // 产生基数
	FromUsername string          `json:"from_username"` // 用户名称
	LogType      int             `json:"log_type"`      // 来源用户id

	LogTypes    string `gorm:"-" json:"log_types"`    // 来源用户id
	MemberLevel int    `gorm:"-" json:"member_level"` // 来源用户id
	Avatar      string `gorm:"-" json:"avatar"`       // 来源用户id
	Page        int    `gorm:"-" json:"page"`         // 来源用户id
}

func (m *CommissionRecord) TableName() string {
	return "commission_record"
}
