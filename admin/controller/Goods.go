package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CGoods struct {
	Services service.SGoods
	BaseController
}

func (this *CGoods) Index(ctx *gin.Context) {

	var req request.PageInfo
	err := ctx.ShouldBind(&req)
	if err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}

	p := req.Page
	if p == 0 {
		p = 1
	}

	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size

	result, count := this.Services.GetAll(req)

	Search := map[string]interface{}{
		"page":         p,
		"limit":        size,
		"kw":           req.Keyword,
		"search_field": req.SearchField,
		"play_id":      req.PlayId,
	}

	ctx.HTML(http.StatusOK, "goods_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *CGoods) Edit(ctx *gin.Context) {

	if ctx.Request.Method == "GET" {
		var id request.GetById

		err := ctx.ShouldBind(&id)
		if err != nil {
			this.ErrorHtml(ctx, err.Error())
		}

		result, err := this.Services.GetByID(id)
		if err != nil {
			this.ErrorHtml(ctx, err.Error())
		}
		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&language)
		var category []model.Category
		global.SHOP_DB.Model(model.Category{}).Where("state=1").Find(&category)
		// 查询权限列表
		ctx.HTML(http.StatusOK, "goods_form.html", gin.H{
			"status":   "200",
			"result":   result,
			"language": language,
			"IsUpdate": true,
			"category": category,
		})
	} else {
		var models model.Goods

		err := ctx.ShouldBind(&models)
		if err != nil {
			this.Error(ctx, err.Error())
			return
		}

		err = this.Services.Save(&models)
		if err != nil {
			this.Error(ctx, err.Error())
			return
		}

		this.Success(ctx, "成功")

	}
}
func (this *CGoods) Delete(ctx *gin.Context) {

	var req request.GetById
	err := ctx.ShouldBind(&req)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	err = this.Services.Delete(req.Uint32())
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "删除成功")

}
func (this *CGoods) Add(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {

		// 查询权限列表
		var models model.Goods
		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&language)
		var category []model.Category
		global.SHOP_DB.Model(model.Category{}).Where("state=1").Find(&category)
		ctx.HTML(http.StatusOK, "goods_form.html", gin.H{
			"status":   "200",
			"result":   models,
			"IsUpdate": false,
			"language": language,
			"category": category,
		})
	} else {
		var models model.Goods

		err := ctx.ShouldBind(&models)

		if err != nil {
			this.Error(ctx, err.Error())
			return
		}
		err = this.Services.Save(&models)
		if err != nil {
			this.Error(ctx, err.Error())
			return
		}

		this.Success(ctx, "成功")
	}

}
func (this *CGoods) Deletebatch(ctx *gin.Context) {
	var req request.IdsReq
	err := ctx.ShouldBind(&req)
	if err != nil {

		this.Error(ctx, err.Error())
		return

	}

	err = this.Services.Deletebatch(req)
	if err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}
	this.Success(ctx, "删除成功")

}
