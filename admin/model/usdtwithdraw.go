package model

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type UsdtWithdrawModel struct {
	Id            uint      `json:"id" form:"id" `
	CreateTime    time.Time `json:"create_time" `
	UpdateTime    time.Time `json:"-" `
	Remarks       string    `json:"-" `
	Username      string
	WalletPath    string          ` form:"wallet_path" json:"wallet_path"`
	Amount        decimal.Decimal ` form:"amount" `
	Status        int
	PathType      string          `form:"path_type" json:"path_type"`
	Type          int             `comment:"1虚拟货币2银行卡" form:"type"`
	Msg           string          `json:"msg" `
	Hash          string          `json:"hash"`
	Fee           decimal.Decimal `json:"fee"`
	TradePassword string          ` form:"trade_password"  gorm:"-" json:"trade_password"`
	ExchangeRate  decimal.Decimal
	OrderNo       string `json:"order_no" gorm:"-" `

	//ModelTime
}

func (*UsdtWithdrawModel) TableName() string {
	return "account_user_withdraw"
}
func (e UsdtWithdrawModel) MarshalJSON() ([]byte, error) {
	type Alias UsdtWithdrawModel // 使用别名来避免递归调用 MarshalJSON
	return json.Marshal(struct {
		*Alias
		CreateTime string `json:"create_time"` // 重写 CreatedAt 的 JSON 标签，使其为字符串类型
	}{
		Alias:      (*Alias)(&e),                               // 将 Event 的字段传递给 Alias，除了 CreatedAt 被重写为字符串格式的时间
		CreateTime: e.CreateTime.Format("2006-01-02 15:04:05"), // 格式化时间并设置为字符串类型
	})
}
