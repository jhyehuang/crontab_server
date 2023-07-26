package models

import "gorm.io/gorm"

type EventSync struct {
	gorm.Model
	Tx              string `gorm:"column:tx;type:varchar(256);uniqueIndex:S_R_1;comment:'Transaction ID'" json:"tx"`                            // Transaction ID
	ContractAddress string `gorm:"column:contract_address;type:varchar(256);comment:'contract address'" json:"contract_address"`                // contract address
	Height          uint64 `gorm:"column:height;type:bigint;comment:'block height'" json:"height"`                                              // block height
	EventSign       string `gorm:"column:event_sign;type:varchar(256);comment:'event sign'" json:"event_sign"`                                  // event sign
	Pyload          string `gorm:"column:pyload;type:text;comment:'Payload'" json:"pyload"`                                                     // Payload
	AssetPackId     uint64 `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R_1;comment:'资产包id'" json:"asset_pack_id,string"`              // 资产包id
	EventSignHash   string `gorm:"column:event_sign_hash;type:varchar(256);uniqueIndex:S_R_1;comment:'event sign hash'" json:"event_sign_hash"` // event sign
}
