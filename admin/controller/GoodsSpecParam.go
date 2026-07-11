package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type CGoodsSpecParam struct {
	Services service.SGoodsSpecParam
	BaseController
}

func (this *CGoodsSpecParam) Index(ctx *gin.Context) {

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
		"status":       req.Status,
	}

	ctx.HTML(http.StatusOK, "goodsspecparam_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *CGoodsSpecParam) Edit(ctx *gin.Context) {

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
		// 查询权限列表
		ctx.HTML(http.StatusOK, "goodsspecparam_form.html", gin.H{
			"status":   "200",
			"result":   result,
			"language": language,
			"IsUpdate": true,
		})
	} else {
		var models model.GoodsSpecParam

		err := ctx.ShouldBind(&models)
		if err != nil {
			this.Error(ctx, err.Error())
			return
		}
		fmt.Println("=======================", models)
		err = this.Services.Save(&models)
		if err != nil {
			this.Error(ctx, err.Error())
			return
		}

		this.Success(ctx, "成功")

	}
}
func (this *CGoodsSpecParam) Delete(ctx *gin.Context) {

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
func (this *CGoodsSpecParam) Add(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {

		// 查询权限列表
		var models model.GoodsSpecParam
		var language []model.Language
		var Result []*model.Category

		// 如果name条件不为空，追加模糊查询：position('搜索字符' in 字段)
		global.SHOP_DB.Where("state=1").Find(&Result)
		var Scategory service.SCategory
		category := Scategory.GetCategory(true)
		var req request.PageInfo
		req.Count = true
		req.Limit = 1000
		req.Offset = 0
		var server service.SGoodsSpecGroup
		result, _ := server.GetAll(req)

		restotal := this.getSelectTree(category, 0, 0)
		resluts := this.getSelectTrees(result, 0)
		global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&language)
		ctx.HTML(http.StatusOK, "goodsspecparam_form.html", gin.H{
			"status":    "200",
			"result":    models,
			"IsUpdate":  false,
			"language":  language,
			"category":  restotal,
			"sepcgroup": resluts,
		})
	} else {
		var models model.GoodsSpecParam

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
func (this *CGoodsSpecParam) Deletebatch(ctx *gin.Context) {
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

// 格式化为select表单树
func (this *CGoodsSpecParam) getSelectTree(categorys []*model.Category, pid int, level int) string {
	html := ""
	for _, category := range categorys {
		sel := ""
		if category.Id == pid {
			sel = " selected"
		}
		html += fmt.Sprintf(`<option value=%d %s>%s%s</option>`,
			category.Id,
			sel,
			strings.Repeat("  -", level*4),
			utils.Languagebycode("zh-cn", category.Name),
		)

		if category.Id != pid && len(category.Child) > 0 {
			html += this.getSelectTree(category.Child, pid, level+1)
		}
	}

	return html
}
func (this *CGoodsSpecParam) getSelectTrees(categorys []model.GoodsSpecGroup, pid int) string {
	html := ""
	for _, category := range categorys {
		sel := ""
		if category.Id == pid {
			sel = " selected"
		}
		html += fmt.Sprintf(`<option disabled  id="spec_group_%d"  value=%d %s>%s%s</option>`,
			category.Cid,
			category.Id,
			sel,
			strings.Repeat("  -", 4),
			utils.Languagebycode("zh-cn", category.Name),
		)

	}

	return html
}
