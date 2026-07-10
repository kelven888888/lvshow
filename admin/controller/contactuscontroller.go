package controller

import (
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactUsController struct {
	BaseController
}

func (h ContactUsController) List(ctx *gin.Context) {
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

	query := db.Model(&model.ContractUs{})

	var count int64 = 0
	if err = query.Count(&count).Error; err != nil {
		h.ErrorHtml(ctx, err.Error())
		return
	}

	req.Count = count

	list := make([]model.ContractUs, 0)
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

	ctx.HTML(http.StatusOK, "contactus_list.html", gin.H{
		"status":    "200",
		"list":      list,
		"req":       req,
		"webApiUrl": webApiUrl,
	})
}
