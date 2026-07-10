package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountFeedbackController struct {
	BaseController
}

func (h AccountFeedbackController) List(ctx *gin.Context) {
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

	query := db.Model(&model.AccountFeedback{})

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.AccountFeedback, 0)
	if err = query.
		Order("id desc").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&list).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}
	webApiUrl := global.SHOP_CONFIG.System.WebApiURL
	ctx.HTML(http.StatusOK, "account_feedback_list.html", gin.H{
		"status":    "200",
		"list":      list,
		"req":       req,
		"webApiUrl": webApiUrl,
	})
}

func (h AccountFeedbackController) Delete(ctx *gin.Context) {
	db := global.SHOP_DB
	type Request struct {
		Id int64 `form:"id"`
	}

	req := Request{}
	if err := ctx.ShouldBind(&req); err != nil {
		h.Error(ctx, err.Error())
		return
	}

	if err := db.Where("id = ?", req.Id).Delete(&model.AccountFeedback{}).Error; err != nil {
		h.Error(ctx, err.Error())
		return
	}

	h.Success(ctx, "成功")
	return
}
