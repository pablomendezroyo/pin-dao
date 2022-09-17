// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pinner_dao

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
)

// PinnerDaoMetaData contains all meta data concerning the PinnerDao contract.
var PinnerDaoMetaData = &bind.MetaData{
        ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getENSList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PinnerDaoABI is the input ABI used to generate the binding from.
// Deprecated: Use PinnerDaoMetaData.ABI instead.
var PinnerDaoABI = PinnerDaoMetaData.ABI

// PinnerDao is an auto generated Go binding around an Ethereum contract.
type PinnerDao struct {
        PinnerDaoCaller     // Read-only binding to the contract
        PinnerDaoTransactor // Write-only binding to the contract
        PinnerDaoFilterer   // Log filterer for contract events
}

// PinnerDaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type PinnerDaoCaller struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerDaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PinnerDaoTransactor struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerDaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PinnerDaoFilterer struct {
        contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerDaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PinnerDaoSession struct {
        Contract     *PinnerDao        // Generic contract binding to set the session for
        CallOpts     bind.CallOpts     // Call options to use throughout this session
        TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinnerDaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PinnerDaoCallerSession struct {
        Contract *PinnerDaoCaller // Generic contract caller binding to set the session for
        CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PinnerDaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PinnerDaoTransactorSession struct {
        Contract     *PinnerDaoTransactor // Generic contract transactor binding to set the session for
        TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PinnerDaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type PinnerDaoRaw struct {
        Contract *PinnerDao // Generic contract binding to access the raw methods on
}

// PinnerDaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PinnerDaoCallerRaw struct {
        Contract *PinnerDaoCaller // Generic read-only contract binding to access the raw methods on
}

// PinnerDaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PinnerDaoTransactorRaw struct {
        Contract *PinnerDaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinnerDao creates a new instance of PinnerDao, bound to a specific deployed contract.
func NewPinnerDao(address common.Address, backend bind.ContractBackend) (*PinnerDao, error) {
        contract, err := bindPinnerDao(address, backend, backend, backend)
        if err != nil {
                return nil, err
        }
        return &PinnerDao{PinnerDaoCaller: PinnerDaoCaller{contract: contract}, PinnerDaoTransactor: PinnerDaoTransactor{contract: contract}, PinnerDaoFilterer: PinnerDaoFilterer{contract: contract}}, nil
}

// NewPinnerDaoCaller creates a new read-only instance of PinnerDao, bound to a specific deployed contract.
func NewPinnerDaoCaller(address common.Address, caller bind.ContractCaller) (*PinnerDaoCaller, error) {
        contract, err := bindPinnerDao(address, caller, nil, nil)
        if err != nil {
                return nil, err
        }
        return &PinnerDaoCaller{contract: contract}, nil
}

// NewPinnerDaoTransactor creates a new write-only instance of PinnerDao, bound to a specific deployed contract.
func NewPinnerDaoTransactor(address common.Address, transactor bind.ContractTransactor) (*PinnerDaoTransactor, error) {
        contract, err := bindPinnerDao(address, nil, transactor, nil)
        if err != nil {
                return nil, err
        }
        return &PinnerDaoTransactor{contract: contract}, nil
}

// NewPinnerDaoFilterer creates a new log filterer instance of PinnerDao, bound to a specific deployed contract.
func NewPinnerDaoFilterer(address common.Address, filterer bind.ContractFilterer) (*PinnerDaoFilterer, error) {
        contract, err := bindPinnerDao(address, nil, nil, filterer)
        if err != nil {
                return nil, err
        }
        return &PinnerDaoFilterer{contract: contract}, nil
}

// bindPinnerDao binds a generic wrapper to an already deployed contract.
func bindPinnerDao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
        parsed, err := abi.JSON(strings.NewReader(PinnerDaoABI))
        if err != nil {
                return nil, err
        }
        return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinnerDao *PinnerDaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
        return _PinnerDao.Contract.PinnerDaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinnerDao *PinnerDaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _PinnerDao.Contract.PinnerDaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinnerDao *PinnerDaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _PinnerDao.Contract.PinnerDaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinnerDao *PinnerDaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
        return _PinnerDao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinnerDao *PinnerDaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
        return _PinnerDao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinnerDao *PinnerDaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
        return _PinnerDao.Contract.contract.Transact(opts, method, params...)
}

// GetENSList is a free data retrieval call binding the contract method 0xd23e12c2.
//
// Solidity: function getENSList() pure returns(string[])
func (_PinnerDao *PinnerDaoCaller) GetENSList(opts *bind.CallOpts) ([]string, error) {
        var out []interface{}
        err := _PinnerDao.contract.Call(opts, &out, "getENSList")

        if err != nil {
                return *new([]string), err
        }

        out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

        return out0, err

}

// GetENSList is a free data retrieval call binding the contract method 0xd23e12c2.
//
// Solidity: function getENSList() pure returns(string[])
func (_PinnerDao *PinnerDaoSession) GetENSList() ([]string, error) {
        return _PinnerDao.Contract.GetENSList(&_PinnerDao.CallOpts)
}

// GetENSList is a free data retrieval call binding the contract method 0xd23e12c2.
//
// Solidity: function getENSList() pure returns(string[])
func (_PinnerDao *PinnerDaoCallerSession) GetENSList() ([]string, error) {
        return _PinnerDao.Contract.GetENSList(&_PinnerDao.CallOpts)
}