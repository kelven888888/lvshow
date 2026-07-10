package model

import "time"

type TradeAccount struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time
	UpdateTime time.Time
	Remarks    string
	Username   string
	TdAcc      string
	TcType     int
	TdName     string
	TdToken    string
	Status     int8
	Profit     float64

	//ModelTime
}

func (TradeAccount) TableName() string {
	return "trade_account"
}
