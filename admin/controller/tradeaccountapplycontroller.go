package controller

import (
	"bytes"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type TradeAccountApply struct {
	Services service.StradeAccountApply
	BaseController
}

func (this *TradeAccountApply) Index(ctx *gin.Context) {
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

	p := req.Page
	if p == 0 {
		p = 1
	}

	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)
	size = 10
	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	if req.Export {
		query := global.SHOP_DB.Model(model.MTradeAccountApply{})
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
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), "正面")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), "反面")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), "账号类型")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), "名称")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), "电话")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), "状态")
		//_ = f.SetCellValue(sheet1, BuildCellPoint(row+8, col), "时间")

		list := make([]model.MTradeAccountApply, 0)
		query.Order("id desc").FindInBatches(&list, 1000, func(db *gorm.DB, batch int) error {
			for _, info := range list {
				col += 1
				_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), info.Id)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), info.Username)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), info.CarImg)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), info.CarImg2)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), "实名")
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), info.Name)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), info.Phone)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), info.Status)
				//_ = f.SetCellValue(sheet1, BuildCellPoint(row+8, col), info.CreateTime.Format(time.DateTime))
			}
			return nil
		})
		f.Path = fmt.Sprintf("交易账号申请-%s.xlsx", time.Now().Format("20060102150405"))
		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Path))
		ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		var buffer bytes.Buffer
		_ = f.Write(&buffer)
		http.ServeContent(ctx.Writer, ctx.Request, f.Path, time.Now(), bytes.NewReader(buffer.Bytes()))
		return
	}

	result, count := this.Services.GetAll(req.PageInfo)

	Search := map[string]interface{}{
		"page":   p,
		"limit":  size,
		"kw":     req.Keyword,
		"status": req.Status,
	}

	ctx.HTML(http.StatusOK, "tradeaccountapply_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}
func (this *TradeAccountApply) Pass(ctx *gin.Context) {
	var requids request.IdsReq
	var err = ctx.ShouldBind(&requids)
	if err != nil {
		this.Error(ctx, err)
		return

	}
	var models model.MTradeAccountApply

	err = ctx.ShouldBind(&models)
	if err != nil {
		this.Error(ctx, err.Error())
		return

	}
	err = this.Services.Pass(requids, models.Fee, models.Action)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "成功", requids)
}
func (this *TradeAccountApply) Refuse(ctx *gin.Context) {
	var requids request.IdsReq
	var err = ctx.ShouldBind(&requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return

	}

	err = this.Services.Refust(requids)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "拒绝成功", requids)
}
