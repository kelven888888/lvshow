package model

import "time"

type AccountBankContract struct {
	Id           int64          `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime   *time.Time     `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime   *time.Time     `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks      string         `json:"remarks" gorm:"column:remarks;type:text;"`
	UserId       int64          `json:"user_id" gorm:"column:user_id"`
	Username     string         `json:"username" gorm:"column:username"`
	ContractNo   string         `json:"contract_no" gorm:"column:contract_no"`
	Status       int            `json:"status" gorm:"column:status"`
	SignInfo     map[string]any `json:"sign_info" gorm:"column:sign_info;serializer:json"`
	SignTime     time.Time      `json:"sign_time" gorm:"column:sign_time"`
	SignImgPath  string         `json:"sign_img_path" gorm:"column:sign_img_path"`
	TemplatePath string         `json:"template_path" gorm:"column:template_path"`
	DownloadPath string         `json:"download_path" gorm:"column:download_path"`
	Amount       float64        `json:"amount" gorm:"column:amount"`
	Read         bool           `json:"read" gorm:"column:read"`
	HandlerTime  *time.Time     `json:"handler_time" gorm:"column:handler_time"`
	ServiceFee   float64        `json:"service_fee" gorm:"column:service_fee"`
}

func (AccountBankContract) TableName() string {
	return "bank_transfer_contract"
}
