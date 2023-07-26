package filContractEvent

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/jhyehuang/crontab_server/src/pkg/fevm"
)

func AllTopics() ethtypes.EthTopicSpec {
	topics := []ethtypes.EthHashList{{
		ECreateRaisePlanSign(),
		ECreateAssetPack(),
		SpSignWithMiner(),
		EStartRaisePlan(),
		ECloseRaisePlan(),
		ERaiseSecurityFund(),
		EDepositOPSSecurityFund(), //
		EWithdrawRaiseSecurityFund(),
		EWithdrawOPSSecurityFund(),
		EStackFromInvestor(),
		EUnstackFromInverstor(),
		EUnstackFromInverstorV1(),
		EInvestorWithdraw(),
		ESPWithdraw(),
		ERaiseWithdraw(),
		EPushBlockReward(),
		ERaiseSuccess(), //
		ERaiseFailed(),
		EStartSeal(),
		EPushHistoryAssetPack(),
		ESealProgress(),
		ESealEnd(),
		EPushSpFine(),
		EPushPledgeReleased(),
		ENodeDestroy(),
		ESpecifyOpsPayer(),
		EStartPreSeal(),
		EStartPreSealTransfer(),
		EClosePlanToSeal(),
		EAddOpsSecurityFund(),
		EPushFinalProgress(),
	}}

	return topics
}

func AllTopicsEth() [][]common.Hash {
	topics := [][]common.Hash{{
		common.HexToHash(ECreateRaisePlanSign().String()),
		common.HexToHash(ECreateAssetPack().String()),
		common.HexToHash(SpSignWithMiner().String()),
		common.HexToHash(EStartRaisePlan().String()),
		common.HexToHash(ECloseRaisePlan().String()),
		common.HexToHash(ERaiseSecurityFund().String()),
		common.HexToHash(EDepositOPSSecurityFund().String()),
		common.HexToHash(EWithdrawRaiseSecurityFund().String()),
		common.HexToHash(EWithdrawOPSSecurityFund().String()),
		common.HexToHash(EStackFromInvestor().String()),
		common.HexToHash(EUnstackFromInverstor().String()),
		common.HexToHash(EUnstackFromInverstorV1().String()),
		common.HexToHash(EInvestorWithdraw().String()),
		common.HexToHash(ESPWithdraw().String()),
		common.HexToHash(ERaiseWithdraw().String()),
		common.HexToHash(EPushBlockReward().String()),
		common.HexToHash(ERaiseSuccess().String()),
		common.HexToHash(ERaiseFailed().String()),
		common.HexToHash(EStartSeal().String()),
		common.HexToHash(EPushHistoryAssetPack().String()),
		common.HexToHash(ESealProgress().String()),
		common.HexToHash(ESealEnd().String()),
		common.HexToHash(EPushSpFine().String()),
		common.HexToHash(EPushPledgeReleased().String()),
		common.HexToHash(ENodeDestroy().String()),
		common.HexToHash(ESpecifyOpsPayer().String()),
		common.HexToHash(EStartPreSeal().String()),
		common.HexToHash(EStartPreSealTransfer().String()),
		common.HexToHash(EClosePlanToSeal().String()),
		common.HexToHash(EAddOpsSecurityFund().String()),
		common.HexToHash(EPushFinalProgress().String()),
	}}

	return topics
}

func ECreateRaisePlanSign() ethtypes.EthHash {
	sign := "CreateRaisePlan(uint256,address,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string),(uint64,uint256,uint256,uint256,uint256,uint256,address,uint256),(uint256,uint256,uint256,uint256,uint256))"
	return fevm.EthTopicHash(sign)
}

func ECreateAssetPack() ethtypes.EthHash {

	sign := "CreateRaisePlan(uint256,address,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string),(uint64,uint256,uint256,uint256,uint256,uint256,address,uint256),(uint256,uint256,uint256,uint256,uint256))"
	return fevm.EthTopicHash(sign)
}

// 服务商签名
func SpSignWithMiner() ethtypes.EthHash {
	sign := "SpSignWithMiner(address,uint64,address)"
	return fevm.EthTopicHash(sign)
}

func EStartRaisePlan() ethtypes.EthHash {
	sign := "StartRaisePlan(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

func ECloseRaisePlan() ethtypes.EthHash {
	sign := "CloseRaisePlan(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 缴纳募集计划保证金
func ERaiseSecurityFund() ethtypes.EthHash {
	sign := "DepositSecurityFund(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 缴纳运维保证金
func EDepositOPSSecurityFund() ethtypes.EthHash {
	sign := "DepositOpsSecurityFund(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

func EWithdrawRaiseSecurityFund() ethtypes.EthHash {
	sign := "WithdrawSecurityFund(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

func EWithdrawOPSSecurityFund() ethtypes.EthHash {
	sign := "WithdrawOpsSecurityFund(uint256,address,uint256,uint256,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者质押
func EStackFromInvestor() ethtypes.EthHash {
	sign := "Staking(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者取消质押
func EUnstackFromInverstor() ethtypes.EthHash {
	sign := "Unstaking(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者取消质押
func EUnstackFromInverstorV1() ethtypes.EthHash {
	sign := "Unstaking(uint256,address,address,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者提取收益
func EInvestorWithdraw() ethtypes.EthHash {
	sign := "InvestorWithdraw(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 服务商提取收益
func ESPWithdraw() ethtypes.EthHash {
	sign := "SpWithdraw(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 发起人提取收益
func ERaiseWithdraw() ethtypes.EthHash {
	sign := "RaiseWithdraw(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 向合约压入爆块信息
func EPushBlockReward() ethtypes.EthHash {
	sign := "PushBlockReward(uint256,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 募集失败
func ERaiseFailed() ethtypes.EthHash {
	sign := "RaiseFailed(uint256)"
	return fevm.EthTopicHash(sign)
}

// 募集成功
func ERaiseSuccess() ethtypes.EthHash {
	sign := "RaiseSuccess(uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 开始封装
func EStartSeal() ethtypes.EthHash {
	sign := "StartSeal(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 推送历史资产包资料
func EPushHistoryAssetPack() ethtypes.EthHash {
	sign := "PushOldAssetPackValue(uint256,address,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 封装延期  封装进度推送
func ESealProgress() ethtypes.EthHash {
	sign := "PushSealProgress(uint256,uint256,uint8)"
	return fevm.EthTopicHash(sign)
}

// 封装结束
func ESealEnd() ethtypes.EthHash {
	sign := "SealEnd(uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 运维罚金
func EPushSpFine() ethtypes.EthHash {
	sign := "PushSpFine(uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 质押币释放推送
func EPushPledgeReleased() ethtypes.EthHash {
	sign := "PushPledgeReleased(uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 节点销毁
func ENodeDestroy() ethtypes.EthHash {
	sign := "DestroyNode(uint256,uint8)"
	return fevm.EthTopicHash(sign)
}

// 指定运维付款人
func ESpecifyOpsPayer() ethtypes.EthHash {
	sign := "eSpecifyOpsPayer(uint256,address,address,address)"
	return fevm.EthTopicHash(sign)
}

// 开始预封装
func EStartPreSeal() ethtypes.EthHash {
	sign := "StartPreSeal(uint256,address)"
	return fevm.EthTopicHash(sign)
}

// 开始转移质押币到miner
func EStartPreSealTransfer() ethtypes.EthHash {
	sign := "SendToMiner(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 关闭计划到封装
func EClosePlanToSeal() ethtypes.EthHash {
	sign := "ClosePlanToSeal(uint256,address,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 追加运维保证金
func EAddOpsSecurityFund() ethtypes.EthHash {
	sign := "AddOpsSecurityFund(uint256,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 推送最终封装进度
func EPushFinalProgress() ethtypes.EthHash {
	sign := "PushFinalProgress(uint256,uint256,uint8)"
	return fevm.EthTopicHash(sign)
}
