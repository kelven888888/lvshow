package model

import "time"

type Agreement struct {
	Id         int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Key        string     `json:"key" gorm:"column:key;notnull;type:varchar(30)" form:"key"`
	Group      string     `json:"group" gorm:"column:group;notnull;type:varchar(30)"`
	Name       string     `json:"name" gorm:"column:name;notnull;type:varchar(255)"`
	Content    string     `json:"content" gorm:"column:content;notnull;type:text"`
	Language   string     `json:"language" gorm:"column:language;"`
}

func (Agreement) TableName() string {
	return "agreement"
}
