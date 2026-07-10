INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1145, 118, '优惠券列表管理', '', 1, NULL, 'TbCouponList', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1145, '优惠券列表编辑', '', 0, NULL, 'TbCouponList', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1145, '优惠券列表删除', '', 0, NULL, 'TbCouponList', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1145, '优惠券列表添加', '', 0, NULL, 'TbCouponList', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1145, '优惠券列表批量删除', '', 0, NULL, 'TbCouponList', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type TbCouponList struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
CouponCode string `comment:"券码" types:"" text:"" json:"coupon_code" form:"coupon_code" range:"" edit:""  `
Status int `comment:"状态" types:"" text:"" json:"status" form:"status" range:"" edit:""  `
Price decimal.Decimal `comment:"价值" types:"" text:"" json:"price" form:"price" range:"" edit:""  `
CouponId int `comment:"优惠券id" types:"" text:"" json:"coupon_id" form:"coupon_id" range:"" edit:""  `
Uid int `comment:"用户id" types:"" text:"" json:"uid" form:"uid" range:"" edit:""  `
ActivateTime string `comment:"" types:"" text:"" json:"activate_time" form:"activate_time" range:"" edit:""  `
CouponTitle string `comment:"券码" types:"" text:"" json:"coupon_title" form:"coupon_title" range:"" edit:""  `
Username string `comment:"用户" types:"" text:"" json:"username" form:"username" range:"" edit:""  `
ExpDate time.Time `comment:"失效时间" types:"" text:"" json:"exp_date" form:"exp_date" range:"" edit:""  `
}