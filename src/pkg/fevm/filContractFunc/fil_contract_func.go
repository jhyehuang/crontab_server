package filContractFunc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/global/constant"
	"github.com/jhyehuang/crontab_server/src/models"
	"github.com/jhyehuang/crontab_server/src/pkg/fevm/contract"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"math/big"
)

func PostBlockChainData(cli *ethclient.Client, conAddress string, methmod string, payload ...interface{}) (receipt *models.PushContractTxRecord, err error) {

	// 创建交易
	log.Infof("payload: %+v", payload)
	log.Warnf("Methmod: %s", methmod)
	log.Infof("conAddress: %s", conAddress)

	auth, err := GetBlockChainTransactAuth(cli)
	if err != nil {
		log.Errorf("Blockchain - Failed to create authorized transactor: %s", err.Error())
		return nil, err
	}

	var tx *types.Transaction
	switch methmod {
	case "putFilReward":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		resealedAmt := payload[1].(*big.Int)
		willResealedAmt := payload[2].(*big.Int)

		// 调用 PushRewardParams 方法
		tx, err = contract.PushBlockReward(auth, assetPackId, resealedAmt, willResealedAmt)
		if err != nil {
			log.Errorf("Blockchain - Failed PushRewardParams: %s", err.Error())
			return nil, err
		}

	case "StartSeal":
		// 创建指向智能合约地址的实例
		//contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		//if err != nil {
		//	return nil, err
		//}
		//beginSealTime := payload[0].(*big.Int)
		//// 调用 StartSeal 方法
		//tx, err = contract.StartSeal(auth, beginSealTime)
		//if err != nil {
		//	log.Errorf("Blockchain - Failed StartSeal: %s", err.Error())
		//	return nil, err
		//}
		//log.Infof("tx: %+v", tx)
		//log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "EndSeal":

	case "PushSealProgress":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		curAmt := payload[1].(*big.Int)
		assetPackId := payload[0].(*big.Int)

		tx, err = contract.PushSealProgress(auth, assetPackId, curAmt)
		if err != nil {
			log.Errorf("Blockchain - Failed PushSealProgress: %s", err.Error())
			return nil, err
		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "pushFinalProgress":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		curAmt := payload[1].(*big.Int)
		assetPackId := payload[0].(*big.Int)

		tx, err = contract.PushFinalProgress(auth, assetPackId, curAmt)
		if err != nil {
			log.Errorf("Blockchain - Failed PushFinalProgress: %s", err.Error())
			return nil, err
		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "EndMiner":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			log.Errorf("Blockchain - Failed EndMiner: %s", err.Error())
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		// 调用 DestroyNode 方法
		tx, err = contract.DestroyNode(auth, assetPackId)
		if err != nil {
			log.Errorf("Blockchain - Failed EndMiner: %s", err.Error())
			return nil, err
		}
	case "RaiseDealLine":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			log.Errorf("Blockchain - Failed RaiseDealLine: %s", err.Error())
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		// 调用 raiseExpire 方法
		tx, err = contract.RaiseExpire(auth, assetPackId)
		if err != nil {
			log.Errorf("Blockchain - Failed RaiseDealLine: %s", err.Error())
			return nil, err
		}

	case "PushSpFine":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			log.Errorf("Blockchain - Failed SPPenalty: %s", err.Error())
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		amt := payload[1].(*big.Int)

		// 调用 SPPenalty 方法
		tx, err = contract.PushSpFine(auth, assetPackId, amt)
		if err != nil {
			log.Errorf("Blockchain - Failed SPPenalty: %s", err.Error())
			return nil, err
		}
	case "pushPledgeReleased":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			log.Errorf("Blockchain - Failed PledgeRelease: %s", err.Error())
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		amt := payload[1].(*big.Int)

		// 调用 PledgeRelease 方法
		tx, err = contract.PushPledgeReleased(auth, assetPackId, amt)
		if err != nil {
			log.Errorf("Blockchain - Failed PledgeRelease: %s", err.Error())
			return nil, err
		}
	case "pushOldAssetPackValue":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			log.Errorf("Blockchain - Failed OldAssetPackValue: %s", err.Error())
			return nil, err
		}
		assetPackId := payload[0].(*big.Int)
		totalPledge := payload[1].(*big.Int)
		released := payload[2].(*big.Int)
		willRelease := payload[3].(*big.Int)
		// 调用 OldAssetPackValue 方法
		tx, err = contract.PushOldAssetPackValue(auth, assetPackId, totalPledge, released, willRelease)
		if err != nil {
			log.Errorf("Blockchain - Failed OldAssetPackValue: %s", err.Error())
			return nil, err
		}

	}
	log.Infof("tx: %s", tx.Hash().String())

	// 登记 PushContractTxRecord 记录表
	var txRecord = &models.PushContractTxRecord{}
	txRecord.TxHash = tx.Hash().Hex()
	txRecord.Method = methmod
	txRecord.From = auth.From.String()
	txRecord.To = conAddress
	txRecord.Value = tx.Value().String()
	txRecord.Params = ""
	txRecord.Status = constant.ChainTxStatusUnconfirmed
	txRecord.Nonce = tx.Nonce()

	return txRecord, nil

	// 保存到数据库

}

func GetBlockChainTransactAuth(cli *ethclient.Client) (receipt *bind.TransactOpts, err error) {

	privateKey, err := ethcrypto.HexToECDSA(configs.GetValue(configs.Private_Key, ""))
	if err != nil {
		return nil, err
	}
	nonce, err := cli.PendingNonceAt(context.Background(), ethcrypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	log.Infof("gasPrice: %s", gasPrice.String())

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))
	auth.GasTipCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))
	return auth, nil
}

func GetBlockChainCallAuth(cli *ethclient.Client) (receipt *bind.CallOpts, err error) {

	privateKey, err := ethcrypto.HexToECDSA(configs.GetValue(configs.Private_Key, ""))
	if err != nil {
		return nil, err
	}
	nonce, err := cli.PendingNonceAt(context.Background(), ethcrypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	//log.Infof("gasPrice: %s", gasPrice.String())

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))
	auth.GasTipCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))

	bopt := bind.CallOpts{
		Pending:     false,
		From:        auth.From,
		BlockNumber: nil,
		Context:     context.Background(),
	}

	return &bopt, nil
}
