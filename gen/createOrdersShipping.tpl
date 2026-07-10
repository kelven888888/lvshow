INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1172, 54, '发货订单管理', '', 1, NULL, 'OrdersShipping', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1172, '发货订单编辑', '', 0, NULL, 'OrdersShipping', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1172, '发货订单删除', '', 0, NULL, 'OrdersShipping', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1172, '发货订单添加', '', 0, NULL, 'OrdersShipping', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1172, '发货订单批量删除', '', 0, NULL, 'OrdersShipping', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type OrdersShipping struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
UserId int `comment:"会员ID" types:"" text:"" json:"user_id" form:"user_id" range:"" edit:""  `
OrderSn string `comment:"订单号" types:"" text:"" json:"order_sn" form:"order_sn" range:"" edit:""  `
Status *int `comment:"状态" types:"" text:"" json:"status" form:"status" range:"" edit:""  `
ReceiverName string `comment:"收货人姓名" types:"" text:"" json:"receiver_name" form:"receiver_name" range:"" edit:""  `
ReceiverPhone string `comment:"收货人电话" types:"" text:"" json:"receiver_phone" form:"receiver_phone" range:"" edit:""  `
ReceiverAddress string `comment:"收货地址快照" types:"" text:"" json:"receiver_address" form:"receiver_address" range:"" edit:""  `
Remark string `comment:"" types:"" text:"" json:"remark" form:"remark" range:"" edit:""  `
TotalQty int `comment:"订单总数量" types:"" text:"" json:"total_qty" form:"total_qty" range:"" edit:""  `
TotalProductQty int `comment:"订单总数量" types:"" text:"" json:"total_product_qty" form:"total_product_qty" range:"" edit:""  `
Chindren []*model.OrderItems `comment:"" types:"" text:"" json:"chindren" form:"chindren" range:"" edit:""  `
}