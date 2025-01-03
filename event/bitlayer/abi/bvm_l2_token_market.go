// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// L2TokenMarketMetaData contains all meta data concerning the L2TokenMarket contract.
var L2TokenMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PerFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TokenTicker\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_MultisigWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"list\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unList\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"List\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderConfirm\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnList\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// L2TokenMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use L2TokenMarketMetaData.ABI instead.
var L2TokenMarketABI = L2TokenMarketMetaData.ABI

// L2TokenMarket is an auto generated Go binding around an Ethereum contract.
type L2TokenMarket struct {
	L2TokenMarketCaller     // Read-only binding to the contract
	L2TokenMarketTransactor // Write-only binding to the contract
	L2TokenMarketFilterer   // Log filterer for contract events
}

// L2TokenMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2TokenMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2TokenMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2TokenMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2TokenMarketSession struct {
	Contract     *L2TokenMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2TokenMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2TokenMarketCallerSession struct {
	Contract *L2TokenMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2TokenMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TokenMarketTransactorSession struct {
	Contract     *L2TokenMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2TokenMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2TokenMarketRaw struct {
	Contract *L2TokenMarket // Generic contract binding to access the raw methods on
}

// L2TokenMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2TokenMarketCallerRaw struct {
	Contract *L2TokenMarketCaller // Generic read-only contract binding to access the raw methods on
}

// L2TokenMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TokenMarketTransactorRaw struct {
	Contract *L2TokenMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2TokenMarket creates a new instance of L2TokenMarket, bound to a specific deployed contract.
func NewL2TokenMarket(address common.Address, backend bind.ContractBackend) (*L2TokenMarket, error) {
	contract, err := bindL2TokenMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarket{L2TokenMarketCaller: L2TokenMarketCaller{contract: contract}, L2TokenMarketTransactor: L2TokenMarketTransactor{contract: contract}, L2TokenMarketFilterer: L2TokenMarketFilterer{contract: contract}}, nil
}

// NewL2TokenMarketCaller creates a new read-only instance of L2TokenMarket, bound to a specific deployed contract.
func NewL2TokenMarketCaller(address common.Address, caller bind.ContractCaller) (*L2TokenMarketCaller, error) {
	contract, err := bindL2TokenMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketCaller{contract: contract}, nil
}

// NewL2TokenMarketTransactor creates a new write-only instance of L2TokenMarket, bound to a specific deployed contract.
func NewL2TokenMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*L2TokenMarketTransactor, error) {
	contract, err := bindL2TokenMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketTransactor{contract: contract}, nil
}

// NewL2TokenMarketFilterer creates a new log filterer instance of L2TokenMarket, bound to a specific deployed contract.
func NewL2TokenMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*L2TokenMarketFilterer, error) {
	contract, err := bindL2TokenMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketFilterer{contract: contract}, nil
}

// bindL2TokenMarket binds a generic wrapper to an already deployed contract.
func bindL2TokenMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2TokenMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenMarket *L2TokenMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenMarket.Contract.L2TokenMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenMarket *L2TokenMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.L2TokenMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenMarket *L2TokenMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.L2TokenMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenMarket *L2TokenMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenMarket *L2TokenMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenMarket *L2TokenMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2TokenMarket.Contract.DEFAULTADMINROLE(&_L2TokenMarket.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2TokenMarket.Contract.DEFAULTADMINROLE(&_L2TokenMarket.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "Nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketSession) Nonce() (*big.Int, error) {
	return _L2TokenMarket.Contract.Nonce(&_L2TokenMarket.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCallerSession) Nonce() (*big.Int, error) {
	return _L2TokenMarket.Contract.Nonce(&_L2TokenMarket.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketSession) PerFee() (*big.Int, error) {
	return _L2TokenMarket.Contract.PerFee(&_L2TokenMarket.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCallerSession) PerFee() (*big.Int, error) {
	return _L2TokenMarket.Contract.PerFee(&_L2TokenMarket.CallOpts)
}

// TokenTicker is a free data retrieval call binding the contract method 0x914d3407.
//
// Solidity: function TokenTicker(address , uint256 ) view returns(uint256 amount, uint256 price, address seller)
func (_L2TokenMarket *L2TokenMarketCaller) TokenTicker(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Price  *big.Int
	Seller common.Address
}, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "TokenTicker", arg0, arg1)

	outstruct := new(struct {
		Amount *big.Int
		Price  *big.Int
		Seller common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenTicker is a free data retrieval call binding the contract method 0x914d3407.
//
// Solidity: function TokenTicker(address , uint256 ) view returns(uint256 amount, uint256 price, address seller)
func (_L2TokenMarket *L2TokenMarketSession) TokenTicker(arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Price  *big.Int
	Seller common.Address
}, error) {
	return _L2TokenMarket.Contract.TokenTicker(&_L2TokenMarket.CallOpts, arg0, arg1)
}

// TokenTicker is a free data retrieval call binding the contract method 0x914d3407.
//
// Solidity: function TokenTicker(address , uint256 ) view returns(uint256 amount, uint256 price, address seller)
func (_L2TokenMarket *L2TokenMarketCallerSession) TokenTicker(arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Price  *big.Int
	Seller common.Address
}, error) {
	return _L2TokenMarket.Contract.TokenTicker(&_L2TokenMarket.CallOpts, arg0, arg1)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketSession) GetFee() (*big.Int, error) {
	return _L2TokenMarket.Contract.GetFee(&_L2TokenMarket.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_L2TokenMarket *L2TokenMarketCallerSession) GetFee() (*big.Int, error) {
	return _L2TokenMarket.Contract.GetFee(&_L2TokenMarket.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2TokenMarket.Contract.GetRoleAdmin(&_L2TokenMarket.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2TokenMarket *L2TokenMarketCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2TokenMarket.Contract.GetRoleAdmin(&_L2TokenMarket.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2TokenMarket *L2TokenMarketCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2TokenMarket.Contract.HasRole(&_L2TokenMarket.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2TokenMarket *L2TokenMarketCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2TokenMarket.Contract.HasRole(&_L2TokenMarket.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2TokenMarket *L2TokenMarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) Paused() (bool, error) {
	return _L2TokenMarket.Contract.Paused(&_L2TokenMarket.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2TokenMarket *L2TokenMarketCallerSession) Paused() (bool, error) {
	return _L2TokenMarket.Contract.Paused(&_L2TokenMarket.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2TokenMarket *L2TokenMarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2TokenMarket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2TokenMarket.Contract.SupportsInterface(&_L2TokenMarket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2TokenMarket *L2TokenMarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2TokenMarket.Contract.SupportsInterface(&_L2TokenMarket.CallOpts, interfaceId)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _nonce) payable returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactor) Buy(opts *bind.TransactOpts, _token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "buy", _token, _nonce)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _nonce) payable returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) Buy(_token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Buy(&_L2TokenMarket.TransactOpts, _token, _nonce)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _nonce) payable returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactorSession) Buy(_token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Buy(&_L2TokenMarket.TransactOpts, _token, _nonce)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.GrantRole(&_L2TokenMarket.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.GrantRole(&_L2TokenMarket.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2TokenMarket *L2TokenMarketTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "initialize", _MultisigWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2TokenMarket *L2TokenMarketSession) Initialize(_MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Initialize(&_L2TokenMarket.TransactOpts, _MultisigWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) Initialize(_MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Initialize(&_L2TokenMarket.TransactOpts, _MultisigWallet)
}

// List is a paid mutator transaction binding the contract method 0xdda342bb.
//
// Solidity: function list(address _token, uint256 _amount, uint256 _price) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactor) List(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "list", _token, _amount, _price)
}

// List is a paid mutator transaction binding the contract method 0xdda342bb.
//
// Solidity: function list(address _token, uint256 _amount, uint256 _price) returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) List(_token common.Address, _amount *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.List(&_L2TokenMarket.TransactOpts, _token, _amount, _price)
}

// List is a paid mutator transaction binding the contract method 0xdda342bb.
//
// Solidity: function list(address _token, uint256 _amount, uint256 _price) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactorSession) List(_token common.Address, _amount *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.List(&_L2TokenMarket.TransactOpts, _token, _amount, _price)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2TokenMarket *L2TokenMarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2TokenMarket *L2TokenMarketSession) Pause() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Pause(&_L2TokenMarket.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) Pause() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Pause(&_L2TokenMarket.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2TokenMarket *L2TokenMarketTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2TokenMarket *L2TokenMarketSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.RenounceRole(&_L2TokenMarket.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.RenounceRole(&_L2TokenMarket.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.RevokeRole(&_L2TokenMarket.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.RevokeRole(&_L2TokenMarket.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.SetFee(&_L2TokenMarket.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.SetFee(&_L2TokenMarket.TransactOpts, _fee)
}

// UnList is a paid mutator transaction binding the contract method 0x1ce99f34.
//
// Solidity: function unList(address _token, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactor) UnList(opts *bind.TransactOpts, _token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "unList", _token, _nonce)
}

// UnList is a paid mutator transaction binding the contract method 0x1ce99f34.
//
// Solidity: function unList(address _token, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) UnList(_token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.UnList(&_L2TokenMarket.TransactOpts, _token, _nonce)
}

// UnList is a paid mutator transaction binding the contract method 0x1ce99f34.
//
// Solidity: function unList(address _token, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactorSession) UnList(_token common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.UnList(&_L2TokenMarket.TransactOpts, _token, _nonce)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2TokenMarket *L2TokenMarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2TokenMarket *L2TokenMarketSession) Unpause() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Unpause(&_L2TokenMarket.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Unpause(&_L2TokenMarket.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xd2c3aaf2.
//
// Solidity: function update(address _token, uint256 _amount, uint256 _price, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactor) Update(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _price *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "update", _token, _amount, _price, _nonce)
}

// Update is a paid mutator transaction binding the contract method 0xd2c3aaf2.
//
// Solidity: function update(address _token, uint256 _amount, uint256 _price, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketSession) Update(_token common.Address, _amount *big.Int, _price *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Update(&_L2TokenMarket.TransactOpts, _token, _amount, _price, _nonce)
}

// Update is a paid mutator transaction binding the contract method 0xd2c3aaf2.
//
// Solidity: function update(address _token, uint256 _amount, uint256 _price, uint256 _nonce) returns(bool)
func (_L2TokenMarket *L2TokenMarketTransactorSession) Update(_token common.Address, _amount *big.Int, _price *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2TokenMarket.Contract.Update(&_L2TokenMarket.TransactOpts, _token, _amount, _price, _nonce)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_L2TokenMarket *L2TokenMarketTransactor) WithdrawFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenMarket.contract.Transact(opts, "withdrawFee")
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_L2TokenMarket *L2TokenMarketSession) WithdrawFee() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.WithdrawFee(&_L2TokenMarket.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_L2TokenMarket *L2TokenMarketTransactorSession) WithdrawFee() (*types.Transaction, error) {
	return _L2TokenMarket.Contract.WithdrawFee(&_L2TokenMarket.TransactOpts)
}

// L2TokenMarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2TokenMarket contract.
type L2TokenMarketInitializedIterator struct {
	Event *L2TokenMarketInitialized // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketInitialized)
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
		it.Event = new(L2TokenMarketInitialized)
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
func (it *L2TokenMarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketInitialized represents a Initialized event raised by the L2TokenMarket contract.
type L2TokenMarketInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2TokenMarketInitializedIterator, error) {

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketInitializedIterator{contract: _L2TokenMarket.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2TokenMarketInitialized) (event.Subscription, error) {

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketInitialized)
				if err := _L2TokenMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseInitialized(log types.Log) (*L2TokenMarketInitialized, error) {
	event := new(L2TokenMarketInitialized)
	if err := _L2TokenMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketListIterator is returned from FilterList and is used to iterate over the raw logs and unpacked data for List events raised by the L2TokenMarket contract.
type L2TokenMarketListIterator struct {
	Event *L2TokenMarketList // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketList)
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
		it.Event = new(L2TokenMarketList)
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
func (it *L2TokenMarketListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketList represents a List event raised by the L2TokenMarket contract.
type L2TokenMarketList struct {
	Token  common.Address
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterList is a free log retrieval operation binding the contract event 0xed0678e15c07b5471c509ea23b0f1e212d44a4f0c5017e2a958b5fdbb774ab1a.
//
// Solidity: event List(address indexed token, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterList(opts *bind.FilterOpts, token []common.Address, seller []common.Address) (*L2TokenMarketListIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "List", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketListIterator{contract: _L2TokenMarket.contract, event: "List", logs: logs, sub: sub}, nil
}

// WatchList is a free log subscription operation binding the contract event 0xed0678e15c07b5471c509ea23b0f1e212d44a4f0c5017e2a958b5fdbb774ab1a.
//
// Solidity: event List(address indexed token, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchList(opts *bind.WatchOpts, sink chan<- *L2TokenMarketList, token []common.Address, seller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "List", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketList)
				if err := _L2TokenMarket.contract.UnpackLog(event, "List", log); err != nil {
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

// ParseList is a log parse operation binding the contract event 0xed0678e15c07b5471c509ea23b0f1e212d44a4f0c5017e2a958b5fdbb774ab1a.
//
// Solidity: event List(address indexed token, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseList(log types.Log) (*L2TokenMarketList, error) {
	event := new(L2TokenMarketList)
	if err := _L2TokenMarket.contract.UnpackLog(event, "List", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketOrderConfirmIterator is returned from FilterOrderConfirm and is used to iterate over the raw logs and unpacked data for OrderConfirm events raised by the L2TokenMarket contract.
type L2TokenMarketOrderConfirmIterator struct {
	Event *L2TokenMarketOrderConfirm // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketOrderConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketOrderConfirm)
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
		it.Event = new(L2TokenMarketOrderConfirm)
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
func (it *L2TokenMarketOrderConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketOrderConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketOrderConfirm represents a OrderConfirm event raised by the L2TokenMarket contract.
type L2TokenMarketOrderConfirm struct {
	Token  common.Address
	Buyer  common.Address
	Seller common.Address
	Amount *big.Int
	Price  *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderConfirm is a free log retrieval operation binding the contract event 0xb22ab3de24474240366ab0f5d203176ac09caef1a5239244228a3af5151d3357.
//
// Solidity: event OrderConfirm(address indexed token, address indexed buyer, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterOrderConfirm(opts *bind.FilterOpts, token []common.Address, buyer []common.Address, seller []common.Address) (*L2TokenMarketOrderConfirmIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "OrderConfirm", tokenRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketOrderConfirmIterator{contract: _L2TokenMarket.contract, event: "OrderConfirm", logs: logs, sub: sub}, nil
}

// WatchOrderConfirm is a free log subscription operation binding the contract event 0xb22ab3de24474240366ab0f5d203176ac09caef1a5239244228a3af5151d3357.
//
// Solidity: event OrderConfirm(address indexed token, address indexed buyer, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchOrderConfirm(opts *bind.WatchOpts, sink chan<- *L2TokenMarketOrderConfirm, token []common.Address, buyer []common.Address, seller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "OrderConfirm", tokenRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketOrderConfirm)
				if err := _L2TokenMarket.contract.UnpackLog(event, "OrderConfirm", log); err != nil {
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

// ParseOrderConfirm is a log parse operation binding the contract event 0xb22ab3de24474240366ab0f5d203176ac09caef1a5239244228a3af5151d3357.
//
// Solidity: event OrderConfirm(address indexed token, address indexed buyer, address indexed seller, uint256 amount, uint256 price, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseOrderConfirm(log types.Log) (*L2TokenMarketOrderConfirm, error) {
	event := new(L2TokenMarketOrderConfirm)
	if err := _L2TokenMarket.contract.UnpackLog(event, "OrderConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2TokenMarket contract.
type L2TokenMarketPausedIterator struct {
	Event *L2TokenMarketPaused // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketPaused)
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
		it.Event = new(L2TokenMarketPaused)
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
func (it *L2TokenMarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketPaused represents a Paused event raised by the L2TokenMarket contract.
type L2TokenMarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterPaused(opts *bind.FilterOpts) (*L2TokenMarketPausedIterator, error) {

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketPausedIterator{contract: _L2TokenMarket.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2TokenMarketPaused) (event.Subscription, error) {

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketPaused)
				if err := _L2TokenMarket.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) ParsePaused(log types.Log) (*L2TokenMarketPaused, error) {
	event := new(L2TokenMarketPaused)
	if err := _L2TokenMarket.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L2TokenMarket contract.
type L2TokenMarketRoleAdminChangedIterator struct {
	Event *L2TokenMarketRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketRoleAdminChanged)
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
		it.Event = new(L2TokenMarketRoleAdminChanged)
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
func (it *L2TokenMarketRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketRoleAdminChanged represents a RoleAdminChanged event raised by the L2TokenMarket contract.
type L2TokenMarketRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L2TokenMarketRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketRoleAdminChangedIterator{contract: _L2TokenMarket.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L2TokenMarketRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketRoleAdminChanged)
				if err := _L2TokenMarket.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseRoleAdminChanged(log types.Log) (*L2TokenMarketRoleAdminChanged, error) {
	event := new(L2TokenMarketRoleAdminChanged)
	if err := _L2TokenMarket.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2TokenMarket contract.
type L2TokenMarketRoleGrantedIterator struct {
	Event *L2TokenMarketRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketRoleGranted)
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
		it.Event = new(L2TokenMarketRoleGranted)
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
func (it *L2TokenMarketRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketRoleGranted represents a RoleGranted event raised by the L2TokenMarket contract.
type L2TokenMarketRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2TokenMarketRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketRoleGrantedIterator{contract: _L2TokenMarket.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2TokenMarketRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketRoleGranted)
				if err := _L2TokenMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseRoleGranted(log types.Log) (*L2TokenMarketRoleGranted, error) {
	event := new(L2TokenMarketRoleGranted)
	if err := _L2TokenMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2TokenMarket contract.
type L2TokenMarketRoleRevokedIterator struct {
	Event *L2TokenMarketRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketRoleRevoked)
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
		it.Event = new(L2TokenMarketRoleRevoked)
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
func (it *L2TokenMarketRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketRoleRevoked represents a RoleRevoked event raised by the L2TokenMarket contract.
type L2TokenMarketRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2TokenMarketRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketRoleRevokedIterator{contract: _L2TokenMarket.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2TokenMarketRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketRoleRevoked)
				if err := _L2TokenMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseRoleRevoked(log types.Log) (*L2TokenMarketRoleRevoked, error) {
	event := new(L2TokenMarketRoleRevoked)
	if err := _L2TokenMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketUnListIterator is returned from FilterUnList and is used to iterate over the raw logs and unpacked data for UnList events raised by the L2TokenMarket contract.
type L2TokenMarketUnListIterator struct {
	Event *L2TokenMarketUnList // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketUnListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketUnList)
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
		it.Event = new(L2TokenMarketUnList)
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
func (it *L2TokenMarketUnListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketUnListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketUnList represents a UnList event raised by the L2TokenMarket contract.
type L2TokenMarketUnList struct {
	Token  common.Address
	Seller common.Address
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnList is a free log retrieval operation binding the contract event 0x123308a699bb6bc81b0d867240aab5a60ffdf26cf71f4937e476645eac7eed98.
//
// Solidity: event UnList(address indexed token, address indexed seller, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterUnList(opts *bind.FilterOpts, token []common.Address, seller []common.Address) (*L2TokenMarketUnListIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "UnList", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketUnListIterator{contract: _L2TokenMarket.contract, event: "UnList", logs: logs, sub: sub}, nil
}

// WatchUnList is a free log subscription operation binding the contract event 0x123308a699bb6bc81b0d867240aab5a60ffdf26cf71f4937e476645eac7eed98.
//
// Solidity: event UnList(address indexed token, address indexed seller, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchUnList(opts *bind.WatchOpts, sink chan<- *L2TokenMarketUnList, token []common.Address, seller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "UnList", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketUnList)
				if err := _L2TokenMarket.contract.UnpackLog(event, "UnList", log); err != nil {
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

// ParseUnList is a log parse operation binding the contract event 0x123308a699bb6bc81b0d867240aab5a60ffdf26cf71f4937e476645eac7eed98.
//
// Solidity: event UnList(address indexed token, address indexed seller, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseUnList(log types.Log) (*L2TokenMarketUnList, error) {
	event := new(L2TokenMarketUnList)
	if err := _L2TokenMarket.contract.UnpackLog(event, "UnList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2TokenMarket contract.
type L2TokenMarketUnpausedIterator struct {
	Event *L2TokenMarketUnpaused // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketUnpaused)
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
		it.Event = new(L2TokenMarketUnpaused)
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
func (it *L2TokenMarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketUnpaused represents a Unpaused event raised by the L2TokenMarket contract.
type L2TokenMarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2TokenMarketUnpausedIterator, error) {

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketUnpausedIterator{contract: _L2TokenMarket.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2TokenMarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketUnpaused)
				if err := _L2TokenMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseUnpaused(log types.Log) (*L2TokenMarketUnpaused, error) {
	event := new(L2TokenMarketUnpaused)
	if err := _L2TokenMarket.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TokenMarketUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the L2TokenMarket contract.
type L2TokenMarketUpdateIterator struct {
	Event *L2TokenMarketUpdate // Event containing the contract specifics and raw log

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
func (it *L2TokenMarketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TokenMarketUpdate)
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
		it.Event = new(L2TokenMarketUpdate)
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
func (it *L2TokenMarketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TokenMarketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TokenMarketUpdate represents a Update event raised by the L2TokenMarket contract.
type L2TokenMarketUpdate struct {
	Token  common.Address
	Seller common.Address
	Price  *big.Int
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xe411526a25ed63765069eba7a21c0a07d38b4487b243f295662878c7437da6c1.
//
// Solidity: event Update(address indexed token, address indexed seller, uint256 price, uint256 amount, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) FilterUpdate(opts *bind.FilterOpts, token []common.Address, seller []common.Address) (*L2TokenMarketUpdateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.FilterLogs(opts, "Update", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &L2TokenMarketUpdateIterator{contract: _L2TokenMarket.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xe411526a25ed63765069eba7a21c0a07d38b4487b243f295662878c7437da6c1.
//
// Solidity: event Update(address indexed token, address indexed seller, uint256 price, uint256 amount, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *L2TokenMarketUpdate, token []common.Address, seller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _L2TokenMarket.contract.WatchLogs(opts, "Update", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TokenMarketUpdate)
				if err := _L2TokenMarket.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0xe411526a25ed63765069eba7a21c0a07d38b4487b243f295662878c7437da6c1.
//
// Solidity: event Update(address indexed token, address indexed seller, uint256 price, uint256 amount, uint256 nonce)
func (_L2TokenMarket *L2TokenMarketFilterer) ParseUpdate(log types.Log) (*L2TokenMarketUpdate, error) {
	event := new(L2TokenMarketUpdate)
	if err := _L2TokenMarket.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
