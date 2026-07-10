INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (1183, 54, '收货地址管理', '', 1, NULL, 'ShippingAddresses', 'Index', 1183, 0, '2024-07-13 12:51:12.094', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1183, '收货地址编辑', '', 0, NULL, 'ShippingAddresses', 'Edit', 1183, 0, '2024-07-13 12:51:12.112', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1183, '收货地址删除', '', 0, NULL, 'ShippingAddresses', 'Delete', 1183, 0, '2024-07-13 12:51:12.114', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1183, '收货地址添加', '', 0, NULL, 'ShippingAddresses', 'Add', 1183, 0, '2024-07-13 12:51:12.115', '2024-06-09 12:03:11.987', NULL);
INSERT INTO `nov_role` (`id`, `pid`, `name`, `icon`, `is_menu`, `desc`, `module`, `action`, `sort`, `is_default`, `created_at`, `updated_at`, `deleted_at`) VALUES (NULL, 1183, '收货地址批量删除', '', 0, NULL, 'ShippingAddresses', 'Deletebatch', 1183, 0, '2024-07-13 12:51:12.117', '2024-06-09 12:03:11.987', NULL);
package model

import "time" 
type ShippingAddresses struct {
Model model.Model `comment:"" types:"" text:"" json:"model" form:"model" range:"" edit:""  `
ReceiverName string `comment:"收货人" types:"" text:"" json:"receiver_name" form:"receiver_name" range:"" edit:""  `
ReceiverPhone string `comment:"收货人电话" types:"" text:"" json:"receiver_phone" form:"receiver_phone" range:"" edit:""  `
AddressLine1 string `comment:"地址1" types:"" text:"" json:"address_line1" form:"address_line1" range:"" edit:""  `
AddressLine2 string `comment:"地址2" types:"" text:"" json:"address_line2" form:"address_line2" range:"" edit:""  `
Area string `comment:"" types:"" text:"" json:"area" form:"area" range:"" edit:""  `
Username string `comment:"用户名" types:"" text:"" json:"username" form:"username" range:"" edit:""  `
City string `comment:"城市" types:"" text:"" json:"city" form:"city" range:"" edit:""  `
State string `comment:"州/联邦直辖区" types:"" text:"" json:"state" form:"state" range:"" edit:""  `
PostalCode string `comment:"邮编" types:"" text:"" json:"postal_code" form:"postal_code" range:"" edit:""  `
CountryCode string `comment:"国家ISO代码" types:"" text:"" json:"country_code" form:"country_code" range:"" edit:""  `
Latitude string `comment:"" types:"" text:"" json:"latitude" form:"latitude" range:"" edit:""  `
Longitude string `comment:"" types:"" text:"" json:"longitude" form:"longitude" range:"" edit:""  `
IsDefault int `comment:"是否默认" types:"radio" text:"否,是" json:"is_default" form:"is_default" range:"0,1" edit:""  `
Page int `comment:"" types:"" text:"" json:"page" form:"page" range:"" edit:""  `
}