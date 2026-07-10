package crondtab

import (
	"fmt"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/admin/service"
	"ginshop.com/utils"
	"github.com/shopspring/decimal"
	"log"
	"time"

	"ginshop.com/admin/model"
	"ginshop.com/global"
	"github.com/go-co-op/gocron"
)

func Initcrond() {

	s := gocron.NewScheduler(time.UTC)
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Error loading location: %s", err)
	}
	s.ChangeLocation(location)

	s.Every(30).Second().Do(iptoaddr)

	//s.Every(1).Day().At("16:00").Do(Initquanaccount)

	go func() {
		for {

			//if global.SHOP_CONFIG.System.Env == "pro" {
			//	SendCode()
			//}
			if global.SHOP_CONFIG.System.Env == "debug" {

				global.SHOP_DB.Model(model.AccountCheckCode{}).Where("status=0").Updates(model.AccountCheckCode{
					Status:  1,
					Captcha: "123456",
				})

			}
			time.Sleep(time.Second)
		}
	}()
	s.StartAsync()

}

func WalletAddCreate() {
	var wallet WalletServer
	wallet.CrondCreateAdd()

}
func WalletWithdraw() {
	var wallet WalletServer
	wallet.CrondWithPassToUdun()

}
func Levlehook() {
	for {
		var levellog model.LevelUpdateLog
		global.SHOP_DB.Model(model.LevelUpdateLog{}).Where("status=0").Order("rand()").Find(&levellog)

		if levellog.Id != 0 {
			couponremark := ""
			points := decimal.Zero
			var user model.User
			global.SHOP_DB.Model(model.User{}).Where("username=?", levellog.Username).Find(&user)
			var levle model.MemberLevel
			global.SHOP_DB.Model(model.MemberLevel{}).Where("level=?", levellog.Newlevel).Find(&levle)
			zero := 0
			if levle.Id != 0 {
				if *levle.CouponId != zero && *levle.CouponNum != zero && levle.Id != 0 {

					//送coupon送积分
					var coupon model.TbCoupon

					global.SHOP_DB.Where("id=?", levle.CouponId).Find(&coupon)
					if coupon.Id == 0 {
						global.SHOP_LOG.Error("找不到coupon")
						continue
					}

					for i := 0; i < *levle.CouponNum; i++ {
						var couponlist model.TbCouponList
						//timeexp := time.Now()
						global.SHOP_DB.Where("coupon_id=?  and status=0 ", levle.CouponId).Limit(1).Find(&couponlist)
						if couponlist.Id == 0 {
							var couponserver service.STbCoupon
							var id request.GetById
							id.ID = uint(coupon.Id)
							//生成一批
							couponserver.Coupongen(id)
							global.SHOP_LOG.Error("找不到couponlist")
							continue
						}
						now := model.LocalTime(time.Now())
						couponlist.Username = user.Username
						couponlist.Uid = user.Id
						couponlist.Status = 2
						extime := model.LocalTime(time.Now().AddDate(0, 0, coupon.Cycle))
						couponlist.ActivateTime = &now
						couponlist.ExpDate = &extime
						couponlist.Remarks = "升级赠送"

						couponlist.UpdatedAt = &now
						err := global.SHOP_DB.Updates(&couponlist).Error
						if err != nil {
							global.SHOP_LOG.Error(err.Error())
							continue
						}
						couponremark = couponremark + "Code:(" + couponlist.CouponCode + ")"
					}
				}
			}
			if levle.Points.GreaterThan(decimal.Zero) {
				err := global.SHOP_DB.Exec("update account_funds set points=points+? where username=?", levle.Points, user.Username).Error
				if err != nil {

					global.SHOP_LOG.Error(err.Error())
					continue
				}
				points = levle.Points
				var saccountfundslog service.AccountFundsLog

				err, _ = saccountfundslog.Createlog(user.Username, levle.Points, utils.Updatelevels, fmt.Sprintf("升级送积分"), 2)

			}
			levellog.Status = 1
			levellog.Remark = levellog.Remark + fmt.Sprintf("升级送积分%s,coupon%s", points.Round(4).StringFixed(4), couponremark)
			global.SHOP_DB.Updates(&levellog)
		}
		time.Sleep(5 * time.Second)

	}
}
func Couponcheck() {
	for {
		var coupon []model.TbCoupon

		global.SHOP_DB.Where("status=1").Find(&coupon)
		if len(coupon) == 0 {
			continue
		}
		for _, v := range coupon {
			var sum int64
			//err = global.SHOP_DB.Model(model.OrderDeal{}).Select("COALESCE(SUM(profit), 0)").Where("quan_account_id=?", v.id).Scan(&sum).Error
			err := global.SHOP_DB.Model(model.TbCouponList{}).Where("coupon_id=? and status=0", v.Id).Count(&sum).Error

			if err != nil {
				global.SHOP_LOG.Error(err.Error())
				continue
			}
			if sum < int64(v.Numbers) {
				var couponserver service.STbCoupon
				var id request.GetById
				id.ID = uint(v.Id)
				//生成一批
				couponserver.Coupongen(id)
			}
		}
		var excoupon model.TbCouponList
		//删除过期的未使用的coupon
		global.SHOP_DB.Model(model.TbCouponList{}).Where(" status in(0,2) and exp_date<?", time.Now()).Delete(&excoupon)
		time.Sleep(10 * time.Second)
	}
}
