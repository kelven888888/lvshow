package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CPointRedeemRecord struct {
	Services service.SPointRedeemRecord
	BaseController
}

func (this *CPointRedeemRecord) Index(ctx *gin.Context) {

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
    	}

	ctx.HTML(http.StatusOK, "pointredeemrecord_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *CPointRedeemRecord) Edit(ctx *gin.Context) {

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
		global.SHOP_DB.Model(model.Language{}).Find(&language)
		// 查询权限列表
		ctx.HTML(http.StatusOK, "pointredeemrecord_form.html", gin.H{
			"status":   "200",
			"result":   result,
			"language": language,
			"IsUpdate": true,
		})
	} else {
		var models model.PointRedeemRecord

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
func (this *CPointRedeemRecord) Delete(ctx *gin.Context) {

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
func (this *CPointRedeemRecord) Add(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {

		// 查询权限列表
		var models model.PointRedeemRecord
		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Find(&language)
		ctx.HTML(http.StatusOK, "pointredeemrecord_form.html", gin.H{
			"status":   "200",
			"result":   models,
			"IsUpdate": false,
			"language": language,
		})
	} else {
		var models model.PointRedeemRecord

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
func (this *CPointRedeemRecord) Deletebatch(ctx *gin.Context) {
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
