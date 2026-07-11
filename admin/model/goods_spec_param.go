package model

type GoodsSpecParam struct {
	Model
	Cid       int    `json:"cid" gorm:"column:cid;not null;type:bigint(20)" form:"cid" comment:"分类id"   `
	GroupId   int    `json:"group_id" gorm:"column:group_id;not null;type:bigint(20)" form:"cid" comment:"属性组id"`
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(255)" form:"name" comment:"名称"`
	Numeric   int    `json:"numeric" gorm:"column:numeric;not null;type:tinyint(1)" form:"numeric" comment:"数字类型参数" types:"radio" text:"是,否" range:"1,2" `
	Unit      string `json:"unit" gorm:"column:unit;type:varchar(255)" form:"unit" comment:"单位"`
	Generic   int    `json:"generic" gorm:"column:generic;not null;type:tinyint(1)" form:"generic" comment:"sku通用属性" types:"radio" text:"是,否" range:"1,2"`
	Searching int    `json:"searching" gorm:"column:searching;not null;type:tinyint(1)" form:"searching" comment:"用于搜索过滤" types:"radio" text:"是,否" range:"1,2"`
	Segments  string `json:"segments" gorm:"column:segments;type:varchar(1000)" form:"segments" comment:"数值类型参数"`

	//ModelTime
}

func (*GoodsSpecParam) TableName() string {
	return "goods_spec_param"
}
