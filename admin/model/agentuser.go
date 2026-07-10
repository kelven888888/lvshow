package model

import (
	"time"
)

type AgentUser struct {
	Id         uint `json:"id" form:"id" `
	CreateTime time.Time

	UpdateTime time.Time
	Remarks    string
	Username   string
	AgCode     string
	//ModelTime
}

func (*AgentUser) TableName() string {
	return "agent_user"
}
