package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/global/constant"
)

type RpcClient struct {
}

func NewRpcClient() *RpcClient {
	return &RpcClient{}
}

func (this *RpcClient) InitRpcClient() *api.FullNodeStruct {
	headers := http.Header{constant.RpcHeaderKey: []string{constant.RpcHeaderTokenPrefix + configs.GetValue(configs.LotusAuthToken, configs.Get().Lotus.AuthToken)}}
	var api api.FullNodeStruct
	_, err := jsonrpc.NewMergeClient(context.Background(), configs.GetValue(configs.LotusUrl, configs.Get().Lotus.Url), configs.Get().Lotus.NameSpace,
		[]interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}
	return &api
}
