package models

import (
	"gorm.io/gorm"
)

type MinerInfo struct {
	gorm.Model
	MinerId        string `gorm:"column:miner_id;type:varchar(128);uniqueIndex:S_R;comment:'Miner ID'" json:"miner_id"`          // Miner ID
	OwnerId        string `gorm:"column:owner_id;type:varchar(128);comment:'Owner ID'" json:"owner_id"`                          // Owner ID
	OriOwnerId     string `gorm:"column:ori_owner_id;type:varchar(128);comment:'Ori Owner ID'" json:"ori_owner_id"`              // Ori Owner ID
	ManagerAddress string `gorm:"column:manager_address;type:varchar(128);comment:'Manager address'" json:"manager_address"`     // Manager address
	Height         uint64 `gorm:"column:height;type:bigint;comment:'Epoch this rewards summary applies to.'" json:"height"`      // Epoch this rewards summary applies to.
	EventHeight    uint64 `gorm:"column:event_height;type:bigint;comment:'事件同步高度'" json:"event_height"`                          // 事件同步高度
	AccOwnerHeight uint64 `gorm:"column:acc_owner_height;type:bigint;comment:'接管owner的epoch'" json:"acc_owner_height"`           // Acc Owner Height
	IsSyncSector   uint8  `gorm:"column:is_sync_sector;type:tinyint(1);comment:'是否需要同步扇区信息：0：不同步 1:需要同步'" json:"is_sync_sector"` // 是否需要同步扇区信息
	Region         string `gorm:"column:region;type:varchar(128);comment:'节点所在区域'" json:"region"`                                // 节点所在区域
}
