package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

type AccountFunds struct {
	Services     service.SaccountFund
	AgentService service.AgentUser
	Serviceslog  service.TradeAccountFundsLog
	BaseController
}

func (this *AccountFunds) Index(ctx *gin.Context) {

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
		"page":   p,
		"limit":  size,
		"kw":     req.Keyword,
		"status": req.Status,
	}

	ctx.HTML(http.StatusOK, "accountfunds_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *AccountFunds) FundsEdit(ctx *gin.Context) {
	method := ctx.Request.Method

	if method == "GET" {
		var User model.AccountFunds

		if err := ctx.ShouldBind(&User); err == nil {

			global.SHOP_DB.First(&User)

			var agentinfo model.AgentInfo
			agentinfo, err = this.AgentService.GetCode(User.Username)
			ctx.HTML(http.StatusOK, "accountfunds_fundsedit.html", gin.H{

				"admininfo": User,

				"agentinfo": agentinfo,

				"IsUpdate": true,
			})
		} else {
			this.Error(ctx, "查询错误", err.Error())
		}
	} else {
		var request request.AccountFunds
		if err := ctx.ShouldBindJSON(&request); err == nil {

			//global.SHOP_DB.Where("username = ?", adminUser.Username).First(&adminUser)
			//id := adminUser.Id

			//if request.Amount == 0 {
			//	this.Error(ctx, "失败", "数量不能为0")
			//	return
			//}
			tran := global.SHOP_DB.Begin()
			if request.Amount != 0 {

				var accountfunds model.AccountFunds
				err = tran.Where("id= ?", request.Id).Find(&accountfunds).Error
				if err != nil {
					tran.Rollback()
					this.Error(ctx, "失败", err.Error())
					return
				}
				//fundold := accountfunds.AvaFunds
				value := decimal.NewFromFloat(request.Amount)
				if value.Add(accountfunds.AvaFunds).LessThan(decimal.NewFromFloat(0)) {
					tran.Rollback()
					this.Error(ctx, "余额不足", "")
					return
				}
				accountfunds.AvaFunds = value.Add(accountfunds.AvaFunds)

				//err = tran.Where("ava_funds=?", fundold).Updates(&accountfunds).Error
				err = tran.Exec("update account_funds set ava_funds=ava_funds+? where id=?", request.Amount, request.Id).Error
				if err != nil {
					tran.Rollback()
					this.Error(ctx, "失败", err.Error())
					return
				}
				var saccountfundslog service.AccountFundsLog
				session := sessions.Default(ctx)
				adminId := session.Get("adminId")
				userId := adminId.(uint)
				admin := model.Admin{}
				global.SHOP_DB.Where("id=?", userId).Find(&admin)
				superusername := admin.Account

				err, _ := saccountfundslog.Createlog(accountfunds.Username, value, utils.Adminlogtype, fmt.Sprintf("%s手工操作", superusername), 1)
				if err != nil {
					this.Error(ctx, "修改失败", err.Error())
					return
				}
			}
			if request.Points != 0 {

				var accountfunds model.AccountFunds
				err = tran.Where("id= ?", request.Id).Find(&accountfunds).Error
				if err != nil {
					tran.Rollback()
					this.Error(ctx, "失败", err.Error())
					return
				}
				//pointold := accountfunds.Points
				value := decimal.NewFromFloat(request.Points)
				if value.Add(accountfunds.Points).LessThan(decimal.NewFromFloat(0)) {
					tran.Rollback()
					this.Error(ctx, "余额不足", "")
					return
				}
				accountfunds.Points = value.Add(accountfunds.Points)

				//err = tran.Where("points=?", pointold).Updates(&accountfunds).Error
				err = tran.Exec("update account_funds set points=points+? where id=?", request.Points, request.Id).Error
				if err != nil {
					tran.Rollback()
					this.Error(ctx, "失败", err.Error())
					return
				}
				var saccountfundslog service.AccountFundsLog
				session := sessions.Default(ctx)
				adminId := session.Get("adminId")
				userId := adminId.(uint)
				admin := model.Admin{}
				global.SHOP_DB.Where("id=?", userId).Find(&admin)
				superusername := admin.Account

				err, _ := saccountfundslog.Createlog(accountfunds.Username, value, utils.Adminlogtype, fmt.Sprintf("%s手工操作", superusername), 2)
				if err != nil {
					this.Error(ctx, "修改失败", err.Error())
					return
				}
			}
			tran.Commit()

			//service.RunPublisher(ctx, "shop_message", "修改"+fmt.Sprintf("%v", id)+"成功")
			this.Success(ctx, "修改成功")
		} else {
			fmt.Println(err.Error())
			this.Error(ctx, "修改失败", err.Error())
		}
	}

}
