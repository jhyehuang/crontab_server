package glif

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/jhyehuang/crontab_server/src/configs"
)

//var LotusNodeAddr = "https://api.node.glif.io/rpc/v0"

//var LotusNodeAddr = "https://api.hyperspace.node.glif.io/rpc/v1"

var fevmApi api.FullNodeStruct
var glifApiV0 api.FullNodeStruct
var glifApiV1 api.FullNodeStruct
var ctx = context.Background()
var httpsClient *ethclient.Client
var glifHttpsClient *ethclient.Client

func InitNode() {
	var Token = configs.GetValue(configs.FilEvmToken, configs.Get().FilFevm.AuthToken)
	var LotusNodeAddr = configs.GetValue(configs.FileEvmHttpsUrl, configs.Get().FilFevm.HttpsUrl)
	var headers http.Header
	if Token == "" {
		headers = http.Header{
			"content-type": []string{"application/json"},
		}
	} else {
		headers = http.Header{
			"content-type":  []string{"application/json"},
			"Authorization": []string{"Bearer " + Token},
		}
	}
	closer, err := jsonrpc.NewMergeClient(context.Background(), LotusNodeAddr, "Filecoin", []interface{}{&fevmApi.Internal, &fevmApi.CommonStruct.Internal}, headers)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}
	defer closer()

	httpsClient, err = ethclient.Dial(LotusNodeAddr)
	//httpsClient, err = ethclient.Dial(configs.Get().FilFevm.HttpsUrl)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}

	glifHttpsClient, err = ethclient.Dial(configs.GetValue(configs.GlifUrlV1, configs.Get().Glif.UrlV1))
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}

	headers2 := http.Header{
		"content-type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + configs.GetValue(configs.GlifAuthToken, configs.Get().Glif.AuthToken)},
	}

	closer2, err := jsonrpc.NewMergeClient(context.Background(), configs.GetValue(configs.GlifUrlV0, configs.Get().Glif.UrlV0), "Filecoin", []interface{}{&glifApiV0.Internal, &glifApiV0.CommonStruct.Internal}, headers2)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}
	defer closer2()

	closer3, err := jsonrpc.NewMergeClient(context.Background(), configs.GetValue(configs.GlifUrlV1, configs.Get().Glif.UrlV1), "Filecoin", []interface{}{&glifApiV1.Internal, &glifApiV1.CommonStruct.Internal}, headers2)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}
	defer closer3()
}
