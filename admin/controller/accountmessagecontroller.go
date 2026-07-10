package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/demdxx/gocast"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountMessageController struct {
	BaseController
}

func (am *AccountMessageController) List(ctx *gin.Context) {
	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Count    int64  `json:"count" form:"count"`
		UserName string `json:"username" form:"username"`
		Group    string `json:"group" form:"group"`
		Type     string `json:"type" form:"type"`
	}

	req := Request{
		PageInfo: request.PageInfo{
			Limit: 20,
			Page:  1,
		},
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		am.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.AccountMessage{})

	if req.UserName != "" {
		query = query.Where("username LIKE ?", "%"+req.UserName+"%")
	}
	if req.Group != "" {
		query = query.Where("`group` = ?", gocast.ToInt(req.Group))
	}
	if req.Type != "" {
		query = query.Where("`type` = ?", gocast.ToInt(req.Type))
	}

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		am.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.AccountMessage, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		am.ErrorHtml(ctx, err.Error())
		return
	}

	type Row struct {
		model.AccountMessage
		TypeOption  *TypeOption
		GroupOption *GroupOption
	}
	targetList := make([]Row, 0)
	for _, v := range list {
		row := Row{
			AccountMessage: v,
			TypeOption:     nil,
			GroupOption:    nil,
		}
		if row.Type != nil {
			for _, o := range am.TypeOption() {
				if o.Value == *row.Type {
					row.TypeOption = &o
				}
			}
		}
		if row.Group != nil {
			for _, o := range am.GroupOption() {
				if o.Value == *row.Group {
					row.GroupOption = &o
				}
			}
		}

		targetList = append(targetList, row)
	}

	ctx.HTML(http.StatusOK, "accountmessage_list.html", gin.H{
		"status":      "200",
		"list":        targetList,
		"req":         req,
		"typeOption":  am.TypeOption(),
		"groupOption": am.GroupOption(),
	})
}

func (am *AccountMessageController) Edit(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			am.ErrorHtml(ctx, err.Error())
			return
		}

		accountMessage := &model.AccountMessage{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(accountMessage).Error; err != nil {
				am.ErrorHtml(ctx, err.Error)
				return
			}
		}

		type Option struct {
			Name   string `json:"name"`
			Value  int    `json:"value"`
			Choose bool   `json:"choose"`
		}

		typeOption := make([]Option, 0)
		for _, v := range am.TypeOption() {
			op := Option{
				Name:   v.Name,
				Value:  v.Value,
				Choose: false,
			}
			if accountMessage.Type != nil && op.Value == *accountMessage.Type {
				op.Choose = true
			}
			typeOption = append(typeOption, op)
		}
		groupOption := make([]Option, 0)
		for _, v := range am.TypeOption() {
			op := Option{
				Name:   v.Name,
				Value:  v.Value,
				Choose: false,
			}
			if accountMessage.Group != nil && op.Value == *accountMessage.Group {
				op.Choose = true
			}
			groupOption = append(groupOption, op)
		}

		ctx.HTML(200, "accountmessage_edit.html", gin.H{
			"typeOption":     typeOption,
			"groupOption":    groupOption,
			"accountMessage": accountMessage,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id       int64          `form:"id"`
			Type     *int           `json:"type" gorm:"column:type"`
			Group    *int           `json:"group" gorm:"column:group"`
			Title    string         `json:"title" gorm:"column:title"`
			Content  string         `json:"content" gorm:"column:content"`
			Extends  map[string]any `json:"extends" gorm:"column:extends;serializer:json"`
			Username string         `json:"username" gorm:"column:username"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			am.Error(ctx, err.Error())
			return
		}
		fmt.Print(req)
		accountMessage := &model.AccountMessage{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(accountMessage).Error; err != nil {
				am.Error(ctx, err.Error())
				return
			}
			accountMessage.Type = req.Type
			accountMessage.Group = req.Group
			accountMessage.Title = req.Title
			accountMessage.Content = req.Content
			accountMessage.Extends = req.Extends
			accountMessage.Username = req.Username
			if err := db.Save(accountMessage).Error; err != nil {
				am.Error(ctx, err.Error())
				return
			}
		} else {
			accountMessage.Type = req.Type
			accountMessage.Group = req.Group
			accountMessage.Title = req.Title
			accountMessage.Content = req.Content
			accountMessage.Extends = req.Extends
			accountMessage.Username = req.Username
			if err := db.Create(accountMessage).Error; err != nil {
				am.Error(ctx, err.Error())
				return
			}
		}

		am.Success(ctx, "成功")
		return
	}
}

func (am *AccountMessageController) Delete(ctx *gin.Context) {
	db := global.SHOP_DB
	type Request struct {
		Id int64 `form:"id"`
	}

	req := Request{}
	if err := ctx.ShouldBind(&req); err != nil {
		am.Error(ctx, err.Error())
		return
	}

	if err := db.Where("id = ?", req.Id).Delete(&model.AccountMessage{}).Error; err != nil {
		am.Error(ctx, err.Error())
		return
	}

	am.Success(ctx, "成功")
	return
}

type TypeOption struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type GroupOption struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (am *AccountMessageController) TypeOption() []TypeOption {
	return []TypeOption{
		{"系统信息", 0},
		{"活动信息", 1},

		{"慈善捐款开启", 601},
		{"慈善捐款关闭", 602},
		{"慈善捐款成功", 603},
	}
}

func (am *AccountMessageController) GroupOption() []GroupOption {
	return []GroupOption{
		{"系统分组", 0},
		{"活动分组", 1},
		{"交易分组", 2},
		{"订单分组", 3},
		{"量化分组", 4},
		{"基金分组", 5},
		{"慈善分组", 6},
	}
}
