package models

import (
	"github.com/jhyehuang/crontab_server/src/global/constant"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RaisingAssetPack struct {
	gorm.Model
	ManagerAddress      string          `gorm:"column:manager_address;type:varchar(128);uniqueIndex:S_R;comment:'资产Manager address，合约地址'" json:"manager_address"` // 资产Manager address，合约地址
	MinerID             string          `gorm:"column:miner_id;type:varchar(128);comment:'Miner ID'" json:"miner_id" `                                            // Miner ID
	AssetPackId         uint64          `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R;comment:'资产包id'" json:"asset_pack_id,string" `                    // 资产包id
	AssetPackName       string          `gorm:"column:asset_pack_name;type:varchar(256);comment:'资产包名称'" json:"asset_pack_name"`                                  // 资产包名称
	RaisingId           uint64          `gorm:"column:raising_id;type:bigint;comment:'募集计划id'" json:"raising_id,string" `                                         // 募集计划id
	AssetPackStartEpoch uint64          `gorm:"column:asset_pack_start_epoch;type:bigint;comment:'资产包开始epoch'" json:"asset_pack_start_epoch" `                    // 资产包开始epoch
	AssetPackEndEpoch   uint64          `gorm:"column:asset_pack_end_epoch;type:bigint;comment:'资产包结束epoch'" json:"asset_pack_end_epoch" `                        // 资产包结束epoch
	AssetPackStatus     int8            `gorm:"column:asset_pack_status;type:tinyint(2);comment:'资产包状态 0:未开始 1:进行中 2:已完成 3:延长封装'" json:"asset_pack_status" `      // 资产包状态 0:未开始 1:进行中 2:已完成 3:延长封装
	AssetPackType       int8            `gorm:"column:asset_pack_type;type:tinyint(2);comment:'资产包类型 0:普通资产包 1:募集计划资产包'" json:"asset_pack_type" `                 // 资产包类型 0:普通资产包 1:募集计划资产包
	TotalPledgeAmount   decimal.Decimal `gorm:"column:total_pledge_amount;type:decimal(65);comment:'总质押金额'" json:"total_pledge_amount" `                          // 总质押金额
	TotalPower          decimal.Decimal `gorm:"column:total_power;type:decimal(65);comment:'总算力'" json:"total_power" `                                            // 总算力
	TotalSector         uint64          `gorm:"column:total_sector;type:bigint;comment:'总扇区数'" json:"total_sector" `                                              // 总扇区数
	MinExpirationEpoch  uint64          `gorm:"column:min_expiration_epoch;type:bigint;comment:'最小到期epoch'" json:"min_expiration_epoch" `                         // 最小到期epoch
	MaxExpirationEpoch  uint64          `gorm:"column:max_expiration_epoch;type:bigint;comment:'最大到期epoch'" json:"max_expiration_epoch" `                         // 最大到期epoch
	SectorSize          uint64          `gorm:"column:sector_size;type:bigint;comment:'扇区大小'" json:"sector_size" `                                                // 扇区大小
	SyncSectorStatus    uint8           `gorm:"column:sync_sector_status;type:tinyint(1);comment:'同步扇区状态 0:正在同步 1:同步完成'" json:"sync_sector_status" `              // 同步扇区状态 0:正在同步 1:同步完成
	StatHeight          uint64          `gorm:"column:stat_height;type:bigint;comment:'统计高度'" json:"stat_height" `                                                // 统计高度
}

// 资产包和扇区对应关系表
type AssetPackSector struct {
	gorm.Model
	MinerID               string          `gorm:"column:miner_id;type:varchar(128);comment:'Miner ID'" json:"miner_id"`                            // Miner ID
	AssetPackId           uint64          `gorm:"column:asset_pack_id;type:bigint;unique_index:S_R;comment:'资产包id'" json:"asset_pack_id,string"`   // 资产包id
	SectorNumber          uint64          `gorm:"column:sector_number;type:bigint;unique_index:S_R;comment:'扇区号'" json:"sector_number"`            // 扇区号
	SectorType            uint8           `gorm:"column:sector_type;type:tinyint(1);comment:'扇区类型 0:cc 1:dc'" json:"sector_type"`                  // 扇区类型 0:cc 1:dc
	SealedCID             string          `gorm:"column:sealed_cid;type:text;comment:'扇区cid'" json:"sealed_cid"`                                   // 扇区cid
	DealIDs               string          `gorm:"column:deal_ids;type:text;comment:'交易id'" json:"deal_ids"`                                        // 交易id
	Activation            uint64          `gorm:"column:activation;type:bigint;comment:'激活epoch'" json:"activation"`                               // 激活epoch
	Expiration            uint64          `gorm:"column:expiration;type:bigint;index:idx_expiration;comment:'过期epoch'" json:"expiration"`          // 过期epoch
	DealWeight            decimal.Decimal `gorm:"column:deal_weight;type:decimal(65);comment:'交易权重'" json:"deal_weight"`                           // 交易权重
	VerifiedDealWeight    decimal.Decimal `gorm:"column:verified_deal_weight;type:decimal(65);comment:'交易权重'" json:"verified_deal_weight"`         // 交易权重
	InitialPledge         decimal.Decimal `gorm:"column:initial_pledge;type:decimal(65);comment:'初始质押'" json:"initial_pledge"`                     // 初始质押
	OldInitialPledge      decimal.Decimal `gorm:"column:old_initial_pledge;type:decimal(65);comment:'旧初始质押'" json:"old_initial_pledge"`            // 旧初始质押
	ExpectedDayReward     decimal.Decimal `gorm:"column:expected_day_reward;type:decimal(65);comment:'预计每日奖励'" json:"expected_day_reward"`         // 预计每日奖励
	ExpectedStoragePledge decimal.Decimal `gorm:"column:expected_storage_pledge;type:decimal(65);comment:'预计存储质押'" json:"expected_storage_pledge"` // 预计存储质押
	ReplacedSectorAge     uint64          `gorm:"column:replaced_sector_age;type:varchar(256);comment:'替换扇区年龄'" json:"replaced_sector_age"`        // 替换扇区年龄
	ReplacedDayReward     decimal.Decimal `gorm:"column:replaced_day_reward;type:decimal(65);comment:'替换扇区每日奖励'" json:"replaced_day_reward"`       // 替换扇区每日奖励
	SimpleQAPower         bool            `gorm:"column:simple_qa_power;type:tinyint(1);comment:'简单qa power'" json:"simple_qa_power"`              // 简单qa power
	Power                 decimal.Decimal `gorm:"column:power;type:decimal(65);comment:'算力'" json:"power"`                                         // 算力
	DealExpiration        uint64          `gorm:"column:deal_expiration;type:bigint;comment:'交易过期epoch'" json:"deal_expiration"`                   // 交易过期epoch
	Status                uint8           `gorm:"column:status;type:tinyint(1);comment:'是否有效:0:无效，1:有效，2:待释放'" json:"status"`                      // 是否有效
}

var AssetPackSectorName = "asset_pack_sector"

func (aps *AssetPackSector) SetSectorType(verifiedDealWeight uint64) {
	if verifiedDealWeight > 0 {
		aps.SectorType = constant.SectorTypeDC
	} else {
		aps.SectorType = constant.SectorTypeCC
	}
}

// 扇区释放记录表
type SectorReleaseRecord struct {
	gorm.Model
	MinerId       string          `gorm:"column:miner_id;type:varchar(128);comment:'Miner ID'" json:"miner_id"`                          // Miner ID
	AssetPackId   uint64          `gorm:"column:asset_pack_id;type:bigint;unique_index:S_R;comment:'资产包id'" json:"asset_pack_id,string"` // 资产包id
	SectorNumber  uint64          `gorm:"column:sector_number;type:bigint;unique_index:S_R;comment:'扇区号'" json:"sector_number"`          // 扇区号
	ReleaseEpoch  uint64          `gorm:"column:release_epoch;type:bigint;comment:'释放epoch'" json:"release_epoch"`                       // 释放epoch
	ReleaseAmount decimal.Decimal `gorm:"column:release_amount;type:decimal(65);comment:'释放质押币数量'" json:"release_amount"`                // 释放质押币数量
}

// 扇区封装进度表
type SectorSealedProgress struct {
	gorm.Model
	AssetPackId    uint64          `gorm:"column:asset_pack_id;type:bigint;comment:'资产包id'" json:"asset_pack_id,string"`   // 资产包id
	SealedEpoch    uint64          `gorm:"column:sealed_epoch;type:bigint;comment:'封装epoch'" json:"sealed_epoch"`          // 封装epoch
	SealedAmount   decimal.Decimal `gorm:"column:sealed_amount;type:decimal(65);comment:'封装金额'" json:"sealed_amount"`      // 封装金额
	UnsealedAmount decimal.Decimal `gorm:"column:unsealed_amount;type:decimal(65);comment:'未封装金额'" json:"unsealed_amount"` // 未封装金额
}

// 客户和资产包关系表
type CustomerAssetPack struct {
	gorm.Model
	Address string `gorm:"column:address;type:varchar(128);uniqueIndex:S_R;comment:'客户地址'" json:"address"` // 客户地址

	AssetPackId uint64 `gorm:"column:asset_pack_id;type:bigint;uniqueIndex:S_R;comment:'资产包id'" json:"asset_pack_id,string"` // 资产包id
	MinerId     string `gorm:"column:miner_id;type:varchar(128);comment:'矿工id'" json:"miner_id"`                             // 矿工id
	ProviderId  int32  `gorm:"column:provider_id;type:int;comment:'服务商id'" json:"provider_id"`                               // 服务商id
	Sponsor     string `gorm:"column:sponsor;type:varchar(128);comment:'发起人'" json:"sponsor"`                                // 发起人

}

// dc交易表
type DCDeal struct {
	gorm.Model

	MinerId              string          `gorm:"column:miner_id;type:varchar(128);comment:'矿工id'" json:"miner_id"`                                    // 矿工id
	StartEpoch           uint64          `gorm:"column:start_epoch;type:bigint;comment:'开始epoch'" json:"start_epoch"`                                 // 开始epoch
	EndEpoch             uint64          `gorm:"column:end_epoch;type:bigint;comment:'结束epoch'" json:"end_epoch"`                                     // 结束epoch
	Client               string          `gorm:"column:client;type:varchar(128);comment:'客户'" json:"client"`                                          // 客户
	PieceSize            uint64          `gorm:"column:piece_size;type:bigint;comment:'piece大小'" json:"piece_size"`                                   // piece大小
	StoragePricePerEpoch decimal.Decimal `gorm:"column:storage_price_per_epoch;type:decimal(65);comment:'每epoch存储价格'" json:"storage_price_per_epoch"` // 每epoch存储价格
	SectorStartEpoch     uint64          `gorm:"column:sector_start_epoch;type:bigint;comment:'扇区开始epoch'" json:"sector_start_epoch"`                 // 扇区开始epoch
	LastUpdatedEpoch     uint64          `gorm:"column:last_updated_epoch;type:bigint;comment:'最后更新epoch'" json:"last_updated_epoch"`                 // 最后更新epoch
	SlashEpoch           uint64          `gorm:"column:slash_epoch;type:bigint;comment:'SlashEpoch'" json:"slash_epoch"`                              // SlashEpoch
	VerifiedClaim        uint64          `gorm:"column:verified_claim;type:bigint;comment:'verified_claim'" json:"verified_claim"`                    // verified_claim
}
