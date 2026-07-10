package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	Id                   int    `json:"id" form:"id" `
	Password             string `json:"password" form:"password" `
	LastLogin            time.Time
	IsSuperuser          int8
	Username             string `json:"username" form:"username" `
	Nickname             string `json:"nickname" form:"nickname" `
	FirstName            string
	LastName             string
	Email                string `json:"email" form:"email"`
	IsStaff              int8
	IsActive             *int `json:"is_active" form:"is_active" gorm:"default:1;"`
	DateJoined           time.Time
	Level                *int   `json:"level" form:"level" gorm:"default:0;"`
	IsTest               int    `json:"is_test" form:"is_test" `
	IsAuth               *int   `json:"is_auth" form:"is_auth" gorm:"default:0;"`
	Language             string `json:"language" form:"language" `
	Phone                string `json:"phone" form:"phone" `
	LoginIp              string `json:"login_ip" form:"login_ip" `
	InviteCode           string `json:"invite_code" form:"invite_code" `
	Captcha              string `json:"captcha" form:"captcha" gorm:"-" `
	ConfirmPassword      string `json:"confirm_password" form:"confirm_password" gorm:"-" `
	AreaCode             string `json:"area_code" form:"area_code" `
	PathId               string `json:"path_id" form:"path_id" `
	Pid                  int    `json:"pid" form:"pid" gorm:"Pid"`
	InviteCount          int    `json:"invite_count" form:"invite_count" gorm:"default:0;"`
	Avatar               string `json:"avatar" form:"avatar" `
	ConfirmTradePassword string `json:"confirm_trade_password" form:"confirm_trade_password" gorm:"-" `
	TradePassword        string `json:"trade_password" form:"trade_password" `
	BlindBoxNum          uint64 `json:"blind_box_num" form:"blind_box_num" gorm:"blind_box_num"`
	Exp                  decimal.Decimal

	//ModelTime
}

func (User) TableName() string {
	return "auth_user"
}

// 绑定手机号struct
type BindPhone struct {
	Phone  string `form:"phone" binding:"required"`
	UserID string `form:"userID" binding:"required"`
}
