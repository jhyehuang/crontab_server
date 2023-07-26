package models

import (
	"gorm.io/gorm"
)

type BannerBackground struct {
	gorm.Model
	Url       string `gorm:"column:url;type:varchar(256)" json:"url"`                            // url
	IsDefault bool   `gorm:"column:is_default;type:tinyint(1);comment:'是否默认'" json:"is_default"` // is_default

}
