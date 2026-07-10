package controller

import (
	"errors"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AgreementController struct {
	BaseController
}

func (ac *AgreementController) List(ctx *gin.Context) {
	db := global.SHOP_DB

	type Request struct {
		request.PageInfo
		Count int64  `json:"count" form:"count"`
		Key   string `form:"key" form:"key"`
		Name  string `json:"name" form:"name"`
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

	query := db.Model(&model.Agreement{})

	if strings.TrimSpace(req.Key) != "" {
		query = query.Where("`key` like ?", "%"+req.Key+"%")
	}

	if strings.TrimSpace(req.Name) != "" {
		query = query.Where("`name` like ?", "%"+req.Name+"%")
	}

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		ac.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.Agreement, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		ac.ErrorHtml(ctx, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "agreement_list.html", gin.H{
		"status": "200",
		"list":   list,
		"req":    req,
	})
}

func (ac *AgreementController) Edit(ctx *gin.Context) {
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

		agreement := &model.Agreement{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(agreement).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}

		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Find(&language)

		ctx.HTML(200, "agreement_edit.html", gin.H{
			"agreement": agreement,
			"language":  language,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		type Request struct {
			Id       int64  `form:"id"`
			Key      string `form:"key"`
			Name     string `form:"name"`
			Group    string `form:"group"`
			Content  string `form:"content"`
			Language string `form:"language"`
		}
		req := Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ac.Error(ctx, err.Error())
			return
		}
		agreement := &model.Agreement{}
		if req.Id != 0 {
			if err := db.Where("id = ?", req.Id).First(agreement).Error; err != nil {
				ac.Error(ctx, err.Error())
				return
			}

			find, err := ac.FindByKey(db, req.Key, req.Language)
			if err != nil {
				ac.Error(ctx, err.Error())
				return
			}

			if find != nil && find.Id != req.Id {
				ac.Error(ctx, "该key已存在,修改失败")
				return
			}

			agreement.Key = req.Key
			agreement.Name = req.Name
			agreement.Group = req.Group
			agreement.Content = req.Content
			agreement.Language = req.Language
			if err := db.Save(agreement).Error; err != nil {
				ac.Error(ctx, err.Error())
				return
			}

		} else {
			find, err := ac.FindByKey(db, req.Key, req.Language)
			if err != nil {
				ac.Error(ctx, err.Error())
				return
			}

			if find != nil {
				ac.Error(ctx, "该key已存在,添加失败")
				return
			}

			agreement.Key = req.Key
			agreement.Name = req.Name
			agreement.Group = req.Group
			agreement.Content = req.Content
			agreement.Language = req.Language

			if err := db.Create(agreement).Error; err != nil {
				ac.ErrorHtml(ctx, err.Error())
				return
			}
		}

		ac.Success(ctx, "成功")
		return
	}
}

func (ac *AgreementController) Delete(ctx *gin.Context) {
	db := global.SHOP_DB
	type Request struct {
		Id int64 `form:"id"`
	}

	req := Request{}
	if err := ctx.ShouldBind(&req); err != nil {
		ac.Error(ctx, err.Error())
		return
	}

	if err := db.Where("id = ?", req.Id).Delete(&model.Agreement{}).Error; err != nil {
		ac.Error(ctx, err.Error())
		return
	}

	ac.Success(ctx, "成功")
	return
}

func (ac *AgreementController) FindByKey(db *gorm.DB, key string, lang string) (*model.Agreement, error) {
	find := &model.Agreement{}
	err := db.Where("`key` = ?", key).Where("`language` = ?", lang).First(find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return find, nil
}
