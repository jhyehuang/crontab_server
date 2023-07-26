package glif

import (
	"github.com/filecoin-project/go-state-types/abi"
	"strconv"
	"strings"

	"github.com/jhyehuang/crontab_server/src/pkg/log"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/api"
	lminer "github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
)

func SyncMinerSectors(minerId address.Address, sectorNums *bitfield.BitField) []*miner.SectorOnChainInfo {
	//fevmApi 获取所有扇区 todo...优化，直接从Letsfil-miner获取可用扇区号
	sectors, err := fevmApi.StateMinerSectors(ctx, minerId, sectorNums, chaintypes.EmptyTSK)
	if err != nil {
		spt := strings.Split(err.Error(), " ")

		sn, err2 := strconv.Atoi(spt[len(spt)-1])
		log.Infof("SyncMinerSectors sn: %v, err: %v", sn, err2)
		if err2 != nil {
			return nil
		}
		sectorNums.Unset(uint64(sn))
		sectors = SyncMinerSectors(minerId, sectorNums)
	}

	return sectors
}

func SyncMinerActiveSectors(minerId address.Address) []*miner.SectorOnChainInfo {
	sectors, err := fevmApi.StateMinerActiveSectors(ctx, minerId, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("SyncMinerActiveSectors %v", err)
		return nil
	}

	return sectors
}

func GetActiveSectorNos(minerId string) map[uint64]interface{} {
	addr, _ := address.NewFromString(minerId)
	sectors, err := fevmApi.StateMinerActiveSectors(ctx, addr, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetActiveSectorNos %v", err)
		return nil
	}
	result := make(map[uint64]interface{})
	for _, sector := range sectors {
		result[uint64(sector.SectorNumber)] = struct{}{}
	}
	return result
}

func GetFaultSectors(minerId address.Address) bitfield.BitField {
	faultSectors, err := fevmApi.StateMinerFaults(ctx, minerId, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return bitfield.New()
	}

	return faultSectors
}

func GetSectorCount(minerId string) (api.MinerSectors, error) {
	addr, _ := address.NewFromString(minerId)
	actorSectors, err := fevmApi.StateMinerSectorCount(ctx, addr, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetSectorCount %v", err)
		return api.MinerSectors{}, err
	}
	return actorSectors, nil
}

func GetRecoveriesSectors(minerId address.Address) (bitfield.BitField, error) {
	faultSectors, err := fevmApi.StateMinerRecoveries(ctx, minerId, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return bitfield.New(), err
	}

	return faultSectors, nil
}

func GetStateMinerPartitions(minerId address.Address, dl uint64) ([]api.Partition, error) {
	faultSectors, err := fevmApi.StateMinerPartitions(ctx, minerId, dl, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return []api.Partition{}, err
	}

	return faultSectors, nil
}

func GetStateSectorPartition(minerId address.Address, sectorNo uint64) (*lminer.SectorLocation, error) {
	faultSectors, err := fevmApi.StateSectorPartition(ctx, minerId, abi.SectorNumber(sectorNo), chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return nil, err
	}

	return faultSectors, nil
}

func GetAllLiveSectorNos(minerId address.Address) (*[]uint64, *[]uint64, error) {
	AllsectorNos := make([]uint64, 0)
	LivesectorNos := make([]uint64, 0)

	//minerDl, err := fevmApi.StateMinerDeadlines(ctx, minerId, chaintypes.EmptyTSK)
	//if err != nil {
	//	log.Errorf("GetFaultSectors %v", err)
	//	return nil, nil, err
	//}
	//log.Infof("GetAllLiveSectorNos minerDl: %+v", len(minerDl))

	// 循环 到48
	for dlIdx := uint64(0); dlIdx < 48; dlIdx++ {

		log.Infof("GetAllLiveSectorNos dlIdx: %v", dlIdx)
		minerDl, err := fevmApi.StateMinerPartitions(ctx, minerId, dlIdx, chaintypes.EmptyTSK)
		if err != nil {
			log.Errorf("GetFaultSectors %v", err)
			return nil, nil, err
		}
		for _, part := range minerDl {
			tmp1, _ := part.AllSectors.Count()
			log.Infof("GetAllLiveSectorNos part AllSectors : %d", tmp1)
			tmp1, _ = part.LiveSectors.Count()
			log.Infof("GetAllLiveSectorNos part LiveSectors : %d", tmp1)
			tmp2, err := part.AllSectors.All(2349)
			if err != nil {
				log.Errorf("GetFaultSectors %v", err)
				return nil, nil, err
			}
			AllsectorNos = append(AllsectorNos, tmp2...)

			tmp3, err := part.LiveSectors.All(2349)
			if err != nil {
				log.Errorf("GetFaultSectors %v", err)
				return nil, nil, err
			}
			LivesectorNos = append(LivesectorNos, tmp3...)
		}
	}

	return &AllsectorNos, &LivesectorNos, nil
}

func GetStateMinerSectors(minerId address.Address, sectorNos []uint64) ([]*miner.SectorOnChainInfo, error) {

	bt := bitfield.NewFromSet(sectorNos)
	faultSectors, err := fevmApi.StateMinerSectors(ctx, minerId, &bt, chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return nil, err
	}

	return faultSectors, nil
}

func GetStateSectorExpiration(minerId address.Address, sectorNo uint64) (*lminer.SectorExpiration, error) {
	faultSectors, err := fevmApi.StateSectorExpiration(ctx, minerId, abi.SectorNumber(sectorNo), chaintypes.EmptyTSK)
	if err != nil {
		log.Errorf("GetFaultSectors %v", err)
		return nil, err
	}

	return faultSectors, nil
}
