INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1110, 54, '积分兑换记录管理', '', 1, NULL, 'PointRedeemRecord', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1110, '积分兑换记录编辑', '', 0, NULL, 'PointRedeemRecord', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1110, '积分兑换记录删除', '', 0, NULL, 'PointRedeemRecord', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1110, '积分兑换记录添加', '', 0, NULL, 'PointRedeemRecord', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1110, '积分兑换记录批量删除', '', 0, NULL, 'PointRedeemRecord', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type PointRedeemRecord struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Remarks string `comment:"" types:"" text:"" json:"remarks" form:"remarks" range:"" edit:""  `
ProductId int `comment:"产品id" types:"" text:"" json:"product_id" form:"product_id" range:"" edit:""  `
Username string `comment:"用户名称" types:"" text:"" json:"username" form:"username" range:"" edit:""  `
UserId int `comment:"用户id" types:"" text:"" json:"user_id" form:"user_id" range:"" edit:""  `
}