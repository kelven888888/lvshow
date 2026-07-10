INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1094, 54, '订单管理', '', 1, NULL, 'Orders', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1094, '订单编辑', '', 0, NULL, 'Orders', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1094, '订单删除', '', 0, NULL, 'Orders', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1094, '订单添加', '', 0, NULL, 'Orders', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1094, '订单批量删除', '', 0, NULL, 'Orders', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type Orders struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
UserIdSell int `comment:"出售会员ID" types:"" text:"" json:"user_id_sell" form:"user_id_sell" range:"" edit:""  `
UserIdBuy int `comment:"购买会员ID" types:"" text:"" json:"user_id_buy" form:"user_id_buy" range:"" edit:""  `
OrderSn string `comment:"订单号" types:"" text:"" json:"order_sn" form:"order_sn" range:"" edit:""  `
TotalAmount decimal.Decimal `comment:"总价" types:"" text:"" json:"total_amount" form:"total_amount" range:"" edit:""  `
PayAmount decimal.Decimal `comment:"支付金额" types:"" text:"" json:"pay_amount" form:"pay_amount" range:"" edit:""  `
Status int `comment:"状态" types:"" text:"" json:"status" form:"status" range:"" edit:""  `
ReceiverName string `comment:"" types:"" text:"" json:"receiver_name" form:"receiver_name" range:"" edit:""  `
ReceiverPhone string `comment:"" types:"" text:"" json:"receiver_phone" form:"receiver_phone" range:"" edit:""  `
ReceiverAddress string `comment:"" types:"" text:"" json:"receiver_address" form:"receiver_address" range:"" edit:""  `
Remark string `comment:"" types:"" text:"" json:"remark" form:"remark" range:"" edit:""  `
PayTime *time.Time `comment:"成交时间" types:"" text:"" json:"pay_time" form:"pay_time" range:"" edit:""  `
OrderType int `comment:"用户名" types:"" text:"" json:"order_type" form:"order_type" range:"" edit:""  `
}