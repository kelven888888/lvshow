package service

import (
	"context"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"ginshop.com/utils"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	rand2 "math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SGame struct {
}

func (this *SGame) LottyBlindBox(models model.Playsetting, req request.GetByLottyId, user model.User, lianjicount int, combocount int, combocounts int, discountmoney decimal.Decimal) (model.Goods, int, int, int, error) {
	fmt.Println("LottyBlindBox")
	var modelgood model.Goods
	//for i := 0; i < 100; i++ {
	reward, point, remark, err := this.Calprice(models, req.LotteryType, user)
	if err != nil {
		return modelgood, 0, 0, 0, err
	}
	//}
	now := model.LocalTime(time.Now())
	tran := global.SHOP_DB.Begin()
	//盲盒扣除盲盒次数
	if models.Id == 6 {
		err = tran.Exec("update auth_user set blind_box_num=blind_box_num-1 where id=?", user.Id).Error
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			tran.Rollback()
			return modelgood, 0, 0, 0, err
		}
	}
	fmt.Println(point)
	goodsid, err := this.GetAwardGoods(models, reward)
	if err != nil {
		global.SHOP_LOG.Error("获取不到奖品")
		tran.Rollback()
		return modelgood, 0, 0, 0, err
	}

	global.SHOP_DB.Where("id=?", goodsid).Find(&modelgood)
	if modelgood.Id == 0 {
		tran.Rollback()
		return modelgood, 0, 0, 0, errors.New("产品不存在")
	}
	var stores model.UserStore
	global.SHOP_DB.Where("user_id=? and goods_id=?", user.Id, goodsid).Find(&stores)
	if stores.Id == 0 {
		stores.Price = modelgood.UnitPrice
		stores.UserId = user.Id
		stores.Username = user.Username
		stores.GoodsName = modelgood.GoodsName
		stores.GoodsId = modelgood.Id
		stores.RewardType = modelgood.RewardType
		stores.Qty = 1
		var storeserver SUserStore
		err = storeserver.Save(&stores)
		if err != nil {
			global.SHOP_LOG.Error(err.Error())
			tran.Rollback()
			return modelgood, 0, 0, 0, errors.New("失败")
		}
	} else {
		err = global.SHOP_DB.Exec("update user_store set qty=qty+1 where user_id=? and goods_id=?", user.Id, goodsid).Error
		if err != nil {
			global.SHOP_LOG.Error("获取不到奖品")
			tran.Rollback()
			return modelgood, 0, 0, 0, errors.New("失败")
		}
	}

	var record model.GameLottyRecord
	record.Model.CreatedAt = &now
	record.Remark = remark
	record.PlayId = models.Id
	record.GoodsId = goodsid
	record.UserName = user.Username
	record.GoodsType = reward
	record.GoodsName = modelgood.GoodsName
	record.UserId = user.Id
	record.PlayName = models.Name
	record.RewardType = reward
	record.CouponCode = req.CouponCode
	record.Discount = discountmoney.Div(decimal.NewFromInt(int64(req.LotteryNum)))
	record.Money = models.Price
	record.LotteryType = req.LotteryType
	result := tran.Save(&record)
	if result.Error != nil {
		tran.Rollback()
		global.SHOP_LOG.Error(err.Error())
		return modelgood, 0, 0, 0, errors.New("失败")
	}

	//除了盲盒送积分
	var accountfunds model.AccountFunds
	err = tran.Where("id= ?", user.Id).Find(&accountfunds).Error
	if err != nil {
		tran.Rollback()

		return modelgood, 0, 0, 0, errors.New("失败")
	}
	logidarr := []uint{}

	zero := decimal.Zero
	//需要钱扣钱
	costmoney := decimal.Zero
	costmoney = models.Price.Sub(discountmoney.Div(decimal.NewFromInt(int64(req.LotteryNum))))
	if costmoney.GreaterThan(zero) {

		if point.GreaterThan(zero) {
			err = tran.Exec("update account_funds set ava_funds=ava_funds-?, points=points+? where uid=?", costmoney, point, user.Id).Error
		} else {
			err = tran.Exec("update account_funds set ava_funds=ava_funds-? where uid=?", costmoney, user.Id).Error
		}
		if err != nil {
			tran.Rollback()
			return modelgood, 0, 0, 0, errors.New("失败")
		}
		var saccountfundslog AccountFundsLog

		err, logid := saccountfundslog.Createlog(user.Username, costmoney.Neg(), utils.Lottery, fmt.Sprintf("抽奖%s", utils.Languagebycode("zh-hant", models.Name)), 1)
		logidarr = append(logidarr, logid)
		if err != nil {
			var mlog model.AccountFundsLog
			global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
			tran.Rollback()

			return modelgood, 0, 0, 0, errors.New("失败")
		}
		if point.GreaterThan(zero) {
			err, logid := saccountfundslog.Createlog(user.Username, point, utils.Lottery, fmt.Sprintf("抽奖%s", utils.Languagebycode("zh-hant", models.Name)), 2)
			logidarr = append(logidarr, logid)
			if err != nil {
				tran.Rollback()
				var mlog model.AccountFundsLog
				global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
				return modelgood, 0, 0, 0, errors.New("失败")
			}
		}
	} else {
		//不花钱单纯赚取积分
		if point.GreaterThan(zero) {

			err = tran.Exec("update account_funds set points=points+? where uid=?", point, user.Id).Error
			if err != nil {
				tran.Rollback()
				return modelgood, 0, 0, 0, errors.New("失败")
			}
			var saccountfundslog AccountFundsLog

			err, logid := saccountfundslog.Createlog(user.Username, point, utils.Lottery, fmt.Sprintf("抽奖%s", utils.Languagebycode("zh-hant", models.Name)), 2)
			logidarr = append(logidarr, logid)
			if err != nil {
				tran.Rollback()
				var mlog model.AccountFundsLog
				global.SHOP_DB.Model(model.AccountFundsLog{}).Where("id in ?", logidarr).Delete(&mlog)
				return modelgood, 0, 0, 0, errors.New("失败")
			}
		}
	}
	lijicouns := 0
	//保底奖才有
	if models.Id == 1 {
		lianjikey := fmt.Sprintf("lianji_%d_%d_%d", user.Id, models.Id, req.LotteryType)
		combokey := fmt.Sprintf("combo_%d_%d_%d", user.Id, models.Id, req.LotteryType)
		if reward == "D" {
			lijicouns = int(global.SHOP_REDIS.Incr(context.Background(), lianjikey).Val())
		} else {
			global.SHOP_REDIS.Set(context.Background(), lianjikey, 0, time.Second*3600*24*360)
			//global.SHOP_REDIS.Del(context.Background(), lianjikey)
		}
		vars, _ := global.SHOP_REDIS.Get(context.Background(), combokey).Result()
		combocounts, _ = strconv.Atoi(vars)
		//连击奖励 六个d 一个combo16个 连接
		if lijicouns == *models.Dhit {
			global.SHOP_REDIS.Del(context.Background(), lianjikey)
			lianjicount = lianjicount + 1
			combocounts = int(global.SHOP_REDIS.Incr(context.Background(), combokey).Val())
		}
		fmt.Println(combocounts, *models.Combo, combocounts == *models.Combo, "dddddddddddddd")
		if combocounts == *models.Combo {
			combocount = combocount + 1
			global.SHOP_REDIS.Set(context.Background(), combokey, 0, time.Second*3600*24*360)
		}
		//16连击赏
		if combocount == 1 {
			var goods model.Goods
			goodsid, err := this.GetAwardGoods(models, "COMBO")
			if err != nil {
				global.SHOP_LOG.Error("获取不到奖品")
				tran.Rollback()
				return modelgood, 0, 0, 0, err
			}

			global.SHOP_DB.Where("id=?", goodsid).Find(&goods)
			if goods.Id == 0 {
				tran.Rollback()
				return modelgood, 0, 0, 0, errors.New("产品不存在")
			}
			var storess model.UserStore
			global.SHOP_DB.Where("user_id=? and goods_id=?", user.Id, goodsid).Find(&storess)
			if storess.Id == 0 {
				storess.Price = goods.UnitPrice
				storess.UserId = user.Id
				storess.Username = user.Username
				storess.GoodsName = goods.GoodsName
				storess.GoodsId = goods.Id
				storess.RewardType = goods.RewardType
				storess.Qty = 1
				var storeserver SUserStore
				err = storeserver.Save(&storess)
				if err != nil {
					global.SHOP_LOG.Error(err.Error())
					tran.Rollback()
					return modelgood, 0, 0, 0, errors.New("失败")
				}
			} else {
				err = global.SHOP_DB.Exec("update user_store set qty=qty+1 where user_id=? and goods_id=?", user.Id, goods.Id).Error
				if err != nil {
					global.SHOP_LOG.Error("获取不到奖品")
					tran.Rollback()
					return modelgood, 0, 0, 0, errors.New("失败")
				}
			}
			remark = "COMBO奖励"
			reward = "COMBO"
			var record model.GameLottyRecord
			record.Model.CreatedAt = &now
			record.Remark = remark
			record.PlayId = models.Id
			record.GoodsId = goods.Id
			record.UserName = user.Username
			record.GoodsType = reward
			record.GoodsName = goods.GoodsName
			record.UserId = user.Id
			record.PlayName = models.Name
			record.RewardType = reward
			record.LotteryType = req.LotteryType
			result := tran.Save(&record)
			if result.Error != nil {
				tran.Rollback()
				global.SHOP_LOG.Error(err.Error())
				return modelgood, 0, 0, 0, errors.New("失败")
			}
		}
	}

	tran.Commit()
	if !models.Price.LessThan(zero) {
		go func() {
			//this.Updatelevel(models.Price, user.Id, models, record.Id)
			this.Lotteryhook(costmoney, user.Id, models, record.Id)

		}()
	}

	return modelgood, lianjicount, combocount, combocounts, nil
}
func (this *SGame) Calprice(models model.Playsetting, LotteryType uint64, user model.User) (string, decimal.Decimal, string, error) { //返回奖励产品和积分
	fmt.Println("Calprice")
	rate_arr := strings.Split(models.RateArr, ",")
	reward_arr := strings.Split(models.RewardArr, ",")
	points_arr := strings.Split(models.PointArr, ",")
	var slice []int
	rands := 0
	if LotteryType == 2 {

		rands = rand2.Intn(models.Div) + 1
	} else {

		lotkeys := fmt.Sprintf("lottery_%d_%d", user.Id, models.Id)
		//res, _ := global.SHOP_REDIS.Get(context.Background(), lotkeys).Result()
		length, err := global.SHOP_REDIS.SCard(context.Background(), lotkeys).Result()
		//if err != nil {
		//	fmt.Println("获取长度错误:", err)
		//	return "", 0.0, "", errors.New("错误")
		//}
		fmt.Println(length, "长度长度长度长度长度")
		wg := sync.WaitGroup{}
		if length == 0 || err != nil {
			for i := 1; i <= models.Div; i++ {
				wg.Add(1)
				go func(v int) {
					global.SHOP_REDIS.SAdd(context.Background(), lotkeys, i)
					wg.Done()
				}(i)

			}
			wg.Wait()
		}
		// 随机获取一个Set中的元素并从集合中移除它
		randomElement, err := global.SHOP_REDIS.SPop(context.Background(), lotkeys).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("返回数字是:", randomElement)

		rands, _ = strconv.Atoi(randomElement)

	}
	sum := 0
	//1.1,2,13,83.9
	for _, v := range rate_arr {
		floatVal, err := strconv.ParseFloat(v, 32)
		if err != nil {
			fmt.Println("转换错误:", err)
			if err != nil {
				return "", decimal.Zero, "", err
			}
		}
		value := int(floatVal * float64(models.Div) / 100)
		if err != nil {
			return "", decimal.Zero, "", err
		}
		sum = value + sum
		fmt.Println(v, int(value))
		slice = append(slice, sum)

	}
	len := len(slice)
	index := 0
	for i := 0; i < len; i++ {
		if i < len-1 {
			if slice[i] >= rands && rands < slice[i+1] {

				index = i
				break
			}
		} else {
			index = i
		}

	}
	floatVal, err := strconv.ParseFloat(points_arr[index], 32)
	if err != nil {
		fmt.Println("转换错误:", err)
		if err != nil {
			return "", decimal.Zero, "", err
		}
	}
	rewardval := reward_arr[index]
	if err != nil {
		fmt.Println("转换错误:", err)
		if err != nil {
			return "", decimal.Zero, "", err
		}
	}
	remark := fmt.Sprintf("Cal%+v%+v%+v%+v%+v%+v%+v", rands, slice, rate_arr, reward_arr, points_arr, rewardval, utils.Float64ToDecimal(floatVal))
	global.SHOP_LOG.Info(remark)
	fmt.Println(rewardval, floatVal)
	return rewardval, utils.Float64ToDecimal(floatVal), remark, nil

	return "", decimal.Zero, "", errors.New("错误")

}
func (this *SGame) GetAwardGoods(models model.Playsetting, reward_type string) (int, error) { //返回奖励产品和积分
	var modelgoods model.PlayGoods
	global.SHOP_DB.Raw("select  * from play_goods where play_id=? and reward_type=? order by rand()  limit 1", models.Id, reward_type).Scan(&modelgoods)

	return modelgoods.GoodsId, nil
}
func (this *SGame) Lotteryhook(money decimal.Decimal, userid int, models model.Playsetting, recordid int) error { //返回奖励产品和积分
	//1积分0.1

	tran := global.SHOP_DB.Begin()
	exp := money.Mul(models.MoneyToExp)
	var user model.User
	global.SHOP_DB.Where("id=?", userid).Find(&user)
	now := model.LocalTime(time.Now())
	if exp.GreaterThan(decimal.Zero) {

		err := tran.Exec("update auth_user set exp=exp+? where id=?", exp, user.Id).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			return nil
		}

		totalexp := exp.Add(user.Exp)
		var levle []model.MemberLevel
		oldlevel := user.Level
		var expre model.ExpRecord
		expre.Exp = exp
		expre.Oldexp = user.Exp
		expre.Newexp = totalexp
		expre.Model.CreatedAt = &now
		expre.RecordId = recordid
		expre.Username = user.Username
		err = tran.Save(&expre).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			return nil
		}
		tran.Commit()
		global.SHOP_DB.Where("`level`>? and exp<= ? ", user.Level, totalexp).Order("`level` asc ").Find(&levle)
		for _, value := range levle {
			if value.Id != 0 {
				tran = global.SHOP_DB.Begin()
				lev := value.Level
				err := global.SHOP_DB.Model(model.User{}).Where("id=?", user.Id).Updates(&model.User{
					Level: lev,
				}).Error
				if err != nil {
					global.SHOP_LOG.Error(err.Error())
					return nil
				}
				remark := fmt.Sprintf("会员等级%d升级%d-%s", *oldlevel, *lev, totalexp.Round(2).StringFixed(2))

				err = tran.Model(model.LevelUpdateLog{}).Save(&model.LevelUpdateLog{
					Model: model.Model{
						CreatedAt: &now,
					},
					Oldlevel: *oldlevel,
					Newlevel: *lev,
					Exp:      totalexp,
					Remark:   remark,
					Username: user.Username,
					Status:   0,
				}).Error
				if err != nil {
					tran.Rollback()
					global.SHOP_LOG.Error(err.Error())
				}
			}
		}

	}
	if money.LessThan(decimal.NewFromInt(1)) { //小于1不计入佣金
		return nil
	}
	//返佣
	commistion := money.Mul(models.CommissionRate.Div(decimal.NewFromInt(100)))
	if commistion.GreaterThan(decimal.Zero) {
		tran = global.SHOP_DB.Begin()
		err := tran.Exec("update account_funds set ava_funds=ava_funds+? where uid=?", commistion, user.Pid).Error
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			return nil
		}
		var saccountfundslog AccountFundsLog
		var userP model.User
		global.SHOP_DB.Where("id=?", user.Pid).Find(&userP)
		if userP.Id == 0 {
			global.SHOP_LOG.Error("没有推荐人,无须返佣")
			return nil
		}

		var commiss model.CommissionRecord
		commiss.Username = userP.Username
		commiss.UserId = userP.Id
		commiss.CreatedAt = &now
		commiss.Amount = commistion
		commiss.FromAmount = money
		commiss.FromUserId = user.Id
		commiss.FromUsername = user.Username
		commiss.LogType = utils.Commission
		tran.Save(&commiss)
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			return nil
		}
		err, _ = saccountfundslog.Createlog(userP.Username, commistion, utils.Commission, fmt.Sprintf("佣金抽奖ID:%d", recordid), 1)
		if err != nil {
			tran.Rollback()
			global.SHOP_LOG.Error(err.Error())
			return nil
		}
		tran.Commit()

	}

	return nil

}
