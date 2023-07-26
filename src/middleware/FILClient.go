package middleware

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
)

type HttpsFILClient struct {
}

func NewHttpsFILClient() *HttpsFILClient {
	return &HttpsFILClient{}
}

func (this *HttpsFILClient) HttpsFILClient() *ethclient.Client {
	//httpsClient, err := ethclient.Dial(configs.GetValue(configs.FileEvmHttpsUrl, configs.Get().FilFevm.HttpsUrl))
	httpsClient, err := ethclient.Dial(configs.GetValue(configs.GlifUrlV1, configs.Get().Glif.UrlV1))
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	return httpsClient
}

type WsFILClient struct {
}

func NewWsFILClient() *WsFILClient {
	return &WsFILClient{}
}

func (this *WsFILClient) WsFILClient() *ethclient.Client {
	httpsClient, err := ethclient.Dial(configs.GetValue(configs.FilEvmWsUrl, configs.Get().FilFevm.WsUrl))
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	return httpsClient
}
