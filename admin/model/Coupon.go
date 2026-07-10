package model

import "github.com/shopspring/decimal"

type TbCoupon struct {
	Model
	Title string `gorm:"column:title"  comment:"标题" form:"title"` // 标题

	Price    decimal.Decimal `gorm:"column:price" text:"价值" comment:"价值" form:"price"`                                          // 价值
	Content  string          `gorm:"column:content"`                                                                            // 备注
	Numbers  int             `gorm:"column:numbers" text:"生成个数" comment:"生成个数" form:"numbers"`                                  // 生成个数
	Isdelete int             `gorm:"column:isdelete" `                                                                          // 0正常1删除
	Cycle    int             `gorm:"column:cycle" text:"周期" comment:"周期" form:"cycle"`                                          // 有效期(天)
	Status   int             `gorm:"column:status;default:1" types:"radio" text:"关闭,启用" range:"0,1" comment:"状态" form:"status"` // 是否启用0禁用1启用

}

func (m *TbCoupon) TableName() string {
	return "tb_coupon"
}
