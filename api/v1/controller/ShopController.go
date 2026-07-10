package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// GoodsDetails 商品详情
func Store(ctx *gin.Context) {
	var req request.PageInfo
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	p := req.Page
	if p == 0 {
		p = 1
	}
	var Services service.SUserStore
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {

		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	req.Username = user.Username

	result, count := Services.GetAll(req)
	data := map[string]any{
		"result": result,
		"count":  count,
	}
	language, _ := ctx.Get("Language")
	var goodserver service.SGoods
	for k, v := range result {
		result[k].CreatedAt = v.CreatedAt
		var goods *model.Goods
		var id request.GetById
		id.ID = uint(v.GoodsId)
		goods, err := goodserver.GetByID(id)
		if err != nil {
			utils.Fail(ctx, "商品不存在", nil)
		}

		result[k].GoodsName = utils.Languagebycode(language.(string), result[k].GoodsName)
		result[k].GoodsCover = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, goods.GoodsCover)
	}
	utils.Success(ctx, "成功", data)
	return
}
func Ordersell(ctx *gin.Context) {

	var req request.IdsReqgood
	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return

	}
	fmt.Println(fmt.Sprintf("%+v", req))

	if len(req.Qtys) == 0 {

		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if len(req.Qtys) != len(req.Ids) {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if req.Prices == decimal.Zero {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	lockkey := fmt.Sprintf("order_%s", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	tran := global.SHOP_DB.Begin()
	var orders model.Orders

	orders.OrderSn = utils.GenerateOrderID()
	now := model.LocalTime(time.Now())
	orders.CreatedAt = &now

	orders.UserIdSell = user.Id
	orders.UserNameSell = user.Username
	orderinsert := tran.Save(&orders)
	if orderinsert.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(orderinsert.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	var goodserver service.SGoods
	var totalprice = decimal.Zero
	var totalqty = 0
	var totalproductqty = 0
	for k, v := range req.Ids {
		var id request.GetById
		var stores model.UserStore
		global.SHOP_DB.Where("id=? and user_id=?", v, user.Id).Find(&stores)
		if stores.Id == 0 {
			tran.Rollback()
			utils.Fail(ctx, "商品不存在", nil)
			return
		}
		id.ID = uint(stores.GoodsId)
		goods, err := goodserver.GetByID(id)
		if err != nil {
			utils.Fail(ctx, "商品不存在", nil)
			tran.Rollback()
			return
		}
		if req.Qtys[k] == 0 {
			tran.Rollback()
			utils.Fail(ctx, "参数错误", nil)
			return
		}
		var orderitem model.OrderItems
		orderitem.Qty = req.Qtys[k]
		orderitem.ProductId = goods.Id
		orderitem.ProductImg = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, goods.GoodsCover)
		orderitem.ProductName = goods.GoodsName
		orderitem.UnitPrice = goods.UnitPrice
		//orderitem.SellPrice = UnitPrice
		orderitem.OrderId = orders.Id
		orderitem.RewardType = goods.RewardType
		//orderitem.TotalPrice = req.Prices[k].Mul(decimal.NewFromInt(int64(req.Qtys[k])))
		totalprice = totalprice.Add(orderitem.TotalPrice)
		totalqty = totalqty + orderitem.Qty
		totalproductqty = totalproductqty + 1

		err = tran.Save(&orderitem).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return

		}
		err = tran.Exec("update user_store set qty=qty-? where user_id=? and id=?", req.Qtys[k], user.Id, v).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "数量不足", nil)
			return
		}
	}
	orders.TotalQty = totalqty
	orders.TotalAmount = req.Prices
	if req.Prices == decimal.Zero {
		tran.Rollback()
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	orders.TotalProductQty = totalproductqty
	err = tran.Save(&orders).Error
	if orderinsert.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(orderinsert.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	tran.Commit()

	utils.Success(ctx, "成功", "")

}
func Orders(ctx *gin.Context) {
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
	fmt.Println(fmt.Sprintf("%+v", req))
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	req.UserId = user.Id

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
		result[k].Chindren = orderitem

	}
	//if req.OrderId != 0 && len(data) != 0 {
	//	data = data[0]
	//}
	utils.Success(ctx, "成功", data)
	return
}
func OrdersCancel(ctx *gin.Context) {
	var req request.GetById
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	var Services service.SOrders

	result, err := Services.GetByID(req)
	if err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if result.Id == 0 {
		utils.Fail(ctx, "订单不存在", nil)
		return
	}
	//if *result.Status != 0 {
	//	utils.Fail(ctx, "订单状态错误", nil)
	//	return
	//}
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	lockkey := fmt.Sprintf("order_%s", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	if result.UserIdSell != user.Id {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	tran := global.SHOP_DB.Begin()
	*result.Status = 4

	results := tran.Model(model.Orders{}).Where("status=0 and id=?", result.Id).Updates(map[string]interface{}{
		"status": 4,
	})
	if results.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(results.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if results.RowsAffected != 1 {
		tran.Rollback()
		global.SHOP_LOG.Error(fmt.Sprintf("影响记录%d", results.RowsAffected))
		utils.Fail(ctx, "失败", nil)
		return
	}
	var orderitemserver service.SOrderItems
	var id request.GetById
	id.ID = uint(result.Id)
	orderitem, err := orderitemserver.GetByOrderID(id)

	for _, value := range orderitem {
		err = tran.Exec("update user_store set qty=qty+? where user_id=? and goods_id=?", value.Qty, user.Id, value.ProductId).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
	}
	tran.Commit()
	utils.Success(ctx, "成功", nil)

}
func Orderbuy(ctx *gin.Context) {

	var req request.GetById
	err := ctx.ShouldBind(&req)
	if err != nil {

		utils.Fail(ctx, "参数错误", nil)
		return

	}
	if req.ID == 0 {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if req.TradePassword == "" {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	tran := global.SHOP_DB.Begin()
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = tran.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.TradePassword), []byte(req.TradePassword))
	if err != nil {
		utils.Fail(ctx, "交易密码错误", nil)
		return
	}
	lockkey := fmt.Sprintf("order_%s", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	var orders model.Orders
	tran.Where("id=? and status=0", req.ID).Find(&orders)
	if orders.Id == 0 {
		tran.Rollback()

		utils.Fail(ctx, "订单状态错误", nil)
		return
	}
	if user.Id == orders.UserIdSell {
		tran.Rollback()

		utils.Fail(ctx, "不能购买自己订单", nil)
		return
	}

	var funds model.AccountFunds
	tran.Where("username=?", user.Username).Find(&funds)
	if funds.Id == 0 {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if funds.AvaFunds.LessThan(orders.TotalAmount) {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "余额不足", nil)
		return
	}
	var orderitemserver service.SOrderItems

	var id request.GetById
	id.ID = uint(orders.Id)
	orderitem, err := orderitemserver.GetByOrderID(id)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	//更新购买者余额及添加流水及更仓库
	err = tran.Exec("update account_funds set ava_funds=ava_funds-? where username=?", orders.TotalAmount, user.Username).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "余额不足", nil)
		return
	}
	var saccountfundslog service.AccountFundsLog
	logidarr := []uint{}
	err, logid := saccountfundslog.Createlog(user.Username, orders.TotalAmount.Neg(), utils.BUYGOOD, fmt.Sprintf("支付订单-id%d", orders.Id), 1)
	logidarr = append(logidarr, logid)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	for _, value := range orderitem {

		var storess model.UserStore
		global.SHOP_DB.Where("user_id=? and goods_id=?", user.Id, value.ProductId).Find(&storess)
		if storess.Id == 0 {
			storess.Price = value.UnitPrice
			storess.UserId = user.Id
			storess.Username = user.Username
			storess.GoodsName = value.ProductName
			storess.GoodsId = value.ProductId
			storess.RewardType = value.RewardType
			storess.Qty = 1
			var storeserver service.SUserStore
			err = storeserver.Save(&storess)
			if err != nil {
				tran.Rollback()
				var mlog model.AccountFundsLog
				global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
				global.SHOP_LOG.Error(err.Error())
				utils.Fail(ctx, "失败", nil)
				return
			}
		} else {
			err = global.SHOP_DB.Exec("update user_store set qty=qty+1 where user_id=? and goods_id=?", user.Id, value.Id).Error
			if err != nil {
				var mlog model.AccountFundsLog
				global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
				tran.Rollback()
				global.SHOP_LOG.Error(err.Error())
				utils.Fail(ctx, "失败", nil)
				return
			}
		}
	}
	var usersell model.User
	tran.Where("id=?", orders.UserIdSell).Find(&usersell)
	if usersell.Id == 0 {
		tran.Rollback()
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		global.SHOP_LOG.Error("用户不存在")
		utils.Fail(ctx, "失败", nil)
		return
	}
	//更新出售者余额及添加流水及更仓库
	err = tran.Exec("update account_funds set ava_funds=ava_funds+? where username=?", orders.TotalAmount, usersell.Username).Error
	if err != nil {
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "余额不足", nil)
		return
	}
	//var saccountfundslog service.AccountFundsLog

	err, logid = saccountfundslog.Createlog(usersell.Username, orders.TotalAmount, utils.SELLGOOD, fmt.Sprintf("订单收入-id%d", orders.Id), 1)
	logidarr = append(logidarr, logid)
	if err != nil {
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	//orders.UserIdBuy = user.Id
	//status := 1
	//orders.Status = &status
	//orders.PayAmount = orders.TotalAmount

	now := model.LocalTime(time.Now())
	//orders.PayTime = &now
	//err = tran.Save(&orders).Error
	results := tran.Model(model.Orders{}).Where("status=0 and id=?", orders.Id).Updates(map[string]interface{}{
		"status":        1,
		"user_id_buy":   user.Id,
		"pay_amount":    orders.TotalAmount,
		"user_name_buy": user.Username,
		"pay_time":      &now,
	})
	if results.Error != nil {
		tran.Rollback()
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		global.SHOP_LOG.Error(results.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if results.RowsAffected != 1 {
		tran.Rollback()
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		global.SHOP_LOG.Error(fmt.Sprintf("影响记录%d", results.RowsAffected))
		utils.Fail(ctx, "失败", nil)
		return
	}

	tran.Commit()
	utils.Success(ctx, "成功", nil)

}
func PointsRedeem(ctx *gin.Context) {

	var req request.GetById
	err := ctx.ShouldBind(&req)
	if err != nil {

		utils.Fail(ctx, "参数错误", nil)
		return

	}
	if req.ID == 0 {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	tran := global.SHOP_DB.Begin()
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = tran.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	lockkey := fmt.Sprintf("order_%s", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	var goods model.Goods
	tran.Where("id=? and goods_status=1", req.ID).Find(&goods)
	if goods.Id == 0 {
		tran.Rollback()

		utils.Fail(ctx, "商品不存在", nil)
		return
	}

	var funds model.AccountFunds
	tran.Where("username=?", user.Username).Find(&funds)
	if funds.Id == 0 {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	if goods.Points.Equal(decimal.Zero) {
		tran.Rollback()
		utils.Fail(ctx, "不支持兑换", nil)
		return
	}
	if funds.AvaFunds.LessThan(goods.Points) {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "余额不足", nil)
		return
	}

	//更新购买者余额及添加流水及更仓库
	err = tran.Exec("update account_funds set points=points-? where uid=?", goods.Points, user.Id).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "余额不足", nil)
		return
	}
	var saccountfundslog service.AccountFundsLog
	logidarr := []uint{}
	err, logid := saccountfundslog.Createlog(user.Username, goods.Points.Neg(), utils.PointRedeem, fmt.Sprintf("积分兑换产品id%d", goods.Id), 2)
	logidarr = append(logidarr, logid)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	var storess model.UserStore
	global.SHOP_DB.Where("user_id=? and goods_id=?", user.Id, goods.Id).Find(&storess)
	if storess.Id == 0 {
		storess.Price = goods.UnitPrice
		storess.UserId = user.Id
		storess.Username = user.Username
		storess.GoodsName = goods.GoodsName
		storess.GoodsId = goods.Id
		storess.RewardType = goods.RewardType
		storess.Qty = 1
		var storeserver service.SUserStore
		err = storeserver.Save(&storess)
		if err != nil {
			tran.Rollback()
			var mlog model.AccountFundsLog
			global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
	} else {
		err = global.SHOP_DB.Exec("update user_store set qty=qty+1 where user_id=? and goods_id=?", user.Id, goods.Id).Error
		if err != nil {
			var mlog model.AccountFundsLog
			global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
	}
	now := model.LocalTime(time.Now())
	var pointrecord model.PointRedeemRecord
	pointrecord.CreatedAt = &now
	pointrecord.ProductId = goods.Id
	pointrecord.UserId = user.Id
	pointrecord.Username = user.Username
	pointrecord.Points = goods.Points
	pointrecord.ProductName = goods.GoodsName
	err = tran.Model(model.PointRedeemRecord{}).Save(&pointrecord).Error
	if err != nil {
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	//更新库存
	err = tran.Exec("update goods set goods_stock=goods_stock-1 where id=?", goods.Id).Error
	if err != nil {
		var mlog model.AccountFundsLog
		global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "库存不足", nil)
		return
	}

	tran.Commit()
	utils.Success(ctx, "成功", nil)

}
func PointsRedeemrecord(ctx *gin.Context) {
	var req request.PageInfo
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	p := req.Page
	if p == 0 {
		p = 1
	}
	var Services service.SPointRedeemRecord
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	req.Username = user.Username

	result, count := Services.GetAll(req)
	data := map[string]any{
		"result": result,
		"count":  count,
	}
	language, _ := ctx.Get("Language")
	for k, _ := range result {

		var good model.Goods
		global.SHOP_DB.Where("id=?", result[k].ProductId).Find(&good)
		result[k].ProductName = utils.Languagebycode(language.(string), result[k].ProductName)

		result[k].Image = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, good.GoodsCover)

	}
	//var withdraw []model.FundRecharge
	//global.SHOP_DB.Where("username=?", user.Username).Order("id desc ").Find(&withdraw)

	utils.Success(ctx, "成功", data)
}

func DoOrdershipping(ctx *gin.Context) {

	var req request.IdsReqgood
	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return

	}
	fmt.Println(fmt.Sprintf("%+v", req))

	if len(req.Qtys) == 0 {

		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if len(req.Qtys) != len(req.Ids) {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if req.ShippingaddrId == 0 {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	lockkey := fmt.Sprintf("order_%s", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	tran := global.SHOP_DB.Begin()
	var orders model.OrdersShipping

	orders.OrderSn = utils.GenerateOrderID()
	now := model.LocalTime(time.Now())
	orders.CreatedAt = &now

	orders.UserId = user.Id
	orderinsert := tran.Save(&orders)
	if orderinsert.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(orderinsert.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	var goodserver service.SGoods

	var totalqty = 0
	var totalproductqty = 0
	for k, v := range req.Ids {
		var id request.GetById
		var stores model.UserStore
		global.SHOP_DB.Where("id=?", v).Find(&stores)
		if stores.Id == 0 {
			tran.Rollback()
			global.SHOP_LOG.Error(orderinsert.Error.Error())
			utils.Fail(ctx, "商品不存在", nil)
			return
		}
		id.ID = uint(stores.Id)
		goods, err := goodserver.GetByID(id)
		if err != nil {
			utils.Fail(ctx, "商品不存在", nil)
			tran.Rollback()
		}
		var orderitem model.OrderItemsShipping
		orderitem.Qty = req.Qtys[k]
		orderitem.ProductId = goods.Id
		orderitem.ProductImg = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, goods.GoodsCover)
		orderitem.ProductName = goods.GoodsName

		//orderitem.SellPrice = UnitPrice
		orderitem.OrderId = orders.Id
		orderitem.RewardType = goods.RewardType
		//orderitem.TotalPrice = req.Prices[k].Mul(decimal.NewFromInt(int64(req.Qtys[k])))

		totalqty = totalqty + orderitem.Qty
		totalproductqty = totalproductqty + 1

		err = tran.Save(&orderitem).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return

		}
		err = tran.Exec("update user_store set qty=qty-? where user_id=? and id=?", req.Qtys[k], user.Id, v).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "数量不足", nil)
			return
		}
	}
	var shipaddr model.ShippingAddresses
	err = global.SHOP_DB.Where("id=? and username=?", req.ShippingaddrId, user.Username).Find(&shipaddr).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		tran.Rollback()
		utils.Fail(ctx, "收货地址不存在", nil)
		return
	}
	if shipaddr.Id == 0 {
		tran.Rollback()
		utils.Fail(ctx, "收货地址不存在", nil)
		return
	}
	orders.TotalQty = totalqty
	orders.ReceiverPhone = shipaddr.ReceiverPhone
	orders.ReceiverName = shipaddr.ReceiverName
	addr := fmt.Sprintf("%s %s %s", shipaddr.Area, shipaddr.City, shipaddr.AddressLine1)
	orders.ReceiverAddress = addr
	orders.TotalProductQty = totalproductqty
	orders.Username = user.Username
	orders.PostalCode = shipaddr.PostalCode
	orders.Remarks = req.Remarks
	err = tran.Save(&orders).Error
	if orderinsert.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(orderinsert.Error.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	//扣邮费
	var conf service.Config
	rate, err := conf.GetKeyValue("SHIPPING_FEE")
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	f, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	shipping_fee := utils.Float64ToDecimal(f)
	if shipping_fee.GreaterThan(decimal.Zero) {
		err = tran.Exec("update account_funds set ava_funds=ava_funds-? where username=?", shipping_fee, user.Username).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "余额不足", nil)
			return
		}
		orders.ShippingFee = shipping_fee
		err = tran.Save(&orders).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
		var saccountfundslog service.AccountFundsLog

		err, _ = saccountfundslog.Createlog(user.Username, shipping_fee.Neg(), utils.SHIPPINGFEE, fmt.Sprintf("邮费"), 1)

		if err != nil {
			tran.Rollback()
			utils.Fail(ctx, "失败", nil)
			return
		}

	}
	tran.Commit()

	utils.Success(ctx, "成功", "")

}

func ShippingOrders(ctx *gin.Context) {
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
	var Services service.SOrdersShipping
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	fmt.Println(fmt.Sprintf("%+v", req))
	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	req.UserId = user.Id

	result, count := Services.GetAll(req)
	data := map[string]any{
		"result": result,
		"count":  count,
	}
	language, _ := ctx.Get("Language")
	var orderitemserver service.SOrderItemsShipping
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
		result[k].Chindren = orderitem

	}
	//if req.OrderId != 0 && len(data) != 0 {
	//	data = data[0]
	//}
	utils.Success(ctx, "成功", data)
	return
}
func ShippingOrdersConfirm(ctx *gin.Context) {
	var req request.GetById
	if err := ctx.ShouldBind(&req); err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	user_id, _ := ctx.Get("user_id")
	var user model.User
	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}

	row := global.SHOP_DB.Model(model.OrdersShipping{}).Where("user_id=? and status=2", user.Id).Updates(map[string]interface{}{"status": 3, "finish_time": time.Now()}).RowsAffected

	if row == 0 {
		utils.Success(ctx, "失败", nil)
		return
	}
	//if req.OrderId != 0 && len(data) != 0 {
	//	data = data[0]
	//}
	utils.Success(ctx, "成功", nil)
	return
}
