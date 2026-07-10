package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Community struct {
}

func Index(ctx *gin.Context) {
	var req request.PageInfo
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(400, gin.H{"error": "ParseForm failed: " + err.Error()})
		return
	}

	// 2. 打印解析后的表单数据，确认数据是否存在
	fmt.Println("Form Data:", ctx.Request.Form)
	fmt.Println("PostForm Data:", ctx.Request.PostForm)

	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	p := req.Page
	if p == 0 {
		p = 1
	}
	var Services service.SOrders
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size

	result, count := Services.GetAll(req)
	data := map[string]any{
		"result": result,
		"count":  count,
	}
	language, _ := ctx.Get("Language")
	var orderitemserver service.SOrderItems
	for k, v := range result {
		result[k].CreatedAt = v.CreatedAt
		//var orderitem []*model.OrderItems
		var id request.GetById
		id.ID = uint(v.Id)
		orderitem, err := orderitemserver.GetByOrderID(id)
		if err != nil {
			//utils.Fail(ctx, "产品不存在", nil)
			continue
		}
		for key, value := range orderitem {
			orderitem[key].ProductName = utils.Languagebycode(language.(string), value.ProductName)
		}
		var user model.User
		uid := v.UserIdSell
		err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())

			continue
		}

		result[k].UserNameSell = utils.MaskString(user.Username, 1, 8)
		result[k].Chindren = orderitem
		result[k].Avatar = user.Avatar
		result[k].MemberLevel = *user.Level

	}
	//if req.OrderId != 0 && len(data) != 0 {
	//	data = data[0]
	//}
	utils.Success(ctx, "成功", data)
	return
}
