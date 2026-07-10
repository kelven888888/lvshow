package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
)

type VersionController struct {
	BaseController
}

func (v *VersionController) List(ctx *gin.Context) {
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
		v.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.Version{})

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		v.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.Version, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		v.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(200, "version_list.html", gin.H{
		"status": "200",
		"list":   list,
		"req":    req,
	})
	return
}

func (v *VersionController) Edit(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			v.ErrorHtml(ctx, err.Error())
			return
		}

		version := &model.Version{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(version).Error; err != nil {
				v.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ctx.HTML(200, "version_edit.html", gin.H{
			"version": version,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id      int64  `form:"id"`
			Version string `form:"version"`
			Type    string `form:"type"`
			Enable  bool   `form:"enable"`
			Force   bool   `form:"force"`
			Url     string `form:"url"`
			Content string `form:"content"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			v.ErrorHtml(ctx, err.Error())
			return
		}
		version := &model.Version{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(version).Error; err != nil {
				v.ErrorHtml(ctx, err.Error())
				return
			}

			version.Version = req.Version
			version.Type = req.Type
			version.Content = req.Content
			version.Url = req.Url
			version.Enable = req.Enable
			version.Force = req.Force
			if err := db.Save(version).Error; err != nil {
				v.ErrorHtml(ctx, err.Error())
				return
			}

		} else {
			version.Version = req.Version
			version.Type = req.Type
			version.Content = req.Content
			version.Url = req.Url
			version.Enable = req.Enable
			version.Force = req.Force

			if err := db.Create(version).Error; err != nil {
				v.ErrorHtml(ctx, err.Error())
				return
			}
		}

		v.Success(ctx, "成功")
		return
	}
}
