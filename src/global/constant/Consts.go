package constant

import "github.com/filecoin-project/lotus/build"

const (
	ENV_MODE = "NETWORK_ENV"

	PageSize        = "pageSize"
	PageNum         = "pageNum"
	DefaultPageNum  = "1"
	DefaultPageSize = "10"
	TimeStart       = "2006-01-02T15:04:05Z"
	SimpleTimeStart = "2006-01-02"

	RewardVestReleaseDays = 180

	// StatusEnable 状态  0：禁用 1：启用
	StatusEnable  = 1
	StatusDisable = 0

	//Rpc
	RpcHeaderKey         = "Authorization"
	RpcHeaderTokenPrefix = "Bearer "

	RpcMethodWithdraw         = "16"
	RpcMethodSend             = "0"
	RpcMethodWithdrawCategory = 0
	RpcMethodSendCategory     = 1

	// 募集计划 状态  9:节点结束 10:等待上链
	RAISING_PLAN_STATUS_NODE_END   = 9
	RAISING_PLAN_STATUS_WAIT_CHAIN = 10 //10:等待上链

	// 资产包状态 	0:未启用 1:封装中 2:封装完成  3.：延长封装 9: 结束
	ASSET_PACKAGE_STATUS_NOT_START = 0
	ASSET_PACKAGE_STATUS_SEALING   = 1
	ASSET_PACKAGE_STATUS_SEALED    = 2
	ASSET_PACKAGE_STATUS_DELAY     = 3
	ASSET_PACKAGE_STATUS_END       = 9

	// 资产包类型 0:普通资产包 1:募集资产包
	ASSET_PACKAGE_TYPE_NORMAL = 0
	ASSET_PACKAGE_TYPE_RAISE  = 1

	// 资金类型
	FUND_TYPE_RAISE_SECURITY_FUND = 1 // 募集保证金
	FUND_TYPE_OPS_SECURITY_FUND   = 2 // 运维保证金
	FUND_TYPE_RAISE_FUND          = 3 // 募集资金
	FUND_TYPE_PROFIT              = 4 // 收益

	// 交易类型
	TX_TYPE_PLEDGE   = 1 // 质押
	TX_TYPE_REDEEM   = 2 // 赎回
	TX_TYPE_WITHDRAW = 3 // 提取收益
	TX_TYPE_SEND     = 4 // 转账
	TX_TYPE_ADD      = 5 // 追加

	// 资金支付状态
	StatusWaitingPay = 0
	StatusPayed      = 1

	//是否需要管理miner扇区
	IsSyncSectorNo   = 0
	IsSyncSectorTrue = 1
)
const CALC_ACCURACY = 1e6
const SealedPlageAmt = 7 * build.FilecoinPrecision
const SecPerDay = 86400
const RewardEachBlock = 7 * build.FilecoinPrecision
const RaisePlanCheckSQL = " status in (0,1,2,3,5) "
const DefaultEndEpoch = 9999999999

// 间隔高度
const FactoryCheckHeightInterval = 1000

const (
	RaiseStateWaitingStart = iota
	RaiseStateRaising
	RaiseStateClosed
	RaiseStateSuccess
	RaiseStateFailure
)

const (
	NodeStateWaitingStart = iota
	NodeStateStarted
	NodeStateDelayed
	NodeStateEnd
	NodeStateDestroy
	NodeStatePreSeal = 11
)

const (
	AssetPackSyncSectorRunning = iota
	AssetPackSyncSectorEnd
)

const (
	SectorTypeCC = iota
	SectorTypeDC
)

const (
	//无效
	SectorStatusDisable = 0
	//有效
	SectorStatusEnable = 1
	//待释放
	SectorStatusWaitRelease = 2
)

const (
	ChainPenaltyReleaseTrue  = true
	ChainPenaltyReleaseFalse = false

	ChainRewardReleaseTrue  = 1
	ChainRewardReleaseFalse = 0
)

const (
	// 0:未确认 1:已确认 2:失败
	ChainTxStatusUnconfirmed = 0
	ChainTxStatusConfirmed   = 1
	ChainTxStatusFailed      = 2
)
