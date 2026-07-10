package model

import (
	"time"
)

type HotShare struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time
	UpdateTime *time.Time
	Remarks    string
	Security   string `json:"security" form:"security"`
	Market     string
	Sort       int
	StockTypes int `json:"stock_types" form:"stock_types"`

	//ModelTime
}

func (*HotShare) TableName() string {
	return "market_hot_share"
}
