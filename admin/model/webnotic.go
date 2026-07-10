package model

type WebNotic struct {
	Model
	Title   string `json:"title" form:"title" comment:"标题"`                                           // 名称
	Content string `json:"content"  comment:"内容" form:"content"`                                      // 内容
	Status  *int   `json:"status"  comment:"状态" form:"status" types:"radio" text:"关闭,启用" range:"0,1"` // 0不显示1显示
	Orders  int    `json:"orders"  comment:"排序" form:"orders"`                                        // 排序
}

func (*WebNotic) TableName() string {
	return "web_notic"
}
