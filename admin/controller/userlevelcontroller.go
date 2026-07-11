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

type UserlevelController struct {
	Services service.Userlevel
	BaseController
}

func (this *UserlevelController) Index(ctx *gin.Context) {

	//for i := 10; i < 101; i++ {
	//	var level model.MemberLevel
	//	global.SHOP_DB.Where("level=?", i).Find(&level)
	//	if level.Id == 0 {
	//		var levels model.MemberLevel
	//		fmt.Println(fmt.Sprintf("%+v", levels))
	//		var value int = i
	//		levels.Level = &value
	//		levels.CreateTime = time.Now()
	//		var display int = 1
	//		levels.IsDisplay = &display
	//		levels.OrderBy = 200
	//
	//		levels.Title = fmt.Sprintf("{\"zh-hant\":\"VIP%d\",\"en\":\"VIP%d\",\"ml\":\"VIP%d\"}", i, i, i)
	//
	//
	//		global.SHOP_DB.Model(model.MemberLevel{}).Save(&levels)
	//	}
	//
	//}

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
		"page":  p,
		"limit": size,
		"kw":    req.Keyword,
	}

	ctx.HTML(http.StatusOK, "userlevel_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *UserlevelController) Edit(ctx *gin.Context) {

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
		var coupon []model.TbCoupon
		global.SHOP_DB.Where("status=1").Find(&coupon)
		// 查询权限列表
		ctx.HTML(http.StatusOK, "userlevel_form.html", gin.H{
			"status":   "200",
			"result":   result,
			"language": language,
			"IsUpdate": true,
			"Coupon":   coupon,
		})
	} else {
		var models model.MemberLevel

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
func (this *UserlevelController) Delete(ctx *gin.Context) {

	var req request.GetById
	err := ctx.ShouldBind(&req)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	err = this.Services.Delete(req.Uint32())
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		this.Error(ctx, err.Error())

		return
	}

	this.Success(ctx, "删除成功")

}
func (this *UserlevelController) Add(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {

		// 查询权限列表
		var models model.MemberLevel
		var language []model.Language
		global.SHOP_DB.Model(model.Language{}).Where("status=1").Find(&language)
		ctx.HTML(http.StatusOK, "userlevel_form.html", gin.H{
			"status":   "200",
			"result":   models,
			"IsUpdate": false,
			"language": language,
		})
	} else {
		var models model.MemberLevel

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
