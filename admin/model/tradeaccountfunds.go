package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TradeAccountFunds struct {
	Id          uint `json:"id" form:"id" `
	CreateTime  time.Time
	UpdateTime  time.Time
	Remarks     string
	TdAcc       string
	AvaFunds    decimal.Decimal
	LockFunds   decimal.Decimal
	Currency    string
	Funds       float64
	ProfitFunds float64
	LevelCode   string
	AgCode      string
	Type        *int
	Username    string

	//ModelTime
}

func (*TradeAccountFunds) TableName() string {
	return "trade_account_funds"
}
