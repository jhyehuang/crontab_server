package models

import "gorm.io/gorm"

type PushContractTxRecord struct {
	gorm.Model
	// 交易hash
	TxHash string `gorm:"column:tx_hash;type:varchar(255);not null;uniqueIndex:idx_tx_hash" json:"tx_hash"`
	// 调用方法
	Method string `gorm:"column:method;type:varchar(255);not null;" json:"method"`
	// 发起人
	From string `gorm:"column:from;type:varchar(255);not null;" json:"from"`
	//目标地址
	To string `gorm:"column:to;type:varchar(255);not null;" json:"to"`
	// 交易金额
	Value string `gorm:"column:value;type:varchar(255);not null;" json:"value"`
	// 参数
	Params string `gorm:"column:params;type:varchar(255);not null;" json:"params"`
	// 交易状态
	Status int `gorm:"column:status;type:int;not null;" json:"status"` // 0:未确认 1:已确认 2:失败
	// 高度
	Height uint64 `gorm:"column:height;type:bigint;not null;" json:"height"`
	// nonce
	Nonce uint64 `gorm:"column:nonce;type:bigint;not null;" json:"nonce"`
}
