package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentController struct {
	BaseController
}

func (ac *AgentController) List(ctx *gin.Context) {
	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Count int64 `json:"count" form:"count"`
	}

	req := Request{
		PageInfo: request.PageInfo{
			Limit: 20,
			Page:  1,
		},
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ac.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.AgentInfo{})

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		ac.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.AgentInfo, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		ac.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "agent_list.html", gin.H{
		"status": "200",
		"list":   list,
		"req":    req,
	})
}

func (ac *AgentController) Edit(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ac.ErrorHtml(ctx, err.Error())
			return
		}

		agentInfo := &model.AgentInfo{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(agentInfo).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ctx.HTML(200, "agent_edit.html", gin.H{
			"agent": agentInfo,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id         int64  `form:"id"`
			Acc        string `form:"acc"`
			Status     int    `form:"status"`
			AgPassword string `form:"ag_password"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ac.ErrorHtml(ctx, err.Error())
			return
		}
		agentInfo := &model.AgentInfo{}
		if req.Id != 0 {
			if err := db.Model(agentInfo).Where("id = ?", req.Id).Updates(map[string]any{
				"status": req.Status,
				"acc":    req.Acc,
			}).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}

		} else {
			agentInfo.Acc = req.Acc
			agentInfo.AgPassword = req.AgPassword
			agentInfo.Status = req.Status

			if err := db.Create(agentInfo).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ac.Success(ctx, "成功")
		return
	}
}
func (ac *AgentController) EditPassword(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ac.ErrorHtml(ctx, err.Error())
			return
		}

		agentInfo := &model.AgentInfo{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(agentInfo).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ctx.HTML(200, "agent_edit_password.html", gin.H{
			"agent": agentInfo,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id         int64  `form:"id"`
			AgPassword string `form:"ag_password"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ac.ErrorHtml(ctx, err.Error())
			return
		}
		agentInfo := &model.AgentInfo{}
		if req.Id != 0 {
			if err := db.Model(agentInfo).Where("id = ?", req.Id).Updates(map[string]any{
				"ag_password": req.AgPassword,
			}).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}
		ac.Success(ctx, "成功")
		return
	}
}
