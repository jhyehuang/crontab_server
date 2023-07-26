package filContractEvent

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/jhyehuang/crontab_server/src/pkg/glif"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"math/big"
	"strconv"
)

func FilterEventFIL(beginBlockNumber uint64, endBlockNumber uint64, conAddress []string) (mbs *ethtypes.EthFilterResult, err error) {

	startBigStr := strconv.FormatUint(uint64(beginBlockNumber), 16)
	endBigStr := strconv.FormatUint(uint64(endBlockNumber), 16)

	contractAddresses := make(ethtypes.EthAddressList, 0)

	var contractAddress ethtypes.EthAddress
	for _, v := range conAddress {
		contractAddress, err = ethtypes.ParseEthAddress(v)
		if err != nil {
			return nil, err
		}
		contractAddresses = append(contractAddresses, contractAddress)
	}

	query := ethtypes.EthFilterSpec{
		FromBlock: &startBigStr,
		ToBlock:   &endBigStr,
		Address:   contractAddresses,
		Topics:    AllTopics(), // [][]common.Hash
	}

	ret, err := glif.GetGetFilterLogs(query)
	if err != nil {
		log.Errorf("GetGetFilterLogs error: %v", err)
		return nil, err
	}

	//ret = 1;
	//*ethtypes.EthFilterResult
	//var mbs *[]types.Log
	return ret, nil
}

func FilterEventEth(beginBlockNumber uint64, endBlockNumber uint64, conAddress []string) (mbs *[]types.Log, err error) {

	contractAddresses := make([]common.Address, 0)

	var contractAddress common.Address
	for _, v := range conAddress {
		contractAddress = common.HexToAddress(v)
		if err != nil {
			return nil, err
		}
		contractAddresses = append(contractAddresses, contractAddress)
	}
	//sign_str := "PutValue(uint256)"
	//sign := crypto.Keccak256Hash([]byte(sign_str))
	//
	//log.Infof("sign: %s", sign.String())

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(beginBlockNumber)),
		ToBlock:   big.NewInt(int64(endBlockNumber)),
		Addresses: contractAddresses,
		Topics:    AllTopicsEth(),
		//Topics: [][]common.Hash{
		//	{sign},
		//},
	}

	ret, err := glif.GetGetFilterLogsEth(query)
	if err != nil {
		log.Errorf("GetGetFilterLogs error: %v", err)
		return nil, err
	}
	return ret, nil
}

func FilterEventEthGlif(beginBlockNumber uint64, endBlockNumber uint64, conAddress []string) (mbs *[]types.Log, err error) {

	contractAddresses := make([]common.Address, 0)

	var contractAddress common.Address
	for _, v := range conAddress {
		contractAddress = common.HexToAddress(v)
		if err != nil {
			return nil, err
		}
		contractAddresses = append(contractAddresses, contractAddress)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(beginBlockNumber)),
		ToBlock:   big.NewInt(int64(endBlockNumber)),
		Addresses: contractAddresses,
		Topics:    AllTopicsEth(),
	}

	ret, err := glif.GetGetFilterLogsEthGlif(query)
	if err != nil {
		log.Errorf("GetGetFilterLogs error: %v", err)
		return nil, err
	}
	return ret, nil
}

func FilterEventFEVM(beginBlockNumber uint64, endBlockNumber uint64, conAddress []string) (mbs *[]types.Log, err error) {

	contractAddresses := make([]common.Address, 0)

	var contractAddress common.Address
	for _, v := range conAddress {
		contractAddress = common.HexToAddress(v)
		if err != nil {
			return nil, err
		}
		contractAddresses = append(contractAddresses, contractAddress)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(beginBlockNumber)),
		ToBlock:   big.NewInt(int64(endBlockNumber)),
		Addresses: contractAddresses,
		Topics:    AllTopicsEth(),
	}

	ret, err := glif.GetGetFilterLogsEth(query)
	if err != nil {
		log.Errorf("GetGetFilterLogs error: %v", err)
		return nil, err
	}
	return ret, nil
}
