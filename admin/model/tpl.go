package model

import (
	"gorm.io/gorm"
	"time"
)

type Tplm struct {
	Id        int `json:"id"  form:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
	Status    int            `json:"status" gorm:"status"  form:"status"`
	Keys      string         `comment:"键" json:"keys" gorm:"keys"  form:"keys"`
	Content   string         `json:"content" gorm:"content" form:"content"`
	Title     string         `json:"title" gorm:"title" form:"title"`
}

func (Tplm) TableName() string {
	return "web_tpl"
}
