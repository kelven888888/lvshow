package service

import (
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
)

type OrderDeal struct {
}

func (this *OrderDeal) GetAll(pageInfo request.PageInfo) ([]model.OrderDeal, int64) {
	var models []model.OrderDeal

	query := global.SHOP_DB.Model(model.OrderDeal{})
	if pageInfo.Keyword != "" {

		query.Where("td_acc = ? ", pageInfo.Keyword)
	}
	if pageInfo.OrderTradeType != 0 {

		query.Where("order_trade_type =? ", pageInfo.OrderTradeType)
	}
	if pageInfo.OrderType != 0 {

		query.Where("order_type =? ", pageInfo.OrderType)
	}
	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}

	return models, count

}
