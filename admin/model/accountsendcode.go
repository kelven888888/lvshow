package model

import "time"

type AccountCheckCode struct {
	Id         int64
	CreateTime time.Time
	UpdateTime time.Time
	Remarks    string
	Name       string `json:"name" form:"name"`
	Captcha    string
	CreateMap  int
	Validity   int
	Type       int8 `comment:"1:注册2修改密码,3忘记密码 4认证 5注销" json:"type" form:"type"`
	Used       int8 `comment:"是否已使用"`
	Status     int8
	Errmsg     string
	Retry      int
	Language   string
}

func (AccountCheckCode) TableName() string {
	return "account_check_code"
}
