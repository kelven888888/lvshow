package model

type MTradeAccountApply struct {
	Id         uint       `json:"id" form:"id" `
	CreateTime *LocalTime `gorm:"column:create_time"  json:"create_time" `
	UpdateTime *LocalTime
	Remarks    string
	Username   string
	AccType    string
	Status     *int
	Answer     string
	CarId      string `json:"car_id" form:"car_id" `
	Phone      string
	CarImg     string `json:"car_img" form:"car_img" `
	FirstName  string
	SecondName string
	StreetOne  string
	StreetTwo  string
	Province   string
	City       string
	Name       string `json:"name" form:"name" `
	PostalCode string
	CarImg2    string  `json:"car_img2" form:"car_img2" `
	Fee        float64 `json:"fee" form:"fee" `
	Action     string  `json:"action" form:"action" gorm:"-" `
	AreaCode   string  `json:"area_code" form:"area_code"`

	//ModelTime
}

func (*MTradeAccountApply) TableName() string {
	return "account_account_apply"
}
