package model

import "github.com/shopspring/decimal"

type Playsetting struct {
	Model
	RateArr         string          `gorm:"column:rate_arr" comment:"概率配置" form:"rate_arr"`                                                    // 概率配置
	RewardArr       string          `gorm:"column:reward_arr" comment:"奖励配置(积分)" form:"reward_arr"`                                            // 奖励配置
	SingelStatus    *int            `gorm:"column:singel_status" types:"radio" text:"关闭,启用" range:"0,1" comment:"单人模式开关" form:"singel_status"` // 0关1开
	DoubleStatus    *int            `gorm:"column:double_status" types:"radio" text:"关闭,启用" range:"0,1" comment:"双人模式开关" form:"double_status"` // 0关1开
	Price           decimal.Decimal `gorm:"column:price" comment:"参加价格" form:"price"`                                                          // 参加价格
	PointArr        string          `gorm:"column:point_arr" form:"point_arr" form:"point_arr"`                                                // 积分配置
	GoodsIds        string          `gorm:"column:goods_ids" form:"goods_ids"`                                                                 // 产品配置
	Type            int             `gorm:"column:type" types:"radio" text:"高爆,保底,魔王,PK,随机" range:"1,2,3,4,5" comment:"类型" form:"type"`        // 1高爆2保底3魔王4pk5随机
	Name            string          `gorm:"column:name" comment:"名称"  form:"name"`
	JoinPointReward string          `gorm:"column:join_point_reward" comment:"参与获得积分奖励"  form:"join_point_reward"`
	Actions         string          `gorm:"-" comment:"actions"  form:"actions"`
	Ids             string          `gorm:"-" comment:"actions"  form:"actions"`
	MoneyToExp      decimal.Decimal `gorm:"column:money_to_exp" comment:"消费1rm得多少经验" form:"money_to_exp"`
	Dhit            *int            `gorm:"column:dhit" comment:"d连击" form:"dhit"`
	Combo           *int            `gorm:"column:combo" comment:"combo连击" form:"combo"`
	CommissionRate  decimal.Decimal `gorm:"column:commission_rate" comment:"佣金比例" form:"commission_rate"`
	SumPlay         int64           `gorm:"-"`
	TotalPlay       int64           `gorm:"-"`
	Img             string          `gorm:"column:img" comment:"图片"  form:"img"`
	Div             int             `gorm:"column:div"`
	SDhit           *int            `gorm:"-" `
	SCombo          *int            `gorm:"-" `
	DDhit           *int            `gorm:"-" `
	DCombo          *int            `gorm:"-" `
}

func (m *Playsetting) TableName() string {
	return "play_config"
}
