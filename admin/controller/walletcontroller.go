package controller

import (
	"fmt"
	"strconv"
	"strings"

	"ginshop.com/admin/model"
	"ginshop.com/admin/service"
	"ginshop.com/crondtab/wallet_api/util"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/gin-gonic/gin"
)

type WalletCtr struct {
	Services service.WalletPath
	BaseController
}
type Rechargresp struct {
	Pid         int64  `json:"pid"  mapstructure:"pid"`
	Cid         int64  `json:"cid"  mapstructure:"cid"`
	ChainID     string `json:"chain_id" mapstructure:"chain_id"`
	TokenID     string `json:"token_id" mapstructure:"token_id"`
	Currency    string `json:"currency" mapstructure:"currency"`
	Amount      string `json:"amount" `
	Address     string `json:"address"`
	Status      string `json:"status"`
	Txid        string `json:"txid"`
	BlockHeight string `json:"block_height" mapstructure:"block_height"`
	BlockTime   string `json:"block_time" mapstructure:"block_time"`
	Nonce       string `json:"nonce"`
	Timestamp   int64  `json:"timestamp"`
	Sign        string `json:"sign"`
}

func (this *WalletCtr) RechargeCallBack(ctx *gin.Context) {
	var resp Rechargresp
	if err := ctx.ShouldBindJSON(&resp); err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	//验签
	data := map[string]interface{}{}
	for k, v := range utils.StructToMap(resp) {
		data[strings.ToLower(k)] = v
	}
	fmt.Println("RechargeCallBack", data)
	signs, _ := util.DoSign(data, global.SHOP_CONFIG.Wallet.Appkey)
	fmt.Println(signs)
	sign, err := util.VerifySign(data, global.SHOP_CONFIG.Wallet.Appkey, resp.Sign)
	if err != nil || !sign {
		global.SHOP_LOG.Log(0, fmt.Sprintf("充值回调验签失败%s", err.Error()))
		this.Success(ctx, "充值回调验签失败")
		return
	}

	// TODO: 处理充值回调
	// 1. 验证签名
	// 2. 处理转账
	// 3. 记录转账信息
	var rechargelog model.FundRecharge
	err = global.SHOP_DB.Where("hash=?", resp.Txid).Find(&rechargelog).Error
	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	if rechargelog.Id > 0 {
		this.Success(ctx, "充值记录已存在")
		global.SHOP_LOG.Log(0, "充值记录已存在")
		return
	}
	currentcymap := map[string]string{
		//"BTC":  "0",
		"60": "ETH",
		"0":  "BTC",
		"0xdac17f958d2ee523a2206206994597c13d831ec7":   "USDT Ethereum(ERC20)",
		"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48":   "USDC Ethereum(ERC20)",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t":           "USDT TRX(TRC20)",
		"TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8":           "USDC TRX(TRC20)",
		"Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB": "USDT Solana",
		"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v": "USDC Solana",
		"0x833589fcd6edb6e08f4c7c32d4f71b54bda02913":   "USDC Base",
	}
	wallet_type, ok := currentcymap[resp.TokenID]
	if !ok {
		global.SHOP_LOG.Log(0, "未找到该TokenID对应的钱包类型")
		return
	}
	price := 1.0
	var rechargeamount float64
	if wallet_type == "BTC" || wallet_type == "ETH" {
		price, err = utils.Get_crypto_current_price(wallet_type)
		if err != nil {
			global.SHOP_LOG.Log(0, err.Error())
			return
		}
		amount, err := strconv.ParseFloat(resp.Amount, 64)
		if err != nil {
			global.SHOP_LOG.Log(0, err.Error())
			return
		}
		tranamount := int(amount * price)
		rechargeamount = float64(tranamount)
	} else {
		amount, err := strconv.ParseFloat(resp.Amount, 64)
		if err != nil {
			global.SHOP_LOG.Log(0, err.Error())
			return
		}
		rechargeamount = amount
	}

	var userpath model.WalletPath
	err = global.SHOP_DB.Where("wallet_path=?", resp.Address).Find(&userpath).Error
	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	if userpath.Id == 0 {
		global.SHOP_LOG.Log(0, fmt.Sprintf("充值地址不存在%s", resp.Address))

		return
	}
	if rechargeamount < 1 {
		global.SHOP_LOG.Log(0, fmt.Sprintf("充值小于1不上分%s", resp.Address))
		return
	}
	username := userpath.Username
	timenow := utils.Us_datatimecon()
	tran := global.SHOP_DB.Begin()
	rechargelog.Address = resp.Address
	rechargelog.Hash = resp.Txid
	rechargelog.Amount = utils.Float64ToDecimal(rechargeamount)
	rechargelog.Remarks = "success"
	rechargelog.PathType = wallet_type
	rechargelog.CreateTime = timenow
	//rechargelog.UpdateTime = timenow
	rechargelog.Username = username

	err = tran.Save(&rechargelog).Error
	if err != nil {
		global.SHOP_LOG.Log(0, fmt.Sprintf("错误%s", err.Error()))
		tran.Rollback()
		return
	}

	var accountfunds model.AccountFunds
	err = global.SHOP_DB.Where("username=?", username).First(&accountfunds).Error
	if err != nil {
		global.SHOP_LOG.Log(0, fmt.Sprintf("账号错误%s", err.Error()))
		tran.Rollback()
		return
	}
	//fundold := accountfunds.AvaFunds
	//
	//accountfunds.AvaFunds = utils.Float64ToDecimal(rechargeamount).Add(accountfunds.AvaFunds)

	err = tran.Exec("UPDATE  account_funds set ava_funds=ava_funds+? where id=?", rechargeamount, accountfunds.Id).Error
	if err != nil {

		global.SHOP_LOG.Log(0, fmt.Sprintf("更新accountfunds错误%s", err.Error()))
		tran.Rollback()
		return
	}
	var saccountfundslog service.AccountFundsLog
	remark := fmt.Sprintf("%s%s%s/%.2f 汇率%f", "充值", wallet_type, resp.Amount, rechargeamount, price)
	err, _ = saccountfundslog.Createlog(username, utils.Float64ToDecimal(rechargeamount), utils.Rechargetype, remark, 1)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	var message service.AccountMsgServer
	var modelmsg model.AccountUserMessage
	modelmsg.Username = username
	modelmsg.CreateTime = timenow
	modelmsg.UpdateTime = timenow
	group := 2
	keys := utils.Get_Code_Key(5)
	parms := make(map[string]string)
	parms["{wallet_type}"] = wallet_type
	parms["{amount}"] = resp.Amount
	parms["{amount}"] = resp.Address
	err = message.Save(&modelmsg, group, group, keys, parms)
	if err != nil {
		tran.Rollback()
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	tran.Commit()
	this.Success(ctx, "success")

}

type Withresp struct {
	Pid          int64  `json:"pid"`
	Cid          int64  `json:"cid"`
	Address      string `json:"address"`
	ChainID      string `json:"chain_id" mapstructure:"chain_id"`
	TokenID      string `json:"token_id" mapstructure:"token_id"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	ThirdPartyID string `json:"third_party_id" mapstructure:"third_party_id"`
	Remark       string `json:"remark"`
	Status       int64  `json:"status"`
	Txid         string `json:"txid"`
	BlockHeight  string `json:"block_height" mapstructure:"block_height"`
	BlockTime    string `json:"block_time" mapstructure:"block_time"`
	Nonce        string `json:"nonce"`
	Timestamp    int64  `json:"timestamp"`
	Sign         string `json:"sign"`
}

func (this *WalletCtr) PayoutBack(ctx *gin.Context) {
	// TODO: 处理提现回调
	// 1. 验证签名
	// 2. 处理转账
	// 3. 记录转账信息
	// 4. ��除手续费
	// 5. 记录手续费
	// 6. 推送消息
	//验签
	var resp Withresp
	if err := ctx.ShouldBindJSON(&resp); err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		return
	}
	data := map[string]interface{}{}
	for k, v := range utils.StructToMap(resp) {
		data[strings.ToLower(k)] = v
	}
	fmt.Println("PayoutBack", data)
	signs, _ := util.DoSign(data, global.SHOP_CONFIG.Wallet.Appkey)
	fmt.Println(signs)
	sign, err := util.VerifySign(data, global.SHOP_CONFIG.Wallet.Appkey, resp.Sign)
	if err != nil || !sign {
		global.SHOP_LOG.Log(0, fmt.Sprintf("充值回调验签失败%s", err.Error()))
		this.Success(ctx, "充值回调验签失败")
		return
	}
	var withresult model.UsdtWithdrawModel
	err = global.SHOP_DB.Where("wallet_path=? and id=?", resp.Address, resp.ThirdPartyID).Find(&withresult).Error
	if err != nil {
		global.SHOP_LOG.Log(0, err.Error())
		this.Error(ctx, err.Error())
		return
	}
	if withresult.Id == 0 {
		this.Error(ctx, "没有记录")
		return
	} else {
		withresult.Status = 1
		withresult.Hash = resp.Txid
		global.SHOP_DB.Updates(&withresult)
	}
	this.Success(ctx, "success")

}
