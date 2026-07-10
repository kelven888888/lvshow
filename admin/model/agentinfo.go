package model

import (
	"ginshop.com/global"
	"time"
)

type AgentInfo struct {
	Id          uint      `json:"id" form:"id" `
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Remarks     string
	Acc         string
	AgCode      string
	InCode      string
	Status      int
	AgPassword  string
	LevelCodeId int64 `comment:"管理编号"`
	LevelCode   string
	Level1Code  string
	ParentId    int64 `json:"parent_id" gorm:"column:parent_id"`
	InviteCount int   `json:"invite_count" gorm:"column:invite_count"`
	//ModelTime
}

func (*AgentInfo) TableName() string {
	return "agent_info"
}
func (*AgentInfo) Get_levelcode_from_username(username string) (agcode string, aglevel string, err error) {
	var AgentInfos AgentInfo
	global.SHOP_DB.Where("username=?", username).Find(&AgentInfos)
	return AgentInfos.AgCode, AgentInfos.Level1Code, nil

}
