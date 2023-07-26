package filContractFunc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/pkg/bigInt"
	"github.com/jhyehuang/crontab_server/src/pkg/fevm/contract"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"testing"
)

func TestTotalReward(t *testing.T) {
	t.Log("test")
	cli, err := ethclient.Dial("https://api.hyperspace.node.glif.io")
	if err != nil {
		t.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	t.Log("httpsClient: ", cli)
	conAddress := "0x01040d7a134fba193755F87b40Dc4B0dcf89eFc9"
	contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
	if err != nil {
		t.Log(err)
	}

	// 创建交易
	log.Infof("conAddress: %s", conAddress)
	privateKey, err := ethcrypto.HexToECDSA(configs.GetValue(configs.Private_Key, ""))

	auth := bind.CallOpts{
		From:    ethcrypto.PubkeyToAddress(privateKey.PublicKey),
		Context: context.Background(),
	}

	// 调用 PutValue 方法
	totalReward, err := contract.TotalRewardAmount(&auth, bigInt.NewBigInt(0).Int)
	if err != nil {
		log.Errorf("Blockchain - Failed EndSeal： err: %s", err.Error())
	}
	log.Infof("totalReward: %+v", totalReward)
}
