package model

import "time"

type AccountFeedback struct {
	Id         int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	UserId     int64      `json:"user_id" gorm:"column:user_id"`
	Username   string     `json:"username" gorm:"column:username"`
	Name       string     `json:"name" gorm:"column:name"`
	Email      string     `json:"email" gorm:"column:email"`
	ImageUrl   string     `json:"image_url" gorm:"column:image_url"`
	Content    string     `json:"content" gorm:"column:content"`
}

func (AccountFeedback) TableName() string {
	return "account_feedback"
}
