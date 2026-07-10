package model

import (
	"github.com/shopspring/decimal"
)

// 订单主表
type Orders struct {
	Model
	UserIdSell      int             `gorm:"column:user_id_sell;NOT NULL" comment:"出售会员ID"`  // 下单用户ID
	UserIdBuy       int             `gorm:"column:user_id_buy;NOT NULL" comment:"购买会员ID"`   // 下单用户ID
	UserNameSell    string          `gorm:"column:user_name_sell;NOT NULL" comment:"出售会员名"` // 下单用户ID
	UserNameBuy     string          `gorm:"column:user_name_buy;NOT NULL" comment:"购买会员名"`
	OrderSn         string          `gorm:"column:order_sn;NOT NULL" comment:"订单号"`                      // 业务订单号(对外展示)
	TotalAmount     decimal.Decimal `gorm:"column:total_amount;default:0.00;NOT NULL" comment:"总价"`      // 订单总金额
	PayAmount       decimal.Decimal `gorm:"column:pay_amount;default:0.00;NOT NULL" comment:"支付金额" `     // 实际支付金额(含优惠)
	Status          *int            `gorm:"column:status;default:0;NOT NULL" comment:"状态" form:"status"` // 订单状态: 0-带成交, 1-已支付, 2-已发货, 3-已完成, 4-已取消
	ReceiverName    string          `gorm:"column:receiver_name"`                                        // 收货人姓名
	ReceiverPhone   string          `gorm:"column:receiver_phone"`                                       // 收货人电话
	ReceiverAddress string          `gorm:"column:receiver_address"`                                     // 收货地址快照
	Remark          string          `gorm:"column:remark" `                                              // 用户备注
	PayTime         *LocalTime      `gorm:"column:pay_time" comment:"成交时间"`                              // 支付时间

	OrderType       int           `gorm:"column:order_type;NOT NULL" comment:"用户名"`  // 1卖2买
	TotalQty        int           `gorm:"column:total_qty;NOT NULL" comment:"订单总数量"` //
	TotalProductQty int           `gorm:"column:total_product_qty;NOT NULL" comment:"订单总数量"`
	Chindren        []*OrderItems `gorm:"-"`
	//UserNameSell    string        `gorm:"-"`
	Avatar      string `gorm:"-"`
	MemberLevel int    `gorm:"-"`
}

func (Orders) TableName() string {
	return "orders"
}

type OrderItems struct {
	Model
	OrderId     int    `gorm:"column:order_id;NOT NULL" comment:"订单id"`     // 关联订单ID
	ProductId   int    `gorm:"column:product_id;NOT NULL" comment:"产品id"`   // 商品ID
	ProductName string `gorm:"column:product_name;NOT NULL" comment:"产品名称"` // 商品名称快照
	ProductImg  string `gorm:"column:product_img" comment:"产品图片"`           // 商品主图快照

	TotalPrice decimal.Decimal `gorm:"column:total_price;NOT NULL" comment:"总价格"`   // 小计金额(单价*数量)
	UnitPrice  decimal.Decimal `gorm:"column:unit_price;NOT NULL" comment:"原价"`     // 单价
	SellPrice  decimal.Decimal `gorm:"column:sell_price;NOT NULL" comment:"sell单价"` // 单价
	Qty        int             `gorm:"column:qty;NOT NULL" comment:"数量"`            // 数量
	RewardType string          `form:"reward_type"  json:"reward_type" gorm:"size:10;not null;default:0;comment:'商品奖励类型'" comment:"商品奖励类型"`
}

func (OrderItems) TableName() string {
	return "order_items"
}

type OrdersShipping struct {
	Model
	UserId int `gorm:"column:user_id;NOT NULL" comment:"会员ID"` // 下单用户ID

	OrderSn         string `gorm:"column:order_sn;NOT NULL" comment:"订单号"` // 业务订单号(对外展示)
	Username        string `gorm:"column:username;NOT NULL" comment:"会员名称"`
	Status          *int   `gorm:"column:status;default:0;NOT NULL" comment:"状态" form:"status"` // 订单状态: 0待发货, 2-已发货, 3-已完成, 4-已取消
	ReceiverName    string `gorm:"column:receiver_name" comment:"收货人姓名"`                        // 收货人姓名
	ReceiverPhone   string `gorm:"column:receiver_phone" comment:"收货人电话"`                       // 收货人电话
	ReceiverAddress string `gorm:"column:receiver_address" comment:"收货地址快照"`                    // 收货地址快照
	Remarks         string `gorm:"column:remarks" form:"remarks"`                               // 用户备注

	TotalQty        int                   `gorm:"column:total_qty;NOT NULL" comment:"订单总数量"` //
	TotalProductQty int                   `gorm:"column:total_product_qty;NOT NULL" comment:"订单总数量"`
	Chindren        []*OrderItemsShipping `gorm:"-"`
	ShippingTime    *LocalTime            `gorm:"column:shipping_time" comment:"发货时间"` // 支付时间
	PostalCode      string                `gorm:"column:postal_code" `
	ShippingFee     decimal.Decimal       `gorm:"column:shipping_fee" comment:"邮费"`  // 支付时间
	FinishTime      *LocalTime            `gorm:"column:finish_time" comment:"完成时间"` // 支付时间
	TotalMoney      decimal.Decimal       `gorm:"-"`                                 // 支付时间
}

func (OrdersShipping) TableName() string {
	return "orders_shipping"
}

type OrderItemsShipping struct {
	Model
	OrderId     int    `gorm:"column:order_id;NOT NULL" comment:"订单id"`     // 关联订单ID
	ProductId   int    `gorm:"column:product_id;NOT NULL" comment:"产品id"`   // 商品ID
	ProductName string `gorm:"column:product_name;NOT NULL" comment:"产品名称"` // 商品名称快照
	ProductImg  string `gorm:"column:product_img" comment:"产品图片"`           // 商品主图快照

	Qty        int    `gorm:"column:qty;NOT NULL" comment:"数量"` // 数量
	RewardType string `form:"reward_type"  json:"reward_type" gorm:"size:10;not null;default:0;comment:'商品奖励类型'" comment:"商品奖励类型"`
}

func (OrderItemsShipping) TableName() string {
	return "order_items_shipping"
}
