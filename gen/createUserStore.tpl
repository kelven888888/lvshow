INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1071, 54, '仓库管理', '', 1, NULL, 'UserStore', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1071, '仓库编辑', '', 0, NULL, 'UserStore', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1071, '仓库删除', '', 0, NULL, 'UserStore', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1071, '仓库添加', '', 0, NULL, 'UserStore', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1071, '仓库批量删除', '', 0, NULL, 'UserStore', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type UserStore struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
UserId int `comment:"用户id" types:"" text:"" json:"user_id" form:"user_id" range:"" edit:""  `
Username string `comment:"用户名" types:"" text:"" json:"username" form:"username" range:"" edit:""  `
GoodsName string `comment:"产品名称" types:"" text:"" json:"goods_name" form:"goods_name" range:"" edit:""  `
GoodsId int `comment:"产品id" types:"" text:"" json:"goods_id" form:"goods_id" range:"" edit:""  `
Qty int `comment:"数量" types:"" text:"" json:"qty" form:"qty" range:"" edit:""  `
Price decimal.Decimal `comment:"价格" types:"" text:"" json:"price" form:"price" range:"" edit:""  `
RewardType string `comment:"商品奖励类型" types:"" text:"" json:"reward_type" form:"reward_type" range:"" edit:""  `
}