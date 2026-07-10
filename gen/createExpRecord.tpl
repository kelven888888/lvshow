INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1086, 54, '经验记录管理', '', 1, NULL, 'ExpRecord', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1086, '经验记录编辑', '', 0, NULL, 'ExpRecord', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1086, '经验记录删除', '', 0, NULL, 'ExpRecord', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1086, '经验记录添加', '', 0, NULL, 'ExpRecord', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1086, '经验记录批量删除', '', 0, NULL, 'ExpRecord', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type ExpRecord struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Oldexp decimal.Decimal `comment:"老经验" types:"" text:"" json:"oldexp" form:"oldexp" range:"" edit:""  `
Exp decimal.Decimal `comment:"获取经验" types:"" text:"" json:"exp" form:"exp" range:"" edit:""  `
Newexp decimal.Decimal `comment:"新经验" types:"" text:"" json:"newexp" form:"newexp" range:"" edit:""  `
RecordId int `comment:"抽奖记录id" types:"" text:"" json:"record_id" form:"record_id" range:"" edit:""  `
}