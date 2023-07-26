package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RaisingPlan struct {
	gorm.Model
	TxHash                 string          `gorm:"column:tx_hash;type:varchar(256);comment:'交易哈希'" json:"tx_hash"`                                     // 交易哈希
	FactoryContract        string          `gorm:"column:factory_contract;type:varchar(256);uniqueIndex:S_R;comment:'工厂合约地址'" json:"factory_contract"` // 工厂合约地址
	RaisingId              string          `gorm:"column:raising_id;type:varchar(256);uniqueIndex:S_R;comment:'募集计划id'" json:"raising_id"`             // 募集计划id
	Raiser                 string          `gorm:"column:raiser;type:varchar(256);comment:'募集人 发起人地址'" json:"raiser"`                                  // 募集人 发起人地址
	RaisingName            string          `gorm:"column:raising_name;type:varchar(256);comment:'募集计划名称'" json:"raising_name"`                         // 募集计划名称
	SponsorCompany         string          `gorm:"column:sponsor_company;type:varchar(1024);comment:'发起单位'" json:"sponsor_company"`                    // 发起单位
	SponsorLogo            string          `gorm:"column:sponsor_logo;type:varchar(1024);comment:'发起单位logo'" json:"sponsor_logo"`                      // 发起单位logo
	MinerId                string          `gorm:"column:miner_id;type:varchar(128);comment:'miner id'" json:"miner_id"`                               // miner id
	MinerType              int8            `gorm:"column:miner_type;type:tinyint(2);comment:'miner 1- 新建节点  2-现有cc转dc'" json:"miner_type"`             // miner 1- 新建节点  2-现有cc转dc
	SectorSize             uint64          `gorm:"column:sector_size;type:bigint;comment:'扇区大小'" json:"sector_size"`                                   // 扇区大小
	SectorPeriod           uint64          `gorm:"column:sector_period;type:bigint;comment:'扇区有效期'" json:"sector_period"`                              // 扇区有效期
	ServiceProviderAddress string          `gorm:"column:service_provider_address;type:varchar(256);comment:'服务商地址'" json:"service_provider_address"`  // 服务商地址
	ServiceId              int32           `gorm:"column:service_id;type:int;comment:'服务商id'" json:"service_id"`                                       // 服务商id
	PlanOpen               int8            `gorm:"column:plan_open;type:tinyint(2);comment:'计划开放类型 1 公募 2 私募'" json:"plan_open"`                       // 计划开放类型 1 公募 2 私募
	TargetAmount           decimal.Decimal `gorm:"column:target_amount;type:decimal(65);comment:'目标金额'" json:"target_amount"`                          // 目标金额
	TargetPower            decimal.Decimal `gorm:"column:target_power;type:decimal(65);comment:'目标算力'" json:"target_power"`                            // 目标算力
	MinRaiseRate           int32           `gorm:"column:min_raise_rate;type:int;comment:'最低募集率'" json:"min_raise_rate"`                               // 最低募集率
	RaiseDays              int32           `gorm:"column:raise_days;type:int;comment:'募集天数'" json:"raise_days"`                                        // 募集天数
	SealDays               int32           `gorm:"column:seal_days;type:int;comment:'封存天数'" json:"seal_days"`                                          // 封存天数

	RaiseSecurityFund     decimal.Decimal `gorm:"column:raise_security_fund;type:decimal(65);comment:'募集保证金'" json:"raise_security_fund"`                 // 募集保证金
	FFIProtocolFeePayMeth int8            `gorm:"column:ffi_protocol_fee_pay_meth;type:tinyint(2);comment:'ffi协议费支付方式'" json:"ffi_protocol_fee_pay_meth"` // ffi协议费支付方式
	FFIProtocolFee        decimal.Decimal `gorm:"column:ffi_protocol_fee;type:decimal(65);comment:'ffi协议费'" json:"ffi_protocol_fee"`                      // ffi协议费
	OpsSecurityFundRate   int32           `gorm:"column:ops_security_fund_rate;type:int;comment:'运维保证金占比'" json:"ops_security_fund_rate"`                 // 运维保证金占比
	OpsSecurityFund       decimal.Decimal `gorm:"column:ops_security_fund;type:decimal(65);comment:'运维保证金'" json:"ops_security_fund"`                     // 运维保证金
	OpsSecurityFundAddr   string          `gorm:"column:ops_security_fund_addr;type:varchar(256);comment:'运维保证金地址'" json:"ops_security_fund_addr"`        // 运维保证金地址

	RaiserCionShare           decimal.Decimal `gorm:"column:raiser_coin_share;type:decimal(20,8);comment:'募集人币份额  70 %'" json:"raiser_coin_share"`             // 募集人币份额  70 %
	OPServerShare             decimal.Decimal `gorm:"column:op_server_share;type:decimal(20,8);comment:'运维币份额   5%'" json:"op_server_share"`                   // 运维币份额   5%
	HisPower                  decimal.Decimal `gorm:"column:his_power;type:decimal(65);comment:'历史算力'" json:"his_power"`                                       // 历史算力
	HisInitialPledge          decimal.Decimal `gorm:"column:his_initial_pledge;type:decimal(65);comment:'历史初始抵押'" json:"his_initial_pledge"`                   // 历史初始抵押
	HisBlance                 decimal.Decimal `gorm:"column:his_blance;type:decimal(65);comment:'历史余额'" json:"his_blance"`                                     // 历史余额
	HisSectorCount            int32           `gorm:"column:his_sector_count;type:int;comment:'历史扇区数'" json:"his_sector_count"`                                // 历史扇区数
	RaiseHisAssetRate         int32           `gorm:"column:raise_his_asset_rate;type:int;comment:'募集人历史资产占比'" json:"raise_his_asset_rate"`                    // 募集人历史资产占比
	RaiseHisPowerRate         int32           `gorm:"column:raise_his_power_rate;type:int;comment:'募集人历史算力占比'" json:"raise_his_power_rate"`                    // 募集人历史算力占比
	RaiseHisInitialPledgeRate int32           `gorm:"column:raise_his_initial_pledge_rate;type:int;comment:'募集人历史质押币占比'" json:"raise_his_initial_pledge_rate"` // 募集人历史质押币占比

	RaiseCreateTime   uint64 `gorm:"column:raise_create_time;type:bigint;comment:'募集计划创建时间'" json:"raise_create_time"`   // 募集计划创建时间
	BeginTime         uint64 `gorm:"column:begin_time;type:bigint;comment:'开始时间'" json:"begin_time"`                     // 开始时间
	ClosingTime       uint64 `gorm:"column:closing_time;type:bigint;comment:'结束时间'" json:"closing_time"`                 // 结束时间
	RaiseEndTime      uint64 `gorm:"column:raise_end_time;type:bigint;default:0;comment:'募集结束时间'" json:"raise_end_time"` // 募集结束时间
	BeginSealTime     uint64 `gorm:"column:begin_seal_time;type:bigint;comment:'开始封存时间'" json:"begin_seal_time"`
	BeginPreSealTime  uint64 `gorm:"column:begin_pre_seal_time;type:bigint;comment:'开始预封存时间'" json:"begin_pre_seal_time"`                 // 开始预封存时间
	EndSealTime       uint64 `gorm:"column:end_seal_time;type:bigint;comment:'结束封存时间'" json:"end_seal_time"`                              // 结束封存时间
	DesignEndSealTime uint64 `gorm:"column:design_end_seal_time;type:bigint;default:0;comment:'设计结束封存时间'" json:"design_end_seal_time"`    // 设计结束封存时间
	DelaySealTime     uint64 `gorm:"column:delay_seal_time;type:bigint;comment:'延迟封存时间'" json:"delay_seal_time"`                          // 延迟封存时间
	Status            int8   `gorm:"column:status;type:tinyint(2);comment:'状态 0:等待开始 1:募集中 2:关闭 3:，募集成功  9:节点结束'" json:"status"`          // 状态 0:募集中 1:募集成功 2:募集失败 3:等待服务商签名 4:封装中 5:运行中 9:节点结束
	Progress          uint64 `gorm:"column:progress;type:bigint;comment:'进度 根据质押币计算'" json:"progress"`                                    // 进度 根据质押币计算
	PowerProgress     uint64 `gorm:"column:power_progress;type:bigint;comment:'进度 根据算力计算'" json:"power_progress"`                         // 进度 根据算力计算
	IncomeRate        int32  `gorm:"column:income_rate;type:int;comment:'收益率'" json:"income_rate"`                                        // 收益率
	SealedStatus      int8   `gorm:"column:sealed_status;type:tinyint(2);comment:'封装开始 0:未开始 1:进行中 2:已完成 3:延长封装'" json:"sealed_status"`   // 封装开始 0:未开始 1:进行中 2:已完成 3:延长封装
	LatestBlockNumber uint64 `gorm:"column:latest_block_number;type:bigint;comment:'最新区块高度'" json:"latest_block_number"`                  // 最新区块高度
	RaiseMarginStatus int8   `gorm:"column:raise_margin_status;type:tinyint(2);comment:'募集保证金状态 0:未缴纳 1:已缴纳'" json:"raise_margin_status"` // 募集保证金状态 0:未缴纳 1:已缴纳
	SPMarginStatus    int8   `gorm:"column:sp_margin_status;type:tinyint(2);comment:'运维保证金状态 0:未缴纳 1:已缴纳'" json:"sp_margin_status"`       // 运维保证金状态 0:未缴纳 1:已缴纳
	SPSignStatus      int8   `gorm:"column:sp_sign_status;type:tinyint(2);comment:'运维签名状态 0:未签名 1:已签名'" json:"sp_sign_status"`            // 运维签名状态 0:未签名 1:已签名

	RaiseAddress string          `gorm:"column:raise_address;type:varchar(256);comment:'募集计划合约地址，上链之后更新'" json:"raise_address"` // 募集计划合约地址，上链之后更新
	ActualAmount decimal.Decimal `gorm:"column:actual_amount;type:decimal(65);comment:'实际募集金额，募集成功后更新'" json:"actual_amount"`   //实际募集金额，募集成功后更新

	// fil_per_tera_day
	FilPerTeraDay decimal.Decimal `gorm:"column:fil_per_tera_day;type:decimal(64,18);default:0;comment:'每T每天收益'" json:"fil_per_tera_day"` // 每T每天收益
	// pledge_per_tera_day
	PledgePerTeraDay decimal.Decimal `gorm:"column:pledge_per_tera_day;type:decimal(64,18);default:0;comment:'每T每天质押币'" json:"pledge_per_tera_day"` // 每T每天质押币

}

// 募集计划各个状态同步的次数
type RaiseSyncCount struct {
	// 募集计划id
	RaiseId string `gorm:"column:raise_id;type:varchar(256);primary_key;comment:'募集计划id'" json:"raise_id"` // 募集计划id
	// 募集开始同步次数
	RaiseBeginSyncCount uint64 `gorm:"column:raise_begin_sync_count;type:bigint;default:0;comment:'募集开始同步次数'" json:"raise_begin_sync_count"` // 募集开始同步次数
	// 募集结束同步次数
	RaiseEndSyncCount uint64 `gorm:"column:raise_end_sync_count;type:bigint;default:0;comment:'募集结束同步次数'" json:"raise_end_sync_count"` // 募集结束同步次数
	// 封装开始同步次数
	SealBeginSyncCount uint64 `gorm:"column:seal_begin_sync_count;type:bigint;default:0;comment:'封装开始同步次数'" json:"seal_begin_sync_count"` // 封装开始同步次数
	// 封装结束同步次数
	SealEndSyncCount uint64 `gorm:"column:seal_end_sync_count;type:bigint;default:0;comment:'封装结束同步次数'" json:"seal_end_sync_count"` // 封装结束同步次数
	// 封装延长T+2同步次数
	SealDelaySyncCount uint64 `gorm:"column:seal_delay_sync_count;type:bigint;default:0;comment:'封装延长T+2同步次数'" json:"seal_delay_sync_count"` // 封装延长T+2同步次数

}
