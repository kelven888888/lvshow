INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1081, 54, '升级记录管理', '', 1, NULL, 'LevelUpdateLog', 'Index', 0, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1081, '升级记录编辑', '', 0, NULL, 'LevelUpdateLog', 'Edit', 0, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1081, '升级记录删除', '', 0, NULL, 'LevelUpdateLog', 'Delete', 0, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1081, '升级记录添加', '', 0, NULL, 'LevelUpdateLog', 'Add', 0, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1081, '升级记录批量删除', '', 0, NULL, 'LevelUpdateLog', 'Deletebatch', 0, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type LevelUpdateLog struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Oldlevel int `comment:"旧等级" types:"" text:"" json:"oldlevel" form:"oldlevel" range:"" edit:""  `
Newlevel int `comment:"新等级" types:"" text:"" json:"newlevel" form:"newlevel" range:"" edit:""  `
Exp decimal.Decimal `comment:"当前积分" types:"" text:"" json:"exp" form:"exp" range:"" edit:""  `
Remark string `comment:"备注" types:"" text:"" json:"remark" form:"remark" range:"" edit:""  `
}