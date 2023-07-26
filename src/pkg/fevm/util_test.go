package fevm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestSortLogs(t *testing.T) {
	logs := []types.Log{
		{
			BlockNumber: 2,
			Index:       0,
			Address:     common.HexToAddress("0x1234"),
		},
		{
			BlockNumber: 1,
			Index:       1,
			Address:     common.HexToAddress("0x5678"),
		},
		{
			BlockNumber: 1,
			Index:       0,
			Address:     common.HexToAddress("0x1234"),
		},
		{
			BlockNumber: 2,
			Index:       1,
			Address:     common.HexToAddress("0x5678"),
		},
	}

	expected := []types.Log{
		{
			BlockNumber: 1,
			Index:       0,
			Address:     common.HexToAddress("0x1234"),
		},
		{
			BlockNumber: 1,
			Index:       1,
			Address:     common.HexToAddress("0x5678"),
		},
		{
			BlockNumber: 2,
			Index:       0,
			Address:     common.HexToAddress("0x1234"),
		},
		{
			BlockNumber: 2,
			Index:       1,
			Address:     common.HexToAddress("0x5678"),
		},
	}

	sortedLogs := SortLogs(logs)

	for i := range sortedLogs {
		if sortedLogs[i].BlockNumber != expected[i].BlockNumber || sortedLogs[i].Index != expected[i].Index {
			t.Errorf("Expected log at index %d to have BlockNumber %d and Index %d, but got BlockNumber %d and Index %d",
				i, expected[i].BlockNumber, expected[i].Index, sortedLogs[i].BlockNumber, sortedLogs[i].Index)
		}
	}
}
