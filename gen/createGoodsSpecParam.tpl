INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1225, 54, '商品规格参数管理', '', 1, NULL, 'GoodsSpecParam', 'Index', 1225, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1225, '商品规格参数编辑', '', 0, NULL, 'GoodsSpecParam', 'Edit', 1225, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1225, '商品规格参数删除', '', 0, NULL, 'GoodsSpecParam', 'Delete', 1225, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1225, '商品规格参数添加', '', 0, NULL, 'GoodsSpecParam', 'Add', 1225, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1225, '商品规格参数批量删除', '', 0, NULL, 'GoodsSpecParam', 'Deletebatch', 1225, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type GoodsSpecParam struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
Cid int `comment:"" types:"" text:"" json:"cid" form:"cid" range:"" edit:""  `
GroupId int `comment:"属性组id" types:"" text:"" json:"group_id" form:"group_id" range:"" edit:""  `
Name string `comment:"名称" types:"" text:"" json:"name" form:"name" range:"" edit:""  `
Numeric int `comment:"数字类型参数" types:"radio" text:"是,否" json:"numeric" form:"numeric" range:"1,2" edit:""  `
Unit string `comment:"单位" types:"" text:"" json:"unit" form:"unit" range:"" edit:""  `
Generic int `comment:"sku通用属性" types:"radio" text:"是,否" json:"generic" form:"generic" range:"1,2" edit:""  `
Searching int `comment:"用于搜索过滤" types:"radio" text:"是,否" json:"searching" form:"searching" range:"1,2" edit:""  `
Segments string `comment:"数值类型参数" types:"" text:"" json:"segments" form:"segments" range:"" edit:""  `
}