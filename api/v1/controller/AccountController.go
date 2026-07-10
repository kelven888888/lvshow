package controller

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/global"
	"ginshop.com/utils"
	"ginshop.com/utils/Paginate"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Login 登录

func ChangePwd(ctx *gin.Context) {
	var params struct {
		ConfirmPassword string `json:"confirm_password" form:"confirm_password" `
		Password        string `json:"password" form:"password"  `
		OldPassword     string `json:"old_password" form:"old_password"  `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	if params.Password != params.ConfirmPassword {
		utils.Fail(ctx, "密码与确认密码不一致", nil)
		return
	}
	if !utils.IsValidPasswd(params.Password) {
		utils.Fail(ctx, "密码必须8到16个字符,包含大小写数字及字母", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.OldPassword))
	if err != nil {
		utils.Fail(ctx, "旧密码错误", nil)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(ctx, "加密错误", nil)
		return
	}
	user.Password = string(hashPassword)
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
	return
}

func ChangeTradePwd(ctx *gin.Context) {
	var params struct {
		ConfirmPassword string `json:"confirm_trade_password" form:"confirm_trade_password" `
		Password        string `json:"trade_password" form:"trade_password"  `
		OldPassword     string `json:"old_trade_password" form:"old_trade_password"  `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	if params.Password != params.ConfirmPassword {
		utils.Fail(ctx, "密码与确认密码不一致", nil)
		return
	}
	if !utils.IsValidTradePasswd(params.Password) {
		utils.Fail(ctx, "交易密码必须6位纯数字", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.TradePassword), []byte(params.OldPassword))
	if err != nil {
		utils.Fail(ctx, "旧密码错误", nil)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(ctx, "加密错误", nil)
		return
	}
	user.TradePassword = string(hashPassword)
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
	return
}
func Findtradepwd(ctx *gin.Context) {
	var params struct {
		ConfirmPassword string `json:"confirm_password" form:"confirm_password" `
		Password        string `json:"password" form:"password"  `
		Captcha         string `json:"captcha" form:"captcha"  `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	if params.Password != params.ConfirmPassword {
		utils.Fail(ctx, "密码与确认密码不一致", nil)
		return
	}
	if !utils.IsValidTradePasswd(params.Password) {
		utils.Fail(ctx, "交易密码必须6位纯数字", nil)
		return
	}
	if params.Captcha == "" {
		utils.Fail(ctx, "验证码不能为空", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB

	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	key := fmt.Sprintf("code_%s", user.Username)
	result, _ := global.SHOP_REDIS.Get(ctx, key).Result()
	if result == "" {
		utils.Fail(ctx, "验证码已过期", nil)
		return
	}
	if params.Captcha != result {
		utils.Fail(ctx, "验证码错误", nil)
		return
	}
	err = DB.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(ctx, "加密错误", nil)
		return
	}
	user.TradePassword = string(hashPassword)
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	global.SHOP_REDIS.Del(ctx, key)
	utils.Success(ctx, "成功", nil)
	return
}
func UserInfo(ctx *gin.Context) {
	var user model.User
	userid, _ := ctx.Get("user_id")

	uid := userid.(string)
	err := global.SHOP_DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	var funds model.AccountFunds
	err = global.SHOP_DB.Where("uid = ?", uid).Find(&funds).Error
	if err != nil {
		utils.Fail(ctx, "资金账号不存在", nil)
		return
	}
	var sum float64
	//err = global.SHOP_DB.Model(model.OrderDeal{}).Select("COALESCE(SUM(profit), 0)").Where("quan_account_id=?", v.id).Scan(&sum).Error
	err = global.SHOP_DB.Model(model.AccountFundsLog{}).Select(" ROUND(COALESCE(SUM(amount),0),2)").Where("uid=? and log_type=9", user.Id).Scan(&sum).Error
	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	var nextlevel model.MemberLevel
	var nowlevel model.MemberLevel
	global.SHOP_DB.Where("level=?", *user.Level+1).Find(&nextlevel)
	global.SHOP_DB.Where("level=?", *user.Level).Find(&nowlevel)
	nextexpleft := nextlevel.Exp.Sub(user.Exp)
	var BlindNum int64
	datenow := time.Now().Format("2006-01-02")
	err = global.SHOP_DB.Model(model.GameLottyRecord{}).Where("user_id=? and created_at>?", user.Id, fmt.Sprintf("%s 00:00:00", datenow)).Count(&BlindNum).Error

	response := make(map[string]any)
	response["Username"] = user.Username
	response["BlindNumToday"] = BlindNum
	response["DayBlindNumLimit"] = nowlevel.DayBlindNumLimit

	response["Id"] = user.Id + 80000000
	response["InviteCode"] = user.InviteCode
	response["IsAuth"] = user.IsAuth
	response["Email"] = user.Email
	response["Phone"] = user.Phone
	response["Level"] = user.Level
	response["Nickname"] = user.Nickname
	response["Avatar"] = user.Avatar
	response["InviteCount"] = user.InviteCount
	AvaFunds, _ := funds.AvaFunds.Float64()
	response["AvaFunds"] = AvaFunds
	Points, _ := funds.Points.Float64()
	response["Point"] = Points
	Exp, _ := user.Exp.Float64()
	response["Exp"] = Exp
	response["BlindBoxNum"] = user.BlindBoxNum
	response["TotalCommistion"] = sum
	response["nextexpleft"] = nextexpleft

	utils.Success(ctx, "成功", response)

}
func ChangeNickName(ctx *gin.Context) {
	var params struct {
		Nickname string `json:"nickname" form:"nickname" `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if params.Nickname == "" {
		utils.Fail(ctx, "请输入昵称", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	user.Nickname = params.Nickname
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
	return
}
func ChangeAvatar(ctx *gin.Context) {
	var params struct {
		Avatar string `json:"avatar" form:"avatar" `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if params.Avatar == "" {
		utils.Fail(ctx, "请上传头像", nil)
		return
	}
	ok, _ := utils.IsNetworkImage(params.Avatar)
	if !ok {
		utils.Fail(ctx, "请上传头像", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")

	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	user.Avatar = params.Avatar
	resErr := DB.Save(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
	return
}
func GetWallet(ctx *gin.Context) {
	var params struct {
		BusinessType string `json:"business_type" form:"business_type" `
		Types        string `json:"type" form:"type" `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	var WalletChain []*model.WalletChain

	if params.BusinessType == "1" {

		//err := global.SHOP_DB.Where("pid = ? and statuswithdraw=1 and type=?", 0, params.Types).Preload("SubCategory").Find(&WalletChain).Error
		err := global.SHOP_DB.Where("statuswithdraw=1 and type=?", params.Types).Find(&WalletChain).Error
		if err != nil {
			utils.Fail(ctx, "失败", nil)
			return
		}
	}
	if params.BusinessType == "2" {
		//	err := global.SHOP_DB.Where("pid = ? and statusrecharge=1 and type=?", 0, params.Types).Preload("SubCategory").Find(&WalletChain).Error
		err := global.SHOP_DB.Where(" statusrecharge=1 and type=?", params.Types).Find(&WalletChain).Error

		if err != nil {
			utils.Fail(ctx, "失败", nil)
			return
		}
	}
	var res model.WalletChain
	language, _ := ctx.Get("Language")
	restotal := res.ToTree(WalletChain, language.(string))
	//response := make(map[string]any)
	//response["Username"] = user.Username
	//
	//response["Id"] = user.Id + 30000000
	//response["InviteCode"] = user.InviteCode
	//response["IsAuth"] = user.IsAuth
	//response["Email"] = user.Email
	//response["Phone"] = user.Phone
	//response["Level"] = user.Level
	//response["Nickname"] = user.Nickname
	//response["Avatar"] = user.Avatar
	//AvaFunds, _ := funds.AvaFunds.Float64()
	//response["AvaFunds"] = AvaFunds
	//Points, _ := funds.Points.Float64()
	//response["Point"] = Points
	utils.Success(ctx, "成功", restotal)

}
func Recharge(ctx *gin.Context) {

	var recharge model.FundRecharge
	if err := ctx.ShouldBind(&recharge); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	//fmt.Println(fmt.Sprintf("%+v", recharge))
	//if err := ctx.Request.ParseForm(); err != nil {
	//	ctx.JSON(400, gin.H{"error": "ParseForm failed: " + err.Error()})
	//	return
	//}
	//
	//// 2. 打印解析后的表单数据，确认数据是否存在
	//fmt.Println("Form Data:", ctx.Request.Form)
	//fmt.Println("PostForm Data:", ctx.Request.PostForm)
	if recharge.PathType == "" {
		utils.Fail(ctx, "类型不能为空", nil)
		return
	}

	var models model.WalletChain
	global.SHOP_DB.Where("label=? and type=?", recharge.PathType, recharge.Type).Find(&models)
	if models.Id == 0 {
		utils.Fail(ctx, "类型不支持", nil)
		return
	}

	if math.Abs(utils.DecimalToFloat(recharge.Amount)) < *models.MinRechargeAmount || recharge.Amount.LessThan(decimal.Zero) {

		parms := utils.Parms{Keys: "{amount}", Value: fmt.Sprintf("%0.f", *models.MinRechargeAmount)}

		//parms.Value = fmt.Sprintf("%0.f", *models.MinRechargeAmount)

		//parms := map[string]string{
		//	"{amount}": fmt.Sprintf("%0.f", *models.MinRechargeAmount),
		//}
		//myMap := make(map[string]string)
		utils.Fail(ctx, "数量小于最小金额{amount}", nil, parms)
		return
	}
	ok, _ := utils.IsNetworkImage(recharge.Pic)
	if !ok {
		utils.Fail(ctx, "请上传凭证", nil)
		return
	}
	var user model.User
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	if *user.IsAuth != 1 {
		utils.Fail(ctx, "请先实名", nil)
		return
	}

	var rechargepending model.FundRecharge
	global.SHOP_DB.Where("status=0 and username=?", user.Username).Find(&rechargepending)
	if rechargepending.Id > 0 {
		utils.Fail(ctx, "已有待审核的申请,请稍后再试", nil)
		return
	}
	var conf service.Config
	rate, err := conf.GetKeyValue("MLECHANGERATE")
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	f, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	recharge.Username = user.Username
	recharge.Status = 0
	now := time.Now()

	recharge.CreateTime = now
	recharge.UpdateTime = &now
	recharge.ExchangeRate = utils.Float64ToDecimal(f)
	//recharge.Type = 1
	err = global.SHOP_DB.Save(&recharge).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
}
func Rechargelist(ctx *gin.Context) {

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
	var Services service.FundsRecharge
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
	//var withdraw []model.FundRecharge
	//global.SHOP_DB.Where("username=?", user.Username).Order("id desc ").Find(&withdraw)

	utils.Success(ctx, "成功", data)
}
func Withdraw(ctx *gin.Context) {

	var withdraw model.UsdtWithdrawModel
	if err := ctx.ShouldBind(&withdraw); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if withdraw.PathType == "" {
		utils.Fail(ctx, "类型不能为空", nil)
		return
	}
	if withdraw.TradePassword == "" {
		utils.Fail(ctx, "交易密码不能为空", nil)
		return
	}
	if withdraw.WalletPath == "" {
		utils.Fail(ctx, "地址不能为空", nil)
		return
	}

	var models model.WalletChain
	global.SHOP_DB.Where("label=? and type=?", withdraw.PathType, withdraw.Type).Find(&models)
	if models.Id == 0 {
		utils.Fail(ctx, "类型不支持", nil)
		return
	}
	var user model.User
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	if *user.IsAuth != 1 {
		utils.Fail(ctx, "请先实名", nil)
		return
	}

	var withdrawpending model.UsdtWithdrawModel
	global.SHOP_DB.Where("status=0 and username=?", user.Username).Find(&withdrawpending)
	if withdrawpending.Id > 0 {
		utils.Fail(ctx, "已有待审核的申请,请稍后再试", nil)
		return
	}
	var sum float64
	//err = global.SHOP_DB.Model(model.OrderDeal{}).Select("COALESCE(SUM(profit), 0)").Where("quan_account_id=?", v.id).Scan(&sum).Error
	times := time.Now().Format("2006-01-02")
	err = global.SHOP_DB.Model(model.UsdtWithdrawModel{}).Select("COALESCE(SUM(CASE WHEN type=1 THEN amount*exchange_rate ELSE amount END),0)").Where("username=? and create_time>? and status!=2", user.Username, times).Scan(&sum).Error
	if err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	var conf service.Config
	rate, err := conf.GetKeyValue("MLECHANGERATE")
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	f, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	//exchange_rang:=utils.Float64ToDecimal(f)
	amount := decimal.Zero
	if withdraw.Type == 1 {
		amount = withdraw.Amount.Mul(utils.Float64ToDecimal(f))
	}
	if withdraw.Type == 2 {
		amount = withdraw.Amount
	}

	var levelmodel model.MemberLevel
	global.SHOP_DB.Where("level=?", user.Level).Find(&levelmodel)
	if levelmodel.Id == 0 {
		utils.Fail(ctx, "会员等级不存在", nil)
		return
	}
	amounttotal := utils.DecimalToFloat(amount) + sum
	if amounttotal > float64(levelmodel.DayWithdrawLimit) {
		utils.Fail(ctx, fmt.Sprintf("超过每日限额"), nil)
		return
	}

	//fmt.Println(user.TradePassword)
	//fmt.Println(withdraw.TradePassword)
	//hashtradePassword, err := bcrypt.GenerateFromPassword([]byte(withdraw.TradePassword), bcrypt.DefaultCost)
	//fmt.Println(string(hashtradePassword), "3333333333333333")
	//if err != nil {
	//	utils.Fail(ctx, "加密错误", nil)
	//	return
	//}

	err = bcrypt.CompareHashAndPassword([]byte(user.TradePassword), []byte(withdraw.TradePassword))
	if err != nil {
		utils.Fail(ctx, "交易密码错误", nil)
		return
	}
	//fmt.Println(withdraw.Amount, *models.MinWithdrawAmount, "ddddddddddddd")
	if math.Abs(utils.DecimalToFloat(withdraw.Amount)) < *models.MinWithdrawAmount || withdraw.Amount.LessThan(decimal.Zero) {

		parms := utils.Parms{Keys: "{amount}", Value: fmt.Sprintf("%0.f", *models.MinWithdrawAmount)}

		//parms.Value = fmt.Sprintf("%0.f", *models.MinRechargeAmount)

		//parms := map[string]string{
		//	"{amount}": fmt.Sprintf("%0.f", *models.MinRechargeAmount),
		//}
		//myMap := make(map[string]string)
		utils.Fail(ctx, "提币数量小于最小金额{amount}", nil, parms)
		return
	}

	tran := global.SHOP_DB.Begin()
	var fund model.AccountFunds
	err = global.SHOP_DB.Where("username = ?", user.Username).Find(&fund).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	if fund.AvaFunds.Sub(amount).LessThan(decimal.Zero) {

		utils.Fail(ctx, "余额不足", nil)
		return
	}
	//err = global.SHOP_DB.Model(model.AccountFunds{}).Where("id=? ", fund.Id).Updates(model.AccountFunds{
	//	AvaFunds: fund.AvaFunds.Sub(withdraw.Amount),
	//}).Error
	err = tran.Exec("update account_funds set ava_funds=ava_funds-? where id=?", amount, fund.Id).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	withdraw.Username = user.Username
	withdraw.Status = 0
	now := time.Now()

	withdraw.CreateTime = now
	withdraw.UpdateTime = now
	//withdraw.Type = 1
	withdraw.ExchangeRate = utils.Float64ToDecimal(f)
	withdraw.Fee = utils.Float64ToDecimal(*models.WithdrawFee)
	//withdraw.Amount = amount
	err = global.SHOP_DB.Save(&withdraw).Error
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	var saccountfundslog service.AccountFundsLog
	remark := ""
	if withdraw.Type == 1 {
		remark = fmt.Sprintf("%sU-%s(RM-%s)", "提币申请", withdraw.Amount.Round(4).StringFixed(4), amount.Round(4).StringFixed(4))
	}
	if withdraw.Type == 2 {
		remark = fmt.Sprintf("%sRM-%s", "提币申请", amount.Round(4).StringFixed(4))
	}
	saccountfundslog.Createlog(user.Username, amount.Neg(), utils.Withdrayapplytype, remark, 1)
	tran.Commit()

	utils.Success(ctx, "成功", nil)
}
func Withdrawlist(ctx *gin.Context) {

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
	var Services service.SUsdtWithdraw
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
	utils.Success(ctx, "成功", data)
}
func Walletaddr(ctx *gin.Context) {
	var walletadd model.WalletPath
	if err := ctx.ShouldBind(&walletadd); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if walletadd.PathType == "" {
		utils.Fail(ctx, "类型不能为空", nil)
		return
	}
	if walletadd.WalletPath == "" {
		utils.Fail(ctx, "地址不能为空", nil)
		return
	}
	if walletadd.WalletType == "" {
		utils.Fail(ctx, "链不能为空", nil)
		return
	}
	var models model.WalletChain
	global.SHOP_DB.Where("label=? and chain=?", walletadd.PathType, walletadd.WalletType).Find(&models)
	if models.Id == 0 {
		utils.Fail(ctx, "类型不支持", nil)
		return
	}
	var user model.User
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	walletadd.Username = user.Username

	var walletaddr model.WalletPath
	global.SHOP_DB.Where("username=? and path_type=?", walletadd.Username, walletadd.PathType).Find(&walletaddr)
	if walletaddr.Id > 1 {
		walletadd.UpdateTime = time.Now()
		err = global.SHOP_DB.Where("id=?", walletaddr.Id).Updates(&walletadd).Error
	} else {
		walletadd.CreateTime = time.Now()
		err = global.SHOP_DB.Create(&walletadd).Error
	}

	if err != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)

}
func Walletaddrlist(ctx *gin.Context) {

	var user model.User
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}

	var walletaddr []model.WalletPath
	global.SHOP_DB.Where("username=?", user.Username).Order("id desc").Find(&walletaddr)

	if err != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", walletaddr)

}
func Authsubmit(ctx *gin.Context) {
	var authapply model.MTradeAccountApply
	if err := ctx.ShouldBind(&authapply); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	ok, _ := utils.IsNetworkImage(authapply.CarImg)
	if !ok {
		utils.Fail(ctx, "请上传证件正面", nil)
		return
	}
	ok, _ = utils.IsNetworkImage(authapply.CarImg2)
	if !ok {
		utils.Fail(ctx, "请上传证件反面", nil)
		return
	}
	if authapply.Name == "" {
		utils.Fail(ctx, "名字不能为空", nil)
		return
	}
	if authapply.CarId == "" {
		utils.Fail(ctx, "身份证号码不能为空", nil)
		return
	}
	var user model.User

	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	if *user.IsAuth == 1 {
		utils.Fail(ctx, "你已审核通过,不必提交", nil)
		return
	}

	global.SHOP_DB.Where("username=?", user.Username).Order("id desc").Limit(1).Find(&authapply)

	if authapply.Id > 0 {
		if *authapply.Status == 0 {
			utils.Fail(ctx, "你已有待审核的申请,请稍后再试", nil)
			return

		}

	}

	var num = 0
	var status = &num
	authapply.Status = status
	now := model.LocalTime(time.Now())
	if authapply.Id == 0 {
		authapply.CreateTime = &now
	} else {
		authapply.UpdateTime = &now
	}
	authapply.Username = user.Username

	err = global.SHOP_DB.Save(&authapply).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	if err != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)

}
func Authlist(ctx *gin.Context) {
	var authapply model.MTradeAccountApply

	var user model.User

	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	err := global.SHOP_DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}

	global.SHOP_DB.Where("username=?", user.Username).Order("id desc").First(&authapply)
	//create := authapply.CreateTime.Format("2006-01-02 15:04:05")
	//const customLayout = "2006-01-02 15:04:05" // 注意这里的数字要和Go的时间常量对应
	//t2, err := time.Parse(customLayout, create)
	//
	//authapply.CreateTime = t2

	utils.Success(ctx, "成功", authapply)

}
func Fundsrecord(ctx *gin.Context) {
	var params model.AccountFundsLog
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	query := global.SHOP_DB.Model(model.AccountFundsLog{})

	var funds []model.AccountFundsLog
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	query = query.Where("uid = ?", uid)
	if params.DataType == 1 {
		query = query.Where("amount>0")
	}
	if params.DataType == 2 {
		query = query.Where("amount<0")
	}
	if params.MoneyType != 0 {
		query = query.Where("money_type=?", params.MoneyType)
	}
	if params.LogType != 0 {
		query = query.Where("log_type=?", params.LogType)
	}
	var count int64 = 0
	pageUp := strconv.Itoa(params.Page)

	if err := query.Count(&count).Error; err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	query = query.Scopes(Paginate.Paginate(pageUp, global.SHOP_CONFIG.System.PageSize)).Order("id desc ")
	err := query.Find(&funds).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	language, _ := ctx.Get("Language")
	for k, _ := range funds {
		code := utils.Get_funtype_Key(funds[k].LogType)
		funds[k].LogTypes = utils.Languageresponse(code, language.(string))
	}

	utils.Success(ctx, "获取成功", gin.H{
		"count": count,

		"data": funds,
	})

}
func Setlanguage(ctx *gin.Context) {
	var params struct {
		Language string `json:"language" form:"language" `
	}
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	if params.Language == "" {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	allowlanguage := strings.Split(global.SHOP_CONFIG.System.Language_Array, ",")
	if !utils.InArray(params.Language, allowlanguage) {
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	err := DB.Where("id = ?", uid).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	user.Language = params.Language
	resErr := DB.Updates(&user).Error
	if resErr != nil {
		utils.Fail(ctx, "失败", nil)
		return
	}
	utils.Success(ctx, "成功", nil)
	return
}
func Memberlevel(ctx *gin.Context) {

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

	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	req.IsDisplay = 1

	var memeberserver service.Userlevel

	memberlevel, number := memeberserver.GetAll(req)

	language, _ := ctx.Get("Language")
	for k, v := range memberlevel {

		memberlevel[k].Title = utils.Languagebycode(language.(string), v.Title)
	}
	var user model.User
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")

	uid := userid.(string)
	err = DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	var sum float64
	//err = global.SHOP_DB.Model(model.OrderDeal{}).Select("COALESCE(SUM(profit), 0)").Where("quan_account_id=?", v.id).Scan(&sum).Error
	err = global.SHOP_DB.Model(model.AccountFundsLog{}).Select(" ROUND(COALESCE(SUM(amount),0),2)").Where("uid=? and log_type=9", user.Id).Scan(&sum).Error
	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	var nextlevel model.MemberLevel
	global.SHOP_DB.Where("level=?", *user.Level+1).Find(&nextlevel)
	nextexpleft := nextlevel.Exp.Sub(user.Exp)
	data := map[string]any{
		"result":      memberlevel,
		"count":       number,
		"level":       user.Level,
		"nowexp":      user.Exp,
		"nextexp":     nextlevel.Exp,
		"nextlevel":   *user.Level + 1,
		"nextexpleft": nextexpleft,
	}
	utils.Success(ctx, "成功", data)
	return
}
func Getcoupon(ctx *gin.Context) {

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

	userid, _ := ctx.Get("user_id")

	uid := userid.(string)

	size, _ := strconv.Atoi(global.SHOP_CONFIG.System.PageSize)

	req.Count = true
	req.Limit = size
	req.Offset = (p - 1) * size
	req.UserId, _ = strconv.Atoi(uid)
	//two := 2
	//req.Status = &two

	var services service.STbCouponList
	if req.Status != nil {
	}
	if req.Status != nil {
		if *req.Status == 2 {
			req.Expdate = time.Now()
		}
	}

	memberlevel, number := services.GetAll(req)

	language, _ := ctx.Get("Language")
	for k, v := range memberlevel {

		memberlevel[k].CouponTitle = utils.Languagebycode(language.(string), v.CouponTitle)
	}

	data := map[string]any{

		"count": number,
		"data":  memberlevel,
	}
	utils.Success(ctx, "成功", data)
	return
}
func Getaddressselect(ctx *gin.Context) {

	var states []model.MyStates
	global.SHOP_DB.Find(&states)

	for key, value := range states {
		var areas []model.MyAreas
		global.SHOP_DB.Where("state_code=?", value.Code).Find(&areas)

		for k, v := range areas {
			var postcode []model.MyPostcodes
			global.SHOP_DB.Where("area_code=?", v.Code).Find(&postcode)
			areas[k].MyPostcodes = postcode
		}
		states[key].MyAreas = areas

	}

	utils.Success(ctx, "成功", states)
	return
}
func Doshippingaddr(ctx *gin.Context) {

	var req model.ShippingAddresses
	err := ctx.ShouldBind(&req)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	var user model.User
	err = DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}
	fmt.Println(req.ReceiverName == "", req.AddressLine1 == "", req.Area == "", req.City == "", req.ReceiverPhone == "")
	if req.ReceiverName == "" || req.AddressLine1 == "" || req.Area == "" || req.City == "" || req.ReceiverPhone == "" {
		utils.Fail(ctx, "必填项未填", nil)
		return
	}

	if req.Id > 0 {

		now := model.LocalTime(time.Now())
		req.Model.UpdatedAt = &now
		roweff := global.SHOP_DB.Where("username=?", user.Username).Updates(&req).RowsAffected
		if roweff == 0 {

			utils.Fail(ctx, "失败", nil)
			return
		}

		utils.Success(ctx, "成功", nil)
		return
	} else {
		var sum int64
		global.SHOP_DB.Model(model.ShippingAddresses{}).Where("username=?", user.Username).Count(&sum)
		if sum >= 3 {
			utils.Fail(ctx, "地址最多保存三个", nil)
			return
		}
		now := model.LocalTime(time.Now())
		req.Model.CreatedAt = &now
		req.Username = user.Username

		err = global.SHOP_DB.Save(&req).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			utils.Fail(ctx, "失败", nil)
			return

		}
		utils.Success(ctx, "成功", nil)
		return
	}
}
func Handleshippingaddr(ctx *gin.Context) {

	type parms struct {
		Id        int `json:"id" form:"id" `
		IsDelete  int `json:"is_delete" form:"is_delete" `
		IsDefault int `json:"is_default" form:"is_default" `
	}
	var parm parms

	err := ctx.ShouldBind(&parm)
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "参数错误", nil)
		return
	}
	DB := global.SHOP_DB
	userid, _ := ctx.Get("user_id")
	fmt.Println(userid)
	uid := userid.(string)
	var user model.User
	err = DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在", nil)
		return
	}

	var address model.ShippingAddresses
	global.SHOP_DB.Where("id=? and username=?", parm.Id, user.Username).Find(&address)
	if address.Id == 0 {
		utils.Fail(ctx, "参数错误", nil)
		return
	}

	if parm.IsDelete == 1 {
		global.SHOP_DB.Where("id=? and username=?", parm.Id, user.Username).Delete(&address)
	}
	if parm.IsDefault == 1 {
		global.SHOP_DB.Model(model.ShippingAddresses{}).Where("username=?", user.Username).Updates(map[string]interface{}{"is_default": 0})
		global.SHOP_DB.Model(model.ShippingAddresses{}).Where("id=?", parm.Id).Updates(model.ShippingAddresses{
			IsDefault: 1,
		})
	}

	utils.Success(ctx, "成功", nil)
	return

}

func Myshippingaddr(ctx *gin.Context) {

	var params model.ShippingAddresses
	if err := ctx.ShouldBind(&params); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	query := global.SHOP_DB.Model(model.ShippingAddresses{})

	var addr []model.ShippingAddresses
	user_id, _ := ctx.Get("user_id")

	var user model.User
	global.SHOP_DB.Where("id=?", user_id).Find(&user)

	query = query.Where("username = ?", user.Username)

	var count int64 = 0
	pageUp := strconv.Itoa(params.Page)

	if err := query.Count(&count).Error; err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	query = query.Scopes(Paginate.Paginate(pageUp, global.SHOP_CONFIG.System.PageSize)).Order("id desc ")
	err := query.Find(&addr).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	utils.Success(ctx, "获取成功", gin.H{
		"count": count,
		"data":  addr,
	})

}
func Commissionrecord(ctx *gin.Context) {
	type parm struct {
		Page    int    `json:"page" form:"page" `
		EndTime string `json:"endtime" form:"endtime" `
	}
	var parms parm

	if err := ctx.ShouldBind(&parms); err != nil {
		utils.Fail(ctx, "参数错误："+err.Error(), nil)
		return
	}

	query := global.SHOP_DB.Model(model.CommissionRecord{})

	var funds []model.CommissionRecord
	user_id, _ := ctx.Get("user_id")

	uid := user_id.(string)
	query = query.Where("user_id = ?", uid)
	if parms.EndTime != "" {
		query = query.Where("created_at>?", fmt.Sprintf("%s 00:00:00", parms.EndTime))
	}
	var sum float64
	//err = global.SHOP_DB.Model(model.OrderDeal{}).Select("COALESCE(SUM(profit), 0)").Where("quan_account_id=?", v.id).Scan(&sum).Error
	err := global.SHOP_DB.Model(model.CommissionRecord{}).Select(" ROUND(COALESCE(SUM(amount),0),2)").Where("user_id=? ", uid).Scan(&sum).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())

		return
	}
	var count int64 = 0
	pageUp := strconv.Itoa(parms.Page)

	if err := query.Count(&count).Error; err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}

	query = query.Scopes(Paginate.Paginate(pageUp, global.SHOP_CONFIG.System.PageSize)).Order("id desc ")
	err = query.Find(&funds).Error
	if err != nil {
		global.SHOP_LOG.Error(err.Error())
		utils.Fail(ctx, "失败", nil)
		return
	}
	language, _ := ctx.Get("Language")
	for k, _ := range funds {
		code := utils.Get_funtype_Key(funds[k].LogType)
		funds[k].LogTypes = utils.Languageresponse(code, language.(string))
		var user model.User
		global.SHOP_DB.Where("id=?", funds[k].FromUserId).Find(&user)

		funds[k].FromUsername = utils.MaskString(funds[k].FromUsername, 2, 4)
		funds[k].Avatar = user.Avatar
		global.SHOP_DB.Where("id=?", funds[k].FromUserId).Find(&user)
		funds[k].MemberLevel = *user.Level
	}
	data := map[string]any{
		"count":           count,
		"totalcommission": sum,

		"data": funds,
	}

	utils.Success(ctx, "获取成功", data)

}
