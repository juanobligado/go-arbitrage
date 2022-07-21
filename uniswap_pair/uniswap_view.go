// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap_pair

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniswapViewABI is the input ABI used to generate the binding from.
const UniswapViewABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_pair\",\"type\":\"address[]\"}],\"name\":\"viewPair\",\"outputs\":[{\"internalType\":\"uint112[]\",\"name\":\"\",\"type\":\"uint112[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UniswapViewBin is the compiled bytecode used for deploying new contracts.
var UniswapViewBin = "0x608060405234801561001057600080fd5b506106f6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80632245f98614610030575b600080fd5b61004a600480360381019061004591906102a4565b610060565b60405161005791906103c9565b60405180910390f35b6060600083839050905060008160026100799190610424565b67ffffffffffffffff8111156100925761009161047e565b5b6040519080825280602002602001820160405280156100c05781602001602082028036833780820191505090505b50905060005b82811015610229576000808787848181106100e4576100e36104ad565b5b90506020020160208101906100f9919061053a565b73ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa158015610143573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016791906105cf565b5091509150818484600261017b9190610424565b8151811061018c5761018b6104ad565b5b60200260200101906dffffffffffffffffffffffffffff1690816dffffffffffffffffffffffffffff1681525050808460018560026101cb9190610424565b6101d59190610622565b815181106101e6576101e56104ad565b5b60200260200101906dffffffffffffffffffffffffffff1690816dffffffffffffffffffffffffffff16815250505050808061022190610678565b9150506100c6565b50809250505092915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126102645761026361023f565b5b8235905067ffffffffffffffff81111561028157610280610244565b5b60208301915083602082028301111561029d5761029c610249565b5b9250929050565b600080602083850312156102bb576102ba610235565b5b600083013567ffffffffffffffff8111156102d9576102d861023a565b5b6102e58582860161024e565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006dffffffffffffffffffffffffffff82169050919050565b6103408161031d565b82525050565b60006103528383610337565b60208301905092915050565b6000602082019050919050565b6000610376826102f1565b61038081856102fc565b935061038b8361030d565b8060005b838110156103bc5781516103a38882610346565b97506103ae8361035e565b92505060018101905061038f565b5085935050505092915050565b600060208201905081810360008301526103e3818461036b565b905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061042f826103eb565b915061043a836103eb565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610473576104726103f5565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610507826104dc565b9050919050565b610517816104fc565b811461052257600080fd5b50565b6000813590506105348161050e565b92915050565b6000602082840312156105505761054f610235565b5b600061055e84828501610525565b91505092915050565b6105708161031d565b811461057b57600080fd5b50565b60008151905061058d81610567565b92915050565b600063ffffffff82169050919050565b6105ac81610593565b81146105b757600080fd5b50565b6000815190506105c9816105a3565b92915050565b6000806000606084860312156105e8576105e7610235565b5b60006105f68682870161057e565b93505060206106078682870161057e565b9250506040610618868287016105ba565b9150509250925092565b600061062d826103eb565b9150610638836103eb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561066d5761066c6103f5565b5b828201905092915050565b6000610683826103eb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106b5576106b46103f5565b5b60018201905091905056fea26469706673582212201ebb62fce551be842a40b4c69199b724af0e9c9fdf059c985242572a520a985164736f6c634300080f0033"

// DeployUniswapView deploys a new Ethereum contract, binding an instance of UniswapView to it.
func DeployUniswapView(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapView, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapViewABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniswapViewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapView{UniswapViewCaller: UniswapViewCaller{contract: contract}, UniswapViewTransactor: UniswapViewTransactor{contract: contract}, UniswapViewFilterer: UniswapViewFilterer{contract: contract}}, nil
}

// UniswapView is an auto generated Go binding around an Ethereum contract.
type UniswapView struct {
	UniswapViewCaller     // Read-only binding to the contract
	UniswapViewTransactor // Write-only binding to the contract
	UniswapViewFilterer   // Log filterer for contract events
}

// UniswapViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapViewSession struct {
	Contract     *UniswapView      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapViewCallerSession struct {
	Contract *UniswapViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UniswapViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapViewTransactorSession struct {
	Contract     *UniswapViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniswapViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapViewRaw struct {
	Contract *UniswapView // Generic contract binding to access the raw methods on
}

// UniswapViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapViewCallerRaw struct {
	Contract *UniswapViewCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapViewTransactorRaw struct {
	Contract *UniswapViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapView creates a new instance of UniswapView, bound to a specific deployed contract.
func NewUniswapView(address common.Address, backend bind.ContractBackend) (*UniswapView, error) {
	contract, err := bindUniswapView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapView{UniswapViewCaller: UniswapViewCaller{contract: contract}, UniswapViewTransactor: UniswapViewTransactor{contract: contract}, UniswapViewFilterer: UniswapViewFilterer{contract: contract}}, nil
}

// NewUniswapViewCaller creates a new read-only instance of UniswapView, bound to a specific deployed contract.
func NewUniswapViewCaller(address common.Address, caller bind.ContractCaller) (*UniswapViewCaller, error) {
	contract, err := bindUniswapView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapViewCaller{contract: contract}, nil
}

// NewUniswapViewTransactor creates a new write-only instance of UniswapView, bound to a specific deployed contract.
func NewUniswapViewTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapViewTransactor, error) {
	contract, err := bindUniswapView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapViewTransactor{contract: contract}, nil
}

// NewUniswapViewFilterer creates a new log filterer instance of UniswapView, bound to a specific deployed contract.
func NewUniswapViewFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapViewFilterer, error) {
	contract, err := bindUniswapView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapViewFilterer{contract: contract}, nil
}

// bindUniswapView binds a generic wrapper to an already deployed contract.
func bindUniswapView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapView *UniswapViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapView.Contract.UniswapViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapView *UniswapViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapView.Contract.UniswapViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapView *UniswapViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapView.Contract.UniswapViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapView *UniswapViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapView *UniswapViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapView *UniswapViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapView.Contract.contract.Transact(opts, method, params...)
}

// ViewPair is a free data retrieval call binding the contract method 0x2245f986.
//
// Solidity: function viewPair(address[] _pair) view returns(uint112[])
func (_UniswapView *UniswapViewCaller) ViewPair(opts *bind.CallOpts, _pair []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapView.contract.Call(opts, &out, "viewPair", _pair)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ViewPair is a free data retrieval call binding the contract method 0x2245f986.
//
// Solidity: function viewPair(address[] _pair) view returns(uint112[])
func (_UniswapView *UniswapViewSession) ViewPair(_pair []common.Address) ([]*big.Int, error) {
	return _UniswapView.Contract.ViewPair(&_UniswapView.CallOpts, _pair)
}

// ViewPair is a free data retrieval call binding the contract method 0x2245f986.
//
// Solidity: function viewPair(address[] _pair) view returns(uint112[])
func (_UniswapView *UniswapViewCallerSession) ViewPair(_pair []common.Address) ([]*big.Int, error) {
	return _UniswapView.Contract.ViewPair(&_UniswapView.CallOpts, _pair)
}
