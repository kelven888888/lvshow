package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type AccountFunds struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time

	UpdateTime       time.Time
	Remarks          string
	Username         string
	AvaFunds         decimal.Decimal
	Currency         string
	LockFunds        float64
	VipBorrowFunds   float64 `comment:"vip可借"`
	VipBorrowedFunds float64 `comment:"vip已借"`
	LoadInterest     float64 `comment:"借款利息"`
	LevelCode        string
	AgCode           string
	GiftMoney        float64
	Uid              int
	Points           decimal.Decimal

	//ModelTime
}

func (*AccountFunds) TableName() string {
	return "account_funds"
}
