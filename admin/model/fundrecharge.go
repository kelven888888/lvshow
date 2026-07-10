package model

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type FundRecharge struct {
	Id uint `json:"id" form:"id" `

	CreateTime time.Time  `json:"create_time"`
	UpdateTime *time.Time `json:"-" `
	Remarks    string
	Username   string
	Amount     decimal.Decimal `form:"amount"`
	Address    string
	PathType   string `form:"path_type" json:"path_type"`
	//WalletPath string `form:"wallet_path"`
	Type         int `comment:"1虚拟货币2银行卡" form:"type"`
	Cms          float64
	LevelCode    string
	AgCode       string
	Hash         string `comment:"hash"`
	Status       int
	Pic          string `form:"pic"`
	ExchangeRate decimal.Decimal

	//ModelTime
}

func (*FundRecharge) TableName() string {
	return "account_funds_recharge_log"
}
func (e FundRecharge) MarshalJSON() ([]byte, error) {
	type Alias FundRecharge // 使用别名来避免递归调用 MarshalJSON
	return json.Marshal(struct {
		*Alias
		CreateTime string `json:"create_time"` // 重写 CreatedAt 的 JSON 标签，使其为字符串类型
	}{
		Alias:      (*Alias)(&e),                               // 将 Event 的字段传递给 Alias，除了 CreatedAt 被重写为字符串格式的时间
		CreateTime: e.CreateTime.Format("2006-01-02 15:04:05"), // 格式化时间并设置为字符串类型
	})
}
