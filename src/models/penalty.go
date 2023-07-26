package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 浏览器惩罚
type Penalty struct {
	gorm.Model
	PenaltyType  int8            `gorm:"column:penalty_type;type:tinyint;comment:'惩罚类型'" json:"penalty_type"`                         // 惩罚类型
	Height       int64           `gorm:"column:height;type:bigint;comment:'区块高度'" json:"height"`                                      // 区块高度
	MinerAddr    string          `gorm:"column:miner_addr;type:varchar(256);comment:'矿工地址'" json:"miner_addr"`                        // 矿工地址
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(65);comment:'惩罚金额'" json:"amount"`                                 // 惩罚金额
	Source       string          `gorm:"column:source;type:varchar(256);comment:'惩罚来源'" json:"source"`                                // 惩罚来源
	MessageCID   string          `gorm:"column:message_cid;type:varchar(256);uniqueIndex;comment:'消息cid'" json:"message_cid"`         // 消息cid
	MessageTime  time.Time       `gorm:"column:message_time;type:datetime;comment:'消息时间'" json:"message_time"`                        // 消息时间
	VMMessageCID string          `gorm:"column:vm_message_cid;type:varchar(256);uniqueIndex;comment:'vm消息cid'" json:"vm_message_cid"` // vm消息cid
	From         string          `gorm:"column:from;type:varchar(256);comment:'发送地址'" json:"from"`                                    // from
	To           string          `gorm:"column:to;type:varchar(256);comment:'接收地址'" json:"to"`                                        // to
	IsRelease    bool            `gorm:"column:is_release;type:tinyint(1);comment:'是否已经释放'" json:"is_release"`
}

// 运维罚金
type SPPenalty struct {
	gorm.Model
	MinerId        string          `gorm:"column:miner_id;type:varchar(128);comment:'Miner ID'" json:"miner_id"`                                 // Miner ID
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:asset_tx;comment:'资产包id'" json:"asset_pack_id,string"`    // 资产包id
	ChainPenaltyId uint64          `gorm:"column:chain_penalty_id;type:varchar(256);comment:'penalty id'" json:"chain_penalty_id"`               // 方法
	Height         int64           `gorm:"column:height;type:bigint;comment:'Block height'" json:"height"`                                       // Block height
	TxHash         string          `gorm:"column:tx_hash;type:varchar(256);comment:'TxHash'" json:"tx_hash"`                                     // TxHash
	TxTime         time.Time       `gorm:"column:tx_time;type:datetime;comment:'TxTime'" json:"tx_time"`                                         // TxTime
	VMMessageCID   string          `gorm:"column:vm_message_cid;type:varchar(256);uniqueIndex:asset_tx;comment:'vm消息cid'" json:"vm_message_cid"` // vm消息cid
	From           string          `gorm:"column:from;type:varchar(256);comment:'from'" json:"from"`                                             // from
	To             string          `gorm:"column:to;type:varchar(256);comment:'to'" json:"to"`                                                   // to
	Amount         decimal.Decimal `gorm:"column:amount;type:decimal(65);comment:'惩罚金额'" json:"amount"`                                          // 交易金额
}
