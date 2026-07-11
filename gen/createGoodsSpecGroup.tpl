INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1216, 54, '商品规格管理', '', 1, NULL, 'GoodsSpecGroup', 'Index', 1216, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1216, '商品规格编辑', '', 0, NULL, 'GoodsSpecGroup', 'Edit', 1216, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1216, '商品规格删除', '', 0, NULL, 'GoodsSpecGroup', 'Delete', 1216, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1216, '商品规格添加', '', 0, NULL, 'GoodsSpecGroup', 'Add', 1216, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1216, '商品规格批量删除', '', 0, NULL, 'GoodsSpecGroup', 'Deletebatch', 1216, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type GoodsSpecGroup struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Cid int64 `comment:"类别" types:"" text:"" json:"cid" form:"cid" range:"" edit:""  `
Name string `comment:"名称" types:"" text:"" json:"name" form:"name" range:"" edit:""  `
}