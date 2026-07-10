package model

import "time"

type Version struct {
	Model
	Id         int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Type       string     `json:"type" gorm:"type:varchar(30);not null;comment:类型"`
	Version    string     `json:"version" gorm:"type:varchar(30);not null;comment:版本号"`
	Enable     bool       `json:"enable" gorm:"type:tinyint(1);not null;comment:是否启用"`
	Force      bool       `json:"force" gorm:"type:tinyint(1);not null;comment:是否强制更新"`
	Url        string     `json:"url" gorm:"type:varchar(255);not null;comment:下载地址"`
	Content    string     `json:"content" gorm:"type:text;not null;comment:更新内容"`
}

func (Version) TableName() string {
	return "version"
}
