INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1115, 118, '优惠券管理', '', 1, NULL, 'TbCoupon', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1115, '优惠券编辑', '', 0, NULL, 'TbCoupon', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1115, '优惠券删除', '', 0, NULL, 'TbCoupon', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1115, '优惠券添加', '', 0, NULL, 'TbCoupon', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1115, '优惠券批量删除', '', 0, NULL, 'TbCoupon', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type TbCoupon struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Title string `comment:"标题" types:"" text:"" json:"title" form:"title" range:"" edit:""  `
Price decimal.Decimal `comment:"价值" types:"" text:"价值" json:"price" form:"price" range:"" edit:""  `
Content string `comment:"" types:"" text:"" json:"content" form:"content" range:"" edit:""  `
Numbers int `comment:"生成个数" types:"" text:"生成个数" json:"numbers" form:"numbers" range:"" edit:""  `
Isdelete int `comment:"" types:"" text:"" json:"isdelete" form:"isdelete" range:"" edit:""  `
Cycle int `comment:"周期" types:"" text:"周期" json:"cycle" form:"cycle" range:"" edit:""  `
Status int `comment:"状态" types:"radio" text:"关闭,启用" json:"status" form:"status" range:"0,1" edit:""  `
}