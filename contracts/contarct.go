// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type ContractJson struct {
	Bytecode string `json:"bytecode"`
}

// DeployMyToken deploys a new Ethereum contract, binding an instance of MyToken to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *types.Transaction, *MyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiJson))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Bytecode), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}}, nil
}

// MyToken is an auto generated Go binding around an Ethereum contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
}

// MyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	WatchOpts    bind.WatchOpts    // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract  *MyTokenCaller // Generic contract caller binding to set the session for
	WatchOpts bind.WatchOpts // Call options to use throughout this session
	CallOpts  bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address common.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var ret0 *big.Int
	err := _MyToken.contract.Call(opts, &ret0, "balanceOf", arg0)
	fmt.Println(*ret0, err, "====181=====", arg0)
	return ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenCallerSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCaller) SpentAllowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "spentAllowance", arg0, arg1)
	return *ret0, err
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.SpentAllowance(&_MyToken.CallOpts, arg0, arg1)
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.SpentAllowance(&_MyToken.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.Contract.ApproveAndCall(&_MyToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.Contract.ApproveAndCall(&_MyToken.TransactOpts, _spender, _value, _extraData)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenSession) Transfer(_to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenTransactorSession) Transfer(_to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, _from, _to, _value)
}

func (_MyToken *MyTokenCaller) EventSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber, eventName string) (<-chan types.Log, ethereum.Subscription, error) {
	return _MyToken.contract.WatchLogs(opts, eventName)
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventIntArgSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "IntArg")
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventIntArgSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventIntArgSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventIntArgSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventIntArgSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventReceiveApprovalSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "ReceiveApproval")
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventReceiveApprovalSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventReceiveApprovalSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventReceiveApprovalSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventReceiveApprovalSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventStringArgSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "StringArg")
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventStringArgSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventStringArgSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventStringArgSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventStringArgSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventTransferSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "Transfer")
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventTransferSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventTransferSubscribe(&t.WatchOpts, fromBlock, toBlock)
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventTransferSubscribe(opts *bind.WatchOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventTransferSubscribe(&t.WatchOpts, fromBlock, toBlock)
}
