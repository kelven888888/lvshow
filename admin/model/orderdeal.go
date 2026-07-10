package model

import (
	"time"
)

type OrderDeal struct {
	Id              uint `json:"id" form:"id" `
	CreateTime      time.Time
	UpdateTime      time.Time
	Remarks         string
	TdAcc           string
	ENum            string
	DNum            string
	Price           float64
	Amount          int64
	Security        string
	Cms             float64
	DType           int
	Direction       int
	Profit          float64
	QuanAccountId   int `comment:"购买量化产品id"`
	AiquanAccountId int `comment:""`
	OrderTradeType  int
	AgPrice         float64
	Dire            int
	FundNow         float64
	PosNow          int
	Status          int
	OrderType       int
	Premium         float64
	PremiumRate     float64

	//ModelTime
}

func (OrderDeal) TableName() string {
	return "trade_entrust_deal"
}
