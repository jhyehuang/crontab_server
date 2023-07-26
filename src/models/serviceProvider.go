package models

import "gorm.io/gorm"

type ServiceProvider struct {
	gorm.Model
	FullName string `gorm:"column:full_name;type:varchar(256);comment:'全称'" json:"full_name"` // full_name
	// 简称
	ShortName string `gorm:"column:short_name;type:varchar(256);comment:'简称'" json:"short_name"` // short_name
	// 简介
	Introduction string `gorm:"column:introduction;type:text;comment:'简介'" json:"introduction"` // introduction
	// logo
	LogoUrl string `gorm:"column:logo_url;type:text;comment:'logo'" json:"logo_url"` // logo
	// 钱包地址
	WalletAddress string `gorm:"column:wallet_address;type:varchar(256);comment:'钱包地址'" json:"wallet_address"` // wallet_address

	// 默认服务商
	IsDefault bool `gorm:"column:is_default;type:tinyint(1);comment:'是否默认服务商'" json:"is_default"` // is_default
}
