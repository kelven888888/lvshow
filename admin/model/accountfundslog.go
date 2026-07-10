package model

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type AccountFundsLog struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time
	UpdateTime *time.Time `json:"-" `
	Remarks    string     `json:"-" `
	Username   string
	Amount     decimal.Decimal `json:"amount" form:"amount"`
	LogType    int             `json:"log_type" form:"log_type" `
	MoneyType  int             `json:"money_type" form:"money_type" `
	AgCode     string
	LevelCode  string
	FundNow    decimal.Decimal
	Page       int    `json:"-" form:"page" gorm:"-" `
	DataType   int    `json:"data_type" form:"data_type" gorm:"-" `
	Uid        int    `json:"uid" form:"uid"  `
	LogTypes   string ` gorm:"-"  `
	//ModelTime
}

func (*AccountFundsLog) TableName() string {
	return "account_funds_log"
}
func (e AccountFundsLog) MarshalJSON() ([]byte, error) {
	type Alias AccountFundsLog // 使用别名来避免递归调用 MarshalJSON
	return json.Marshal(struct {
		*Alias
		CreateTime string `json:"CreateTime"` // 重写 CreatedAt 的 JSON 标签，使其为字符串类型
	}{
		Alias:      (*Alias)(&e),                               // 将 Event 的字段传递给 Alias，除了 CreatedAt 被重写为字符串格式的时间
		CreateTime: e.CreateTime.Format("2006-01-02 15:04:05"), // 格式化时间并设置为字符串类型
	})
}
