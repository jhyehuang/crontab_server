package main

import (
	"context"
	"encoding/json"
	"github.com/filecoin-project/lotus/build"
	"github.com/jhyehuang/crontab_server/src/middleware"
	"github.com/jhyehuang/crontab_server/src/pkg/filscan"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/urfave/cli"
	"os"
)

func main() {

	log.Info("Starting filsacn-chain-stat...")

	app := &cli.App{
		Name:    "get-chain-stat",
		Usage:   "Benchmark performance of lotus on your hardware",
		Version: build.UserVersion(),
		Commands: []cli.Command{
			getCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}

var getCmd = cli.Command{
	Name:      "get-chain-stat",
	Usage:     "Benchmark a proof computation",
	ArgsUsage: "[input.json]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "version",
			Usage:    "pass miner address (only necessary if using existing sectorbuilder)",
			Value:    "v1",
			Required: false,
		},
	},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		version := c.String("version")

		var ret = &filscan.Response{}
		var err error
		if version == "v1" {
			ret, err = filscan.GetStatChainInfoV1()
		} else if version == "v2" {
			ret, err = filscan.GetStatChainInfoV2()
		} else {
			log.Warn("version error")
			return nil
		}
		if err != nil {
			log.Warnf("get stat chain info error: %+v", err)
			return nil
		}
		log.Infof("get stat chain info success: %+v", ret)

		rc := middleware.NewRedisClient().RedisClient()
		if rc == nil {
			log.Errorf("Blockchain - Failed to connect redis")
			return nil
		}

		if ret != nil {
			// 保存到redis
			var result []byte
			result, err := json.Marshal(ret)
			if err != nil {
				log.Errorf("Blockchain - Failed to unmarshal response: %s", err.Error())
				return nil
			}
			// redis 设置 filscan_stat_chain_info
			redisStr := rc.Set(ctx, "filscan_stat_chain_info", result, 0)
			if err := redisStr.Err(); err != nil {
				log.Errorf("Blockchain - Failed to set redis: %s", err.Error())
				return nil
			}
			// todo.... 实时更新统计值，在募集计划创建的时候直接更新到募集计划中
			// redis 设置 PledgePerTera
			pledgePerTera := rc.Set(ctx, "statChain:MainNet:ResultData:PledgePerTera", ret.Result.Data.PledgePerTera, 0)
			if err := pledgePerTera.Err(); err != nil {
				log.Errorf("Blockchain - Failed to set redis: %s", err.Error())
				return nil
			}

			// redis 设置 PledgePerTera
			filPerTera := rc.Set(ctx, "statChain:MainNet:ResultData:FilPerTera", ret.Result.Data.FilPerTera, 0)
			if err := filPerTera.Err(); err != nil {
				log.Errorf("Blockchain - Failed to set redis: %s", err.Error())
				return nil
			}
		}

		return nil
	},
}
