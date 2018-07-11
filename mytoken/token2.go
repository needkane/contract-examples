// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mytoken

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

// MyTokenABI is the input ABI used to generate the binding from.
const MyTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"spentAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"s\",\"type\":\"string\"}],\"name\":\"StringArg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"i8\",\"type\":\"int8\"},{\"indexed\":true,\"name\":\"i64\",\"type\":\"int64\"}],\"name\":\"IntArg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"ReceiveApproval\",\"type\":\"event\"}]"

// MyTokenBin is the compiled bytecode used for deploying new contracts.
const MyTokenBin = `606060405260126006556000600790600019169055341561001c57fe5b604051602080611c5d833981016040528080519060200190919050505b5b60005b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806000819055505b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405180905060405180910390a25b80600581600019169055505b505b611b2e8061012f6000396000f3006060604052361561011b576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461011d57806307da68f51461014b578063095ea7b31461015d57806313af4035146101b457806318160ddd146101ea57806323b872dd14610210578063313ce567146102865780633452f51d146102ac5780635ac801fe1461031557806369d3e20e1461033957806370a082311461036b57806375f12b21146103b55780637a9e5e4b146103df5780638402181f146104155780638da5cb5b1461047e57806390bc1693146104d057806395d89b4114610502578063a9059cbb14610530578063be9a655514610587578063bf7e214f14610599578063dd62ed3e146105eb575bfe5b341561012557fe5b61012d610654565b60405180826000191660001916815260200191505060405180910390f35b341561015357fe5b61015b61065a565b005b341561016557fe5b61019a600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061075e565b604051808215151515815260200191505060405180910390f35b34156101bc57fe5b6101e8600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061083c565b005b34156101f257fe5b6101fa610920565b6040518082815260200191505060405180910390f35b341561021857fe5b61026c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061092b565b604051808215151515815260200191505060405180910390f35b341561028e57fe5b610296610a0b565b6040518082815260200191505060405180910390f35b34156102b457fe5b6102fb600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080356fffffffffffffffffffffffffffffffff16906020019091905050610a11565b604051808215151515815260200191505060405180910390f35b341561031d57fe5b610337600480803560001916906020019091905050610a38565b005b341561034157fe5b61036960048080356fffffffffffffffffffffffffffffffff16906020019091905050610a7e565b005b341561037357fe5b61039f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610c44565b6040518082815260200191505060405180910390f35b34156103bd57fe5b6103c5610c8e565b604051808215151515815260200191505060405180910390f35b34156103e757fe5b610413600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ca1565b005b341561041d57fe5b610464600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080356fffffffffffffffffffffffffffffffff16906020019091905050610d85565b604051808215151515815260200191505060405180910390f35b341561048657fe5b61048e610dad565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104d857fe5b61050060048080356fffffffffffffffffffffffffffffffff16906020019091905050610dd3565b005b341561050a57fe5b610512610f99565b60405180826000191660001916815260200191505060405180910390f35b341561053857fe5b61056d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610f9f565b604051808215151515815260200191505060405180910390f35b341561058f57fe5b61059761107d565b005b34156105a157fe5b6105a9611181565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105f357fe5b61063e600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506111a7565b6040518082815260200191505060405180910390f35b60075481565b61069061068b336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a46001600460146101000a81548160ff0219169083151502179055505b5b50505b565b6000610779600460149054906101000a900460ff1615611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a461082f85856114a2565b92505b5b50505b92915050565b61087261086d336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405180905060405180910390a25b5b50565b600060005490505b90565b6000610946600460149054906101000a900460ff1615611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a46109fd868686611595565b92505b5b50505b9392505050565b60065481565b6000610a2f83836fffffffffffffffffffffffffffffffff16610f9f565b90505b92915050565b610a6e610a69336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b80600781600019169055505b5b50565b610ab4610aaf336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b610acd600460149054906101000a900460ff1615611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a4610bd4600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054846fffffffffffffffffffffffffffffffff166118f9565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610c35600054846fffffffffffffffffffffffffffffffff166118f9565b6000819055505b5b50505b5b50565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b919050565b600460149054906101000a900460ff1681565b610cd7610cd2336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405180905060405180910390a25b5b50565b6000610da48333846fffffffffffffffffffffffffffffffff1661092b565b90505b92915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610e09610e04336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b610e22600460149054906101000a900460ff1615611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a4610f29600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054846fffffffffffffffffffffffffffffffff16611913565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f8a600054846fffffffffffffffffffffffffffffffff16611913565b6000819055505b5b50505b5b50565b60055481565b6000610fba600460149054906101000a900460ff1615611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a4611070858561192d565b92505b5b50505b92915050565b6110b36110ae336000357fffffffff000000000000000000000000000000000000000000000000000000001661122f565b611491565b6000600060043591506024359050806000191682600019163373ffffffffffffffffffffffffffffffffffffffff166000357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19163460003660405180848152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a46000600460146101000a81548160ff0219169083151502179055505b5b50505b565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b92915050565b60003073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561126e576001905061148b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156112cd576001905061148b565b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561132d576000905061148b565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b70096138430856000604051602001526040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019350505050602060405180830381600087803b151561146957fe5b6102c65a03f1151561147757fe5b50505060405180519050905061148b565b5b5b5b92915050565b80151561149e5760006000fd5b5b50565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a3600190505b92915050565b600081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156115e257fe5b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561166a57fe5b6116f0600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611913565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506117b9600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611913565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611845600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836118f9565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b9392505050565b6000828284019150811015151561190c57fe5b5b92915050565b6000828284039150811115151561192657fe5b5b92915050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015151561197a57fe5b6119c3600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611913565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611a4f600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836118f9565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b929150505600a165627a7a723058204432a84cfe06a995bd95935559d84003b39a006720c2eabd96115f376347f9b80029454f530000000000000000000000000000000000000000000000000000000000`

// DeployMyToken deploys a new Ethereum contract, binding an instance of MyToken to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *MyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println(parsed)
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyTokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
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
