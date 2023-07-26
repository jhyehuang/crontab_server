package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type FundDetails struct {
	gorm.Model
	TxHash      string          `gorm:"column:tx_hash;type:varchar(256);uniqueIndex:S_R;comment:'交易哈希'" json:"tx_hash"`                          // 交易哈希
	FundType    int8            `gorm:"column:fund_type;type:tinyint(2);uniqueIndex:S_R;comment:'资金类型 1 募集保证金 2 运维保证金 3 募集资金'" json:"fund_type"` // 资金类型 1 募集保证金 2 运维保证金 3 募集资金
	TxType      int8            `gorm:"column:tx_type;type:tinyint(2);uniqueIndex:S_R;comment:'交易类型 1 质押 2 赎回 3 提取收益'" json:"tx_type"`           // 交易类型 1 质押 2 赎回 3 提取收益
	FromAddress string          `gorm:"column:from_address;type:varchar(256);comment:'from'" json:"from_address"`                                // from
	ToAddress   string          `gorm:"column:to_address;type:varchar(256);comment:'to'" json:"to_address"`                                      // to
	Value       decimal.Decimal `gorm:"column:value;type:decimal(65);comment:'交易金额'" json:"value"`                                               // 交易金额
	RaiserId    uint64          `gorm:"column:raiser_id;type:bigint;comment:'募集计划id'" json:"raiser_id"`                                          // 募集计划id
	AssetPackId uint64          `gorm:"column:asset_pack_id;type:bigint;comment:'资产包id'" json:"asset_pack_id,string"`                            // 资产包id
	TxTime      uint64          `gorm:"column:tx_time;type:bigint;comment:'交易时间'" json:"tx_time"`                                                // 交易时间
}
