package glif

import (
	"fmt"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
)

func GetGasFee() (gasFee *string, err error) {

	//log.Infof("get gas fee from glif %+v", glifApi)

	gasfee, err := fevmApi.EthMaxPriorityFeePerGas(ctx)
	if err != nil {
		log.Errorf("failed to get gas fee: %s", err)
		return nil, err
	}
	gasFeeString := fmt.Sprintf("%d", gasfee.Uint64())

	return &gasFeeString, nil

}
