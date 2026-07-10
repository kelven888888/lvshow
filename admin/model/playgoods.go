package model

import "github.com/shopspring/decimal"

type PlayGoods struct {
	Model
	PlayId     int
	GoodsId    int
	RewardType string
	Price      decimal.Decimal

	//ModelTime
}

func (*PlayGoods) TableName() string {
	return "play_goods"
}
