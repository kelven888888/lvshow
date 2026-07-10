package model

import "time"

type AgentInvite struct {
	Id          int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime  *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime  *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks     string     `json:"remarks" gorm:"column:remarks;type:text;"`
	AgentId     int64      `json:"agent_id" gorm:"column:agent_id"`
	InviteCode  string     `json:"invite_code" gorm:"column:invite_code"`
	InviteCount int        `json:"invite_count" gorm:"column:invite_count"`
	TeamCount   int        `json:"team_count" gorm:"column:team_count"`
}

func (AgentInvite) TableName() string {
	return "agent_invite"
}
