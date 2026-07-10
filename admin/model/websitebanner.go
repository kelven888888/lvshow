package model

import "time"

type WebsiteBanner struct {
	Id         int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks    string     `json:"remarks" gorm:"column:remarks;type:text;"`
	Type       int        `json:"type" gorm:"column:type;"`
	Title      string     `json:"title" gorm:"column:title;type:varchar(255);"`
	Content    string     `json:"content" gorm:"column:content;type:varchar(255)"`
	Image      string     `json:"image" gorm:"column:image;type:varchar(255);"`
	PointUrl   string     `json:"point_url" gorm:"column:point_url;type:varchar(255);"`
	Sort       int        `json:"sort" gorm:"column:sort;"`
	Status     int        `json:"status" gorm:"column:status;"`
}

func (WebsiteBanner) TableName() string {
	return "pc_banner_image"
}
