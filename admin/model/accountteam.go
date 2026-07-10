package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type AccountTeam struct {
	Id            int64           `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime    *time.Time      `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime    *time.Time      `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks       string          `json:"remarks" gorm:"column:remarks;type:text;"`
	UserId        int64           `json:"user_id" gorm:"column:user_id"`
	ParentId      int64           `json:"parent_id" gorm:"column:parent_id"`
	ParentIds     []int64         `json:"parent_ids" gorm:"column:parent_ids;type:json;serializer:json;"`
	AgentId       int64           `json:"agent_id" gorm:"column:agent_id"`
	AgentInviteId int64           `json:"agent_invite_id" gorm:"column:agent_invite_id"`
	InviteCount   int64           `json:"invite_count" gorm:"column:invite_count"`
	TeamCount     int64           `json:"team_count" gorm:"column:team_count"`
	InviteCode    string          `json:"invite_code" gorm:"column:invite_code"`
	Status        int             `json:"status" gorm:"column:status"`
	OutputAward   decimal.Decimal `json:"output_award" gorm:"column:output_award"`
	InviteAward   decimal.Decimal `json:"invite_award" gorm:"column:invite_award"`
	TeamAward     decimal.Decimal `json:"team_award" gorm:"column:team_award"`
}

func (AccountTeam) TableName() string {
	return "account_team"
}
