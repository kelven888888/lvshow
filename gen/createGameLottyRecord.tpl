INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1076, 54, '抽奖记录管理', '', 1, NULL, 'GameLottyRecord', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1076, '抽奖记录编辑', '', 0, NULL, 'GameLottyRecord', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1076, '抽奖记录删除', '', 0, NULL, 'GameLottyRecord', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1076, '抽奖记录添加', '', 0, NULL, 'GameLottyRecord', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1076, '抽奖记录批量删除', '', 0, NULL, 'GameLottyRecord', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type GameLottyRecord struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
UserName string `comment:"用户名称" types:"" text:"" json:"user_name" form:"user_name" range:"" edit:""  `
GoodsId int `comment:"产品id" types:"" text:"" json:"goods_id" form:"goods_id" range:"" edit:""  `
GoodsType string `comment:"产品类型" types:"" text:"" json:"goods_type" form:"goods_type" range:"" edit:""  `
UserId int `comment:"用户id" types:"" text:"" json:"user_id" form:"user_id" range:"" edit:""  `
PlayId int `comment:"玩法id" types:"" text:"" json:"play_id" form:"play_id" range:"" edit:""  `
PlayName string `comment:"玩法名称" types:"" text:"" json:"play_name" form:"play_name" range:"" edit:""  `
Remark string `comment:"备注" types:"" text:"" json:"remark" form:"remark" range:"" edit:""  `
RewardType string `comment:"商品奖励类型" types:"" text:"" json:"reward_type" form:"reward_type" range:"" edit:""  `
}