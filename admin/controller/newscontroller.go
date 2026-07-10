package controller

import (
	"net/http"
	"time"

	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/demdxx/gocast"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	BaseController
}

func (f *NewsController) List(ctx *gin.Context) {
	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Username string `form:"username"`
		Count    int64  `json:"count" form:"count"`
		Category string `json:"category" form:"category"`
	}

	req := Request{
		PageInfo: request.PageInfo{
			Limit: 20,
			Page:  1,
		},
	}
	err := ctx.ShouldBind(&req)
	if err != nil {
		f.ErrorHtml(ctx, err.Error())
		return
	}

	req.Offset = (req.Page - 1) * req.Limit

	query := db.Model(&model.News{})

	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		f.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	orderList := make([]model.News, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&orderList).Error; err != nil {
		f.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "news_list.html", gin.H{
		"status":    "200",
		"orderList": orderList,
		"req":       req,
	})
}

func (f *NewsController) Edit(ctx *gin.Context) {
	db := global.SHOP_DB

	if ctx.Request.Method == "GET" {
		type Request struct {
			Id int64 `form:"id"`
		}

		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			f.ErrorHtml(ctx, err.Error())
			return
		}

		news := &model.News{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(news).Error; err != nil {
				f.ErrorHtml(ctx, err.Error())
				return
			}
		}
		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Find(&language)
		version := global.SHOP_CONFIG.System.Version
		ctx.HTML(200, "news_edit.html", gin.H{
			"news":     news,
			"version":  version,
			"language": language,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id       int64  `form:"id"`
			Headline string `form:"headline"`
			Category string `form:"category"`
			Datetime string `form:"datetime"`
			Show     bool   `form:"show"`
			Image    string `form:"image"`
			Summary  string `form:"summary"`
			Content  string `form:"content"`
			Language string `form:"language"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			f.ErrorHtml(ctx, err.Error())
			return
		}
		news := &model.News{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(news).Error; err != nil {
				f.ErrorHtml(ctx, err.Error())
				return
			}

			news.Category = req.Category
			datetime, _ := time.Parse("2006-01-02 15:04:05", req.Datetime)
			news.Datetime = datetime.Unix()
			news.Headline = req.Headline
			news.Show = req.Show
			news.Image = &req.Image
			news.Summary = &req.Summary
			news.Content = &req.Content
			news.Language = req.Language
			if err := db.Save(news).Error; err != nil {
				f.ErrorHtml(ctx, err.Error())
				return
			}

		} else {
			news.Category = req.Category
			news.NewsId = gocast.ToString(time.Now().Unix())
			news.Source = &req.Category
			datetime, _ := time.Parse("2006-01-02 15:04:05", req.Datetime)
			news.Datetime = datetime.Unix()
			news.Headline = req.Headline
			news.Show = req.Show
			news.Image = &req.Image
			news.Summary = &req.Summary
			news.Content = &req.Content
			news.Language = req.Language

			if err := db.Create(news).Error; err != nil {
				f.ErrorHtml(ctx, err.Error())
				return
			}
		}

		f.Success(ctx, "成功")
		return
	}
}
