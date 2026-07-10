package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
)

type GoodsDetailsParamsRequest struct {
	ID int64 `form:"id" json:"id"`
}

type GoodsDetailsResult struct {
	model.Goods
	CateName string `json:"cate_name"` // 分类名称
	//GoodsPropertyList []string `json:"goodsPropertyList"` // 商品属性列表
	Redeemable int `json:"redeemable"`
}

// GoodsDetails 商品详情
func GoodsDetails(ctx *gin.Context) {
	var params GoodsDetailsParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}
	var GoodDetail GoodsDetailsResult
	DB := global.SHOP_DB
	resErr := DB.Model(&model.Goods{}).Select("goods.*, categories.name as cate_name").Joins("left join categories on categories.id = goods.goods_cate").Where("goods.id = ?", params.ID).Scan(&GoodDetail).Error
	if resErr != nil {
		utils.Fail(ctx, resErr.Error(), nil)
		return
	}
	language, _ := ctx.Get("Language")

	GoodDetail.GoodsContent = utils.Languagebycode(language.(string), GoodDetail.GoodsContent)
	GoodDetail.GoodsCover = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, GoodDetail.GoodsCover)
	GoodDetail.GoodsName = utils.Languagebycode(language.(string), GoodDetail.GoodsName)
	GoodDetail.CateName = utils.Languagebycode(language.(string), GoodDetail.CateName)
	GoodDetail.GoodsContent = utils.Languagebycode(language.(string), GoodDetail.GoodsContent)
	redeemable := 0
	user_id, exit := ctx.Get("user_id")
	if exit {
		var fund model.AccountFunds
		global.SHOP_DB.Where("uid=?", user_id).Find(&fund)
		if fund.Points.GreaterThan(GoodDetail.Points) {
			redeemable = 1
		}
	} else {
		redeemable = 0
	}
	GoodDetail.Redeemable = redeemable
	//GoodDetail.GoodsPropertyList = strings.Split(GoodDetail.GoodsProperty, ";")

	utils.Success(ctx, "获取成功", GoodDetail)
	return
}
