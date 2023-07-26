// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ILetsFilPackInfoExtendInfo is an auto generated low-level Go binding around an user-defined struct.
type ILetsFilPackInfoExtendInfo struct {
	OldId                 *big.Int
	RaiserOldShare        *big.Int
	SpOldShare            *big.Int
	SponsorOldRewardShare *big.Int
	SpOldRewardShare      *big.Int
}

// ILetsFilPackInfoNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type ILetsFilPackInfoNodeInfo struct {
	MinerId         uint64
	NodeSize        *big.Int
	SectorSize      *big.Int
	SealDays        *big.Int
	NodeDays        *big.Int
	OpsSecurityFund *big.Int
	SpAddr          common.Address
	CompanyId       *big.Int
}

// ILetsFilPackInfoRaiseInfo is an auto generated low-level Go binding around an user-defined struct.
type ILetsFilPackInfoRaiseInfo struct {
	Id            *big.Int
	TargetAmount  *big.Int
	MinRaiseRate  *big.Int
	SecurityFund  *big.Int
	RaiseDays     *big.Int
	InvestorShare *big.Int
	SpFundShare   *big.Int
	RaiserShare   *big.Int
	ServicerShare *big.Int
	FilFiShare    *big.Int
	Sponsor       common.Address
	RaiseCompany  string
}

// LetsFilControlerMetaData contains all meta data concerning the LetsFilControler contract.
var LetsFilControlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddOpsSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toSealAmount\",\"type\":\"uint256\"}],\"name\":\"ClosePlanToSeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closeTime\",\"type\":\"uint256\"}],\"name\":\"CloseRaisePlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositOpsSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumNodeState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DestroyNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvestorWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumNodeState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"NodeStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"PushBlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumNodeState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"PushFinalProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPledge\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"PushOldAssetPackValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"}],\"name\":\"PushPledgeReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumNodeState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"PushSealProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fineAmount\",\"type\":\"uint256\"}],\"name\":\"PushSpFine\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"}],\"name\":\"RaiseFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumRaiseState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"RaiseStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"}],\"name\":\"RaiseSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaiseWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SealEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendToMiner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sp\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"SpSignWithMiner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SpWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"StartPreSeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartRaisePlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartSeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"}],\"name\":\"Unstaking\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"WithdrawOpsSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawSecurityFund\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"PLEDGE_MIN_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"addOpsSecurityFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"availableRewardOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_actorId\",\"type\":\"uint64\"}],\"name\":\"changeOwnerById\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"closeRaisePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"destroyNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extendInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oldId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sponsorOldRewardShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldRewardShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundBack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundSealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToolAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"process\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gotFilFiReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gotMiner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gotRaiserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gotSpFundReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gotSpReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raiseID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRaiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spFundShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filFiShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"raiseCompany\",\"type\":\"string\"}],\"internalType\":\"structILetsFilPackInfo.RaiseInfo\",\"name\":\"_raiseInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"spAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"companyId\",\"type\":\"uint256\"}],\"internalType\":\"structILetsFilPackInfo.NodeInfo\",\"name\":\"_nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sponsorOldRewardShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldRewardShare\",\"type\":\"uint256\"}],\"internalType\":\"structILetsFilPackInfo.ExtendInfo\",\"name\":\"_extendInfo\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"investorFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"investorInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pledgeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledgeCalcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"investorWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"spAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"companyId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeState\",\"outputs\":[{\"internalType\":\"enumNodeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opsCalcFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"opsFundReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opsSecurityFundRemain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"payOpsSecurityFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"paySecurityFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pledgeReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pledgeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pledgeTotalCalcAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pledgeTotalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"pushBlockReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pushFinalProgress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"pushOldAssetPackValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"}],\"name\":\"pushPledgeReleased\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pushSealProgress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fineAmount\",\"type\":\"uint256\"}],\"name\":\"pushSpFine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raiseCloseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiseExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raiseInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRaiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spFundShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filFiShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"raiseCompany\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raiseStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raiseState\",\"outputs\":[{\"internalType\":\"enumRaiseState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raiser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raiserFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiserRewardAvailableLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiserWillReleaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiserWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"returnAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeFundFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeSealFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sealEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sealedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secPack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"securityFundRemain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"minerOwner\",\"type\":\"bytes\"}],\"name\":\"setMinerBackOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setSectorPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_raiseStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startSealTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sealEndTime\",\"type\":\"uint256\"}],\"name\":\"setTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spFundFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spFundRewardFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spRemainFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"spRewardAvailableLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReturn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spRewardFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spRewardLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spSignWithMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"spWillReleaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"spWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"staking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"startPreSeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"startRaisePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"startSealTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toSealAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalReleasedRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"totalRewardOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"unStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"willReleaseOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdrawOpsSecurityFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdrawSecurityFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LetsFilControlerABI is the input ABI used to generate the binding from.
// Deprecated: Use LetsFilControlerMetaData.ABI instead.
var LetsFilControlerABI = LetsFilControlerMetaData.ABI

// LetsFilControler is an auto generated Go binding around an Ethereum contract.
type LetsFilControler struct {
	LetsFilControlerCaller     // Read-only binding to the contract
	LetsFilControlerTransactor // Write-only binding to the contract
	LetsFilControlerFilterer   // Log filterer for contract events
}

// LetsFilControlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LetsFilControlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilControlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LetsFilControlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilControlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LetsFilControlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilControlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LetsFilControlerSession struct {
	Contract     *LetsFilControler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LetsFilControlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LetsFilControlerCallerSession struct {
	Contract *LetsFilControlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LetsFilControlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LetsFilControlerTransactorSession struct {
	Contract     *LetsFilControlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LetsFilControlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LetsFilControlerRaw struct {
	Contract *LetsFilControler // Generic contract binding to access the raw methods on
}

// LetsFilControlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LetsFilControlerCallerRaw struct {
	Contract *LetsFilControlerCaller // Generic read-only contract binding to access the raw methods on
}

// LetsFilControlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LetsFilControlerTransactorRaw struct {
	Contract *LetsFilControlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLetsFilControler creates a new instance of LetsFilControler, bound to a specific deployed contract.
func NewLetsFilControler(address common.Address, backend bind.ContractBackend) (*LetsFilControler, error) {
	contract, err := bindLetsFilControler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LetsFilControler{LetsFilControlerCaller: LetsFilControlerCaller{contract: contract}, LetsFilControlerTransactor: LetsFilControlerTransactor{contract: contract}, LetsFilControlerFilterer: LetsFilControlerFilterer{contract: contract}}, nil
}

// NewLetsFilControlerCaller creates a new read-only instance of LetsFilControler, bound to a specific deployed contract.
func NewLetsFilControlerCaller(address common.Address, caller bind.ContractCaller) (*LetsFilControlerCaller, error) {
	contract, err := bindLetsFilControler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerCaller{contract: contract}, nil
}

// NewLetsFilControlerTransactor creates a new write-only instance of LetsFilControler, bound to a specific deployed contract.
func NewLetsFilControlerTransactor(address common.Address, transactor bind.ContractTransactor) (*LetsFilControlerTransactor, error) {
	contract, err := bindLetsFilControler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerTransactor{contract: contract}, nil
}

// NewLetsFilControlerFilterer creates a new log filterer instance of LetsFilControler, bound to a specific deployed contract.
func NewLetsFilControlerFilterer(address common.Address, filterer bind.ContractFilterer) (*LetsFilControlerFilterer, error) {
	contract, err := bindLetsFilControler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerFilterer{contract: contract}, nil
}

// bindLetsFilControler binds a generic wrapper to an already deployed contract.
func bindLetsFilControler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LetsFilControlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilControler *LetsFilControlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilControler.Contract.LetsFilControlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilControler *LetsFilControlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilControler.Contract.LetsFilControlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilControler *LetsFilControlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilControler.Contract.LetsFilControlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilControler *LetsFilControlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilControler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilControler *LetsFilControlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilControler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilControler *LetsFilControlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilControler.Contract.contract.Transact(opts, method, params...)
}

// PLEDGEMINAMOUNT is a free data retrieval call binding the contract method 0x4ccd44d3.
//
// Solidity: function PLEDGE_MIN_AMOUNT() view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) PLEDGEMINAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "PLEDGE_MIN_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PLEDGEMINAMOUNT is a free data retrieval call binding the contract method 0x4ccd44d3.
//
// Solidity: function PLEDGE_MIN_AMOUNT() view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) PLEDGEMINAMOUNT() (*big.Int, error) {
	return _LetsFilControler.Contract.PLEDGEMINAMOUNT(&_LetsFilControler.CallOpts)
}

// PLEDGEMINAMOUNT is a free data retrieval call binding the contract method 0x4ccd44d3.
//
// Solidity: function PLEDGE_MIN_AMOUNT() view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) PLEDGEMINAMOUNT() (*big.Int, error) {
	return _LetsFilControler.Contract.PLEDGEMINAMOUNT(&_LetsFilControler.CallOpts)
}

// AvailableRewardOf is a free data retrieval call binding the contract method 0xbd7cafcb.
//
// Solidity: function availableRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) AvailableRewardOf(opts *bind.CallOpts, id *big.Int, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "availableRewardOf", id, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableRewardOf is a free data retrieval call binding the contract method 0xbd7cafcb.
//
// Solidity: function availableRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) AvailableRewardOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.AvailableRewardOf(&_LetsFilControler.CallOpts, id, addr)
}

// AvailableRewardOf is a free data retrieval call binding the contract method 0xbd7cafcb.
//
// Solidity: function availableRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) AvailableRewardOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.AvailableRewardOf(&_LetsFilControler.CallOpts, id, addr)
}

// ExtendInfo is a free data retrieval call binding the contract method 0xdc9a9bbb.
//
// Solidity: function extendInfo(uint256 ) view returns(uint256 oldId, uint256 raiserOldShare, uint256 spOldShare, uint256 sponsorOldRewardShare, uint256 spOldRewardShare)
func (_LetsFilControler *LetsFilControlerCaller) ExtendInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OldId                 *big.Int
	RaiserOldShare        *big.Int
	SpOldShare            *big.Int
	SponsorOldRewardShare *big.Int
	SpOldRewardShare      *big.Int
}, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "extendInfo", arg0)

	outstruct := new(struct {
		OldId                 *big.Int
		RaiserOldShare        *big.Int
		SpOldShare            *big.Int
		SponsorOldRewardShare *big.Int
		SpOldRewardShare      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OldId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RaiserOldShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SpOldShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SponsorOldRewardShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SpOldRewardShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExtendInfo is a free data retrieval call binding the contract method 0xdc9a9bbb.
//
// Solidity: function extendInfo(uint256 ) view returns(uint256 oldId, uint256 raiserOldShare, uint256 spOldShare, uint256 sponsorOldRewardShare, uint256 spOldRewardShare)
func (_LetsFilControler *LetsFilControlerSession) ExtendInfo(arg0 *big.Int) (struct {
	OldId                 *big.Int
	RaiserOldShare        *big.Int
	SpOldShare            *big.Int
	SponsorOldRewardShare *big.Int
	SpOldRewardShare      *big.Int
}, error) {
	return _LetsFilControler.Contract.ExtendInfo(&_LetsFilControler.CallOpts, arg0)
}

// ExtendInfo is a free data retrieval call binding the contract method 0xdc9a9bbb.
//
// Solidity: function extendInfo(uint256 ) view returns(uint256 oldId, uint256 raiserOldShare, uint256 spOldShare, uint256 sponsorOldRewardShare, uint256 spOldRewardShare)
func (_LetsFilControler *LetsFilControlerCallerSession) ExtendInfo(arg0 *big.Int) (struct {
	OldId                 *big.Int
	RaiserOldShare        *big.Int
	SpOldShare            *big.Int
	SponsorOldRewardShare *big.Int
	SpOldRewardShare      *big.Int
}, error) {
	return _LetsFilControler.Contract.ExtendInfo(&_LetsFilControler.CallOpts, arg0)
}

// FundBack is a free data retrieval call binding the contract method 0x8813bd59.
//
// Solidity: function fundBack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) FundBack(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "fundBack", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundBack is a free data retrieval call binding the contract method 0x8813bd59.
//
// Solidity: function fundBack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) FundBack(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.FundBack(&_LetsFilControler.CallOpts, arg0)
}

// FundBack is a free data retrieval call binding the contract method 0x8813bd59.
//
// Solidity: function fundBack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) FundBack(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.FundBack(&_LetsFilControler.CallOpts, arg0)
}

// FundSealed is a free data retrieval call binding the contract method 0xa6a45a8b.
//
// Solidity: function fundSealed(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) FundSealed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "fundSealed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundSealed is a free data retrieval call binding the contract method 0xa6a45a8b.
//
// Solidity: function fundSealed(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) FundSealed(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.FundSealed(&_LetsFilControler.CallOpts, arg0)
}

// FundSealed is a free data retrieval call binding the contract method 0xa6a45a8b.
//
// Solidity: function fundSealed(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) FundSealed(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.FundSealed(&_LetsFilControler.CallOpts, arg0)
}

// GetBack is a free data retrieval call binding the contract method 0x2c2c868f.
//
// Solidity: function getBack(uint256 id, address account) view returns(uint256, uint256)
func (_LetsFilControler *LetsFilControlerCaller) GetBack(opts *bind.CallOpts, id *big.Int, account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "getBack", id, account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBack is a free data retrieval call binding the contract method 0x2c2c868f.
//
// Solidity: function getBack(uint256 id, address account) view returns(uint256, uint256)
func (_LetsFilControler *LetsFilControlerSession) GetBack(id *big.Int, account common.Address) (*big.Int, *big.Int, error) {
	return _LetsFilControler.Contract.GetBack(&_LetsFilControler.CallOpts, id, account)
}

// GetBack is a free data retrieval call binding the contract method 0x2c2c868f.
//
// Solidity: function getBack(uint256 id, address account) view returns(uint256, uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) GetBack(id *big.Int, account common.Address) (*big.Int, *big.Int, error) {
	return _LetsFilControler.Contract.GetBack(&_LetsFilControler.CallOpts, id, account)
}

// GetToolAddr is a free data retrieval call binding the contract method 0xcc6b88f6.
//
// Solidity: function getToolAddr() pure returns(address tool, address miner, address process)
func (_LetsFilControler *LetsFilControlerCaller) GetToolAddr(opts *bind.CallOpts) (struct {
	Tool    common.Address
	Miner   common.Address
	Process common.Address
}, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "getToolAddr")

	outstruct := new(struct {
		Tool    common.Address
		Miner   common.Address
		Process common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tool = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Process = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetToolAddr is a free data retrieval call binding the contract method 0xcc6b88f6.
//
// Solidity: function getToolAddr() pure returns(address tool, address miner, address process)
func (_LetsFilControler *LetsFilControlerSession) GetToolAddr() (struct {
	Tool    common.Address
	Miner   common.Address
	Process common.Address
}, error) {
	return _LetsFilControler.Contract.GetToolAddr(&_LetsFilControler.CallOpts)
}

// GetToolAddr is a free data retrieval call binding the contract method 0xcc6b88f6.
//
// Solidity: function getToolAddr() pure returns(address tool, address miner, address process)
func (_LetsFilControler *LetsFilControlerCallerSession) GetToolAddr() (struct {
	Tool    common.Address
	Miner   common.Address
	Process common.Address
}, error) {
	return _LetsFilControler.Contract.GetToolAddr(&_LetsFilControler.CallOpts)
}

// GotFilFiReward is a free data retrieval call binding the contract method 0x4565870f.
//
// Solidity: function gotFilFiReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) GotFilFiReward(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "gotFilFiReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GotFilFiReward is a free data retrieval call binding the contract method 0x4565870f.
//
// Solidity: function gotFilFiReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) GotFilFiReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotFilFiReward(&_LetsFilControler.CallOpts, arg0)
}

// GotFilFiReward is a free data retrieval call binding the contract method 0x4565870f.
//
// Solidity: function gotFilFiReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) GotFilFiReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotFilFiReward(&_LetsFilControler.CallOpts, arg0)
}

// GotMiner is a free data retrieval call binding the contract method 0x369af240.
//
// Solidity: function gotMiner() view returns(bool)
func (_LetsFilControler *LetsFilControlerCaller) GotMiner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "gotMiner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GotMiner is a free data retrieval call binding the contract method 0x369af240.
//
// Solidity: function gotMiner() view returns(bool)
func (_LetsFilControler *LetsFilControlerSession) GotMiner() (bool, error) {
	return _LetsFilControler.Contract.GotMiner(&_LetsFilControler.CallOpts)
}

// GotMiner is a free data retrieval call binding the contract method 0x369af240.
//
// Solidity: function gotMiner() view returns(bool)
func (_LetsFilControler *LetsFilControlerCallerSession) GotMiner() (bool, error) {
	return _LetsFilControler.Contract.GotMiner(&_LetsFilControler.CallOpts)
}

// GotRaiserReward is a free data retrieval call binding the contract method 0x5fc51b89.
//
// Solidity: function gotRaiserReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) GotRaiserReward(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "gotRaiserReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GotRaiserReward is a free data retrieval call binding the contract method 0x5fc51b89.
//
// Solidity: function gotRaiserReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) GotRaiserReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotRaiserReward(&_LetsFilControler.CallOpts, arg0)
}

// GotRaiserReward is a free data retrieval call binding the contract method 0x5fc51b89.
//
// Solidity: function gotRaiserReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) GotRaiserReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotRaiserReward(&_LetsFilControler.CallOpts, arg0)
}

// GotSpFundReward is a free data retrieval call binding the contract method 0xdbde1c60.
//
// Solidity: function gotSpFundReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) GotSpFundReward(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "gotSpFundReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GotSpFundReward is a free data retrieval call binding the contract method 0xdbde1c60.
//
// Solidity: function gotSpFundReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) GotSpFundReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotSpFundReward(&_LetsFilControler.CallOpts, arg0)
}

// GotSpFundReward is a free data retrieval call binding the contract method 0xdbde1c60.
//
// Solidity: function gotSpFundReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) GotSpFundReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotSpFundReward(&_LetsFilControler.CallOpts, arg0)
}

// GotSpReward is a free data retrieval call binding the contract method 0xdbf12fbd.
//
// Solidity: function gotSpReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) GotSpReward(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "gotSpReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GotSpReward is a free data retrieval call binding the contract method 0xdbf12fbd.
//
// Solidity: function gotSpReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) GotSpReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotSpReward(&_LetsFilControler.CallOpts, arg0)
}

// GotSpReward is a free data retrieval call binding the contract method 0xdbf12fbd.
//
// Solidity: function gotSpReward(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) GotSpReward(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.GotSpReward(&_LetsFilControler.CallOpts, arg0)
}

// InvestorFine is a free data retrieval call binding the contract method 0x89cbe074.
//
// Solidity: function investorFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) InvestorFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "investorFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvestorFine is a free data retrieval call binding the contract method 0x89cbe074.
//
// Solidity: function investorFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) InvestorFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.InvestorFine(&_LetsFilControler.CallOpts, arg0)
}

// InvestorFine is a free data retrieval call binding the contract method 0x89cbe074.
//
// Solidity: function investorFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) InvestorFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.InvestorFine(&_LetsFilControler.CallOpts, arg0)
}

// InvestorInfo is a free data retrieval call binding the contract method 0x676813a9.
//
// Solidity: function investorInfo(uint256 , address ) view returns(uint256 pledgeAmount, uint256 pledgeCalcAmount, uint256 interestDebt, uint256 withdrawAmount)
func (_LetsFilControler *LetsFilControlerCaller) InvestorInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	PledgeAmount     *big.Int
	PledgeCalcAmount *big.Int
	InterestDebt     *big.Int
	WithdrawAmount   *big.Int
}, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "investorInfo", arg0, arg1)

	outstruct := new(struct {
		PledgeAmount     *big.Int
		PledgeCalcAmount *big.Int
		InterestDebt     *big.Int
		WithdrawAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PledgeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PledgeCalcAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InterestDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WithdrawAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InvestorInfo is a free data retrieval call binding the contract method 0x676813a9.
//
// Solidity: function investorInfo(uint256 , address ) view returns(uint256 pledgeAmount, uint256 pledgeCalcAmount, uint256 interestDebt, uint256 withdrawAmount)
func (_LetsFilControler *LetsFilControlerSession) InvestorInfo(arg0 *big.Int, arg1 common.Address) (struct {
	PledgeAmount     *big.Int
	PledgeCalcAmount *big.Int
	InterestDebt     *big.Int
	WithdrawAmount   *big.Int
}, error) {
	return _LetsFilControler.Contract.InvestorInfo(&_LetsFilControler.CallOpts, arg0, arg1)
}

// InvestorInfo is a free data retrieval call binding the contract method 0x676813a9.
//
// Solidity: function investorInfo(uint256 , address ) view returns(uint256 pledgeAmount, uint256 pledgeCalcAmount, uint256 interestDebt, uint256 withdrawAmount)
func (_LetsFilControler *LetsFilControlerCallerSession) InvestorInfo(arg0 *big.Int, arg1 common.Address) (struct {
	PledgeAmount     *big.Int
	PledgeCalcAmount *big.Int
	InterestDebt     *big.Int
	WithdrawAmount   *big.Int
}, error) {
	return _LetsFilControler.Contract.InvestorInfo(&_LetsFilControler.CallOpts, arg0, arg1)
}

// MinerId is a free data retrieval call binding the contract method 0x96d7b514.
//
// Solidity: function minerId() view returns(uint64)
func (_LetsFilControler *LetsFilControlerCaller) MinerId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "minerId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinerId is a free data retrieval call binding the contract method 0x96d7b514.
//
// Solidity: function minerId() view returns(uint64)
func (_LetsFilControler *LetsFilControlerSession) MinerId() (uint64, error) {
	return _LetsFilControler.Contract.MinerId(&_LetsFilControler.CallOpts)
}

// MinerId is a free data retrieval call binding the contract method 0x96d7b514.
//
// Solidity: function minerId() view returns(uint64)
func (_LetsFilControler *LetsFilControlerCallerSession) MinerId() (uint64, error) {
	return _LetsFilControler.Contract.MinerId(&_LetsFilControler.CallOpts)
}

// NodeInfo is a free data retrieval call binding the contract method 0xb02439ae.
//
// Solidity: function nodeInfo(uint256 ) view returns(uint64 minerId, uint256 nodeSize, uint256 sectorSize, uint256 sealDays, uint256 nodeDays, uint256 opsSecurityFund, address spAddr, uint256 companyId)
func (_LetsFilControler *LetsFilControlerCaller) NodeInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MinerId         uint64
	NodeSize        *big.Int
	SectorSize      *big.Int
	SealDays        *big.Int
	NodeDays        *big.Int
	OpsSecurityFund *big.Int
	SpAddr          common.Address
	CompanyId       *big.Int
}, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "nodeInfo", arg0)

	outstruct := new(struct {
		MinerId         uint64
		NodeSize        *big.Int
		SectorSize      *big.Int
		SealDays        *big.Int
		NodeDays        *big.Int
		OpsSecurityFund *big.Int
		SpAddr          common.Address
		CompanyId       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinerId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NodeSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SectorSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SealDays = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NodeDays = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OpsSecurityFund = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SpAddr = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.CompanyId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NodeInfo is a free data retrieval call binding the contract method 0xb02439ae.
//
// Solidity: function nodeInfo(uint256 ) view returns(uint64 minerId, uint256 nodeSize, uint256 sectorSize, uint256 sealDays, uint256 nodeDays, uint256 opsSecurityFund, address spAddr, uint256 companyId)
func (_LetsFilControler *LetsFilControlerSession) NodeInfo(arg0 *big.Int) (struct {
	MinerId         uint64
	NodeSize        *big.Int
	SectorSize      *big.Int
	SealDays        *big.Int
	NodeDays        *big.Int
	OpsSecurityFund *big.Int
	SpAddr          common.Address
	CompanyId       *big.Int
}, error) {
	return _LetsFilControler.Contract.NodeInfo(&_LetsFilControler.CallOpts, arg0)
}

// NodeInfo is a free data retrieval call binding the contract method 0xb02439ae.
//
// Solidity: function nodeInfo(uint256 ) view returns(uint64 minerId, uint256 nodeSize, uint256 sectorSize, uint256 sealDays, uint256 nodeDays, uint256 opsSecurityFund, address spAddr, uint256 companyId)
func (_LetsFilControler *LetsFilControlerCallerSession) NodeInfo(arg0 *big.Int) (struct {
	MinerId         uint64
	NodeSize        *big.Int
	SectorSize      *big.Int
	SealDays        *big.Int
	NodeDays        *big.Int
	OpsSecurityFund *big.Int
	SpAddr          common.Address
	CompanyId       *big.Int
}, error) {
	return _LetsFilControler.Contract.NodeInfo(&_LetsFilControler.CallOpts, arg0)
}

// NodeState is a free data retrieval call binding the contract method 0x9600ac5b.
//
// Solidity: function nodeState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerCaller) NodeState(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "nodeState", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NodeState is a free data retrieval call binding the contract method 0x9600ac5b.
//
// Solidity: function nodeState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerSession) NodeState(arg0 *big.Int) (uint8, error) {
	return _LetsFilControler.Contract.NodeState(&_LetsFilControler.CallOpts, arg0)
}

// NodeState is a free data retrieval call binding the contract method 0x9600ac5b.
//
// Solidity: function nodeState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerCallerSession) NodeState(arg0 *big.Int) (uint8, error) {
	return _LetsFilControler.Contract.NodeState(&_LetsFilControler.CallOpts, arg0)
}

// OpsCalcFund is a free data retrieval call binding the contract method 0x29595e7f.
//
// Solidity: function opsCalcFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) OpsCalcFund(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "opsCalcFund", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsCalcFund is a free data retrieval call binding the contract method 0x29595e7f.
//
// Solidity: function opsCalcFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) OpsCalcFund(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsCalcFund(&_LetsFilControler.CallOpts, arg0)
}

// OpsCalcFund is a free data retrieval call binding the contract method 0x29595e7f.
//
// Solidity: function opsCalcFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) OpsCalcFund(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsCalcFund(&_LetsFilControler.CallOpts, arg0)
}

// OpsFundReward is a free data retrieval call binding the contract method 0x02c06d2b.
//
// Solidity: function opsFundReward(uint256 id) view returns(uint256 reward)
func (_LetsFilControler *LetsFilControlerCaller) OpsFundReward(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "opsFundReward", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsFundReward is a free data retrieval call binding the contract method 0x02c06d2b.
//
// Solidity: function opsFundReward(uint256 id) view returns(uint256 reward)
func (_LetsFilControler *LetsFilControlerSession) OpsFundReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsFundReward(&_LetsFilControler.CallOpts, id)
}

// OpsFundReward is a free data retrieval call binding the contract method 0x02c06d2b.
//
// Solidity: function opsFundReward(uint256 id) view returns(uint256 reward)
func (_LetsFilControler *LetsFilControlerCallerSession) OpsFundReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsFundReward(&_LetsFilControler.CallOpts, id)
}

// OpsSecurityFundRemain is a free data retrieval call binding the contract method 0x4b3296c6.
//
// Solidity: function opsSecurityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) OpsSecurityFundRemain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "opsSecurityFundRemain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsSecurityFundRemain is a free data retrieval call binding the contract method 0x4b3296c6.
//
// Solidity: function opsSecurityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) OpsSecurityFundRemain(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsSecurityFundRemain(&_LetsFilControler.CallOpts, arg0)
}

// OpsSecurityFundRemain is a free data retrieval call binding the contract method 0x4b3296c6.
//
// Solidity: function opsSecurityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) OpsSecurityFundRemain(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.OpsSecurityFundRemain(&_LetsFilControler.CallOpts, arg0)
}

// PledgeReleased is a free data retrieval call binding the contract method 0xa1befec0.
//
// Solidity: function pledgeReleased(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) PledgeReleased(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "pledgeReleased", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeReleased is a free data retrieval call binding the contract method 0xa1befec0.
//
// Solidity: function pledgeReleased(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) PledgeReleased(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeReleased(&_LetsFilControler.CallOpts, arg0)
}

// PledgeReleased is a free data retrieval call binding the contract method 0xa1befec0.
//
// Solidity: function pledgeReleased(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) PledgeReleased(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeReleased(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalAmount is a free data retrieval call binding the contract method 0xb98ac724.
//
// Solidity: function pledgeTotalAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) PledgeTotalAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "pledgeTotalAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeTotalAmount is a free data retrieval call binding the contract method 0xb98ac724.
//
// Solidity: function pledgeTotalAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) PledgeTotalAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalAmount(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalAmount is a free data retrieval call binding the contract method 0xb98ac724.
//
// Solidity: function pledgeTotalAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) PledgeTotalAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalAmount(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalCalcAmount is a free data retrieval call binding the contract method 0x38f9706a.
//
// Solidity: function pledgeTotalCalcAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) PledgeTotalCalcAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "pledgeTotalCalcAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeTotalCalcAmount is a free data retrieval call binding the contract method 0x38f9706a.
//
// Solidity: function pledgeTotalCalcAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) PledgeTotalCalcAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalCalcAmount(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalCalcAmount is a free data retrieval call binding the contract method 0x38f9706a.
//
// Solidity: function pledgeTotalCalcAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) PledgeTotalCalcAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalCalcAmount(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalDebt is a free data retrieval call binding the contract method 0x634969b5.
//
// Solidity: function pledgeTotalDebt(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) PledgeTotalDebt(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "pledgeTotalDebt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeTotalDebt is a free data retrieval call binding the contract method 0x634969b5.
//
// Solidity: function pledgeTotalDebt(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) PledgeTotalDebt(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalDebt(&_LetsFilControler.CallOpts, arg0)
}

// PledgeTotalDebt is a free data retrieval call binding the contract method 0x634969b5.
//
// Solidity: function pledgeTotalDebt(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) PledgeTotalDebt(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.PledgeTotalDebt(&_LetsFilControler.CallOpts, arg0)
}

// RaiseCloseTime is a free data retrieval call binding the contract method 0x8f530c90.
//
// Solidity: function raiseCloseTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) RaiseCloseTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiseCloseTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiseCloseTime is a free data retrieval call binding the contract method 0x8f530c90.
//
// Solidity: function raiseCloseTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) RaiseCloseTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiseCloseTime(&_LetsFilControler.CallOpts, arg0)
}

// RaiseCloseTime is a free data retrieval call binding the contract method 0x8f530c90.
//
// Solidity: function raiseCloseTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiseCloseTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiseCloseTime(&_LetsFilControler.CallOpts, arg0)
}

// RaiseInfo is a free data retrieval call binding the contract method 0x6d6d26ed.
//
// Solidity: function raiseInfo(uint256 ) view returns(uint256 id, uint256 targetAmount, uint256 minRaiseRate, uint256 securityFund, uint256 raiseDays, uint256 investorShare, uint256 spFundShare, uint256 raiserShare, uint256 servicerShare, uint256 filFiShare, address sponsor, string raiseCompany)
func (_LetsFilControler *LetsFilControlerCaller) RaiseInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	TargetAmount  *big.Int
	MinRaiseRate  *big.Int
	SecurityFund  *big.Int
	RaiseDays     *big.Int
	InvestorShare *big.Int
	SpFundShare   *big.Int
	RaiserShare   *big.Int
	ServicerShare *big.Int
	FilFiShare    *big.Int
	Sponsor       common.Address
	RaiseCompany  string
}, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiseInfo", arg0)

	outstruct := new(struct {
		Id            *big.Int
		TargetAmount  *big.Int
		MinRaiseRate  *big.Int
		SecurityFund  *big.Int
		RaiseDays     *big.Int
		InvestorShare *big.Int
		SpFundShare   *big.Int
		RaiserShare   *big.Int
		ServicerShare *big.Int
		FilFiShare    *big.Int
		Sponsor       common.Address
		RaiseCompany  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TargetAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinRaiseRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecurityFund = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RaiseDays = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.InvestorShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SpFundShare = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RaiserShare = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ServicerShare = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FilFiShare = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Sponsor = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.RaiseCompany = *abi.ConvertType(out[11], new(string)).(*string)

	return *outstruct, err

}

// RaiseInfo is a free data retrieval call binding the contract method 0x6d6d26ed.
//
// Solidity: function raiseInfo(uint256 ) view returns(uint256 id, uint256 targetAmount, uint256 minRaiseRate, uint256 securityFund, uint256 raiseDays, uint256 investorShare, uint256 spFundShare, uint256 raiserShare, uint256 servicerShare, uint256 filFiShare, address sponsor, string raiseCompany)
func (_LetsFilControler *LetsFilControlerSession) RaiseInfo(arg0 *big.Int) (struct {
	Id            *big.Int
	TargetAmount  *big.Int
	MinRaiseRate  *big.Int
	SecurityFund  *big.Int
	RaiseDays     *big.Int
	InvestorShare *big.Int
	SpFundShare   *big.Int
	RaiserShare   *big.Int
	ServicerShare *big.Int
	FilFiShare    *big.Int
	Sponsor       common.Address
	RaiseCompany  string
}, error) {
	return _LetsFilControler.Contract.RaiseInfo(&_LetsFilControler.CallOpts, arg0)
}

// RaiseInfo is a free data retrieval call binding the contract method 0x6d6d26ed.
//
// Solidity: function raiseInfo(uint256 ) view returns(uint256 id, uint256 targetAmount, uint256 minRaiseRate, uint256 securityFund, uint256 raiseDays, uint256 investorShare, uint256 spFundShare, uint256 raiserShare, uint256 servicerShare, uint256 filFiShare, address sponsor, string raiseCompany)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiseInfo(arg0 *big.Int) (struct {
	Id            *big.Int
	TargetAmount  *big.Int
	MinRaiseRate  *big.Int
	SecurityFund  *big.Int
	RaiseDays     *big.Int
	InvestorShare *big.Int
	SpFundShare   *big.Int
	RaiserShare   *big.Int
	ServicerShare *big.Int
	FilFiShare    *big.Int
	Sponsor       common.Address
	RaiseCompany  string
}, error) {
	return _LetsFilControler.Contract.RaiseInfo(&_LetsFilControler.CallOpts, arg0)
}

// RaiseStartTime is a free data retrieval call binding the contract method 0x5b9a3b67.
//
// Solidity: function raiseStartTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) RaiseStartTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiseStartTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiseStartTime is a free data retrieval call binding the contract method 0x5b9a3b67.
//
// Solidity: function raiseStartTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) RaiseStartTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiseStartTime(&_LetsFilControler.CallOpts, arg0)
}

// RaiseStartTime is a free data retrieval call binding the contract method 0x5b9a3b67.
//
// Solidity: function raiseStartTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiseStartTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiseStartTime(&_LetsFilControler.CallOpts, arg0)
}

// RaiseState is a free data retrieval call binding the contract method 0x46981829.
//
// Solidity: function raiseState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerCaller) RaiseState(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiseState", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RaiseState is a free data retrieval call binding the contract method 0x46981829.
//
// Solidity: function raiseState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerSession) RaiseState(arg0 *big.Int) (uint8, error) {
	return _LetsFilControler.Contract.RaiseState(&_LetsFilControler.CallOpts, arg0)
}

// RaiseState is a free data retrieval call binding the contract method 0x46981829.
//
// Solidity: function raiseState(uint256 ) view returns(uint8)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiseState(arg0 *big.Int) (uint8, error) {
	return _LetsFilControler.Contract.RaiseState(&_LetsFilControler.CallOpts, arg0)
}

// Raiser is a free data retrieval call binding the contract method 0xbe4f3b0b.
//
// Solidity: function raiser() view returns(address)
func (_LetsFilControler *LetsFilControlerCaller) Raiser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Raiser is a free data retrieval call binding the contract method 0xbe4f3b0b.
//
// Solidity: function raiser() view returns(address)
func (_LetsFilControler *LetsFilControlerSession) Raiser() (common.Address, error) {
	return _LetsFilControler.Contract.Raiser(&_LetsFilControler.CallOpts)
}

// Raiser is a free data retrieval call binding the contract method 0xbe4f3b0b.
//
// Solidity: function raiser() view returns(address)
func (_LetsFilControler *LetsFilControlerCallerSession) Raiser() (common.Address, error) {
	return _LetsFilControler.Contract.Raiser(&_LetsFilControler.CallOpts)
}

// RaiserFine is a free data retrieval call binding the contract method 0x12c33f8d.
//
// Solidity: function raiserFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) RaiserFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiserFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiserFine is a free data retrieval call binding the contract method 0x12c33f8d.
//
// Solidity: function raiserFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) RaiserFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserFine(&_LetsFilControler.CallOpts, arg0)
}

// RaiserFine is a free data retrieval call binding the contract method 0x12c33f8d.
//
// Solidity: function raiserFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiserFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserFine(&_LetsFilControler.CallOpts, arg0)
}

// RaiserRewardAvailableLeft is a free data retrieval call binding the contract method 0x08a882e0.
//
// Solidity: function raiserRewardAvailableLeft(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) RaiserRewardAvailableLeft(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiserRewardAvailableLeft", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiserRewardAvailableLeft is a free data retrieval call binding the contract method 0x08a882e0.
//
// Solidity: function raiserRewardAvailableLeft(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) RaiserRewardAvailableLeft(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserRewardAvailableLeft(&_LetsFilControler.CallOpts, id)
}

// RaiserRewardAvailableLeft is a free data retrieval call binding the contract method 0x08a882e0.
//
// Solidity: function raiserRewardAvailableLeft(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiserRewardAvailableLeft(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserRewardAvailableLeft(&_LetsFilControler.CallOpts, id)
}

// RaiserWillReleaseReward is a free data retrieval call binding the contract method 0x8e3a7894.
//
// Solidity: function raiserWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) RaiserWillReleaseReward(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "raiserWillReleaseReward", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiserWillReleaseReward is a free data retrieval call binding the contract method 0x8e3a7894.
//
// Solidity: function raiserWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) RaiserWillReleaseReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserWillReleaseReward(&_LetsFilControler.CallOpts, id)
}

// RaiserWillReleaseReward is a free data retrieval call binding the contract method 0x8e3a7894.
//
// Solidity: function raiserWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) RaiserWillReleaseReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.RaiserWillReleaseReward(&_LetsFilControler.CallOpts, id)
}

// ReturnAmount is a free data retrieval call binding the contract method 0x030f5702.
//
// Solidity: function returnAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) ReturnAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "returnAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReturnAmount is a free data retrieval call binding the contract method 0x030f5702.
//
// Solidity: function returnAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) ReturnAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.ReturnAmount(&_LetsFilControler.CallOpts, arg0)
}

// ReturnAmount is a free data retrieval call binding the contract method 0x030f5702.
//
// Solidity: function returnAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) ReturnAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.ReturnAmount(&_LetsFilControler.CallOpts, arg0)
}

// SafeFundFine is a free data retrieval call binding the contract method 0xe062c074.
//
// Solidity: function safeFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SafeFundFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "safeFundFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SafeFundFine is a free data retrieval call binding the contract method 0xe062c074.
//
// Solidity: function safeFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SafeFundFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SafeFundFine(&_LetsFilControler.CallOpts, arg0)
}

// SafeFundFine is a free data retrieval call binding the contract method 0xe062c074.
//
// Solidity: function safeFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SafeFundFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SafeFundFine(&_LetsFilControler.CallOpts, arg0)
}

// SafeSealFund is a free data retrieval call binding the contract method 0x67326918.
//
// Solidity: function safeSealFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SafeSealFund(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "safeSealFund", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SafeSealFund is a free data retrieval call binding the contract method 0x67326918.
//
// Solidity: function safeSealFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SafeSealFund(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SafeSealFund(&_LetsFilControler.CallOpts, arg0)
}

// SafeSealFund is a free data retrieval call binding the contract method 0x67326918.
//
// Solidity: function safeSealFund(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SafeSealFund(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SafeSealFund(&_LetsFilControler.CallOpts, arg0)
}

// SealEndTime is a free data retrieval call binding the contract method 0x2294d793.
//
// Solidity: function sealEndTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SealEndTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "sealEndTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SealEndTime is a free data retrieval call binding the contract method 0x2294d793.
//
// Solidity: function sealEndTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SealEndTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SealEndTime(&_LetsFilControler.CallOpts, arg0)
}

// SealEndTime is a free data retrieval call binding the contract method 0x2294d793.
//
// Solidity: function sealEndTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SealEndTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SealEndTime(&_LetsFilControler.CallOpts, arg0)
}

// SealedAmount is a free data retrieval call binding the contract method 0x2de513c2.
//
// Solidity: function sealedAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SealedAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "sealedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SealedAmount is a free data retrieval call binding the contract method 0x2de513c2.
//
// Solidity: function sealedAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SealedAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SealedAmount(&_LetsFilControler.CallOpts, arg0)
}

// SealedAmount is a free data retrieval call binding the contract method 0x2de513c2.
//
// Solidity: function sealedAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SealedAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SealedAmount(&_LetsFilControler.CallOpts, arg0)
}

// SecPack is a free data retrieval call binding the contract method 0x0b8581a2.
//
// Solidity: function secPack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SecPack(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "secPack", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecPack is a free data retrieval call binding the contract method 0x0b8581a2.
//
// Solidity: function secPack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SecPack(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SecPack(&_LetsFilControler.CallOpts, arg0)
}

// SecPack is a free data retrieval call binding the contract method 0x0b8581a2.
//
// Solidity: function secPack(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SecPack(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SecPack(&_LetsFilControler.CallOpts, arg0)
}

// SecurityFundRemain is a free data retrieval call binding the contract method 0x283dcdaa.
//
// Solidity: function securityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SecurityFundRemain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "securityFundRemain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecurityFundRemain is a free data retrieval call binding the contract method 0x283dcdaa.
//
// Solidity: function securityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SecurityFundRemain(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SecurityFundRemain(&_LetsFilControler.CallOpts, arg0)
}

// SecurityFundRemain is a free data retrieval call binding the contract method 0x283dcdaa.
//
// Solidity: function securityFundRemain(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SecurityFundRemain(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SecurityFundRemain(&_LetsFilControler.CallOpts, arg0)
}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_LetsFilControler *LetsFilControlerCaller) Sp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "sp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_LetsFilControler *LetsFilControlerSession) Sp() (common.Address, error) {
	return _LetsFilControler.Contract.Sp(&_LetsFilControler.CallOpts)
}

// Sp is a free data retrieval call binding the contract method 0xe1dadf3b.
//
// Solidity: function sp() view returns(address)
func (_LetsFilControler *LetsFilControlerCallerSession) Sp() (common.Address, error) {
	return _LetsFilControler.Contract.Sp(&_LetsFilControler.CallOpts)
}

// SpFine is a free data retrieval call binding the contract method 0xd589b9c4.
//
// Solidity: function spFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpFine is a free data retrieval call binding the contract method 0xd589b9c4.
//
// Solidity: function spFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFine(&_LetsFilControler.CallOpts, arg0)
}

// SpFine is a free data retrieval call binding the contract method 0xd589b9c4.
//
// Solidity: function spFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFine(&_LetsFilControler.CallOpts, arg0)
}

// SpFundFine is a free data retrieval call binding the contract method 0x2b05000a.
//
// Solidity: function spFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpFundFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spFundFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpFundFine is a free data retrieval call binding the contract method 0x2b05000a.
//
// Solidity: function spFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpFundFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFundFine(&_LetsFilControler.CallOpts, arg0)
}

// SpFundFine is a free data retrieval call binding the contract method 0x2b05000a.
//
// Solidity: function spFundFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpFundFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFundFine(&_LetsFilControler.CallOpts, arg0)
}

// SpFundRewardFine is a free data retrieval call binding the contract method 0x01c93f20.
//
// Solidity: function spFundRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpFundRewardFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spFundRewardFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpFundRewardFine is a free data retrieval call binding the contract method 0x01c93f20.
//
// Solidity: function spFundRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpFundRewardFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFundRewardFine(&_LetsFilControler.CallOpts, arg0)
}

// SpFundRewardFine is a free data retrieval call binding the contract method 0x01c93f20.
//
// Solidity: function spFundRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpFundRewardFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpFundRewardFine(&_LetsFilControler.CallOpts, arg0)
}

// SpRemainFine is a free data retrieval call binding the contract method 0x3470773c.
//
// Solidity: function spRemainFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpRemainFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spRemainFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpRemainFine is a free data retrieval call binding the contract method 0x3470773c.
//
// Solidity: function spRemainFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpRemainFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRemainFine(&_LetsFilControler.CallOpts, arg0)
}

// SpRemainFine is a free data retrieval call binding the contract method 0x3470773c.
//
// Solidity: function spRemainFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpRemainFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRemainFine(&_LetsFilControler.CallOpts, arg0)
}

// SpRewardAvailableLeft is a free data retrieval call binding the contract method 0x835eb451.
//
// Solidity: function spRewardAvailableLeft(uint256 id) view returns(uint256 amountReturn)
func (_LetsFilControler *LetsFilControlerCaller) SpRewardAvailableLeft(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spRewardAvailableLeft", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpRewardAvailableLeft is a free data retrieval call binding the contract method 0x835eb451.
//
// Solidity: function spRewardAvailableLeft(uint256 id) view returns(uint256 amountReturn)
func (_LetsFilControler *LetsFilControlerSession) SpRewardAvailableLeft(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardAvailableLeft(&_LetsFilControler.CallOpts, id)
}

// SpRewardAvailableLeft is a free data retrieval call binding the contract method 0x835eb451.
//
// Solidity: function spRewardAvailableLeft(uint256 id) view returns(uint256 amountReturn)
func (_LetsFilControler *LetsFilControlerCallerSession) SpRewardAvailableLeft(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardAvailableLeft(&_LetsFilControler.CallOpts, id)
}

// SpRewardFine is a free data retrieval call binding the contract method 0x6a3af92c.
//
// Solidity: function spRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpRewardFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spRewardFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpRewardFine is a free data retrieval call binding the contract method 0x6a3af92c.
//
// Solidity: function spRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpRewardFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardFine(&_LetsFilControler.CallOpts, arg0)
}

// SpRewardFine is a free data retrieval call binding the contract method 0x6a3af92c.
//
// Solidity: function spRewardFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpRewardFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardFine(&_LetsFilControler.CallOpts, arg0)
}

// SpRewardLock is a free data retrieval call binding the contract method 0x99ac1518.
//
// Solidity: function spRewardLock(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpRewardLock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spRewardLock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpRewardLock is a free data retrieval call binding the contract method 0x99ac1518.
//
// Solidity: function spRewardLock(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpRewardLock(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardLock(&_LetsFilControler.CallOpts, arg0)
}

// SpRewardLock is a free data retrieval call binding the contract method 0x99ac1518.
//
// Solidity: function spRewardLock(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpRewardLock(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpRewardLock(&_LetsFilControler.CallOpts, arg0)
}

// SpWillReleaseReward is a free data retrieval call binding the contract method 0xff9775a5.
//
// Solidity: function spWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SpWillReleaseReward(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "spWillReleaseReward", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpWillReleaseReward is a free data retrieval call binding the contract method 0xff9775a5.
//
// Solidity: function spWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SpWillReleaseReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpWillReleaseReward(&_LetsFilControler.CallOpts, id)
}

// SpWillReleaseReward is a free data retrieval call binding the contract method 0xff9775a5.
//
// Solidity: function spWillReleaseReward(uint256 id) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SpWillReleaseReward(id *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SpWillReleaseReward(&_LetsFilControler.CallOpts, id)
}

// StartSealTime is a free data retrieval call binding the contract method 0xb060548b.
//
// Solidity: function startSealTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) StartSealTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "startSealTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartSealTime is a free data retrieval call binding the contract method 0xb060548b.
//
// Solidity: function startSealTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) StartSealTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.StartSealTime(&_LetsFilControler.CallOpts, arg0)
}

// StartSealTime is a free data retrieval call binding the contract method 0xb060548b.
//
// Solidity: function startSealTime(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) StartSealTime(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.StartSealTime(&_LetsFilControler.CallOpts, arg0)
}

// SubFine is a free data retrieval call binding the contract method 0x2afa0c87.
//
// Solidity: function subFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) SubFine(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "subFine", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubFine is a free data retrieval call binding the contract method 0x2afa0c87.
//
// Solidity: function subFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) SubFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SubFine(&_LetsFilControler.CallOpts, arg0)
}

// SubFine is a free data retrieval call binding the contract method 0x2afa0c87.
//
// Solidity: function subFine(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) SubFine(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.SubFine(&_LetsFilControler.CallOpts, arg0)
}

// ToSealAmount is a free data retrieval call binding the contract method 0x70fe3805.
//
// Solidity: function toSealAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) ToSealAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "toSealAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToSealAmount is a free data retrieval call binding the contract method 0x70fe3805.
//
// Solidity: function toSealAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) ToSealAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.ToSealAmount(&_LetsFilControler.CallOpts, arg0)
}

// ToSealAmount is a free data retrieval call binding the contract method 0x70fe3805.
//
// Solidity: function toSealAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) ToSealAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.ToSealAmount(&_LetsFilControler.CallOpts, arg0)
}

// TotalInterest is a free data retrieval call binding the contract method 0x8f103ce6.
//
// Solidity: function totalInterest(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) TotalInterest(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "totalInterest", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInterest is a free data retrieval call binding the contract method 0x8f103ce6.
//
// Solidity: function totalInterest(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) TotalInterest(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalInterest(&_LetsFilControler.CallOpts, arg0)
}

// TotalInterest is a free data retrieval call binding the contract method 0x8f103ce6.
//
// Solidity: function totalInterest(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) TotalInterest(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalInterest(&_LetsFilControler.CallOpts, arg0)
}

// TotalReleasedRewardAmount is a free data retrieval call binding the contract method 0x0b22f848.
//
// Solidity: function totalReleasedRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) TotalReleasedRewardAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "totalReleasedRewardAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleasedRewardAmount is a free data retrieval call binding the contract method 0x0b22f848.
//
// Solidity: function totalReleasedRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) TotalReleasedRewardAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalReleasedRewardAmount(&_LetsFilControler.CallOpts, arg0)
}

// TotalReleasedRewardAmount is a free data retrieval call binding the contract method 0x0b22f848.
//
// Solidity: function totalReleasedRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) TotalReleasedRewardAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalReleasedRewardAmount(&_LetsFilControler.CallOpts, arg0)
}

// TotalRewardAmount is a free data retrieval call binding the contract method 0xb625315a.
//
// Solidity: function totalRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) TotalRewardAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "totalRewardAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardAmount is a free data retrieval call binding the contract method 0xb625315a.
//
// Solidity: function totalRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) TotalRewardAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalRewardAmount(&_LetsFilControler.CallOpts, arg0)
}

// TotalRewardAmount is a free data retrieval call binding the contract method 0xb625315a.
//
// Solidity: function totalRewardAmount(uint256 ) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) TotalRewardAmount(arg0 *big.Int) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalRewardAmount(&_LetsFilControler.CallOpts, arg0)
}

// TotalRewardOf is a free data retrieval call binding the contract method 0xb80990cc.
//
// Solidity: function totalRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) TotalRewardOf(opts *bind.CallOpts, id *big.Int, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "totalRewardOf", id, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardOf is a free data retrieval call binding the contract method 0xb80990cc.
//
// Solidity: function totalRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) TotalRewardOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalRewardOf(&_LetsFilControler.CallOpts, id, addr)
}

// TotalRewardOf is a free data retrieval call binding the contract method 0xb80990cc.
//
// Solidity: function totalRewardOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) TotalRewardOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.TotalRewardOf(&_LetsFilControler.CallOpts, id, addr)
}

// WillReleaseOf is a free data retrieval call binding the contract method 0xda28211c.
//
// Solidity: function willReleaseOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCaller) WillReleaseOf(opts *bind.CallOpts, id *big.Int, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilControler.contract.Call(opts, &out, "willReleaseOf", id, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WillReleaseOf is a free data retrieval call binding the contract method 0xda28211c.
//
// Solidity: function willReleaseOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerSession) WillReleaseOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.WillReleaseOf(&_LetsFilControler.CallOpts, id, addr)
}

// WillReleaseOf is a free data retrieval call binding the contract method 0xda28211c.
//
// Solidity: function willReleaseOf(uint256 id, address addr) view returns(uint256)
func (_LetsFilControler *LetsFilControlerCallerSession) WillReleaseOf(id *big.Int, addr common.Address) (*big.Int, error) {
	return _LetsFilControler.Contract.WillReleaseOf(&_LetsFilControler.CallOpts, id, addr)
}

// AddOpsSecurityFund is a paid mutator transaction binding the contract method 0x41031798.
//
// Solidity: function addOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) AddOpsSecurityFund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "addOpsSecurityFund", id)
}

// AddOpsSecurityFund is a paid mutator transaction binding the contract method 0x41031798.
//
// Solidity: function addOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerSession) AddOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.AddOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// AddOpsSecurityFund is a paid mutator transaction binding the contract method 0x41031798.
//
// Solidity: function addOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) AddOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.AddOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// BackOwner is a paid mutator transaction binding the contract method 0x6e244f80.
//
// Solidity: function backOwner() returns()
func (_LetsFilControler *LetsFilControlerTransactor) BackOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "backOwner")
}

// BackOwner is a paid mutator transaction binding the contract method 0x6e244f80.
//
// Solidity: function backOwner() returns()
func (_LetsFilControler *LetsFilControlerSession) BackOwner() (*types.Transaction, error) {
	return _LetsFilControler.Contract.BackOwner(&_LetsFilControler.TransactOpts)
}

// BackOwner is a paid mutator transaction binding the contract method 0x6e244f80.
//
// Solidity: function backOwner() returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) BackOwner() (*types.Transaction, error) {
	return _LetsFilControler.Contract.BackOwner(&_LetsFilControler.TransactOpts)
}

// ChangeOwnerById is a paid mutator transaction binding the contract method 0x1c8bf2bb.
//
// Solidity: function changeOwnerById(uint64 _minerId, uint64 _actorId) returns()
func (_LetsFilControler *LetsFilControlerTransactor) ChangeOwnerById(opts *bind.TransactOpts, _minerId uint64, _actorId uint64) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "changeOwnerById", _minerId, _actorId)
}

// ChangeOwnerById is a paid mutator transaction binding the contract method 0x1c8bf2bb.
//
// Solidity: function changeOwnerById(uint64 _minerId, uint64 _actorId) returns()
func (_LetsFilControler *LetsFilControlerSession) ChangeOwnerById(_minerId uint64, _actorId uint64) (*types.Transaction, error) {
	return _LetsFilControler.Contract.ChangeOwnerById(&_LetsFilControler.TransactOpts, _minerId, _actorId)
}

// ChangeOwnerById is a paid mutator transaction binding the contract method 0x1c8bf2bb.
//
// Solidity: function changeOwnerById(uint64 _minerId, uint64 _actorId) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) ChangeOwnerById(_minerId uint64, _actorId uint64) (*types.Transaction, error) {
	return _LetsFilControler.Contract.ChangeOwnerById(&_LetsFilControler.TransactOpts, _minerId, _actorId)
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xbfdc4a1b.
//
// Solidity: function closeRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) CloseRaisePlan(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "closeRaisePlan", id)
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xbfdc4a1b.
//
// Solidity: function closeRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) CloseRaisePlan(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.CloseRaisePlan(&_LetsFilControler.TransactOpts, id)
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xbfdc4a1b.
//
// Solidity: function closeRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) CloseRaisePlan(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.CloseRaisePlan(&_LetsFilControler.TransactOpts, id)
}

// DestroyNode is a paid mutator transaction binding the contract method 0xd2c50bf4.
//
// Solidity: function destroyNode(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) DestroyNode(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "destroyNode", id)
}

// DestroyNode is a paid mutator transaction binding the contract method 0xd2c50bf4.
//
// Solidity: function destroyNode(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) DestroyNode(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.DestroyNode(&_LetsFilControler.TransactOpts, id)
}

// DestroyNode is a paid mutator transaction binding the contract method 0xd2c50bf4.
//
// Solidity: function destroyNode(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) DestroyNode(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.DestroyNode(&_LetsFilControler.TransactOpts, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x06dd3e52.
//
// Solidity: function initialize(uint256 _raiseID, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns()
func (_LetsFilControler *LetsFilControlerTransactor) Initialize(opts *bind.TransactOpts, _raiseID *big.Int, _raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "initialize", _raiseID, _raiseInfo, _nodeInfo, _extendInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x06dd3e52.
//
// Solidity: function initialize(uint256 _raiseID, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns()
func (_LetsFilControler *LetsFilControlerSession) Initialize(_raiseID *big.Int, _raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Initialize(&_LetsFilControler.TransactOpts, _raiseID, _raiseInfo, _nodeInfo, _extendInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x06dd3e52.
//
// Solidity: function initialize(uint256 _raiseID, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) Initialize(_raiseID *big.Int, _raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Initialize(&_LetsFilControler.TransactOpts, _raiseID, _raiseInfo, _nodeInfo, _extendInfo)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0xe6febc9b.
//
// Solidity: function investorWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) InvestorWithdraw(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "investorWithdraw", id)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0xe6febc9b.
//
// Solidity: function investorWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) InvestorWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.InvestorWithdraw(&_LetsFilControler.TransactOpts, id)
}

// InvestorWithdraw is a paid mutator transaction binding the contract method 0xe6febc9b.
//
// Solidity: function investorWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) InvestorWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.InvestorWithdraw(&_LetsFilControler.TransactOpts, id)
}

// PayOpsSecurityFund is a paid mutator transaction binding the contract method 0x0bf17040.
//
// Solidity: function payOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) PayOpsSecurityFund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "payOpsSecurityFund", id)
}

// PayOpsSecurityFund is a paid mutator transaction binding the contract method 0x0bf17040.
//
// Solidity: function payOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerSession) PayOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PayOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// PayOpsSecurityFund is a paid mutator transaction binding the contract method 0x0bf17040.
//
// Solidity: function payOpsSecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PayOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PayOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// PaySecurityFund is a paid mutator transaction binding the contract method 0x0e430cce.
//
// Solidity: function paySecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) PaySecurityFund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "paySecurityFund", id)
}

// PaySecurityFund is a paid mutator transaction binding the contract method 0x0e430cce.
//
// Solidity: function paySecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerSession) PaySecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PaySecurityFund(&_LetsFilControler.TransactOpts, id)
}

// PaySecurityFund is a paid mutator transaction binding the contract method 0x0e430cce.
//
// Solidity: function paySecurityFund(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PaySecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PaySecurityFund(&_LetsFilControler.TransactOpts, id)
}

// PushBlockReward is a paid mutator transaction binding the contract method 0xa0c2e20f.
//
// Solidity: function pushBlockReward(uint256 id, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushBlockReward(opts *bind.TransactOpts, id *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushBlockReward", id, released, willRelease)
}

// PushBlockReward is a paid mutator transaction binding the contract method 0xa0c2e20f.
//
// Solidity: function pushBlockReward(uint256 id, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerSession) PushBlockReward(id *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushBlockReward(&_LetsFilControler.TransactOpts, id, released, willRelease)
}

// PushBlockReward is a paid mutator transaction binding the contract method 0xa0c2e20f.
//
// Solidity: function pushBlockReward(uint256 id, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushBlockReward(id *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushBlockReward(&_LetsFilControler.TransactOpts, id, released, willRelease)
}

// PushFinalProgress is a paid mutator transaction binding the contract method 0x8e05c96b.
//
// Solidity: function pushFinalProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushFinalProgress(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushFinalProgress", id, amount)
}

// PushFinalProgress is a paid mutator transaction binding the contract method 0x8e05c96b.
//
// Solidity: function pushFinalProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerSession) PushFinalProgress(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushFinalProgress(&_LetsFilControler.TransactOpts, id, amount)
}

// PushFinalProgress is a paid mutator transaction binding the contract method 0x8e05c96b.
//
// Solidity: function pushFinalProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushFinalProgress(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushFinalProgress(&_LetsFilControler.TransactOpts, id, amount)
}

// PushOldAssetPackValue is a paid mutator transaction binding the contract method 0xc89d7a20.
//
// Solidity: function pushOldAssetPackValue(uint256 id, uint256 totalPledge, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushOldAssetPackValue(opts *bind.TransactOpts, id *big.Int, totalPledge *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushOldAssetPackValue", id, totalPledge, released, willRelease)
}

// PushOldAssetPackValue is a paid mutator transaction binding the contract method 0xc89d7a20.
//
// Solidity: function pushOldAssetPackValue(uint256 id, uint256 totalPledge, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerSession) PushOldAssetPackValue(id *big.Int, totalPledge *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushOldAssetPackValue(&_LetsFilControler.TransactOpts, id, totalPledge, released, willRelease)
}

// PushOldAssetPackValue is a paid mutator transaction binding the contract method 0xc89d7a20.
//
// Solidity: function pushOldAssetPackValue(uint256 id, uint256 totalPledge, uint256 released, uint256 willRelease) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushOldAssetPackValue(id *big.Int, totalPledge *big.Int, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushOldAssetPackValue(&_LetsFilControler.TransactOpts, id, totalPledge, released, willRelease)
}

// PushPledgeReleased is a paid mutator transaction binding the contract method 0xa6405107.
//
// Solidity: function pushPledgeReleased(uint256 id, uint256 released) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushPledgeReleased(opts *bind.TransactOpts, id *big.Int, released *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushPledgeReleased", id, released)
}

// PushPledgeReleased is a paid mutator transaction binding the contract method 0xa6405107.
//
// Solidity: function pushPledgeReleased(uint256 id, uint256 released) returns()
func (_LetsFilControler *LetsFilControlerSession) PushPledgeReleased(id *big.Int, released *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushPledgeReleased(&_LetsFilControler.TransactOpts, id, released)
}

// PushPledgeReleased is a paid mutator transaction binding the contract method 0xa6405107.
//
// Solidity: function pushPledgeReleased(uint256 id, uint256 released) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushPledgeReleased(id *big.Int, released *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushPledgeReleased(&_LetsFilControler.TransactOpts, id, released)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0xea875a38.
//
// Solidity: function pushSealProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushSealProgress(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushSealProgress", id, amount)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0xea875a38.
//
// Solidity: function pushSealProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerSession) PushSealProgress(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushSealProgress(&_LetsFilControler.TransactOpts, id, amount)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0xea875a38.
//
// Solidity: function pushSealProgress(uint256 id, uint256 amount) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushSealProgress(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushSealProgress(&_LetsFilControler.TransactOpts, id, amount)
}

// PushSpFine is a paid mutator transaction binding the contract method 0x2f6d00a2.
//
// Solidity: function pushSpFine(uint256 id, uint256 fineAmount) returns()
func (_LetsFilControler *LetsFilControlerTransactor) PushSpFine(opts *bind.TransactOpts, id *big.Int, fineAmount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "pushSpFine", id, fineAmount)
}

// PushSpFine is a paid mutator transaction binding the contract method 0x2f6d00a2.
//
// Solidity: function pushSpFine(uint256 id, uint256 fineAmount) returns()
func (_LetsFilControler *LetsFilControlerSession) PushSpFine(id *big.Int, fineAmount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushSpFine(&_LetsFilControler.TransactOpts, id, fineAmount)
}

// PushSpFine is a paid mutator transaction binding the contract method 0x2f6d00a2.
//
// Solidity: function pushSpFine(uint256 id, uint256 fineAmount) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) PushSpFine(id *big.Int, fineAmount *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.PushSpFine(&_LetsFilControler.TransactOpts, id, fineAmount)
}

// RaiseExpire is a paid mutator transaction binding the contract method 0xfc6753da.
//
// Solidity: function raiseExpire(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) RaiseExpire(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "raiseExpire", id)
}

// RaiseExpire is a paid mutator transaction binding the contract method 0xfc6753da.
//
// Solidity: function raiseExpire(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) RaiseExpire(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.RaiseExpire(&_LetsFilControler.TransactOpts, id)
}

// RaiseExpire is a paid mutator transaction binding the contract method 0xfc6753da.
//
// Solidity: function raiseExpire(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) RaiseExpire(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.RaiseExpire(&_LetsFilControler.TransactOpts, id)
}

// RaiserWithdraw is a paid mutator transaction binding the contract method 0x4e3f2e48.
//
// Solidity: function raiserWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) RaiserWithdraw(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "raiserWithdraw", id)
}

// RaiserWithdraw is a paid mutator transaction binding the contract method 0x4e3f2e48.
//
// Solidity: function raiserWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) RaiserWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.RaiserWithdraw(&_LetsFilControler.TransactOpts, id)
}

// RaiserWithdraw is a paid mutator transaction binding the contract method 0x4e3f2e48.
//
// Solidity: function raiserWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) RaiserWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.RaiserWithdraw(&_LetsFilControler.TransactOpts, id)
}

// SetMinerBackOwner is a paid mutator transaction binding the contract method 0x37bd199a.
//
// Solidity: function setMinerBackOwner(bytes minerOwner) returns()
func (_LetsFilControler *LetsFilControlerTransactor) SetMinerBackOwner(opts *bind.TransactOpts, minerOwner []byte) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "setMinerBackOwner", minerOwner)
}

// SetMinerBackOwner is a paid mutator transaction binding the contract method 0x37bd199a.
//
// Solidity: function setMinerBackOwner(bytes minerOwner) returns()
func (_LetsFilControler *LetsFilControlerSession) SetMinerBackOwner(minerOwner []byte) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetMinerBackOwner(&_LetsFilControler.TransactOpts, minerOwner)
}

// SetMinerBackOwner is a paid mutator transaction binding the contract method 0x37bd199a.
//
// Solidity: function setMinerBackOwner(bytes minerOwner) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) SetMinerBackOwner(minerOwner []byte) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetMinerBackOwner(&_LetsFilControler.TransactOpts, minerOwner)
}

// SetSectorPackage is a paid mutator transaction binding the contract method 0x9f0c3e2a.
//
// Solidity: function setSectorPackage(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) SetSectorPackage(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "setSectorPackage", id)
}

// SetSectorPackage is a paid mutator transaction binding the contract method 0x9f0c3e2a.
//
// Solidity: function setSectorPackage(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) SetSectorPackage(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetSectorPackage(&_LetsFilControler.TransactOpts, id)
}

// SetSectorPackage is a paid mutator transaction binding the contract method 0x9f0c3e2a.
//
// Solidity: function setSectorPackage(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) SetSectorPackage(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetSectorPackage(&_LetsFilControler.TransactOpts, id)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 id, uint256 _raiseStartTime, uint256 _startSealTime, uint256 _sealEndTime) returns()
func (_LetsFilControler *LetsFilControlerTransactor) SetTime(opts *bind.TransactOpts, id *big.Int, _raiseStartTime *big.Int, _startSealTime *big.Int, _sealEndTime *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "setTime", id, _raiseStartTime, _startSealTime, _sealEndTime)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 id, uint256 _raiseStartTime, uint256 _startSealTime, uint256 _sealEndTime) returns()
func (_LetsFilControler *LetsFilControlerSession) SetTime(id *big.Int, _raiseStartTime *big.Int, _startSealTime *big.Int, _sealEndTime *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetTime(&_LetsFilControler.TransactOpts, id, _raiseStartTime, _startSealTime, _sealEndTime)
}

// SetTime is a paid mutator transaction binding the contract method 0xf938c417.
//
// Solidity: function setTime(uint256 id, uint256 _raiseStartTime, uint256 _startSealTime, uint256 _sealEndTime) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) SetTime(id *big.Int, _raiseStartTime *big.Int, _startSealTime *big.Int, _sealEndTime *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SetTime(&_LetsFilControler.TransactOpts, id, _raiseStartTime, _startSealTime, _sealEndTime)
}

// SpSignWithMiner is a paid mutator transaction binding the contract method 0xf01c6ee3.
//
// Solidity: function spSignWithMiner() returns()
func (_LetsFilControler *LetsFilControlerTransactor) SpSignWithMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "spSignWithMiner")
}

// SpSignWithMiner is a paid mutator transaction binding the contract method 0xf01c6ee3.
//
// Solidity: function spSignWithMiner() returns()
func (_LetsFilControler *LetsFilControlerSession) SpSignWithMiner() (*types.Transaction, error) {
	return _LetsFilControler.Contract.SpSignWithMiner(&_LetsFilControler.TransactOpts)
}

// SpSignWithMiner is a paid mutator transaction binding the contract method 0xf01c6ee3.
//
// Solidity: function spSignWithMiner() returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) SpSignWithMiner() (*types.Transaction, error) {
	return _LetsFilControler.Contract.SpSignWithMiner(&_LetsFilControler.TransactOpts)
}

// SpWithdraw is a paid mutator transaction binding the contract method 0xa902eb59.
//
// Solidity: function spWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) SpWithdraw(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "spWithdraw", id)
}

// SpWithdraw is a paid mutator transaction binding the contract method 0xa902eb59.
//
// Solidity: function spWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) SpWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SpWithdraw(&_LetsFilControler.TransactOpts, id)
}

// SpWithdraw is a paid mutator transaction binding the contract method 0xa902eb59.
//
// Solidity: function spWithdraw(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) SpWithdraw(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.SpWithdraw(&_LetsFilControler.TransactOpts, id)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) Staking(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "staking", id)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerSession) Staking(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Staking(&_LetsFilControler.TransactOpts, id)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 id) payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) Staking(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Staking(&_LetsFilControler.TransactOpts, id)
}

// StartPreSeal is a paid mutator transaction binding the contract method 0x65d4d78a.
//
// Solidity: function startPreSeal(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) StartPreSeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "startPreSeal", id)
}

// StartPreSeal is a paid mutator transaction binding the contract method 0x65d4d78a.
//
// Solidity: function startPreSeal(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) StartPreSeal(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.StartPreSeal(&_LetsFilControler.TransactOpts, id)
}

// StartPreSeal is a paid mutator transaction binding the contract method 0x65d4d78a.
//
// Solidity: function startPreSeal(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) StartPreSeal(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.StartPreSeal(&_LetsFilControler.TransactOpts, id)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0xb9c01d0a.
//
// Solidity: function startRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) StartRaisePlan(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "startRaisePlan", id)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0xb9c01d0a.
//
// Solidity: function startRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) StartRaisePlan(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.StartRaisePlan(&_LetsFilControler.TransactOpts, id)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0xb9c01d0a.
//
// Solidity: function startRaisePlan(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) StartRaisePlan(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.StartRaisePlan(&_LetsFilControler.TransactOpts, id)
}

// UnStaking is a paid mutator transaction binding the contract method 0x86f43a41.
//
// Solidity: function unStaking(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) UnStaking(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "unStaking", id)
}

// UnStaking is a paid mutator transaction binding the contract method 0x86f43a41.
//
// Solidity: function unStaking(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) UnStaking(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.UnStaking(&_LetsFilControler.TransactOpts, id)
}

// UnStaking is a paid mutator transaction binding the contract method 0x86f43a41.
//
// Solidity: function unStaking(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) UnStaking(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.UnStaking(&_LetsFilControler.TransactOpts, id)
}

// WithdrawOpsSecurityFund is a paid mutator transaction binding the contract method 0x1a1c6ca8.
//
// Solidity: function withdrawOpsSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) WithdrawOpsSecurityFund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "withdrawOpsSecurityFund", id)
}

// WithdrawOpsSecurityFund is a paid mutator transaction binding the contract method 0x1a1c6ca8.
//
// Solidity: function withdrawOpsSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) WithdrawOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.WithdrawOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// WithdrawOpsSecurityFund is a paid mutator transaction binding the contract method 0x1a1c6ca8.
//
// Solidity: function withdrawOpsSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) WithdrawOpsSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.WithdrawOpsSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// WithdrawSecurityFund is a paid mutator transaction binding the contract method 0x7f78482b.
//
// Solidity: function withdrawSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactor) WithdrawSecurityFund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.contract.Transact(opts, "withdrawSecurityFund", id)
}

// WithdrawSecurityFund is a paid mutator transaction binding the contract method 0x7f78482b.
//
// Solidity: function withdrawSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerSession) WithdrawSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.WithdrawSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// WithdrawSecurityFund is a paid mutator transaction binding the contract method 0x7f78482b.
//
// Solidity: function withdrawSecurityFund(uint256 id) returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) WithdrawSecurityFund(id *big.Int) (*types.Transaction, error) {
	return _LetsFilControler.Contract.WithdrawSecurityFund(&_LetsFilControler.TransactOpts, id)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LetsFilControler.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LetsFilControler *LetsFilControlerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Fallback(&_LetsFilControler.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LetsFilControler.Contract.Fallback(&_LetsFilControler.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LetsFilControler *LetsFilControlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilControler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LetsFilControler *LetsFilControlerSession) Receive() (*types.Transaction, error) {
	return _LetsFilControler.Contract.Receive(&_LetsFilControler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LetsFilControler *LetsFilControlerTransactorSession) Receive() (*types.Transaction, error) {
	return _LetsFilControler.Contract.Receive(&_LetsFilControler.TransactOpts)
}

// LetsFilControlerAddOpsSecurityFundIterator is returned from FilterAddOpsSecurityFund and is used to iterate over the raw logs and unpacked data for AddOpsSecurityFund events raised by the LetsFilControler contract.
type LetsFilControlerAddOpsSecurityFundIterator struct {
	Event *LetsFilControlerAddOpsSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerAddOpsSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerAddOpsSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerAddOpsSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerAddOpsSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerAddOpsSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerAddOpsSecurityFund represents a AddOpsSecurityFund event raised by the LetsFilControler contract.
type LetsFilControlerAddOpsSecurityFund struct {
	Id     *big.Int
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddOpsSecurityFund is a free log retrieval operation binding the contract event 0xc6100d8f0803843a8dd4b0e0ad2648158e67511ecbebf4733883cedcf0bf5b22.
//
// Solidity: event AddOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterAddOpsSecurityFund(opts *bind.FilterOpts, id []*big.Int, sender []common.Address) (*LetsFilControlerAddOpsSecurityFundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "AddOpsSecurityFund", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerAddOpsSecurityFundIterator{contract: _LetsFilControler.contract, event: "AddOpsSecurityFund", logs: logs, sub: sub}, nil
}

// WatchAddOpsSecurityFund is a free log subscription operation binding the contract event 0xc6100d8f0803843a8dd4b0e0ad2648158e67511ecbebf4733883cedcf0bf5b22.
//
// Solidity: event AddOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchAddOpsSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilControlerAddOpsSecurityFund, id []*big.Int, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "AddOpsSecurityFund", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerAddOpsSecurityFund)
				if err := _LetsFilControler.contract.UnpackLog(event, "AddOpsSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddOpsSecurityFund is a log parse operation binding the contract event 0xc6100d8f0803843a8dd4b0e0ad2648158e67511ecbebf4733883cedcf0bf5b22.
//
// Solidity: event AddOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseAddOpsSecurityFund(log types.Log) (*LetsFilControlerAddOpsSecurityFund, error) {
	event := new(LetsFilControlerAddOpsSecurityFund)
	if err := _LetsFilControler.contract.UnpackLog(event, "AddOpsSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerClosePlanToSealIterator is returned from FilterClosePlanToSeal and is used to iterate over the raw logs and unpacked data for ClosePlanToSeal events raised by the LetsFilControler contract.
type LetsFilControlerClosePlanToSealIterator struct {
	Event *LetsFilControlerClosePlanToSeal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerClosePlanToSealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerClosePlanToSeal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerClosePlanToSeal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerClosePlanToSealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerClosePlanToSealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerClosePlanToSeal represents a ClosePlanToSeal event raised by the LetsFilControler contract.
type LetsFilControlerClosePlanToSeal struct {
	Id           *big.Int
	SponsorAddr  common.Address
	CloseTime    *big.Int
	ToSealAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClosePlanToSeal is a free log retrieval operation binding the contract event 0x6d1f8436794059cd9e71fbd58e29cd4203980ca8539a543a2891cf6018abd290.
//
// Solidity: event ClosePlanToSeal(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime, uint256 toSealAmount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterClosePlanToSeal(opts *bind.FilterOpts, id []*big.Int, sponsorAddr []common.Address) (*LetsFilControlerClosePlanToSealIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "ClosePlanToSeal", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerClosePlanToSealIterator{contract: _LetsFilControler.contract, event: "ClosePlanToSeal", logs: logs, sub: sub}, nil
}

// WatchClosePlanToSeal is a free log subscription operation binding the contract event 0x6d1f8436794059cd9e71fbd58e29cd4203980ca8539a543a2891cf6018abd290.
//
// Solidity: event ClosePlanToSeal(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime, uint256 toSealAmount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchClosePlanToSeal(opts *bind.WatchOpts, sink chan<- *LetsFilControlerClosePlanToSeal, id []*big.Int, sponsorAddr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "ClosePlanToSeal", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerClosePlanToSeal)
				if err := _LetsFilControler.contract.UnpackLog(event, "ClosePlanToSeal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosePlanToSeal is a log parse operation binding the contract event 0x6d1f8436794059cd9e71fbd58e29cd4203980ca8539a543a2891cf6018abd290.
//
// Solidity: event ClosePlanToSeal(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime, uint256 toSealAmount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseClosePlanToSeal(log types.Log) (*LetsFilControlerClosePlanToSeal, error) {
	event := new(LetsFilControlerClosePlanToSeal)
	if err := _LetsFilControler.contract.UnpackLog(event, "ClosePlanToSeal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerCloseRaisePlanIterator is returned from FilterCloseRaisePlan and is used to iterate over the raw logs and unpacked data for CloseRaisePlan events raised by the LetsFilControler contract.
type LetsFilControlerCloseRaisePlanIterator struct {
	Event *LetsFilControlerCloseRaisePlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerCloseRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerCloseRaisePlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerCloseRaisePlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerCloseRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerCloseRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerCloseRaisePlan represents a CloseRaisePlan event raised by the LetsFilControler contract.
type LetsFilControlerCloseRaisePlan struct {
	Id          *big.Int
	SponsorAddr common.Address
	CloseTime   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCloseRaisePlan is a free log retrieval operation binding the contract event 0xd3cd5416641279beefa998629a2a6c0d13c880ec5c81593a65882352b8b3a33f.
//
// Solidity: event CloseRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime)
func (_LetsFilControler *LetsFilControlerFilterer) FilterCloseRaisePlan(opts *bind.FilterOpts, id []*big.Int, sponsorAddr []common.Address) (*LetsFilControlerCloseRaisePlanIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "CloseRaisePlan", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerCloseRaisePlanIterator{contract: _LetsFilControler.contract, event: "CloseRaisePlan", logs: logs, sub: sub}, nil
}

// WatchCloseRaisePlan is a free log subscription operation binding the contract event 0xd3cd5416641279beefa998629a2a6c0d13c880ec5c81593a65882352b8b3a33f.
//
// Solidity: event CloseRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime)
func (_LetsFilControler *LetsFilControlerFilterer) WatchCloseRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilControlerCloseRaisePlan, id []*big.Int, sponsorAddr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "CloseRaisePlan", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerCloseRaisePlan)
				if err := _LetsFilControler.contract.UnpackLog(event, "CloseRaisePlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCloseRaisePlan is a log parse operation binding the contract event 0xd3cd5416641279beefa998629a2a6c0d13c880ec5c81593a65882352b8b3a33f.
//
// Solidity: event CloseRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 closeTime)
func (_LetsFilControler *LetsFilControlerFilterer) ParseCloseRaisePlan(log types.Log) (*LetsFilControlerCloseRaisePlan, error) {
	event := new(LetsFilControlerCloseRaisePlan)
	if err := _LetsFilControler.contract.UnpackLog(event, "CloseRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerDepositOpsSecurityFundIterator is returned from FilterDepositOpsSecurityFund and is used to iterate over the raw logs and unpacked data for DepositOpsSecurityFund events raised by the LetsFilControler contract.
type LetsFilControlerDepositOpsSecurityFundIterator struct {
	Event *LetsFilControlerDepositOpsSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerDepositOpsSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerDepositOpsSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerDepositOpsSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerDepositOpsSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerDepositOpsSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerDepositOpsSecurityFund represents a DepositOpsSecurityFund event raised by the LetsFilControler contract.
type LetsFilControlerDepositOpsSecurityFund struct {
	Id     *big.Int
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositOpsSecurityFund is a free log retrieval operation binding the contract event 0xe228837561112d10e769edd0f6a44c3b0d6c66ecdc1b152d809e01ac9777803f.
//
// Solidity: event DepositOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterDepositOpsSecurityFund(opts *bind.FilterOpts, id []*big.Int, sender []common.Address) (*LetsFilControlerDepositOpsSecurityFundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "DepositOpsSecurityFund", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerDepositOpsSecurityFundIterator{contract: _LetsFilControler.contract, event: "DepositOpsSecurityFund", logs: logs, sub: sub}, nil
}

// WatchDepositOpsSecurityFund is a free log subscription operation binding the contract event 0xe228837561112d10e769edd0f6a44c3b0d6c66ecdc1b152d809e01ac9777803f.
//
// Solidity: event DepositOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchDepositOpsSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilControlerDepositOpsSecurityFund, id []*big.Int, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "DepositOpsSecurityFund", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerDepositOpsSecurityFund)
				if err := _LetsFilControler.contract.UnpackLog(event, "DepositOpsSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositOpsSecurityFund is a log parse operation binding the contract event 0xe228837561112d10e769edd0f6a44c3b0d6c66ecdc1b152d809e01ac9777803f.
//
// Solidity: event DepositOpsSecurityFund(uint256 indexed id, address indexed sender, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseDepositOpsSecurityFund(log types.Log) (*LetsFilControlerDepositOpsSecurityFund, error) {
	event := new(LetsFilControlerDepositOpsSecurityFund)
	if err := _LetsFilControler.contract.UnpackLog(event, "DepositOpsSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerDepositSecurityFundIterator is returned from FilterDepositSecurityFund and is used to iterate over the raw logs and unpacked data for DepositSecurityFund events raised by the LetsFilControler contract.
type LetsFilControlerDepositSecurityFundIterator struct {
	Event *LetsFilControlerDepositSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerDepositSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerDepositSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerDepositSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerDepositSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerDepositSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerDepositSecurityFund represents a DepositSecurityFund event raised by the LetsFilControler contract.
type LetsFilControlerDepositSecurityFund struct {
	Id          *big.Int
	SponsorAddr common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositSecurityFund is a free log retrieval operation binding the contract event 0x3730cfc08a7873825f9f06f2a0dca1120f60a5b70751f8ff7f843f080edaf556.
//
// Solidity: event DepositSecurityFund(uint256 indexed id, address indexed sponsorAddr, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterDepositSecurityFund(opts *bind.FilterOpts, id []*big.Int, sponsorAddr []common.Address) (*LetsFilControlerDepositSecurityFundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "DepositSecurityFund", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerDepositSecurityFundIterator{contract: _LetsFilControler.contract, event: "DepositSecurityFund", logs: logs, sub: sub}, nil
}

// WatchDepositSecurityFund is a free log subscription operation binding the contract event 0x3730cfc08a7873825f9f06f2a0dca1120f60a5b70751f8ff7f843f080edaf556.
//
// Solidity: event DepositSecurityFund(uint256 indexed id, address indexed sponsorAddr, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchDepositSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilControlerDepositSecurityFund, id []*big.Int, sponsorAddr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "DepositSecurityFund", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerDepositSecurityFund)
				if err := _LetsFilControler.contract.UnpackLog(event, "DepositSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositSecurityFund is a log parse operation binding the contract event 0x3730cfc08a7873825f9f06f2a0dca1120f60a5b70751f8ff7f843f080edaf556.
//
// Solidity: event DepositSecurityFund(uint256 indexed id, address indexed sponsorAddr, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseDepositSecurityFund(log types.Log) (*LetsFilControlerDepositSecurityFund, error) {
	event := new(LetsFilControlerDepositSecurityFund)
	if err := _LetsFilControler.contract.UnpackLog(event, "DepositSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerDestroyNodeIterator is returned from FilterDestroyNode and is used to iterate over the raw logs and unpacked data for DestroyNode events raised by the LetsFilControler contract.
type LetsFilControlerDestroyNodeIterator struct {
	Event *LetsFilControlerDestroyNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerDestroyNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerDestroyNode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerDestroyNode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerDestroyNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerDestroyNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerDestroyNode represents a DestroyNode event raised by the LetsFilControler contract.
type LetsFilControlerDestroyNode struct {
	RaiseID *big.Int
	State   uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDestroyNode is a free log retrieval operation binding the contract event 0x654ca7c7878998229f7d5ff70c3db3f23592db81109d28d8459328f5a6cc0ed7.
//
// Solidity: event DestroyNode(uint256 indexed raiseID, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) FilterDestroyNode(opts *bind.FilterOpts, raiseID []*big.Int) (*LetsFilControlerDestroyNodeIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "DestroyNode", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerDestroyNodeIterator{contract: _LetsFilControler.contract, event: "DestroyNode", logs: logs, sub: sub}, nil
}

// WatchDestroyNode is a free log subscription operation binding the contract event 0x654ca7c7878998229f7d5ff70c3db3f23592db81109d28d8459328f5a6cc0ed7.
//
// Solidity: event DestroyNode(uint256 indexed raiseID, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) WatchDestroyNode(opts *bind.WatchOpts, sink chan<- *LetsFilControlerDestroyNode, raiseID []*big.Int) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "DestroyNode", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerDestroyNode)
				if err := _LetsFilControler.contract.UnpackLog(event, "DestroyNode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDestroyNode is a log parse operation binding the contract event 0x654ca7c7878998229f7d5ff70c3db3f23592db81109d28d8459328f5a6cc0ed7.
//
// Solidity: event DestroyNode(uint256 indexed raiseID, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) ParseDestroyNode(log types.Log) (*LetsFilControlerDestroyNode, error) {
	event := new(LetsFilControlerDestroyNode)
	if err := _LetsFilControler.contract.UnpackLog(event, "DestroyNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerInvestorWithdrawIterator is returned from FilterInvestorWithdraw and is used to iterate over the raw logs and unpacked data for InvestorWithdraw events raised by the LetsFilControler contract.
type LetsFilControlerInvestorWithdrawIterator struct {
	Event *LetsFilControlerInvestorWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerInvestorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerInvestorWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerInvestorWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerInvestorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerInvestorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerInvestorWithdraw represents a InvestorWithdraw event raised by the LetsFilControler contract.
type LetsFilControlerInvestorWithdraw struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInvestorWithdraw is a free log retrieval operation binding the contract event 0xac6f86a4fdc2a496e25800fd9ba0670ca68a475f62e0b0586708615ec2787073.
//
// Solidity: event InvestorWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterInvestorWithdraw(opts *bind.FilterOpts, raiseID []*big.Int, from []common.Address, to []common.Address) (*LetsFilControlerInvestorWithdrawIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "InvestorWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerInvestorWithdrawIterator{contract: _LetsFilControler.contract, event: "InvestorWithdraw", logs: logs, sub: sub}, nil
}

// WatchInvestorWithdraw is a free log subscription operation binding the contract event 0xac6f86a4fdc2a496e25800fd9ba0670ca68a475f62e0b0586708615ec2787073.
//
// Solidity: event InvestorWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchInvestorWithdraw(opts *bind.WatchOpts, sink chan<- *LetsFilControlerInvestorWithdraw, raiseID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "InvestorWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerInvestorWithdraw)
				if err := _LetsFilControler.contract.UnpackLog(event, "InvestorWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvestorWithdraw is a log parse operation binding the contract event 0xac6f86a4fdc2a496e25800fd9ba0670ca68a475f62e0b0586708615ec2787073.
//
// Solidity: event InvestorWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseInvestorWithdraw(log types.Log) (*LetsFilControlerInvestorWithdraw, error) {
	event := new(LetsFilControlerInvestorWithdraw)
	if err := _LetsFilControler.contract.UnpackLog(event, "InvestorWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerNodeStateChangeIterator is returned from FilterNodeStateChange and is used to iterate over the raw logs and unpacked data for NodeStateChange events raised by the LetsFilControler contract.
type LetsFilControlerNodeStateChangeIterator struct {
	Event *LetsFilControlerNodeStateChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerNodeStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerNodeStateChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerNodeStateChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerNodeStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerNodeStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerNodeStateChange represents a NodeStateChange event raised by the LetsFilControler contract.
type LetsFilControlerNodeStateChange struct {
	Id    *big.Int
	State uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNodeStateChange is a free log retrieval operation binding the contract event 0x796499903ac9a6b030f30fdc40bc64bf616accc42309a428cf70a0b4c8ca313e.
//
// Solidity: event NodeStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) FilterNodeStateChange(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerNodeStateChangeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "NodeStateChange", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerNodeStateChangeIterator{contract: _LetsFilControler.contract, event: "NodeStateChange", logs: logs, sub: sub}, nil
}

// WatchNodeStateChange is a free log subscription operation binding the contract event 0x796499903ac9a6b030f30fdc40bc64bf616accc42309a428cf70a0b4c8ca313e.
//
// Solidity: event NodeStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) WatchNodeStateChange(opts *bind.WatchOpts, sink chan<- *LetsFilControlerNodeStateChange, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "NodeStateChange", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerNodeStateChange)
				if err := _LetsFilControler.contract.UnpackLog(event, "NodeStateChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeStateChange is a log parse operation binding the contract event 0x796499903ac9a6b030f30fdc40bc64bf616accc42309a428cf70a0b4c8ca313e.
//
// Solidity: event NodeStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) ParseNodeStateChange(log types.Log) (*LetsFilControlerNodeStateChange, error) {
	event := new(LetsFilControlerNodeStateChange)
	if err := _LetsFilControler.contract.UnpackLog(event, "NodeStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushBlockRewardIterator is returned from FilterPushBlockReward and is used to iterate over the raw logs and unpacked data for PushBlockReward events raised by the LetsFilControler contract.
type LetsFilControlerPushBlockRewardIterator struct {
	Event *LetsFilControlerPushBlockReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushBlockRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushBlockReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushBlockReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushBlockRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushBlockRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushBlockReward represents a PushBlockReward event raised by the LetsFilControler contract.
type LetsFilControlerPushBlockReward struct {
	Id          *big.Int
	Released    *big.Int
	WillRelease *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPushBlockReward is a free log retrieval operation binding the contract event 0xea4f3b6be33453cbf08fec1ba60a0adb24aef2f928c13de4172beb4bab67624e.
//
// Solidity: event PushBlockReward(uint256 indexed id, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushBlockReward(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerPushBlockRewardIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushBlockReward", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushBlockRewardIterator{contract: _LetsFilControler.contract, event: "PushBlockReward", logs: logs, sub: sub}, nil
}

// WatchPushBlockReward is a free log subscription operation binding the contract event 0xea4f3b6be33453cbf08fec1ba60a0adb24aef2f928c13de4172beb4bab67624e.
//
// Solidity: event PushBlockReward(uint256 indexed id, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushBlockReward(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushBlockReward, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushBlockReward", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushBlockReward)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushBlockReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushBlockReward is a log parse operation binding the contract event 0xea4f3b6be33453cbf08fec1ba60a0adb24aef2f928c13de4172beb4bab67624e.
//
// Solidity: event PushBlockReward(uint256 indexed id, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushBlockReward(log types.Log) (*LetsFilControlerPushBlockReward, error) {
	event := new(LetsFilControlerPushBlockReward)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushBlockReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushFinalProgressIterator is returned from FilterPushFinalProgress and is used to iterate over the raw logs and unpacked data for PushFinalProgress events raised by the LetsFilControler contract.
type LetsFilControlerPushFinalProgressIterator struct {
	Event *LetsFilControlerPushFinalProgress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushFinalProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushFinalProgress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushFinalProgress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushFinalProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushFinalProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushFinalProgress represents a PushFinalProgress event raised by the LetsFilControler contract.
type LetsFilControlerPushFinalProgress struct {
	Id     *big.Int
	Amount *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPushFinalProgress is a free log retrieval operation binding the contract event 0x1fb64102908d34f06c6e3fed072f0e43d01bc849e62f0de2ced926fcaa42978a.
//
// Solidity: event PushFinalProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushFinalProgress(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerPushFinalProgressIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushFinalProgress", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushFinalProgressIterator{contract: _LetsFilControler.contract, event: "PushFinalProgress", logs: logs, sub: sub}, nil
}

// WatchPushFinalProgress is a free log subscription operation binding the contract event 0x1fb64102908d34f06c6e3fed072f0e43d01bc849e62f0de2ced926fcaa42978a.
//
// Solidity: event PushFinalProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushFinalProgress(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushFinalProgress, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushFinalProgress", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushFinalProgress)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushFinalProgress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushFinalProgress is a log parse operation binding the contract event 0x1fb64102908d34f06c6e3fed072f0e43d01bc849e62f0de2ced926fcaa42978a.
//
// Solidity: event PushFinalProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushFinalProgress(log types.Log) (*LetsFilControlerPushFinalProgress, error) {
	event := new(LetsFilControlerPushFinalProgress)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushFinalProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushOldAssetPackValueIterator is returned from FilterPushOldAssetPackValue and is used to iterate over the raw logs and unpacked data for PushOldAssetPackValue events raised by the LetsFilControler contract.
type LetsFilControlerPushOldAssetPackValueIterator struct {
	Event *LetsFilControlerPushOldAssetPackValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushOldAssetPackValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushOldAssetPackValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushOldAssetPackValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushOldAssetPackValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushOldAssetPackValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushOldAssetPackValue represents a PushOldAssetPackValue event raised by the LetsFilControler contract.
type LetsFilControlerPushOldAssetPackValue struct {
	Id          *big.Int
	Caller      common.Address
	TotalPledge *big.Int
	Released    *big.Int
	WillRelease *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPushOldAssetPackValue is a free log retrieval operation binding the contract event 0x185c2ab12778701d0fd99b2a631d55b26c77aa33e4a910a7d8a3bcfd1296e557.
//
// Solidity: event PushOldAssetPackValue(uint256 id, address caller, uint256 totalPledge, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushOldAssetPackValue(opts *bind.FilterOpts) (*LetsFilControlerPushOldAssetPackValueIterator, error) {

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushOldAssetPackValue")
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushOldAssetPackValueIterator{contract: _LetsFilControler.contract, event: "PushOldAssetPackValue", logs: logs, sub: sub}, nil
}

// WatchPushOldAssetPackValue is a free log subscription operation binding the contract event 0x185c2ab12778701d0fd99b2a631d55b26c77aa33e4a910a7d8a3bcfd1296e557.
//
// Solidity: event PushOldAssetPackValue(uint256 id, address caller, uint256 totalPledge, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushOldAssetPackValue(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushOldAssetPackValue) (event.Subscription, error) {

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushOldAssetPackValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushOldAssetPackValue)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushOldAssetPackValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushOldAssetPackValue is a log parse operation binding the contract event 0x185c2ab12778701d0fd99b2a631d55b26c77aa33e4a910a7d8a3bcfd1296e557.
//
// Solidity: event PushOldAssetPackValue(uint256 id, address caller, uint256 totalPledge, uint256 released, uint256 willRelease)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushOldAssetPackValue(log types.Log) (*LetsFilControlerPushOldAssetPackValue, error) {
	event := new(LetsFilControlerPushOldAssetPackValue)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushOldAssetPackValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushPledgeReleasedIterator is returned from FilterPushPledgeReleased and is used to iterate over the raw logs and unpacked data for PushPledgeReleased events raised by the LetsFilControler contract.
type LetsFilControlerPushPledgeReleasedIterator struct {
	Event *LetsFilControlerPushPledgeReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushPledgeReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushPledgeReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushPledgeReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushPledgeReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushPledgeReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushPledgeReleased represents a PushPledgeReleased event raised by the LetsFilControler contract.
type LetsFilControlerPushPledgeReleased struct {
	Id       *big.Int
	Released *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPushPledgeReleased is a free log retrieval operation binding the contract event 0x52efdb5d1649c2af90aedbda9584ea6a1d813173da66d7d3b8f141d9f7cbe8d8.
//
// Solidity: event PushPledgeReleased(uint256 indexed id, uint256 released)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushPledgeReleased(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerPushPledgeReleasedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushPledgeReleased", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushPledgeReleasedIterator{contract: _LetsFilControler.contract, event: "PushPledgeReleased", logs: logs, sub: sub}, nil
}

// WatchPushPledgeReleased is a free log subscription operation binding the contract event 0x52efdb5d1649c2af90aedbda9584ea6a1d813173da66d7d3b8f141d9f7cbe8d8.
//
// Solidity: event PushPledgeReleased(uint256 indexed id, uint256 released)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushPledgeReleased(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushPledgeReleased, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushPledgeReleased", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushPledgeReleased)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushPledgeReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushPledgeReleased is a log parse operation binding the contract event 0x52efdb5d1649c2af90aedbda9584ea6a1d813173da66d7d3b8f141d9f7cbe8d8.
//
// Solidity: event PushPledgeReleased(uint256 indexed id, uint256 released)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushPledgeReleased(log types.Log) (*LetsFilControlerPushPledgeReleased, error) {
	event := new(LetsFilControlerPushPledgeReleased)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushPledgeReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushSealProgressIterator is returned from FilterPushSealProgress and is used to iterate over the raw logs and unpacked data for PushSealProgress events raised by the LetsFilControler contract.
type LetsFilControlerPushSealProgressIterator struct {
	Event *LetsFilControlerPushSealProgress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushSealProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushSealProgress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushSealProgress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushSealProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushSealProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushSealProgress represents a PushSealProgress event raised by the LetsFilControler contract.
type LetsFilControlerPushSealProgress struct {
	Id     *big.Int
	Amount *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPushSealProgress is a free log retrieval operation binding the contract event 0x9308d7598133790a5daa9c5c29666eaec72b4f1a827fc2fb1309aba23b7fc74a.
//
// Solidity: event PushSealProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushSealProgress(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerPushSealProgressIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushSealProgress", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushSealProgressIterator{contract: _LetsFilControler.contract, event: "PushSealProgress", logs: logs, sub: sub}, nil
}

// WatchPushSealProgress is a free log subscription operation binding the contract event 0x9308d7598133790a5daa9c5c29666eaec72b4f1a827fc2fb1309aba23b7fc74a.
//
// Solidity: event PushSealProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushSealProgress(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushSealProgress, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushSealProgress", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushSealProgress)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushSealProgress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushSealProgress is a log parse operation binding the contract event 0x9308d7598133790a5daa9c5c29666eaec72b4f1a827fc2fb1309aba23b7fc74a.
//
// Solidity: event PushSealProgress(uint256 indexed id, uint256 amount, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushSealProgress(log types.Log) (*LetsFilControlerPushSealProgress, error) {
	event := new(LetsFilControlerPushSealProgress)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushSealProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerPushSpFineIterator is returned from FilterPushSpFine and is used to iterate over the raw logs and unpacked data for PushSpFine events raised by the LetsFilControler contract.
type LetsFilControlerPushSpFineIterator struct {
	Event *LetsFilControlerPushSpFine // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerPushSpFineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerPushSpFine)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerPushSpFine)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerPushSpFineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerPushSpFineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerPushSpFine represents a PushSpFine event raised by the LetsFilControler contract.
type LetsFilControlerPushSpFine struct {
	Id         *big.Int
	FineAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPushSpFine is a free log retrieval operation binding the contract event 0x110e891d224861a4165664c379982282403a7a8e60d5fd499218462fb9b21828.
//
// Solidity: event PushSpFine(uint256 indexed id, uint256 fineAmount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterPushSpFine(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerPushSpFineIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "PushSpFine", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerPushSpFineIterator{contract: _LetsFilControler.contract, event: "PushSpFine", logs: logs, sub: sub}, nil
}

// WatchPushSpFine is a free log subscription operation binding the contract event 0x110e891d224861a4165664c379982282403a7a8e60d5fd499218462fb9b21828.
//
// Solidity: event PushSpFine(uint256 indexed id, uint256 fineAmount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchPushSpFine(opts *bind.WatchOpts, sink chan<- *LetsFilControlerPushSpFine, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "PushSpFine", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerPushSpFine)
				if err := _LetsFilControler.contract.UnpackLog(event, "PushSpFine", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePushSpFine is a log parse operation binding the contract event 0x110e891d224861a4165664c379982282403a7a8e60d5fd499218462fb9b21828.
//
// Solidity: event PushSpFine(uint256 indexed id, uint256 fineAmount)
func (_LetsFilControler *LetsFilControlerFilterer) ParsePushSpFine(log types.Log) (*LetsFilControlerPushSpFine, error) {
	event := new(LetsFilControlerPushSpFine)
	if err := _LetsFilControler.contract.UnpackLog(event, "PushSpFine", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerRaiseFailedIterator is returned from FilterRaiseFailed and is used to iterate over the raw logs and unpacked data for RaiseFailed events raised by the LetsFilControler contract.
type LetsFilControlerRaiseFailedIterator struct {
	Event *LetsFilControlerRaiseFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerRaiseFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerRaiseFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerRaiseFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerRaiseFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerRaiseFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerRaiseFailed represents a RaiseFailed event raised by the LetsFilControler contract.
type LetsFilControlerRaiseFailed struct {
	RaiseID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaiseFailed is a free log retrieval operation binding the contract event 0xc29d2a5a4ec664c8945e3196c7be82e008f63e1cd1465ad90e950eb8a8406fc2.
//
// Solidity: event RaiseFailed(uint256 indexed raiseID)
func (_LetsFilControler *LetsFilControlerFilterer) FilterRaiseFailed(opts *bind.FilterOpts, raiseID []*big.Int) (*LetsFilControlerRaiseFailedIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "RaiseFailed", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerRaiseFailedIterator{contract: _LetsFilControler.contract, event: "RaiseFailed", logs: logs, sub: sub}, nil
}

// WatchRaiseFailed is a free log subscription operation binding the contract event 0xc29d2a5a4ec664c8945e3196c7be82e008f63e1cd1465ad90e950eb8a8406fc2.
//
// Solidity: event RaiseFailed(uint256 indexed raiseID)
func (_LetsFilControler *LetsFilControlerFilterer) WatchRaiseFailed(opts *bind.WatchOpts, sink chan<- *LetsFilControlerRaiseFailed, raiseID []*big.Int) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "RaiseFailed", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerRaiseFailed)
				if err := _LetsFilControler.contract.UnpackLog(event, "RaiseFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRaiseFailed is a log parse operation binding the contract event 0xc29d2a5a4ec664c8945e3196c7be82e008f63e1cd1465ad90e950eb8a8406fc2.
//
// Solidity: event RaiseFailed(uint256 indexed raiseID)
func (_LetsFilControler *LetsFilControlerFilterer) ParseRaiseFailed(log types.Log) (*LetsFilControlerRaiseFailed, error) {
	event := new(LetsFilControlerRaiseFailed)
	if err := _LetsFilControler.contract.UnpackLog(event, "RaiseFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerRaiseStateChangeIterator is returned from FilterRaiseStateChange and is used to iterate over the raw logs and unpacked data for RaiseStateChange events raised by the LetsFilControler contract.
type LetsFilControlerRaiseStateChangeIterator struct {
	Event *LetsFilControlerRaiseStateChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerRaiseStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerRaiseStateChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerRaiseStateChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerRaiseStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerRaiseStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerRaiseStateChange represents a RaiseStateChange event raised by the LetsFilControler contract.
type LetsFilControlerRaiseStateChange struct {
	Id    *big.Int
	State uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRaiseStateChange is a free log retrieval operation binding the contract event 0x42941183c8fc2993ca42f0ba242b46b502e22df4b921e3e777f3c96daf2e7917.
//
// Solidity: event RaiseStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) FilterRaiseStateChange(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerRaiseStateChangeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "RaiseStateChange", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerRaiseStateChangeIterator{contract: _LetsFilControler.contract, event: "RaiseStateChange", logs: logs, sub: sub}, nil
}

// WatchRaiseStateChange is a free log subscription operation binding the contract event 0x42941183c8fc2993ca42f0ba242b46b502e22df4b921e3e777f3c96daf2e7917.
//
// Solidity: event RaiseStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) WatchRaiseStateChange(opts *bind.WatchOpts, sink chan<- *LetsFilControlerRaiseStateChange, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "RaiseStateChange", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerRaiseStateChange)
				if err := _LetsFilControler.contract.UnpackLog(event, "RaiseStateChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRaiseStateChange is a log parse operation binding the contract event 0x42941183c8fc2993ca42f0ba242b46b502e22df4b921e3e777f3c96daf2e7917.
//
// Solidity: event RaiseStateChange(uint256 indexed id, uint8 state)
func (_LetsFilControler *LetsFilControlerFilterer) ParseRaiseStateChange(log types.Log) (*LetsFilControlerRaiseStateChange, error) {
	event := new(LetsFilControlerRaiseStateChange)
	if err := _LetsFilControler.contract.UnpackLog(event, "RaiseStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerRaiseSuccessIterator is returned from FilterRaiseSuccess and is used to iterate over the raw logs and unpacked data for RaiseSuccess events raised by the LetsFilControler contract.
type LetsFilControlerRaiseSuccessIterator struct {
	Event *LetsFilControlerRaiseSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerRaiseSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerRaiseSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerRaiseSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerRaiseSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerRaiseSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerRaiseSuccess represents a RaiseSuccess event raised by the LetsFilControler contract.
type LetsFilControlerRaiseSuccess struct {
	RaiseID      *big.Int
	PledgeAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRaiseSuccess is a free log retrieval operation binding the contract event 0x551a3c904d95a54b7cd74535a05b062c8b656bd6b35412069a290e690750c16a.
//
// Solidity: event RaiseSuccess(uint256 indexed raiseID, uint256 pledgeAmount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterRaiseSuccess(opts *bind.FilterOpts, raiseID []*big.Int) (*LetsFilControlerRaiseSuccessIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "RaiseSuccess", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerRaiseSuccessIterator{contract: _LetsFilControler.contract, event: "RaiseSuccess", logs: logs, sub: sub}, nil
}

// WatchRaiseSuccess is a free log subscription operation binding the contract event 0x551a3c904d95a54b7cd74535a05b062c8b656bd6b35412069a290e690750c16a.
//
// Solidity: event RaiseSuccess(uint256 indexed raiseID, uint256 pledgeAmount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchRaiseSuccess(opts *bind.WatchOpts, sink chan<- *LetsFilControlerRaiseSuccess, raiseID []*big.Int) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "RaiseSuccess", raiseIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerRaiseSuccess)
				if err := _LetsFilControler.contract.UnpackLog(event, "RaiseSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRaiseSuccess is a log parse operation binding the contract event 0x551a3c904d95a54b7cd74535a05b062c8b656bd6b35412069a290e690750c16a.
//
// Solidity: event RaiseSuccess(uint256 indexed raiseID, uint256 pledgeAmount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseRaiseSuccess(log types.Log) (*LetsFilControlerRaiseSuccess, error) {
	event := new(LetsFilControlerRaiseSuccess)
	if err := _LetsFilControler.contract.UnpackLog(event, "RaiseSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerRaiseWithdrawIterator is returned from FilterRaiseWithdraw and is used to iterate over the raw logs and unpacked data for RaiseWithdraw events raised by the LetsFilControler contract.
type LetsFilControlerRaiseWithdrawIterator struct {
	Event *LetsFilControlerRaiseWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerRaiseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerRaiseWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerRaiseWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerRaiseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerRaiseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerRaiseWithdraw represents a RaiseWithdraw event raised by the LetsFilControler contract.
type LetsFilControlerRaiseWithdraw struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaiseWithdraw is a free log retrieval operation binding the contract event 0x701cddd57745bf6877244212e325888384412128b541d5c9941c79a764b95925.
//
// Solidity: event RaiseWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterRaiseWithdraw(opts *bind.FilterOpts, raiseID []*big.Int, from []common.Address, to []common.Address) (*LetsFilControlerRaiseWithdrawIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "RaiseWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerRaiseWithdrawIterator{contract: _LetsFilControler.contract, event: "RaiseWithdraw", logs: logs, sub: sub}, nil
}

// WatchRaiseWithdraw is a free log subscription operation binding the contract event 0x701cddd57745bf6877244212e325888384412128b541d5c9941c79a764b95925.
//
// Solidity: event RaiseWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchRaiseWithdraw(opts *bind.WatchOpts, sink chan<- *LetsFilControlerRaiseWithdraw, raiseID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "RaiseWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerRaiseWithdraw)
				if err := _LetsFilControler.contract.UnpackLog(event, "RaiseWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRaiseWithdraw is a log parse operation binding the contract event 0x701cddd57745bf6877244212e325888384412128b541d5c9941c79a764b95925.
//
// Solidity: event RaiseWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseRaiseWithdraw(log types.Log) (*LetsFilControlerRaiseWithdraw, error) {
	event := new(LetsFilControlerRaiseWithdraw)
	if err := _LetsFilControler.contract.UnpackLog(event, "RaiseWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerSealEndIterator is returned from FilterSealEnd and is used to iterate over the raw logs and unpacked data for SealEnd events raised by the LetsFilControler contract.
type LetsFilControlerSealEndIterator struct {
	Event *LetsFilControlerSealEnd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerSealEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerSealEnd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerSealEnd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerSealEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerSealEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerSealEnd represents a SealEnd event raised by the LetsFilControler contract.
type LetsFilControlerSealEnd struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSealEnd is a free log retrieval operation binding the contract event 0x28377046d80d0ca8f94e6d451539e6889ca029f4480dc2a892ddb85c157251be.
//
// Solidity: event SealEnd(uint256 indexed id, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterSealEnd(opts *bind.FilterOpts, id []*big.Int) (*LetsFilControlerSealEndIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "SealEnd", idRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerSealEndIterator{contract: _LetsFilControler.contract, event: "SealEnd", logs: logs, sub: sub}, nil
}

// WatchSealEnd is a free log subscription operation binding the contract event 0x28377046d80d0ca8f94e6d451539e6889ca029f4480dc2a892ddb85c157251be.
//
// Solidity: event SealEnd(uint256 indexed id, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchSealEnd(opts *bind.WatchOpts, sink chan<- *LetsFilControlerSealEnd, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "SealEnd", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerSealEnd)
				if err := _LetsFilControler.contract.UnpackLog(event, "SealEnd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSealEnd is a log parse operation binding the contract event 0x28377046d80d0ca8f94e6d451539e6889ca029f4480dc2a892ddb85c157251be.
//
// Solidity: event SealEnd(uint256 indexed id, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseSealEnd(log types.Log) (*LetsFilControlerSealEnd, error) {
	event := new(LetsFilControlerSealEnd)
	if err := _LetsFilControler.contract.UnpackLog(event, "SealEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerSendToMinerIterator is returned from FilterSendToMiner and is used to iterate over the raw logs and unpacked data for SendToMiner events raised by the LetsFilControler contract.
type LetsFilControlerSendToMinerIterator struct {
	Event *LetsFilControlerSendToMiner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerSendToMinerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerSendToMiner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerSendToMiner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerSendToMinerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerSendToMinerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerSendToMiner represents a SendToMiner event raised by the LetsFilControler contract.
type LetsFilControlerSendToMiner struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendToMiner is a free log retrieval operation binding the contract event 0x467851990bc7def96ff6170a886e1856d2c1644621e762169a366f555cfaea10.
//
// Solidity: event SendToMiner(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterSendToMiner(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*LetsFilControlerSendToMinerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "SendToMiner", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerSendToMinerIterator{contract: _LetsFilControler.contract, event: "SendToMiner", logs: logs, sub: sub}, nil
}

// WatchSendToMiner is a free log subscription operation binding the contract event 0x467851990bc7def96ff6170a886e1856d2c1644621e762169a366f555cfaea10.
//
// Solidity: event SendToMiner(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchSendToMiner(opts *bind.WatchOpts, sink chan<- *LetsFilControlerSendToMiner, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "SendToMiner", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerSendToMiner)
				if err := _LetsFilControler.contract.UnpackLog(event, "SendToMiner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendToMiner is a log parse operation binding the contract event 0x467851990bc7def96ff6170a886e1856d2c1644621e762169a366f555cfaea10.
//
// Solidity: event SendToMiner(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseSendToMiner(log types.Log) (*LetsFilControlerSendToMiner, error) {
	event := new(LetsFilControlerSendToMiner)
	if err := _LetsFilControler.contract.UnpackLog(event, "SendToMiner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerSpSignWithMinerIterator is returned from FilterSpSignWithMiner and is used to iterate over the raw logs and unpacked data for SpSignWithMiner events raised by the LetsFilControler contract.
type LetsFilControlerSpSignWithMinerIterator struct {
	Event *LetsFilControlerSpSignWithMiner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerSpSignWithMinerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerSpSignWithMiner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerSpSignWithMiner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerSpSignWithMinerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerSpSignWithMinerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerSpSignWithMiner represents a SpSignWithMiner event raised by the LetsFilControler contract.
type LetsFilControlerSpSignWithMiner struct {
	Sp           common.Address
	MinerId      uint64
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSpSignWithMiner is a free log retrieval operation binding the contract event 0xe83153d2fee620b51a41d993afdc1c9740a91d2758b5c6d05680a5ad051dbf7c.
//
// Solidity: event SpSignWithMiner(address indexed sp, uint64 indexed minerId, address contractAddr)
func (_LetsFilControler *LetsFilControlerFilterer) FilterSpSignWithMiner(opts *bind.FilterOpts, sp []common.Address, minerId []uint64) (*LetsFilControlerSpSignWithMinerIterator, error) {

	var spRule []interface{}
	for _, spItem := range sp {
		spRule = append(spRule, spItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "SpSignWithMiner", spRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerSpSignWithMinerIterator{contract: _LetsFilControler.contract, event: "SpSignWithMiner", logs: logs, sub: sub}, nil
}

// WatchSpSignWithMiner is a free log subscription operation binding the contract event 0xe83153d2fee620b51a41d993afdc1c9740a91d2758b5c6d05680a5ad051dbf7c.
//
// Solidity: event SpSignWithMiner(address indexed sp, uint64 indexed minerId, address contractAddr)
func (_LetsFilControler *LetsFilControlerFilterer) WatchSpSignWithMiner(opts *bind.WatchOpts, sink chan<- *LetsFilControlerSpSignWithMiner, sp []common.Address, minerId []uint64) (event.Subscription, error) {

	var spRule []interface{}
	for _, spItem := range sp {
		spRule = append(spRule, spItem)
	}
	var minerIdRule []interface{}
	for _, minerIdItem := range minerId {
		minerIdRule = append(minerIdRule, minerIdItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "SpSignWithMiner", spRule, minerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerSpSignWithMiner)
				if err := _LetsFilControler.contract.UnpackLog(event, "SpSignWithMiner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpSignWithMiner is a log parse operation binding the contract event 0xe83153d2fee620b51a41d993afdc1c9740a91d2758b5c6d05680a5ad051dbf7c.
//
// Solidity: event SpSignWithMiner(address indexed sp, uint64 indexed minerId, address contractAddr)
func (_LetsFilControler *LetsFilControlerFilterer) ParseSpSignWithMiner(log types.Log) (*LetsFilControlerSpSignWithMiner, error) {
	event := new(LetsFilControlerSpSignWithMiner)
	if err := _LetsFilControler.contract.UnpackLog(event, "SpSignWithMiner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerSpWithdrawIterator is returned from FilterSpWithdraw and is used to iterate over the raw logs and unpacked data for SpWithdraw events raised by the LetsFilControler contract.
type LetsFilControlerSpWithdrawIterator struct {
	Event *LetsFilControlerSpWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerSpWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerSpWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerSpWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerSpWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerSpWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerSpWithdraw represents a SpWithdraw event raised by the LetsFilControler contract.
type LetsFilControlerSpWithdraw struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpWithdraw is a free log retrieval operation binding the contract event 0xc4e50c784ae5e0963f27ab8655817f31dbcb2d9617f96572346cf8fa77bba3f7.
//
// Solidity: event SpWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterSpWithdraw(opts *bind.FilterOpts, raiseID []*big.Int, from []common.Address, to []common.Address) (*LetsFilControlerSpWithdrawIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "SpWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerSpWithdrawIterator{contract: _LetsFilControler.contract, event: "SpWithdraw", logs: logs, sub: sub}, nil
}

// WatchSpWithdraw is a free log subscription operation binding the contract event 0xc4e50c784ae5e0963f27ab8655817f31dbcb2d9617f96572346cf8fa77bba3f7.
//
// Solidity: event SpWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchSpWithdraw(opts *bind.WatchOpts, sink chan<- *LetsFilControlerSpWithdraw, raiseID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "SpWithdraw", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerSpWithdraw)
				if err := _LetsFilControler.contract.UnpackLog(event, "SpWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpWithdraw is a log parse operation binding the contract event 0xc4e50c784ae5e0963f27ab8655817f31dbcb2d9617f96572346cf8fa77bba3f7.
//
// Solidity: event SpWithdraw(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseSpWithdraw(log types.Log) (*LetsFilControlerSpWithdraw, error) {
	event := new(LetsFilControlerSpWithdraw)
	if err := _LetsFilControler.contract.UnpackLog(event, "SpWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerStakingIterator is returned from FilterStaking and is used to iterate over the raw logs and unpacked data for Staking events raised by the LetsFilControler contract.
type LetsFilControlerStakingIterator struct {
	Event *LetsFilControlerStaking // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerStakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerStaking)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerStaking)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerStakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerStakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerStaking represents a Staking event raised by the LetsFilControler contract.
type LetsFilControlerStaking struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaking is a free log retrieval operation binding the contract event 0x1cbb7f116c85157771281056083fc1f1eff1bb69b5ad1cb321227901d16bb222.
//
// Solidity: event Staking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterStaking(opts *bind.FilterOpts, raiseID []*big.Int, from []common.Address, to []common.Address) (*LetsFilControlerStakingIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "Staking", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerStakingIterator{contract: _LetsFilControler.contract, event: "Staking", logs: logs, sub: sub}, nil
}

// WatchStaking is a free log subscription operation binding the contract event 0x1cbb7f116c85157771281056083fc1f1eff1bb69b5ad1cb321227901d16bb222.
//
// Solidity: event Staking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchStaking(opts *bind.WatchOpts, sink chan<- *LetsFilControlerStaking, raiseID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "Staking", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerStaking)
				if err := _LetsFilControler.contract.UnpackLog(event, "Staking", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaking is a log parse operation binding the contract event 0x1cbb7f116c85157771281056083fc1f1eff1bb69b5ad1cb321227901d16bb222.
//
// Solidity: event Staking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseStaking(log types.Log) (*LetsFilControlerStaking, error) {
	event := new(LetsFilControlerStaking)
	if err := _LetsFilControler.contract.UnpackLog(event, "Staking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerStartPreSealIterator is returned from FilterStartPreSeal and is used to iterate over the raw logs and unpacked data for StartPreSeal events raised by the LetsFilControler contract.
type LetsFilControlerStartPreSealIterator struct {
	Event *LetsFilControlerStartPreSeal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerStartPreSealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerStartPreSeal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerStartPreSeal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerStartPreSealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerStartPreSealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerStartPreSeal represents a StartPreSeal event raised by the LetsFilControler contract.
type LetsFilControlerStartPreSeal struct {
	Id     *big.Int
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStartPreSeal is a free log retrieval operation binding the contract event 0xa83fe62287fad634ed5262a648ae68da03515381d3c0bc27e5890846f69421e0.
//
// Solidity: event StartPreSeal(uint256 indexed id, address indexed caller)
func (_LetsFilControler *LetsFilControlerFilterer) FilterStartPreSeal(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*LetsFilControlerStartPreSealIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "StartPreSeal", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerStartPreSealIterator{contract: _LetsFilControler.contract, event: "StartPreSeal", logs: logs, sub: sub}, nil
}

// WatchStartPreSeal is a free log subscription operation binding the contract event 0xa83fe62287fad634ed5262a648ae68da03515381d3c0bc27e5890846f69421e0.
//
// Solidity: event StartPreSeal(uint256 indexed id, address indexed caller)
func (_LetsFilControler *LetsFilControlerFilterer) WatchStartPreSeal(opts *bind.WatchOpts, sink chan<- *LetsFilControlerStartPreSeal, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "StartPreSeal", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerStartPreSeal)
				if err := _LetsFilControler.contract.UnpackLog(event, "StartPreSeal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartPreSeal is a log parse operation binding the contract event 0xa83fe62287fad634ed5262a648ae68da03515381d3c0bc27e5890846f69421e0.
//
// Solidity: event StartPreSeal(uint256 indexed id, address indexed caller)
func (_LetsFilControler *LetsFilControlerFilterer) ParseStartPreSeal(log types.Log) (*LetsFilControlerStartPreSeal, error) {
	event := new(LetsFilControlerStartPreSeal)
	if err := _LetsFilControler.contract.UnpackLog(event, "StartPreSeal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerStartRaisePlanIterator is returned from FilterStartRaisePlan and is used to iterate over the raw logs and unpacked data for StartRaisePlan events raised by the LetsFilControler contract.
type LetsFilControlerStartRaisePlanIterator struct {
	Event *LetsFilControlerStartRaisePlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerStartRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerStartRaisePlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerStartRaisePlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerStartRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerStartRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerStartRaisePlan represents a StartRaisePlan event raised by the LetsFilControler contract.
type LetsFilControlerStartRaisePlan struct {
	Id          *big.Int
	SponsorAddr common.Address
	StartTime   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartRaisePlan is a free log retrieval operation binding the contract event 0x8381ffa4f05dd3bf47b5edc02903a64d2878d617f1d1790d47fe0e48517d158b.
//
// Solidity: event StartRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) FilterStartRaisePlan(opts *bind.FilterOpts, id []*big.Int, sponsorAddr []common.Address) (*LetsFilControlerStartRaisePlanIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "StartRaisePlan", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerStartRaisePlanIterator{contract: _LetsFilControler.contract, event: "StartRaisePlan", logs: logs, sub: sub}, nil
}

// WatchStartRaisePlan is a free log subscription operation binding the contract event 0x8381ffa4f05dd3bf47b5edc02903a64d2878d617f1d1790d47fe0e48517d158b.
//
// Solidity: event StartRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) WatchStartRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilControlerStartRaisePlan, id []*big.Int, sponsorAddr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sponsorAddrRule []interface{}
	for _, sponsorAddrItem := range sponsorAddr {
		sponsorAddrRule = append(sponsorAddrRule, sponsorAddrItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "StartRaisePlan", idRule, sponsorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerStartRaisePlan)
				if err := _LetsFilControler.contract.UnpackLog(event, "StartRaisePlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartRaisePlan is a log parse operation binding the contract event 0x8381ffa4f05dd3bf47b5edc02903a64d2878d617f1d1790d47fe0e48517d158b.
//
// Solidity: event StartRaisePlan(uint256 indexed id, address indexed sponsorAddr, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) ParseStartRaisePlan(log types.Log) (*LetsFilControlerStartRaisePlan, error) {
	event := new(LetsFilControlerStartRaisePlan)
	if err := _LetsFilControler.contract.UnpackLog(event, "StartRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerStartSealIterator is returned from FilterStartSeal and is used to iterate over the raw logs and unpacked data for StartSeal events raised by the LetsFilControler contract.
type LetsFilControlerStartSealIterator struct {
	Event *LetsFilControlerStartSeal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerStartSealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerStartSeal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerStartSeal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerStartSealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerStartSealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerStartSeal represents a StartSeal event raised by the LetsFilControler contract.
type LetsFilControlerStartSeal struct {
	Id        *big.Int
	Caller    common.Address
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartSeal is a free log retrieval operation binding the contract event 0x255dee1360907571e6d70611cf6b4ddb031ce734083d5a47ab67790ea08128a3.
//
// Solidity: event StartSeal(uint256 indexed id, address indexed caller, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) FilterStartSeal(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*LetsFilControlerStartSealIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "StartSeal", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerStartSealIterator{contract: _LetsFilControler.contract, event: "StartSeal", logs: logs, sub: sub}, nil
}

// WatchStartSeal is a free log subscription operation binding the contract event 0x255dee1360907571e6d70611cf6b4ddb031ce734083d5a47ab67790ea08128a3.
//
// Solidity: event StartSeal(uint256 indexed id, address indexed caller, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) WatchStartSeal(opts *bind.WatchOpts, sink chan<- *LetsFilControlerStartSeal, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "StartSeal", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerStartSeal)
				if err := _LetsFilControler.contract.UnpackLog(event, "StartSeal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartSeal is a log parse operation binding the contract event 0x255dee1360907571e6d70611cf6b4ddb031ce734083d5a47ab67790ea08128a3.
//
// Solidity: event StartSeal(uint256 indexed id, address indexed caller, uint256 startTime)
func (_LetsFilControler *LetsFilControlerFilterer) ParseStartSeal(log types.Log) (*LetsFilControlerStartSeal, error) {
	event := new(LetsFilControlerStartSeal)
	if err := _LetsFilControler.contract.UnpackLog(event, "StartSeal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerUnstakingIterator is returned from FilterUnstaking and is used to iterate over the raw logs and unpacked data for Unstaking events raised by the LetsFilControler contract.
type LetsFilControlerUnstakingIterator struct {
	Event *LetsFilControlerUnstaking // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerUnstakingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerUnstaking)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerUnstaking)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerUnstakingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerUnstakingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerUnstaking represents a Unstaking event raised by the LetsFilControler contract.
type LetsFilControlerUnstaking struct {
	RaiseID  *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Interest *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstaking is a free log retrieval operation binding the contract event 0x0680b3735999bc5feb8a878f2a981272d04dbc317da5e4f6ae876bfb0989370d.
//
// Solidity: event Unstaking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount, uint256 interest)
func (_LetsFilControler *LetsFilControlerFilterer) FilterUnstaking(opts *bind.FilterOpts, raiseID []*big.Int, from []common.Address, to []common.Address) (*LetsFilControlerUnstakingIterator, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "Unstaking", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerUnstakingIterator{contract: _LetsFilControler.contract, event: "Unstaking", logs: logs, sub: sub}, nil
}

// WatchUnstaking is a free log subscription operation binding the contract event 0x0680b3735999bc5feb8a878f2a981272d04dbc317da5e4f6ae876bfb0989370d.
//
// Solidity: event Unstaking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount, uint256 interest)
func (_LetsFilControler *LetsFilControlerFilterer) WatchUnstaking(opts *bind.WatchOpts, sink chan<- *LetsFilControlerUnstaking, raiseID []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var raiseIDRule []interface{}
	for _, raiseIDItem := range raiseID {
		raiseIDRule = append(raiseIDRule, raiseIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "Unstaking", raiseIDRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerUnstaking)
				if err := _LetsFilControler.contract.UnpackLog(event, "Unstaking", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaking is a log parse operation binding the contract event 0x0680b3735999bc5feb8a878f2a981272d04dbc317da5e4f6ae876bfb0989370d.
//
// Solidity: event Unstaking(uint256 indexed raiseID, address indexed from, address indexed to, uint256 amount, uint256 interest)
func (_LetsFilControler *LetsFilControlerFilterer) ParseUnstaking(log types.Log) (*LetsFilControlerUnstaking, error) {
	event := new(LetsFilControlerUnstaking)
	if err := _LetsFilControler.contract.UnpackLog(event, "Unstaking", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerWithdrawOpsSecurityFundIterator is returned from FilterWithdrawOpsSecurityFund and is used to iterate over the raw logs and unpacked data for WithdrawOpsSecurityFund events raised by the LetsFilControler contract.
type LetsFilControlerWithdrawOpsSecurityFundIterator struct {
	Event *LetsFilControlerWithdrawOpsSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerWithdrawOpsSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerWithdrawOpsSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerWithdrawOpsSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerWithdrawOpsSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerWithdrawOpsSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerWithdrawOpsSecurityFund represents a WithdrawOpsSecurityFund event raised by the LetsFilControler contract.
type LetsFilControlerWithdrawOpsSecurityFund struct {
	Id       *big.Int
	Caller   common.Address
	Amount   *big.Int
	Interest *big.Int
	Reward   *big.Int
	Fine     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawOpsSecurityFund is a free log retrieval operation binding the contract event 0x102111b811acdc976ae9e78f550eb11fb2c20b44158f38ffb2f7f38f80ec4422.
//
// Solidity: event WithdrawOpsSecurityFund(uint256 indexed id, address indexed caller, uint256 amount, uint256 interest, uint256 reward, uint256 fine)
func (_LetsFilControler *LetsFilControlerFilterer) FilterWithdrawOpsSecurityFund(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*LetsFilControlerWithdrawOpsSecurityFundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "WithdrawOpsSecurityFund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerWithdrawOpsSecurityFundIterator{contract: _LetsFilControler.contract, event: "WithdrawOpsSecurityFund", logs: logs, sub: sub}, nil
}

// WatchWithdrawOpsSecurityFund is a free log subscription operation binding the contract event 0x102111b811acdc976ae9e78f550eb11fb2c20b44158f38ffb2f7f38f80ec4422.
//
// Solidity: event WithdrawOpsSecurityFund(uint256 indexed id, address indexed caller, uint256 amount, uint256 interest, uint256 reward, uint256 fine)
func (_LetsFilControler *LetsFilControlerFilterer) WatchWithdrawOpsSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilControlerWithdrawOpsSecurityFund, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "WithdrawOpsSecurityFund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerWithdrawOpsSecurityFund)
				if err := _LetsFilControler.contract.UnpackLog(event, "WithdrawOpsSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawOpsSecurityFund is a log parse operation binding the contract event 0x102111b811acdc976ae9e78f550eb11fb2c20b44158f38ffb2f7f38f80ec4422.
//
// Solidity: event WithdrawOpsSecurityFund(uint256 indexed id, address indexed caller, uint256 amount, uint256 interest, uint256 reward, uint256 fine)
func (_LetsFilControler *LetsFilControlerFilterer) ParseWithdrawOpsSecurityFund(log types.Log) (*LetsFilControlerWithdrawOpsSecurityFund, error) {
	event := new(LetsFilControlerWithdrawOpsSecurityFund)
	if err := _LetsFilControler.contract.UnpackLog(event, "WithdrawOpsSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilControlerWithdrawSecurityFundIterator is returned from FilterWithdrawSecurityFund and is used to iterate over the raw logs and unpacked data for WithdrawSecurityFund events raised by the LetsFilControler contract.
type LetsFilControlerWithdrawSecurityFundIterator struct {
	Event *LetsFilControlerWithdrawSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilControlerWithdrawSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilControlerWithdrawSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilControlerWithdrawSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilControlerWithdrawSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilControlerWithdrawSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilControlerWithdrawSecurityFund represents a WithdrawSecurityFund event raised by the LetsFilControler contract.
type LetsFilControlerWithdrawSecurityFund struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSecurityFund is a free log retrieval operation binding the contract event 0x29c8190b82df1ee51e5380bc142f4613aacb0bd4078470790d35942959c3d0d5.
//
// Solidity: event WithdrawSecurityFund(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) FilterWithdrawSecurityFund(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*LetsFilControlerWithdrawSecurityFundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.FilterLogs(opts, "WithdrawSecurityFund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilControlerWithdrawSecurityFundIterator{contract: _LetsFilControler.contract, event: "WithdrawSecurityFund", logs: logs, sub: sub}, nil
}

// WatchWithdrawSecurityFund is a free log subscription operation binding the contract event 0x29c8190b82df1ee51e5380bc142f4613aacb0bd4078470790d35942959c3d0d5.
//
// Solidity: event WithdrawSecurityFund(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) WatchWithdrawSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilControlerWithdrawSecurityFund, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LetsFilControler.contract.WatchLogs(opts, "WithdrawSecurityFund", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilControlerWithdrawSecurityFund)
				if err := _LetsFilControler.contract.UnpackLog(event, "WithdrawSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawSecurityFund is a log parse operation binding the contract event 0x29c8190b82df1ee51e5380bc142f4613aacb0bd4078470790d35942959c3d0d5.
//
// Solidity: event WithdrawSecurityFund(uint256 indexed id, address indexed caller, uint256 amount)
func (_LetsFilControler *LetsFilControlerFilterer) ParseWithdrawSecurityFund(log types.Log) (*LetsFilControlerWithdrawSecurityFund, error) {
	event := new(LetsFilControlerWithdrawSecurityFund)
	if err := _LetsFilControler.contract.UnpackLog(event, "WithdrawSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
