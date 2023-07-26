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

// LetsFilFactoryMetaData contains all meta data concerning the LetsFilFactory contract.
var LetsFilFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRaiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spFundShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filFiShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"raiseCompany\",\"type\":\"string\"}],\"internalType\":\"structILetsFilPackInfo.RaiseInfo\",\"name\":\"_raiseInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"spAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"companyId\",\"type\":\"uint256\"}],\"internalType\":\"structILetsFilPackInfo.NodeInfo\",\"name\":\"_nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sponsorOldRewardShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldRewardShare\",\"type\":\"uint256\"}],\"internalType\":\"structILetsFilPackInfo.ExtendInfo\",\"name\":\"_extendInfo\",\"type\":\"tuple\"}],\"name\":\"createRaisePlan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"planAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"raisePool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRaiseRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiseDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spFundShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filFiShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"raiseCompany\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structILetsFilPackInfo.RaiseInfo\",\"name\":\"raiseInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"spAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"companyId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILetsFilPackInfo.NodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sponsorOldRewardShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spOldRewardShare\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILetsFilPackInfo.ExtendInfo\",\"name\":\"extendInfo\",\"type\":\"tuple\"}],\"name\":\"CreateRaisePlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_annual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sealMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sealDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sealFine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_opsMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toolAddr\",\"type\":\"address\"}],\"name\":\"setToolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fineCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkAnnual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opsSecurityFundMinCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plans\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"raiseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"controlerAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolSealFineCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sealDelayCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sealMinCoeff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LetsFilFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LetsFilFactoryMetaData.ABI instead.
var LetsFilFactoryABI = LetsFilFactoryMetaData.ABI

// LetsFilFactory is an auto generated Go binding around an Ethereum contract.
type LetsFilFactory struct {
	LetsFilFactoryCaller     // Read-only binding to the contract
	LetsFilFactoryTransactor // Write-only binding to the contract
	LetsFilFactoryFilterer   // Log filterer for contract events
}

// LetsFilFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LetsFilFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LetsFilFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LetsFilFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LetsFilFactorySession struct {
	Contract     *LetsFilFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LetsFilFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LetsFilFactoryCallerSession struct {
	Contract *LetsFilFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LetsFilFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LetsFilFactoryTransactorSession struct {
	Contract     *LetsFilFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LetsFilFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LetsFilFactoryRaw struct {
	Contract *LetsFilFactory // Generic contract binding to access the raw methods on
}

// LetsFilFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LetsFilFactoryCallerRaw struct {
	Contract *LetsFilFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LetsFilFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LetsFilFactoryTransactorRaw struct {
	Contract *LetsFilFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLetsFilFactory creates a new instance of LetsFilFactory, bound to a specific deployed contract.
func NewLetsFilFactory(address common.Address, backend bind.ContractBackend) (*LetsFilFactory, error) {
	contract, err := bindLetsFilFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LetsFilFactory{LetsFilFactoryCaller: LetsFilFactoryCaller{contract: contract}, LetsFilFactoryTransactor: LetsFilFactoryTransactor{contract: contract}, LetsFilFactoryFilterer: LetsFilFactoryFilterer{contract: contract}}, nil
}

// NewLetsFilFactoryCaller creates a new read-only instance of LetsFilFactory, bound to a specific deployed contract.
func NewLetsFilFactoryCaller(address common.Address, caller bind.ContractCaller) (*LetsFilFactoryCaller, error) {
	contract, err := bindLetsFilFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilFactoryCaller{contract: contract}, nil
}

// NewLetsFilFactoryTransactor creates a new write-only instance of LetsFilFactory, bound to a specific deployed contract.
func NewLetsFilFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LetsFilFactoryTransactor, error) {
	contract, err := bindLetsFilFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilFactoryTransactor{contract: contract}, nil
}

// NewLetsFilFactoryFilterer creates a new log filterer instance of LetsFilFactory, bound to a specific deployed contract.
func NewLetsFilFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LetsFilFactoryFilterer, error) {
	contract, err := bindLetsFilFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LetsFilFactoryFilterer{contract: contract}, nil
}

// bindLetsFilFactory binds a generic wrapper to an already deployed contract.
func bindLetsFilFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LetsFilFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilFactory *LetsFilFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilFactory.Contract.LetsFilFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilFactory *LetsFilFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.LetsFilFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilFactory *LetsFilFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.LetsFilFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilFactory *LetsFilFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilFactory *LetsFilFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilFactory *LetsFilFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.contract.Transact(opts, method, params...)
}

// FeeCoeff is a free data retrieval call binding the contract method 0x31b162ad.
//
// Solidity: function feeCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) FeeCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "feeCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCoeff is a free data retrieval call binding the contract method 0x31b162ad.
//
// Solidity: function feeCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) FeeCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.FeeCoeff(&_LetsFilFactory.CallOpts)
}

// FeeCoeff is a free data retrieval call binding the contract method 0x31b162ad.
//
// Solidity: function feeCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) FeeCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.FeeCoeff(&_LetsFilFactory.CallOpts)
}

// FineCoeff is a free data retrieval call binding the contract method 0x8135c662.
//
// Solidity: function fineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) FineCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "fineCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FineCoeff is a free data retrieval call binding the contract method 0x8135c662.
//
// Solidity: function fineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) FineCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.FineCoeff(&_LetsFilFactory.CallOpts)
}

// FineCoeff is a free data retrieval call binding the contract method 0x8135c662.
//
// Solidity: function fineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) FineCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.FineCoeff(&_LetsFilFactory.CallOpts)
}

// NetworkAnnual is a free data retrieval call binding the contract method 0xa4fe3c1b.
//
// Solidity: function networkAnnual() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) NetworkAnnual(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "networkAnnual")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkAnnual is a free data retrieval call binding the contract method 0xa4fe3c1b.
//
// Solidity: function networkAnnual() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) NetworkAnnual() (*big.Int, error) {
	return _LetsFilFactory.Contract.NetworkAnnual(&_LetsFilFactory.CallOpts)
}

// NetworkAnnual is a free data retrieval call binding the contract method 0xa4fe3c1b.
//
// Solidity: function networkAnnual() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) NetworkAnnual() (*big.Int, error) {
	return _LetsFilFactory.Contract.NetworkAnnual(&_LetsFilFactory.CallOpts)
}

// OpsSecurityFundMinCoeff is a free data retrieval call binding the contract method 0xbc3818b5.
//
// Solidity: function opsSecurityFundMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) OpsSecurityFundMinCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "opsSecurityFundMinCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsSecurityFundMinCoeff is a free data retrieval call binding the contract method 0xbc3818b5.
//
// Solidity: function opsSecurityFundMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) OpsSecurityFundMinCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.OpsSecurityFundMinCoeff(&_LetsFilFactory.CallOpts)
}

// OpsSecurityFundMinCoeff is a free data retrieval call binding the contract method 0xbc3818b5.
//
// Solidity: function opsSecurityFundMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) OpsSecurityFundMinCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.OpsSecurityFundMinCoeff(&_LetsFilFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilFactory *LetsFilFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilFactory *LetsFilFactorySession) Owner() (common.Address, error) {
	return _LetsFilFactory.Contract.Owner(&_LetsFilFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilFactory *LetsFilFactoryCallerSession) Owner() (common.Address, error) {
	return _LetsFilFactory.Contract.Owner(&_LetsFilFactory.CallOpts)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(address sponsor, uint64 minerId, uint256 raiseId, address controlerAddr)
func (_LetsFilFactory *LetsFilFactoryCaller) Plans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sponsor       common.Address
	MinerId       uint64
	RaiseId       *big.Int
	ControlerAddr common.Address
}, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "plans", arg0)

	outstruct := new(struct {
		Sponsor       common.Address
		MinerId       uint64
		RaiseId       *big.Int
		ControlerAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sponsor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinerId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RaiseId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ControlerAddr = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(address sponsor, uint64 minerId, uint256 raiseId, address controlerAddr)
func (_LetsFilFactory *LetsFilFactorySession) Plans(arg0 *big.Int) (struct {
	Sponsor       common.Address
	MinerId       uint64
	RaiseId       *big.Int
	ControlerAddr common.Address
}, error) {
	return _LetsFilFactory.Contract.Plans(&_LetsFilFactory.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(address sponsor, uint64 minerId, uint256 raiseId, address controlerAddr)
func (_LetsFilFactory *LetsFilFactoryCallerSession) Plans(arg0 *big.Int) (struct {
	Sponsor       common.Address
	MinerId       uint64
	RaiseId       *big.Int
	ControlerAddr common.Address
}, error) {
	return _LetsFilFactory.Contract.Plans(&_LetsFilFactory.CallOpts, arg0)
}

// ProtocolSealFineCoeff is a free data retrieval call binding the contract method 0xd4c79a57.
//
// Solidity: function protocolSealFineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) ProtocolSealFineCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "protocolSealFineCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolSealFineCoeff is a free data retrieval call binding the contract method 0xd4c79a57.
//
// Solidity: function protocolSealFineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) ProtocolSealFineCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.ProtocolSealFineCoeff(&_LetsFilFactory.CallOpts)
}

// ProtocolSealFineCoeff is a free data retrieval call binding the contract method 0xd4c79a57.
//
// Solidity: function protocolSealFineCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) ProtocolSealFineCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.ProtocolSealFineCoeff(&_LetsFilFactory.CallOpts)
}

// RateBase is a free data retrieval call binding the contract method 0x014e95ba.
//
// Solidity: function rateBase() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) RateBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "rateBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateBase is a free data retrieval call binding the contract method 0x014e95ba.
//
// Solidity: function rateBase() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) RateBase() (*big.Int, error) {
	return _LetsFilFactory.Contract.RateBase(&_LetsFilFactory.CallOpts)
}

// RateBase is a free data retrieval call binding the contract method 0x014e95ba.
//
// Solidity: function rateBase() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) RateBase() (*big.Int, error) {
	return _LetsFilFactory.Contract.RateBase(&_LetsFilFactory.CallOpts)
}

// SealDelayCoeff is a free data retrieval call binding the contract method 0x6d374f15.
//
// Solidity: function sealDelayCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) SealDelayCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "sealDelayCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SealDelayCoeff is a free data retrieval call binding the contract method 0x6d374f15.
//
// Solidity: function sealDelayCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) SealDelayCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.SealDelayCoeff(&_LetsFilFactory.CallOpts)
}

// SealDelayCoeff is a free data retrieval call binding the contract method 0x6d374f15.
//
// Solidity: function sealDelayCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) SealDelayCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.SealDelayCoeff(&_LetsFilFactory.CallOpts)
}

// SealMinCoeff is a free data retrieval call binding the contract method 0xc3a8d9b2.
//
// Solidity: function sealMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCaller) SealMinCoeff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilFactory.contract.Call(opts, &out, "sealMinCoeff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SealMinCoeff is a free data retrieval call binding the contract method 0xc3a8d9b2.
//
// Solidity: function sealMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactorySession) SealMinCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.SealMinCoeff(&_LetsFilFactory.CallOpts)
}

// SealMinCoeff is a free data retrieval call binding the contract method 0xc3a8d9b2.
//
// Solidity: function sealMinCoeff() view returns(uint256)
func (_LetsFilFactory *LetsFilFactoryCallerSession) SealMinCoeff() (*big.Int, error) {
	return _LetsFilFactory.Contract.SealMinCoeff(&_LetsFilFactory.CallOpts)
}

// CreateRaisePlan is a paid mutator transaction binding the contract method 0xd559c31f.
//
// Solidity: function createRaisePlan((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns(address planAddress)
func (_LetsFilFactory *LetsFilFactoryTransactor) CreateRaisePlan(opts *bind.TransactOpts, _raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilFactory.contract.Transact(opts, "createRaisePlan", _raiseInfo, _nodeInfo, _extendInfo)
}

// CreateRaisePlan is a paid mutator transaction binding the contract method 0xd559c31f.
//
// Solidity: function createRaisePlan((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns(address planAddress)
func (_LetsFilFactory *LetsFilFactorySession) CreateRaisePlan(_raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.CreateRaisePlan(&_LetsFilFactory.TransactOpts, _raiseInfo, _nodeInfo, _extendInfo)
}

// CreateRaisePlan is a paid mutator transaction binding the contract method 0xd559c31f.
//
// Solidity: function createRaisePlan((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) _raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) _nodeInfo, (uint256,uint256,uint256,uint256,uint256) _extendInfo) returns(address planAddress)
func (_LetsFilFactory *LetsFilFactoryTransactorSession) CreateRaisePlan(_raiseInfo ILetsFilPackInfoRaiseInfo, _nodeInfo ILetsFilPackInfoNodeInfo, _extendInfo ILetsFilPackInfoExtendInfo) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.CreateRaisePlan(&_LetsFilFactory.TransactOpts, _raiseInfo, _nodeInfo, _extendInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilFactory *LetsFilFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilFactory *LetsFilFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _LetsFilFactory.Contract.RenounceOwnership(&_LetsFilFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilFactory *LetsFilFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LetsFilFactory.Contract.RenounceOwnership(&_LetsFilFactory.TransactOpts)
}

// SetParam is a paid mutator transaction binding the contract method 0xc134b057.
//
// Solidity: function setParam(uint256 _annual, uint256 _base, uint256 _fine, uint256 _sealMin, uint256 _sealDelay, uint256 _sealFine, uint256 _opsMin, uint256 _fee) returns()
func (_LetsFilFactory *LetsFilFactoryTransactor) SetParam(opts *bind.TransactOpts, _annual *big.Int, _base *big.Int, _fine *big.Int, _sealMin *big.Int, _sealDelay *big.Int, _sealFine *big.Int, _opsMin *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _LetsFilFactory.contract.Transact(opts, "setParam", _annual, _base, _fine, _sealMin, _sealDelay, _sealFine, _opsMin, _fee)
}

// SetParam is a paid mutator transaction binding the contract method 0xc134b057.
//
// Solidity: function setParam(uint256 _annual, uint256 _base, uint256 _fine, uint256 _sealMin, uint256 _sealDelay, uint256 _sealFine, uint256 _opsMin, uint256 _fee) returns()
func (_LetsFilFactory *LetsFilFactorySession) SetParam(_annual *big.Int, _base *big.Int, _fine *big.Int, _sealMin *big.Int, _sealDelay *big.Int, _sealFine *big.Int, _opsMin *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.SetParam(&_LetsFilFactory.TransactOpts, _annual, _base, _fine, _sealMin, _sealDelay, _sealFine, _opsMin, _fee)
}

// SetParam is a paid mutator transaction binding the contract method 0xc134b057.
//
// Solidity: function setParam(uint256 _annual, uint256 _base, uint256 _fine, uint256 _sealMin, uint256 _sealDelay, uint256 _sealFine, uint256 _opsMin, uint256 _fee) returns()
func (_LetsFilFactory *LetsFilFactoryTransactorSession) SetParam(_annual *big.Int, _base *big.Int, _fine *big.Int, _sealMin *big.Int, _sealDelay *big.Int, _sealFine *big.Int, _opsMin *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.SetParam(&_LetsFilFactory.TransactOpts, _annual, _base, _fine, _sealMin, _sealDelay, _sealFine, _opsMin, _fee)
}

// SetToolAddress is a paid mutator transaction binding the contract method 0x14fc9684.
//
// Solidity: function setToolAddress(address toolAddr) returns()
func (_LetsFilFactory *LetsFilFactoryTransactor) SetToolAddress(opts *bind.TransactOpts, toolAddr common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.contract.Transact(opts, "setToolAddress", toolAddr)
}

// SetToolAddress is a paid mutator transaction binding the contract method 0x14fc9684.
//
// Solidity: function setToolAddress(address toolAddr) returns()
func (_LetsFilFactory *LetsFilFactorySession) SetToolAddress(toolAddr common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.SetToolAddress(&_LetsFilFactory.TransactOpts, toolAddr)
}

// SetToolAddress is a paid mutator transaction binding the contract method 0x14fc9684.
//
// Solidity: function setToolAddress(address toolAddr) returns()
func (_LetsFilFactory *LetsFilFactoryTransactorSession) SetToolAddress(toolAddr common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.SetToolAddress(&_LetsFilFactory.TransactOpts, toolAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilFactory *LetsFilFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilFactory *LetsFilFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.TransferOwnership(&_LetsFilFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilFactory *LetsFilFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilFactory.Contract.TransferOwnership(&_LetsFilFactory.TransactOpts, newOwner)
}

// LetsFilFactoryCreateRaisePlanIterator is returned from FilterCreateRaisePlan and is used to iterate over the raw logs and unpacked data for CreateRaisePlan events raised by the LetsFilFactory contract.
type LetsFilFactoryCreateRaisePlanIterator struct {
	Event *LetsFilFactoryCreateRaisePlan // Event containing the contract specifics and raw log

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
func (it *LetsFilFactoryCreateRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilFactoryCreateRaisePlan)
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
		it.Event = new(LetsFilFactoryCreateRaisePlan)
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
func (it *LetsFilFactoryCreateRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilFactoryCreateRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilFactoryCreateRaisePlan represents a CreateRaisePlan event raised by the LetsFilFactory contract.
type LetsFilFactoryCreateRaisePlan struct {
	Id         *big.Int
	RaisePool  common.Address
	Caller     common.Address
	RaiseInfo  ILetsFilPackInfoRaiseInfo
	NodeInfo   ILetsFilPackInfoNodeInfo
	ExtendInfo ILetsFilPackInfoExtendInfo
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateRaisePlan is a free log retrieval operation binding the contract event 0x947c78218cd393fd9b31b2d75ad9afe17be1dfb075fe63502c141212b937b176.
//
// Solidity: event CreateRaisePlan(uint256 id, address raisePool, address caller, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) nodeInfo, (uint256,uint256,uint256,uint256,uint256) extendInfo)
func (_LetsFilFactory *LetsFilFactoryFilterer) FilterCreateRaisePlan(opts *bind.FilterOpts) (*LetsFilFactoryCreateRaisePlanIterator, error) {

	logs, sub, err := _LetsFilFactory.contract.FilterLogs(opts, "CreateRaisePlan")
	if err != nil {
		return nil, err
	}
	return &LetsFilFactoryCreateRaisePlanIterator{contract: _LetsFilFactory.contract, event: "CreateRaisePlan", logs: logs, sub: sub}, nil
}

// WatchCreateRaisePlan is a free log subscription operation binding the contract event 0x947c78218cd393fd9b31b2d75ad9afe17be1dfb075fe63502c141212b937b176.
//
// Solidity: event CreateRaisePlan(uint256 id, address raisePool, address caller, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) nodeInfo, (uint256,uint256,uint256,uint256,uint256) extendInfo)
func (_LetsFilFactory *LetsFilFactoryFilterer) WatchCreateRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilFactoryCreateRaisePlan) (event.Subscription, error) {

	logs, sub, err := _LetsFilFactory.contract.WatchLogs(opts, "CreateRaisePlan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilFactoryCreateRaisePlan)
				if err := _LetsFilFactory.contract.UnpackLog(event, "CreateRaisePlan", log); err != nil {
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

// ParseCreateRaisePlan is a log parse operation binding the contract event 0x947c78218cd393fd9b31b2d75ad9afe17be1dfb075fe63502c141212b937b176.
//
// Solidity: event CreateRaisePlan(uint256 id, address raisePool, address caller, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,string) raiseInfo, (uint64,uint256,uint256,uint256,uint256,uint256,address,uint256) nodeInfo, (uint256,uint256,uint256,uint256,uint256) extendInfo)
func (_LetsFilFactory *LetsFilFactoryFilterer) ParseCreateRaisePlan(log types.Log) (*LetsFilFactoryCreateRaisePlan, error) {
	event := new(LetsFilFactoryCreateRaisePlan)
	if err := _LetsFilFactory.contract.UnpackLog(event, "CreateRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LetsFilFactory contract.
type LetsFilFactoryOwnershipTransferredIterator struct {
	Event *LetsFilFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LetsFilFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilFactoryOwnershipTransferred)
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
		it.Event = new(LetsFilFactoryOwnershipTransferred)
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
func (it *LetsFilFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the LetsFilFactory contract.
type LetsFilFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilFactory *LetsFilFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LetsFilFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LetsFilFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilFactoryOwnershipTransferredIterator{contract: _LetsFilFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilFactory *LetsFilFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LetsFilFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LetsFilFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilFactoryOwnershipTransferred)
				if err := _LetsFilFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilFactory *LetsFilFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LetsFilFactoryOwnershipTransferred, error) {
	event := new(LetsFilFactoryOwnershipTransferred)
	if err := _LetsFilFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
