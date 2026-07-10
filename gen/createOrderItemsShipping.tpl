INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1167, 54, '发货订单详情管理', '', 1, NULL, 'OrderItemsShipping', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1167, '发货订单详情编辑', '', 0, NULL, 'OrderItemsShipping', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1167, '发货订单详情删除', '', 0, NULL, 'OrderItemsShipping', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1167, '发货订单详情添加', '', 0, NULL, 'OrderItemsShipping', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1167, '发货订单详情批量删除', '', 0, NULL, 'OrderItemsShipping', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type OrderItemsShipping struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
OrderId int `comment:"订单id" types:"" text:"" json:"order_id" form:"order_id" range:"" edit:""  `
ProductId int `comment:"产品id" types:"" text:"" json:"product_id" form:"product_id" range:"" edit:""  `
ProductName string `comment:"产品名称" types:"" text:"" json:"product_name" form:"product_name" range:"" edit:""  `
ProductImg string `comment:"产品图片" types:"" text:"" json:"product_img" form:"product_img" range:"" edit:""  `
Qty int `comment:"数量" types:"" text:"" json:"qty" form:"qty" range:"" edit:""  `
RewardType string `comment:"商品奖励类型" types:"" text:"" json:"reward_type" form:"reward_type" range:"" edit:""  `
}