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

type FundsRecharge struct {
	Services service.FundsRecharge
	BaseController
}

func (this *FundsRecharge) Index(ctx *gin.Context) {
	type Request struct {
		request.PageInfo
		Export bool `form:"export"`
	}
	var req Request
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

	if req.Export {
		query := global.SHOP_DB.Model(model.FundRecharge{})
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
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), "创建时间")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), "账号")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), "地址")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), "类型")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), "金额")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), "手续费")
		_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), "hash")

		list := make([]model.FundRecharge, 0)
		query.Order("id desc").FindInBatches(&list, 1000, func(db *gorm.DB, batch int) error {
			for _, info := range list {
				col += 1
				_ = f.SetCellValue(sheet1, BuildCellPoint(row, col), info.Id)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+1, col), info.CreateTime.Format(time.DateTime))
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+2, col), info.Username)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+3, col), info.Address)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+4, col), info.PathType)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+5, col), info.Amount)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+6, col), info.Cms)
				_ = f.SetCellValue(sheet1, BuildCellPoint(row+7, col), info.Hash)
			}
			return nil
		})
		f.Path = fmt.Sprintf("账户充值记录-%s.xlsx", time.Now().Format("20060102150405"))
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

	ctx.HTML(http.StatusOK, "fundsrecharge_index.html", gin.H{
		"status": "200",
		"List":   result,
		"Count":  count,
		"Search": Search,
	})
}

var rows = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func BuildCellPoint(row, col int) string {
	offsetIndex := row % len(rows)
	offsetRow := row / len(rows)

	if offsetIndex == 0 {
		offsetIndex = 26
		offsetRow -= 1
	}
	if offsetRow > 0 {
		return fmt.Sprintf("%s%s%d", rows[offsetRow-1], rows[offsetIndex-1], col)
	}

	return fmt.Sprintf("%s%d", rows[offsetIndex-1], col)
}
func (this *FundsRecharge) Do(ctx *gin.Context) {
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
	err = this.Services.Do(requids, requids.Action)
	if err != nil {
		this.Error(ctx, err.Error())
		return
	}

	this.Success(ctx, "成功", requids)
}
