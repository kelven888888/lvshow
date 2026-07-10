package model

import (
	"time"
)

type AccountUserAuthority struct {
	Id         int64
	CreateTime time.Time
	UpdateTime time.Time
	Remarks    string
	Username   string
	App        int8
	Pc         int8
	IsVip      int8 `comment:"是否获得25000借款资格"`
	Avatar     string
	Language   string `comment:"语言"`
}

func (*AccountUserAuthority) TableName() string {
	return "account_user_authority"
}
