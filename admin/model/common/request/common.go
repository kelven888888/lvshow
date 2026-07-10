package request

import (
	"github.com/shopspring/decimal"
	"time"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page           int       `json:"page" form:"page" `        // 页码
	Limit          int       `json:"limit" form:"limit" `      // 每页大小
	Keyword        string    `json:"kw" form:"kw"`             //关键字
	Count          bool      `json:"count" form:"count"`       //关键字
	Offset         int       `json:"offset" form:"offset"`     //关键字
	PageSize       int       `json:"pageSize" form:"pageSize"` // 每页大小
	Account        string    `json:"account" form:"account"`   //
	GroupId        int       `json:"group_id" form:"group_id"`
	Status         *int      `json:"status" form:"status"` //
	Id             int       `json:"id" form:"id" `
	StockTypes     int       `json:"stock_types" form:"stock_types" `
	IsBulk         int       `json:"is_bulk" form:"is_bulk" `
	MStatus        int       `json:"m_status" form:"m_status" `
	OrderTradeType int       `json:"order_trade_type" form:"order_trade_type" `
	OrderType      int       `json:"order_type" form:"order_type" `
	SearchField    string    `json:"search_field" form:"search_field" `
	Type           int       `json:"type" form:"type" `
	IsTest         int       `json:"is_test" form:"is_test" `
	Active         int       `json:"active" form:"active" `
	PlayId         int       `json:"play_id" form:"play_id" `
	OrderId        int       `json:"order_id" form:"order_id" `
	CouponId       int       `json:"coupon_id" form:"coupon_id" `
	RewardType     string    `json:"reward_type" form:"reward_type" `
	Username       string    `json:"username" form:"username" `
	Date           string    `json:"date" form:"date" `
	UserId         int       `json:"user_id" form:"user_id" `
	IsDisplay      int       `json:"is_display" form:"is_display" `
	LogType        int       `json:"log_type" form:"log_type" `
	MoneyType      int       `json:"money_type" form:"money_type" `
	ISDefault      *int      `json:"is_default" form:"is_default"` //
	Expdate        time.Time `json:"exp_date" form:"exp_date"`     //
	EndTime        string    `json:"endtime" form:"endtime" `
}
type PageInfoApi struct {
	Page int `json:"page" form:"page" ` // 页码

	Count    bool `json:"count" form:"count"`       //关键字
	Offset   int  `json:"offset" form:"offset"`     //关键字
	PageSize int  `json:"pageSize" form:"pageSize"` // 每页大小

	Status   int    `json:"status" form:"status"` //
	Type     int    `json:"type" form:"type" `
	IsTest   int    `json:"is_test" form:"is_test" `
	Limit    int    `json:"limit" form:"limit" ` // 每页大小
	Username string `json:"username" form:"username" `
	Date     string `json:"date" form:"date" `
	UserId   int    `json:"user_id" form:"user_id" `
}

// GetById Find by id structure
type GetById struct {
	ID            uint   `form:"id" bind:"required" json:"id"`         // 主键ID
	TradePassword string `form:"trade_password" json:"trade_password"` // 主键ID
}
type GetByUserId struct {
	Id int `form:"id" bind:"required" json:"id"` // 主键ID
}

func (r *GetByUserId) Uint() uint {
	return uint(r.Id)
}
func (r *GetById) Uint() uint {
	return uint(r.ID)
}
func (r *GetById) Uint32() uint32 {

	return uint32(r.ID)
}

type IdsReq struct {
	Ids    []int  `json:"ids[]" form:"ids[]"`
	Action string `json:"action" form:"action"`
	PlayId int    `json:"play_id" form:"play_id"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
type GetByLottyId struct {
	ID          uint   `form:"id" bind:"required" json:"id"`     // 主键ID
	LotteryNum  uint64 `form:"lottery_num" json:"lottery_num"`   // 主键
	LotteryType uint64 `form:"lottery_type" json:"lottery_type"` // 主键
	CouponCode  string `form:"coupon_code" json:"coupon_code"`   // 主键
	CouponId    int    `form:"coupon_id" json:"coupon_id"`       // 主键
}
type IdsReqgood struct {
	Ids            []int           `json:"ids[]" form:"ids[]"`
	Qtys           []int           `json:"qtys[]" form:"qtys[]"`
	Prices         decimal.Decimal `json:"prices" form:"prices"`
	ShippingaddrId int             `json:"shippingaddr_id" form:"shippingaddr_id"`
	Remarks        string          `form:"remarks" json:"remarks"` // 主键
}
