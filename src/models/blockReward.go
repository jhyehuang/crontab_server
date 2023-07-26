package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Rewards struct {
	gorm.Model
	RewardType   int8            `gorm:"column:reward_type;type:tinyint(1);comment:'奖励类型'" json:"reward_type"`            // 奖励类型
	Height       uint64          `gorm:"column:height;type:bigint;comment:'高度'" json:"height"`                            // 高度
	MinerAddr    string          `gorm:"column:miner_addr;type:varchar(256);comment:'矿工地址'" json:"miner_addr"`            // 矿工地址
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(65);comment:'奖励金额'" json:"amount"`                     // 奖励金额
	BlockCID     string          `gorm:"column:block_cid;type:varchar(256);comment:'区块CID'" json:"block_cid"`             // 区块CID
	Source       string          `gorm:"column:source;type:varchar(256);comment:'数据来源'" json:"source"`                    // 奖励来源
	MessageCID   string          `gorm:"column:message_cid;type:varchar(256);comment:'消息CID, unique'" json:"message_cid"` // 消息CID
	VMMessageCID string          `gorm:"column:vm_message_cid;type:varchar(256);comment:'VM消息CID'" json:"vm_message_cid"` // VM消息CID
	From         string          `gorm:"column:from;type:varchar(256);comment:'发送地址'" json:"from"`                        // 发送地址
	To           string          `gorm:"column:to;type:varchar(256);comment:'接收地址'" json:"to"`                            // 接收地址
	IsRelease    bool            `gorm:"column:is_release;type:tinyint(1);comment:'是否已经释放'" json:"is_release"`
}

// 出块奖励释放记录
type BlockRewardReleaseDetail struct {
	gorm.Model
	MinerId        string          `gorm:"column:miner_id;type:varchar(128);uniqueIndex:S_R;comment:'Miner ID'" json:"miner_id"`         // Miner ID
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R;comment:'资产包id'" json:"asset_pack_id,string"` // 资产包id
	Height         uint64          `gorm:"column:height;type:bigint;uniqueIndex:S_R;comment:'Block height'" json:"height"`               // Block height
	ReleasedReward decimal.Decimal `gorm:"column:release_reward;type:decimal(65);comment:'释放奖励'" json:"release_reward"`                  // 释放奖励
	// ReleaseTime    time.Time       `gorm:"column:release_time;type:timestamp;kcomment:'释放时间'" json:"release_time"`                // 释放时间
}

type BlockRewardReleaseLog struct {
	gorm.Model
	MinerId        string          `gorm:"column:miner_id;type:varchar(128);comment:'Miner ID'" json:"miner_id"`         // Miner ID
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;comment:'资产包id'" json:"asset_pack_id,string"` // 资产包id
	Height         uint64          `gorm:"column:height;type:bigint;comment:'Block height'" json:"height"`               // Block height
	ReleasedReward decimal.Decimal `gorm:"column:release_reward;type:decimal(65);comment:'释放奖励'" json:"release_reward"`  // 释放奖励
	ReleaseCount   uint64          `gorm:"column:release_count;type:bigint;comment:'释放次数'" json:"release_count"`         // 释放次数
}

// 出块奖励释放自然日记录表
type BlockRewardReleaseDayDetail struct {
	gorm.Model
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R;comment:'资产包id'" json:"asset_pack_id,string"` // 资产包id
	Date           string          `gorm:"column:date;type:varchar(256);uniqueIndex:S_R;comment:'日期，00:00 - 24:00'" json:"date"`         // 日期，00:00 - 24:00
	ReleasedReward decimal.Decimal `gorm:"column:release_reward;type:decimal(65);comment:'释放奖励'" json:"release_reward"`                  // 释放奖励
}

// 出块奖励分配参数
type BlockRewardAllocRecord struct {
	gorm.Model
	MinerId        string          `gorm:"column:miner_id;type:varchar(128);uniqueIndex:S_R;comment:'Miner ID'" json:"miner_id"`            // Miner ID
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R;comment:'资产包id'" json:"asset_pack_id,string"`    // 资产包id
	Height         uint64          `gorm:"column:height;type:bigint;uniqueIndex:S_R;comment:'Block height'" json:"height"`                  // Block height
	Power          decimal.Decimal `gorm:"column:power;type:decimal(65);comment:'算力'" json:"power"`                                         // 资产包算力
	MinerPower     decimal.Decimal `gorm:"column:miner_power;type:decimal(65);comment:'矿工算力'" json:"miner_power"`                           // 矿工算力
	AssetPackRate  decimal.Decimal `gorm:"column:asset_pack_rate;type:decimal(6,4);comment:'资产包占比,出块当时高度的资产包质押币数量'" json:"asset_pack_rate"` // 资产包占比,出块当时高度的资产包质押币数量
	AssetPackCount int             `gorm:"column:asset_pack_count;type:tinyint(1);comment:'可用资产包数量'" json:"asset_pack_count"`               // 可用资产包数量
}
