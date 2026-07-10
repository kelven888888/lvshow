package controller

import (
	"encoding/json"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"log"
	"net/http"
	"sync"
	"time"

	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"strconv"
)

func GetGameProductbyid(ctx *gin.Context) {
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
	var Services service.SPlaysetting
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size

	result, count := Services.GetAllProduct(req)

	language, _ := ctx.Get("Language")
	for k, v := range result {
		result[k].CreatedAt = v.CreatedAt
		result[k].GoodsName = utils.Languagebycode(language.(string), result[k].GoodsName)
		result[k].GoodsCover = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, v.GoodsCover)
	}
	var Servicess service.SPlaysetting

	var reqs request.GetById
	reqs.ID = uint(req.PlayId)
	play, _ := Servicess.GetByID(reqs)
	var sum int64
	global.SHOP_DB.Model(model.GameLottyRecord{}).Where("play_id=?", req.PlayId).Count(&sum)
	play.SumPlay = sum
	var totalplay int64
	totalplay = 0

	user_id, exit := ctx.Get("user_id")
	play.Img = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, play.Img)
	play.Name = utils.Languagebycode(language.(string), play.Name)
	if exit {
		var user model.User
		uid := user_id.(string)
		err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "用户不存在", nil)
			return
		}
		global.SHOP_DB.Model(model.GameLottyRecord{}).Where("play_id=? and user_id=?", play.Id, user.Id).Count(&totalplay)
		lianjikey1 := fmt.Sprintf("lianji_%d_%d_1", user.Id, play.Id)
		combokey1 := fmt.Sprintf("combo_%d_%d_1", user.Id, play.Id)
		Combo, _ := global.SHOP_REDIS.Get(ctx, combokey1).Result()
		Dhit, _ := global.SHOP_REDIS.Get(ctx, lianjikey1).Result()
		c, _ := strconv.Atoi(Combo)
		d, _ := strconv.Atoi(Dhit)
		play.SCombo = &c
		play.SDhit = &d
		lianjikey2 := fmt.Sprintf("lianji_%d_%d_2", user.Id, play.Id)
		combokey2 := fmt.Sprintf("combo_%d_%d_2", user.Id, play.Id)
		DCombo, _ := global.SHOP_REDIS.Get(ctx, combokey2).Result()
		DDhit, _ := global.SHOP_REDIS.Get(ctx, lianjikey2).Result()
		c2, _ := strconv.Atoi(DCombo)
		d2, _ := strconv.Atoi(DDhit)
		play.DCombo = &c2
		play.DDhit = &d2
		play.TotalPlay = totalplay
	}
	data := map[string]any{
		"result": result,
		"count":  count,
		"play":   play,
	}
	utils.Success(ctx, "成功", data)
	return

}
func GetGame(ctx *gin.Context) {
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
	var Services service.SPlaysetting
	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	status := 1
	req.Status = &status

	result, count := Services.GetAll(req)
	data := map[string]any{
		"result": result,
		"count":  count,
	}
	language, _ := ctx.Get("Language")
	user_id, exit := ctx.Get("user_id")

	for k, v := range result {
		result[k].CreatedAt = v.CreatedAt
		result[k].Img = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, v.Img)
		result[k].Name = utils.Languagebycode(language.(string), result[k].Name)
		var sum int64
		global.SHOP_DB.Model(model.GameLottyRecord{}).Where("play_id=?", v.Id).Count(&sum)
		result[k].SumPlay = sum
		var totalplay int64
		totalplay = 0
		if exit {
			var user model.User
			uid := user_id.(string)
			err = global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
			if err != nil {
				global.SHOP_LOG.Error(err.Error())
				utils.Fail(ctx, "用户不存在", nil)
				return
			}
			global.SHOP_DB.Model(model.GameLottyRecord{}).Where("play_id=? and user_id=?", v.Id, user.Id).Count(&totalplay)
			lianjikey1 := fmt.Sprintf("lianji_%d_%d_1", user.Id, v.Id)
			combokey1 := fmt.Sprintf("combo_%d_%d_1", user.Id, v.Id)
			Combo, _ := global.SHOP_REDIS.Get(ctx, combokey1).Result()
			Dhit, _ := global.SHOP_REDIS.Get(ctx, lianjikey1).Result()
			c, _ := strconv.Atoi(Combo)
			d, _ := strconv.Atoi(Dhit)
			result[k].SCombo = &c
			result[k].SDhit = &d
			lianjikey2 := fmt.Sprintf("lianji_%d_%d_2", user.Id, v.Id)
			combokey2 := fmt.Sprintf("combo_%d_%d_2", user.Id, v.Id)
			DCombo, _ := global.SHOP_REDIS.Get(ctx, combokey2).Result()
			DDhit, _ := global.SHOP_REDIS.Get(ctx, lianjikey2).Result()
			c2, _ := strconv.Atoi(DCombo)
			d2, _ := strconv.Atoi(DDhit)
			result[k].DCombo = &c2
			result[k].DDhit = &d2
		}
		result[k].TotalPlay = totalplay

	}
	utils.Success(ctx, "成功", data)
	return

}
func LotteryRecord(ctx *gin.Context) {
	var req request.PageInfo
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
	var Services service.SGameLottyRecord
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
	for k, v := range result {
		result[k].CreatedAt = v.CreatedAt
		result[k].GoodsName = utils.Languagebycode(language.(string), result[k].GoodsName)
		result[k].PlayName = utils.Languagebycode(language.(string), result[k].PlayName)
		var good model.Goods
		global.SHOP_DB.Where("id=?", result[k].GoodsId).Find(&good)
		result[k].Remark = "success"
		result[k].Image = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, good.GoodsCover)

	}
	utils.Success(ctx, "成功", data)
	return

}

func Dolottery(ctx *gin.Context) {

	var req request.GetByLottyId
	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if req.LotteryNum > 50 {

		utils.Fail(ctx, "参数错误", nil)
		return
	}

	var models model.Playsetting
	global.SHOP_DB.Where("id=?", req.ID).Find(&models)
	if models.Id == 0 {
		utils.Fail(ctx, "玩法不存在", nil)
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
	var modelfund model.AccountFunds
	err = global.SHOP_DB.Where("username = ?", user.Username).Find(&modelfund).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	var couponlist model.TbCouponList
	discountmoney := decimal.Zero
	if req.CouponCode != "" {

		now := time.Now()
		err := global.SHOP_DB.Where("coupon_code=? and status=2 and exp_date>? and uid=?", req.CouponCode, now, user.Id).Find(&couponlist).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		if couponlist.Id == 0 {

			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		discountmoney = couponlist.Price
	}
	if req.CouponId != 0 {

		now := time.Now()
		err := global.SHOP_DB.Where("id=? and status=2 and exp_date>? and uid=?", req.CouponId, now, user.Id).Find(&couponlist).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		if couponlist.Id == 0 {

			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		req.CouponCode = couponlist.CouponCode
		discountmoney = couponlist.Price
	}
	coustmoney := decimal.NewFromUint64(req.LotteryNum).Mul(models.Price)
	if !coustmoney.LessThan(modelfund.AvaFunds.Sub(discountmoney)) {
		utils.Fail(ctx, fmt.Sprintf("资金不足"), nil)
		return
	}
	if discountmoney.GreaterThan(coustmoney) {
		utils.Fail(ctx, fmt.Sprintf("优惠券价值不能大于所需资金"), nil)
		return
	}
	lockkey := fmt.Sprintf("lotter_%d", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	var serrvice service.SGame
	var good model.Goods
	var goods []model.Goods
	wg := sync.WaitGroup{}
	//盲盒
	if models.Id == 6 {
		if *models.SingelStatus != 1 && req.LotteryType != 1 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if *models.DoubleStatus != 1 && req.LotteryType != 2 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if req.LotteryNum != 1 {
			utils.Fail(ctx, fmt.Sprintf("盲盒一次只能抽一次"), nil)
			return
		}
		if user.BlindBoxNum < req.LotteryNum {
			utils.Fail(ctx, fmt.Sprintf("盲盒次数不足"), nil)
			return
		}
		var levle model.MemberLevel
		global.SHOP_DB.Where("level=?", user.Level).Find(&levle)
		if levle.Id == 0 {
			utils.Fail(ctx, fmt.Sprintf("会员等级找不到"), nil)
			return
		}
		var sum int64
		datenow := time.Now().Format("2006-01-02")
		err := global.SHOP_DB.Model(model.GameLottyRecord{}).Where("user_id=? and created_at>? and play_id=6", user.Id, fmt.Sprintf("%s 00:00:00", datenow)).Count(&sum).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
		fmt.Println(sum, levle.DayBlindNumLimit)
		if sum >= int64(levle.DayBlindNumLimit) {
			utils.Fail(ctx, fmt.Sprintf("超过会员每日盲盒子次数"), nil)
			return
		}

	}

	//保底
	if models.Id == 1 || models.Id == 5 {
		if *models.SingelStatus != 1 && req.LotteryType != 1 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if *models.DoubleStatus != 1 && req.LotteryType != 2 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
	}
	lianjicount := 0
	combocount := 0
	combocounts := 0
	combo := 0

	for i := 0; i < int(req.LotteryNum); i++ {
		wg.Add(1)
		go func() {
			good, lianjicount, combocount, combocounts, err = serrvice.LottyBlindBox(models, req, user, lianjicount, combocount, 0, discountmoney)
			if combocount == 1 {
				combo = combo + 1
				combocount = 0
			}
			if err != nil {
				global.SHOP_LOG.Error(err.Error())

			}
			goods = append(goods, good)
			wg.Done()
		}()

	}
	wg.Wait()
	language, _ := ctx.Get("Language")
	goodsslice := make([]map[string]any, 0)

	for _, v := range goods {

		goodresult := map[string]any{}
		goodresult["reward_type"] = v.RewardType
		goodresult["goods_cover"] = fmt.Sprintf("%s%s", global.SHOP_CONFIG.System.WebApiURL, v.GoodsCover)
		goodresult["goods_name"] = utils.Languagebycode(language.(string), v.GoodsName)
		goodresult["money"] = v.UnitPrice
		goodsslice = append(goodsslice, goodresult)

	}

	if len(goodsslice) == 0 {
		utils.Fail(ctx, fmt.Sprintf("失败"), goodsslice)
		return
	}
	goodsslices := map[string]any{
		"ljcount":     lianjicount,
		"combo":       combo,
		"combocounts": combocounts,
		"data":        goodsslice,
	}
	couponlist.Status = 1
	timenow := model.LocalTime(time.Now())
	couponlist.UsedTime = &timenow
	if req.CouponCode != "" {
		global.SHOP_DB.Updates(&couponlist)
	}

	utils.Success(ctx, fmt.Sprintf("成功"), goodsslices)
	return

}

func Dolotterystr(ctx *gin.Context) {
	var req request.GetByLottyId
	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if req.LotteryNum > 50 {

		utils.Fail(ctx, "参数错误", nil)
		return
	}

	var models model.Playsetting
	global.SHOP_DB.Where("id=?", req.ID).Find(&models)
	if models.Id == 0 {
		utils.Fail(ctx, "玩法不存在", nil)
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
	var modelfund model.AccountFunds
	err = global.SHOP_DB.Where("username = ?", user.Username).Find(&modelfund).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	var couponlist model.TbCouponList
	discountmoney := decimal.Zero
	if req.CouponCode != "" {

		now := time.Now()
		err := global.SHOP_DB.Where("coupon_code=? and status=2 and exp_date>? and uid=?", req.CouponCode, now, user.Id).Find(&couponlist).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		if couponlist.Id == 0 {

			utils.Fail(ctx, "优惠券无效", nil)
			return
		}
		discountmoney = couponlist.Price
	}
	coustmoney := decimal.NewFromUint64(req.LotteryNum).Mul(models.Price)
	if !coustmoney.LessThan(modelfund.AvaFunds.Sub(discountmoney)) {
		utils.Fail(ctx, fmt.Sprintf("资金不足"), nil)
		return
	}
	if discountmoney.GreaterThan(coustmoney) {
		utils.Fail(ctx, fmt.Sprintf("优惠券价值大于所需资金"), nil)
		return
	}
	lockkey := fmt.Sprintf("lotter_%d", user.Id)
	if !utils.AcquireLock(global.SHOP_REDIS, lockkey, time.Second*10) {
		global.SHOP_LOG.Log(0, fmt.Sprintf("获取不到锁：%s"))
		utils.Fail(ctx, "操作频繁,请稍后再试", nil)
		return

	}
	defer utils.ReleaseLock(global.SHOP_REDIS, lockkey)
	var serrvice service.SGame
	var good model.Goods
	var goods []model.Goods
	//wg := sync.WaitGroup{}
	//盲盒
	if models.Id == 6 {
		if *models.SingelStatus != 1 && req.LotteryType != 1 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if *models.DoubleStatus != 1 && req.LotteryType != 2 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if req.LotteryNum != 1 {
			utils.Fail(ctx, fmt.Sprintf("盲盒一次只能抽一次"), nil)
			return
		}
		if user.BlindBoxNum < req.LotteryNum {
			utils.Fail(ctx, fmt.Sprintf("盲盒次数不足"), nil)
			return
		}
		var levle model.MemberLevel
		global.SHOP_DB.Where("level=?", user.Level).Find(&levle)
		if levle.Id == 0 {
			utils.Fail(ctx, fmt.Sprintf("会员等级找不到"), nil)
			return
		}
		var sum int64
		datenow := time.Now().Format("2006-01-02")
		err := global.SHOP_DB.Model(model.GameLottyRecord{}).Where("user_id=? and created_at>? and play_id=6", user.Id, fmt.Sprintf("%s 00:00:00", datenow)).Count(&sum).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return
		}
		fmt.Println(sum, levle.DayBlindNumLimit)
		if sum >= int64(levle.DayBlindNumLimit) {
			utils.Fail(ctx, fmt.Sprintf("超过会员每日盲盒子次数"), nil)
			return
		}

	}

	//保底
	if models.Id == 1 || models.Id == 5 {
		if *models.SingelStatus != 1 && req.LotteryType != 1 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
		if *models.DoubleStatus != 1 && req.LotteryType != 2 {
			utils.Fail(ctx, fmt.Sprintf("尚未开放"), nil)
			return
		}
	}
	lianjicount := 0
	combocount := 0
	combocounts := 0
	combo := 0
	wg := sync.WaitGroup{}
	chgood := make(chan model.Goods)
	done := make(chan int, 1)
	go func() {
		for i := 0; i < int(req.LotteryNum); i++ {
			wg.Add(1)
			go func() {
				good, lianjicount, combocount, combocounts, err = serrvice.LottyBlindBox(models, req, user, lianjicount, combocount, 0, discountmoney)
				if combocount == 1 {
					combo = combo + 1
					combocount = 0
				}
				if err != nil {
					global.SHOP_LOG.Error(err.Error())

				}
				goods = append(goods, good)
				chgood <- good
				wg.Done()

			}()

		}
		wg.Wait()
		done <- 1
		fmt.Println(combocounts)

	}()

	// 1. 设置必要的响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*") // 允许跨域

	// 2. 获取底层的 ResponseWriter 并断言为 Flusher
	w := ctx.Writer
	flusher, ok := w.(http.Flusher)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 3. 监听客户端断开连接
	// 注意：在较新的 Go 版本中，CloseNotify 已废弃，推荐使用 c.Request.Context().Done()
	clientGone := ctx.Request.Context().Done()

	// 4. 创建定时器模拟实时数据推送
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 5. 循环推送数据
	for {
		select {
		case <-clientGone:
			log.Println("Client disconnected")
			return
		case t := <-chgood:
			// 构造 SSE 格式数据
			// 格式要求: data: <message>\n\n
			str, _ := json.Marshal(t)
			msg := fmt.Sprintf("data: Current Time: %s\n\n", str)
			fmt.Println(str)

			_, err := w.Write([]byte(msg))
			if err != nil {
				global.SHOP_LOG.Error(err.Error())
				return
			}

			// 关键步骤：刷新缓冲区，立即发送数据
			flusher.Flush()
		case <-done:
			global.SHOP_LOG.Error("完结完结完结完结完结完结完结完结完结完结完结完结完结完结")
			return
		}
	}

}
