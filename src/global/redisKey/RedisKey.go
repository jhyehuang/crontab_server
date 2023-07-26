package redisKey

const (
	TokenHeaderKey = "Authorization"

	// UserTokenKey 用户token
	UserTokenKey         = "login-token:%s"
	UserTokenExpireHours = 7 * 24
	UserTokenValue       = "userId"

	// CnyPriceKey 人民币价格每日缓存
	CnyPriceKey = "cny-price-cache:%s"
	// CnyPriceTodayKey 当前价格
	CnyPriceCurrentKey = "cny-price:current"
	CnyPricePrevKey    = "cny-price:prev"

	// miner 最小扇区号
	MinerMinSectorNum = "miner-min-sector-num:%s"
	// miner 最大扇区号
	MinerMaxSectorNum = "miner-max-sector-num:%s"

	//同步扇区锁
	SyncSectorLockKey     = "sector-sync-lock"
	SyncPenaltyLockKey    = "penalty-sync-lock"
	SyncRewardLockKey     = "reward-sync-lock"
	SyncTerminatedLockKey = "terminated-sync-lock"

	// 同步Penalty表最后一次id（同步时最大id）
	SyncPenaltyLastId = "sync-penalty-last-id"

	//老资产包线性释放同步锁 - 只同步一次 minerId-assetPackId
	OldPackVestLock = "old-pack-vest-lock:%s-%d"

	//已终止或释放的扇区编号
	SectorNoExist = "disabled-sector-no:%s"
)

const (
	SyncFilRaiseDealLineLock           = "sync-fil-raise-deal-line-lock"
	SyncMinerSealedProgressLock        = "sync-miner-sealed-progress-lock"
	SyncMinerSealedEndLock             = "sync-miner-sealed-end-lock"
	SyncAssetPackDestroyLock           = "sync-asset-pack-destroy-lock"
	SyncFilBlockRewardDistributionLock = "sync-fil-block-reward-distribution-lock"
	SyncFilPledgeReleaseLock           = "sync-fil-pledge-release-lock"
	SyncSPPenaltyLock                  = "sync-sp-penalty-lock"
	SyncContractDataLock               = "sync-contract-data-lock"
	SyncCheckRaiseEventLock            = "sync-check-raise-event-lock"
)
