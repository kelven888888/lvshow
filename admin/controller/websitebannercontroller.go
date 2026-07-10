package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebsiteBannerController struct {
	BaseController
}

func (h WebsiteBannerController) List(ctx *gin.Context) {
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
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.WebsiteBanner{})

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.WebsiteBanner, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}
	webApiUrl := global.SHOP_CONFIG.System.WebApiURL

	ids := make([]int64, 0)
	for _, row := range list {
		ids = append(ids, row.Id)
	}

	ctx.HTML(http.StatusOK, "website_banner_list.html", gin.H{
		"status":    "200",
		"list":      list,
		"req":       req,
		"webApiUrl": webApiUrl,
	})
}

func (h WebsiteBannerController) Edit(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			h.ErrorHtml(ctx, err.Error())
			return
		}

		news := &model.WebsiteBanner{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(news).Error; err != nil {
				h.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ctx.HTML(200, "website_banner_edit.html", gin.H{
			"banner": news,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id      int64  `form:"id"`
			Type    int    `form:"type"`
			Title   string `form:"title"`
			Content string `form:"content"`
			Image   string `form:"image"`
			Sort    int    `form:"sort"`
			Status  int    `form:"status"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			h.Error(ctx, err.Error())
			return
		}
		banner := &model.WebsiteBanner{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(banner).Error; err != nil {
				h.Error(ctx, err.Error())
				return
			}

			banner.Title = req.Title
			banner.Content = req.Content
			banner.Type = req.Type
			banner.Sort = req.Sort
			banner.Image = req.Image
			banner.Status = req.Status
			if err := db.Save(banner).Error; err != nil {
				h.Error(ctx, err.Error())
				return
			}

		} else {
			banner.Title = req.Title
			banner.Content = req.Content
			banner.Type = req.Type
			banner.Sort = req.Sort
			banner.Image = req.Image
			banner.Status = req.Status

			if err := db.Create(banner).Error; err != nil {
				h.Error(ctx, err.Error())
				return
			}
		}

		h.Success(ctx, "成功", banner)
		return
	}
}

func (h WebsiteBannerController) Delete(ctx *gin.Context) {
	db := global.SHOP_DB
	type Request struct {
		Id int64 `form:"id"`
	}

	req := Request{}
	if err := ctx.ShouldBind(&req); err != nil {
		h.Error(ctx, err.Error())
		return
	}

	if err := db.Where("id = ?", req.Id).Delete(&model.WebsiteBanner{}).Error; err != nil {
		h.Error(ctx, err.Error())
		return
	}

	h.Success(ctx, "成功")
	return
}
