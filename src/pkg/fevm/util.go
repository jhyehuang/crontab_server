package fevm

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"golang.org/x/crypto/sha3"
	"sort"
)

func EthFunctionHash(sig string) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(sig))
	return hasher.Sum(nil)[:4]
}

func EthTopicHash(sig string) ethtypes.EthHash {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(sig))
	var hash ethtypes.EthHash
	copy(hash[:], hasher.Sum(nil))
	return hash
}

func ParseEthTopics(topics ethtypes.EthTopicSpec) (map[string][][]byte, error) {
	keys := map[string][][]byte{}
	for idx, vals := range topics {
		if len(vals) == 0 {
			continue
		}
		// Ethereum topics are emitted using `LOG{0..4}` opcodes resulting in topics1..4
		key := fmt.Sprintf("t%d", idx+1)
		for _, v := range vals {
			v := v // copy the ethhash to avoid repeatedly referencing the same one.
			keys[key] = append(keys[key], v[:])
		}
	}
	return keys, nil
}

func ParseEthLogs(logs interface{}) (*ethtypes.EthLog, error) {
	resByre, resByteErr := json.Marshal(logs)
	if resByteErr != nil {
		return nil, resByteErr
	}
	var newData ethtypes.EthLog
	jsonRes := json.Unmarshal(resByre, &newData)
	if jsonRes != nil {
		return nil, jsonRes
	}
	return &newData, nil
}

func SortLogs(logs []types.Log) []types.Log {
	// Define a custom sorting function that compares logs by their BlockNumber and Index fields
	sortFn := func(i, j int) bool {
		if logs[i].BlockNumber == logs[j].BlockNumber {
			return logs[i].Index < logs[j].Index
		}
		return logs[i].BlockNumber < logs[j].BlockNumber
	}

	// Sort the logs slice using the custom sorting function
	sort.Slice(logs, sortFn)

	return logs
}
