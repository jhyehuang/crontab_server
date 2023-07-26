package glif

import (
	"bytes"
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/jhyehuang/crontab_server/src/pkg/log"

	//"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	chaintypes "github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func GetMinerPower(miner string) (mbi *api.MinerPower, err error) {

	mineraddr, err := address.NewFromString(miner)
	if err != nil {
		return nil, err
	}

	//currHead, err := lapi.ChainHead(ctx)

	mbi, err = fevmApi.StateMinerPower(ctx, mineraddr, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("failed to get mining base info: %s", err)
		err = xerrors.Errorf("failed to get mining base info: %w", err)
		return nil, err
	}
	return mbi, nil

}

func GetMinerPowerV1(minerId string) (mbi *api.MiningBaseInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to get miner address: %s", err)
		return nil, err
	}
	currHead, err := fevmApi.ChainHead(ctx)
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

func GetMinerBaseInfo(minerId string) (mbi *api.MiningBaseInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to get miner address: %s", err)
		return nil, err
	}
	currHead, err := fevmApi.ChainHead(ctx)
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
	//log.Warnf("mBaseInfo: %+v,%+v", mBaseInfo, err)

	if mBaseInfo == nil {
		mBaseInfo = &api.MiningBaseInfo{}
	}

	return mBaseInfo, nil

}

func GetMinerState(minerId string) (mbs *miner.State, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		log.Errorf("failed to ChainHead: %s", err)
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

func GetMinerActor(minerId string) (actor *chaintypes.Actor, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		log.Errorf("failed to ChainHead: %s", err)
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

func GetMinerAvailableBalance(minerId string) (actor *chaintypes.BigInt, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		log.Errorf("failed to ChainHead: %s", err)
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

func GetMinerSectors(minerId string) (mbi *api.MinerSectors, err error) {

	log.Debug("GetMinerSectors start")
	log.Infof("minerId: %s\n", minerId)
	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr.String())
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
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

func GetChainHead() (mbs *chaintypes.TipSet, err error) {

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		err = xerrors.Errorf("failed to ChainHead: %w", err)
		log.Infof("failed to ChainHead: %s", err)
		return nil, err
	}

	return currHead, nil

}

func GetBlockHead(cid cid.Cid) (*chaintypes.BlockHeader, error) {

	mbi, err := fevmApi.ChainGetBlock(ctx, cid)
	if err != nil {
		err = xerrors.Errorf("failed to ChainGetBlock: %w", err)
		log.Infof("failed to ChainGetBlock: %s", err)
		return nil, err
	}
	return mbi, nil

}

func GetBlockHeadEth(cid string) (*ethtypes.EthBlock, error) {

	mbi, err := fevmApi.EthGetBlockByNumber(ctx, cid, true)
	if err != nil {
		err = xerrors.Errorf("failed to ChainGetBlock: %w", err)
		log.Infof("failed to ChainGetBlock: %s", err)
		return nil, err
	}
	return &mbi, nil

}

func GetCurrentHeight() (int64, error) {
	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		return -1, err
	}
	return int64(currHead.Height()), nil
}

func GetStateMinerInfo(minerId string) (actor *api.MinerInfo, err error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		log.Errorf("failed to ChainHead: %s", err)
		return nil, err
	}

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateMinerInfo(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to StateMinerInfo: %s", err)
		return nil, err
	}

	return &ma, nil

}

func GetStateMinerOwner(minerId string) (*string, error) {

	mineraddr, err := address.NewFromString(minerId)
	log.Infof("mineraddr: %s\n", mineraddr)
	if err != nil {
		log.Errorf("failed to address.NewFromString: %s", err)
		return nil, err
	}

	currHead, err := fevmApi.ChainHead(ctx)
	if err != nil {
		log.Errorf("failed to ChainHead: %s", err)
		return nil, err
	}

	tsk := chaintypes.NewTipSetKey(currHead.Cids()...)

	ma, err := fevmApi.StateMinerInfo(context.Background(), mineraddr, tsk)
	if err != nil {
		log.Errorf("failed to StateMinerInfo: %s", err)
		return nil, err
	}

	var k = ma.Owner

	if k, err = fevmApi.StateAccountKey(ctx, ma.Owner, chaintypes.EmptyTSK); err != nil {
		log.Errorf("failed to StateAccountKey: %s", err)
		return nil, err
	}

	kstr := k.String()
	return &kstr, nil

}
