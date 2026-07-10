package model

// 收货地址表
type ShippingAddresses struct {
	Model
	ReceiverName  string `gorm:"column:receiver_name;NOT NULL" form:"receiver_name" comment:"收货人" json:"receiver_name"`
	ReceiverPhone string `gorm:"column:receiver_phone;NOT NULL" form:"receiver_phone" comment:"收货人电话" json:"receiver_phone"`
	AddressLine1  string `gorm:"column:address_line1;NOT NULL" form:"address_line1" comment:"地址1" json:"address_line1"`
	AddressLine2  string `gorm:"column:address_line2" form:"address_line2" comment:"地址2" json:"address_line_2"`
	Area          string `gorm:"column:area" form:"area" common:"区域/乡镇" json:"area"`
	Username      string `gorm:"column:username" form:"username" comment:"用户名" json:"username"`
	City          string `gorm:"column:city;NOT NULL" form:"city" comment:"城市" json:"city"`
	State         string `gorm:"column:state" form:"state" comment:"州/联邦直辖区" json:"state"`
	PostalCode    string `gorm:"column:postal_code;NOT NULL" form:"postal_code" comment:"邮编" json:"postal_code"`
	CountryCode   string `gorm:"column:country_code;default:MY;NOT NULL" form:"country_code" comment:"国家ISO代码" json:"country_code"`
	Latitude      string `gorm:"-" json:"-"`
	Longitude     string `gorm:"-" json:"-"`
	IsDefault     int    `gorm:"column:is_default;default:0;NOT NULL" form:"is_default" text:"否,是" range:"0,1" comment:"是否默认" json:"is_default" types:"radio"`
	Page          int    `json:"-" form:"page" gorm:"-" `
}

func (*MyStates) ShippingAddresses() string {
	return "shipping_addresses"
}

// 马来西亚州属字典表
type MyStates struct {
	Id      uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键ID
	Code    string    `gorm:"column:code;NOT NULL"`                 // 州属代码
	NameEn  string    `gorm:"column:name_en;NOT NULL"`              // 英文名称
	NameMs  string    `gorm:"column:name_ms"`                       // 马来名称
	NameCn  string    `gorm:"column:name_cn"`                       // 中文名称
	Type    string    `gorm:"column:type;default:STATE"`            // 类型
	MyAreas []MyAreas `gorm:"-"`
}

func (*MyStates) TableName() string {
	return "my_states"
}

type MyAreas struct {
	Id          uint          `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键ID
	StateCode   string        `gorm:"column:state_code;NOT NULL"`           // 关联州属代码
	Code        string        `gorm:"column:code;NOT NULL"`                 // 地区内部编码
	NameEn      string        `gorm:"column:name_en;NOT NULL"`              // 英文名称
	NameMs      string        `gorm:"column:name_ms"`                       // 马来名称
	NameCn      string        `gorm:"column:name_cn"`                       // 中文名称
	MyPostcodes []MyPostcodes `gorm:"-"`
}

func (*MyAreas) TableName() string {
	return "my_areas"
}

// 马来西亚邮政编码字典表
type MyPostcodes struct {
	Id        uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键ID
	AreaCode  string `gorm:"column:area_code;NOT NULL"`            // 关联地区编码
	Postcode  string `gorm:"column:postcode;NOT NULL"`             // 邮政编码 (5位数字)
	PlaceName string `gorm:"column:place_name"`                    // 邮局地点名称
}

func (*MyPostcodes) TableName() string {
	return "my_postcodes"
}
