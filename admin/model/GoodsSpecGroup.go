package model

type GoodsSpecGroup struct {
	Model
	Cid  int    `json:"cid" gorm:"column:cid;not null;type:bigint(20)" comment:"类别" form:"cid"`
	Name string `json:"name" gorm:"column:name;not null;type:varchar(50)" comment:"名称" form:"name"`
}

func (*GoodsSpecGroup) TableName() string {
	return "goods_spec_group"
}
