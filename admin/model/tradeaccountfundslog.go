package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TradeAccountFundslog struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time
	UpdateTime time.Time
	Remarks    string
	TdAcc      string
	Fee        decimal.Decimal
	FType      string
	LevelCode  string
	AgCode     string
	FundNow    float64

	//ModelTime
}

func (*TradeAccountFundslog) TableName() string {
	return "trade_account_funds_log"
}
