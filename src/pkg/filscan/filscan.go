package filscan

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/go-redis/redis/v8"
	"github.com/jhyehuang/crontab_server/src/pkg"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type FilscanChainInfo struct {
	LatestHeight       int     `json:"latest_height"`
	LatestBlockReward  string  `json:"latest_block_reward"`
	TotalBlocks        int     `json:"total_blocks"`
	TotalRewards       string  `json:"total_rewards"`
	PowerRatio         string  `json:"power_ratio"`
	PowerIncrease24h   string  `json:"power_increase_24h"`
	PowerIncrease1h    string  `json:"power_increase_1h"`
	RewardsIncrease24h string  `json:"rewards_increase_24h"`
	FilPerTera         string  `json:"fil_per_tera"`
	TotalQualityPower  string  `json:"total_quality_power"`
	ActiveMiners       int     `json:"active_miners"`
	BaseFee            string  `json:"base_fee"`
	PledgePerTera      string  `json:"pledge_per_tera"`
	AvgMessageCount    int     `json:"avg_message_count"`
	AvgBlockCount      string  `json:"avg_block_count"`
	GasIn32g           string  `json:"gas_in_32g"`
	GasIn64g           *string `json:"gas_in_64g"`
	AddPowerIn32g      string  `json:"add_power_in_32g"`
	AddPowerIn64g      string  `json:"add_power_in_64g"`
	MinerInitialPledge string  `json:"miner_initial_pledge"`
	Burnt              string  `json:"burnt"`
	CirculatingPercent string  `json:"circulating_percent"`
}

type Result struct {
	Data FilscanChainInfo `json:"data"`
}
type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

type ResponseV2 struct {
	Result struct {
		TotalIndicators struct {
			LatestHeight       int     `json:"latest_height" gorm:"column:latest_height"`
			AvgMessageCount    float64 `json:"avg_message_count" gorm:"column:avg_message_count"`
			CirculatingPercent string  `json:"circulating_percent" gorm:"column:circulating_percent"`
			TotalQualityPower  string  `json:"total_quality_power" gorm:"column:total_quality_power"`
			TotalRewards       string  `json:"total_rewards" gorm:"column:total_rewards"`
			MinerInitialPledge string  `json:"miner_initial_pledge" gorm:"column:miner_initial_pledge"`
			PowerIncrease24h   string  `json:"power_increase_24h" gorm:"column:power_increase_24h"`
			ActiveMiners       int     `json:"active_miners" gorm:"column:active_miners"`
			TotalBlocks        int     `json:"total_blocks" gorm:"column:total_blocks"`
			GasIn32g           string  `json:"gas_in_32g" gorm:"column:gas_in_32g"`
			GasIn64g           string  `json:"gas_in_64g" gorm:"column:gas_in_64g"`
			BaseFee            string  `json:"base_fee" gorm:"column:base_fee"`
			LatestBlockTime    int     `json:"latest_block_time" gorm:"column:latest_block_time"`
			RewardsIncrease24h string  `json:"rewards_increase_24h" gorm:"column:rewards_increase_24h"`
			WinCountReward     string  `json:"win_count_reward" gorm:"column:win_count_reward"`
			AddPowerIn32g      string  `json:"add_power_in_32g" gorm:"column:add_power_in_32g"`
			AddPowerIn64g      string  `json:"add_power_in_64g" gorm:"column:add_power_in_64g"`
			AvgBlockCount      string  `json:"avg_block_count" gorm:"column:avg_block_count"`
			FilPerTera24h      string  `json:"fil_per_tera_24h" gorm:"column:fil_per_tera_24h"`
			Burnt              string  `json:"burnt" gorm:"column:burnt"`
		} `json:"total_indicators" gorm:"column:total_indicators"`
	} `json:"result" gorm:"column:result"`
}

func GetStatChainInfo() (*Response, error) {
	return GetStatChainInfoV2()
}

func GetStatChainInfoV1() (*Response, error) {
	// 构造API请求
	url := "https://api.filscan.io:8700/rpc/v1"
	method := "POST"

	payload := strings.NewReader(`{"id":1,"jsonrpc":"2.0","params":[],"method":"filscan.StatChainInfo"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Errorf("Blockchain - Failed to create request: %s", err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Origin", "https://filscan.io")
	var res *http.Response
	for i := 0; i < 3; i++ {
		res, err = client.Do(req)
		if err == nil {
			// 成功获取响应，退出循环
			break
		}
		time.Sleep(1 * time.Second)
	}
	if res != nil {
		defer res.Body.Close()
	} else {
		log.Errorf("Blockchain - Failed to get response: %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Blockchain - Failed to read response: %s", err.Error())
		return nil, err
	}

	log.Infof("Blockchain - Response: %s", string(body))
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Errorf("Blockchain - Failed to unmarshal response: %s", err.Error())
		return nil, err
	}

	return &result, nil
}

func GetStatChainInfoV2() (*Response, error) {
	// 构造API请求
	url := "https://api-v2.filscan.io:27000/api/v1/TotalIndicators"
	method := "POST"

	payload := strings.NewReader(`{}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Errorf("Blockchain - Failed to create request: %s", err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Origin", "https://filscan.io")
	var res *http.Response
	for i := 0; i < 3; i++ {
		res, err = client.Do(req)
		if err == nil {
			// 成功获取响应，退出循环
			break
		}
		time.Sleep(1 * time.Second)
	}
	if res != nil {
		defer res.Body.Close()
	} else {
		log.Errorf("Blockchain - Failed to get response: %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Blockchain - Failed to read response: %s", err.Error())
		return nil, err
	}

	var result ResponseV2
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Errorf("Blockchain - Failed to unmarshal response: %s", err.Error())
		return nil, err
	}

	var ret Response
	ret.Result.Data.LatestHeight = result.Result.TotalIndicators.LatestHeight
	pledgePerTera, err := big.FromString(result.Result.TotalIndicators.MinerInitialPledge)
	if err != nil {
		log.Errorf("Blockchain - Failed to convert string to big.Int: %s", err.Error())
		return nil, err
	}
	ret.Result.Data.PledgePerTera = fmt.Sprintf("%.18f", pkg.ToFloat64(abi.TokenAmount{pledgePerTera.Int}))

	ret.Result.Data.RewardsIncrease24h = result.Result.TotalIndicators.RewardsIncrease24h
	tmpStr := strings.Split(result.Result.TotalIndicators.FilPerTera24h, ".")

	filPerTera24h, err := big.FromString(tmpStr[0])
	if err != nil {
		log.Errorf("Blockchain - Failed to convert string to big.Int: %s", err.Error())
		return nil, err
	}
	ret.Result.Data.FilPerTera = fmt.Sprintf("%.18f", pkg.ToFloat64(abi.TokenAmount{filPerTera24h.Int}))
	ret.Result.Data.AddPowerIn32g = result.Result.TotalIndicators.AddPowerIn32g
	ret.Result.Data.AddPowerIn64g = result.Result.TotalIndicators.AddPowerIn64g
	ret.Result.Data.AvgBlockCount = result.Result.TotalIndicators.AvgBlockCount
	ret.Result.Data.BaseFee = result.Result.TotalIndicators.BaseFee
	ret.Result.Data.Burnt = result.Result.TotalIndicators.Burnt

	return &ret, nil
}

func GetStatChainInfoCache(c context.Context, red *redis.Client) (*Response, error) {
	var statChain = Response{}

	redisStr := red.Get(c, "filscan_stat_chain_info")
	body, err := redisStr.Bytes()
	if err != nil {
		log.Errorf("Blockchain - Failed to get redis: %s", err.Error())
		return nil, err
	} else {
		err = json.Unmarshal(body, &statChain)
		if err != nil {
			log.Errorf("Blockchain - Failed to unmarshal response: %s", err.Error())
			//return nil, err
		}

		// 产品要求写死，通过redis获取一个新的值然后覆盖
		filPerTeraStr := red.Get(c, "statChain:MainNet:ResultData:FilPerTera")
		body, err := filPerTeraStr.Bytes()
		if err != nil {
			log.Errorf("Blockchain - Failed to get redis: %s", err.Error())
		} else {
			statChain.Result.Data.FilPerTera = string(body)
		}

		return &statChain, nil
	}
}
