package controller

import (
	"bytes"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UsdtWithdraw struct {
	Services service.SUsdtWithdraw
	BaseController
}

func (this *UsdtWithdraw) Index(ctx *gin.Context) {

	type Request struct {
		request.PageInfo
		Export bool `form:"export"`
	}
	req := Request{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		this.ErrorHtml(ctx, err.Error())
		return
	}

	if req.Export {
		query := global.SHOP_DB.Model(model.UsdtWithdrawModel{})
		if req.Keyword != "" {

			query.Where("username LIKE ?  ", "%"+req.Keyword+"%")
		}
		if *req.Status != 0 {

			query.Where("status =? ", req.Status)
		}

		f := excelize.NewFile()
		defer func() {
			f.Close()
		}()
		sheet1 := "Sheet1"
		_, err := f.NewSheet(sheet1)
		if err != nil {
			ctx.Error(err)
			return
		}

		row := 1
		col := 1
		_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), "序号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), "平台账号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), "提币地址")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), "地址类型")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), "金额")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), "状态")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), "添加时间")

		statusEnum := map[int]string{0: "待审核", 1: "通过", 2: "拒绝", 3: "待上链", 4: "上链失败", 5: "钱包上链中"}
		list := make([]model.UsdtWithdrawModel, 0)
		query.Order("id desc").FindInBatches(&list, 1000, func(db *gorm.DB, batch int) error {
			for _, info := range list {
				col += 1
				_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), info.Id)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), info.Username)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), info.WalletPath)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), info.PathType)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), info.Amount)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), statusEnum[info.Status])
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), info.CreateTime.Format(time.DateTime))
			}
			return nil
		})
		f.Path = fmt.Sprintf("提币申请-%s.xlsx", time.Now().Format("20060102150405"))
		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Path))
		ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		var buffer bytes.Buffer
		_ = f.Write(&buffer)
		http.ServeContent(ctx.Writer, ctx.Request, f.Path, time.Now(), bytes.NewReader(buffer.Bytes()))
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

	result, count := this.Services.GetAll(req.PageInfo)

	Search := map[string]interface{}{
		"page":   p,
		"limit":  size,
		"kw":     req.Keyword,
		"status": req.Status,
	}

	ctx.HTML(http.StatusOK, "userwithdraw_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *UsdtWithdraw) Pass(ctx *gin.Context) {
	var requids request.IdsReq
	var err = ctx.ShouldBind(&requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return

	}
	err = this.Services.Pass(requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "审核成功", requids)
}

func (this *UsdtWithdraw) PassUdun(ctx *gin.Context) {
	var requids request.IdsReq
	var err = ctx.ShouldBind(&requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return

	}
	if cap(requids.Ids) > 1 {
		this.Error(ctx, "提交U盾只能一条")
		return
	}
	err = this.Services.PassUdun(requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "审核成功", requids)
}

func (this *UsdtWithdraw) Refuse(ctx *gin.Context) {
	var requids request.IdsReq
	var err = ctx.ShouldBind(&requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return

	}
	session := sessions.Default(ctx)
	adminId := session.Get("adminId")
	userId := adminId.(uint)
	admin := model.Admin{}
	global.SHOP_DB.Where("id=?", userId).Find(&admin)
	superusername := admin.Account
	err = this.Services.Refuse(requids, superusername)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "审核成功", requids)
}
