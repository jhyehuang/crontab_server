package controller

import "github.com/shopspring/decimal"

type RaisingPlanResp struct {
	//募集计划
	TxHash                 string          ` json:"tx_hash"`                  // 交易哈希
	FactoryContract        string          ` json:"factory_contract"`         // 工厂合约地址
	RaisingId              string          ` json:"raising_id"`               // 募集计划id
	Raiser                 string          ` json:"raiser"`                   // 募集人 发起人地址
	RaisingName            string          ` json:"raising_name"`             // 募集计划名称
	SponsorCompany         string          ` json:"sponsor_company"`          // 发起单位
	SponsorLogo            string          ` json:"sponsor_logo"`             // 发起单位logo
	MinerId                string          ` json:"miner_id"`                 // miner id
	MinerType              int8            ` json:"miner_type"`               // miner 1- 新建节点  2-现有cc转dc
	SectorSize             uint64          ` json:"sector_size"`              // 扇区大小
	SectorPeriod           uint64          ` json:"sector_period"`            // 扇区有效期
	ServiceProviderAddress string          ` json:"service_provider_address"` // 服务商地址
	ServiceId              int32           ` json:"service_id"`               // 服务商id
	PlanOpen               int8            ` json:"plan_open"`                // 计划开放类型 1 公募 2 私募
	TargetAmount           decimal.Decimal ` json:"target_amount"`            // 目标金额
	TargetPower            decimal.Decimal ` json:"target_power"`             // 目标算力
	MinRaiseRate           int32           ` json:"min_raise_rate"`           // 最低募集率
	RaiseDays              int32           ` json:"raise_days"`               // 募集天数
	SealDays               int32           ` json:"seal_days"`                // 封存天数

	RaiseSecurityFund     decimal.Decimal ` json:"raise_security_fund"`       // 募集保证金
	FFIProtocolFeePayMeth int8            ` json:"ffi_protocol_fee_pay_meth"` // ffi协议费支付方式
	FFIProtocolFee        decimal.Decimal ` json:"ffi_protocol_fee"`          // ffi协议费
	OpsSecurityFundRate   int32           ` json:"ops_security_fund_rate"`    // 运维保证金占比
	OpsSecurityFund       decimal.Decimal ` json:"ops_security_fund"`         // 运维保证金
	OpsSecurityFundAddr   string          ` json:"ops_security_fund_addr"`    // 运维保证金地址

	RaiserCionShare           decimal.Decimal ` json:"raiser_coin_share"`             // 募集人币份额  70 %
	OPServerShare             decimal.Decimal ` json:"op_server_share"`               // 运维币份额   5%
	HisPower                  decimal.Decimal ` json:"his_power"`                     // 历史算力
	HisInitialPledge          decimal.Decimal ` json:"his_initial_pledge"`            // 历史初始抵押
	HisBlance                 decimal.Decimal ` json:"his_blance"`                    // 历史余额
	HisSectorCount            int32           ` json:"his_sector_count"`              // 历史扇区数
	RaiseHisAssetRate         int32           ` json:"raise_his_asset_rate"`          // 募集人历史资产占比
	RaiseHisPowerRate         int32           ` json:"raise_his_power_rate"`          // 募集人历史算力占比
	RaiseHisInitialPledgeRate int32           ` json:"raise_his_initial_pledge_rate"` // 募集人历史初始抵押占比

	RaiseCreateTime   uint64 ` json:"raise_create_time"`   // 募集计划创建时间
	BeginTime         uint64 ` json:"begin_time"`          // 开始时间
	ClosingTime       uint64 ` json:"closing_time"`        // 结束时间
	BeginSealTime     uint64 ` json:"begin_seal_time"`     // 开始封存时间
	BeginPreSealTime  uint64 ` json:"begin_pre_seal_time"` // 开始预封存时间
	EndSealTime       uint64 ` json:"end_seal_time"`       // 结束封存时间
	DelaySealTime     uint64 ` json:"delay_seal_time"`     // 延迟封存时间
	Status            int8   ` json:"status"`              // 状态 0:募集中 1:募集成功 2:募集失败 3:等待服务商签名 4:封装中 5:运行中 9:节点结束
	Progress          uint64 ` json:"progress"`            // 进度 根据质押币计算
	PowerProgress     uint64 ` json:"power_progress"`      // 进度 根据算力计算
	IncomeRate        int32  ` json:"income_rate"`         // 收益率
	SealedStatus      int8   ` json:"sealed_status"`       // 封装开始 0:未开始 1:进行中 2:已完成 3:延长封装
	LatestBlockNumber uint64 ` json:"latest_block_number"` // 最新区块高度
	RaiseMarginStatus int8   ` json:"raise_margin_status"` // 募集保证金状态 0:未缴纳 1:已缴纳
	SPMarginStatus    int8   ` json:"sp_margin_status"`    // 运维保证金状态 0:未缴纳 1:已缴纳
	SPSignStatus      int8   ` json:"sp_sign_status"`      // 运维签名状态 0:未签名 1:已签名

	RaiseAddress string          ` json:"raise_address"` // 募集计划合约地址，上链之后更新
	ActualAmount decimal.Decimal ` json:"actual_amount"` //实际募集金额，募集成功后更新

	FilPerTeraDay    decimal.Decimal ` json:"fil_per_tera_day"`    //每T每天收益
	PledgePerTeraDay decimal.Decimal ` json:"pledge_per_tera_day"` //每T每天抵押
}

type RaiseAssetReport struct {
	TotalReward    string `json:"total_reward"`     //总奖励
	TotalSpPenalty string `json:"total_sp_penalty"` //总sp罚金
}

type AddressInPlanResp struct {
	//募集计划id
	RaisingId uint64 `json:"raising_id,string"` // 募集计划id
	//合约地址
	RaiseAddress string `json:"raise_address"` // 募集计划合约地址
	//minerID
	MinerId string `json:"miner_id"` // miner

}
