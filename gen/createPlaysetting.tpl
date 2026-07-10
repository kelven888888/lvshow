INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1031, 54, '玩法配置管理', '', 1, NULL, 'Playsetting', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1031, '玩法配置编辑', '', 0, NULL, 'Playsetting', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1031, '玩法配置删除', '', 0, NULL, 'Playsetting', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1031, '玩法配置添加', '', 0, NULL, 'Playsetting', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1031, '玩法配置批量删除', '', 0, NULL, 'Playsetting', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type Playsetting struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
RateArr string `comment:"概率配置" types:"" text:"" json:"rate_arr" form:"rate_arr" range:"" edit:""  `
RewardArr string `comment:"概率配置" types:"" text:"" json:"reward_arr" form:"reward_arr" range:"" edit:""  `
SingelStatus int `comment:"单人模式开关" types:"radio" text:"关闭,启用" json:"singel_status" form:"singel_status" range:"0,1" edit:""  `
DoubleStatus int `comment:"双人模式开关" types:"radio" text:"关闭,启用" json:"double_status" form:"double_status" range:"0,1" edit:""  `
Price int `comment:"参加价格" types:"" text:"" json:"price" form:"price" range:"" edit:""  `
PointArr string `comment:"" types:"" text:"" json:"point_arr" form:"point_arr" range:"" edit:""  `
GoodsIds string `comment:"" types:"" text:"" json:"goods_ids" form:"goods_ids" range:"" edit:""  `
Type int `comment:"类型" types:"radio" text:"高爆,保底,魔王,PK,随机" json:"type" form:"type" range:"1,2,3,4,5" edit:""  `
Name string `comment:"名称" types:"" text:"" json:"name" form:"name" range:"" edit:""  `
}