package glif

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
)

func GetGetFilterLogs(filter ethtypes.EthFilterSpec) (mbs *ethtypes.EthFilterResult, err error) {

	log.Infof("GetGetFilterLogs start %+v", filter)
	filt, err := fevmApi.EthNewFilter(context.TODO(), &filter)
	if err != nil {
		log.Infof("EthNewFilter failed: %s", err)
		return nil, err
	}

	log.Infof("filt: %+v\n", filt)
	ma, err := fevmApi.EthGetFilterLogs(context.Background(), filt)
	if err != nil {
		log.Infof("GetGetFilterLogs failed: %s", err)
		//fmt.Println(err)
		return nil, err
	}
	//fmt.Println(ma)

	return ma, nil

}

func GetGetFilterLogsEth(filter ethereum.FilterQuery) (mbs *[]types.Log, err error) {

	logs, err := httpsClient.FilterLogs(context.Background(), filter)
	if err != nil {
		log.Error("SubscribeFilterLogs ", err)
		return nil, err
	}
	if len(logs) != 0 {
		return RemoveDuplicateLogs(logs), nil
	}

	return &logs, nil

}

func GetGetFilterLogsEthGlif(filter ethereum.FilterQuery) (mbs *[]types.Log, err error) {

	logs, err := glifHttpsClient.FilterLogs(context.Background(), filter)
	if err != nil {
		log.Error("SubscribeFilterLogs ", err)
		return nil, err
	}
	if len(logs) != 0 {
		return RemoveDuplicateLogs(logs), nil
	}

	return &logs, nil

}

// 去重logs
func RemoveDuplicateLogs(logs []types.Log) *[]types.Log {
	var logMap = make(map[string]bool)
	newLogs := make([]types.Log, 0)

	for i, vLog := range logs {
		if _, ok := logMap[vLog.TxHash.String()+fmt.Sprintf("%d", vLog.Index)]; ok {
			continue
		}
		logMap[vLog.TxHash.String()+fmt.Sprintf("%d", vLog.Index)] = true
		newLogs = append(newLogs, logs[i])
	}
	return &newLogs
}
