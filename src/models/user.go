package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Address string `gorm:"column:address;type:varchar(256);uniqueIndex:S_R" json:"address"` // address
	Name    string `gorm:"column:name;type:varchar(256);uniqueIndex:S_R" json:"name"`       // name
	Url     string `gorm:"column:url;type:varchar(256)" json:"url"`                         // url

}

type UserTest struct {
	TotalPledgeAmount string `gorm:"column:total_pledge_amount;type:decimal(65)" json:"total_pledge_amount"` // 总质押金额

}
