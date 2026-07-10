package controller

import (
	"bytes"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/global"
	"ginshop.com/utils"
	"ginshop.com/utils/Paginate"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"strconv"
)

type ParamsRequest struct {
	Name string `form:"name" json:"name"`
	//Pid  int64  `form:"pid" json:"pid"` // pid 上级节点ID
}

type GoodsParamsRequest struct {
	CateId int `form:"cate_id" json:"cate_id"`
	Page   int `form:"page" json:"page"`
	Enable int `form:"enable" json:"enable"`
}

// GetCategoryLists 获取分类列表
func GetCategoryLists(ctx *gin.Context) {
	var params ParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}
	// 过滤page和pageSize
	paramsMap, _ := utils.AnyToMap(params)
	paramsMap["state"] = 1

	ParamsFilter := utils.ParamsFilter("page,pageSize,name", paramsMap)
	// 获取列表
	DB := global.SHOP_DB
	var Result []*model.Category
	var res model.Category
	var resErr error
	// 如果name条件不为空，追加模糊查询：position('搜索字符' in 字段)
	if len(params.Name) > 0 {
		resErr = DB.Where(ParamsFilter).Where("position(? in name)", params.Name).Find(&Result).Error
	} else {
		resErr = DB.Where(ParamsFilter).Order("id asc").Find(&Result).Error
	}
	language, _ := ctx.Get("Language")
	restotal := res.ToTree(Result, language.(string))

	if resErr != nil {
		fmt.Println(resErr.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	//for k, v := range restotal {
	//
	//	restotal[k].Name = utils.Languagebycode(language.(string), v.Name)
	//	if restotal[k].Children != nil {
	//		var models []*model.CategoryTrees
	//		for key, val := range *restotal[k].Children {
	//
	//			models = append(models, restotal[k].Children)
	//			restotal[k].Children = append(restotal[k].Children, models)
	//			//fmt.Println(key, val)
	//		}
	//	}
	//
	//}
	//service, err := utils.NewInviteService("my-secret-salt", 8)
	//if err != nil {
	//	log.Fatalf("Failed to initialize service: %v", err)
	//}
	//
	//userID := 1
	//
	//// 1. 生成邀请码
	//inviteCode, err := service.GenerateCode(userID)
	//if err != nil {
	//	log.Fatalf("Failed to generate code: %v", err)
	//}
	//fmt.Printf("用户ID: %d -> 邀请码: %s\n", userID, inviteCode)
	//
	//// 2. 解密邀请码
	//decodedID, err := service.DecodeCode(inviteCode)
	//if err != nil {
	//	log.Fatalf("Failed to decode code: %v", err)
	//}
	//fmt.Printf("邀请码: %s -> 解析用户ID: %d\n", inviteCode, decodedID)
	//
	//// 验证一致性
	//if userID == decodedID {
	//	fmt.Println("验证成功：ID一致")
	//} else {
	//	fmt.Println("验证失败：ID不一致")
	//}
	utils.Success(ctx, "成功", restotal)
}

// GetCategoryGoodsLists 获取分类下的全部商品
func GetCategoryGoodsLists(ctx *gin.Context) {
	var params GoodsParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	var cate model.Category
	DB := global.SHOP_DB
	pids := cate.GetCategoryIds(DB, params.CateId)
	fmt.Println(pids)

	var GoodsResult []*model.Goods
	var resErr error
	var ids = []int{}
	query := global.SHOP_DB.Model(model.Goods{})
	if len(pids) > 0 {
		pidsByte := new(bytes.Buffer)
		for _, value := range pids {
			ids = append(ids, value)
			_, err := fmt.Fprintf(pidsByte, "'%d',", value)
			if err != nil {
				return
			}
		}
		ids = append(ids, params.CateId)

		query = query.Where("goods_cate in (?)", ids)
	} else {

	}
	if params.Enable == 1 {

	}
	var count int64 = 0
	pageUp := strconv.Itoa(params.Page)

	if err := query.Count(&count).Error; err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	query = query.Scopes(Paginate.Paginate(pageUp, global.SHOP_CONFIG.System.PageSize)).Order("id desc ")
	err := query.Find(&GoodsResult).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	language, _ := ctx.Get("Language")
	for k, v := range GoodsResult {
		GoodsResult[k].GoodsName = utils.Languagebycode(language.(string), v.GoodsName)
		GoodsResult[k].GoodsCover = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, v.GoodsCover)
	}
	utils.Success(ctx, "获取成功", gin.H{
		"count": count,
		"data":  GoodsResult,
	})

}
func GetCategoryGoodsListsenable(ctx *gin.Context) {
	var params GoodsParamsRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	var cate model.Category
	DB := global.SHOP_DB
	pids := cate.GetCategoryIds(DB, params.CateId)
	fmt.Println(pids)

	var GoodsResult []*model.Goods
	var resErr error
	var ids = []int{}
	query := global.SHOP_DB.Model(model.Goods{})
	if len(pids) > 0 {
		pidsByte := new(bytes.Buffer)
		for _, value := range pids {
			ids = append(ids, value)
			_, err := fmt.Fprintf(pidsByte, "'%d',", value)
			if err != nil {
				return
			}
		}
		ids = append(ids, params.CateId)

		query = query.Where("goods_cate in (?)", ids)
	} else {

	}
	var funds model.AccountFunds
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("uid = ?", uid).Find(&funds).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	points := funds.Points
	query = query.Where("points>0")
	if points.GreaterThan(decimal.Zero) {
		query = query.Where("points <?", funds.Points)
	}

	var count int64 = 0
	pageUp := strconv.Itoa(params.Page)

	if err := query.Count(&count).Error; err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	query = query.Scopes(Paginate.Paginate(pageUp, global.SHOP_CONFIG.System.PageSize)).Order("id desc ")
	err = query.Find(&GoodsResult).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	language, _ := ctx.Get("Language")
	for k, v := range GoodsResult {
		GoodsResult[k].GoodsName = utils.Languagebycode(language.(string), v.GoodsName)
		GoodsResult[k].GoodsCover = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, v.GoodsCover)
	}
	utils.Success(ctx, "获取成功", gin.H{
		"count": count,
		"data":  GoodsResult,
	})

}
