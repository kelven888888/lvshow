package model

import "time"

type AccountMessage struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks    string    `json:"remarks" gorm:"column:remarks;default:success"`
	Type       *int      `json:"type" gorm:"column:type"`
	Group      *int      `json:"group" gorm:"column:group"`

	Title    string         `json:"title" gorm:"column:title"`
	Content  string         `json:"content" gorm:"column:content"`
	Username string         `json:"username" gorm:"column:username"`
	Status   int            `json:"status" gorm:"column:status"`
	Read     int            `json:"read" gorm:"column:read"`
	Extends  map[string]any `json:"extends" gorm:"column:extends;serializer:json"`
}

func (AccountMessage) TableName() string {
	return "account_user_message"
}
