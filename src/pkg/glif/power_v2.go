package glif

import (
	"bytes"
	"context"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/jhyehuang/crontab_server/src/pkg/timeUtils"

	//"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
)

func GetMinerPowerV2(minerId string, currEpoch abi.ChainEpoch) (mbi *api.MiningBaseInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to get miner address: %s", err)
		return nil, err
	}
	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	mActor, err := fevmApi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to get miner actor: %s", err)
		return nil, err
	}

	//actorHead, err := service.RequestChainReadObj(ma.Head.String())
	actorHead, err := fevmApi.ChainReadObj(ctx, mActor.Head)
	if err != nil {
		log.Errorf("failed to get miner actor head: %s", err)
		return nil, err
	}

	var minerState miner.State

	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		log.Errorf("failed to unmarshal miner actor head: %s", err)
		return nil, err
	}

	actorInfo, err := fevmApi.ChainReadObj(ctx, minerState.Info)
	if err != nil {
		log.Errorf("failed to get miner actor info: %s", err)
		return nil, err
	}
	var minerInfo miner.MinerInfo

	if err := minerInfo.UnmarshalCBOR(bytes.NewReader(actorInfo)); err != nil {
		log.Errorf("failed to get miner actor info: %s", err)
		return nil, err
	}

	return nil, nil

}

func GetMinerBaseInfoV2(minerId string, currEpoch abi.ChainEpoch) (mbi *api.MiningBaseInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to get miner address: %s", err)
		return nil, err
	}
	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	mBaseInfo, err := fevmApi.MinerGetBaseInfo(context.Background(), mineraddr, currHead.Height(), tsk)
	if err != nil {
		log.Errorf("failed to get miner actor: %s", err)
		return nil, err
	}

	if mBaseInfo == nil {
		mBaseInfo = &api.MiningBaseInfo{
			MinerPower: chaintypes.BigInt{
				Int: big.NewInt(0),
			},
			NetworkPower: chaintypes.BigInt{
				Int: big.NewInt(0),
			},
		}

	}

	return mBaseInfo, nil

}

func GetProvingEpoch(minerId string, currEpoch abi.ChainEpoch) (abi.ChainEpoch, error) {
	mState, err := GetMinerStateV2(minerId, currEpoch)
	if err != nil {
		log.Errorf("failed to get miner state: %s", err)
		return -1, err
	}
	return mState.ProvingPeriodStart, nil
}

func GetVestFundsV2(minerId string, currEpoch abi.ChainEpoch) (*miner.VestingFunds, error) {
	mState, err := GetMinerStateV2(minerId, currEpoch)
	if err != nil {
		log.Errorf("failed to get miner state: %s", err)
		return nil, err
	}

	fundsBytes, err := fevmApi.ChainReadObj(ctx, mState.VestingFunds)
	if err != nil {
		log.Errorf("failed to get miner vesting funds: %s", err)
		return nil, err
	}

	var vestFunds miner.VestingFunds
	if err := vestFunds.UnmarshalCBOR(bytes.NewReader(fundsBytes)); err != nil {
		log.Errorf("failed to unmarshal miner vesting funds: %s", err)
		return nil, err
	}
	return &vestFunds, nil
}

func GetMinerStateV2(minerId string, currEpoch abi.ChainEpoch) (mbs *miner.State, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to StateGetActor: %s", err)
		return nil, err
	}

	actorHead, err := fevmApi.ChainReadObj(ctx, ma.Head)
	if err != nil {
		return nil, err
	}
	//log.Infof("actorHead: %s", actorHead)

	var minerState miner.State

	// 可能会获取到一个 其他的 actor 的信息，非miner的actor信息不能解析出来
	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		log.Errorf("failed to UnmarshalCBOR: %s", err)
		return nil, err
	}

	return &minerState, nil

}

func GetMinerActorV2(minerId string, currEpoch abi.ChainEpoch) (actor *chaintypes.Actor, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to StateGetActor: %s", err)
		return nil, err
	}

	return ma, nil

}

func GetMinerAvailableBalanceV2(minerId string, currEpoch abi.ChainEpoch) (actor *chaintypes.BigInt, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateMinerAvailableBalance(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to StateGetActor: %s", err)
		return nil, err
	}

	return &ma, nil

}

func GetMinerSectorsV2(minerId string, currEpoch abi.ChainEpoch) (mbi *api.MinerSectors, err error) {

	log.Debug("GetMinerSectors start")
	log.Infof("minerId: %s\n", minerId)
	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr.String())
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	//currEpoch := currHead.Height()

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateGetActor(context.Background(), mineraddr, tsk)
	if err != nil {
		return nil, err
	}

	//actorHead, err := service.RequestChainReadObj(ma.Head.String())
	actorHead, err := fevmApi.ChainReadObj(ctx, ma.Head)
	if err != nil {
		return nil, err
	}

	var minerState miner.State

	if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
		return nil, err
	}

	actorSectors, err := fevmApi.StateMinerSectorCount(ctx, mineraddr, tsk)
	if err != nil {
		return nil, err
	}
	log.Infof("actorSectors: %+v", actorSectors)

	return &actorSectors, nil

}

func GetChainHeadV2(currEpoch abi.ChainEpoch) (mbs *chaintypes.TipSet, err error) {

	currHead, err := fevmApi.ChainGetTipSetByHeight(ctx, currEpoch, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}

	return currHead, nil

}

func ChainGetGenesis() (mbs *chaintypes.TipSet, err error) {
	currHead, err := fevmApi.ChainGetGenesis(ctx)
	if err != nil {
		log.Errorf("failed to get chain head: %s", err)
		return nil, err
	}
	return currHead, nil
}

func GetTimeByEpoch(epoch uint64) (uint64, error) {
	gen, err := ChainGetGenesis()
	if err != nil {
		return 0, err
	}
	return timeUtils.GetTimeByEpoch(uint64(epoch), gen.MinTimestamp())
}
