package model

import "time"

type Optionslog struct {
	Id           int       `comment:"ID" types:"" text:"" json:"id" form:"id" range:"" edit:"0"`
	CreateTime   time.Time `comment:"创建时间" types:"" text:"" json:"create_time" form:"create_time" range:"" edit:"0"`
	Security     string    `comment:"股票" types:"" text:"" json:"security" form:"security" range:""`
	Contracttype string    `comment:"期权类型" types:"" text:"" json:"contracttype" form:"contracttype" range:""`
	Strikeprice  float64   `comment:"行权价" types:"" text:"" json:"strikeprice" form:"strikeprice" range:""`
	Price        float64   `comment:"当时价格" types:"" text:"" json:"price" form:"price" range:""`
	Remark       string    `comment:"备注" types:"" text:"" json:"remark" form:"remark" range:""`
	TdAcc        string    `comment:"交易账户" types:"" text:"" json:"td_acc" form:"td_acc" range:""`
	Type         *int      `comment:"类型" types:"radio" text:"提前爆仓,行权成功,行权失败,不行权" json:"type" form:"type" range:"1,2,3,4"`
}

func (Optionslog) TableName() string {
	return "options_log"
}
