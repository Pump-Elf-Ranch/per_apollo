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

// PerTokenBaseMetaData contains all meta data concerning the PerTokenBase contract.
var PerTokenBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PerFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buy_item\",\"inputs\":[{\"name\":\"pid1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pid2\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ranchId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"erc20PayAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"itemContractAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint_nft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nftAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nft_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setErc20PayAddress\",\"inputs\":[{\"name\":\"_erc20PayAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFreelist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isFreelist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setItemContract\",\"inputs\":[{\"name\":\"_itemContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNftAddress\",\"inputs\":[{\"name\":\"_nftAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublicPrice\",\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelist\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistPrice\",\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ItemBought\",\"inputs\":[{\"name\":\"pid1\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pid2\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"itemId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ranchId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"itemType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ItemMinted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"mintType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// PerTokenBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PerTokenBaseMetaData.ABI instead.
var PerTokenBaseABI = PerTokenBaseMetaData.ABI

// PerTokenBase is an auto generated Go binding around an Ethereum contract.
type PerTokenBase struct {
	PerTokenBaseCaller     // Read-only binding to the contract
	PerTokenBaseTransactor // Write-only binding to the contract
	PerTokenBaseFilterer   // Log filterer for contract events
}

// PerTokenBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerTokenBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerTokenBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerTokenBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerTokenBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerTokenBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerTokenBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerTokenBaseSession struct {
	Contract     *PerTokenBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerTokenBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerTokenBaseCallerSession struct {
	Contract *PerTokenBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PerTokenBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerTokenBaseTransactorSession struct {
	Contract     *PerTokenBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PerTokenBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerTokenBaseRaw struct {
	Contract *PerTokenBase // Generic contract binding to access the raw methods on
}

// PerTokenBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerTokenBaseCallerRaw struct {
	Contract *PerTokenBaseCaller // Generic read-only contract binding to access the raw methods on
}

// PerTokenBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerTokenBaseTransactorRaw struct {
	Contract *PerTokenBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerTokenBase creates a new instance of PerTokenBase, bound to a specific deployed contract.
func NewPerTokenBase(address common.Address, backend bind.ContractBackend) (*PerTokenBase, error) {
	contract, err := bindPerTokenBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerTokenBase{PerTokenBaseCaller: PerTokenBaseCaller{contract: contract}, PerTokenBaseTransactor: PerTokenBaseTransactor{contract: contract}, PerTokenBaseFilterer: PerTokenBaseFilterer{contract: contract}}, nil
}

// NewPerTokenBaseCaller creates a new read-only instance of PerTokenBase, bound to a specific deployed contract.
func NewPerTokenBaseCaller(address common.Address, caller bind.ContractCaller) (*PerTokenBaseCaller, error) {
	contract, err := bindPerTokenBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseCaller{contract: contract}, nil
}

// NewPerTokenBaseTransactor creates a new write-only instance of PerTokenBase, bound to a specific deployed contract.
func NewPerTokenBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PerTokenBaseTransactor, error) {
	contract, err := bindPerTokenBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseTransactor{contract: contract}, nil
}

// NewPerTokenBaseFilterer creates a new log filterer instance of PerTokenBase, bound to a specific deployed contract.
func NewPerTokenBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PerTokenBaseFilterer, error) {
	contract, err := bindPerTokenBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseFilterer{contract: contract}, nil
}

// bindPerTokenBase binds a generic wrapper to an already deployed contract.
func bindPerTokenBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PerTokenBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerTokenBase *PerTokenBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerTokenBase.Contract.PerTokenBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerTokenBase *PerTokenBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerTokenBase.Contract.PerTokenBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerTokenBase *PerTokenBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerTokenBase.Contract.PerTokenBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerTokenBase *PerTokenBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerTokenBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerTokenBase *PerTokenBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerTokenBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerTokenBase *PerTokenBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerTokenBase.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PerTokenBase *PerTokenBaseCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PerTokenBase *PerTokenBaseSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PerTokenBase.Contract.DEFAULTADMINROLE(&_PerTokenBase.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PerTokenBase *PerTokenBaseCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PerTokenBase.Contract.DEFAULTADMINROLE(&_PerTokenBase.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "Nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) Nonce() (*big.Int, error) {
	return _PerTokenBase.Contract.Nonce(&_PerTokenBase.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xf8e08389.
//
// Solidity: function Nonce() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) Nonce() (*big.Int, error) {
	return _PerTokenBase.Contract.Nonce(&_PerTokenBase.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) PerFee() (*big.Int, error) {
	return _PerTokenBase.Contract.PerFee(&_PerTokenBase.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) PerFee() (*big.Int, error) {
	return _PerTokenBase.Contract.PerFee(&_PerTokenBase.CallOpts)
}

// Erc20PayAddress is a free data retrieval call binding the contract method 0x66b287fc.
//
// Solidity: function erc20PayAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCaller) Erc20PayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "erc20PayAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20PayAddress is a free data retrieval call binding the contract method 0x66b287fc.
//
// Solidity: function erc20PayAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseSession) Erc20PayAddress() (common.Address, error) {
	return _PerTokenBase.Contract.Erc20PayAddress(&_PerTokenBase.CallOpts)
}

// Erc20PayAddress is a free data retrieval call binding the contract method 0x66b287fc.
//
// Solidity: function erc20PayAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCallerSession) Erc20PayAddress() (common.Address, error) {
	return _PerTokenBase.Contract.Erc20PayAddress(&_PerTokenBase.CallOpts)
}

// Freelist is a free data retrieval call binding the contract method 0x7227548b.
//
// Solidity: function freelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseCaller) Freelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "freelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Freelist is a free data retrieval call binding the contract method 0x7227548b.
//
// Solidity: function freelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseSession) Freelist(arg0 common.Address) (bool, error) {
	return _PerTokenBase.Contract.Freelist(&_PerTokenBase.CallOpts, arg0)
}

// Freelist is a free data retrieval call binding the contract method 0x7227548b.
//
// Solidity: function freelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseCallerSession) Freelist(arg0 common.Address) (bool, error) {
	return _PerTokenBase.Contract.Freelist(&_PerTokenBase.CallOpts, arg0)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) GetFee() (*big.Int, error) {
	return _PerTokenBase.Contract.GetFee(&_PerTokenBase.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) GetFee() (*big.Int, error) {
	return _PerTokenBase.Contract.GetFee(&_PerTokenBase.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PerTokenBase *PerTokenBaseCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PerTokenBase *PerTokenBaseSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PerTokenBase.Contract.GetRoleAdmin(&_PerTokenBase.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PerTokenBase *PerTokenBaseCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PerTokenBase.Contract.GetRoleAdmin(&_PerTokenBase.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PerTokenBase *PerTokenBaseCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PerTokenBase *PerTokenBaseSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PerTokenBase.Contract.HasRole(&_PerTokenBase.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PerTokenBase *PerTokenBaseCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PerTokenBase.Contract.HasRole(&_PerTokenBase.CallOpts, role, account)
}

// ItemContractAddress is a free data retrieval call binding the contract method 0x3089f68f.
//
// Solidity: function itemContractAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCaller) ItemContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "itemContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ItemContractAddress is a free data retrieval call binding the contract method 0x3089f68f.
//
// Solidity: function itemContractAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseSession) ItemContractAddress() (common.Address, error) {
	return _PerTokenBase.Contract.ItemContractAddress(&_PerTokenBase.CallOpts)
}

// ItemContractAddress is a free data retrieval call binding the contract method 0x3089f68f.
//
// Solidity: function itemContractAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCallerSession) ItemContractAddress() (common.Address, error) {
	return _PerTokenBase.Contract.ItemContractAddress(&_PerTokenBase.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCaller) NftAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "nftAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseSession) NftAddress() (common.Address, error) {
	return _PerTokenBase.Contract.NftAddress(&_PerTokenBase.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_PerTokenBase *PerTokenBaseCallerSession) NftAddress() (common.Address, error) {
	return _PerTokenBase.Contract.NftAddress(&_PerTokenBase.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) NftId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "nftId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) NftId() (*big.Int, error) {
	return _PerTokenBase.Contract.NftId(&_PerTokenBase.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) NftId() (*big.Int, error) {
	return _PerTokenBase.Contract.NftId(&_PerTokenBase.CallOpts)
}

// NftCount is a free data retrieval call binding the contract method 0x73b0745f.
//
// Solidity: function nft_count() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) NftCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "nft_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftCount is a free data retrieval call binding the contract method 0x73b0745f.
//
// Solidity: function nft_count() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) NftCount() (*big.Int, error) {
	return _PerTokenBase.Contract.NftCount(&_PerTokenBase.CallOpts)
}

// NftCount is a free data retrieval call binding the contract method 0x73b0745f.
//
// Solidity: function nft_count() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) NftCount() (*big.Int, error) {
	return _PerTokenBase.Contract.NftCount(&_PerTokenBase.CallOpts)
}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) PublicPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "publicPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) PublicPrice() (*big.Int, error) {
	return _PerTokenBase.Contract.PublicPrice(&_PerTokenBase.CallOpts)
}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) PublicPrice() (*big.Int, error) {
	return _PerTokenBase.Contract.PublicPrice(&_PerTokenBase.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerTokenBase *PerTokenBaseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerTokenBase *PerTokenBaseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerTokenBase.Contract.SupportsInterface(&_PerTokenBase.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PerTokenBase *PerTokenBaseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PerTokenBase.Contract.SupportsInterface(&_PerTokenBase.CallOpts, interfaceId)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseSession) Whitelist(arg0 common.Address) (bool, error) {
	return _PerTokenBase.Contract.Whitelist(&_PerTokenBase.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_PerTokenBase *PerTokenBaseCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _PerTokenBase.Contract.Whitelist(&_PerTokenBase.CallOpts, arg0)
}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCaller) WhitelistPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerTokenBase.contract.Call(opts, &out, "whitelistPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseSession) WhitelistPrice() (*big.Int, error) {
	return _PerTokenBase.Contract.WhitelistPrice(&_PerTokenBase.CallOpts)
}

// WhitelistPrice is a free data retrieval call binding the contract method 0xfc1a1c36.
//
// Solidity: function whitelistPrice() view returns(uint256)
func (_PerTokenBase *PerTokenBaseCallerSession) WhitelistPrice() (*big.Int, error) {
	return _PerTokenBase.Contract.WhitelistPrice(&_PerTokenBase.CallOpts)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d362be1.
//
// Solidity: function buy_item(uint256 pid1, uint256 pid2, uint256 itemId, uint256 ranchId) returns()
func (_PerTokenBase *PerTokenBaseTransactor) BuyItem(opts *bind.TransactOpts, pid1 *big.Int, pid2 *big.Int, itemId *big.Int, ranchId *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "buy_item", pid1, pid2, itemId, ranchId)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d362be1.
//
// Solidity: function buy_item(uint256 pid1, uint256 pid2, uint256 itemId, uint256 ranchId) returns()
func (_PerTokenBase *PerTokenBaseSession) BuyItem(pid1 *big.Int, pid2 *big.Int, itemId *big.Int, ranchId *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.BuyItem(&_PerTokenBase.TransactOpts, pid1, pid2, itemId, ranchId)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d362be1.
//
// Solidity: function buy_item(uint256 pid1, uint256 pid2, uint256 itemId, uint256 ranchId) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) BuyItem(pid1 *big.Int, pid2 *big.Int, itemId *big.Int, ranchId *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.BuyItem(&_PerTokenBase.TransactOpts, pid1, pid2, itemId, ranchId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.GrantRole(&_PerTokenBase.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.GrantRole(&_PerTokenBase.TransactOpts, role, account)
}

// MintNft is a paid mutator transaction binding the contract method 0xff4cb021.
//
// Solidity: function mint_nft() returns(bool)
func (_PerTokenBase *PerTokenBaseTransactor) MintNft(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "mint_nft")
}

// MintNft is a paid mutator transaction binding the contract method 0xff4cb021.
//
// Solidity: function mint_nft() returns(bool)
func (_PerTokenBase *PerTokenBaseSession) MintNft() (*types.Transaction, error) {
	return _PerTokenBase.Contract.MintNft(&_PerTokenBase.TransactOpts)
}

// MintNft is a paid mutator transaction binding the contract method 0xff4cb021.
//
// Solidity: function mint_nft() returns(bool)
func (_PerTokenBase *PerTokenBaseTransactorSession) MintNft() (*types.Transaction, error) {
	return _PerTokenBase.Contract.MintNft(&_PerTokenBase.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PerTokenBase *PerTokenBaseTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PerTokenBase *PerTokenBaseSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.RenounceRole(&_PerTokenBase.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.RenounceRole(&_PerTokenBase.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.RevokeRole(&_PerTokenBase.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.RevokeRole(&_PerTokenBase.TransactOpts, role, account)
}

// SetErc20PayAddress is a paid mutator transaction binding the contract method 0xdebd0378.
//
// Solidity: function setErc20PayAddress(address _erc20PayAddress) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetErc20PayAddress(opts *bind.TransactOpts, _erc20PayAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setErc20PayAddress", _erc20PayAddress)
}

// SetErc20PayAddress is a paid mutator transaction binding the contract method 0xdebd0378.
//
// Solidity: function setErc20PayAddress(address _erc20PayAddress) returns()
func (_PerTokenBase *PerTokenBaseSession) SetErc20PayAddress(_erc20PayAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetErc20PayAddress(&_PerTokenBase.TransactOpts, _erc20PayAddress)
}

// SetErc20PayAddress is a paid mutator transaction binding the contract method 0xdebd0378.
//
// Solidity: function setErc20PayAddress(address _erc20PayAddress) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetErc20PayAddress(_erc20PayAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetErc20PayAddress(&_PerTokenBase.TransactOpts, _erc20PayAddress)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_PerTokenBase *PerTokenBaseTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_PerTokenBase *PerTokenBaseSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetFee(&_PerTokenBase.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns(bool)
func (_PerTokenBase *PerTokenBaseTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetFee(&_PerTokenBase.TransactOpts, _fee)
}

// SetFreelist is a paid mutator transaction binding the contract method 0x5411d168.
//
// Solidity: function setFreelist(address _address, bool _isFreelist) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetFreelist(opts *bind.TransactOpts, _address common.Address, _isFreelist bool) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setFreelist", _address, _isFreelist)
}

// SetFreelist is a paid mutator transaction binding the contract method 0x5411d168.
//
// Solidity: function setFreelist(address _address, bool _isFreelist) returns()
func (_PerTokenBase *PerTokenBaseSession) SetFreelist(_address common.Address, _isFreelist bool) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetFreelist(&_PerTokenBase.TransactOpts, _address, _isFreelist)
}

// SetFreelist is a paid mutator transaction binding the contract method 0x5411d168.
//
// Solidity: function setFreelist(address _address, bool _isFreelist) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetFreelist(_address common.Address, _isFreelist bool) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetFreelist(&_PerTokenBase.TransactOpts, _address, _isFreelist)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address _itemContract) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetItemContract(opts *bind.TransactOpts, _itemContract common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setItemContract", _itemContract)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address _itemContract) returns()
func (_PerTokenBase *PerTokenBaseSession) SetItemContract(_itemContract common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetItemContract(&_PerTokenBase.TransactOpts, _itemContract)
}

// SetItemContract is a paid mutator transaction binding the contract method 0xa7120433.
//
// Solidity: function setItemContract(address _itemContract) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetItemContract(_itemContract common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetItemContract(&_PerTokenBase.TransactOpts, _itemContract)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _nftAddress) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetNftAddress(opts *bind.TransactOpts, _nftAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setNftAddress", _nftAddress)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _nftAddress) returns()
func (_PerTokenBase *PerTokenBaseSession) SetNftAddress(_nftAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetNftAddress(&_PerTokenBase.TransactOpts, _nftAddress)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _nftAddress) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetNftAddress(_nftAddress common.Address) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetNftAddress(&_PerTokenBase.TransactOpts, _nftAddress)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetPublicPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setPublicPrice", _price)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseSession) SetPublicPrice(_price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetPublicPrice(&_PerTokenBase.TransactOpts, _price)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetPublicPrice(_price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetPublicPrice(&_PerTokenBase.TransactOpts, _price)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _address, bool _isWhitelisted) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetWhitelist(opts *bind.TransactOpts, _address common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setWhitelist", _address, _isWhitelisted)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _address, bool _isWhitelisted) returns()
func (_PerTokenBase *PerTokenBaseSession) SetWhitelist(_address common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetWhitelist(&_PerTokenBase.TransactOpts, _address, _isWhitelisted)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x53d6fd59.
//
// Solidity: function setWhitelist(address _address, bool _isWhitelisted) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetWhitelist(_address common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetWhitelist(&_PerTokenBase.TransactOpts, _address, _isWhitelisted)
}

// SetWhitelistPrice is a paid mutator transaction binding the contract method 0x717d57d3.
//
// Solidity: function setWhitelistPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseTransactor) SetWhitelistPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "setWhitelistPrice", _price)
}

// SetWhitelistPrice is a paid mutator transaction binding the contract method 0x717d57d3.
//
// Solidity: function setWhitelistPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseSession) SetWhitelistPrice(_price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetWhitelistPrice(&_PerTokenBase.TransactOpts, _price)
}

// SetWhitelistPrice is a paid mutator transaction binding the contract method 0x717d57d3.
//
// Solidity: function setWhitelistPrice(uint256 _price) returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) SetWhitelistPrice(_price *big.Int) (*types.Transaction, error) {
	return _PerTokenBase.Contract.SetWhitelistPrice(&_PerTokenBase.TransactOpts, _price)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PerTokenBase *PerTokenBaseTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PerTokenBase *PerTokenBaseSession) Withdraw() (*types.Transaction, error) {
	return _PerTokenBase.Contract.Withdraw(&_PerTokenBase.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PerTokenBase.Contract.Withdraw(&_PerTokenBase.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_PerTokenBase *PerTokenBaseTransactor) WithdrawFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerTokenBase.contract.Transact(opts, "withdrawFee")
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_PerTokenBase *PerTokenBaseSession) WithdrawFee() (*types.Transaction, error) {
	return _PerTokenBase.Contract.WithdrawFee(&_PerTokenBase.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_PerTokenBase *PerTokenBaseTransactorSession) WithdrawFee() (*types.Transaction, error) {
	return _PerTokenBase.Contract.WithdrawFee(&_PerTokenBase.TransactOpts)
}

// PerTokenBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PerTokenBase contract.
type PerTokenBaseInitializedIterator struct {
	Event *PerTokenBaseInitialized // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseInitialized)
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
		it.Event = new(PerTokenBaseInitialized)
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
func (it *PerTokenBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseInitialized represents a Initialized event raised by the PerTokenBase contract.
type PerTokenBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PerTokenBase *PerTokenBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*PerTokenBaseInitializedIterator, error) {

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseInitializedIterator{contract: _PerTokenBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PerTokenBase *PerTokenBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PerTokenBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseInitialized)
				if err := _PerTokenBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PerTokenBase *PerTokenBaseFilterer) ParseInitialized(log types.Log) (*PerTokenBaseInitialized, error) {
	event := new(PerTokenBaseInitialized)
	if err := _PerTokenBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerTokenBaseItemBoughtIterator is returned from FilterItemBought and is used to iterate over the raw logs and unpacked data for ItemBought events raised by the PerTokenBase contract.
type PerTokenBaseItemBoughtIterator struct {
	Event *PerTokenBaseItemBought // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseItemBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseItemBought)
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
		it.Event = new(PerTokenBaseItemBought)
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
func (it *PerTokenBaseItemBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseItemBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseItemBought represents a ItemBought event raised by the PerTokenBase contract.
type PerTokenBaseItemBought struct {
	Pid1     *big.Int
	Pid2     *big.Int
	ItemId   *big.Int
	Price    *big.Int
	RanchId  *big.Int
	Buyer    common.Address
	ItemType string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemBought is a free log retrieval operation binding the contract event 0x359cb519da2e9b86aee5ca70a3c570a43d8640b883ff3fcd758dbca2ab2e3bfd.
//
// Solidity: event ItemBought(uint256 indexed pid1, uint256 indexed pid2, uint256 indexed itemId, uint256 price, uint256 ranchId, address buyer, string itemType)
func (_PerTokenBase *PerTokenBaseFilterer) FilterItemBought(opts *bind.FilterOpts, pid1 []*big.Int, pid2 []*big.Int, itemId []*big.Int) (*PerTokenBaseItemBoughtIterator, error) {

	var pid1Rule []interface{}
	for _, pid1Item := range pid1 {
		pid1Rule = append(pid1Rule, pid1Item)
	}
	var pid2Rule []interface{}
	for _, pid2Item := range pid2 {
		pid2Rule = append(pid2Rule, pid2Item)
	}
	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "ItemBought", pid1Rule, pid2Rule, itemIdRule)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseItemBoughtIterator{contract: _PerTokenBase.contract, event: "ItemBought", logs: logs, sub: sub}, nil
}

// WatchItemBought is a free log subscription operation binding the contract event 0x359cb519da2e9b86aee5ca70a3c570a43d8640b883ff3fcd758dbca2ab2e3bfd.
//
// Solidity: event ItemBought(uint256 indexed pid1, uint256 indexed pid2, uint256 indexed itemId, uint256 price, uint256 ranchId, address buyer, string itemType)
func (_PerTokenBase *PerTokenBaseFilterer) WatchItemBought(opts *bind.WatchOpts, sink chan<- *PerTokenBaseItemBought, pid1 []*big.Int, pid2 []*big.Int, itemId []*big.Int) (event.Subscription, error) {

	var pid1Rule []interface{}
	for _, pid1Item := range pid1 {
		pid1Rule = append(pid1Rule, pid1Item)
	}
	var pid2Rule []interface{}
	for _, pid2Item := range pid2 {
		pid2Rule = append(pid2Rule, pid2Item)
	}
	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "ItemBought", pid1Rule, pid2Rule, itemIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseItemBought)
				if err := _PerTokenBase.contract.UnpackLog(event, "ItemBought", log); err != nil {
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

// ParseItemBought is a log parse operation binding the contract event 0x359cb519da2e9b86aee5ca70a3c570a43d8640b883ff3fcd758dbca2ab2e3bfd.
//
// Solidity: event ItemBought(uint256 indexed pid1, uint256 indexed pid2, uint256 indexed itemId, uint256 price, uint256 ranchId, address buyer, string itemType)
func (_PerTokenBase *PerTokenBaseFilterer) ParseItemBought(log types.Log) (*PerTokenBaseItemBought, error) {
	event := new(PerTokenBaseItemBought)
	if err := _PerTokenBase.contract.UnpackLog(event, "ItemBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerTokenBaseItemMintedIterator is returned from FilterItemMinted and is used to iterate over the raw logs and unpacked data for ItemMinted events raised by the PerTokenBase contract.
type PerTokenBaseItemMintedIterator struct {
	Event *PerTokenBaseItemMinted // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseItemMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseItemMinted)
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
		it.Event = new(PerTokenBaseItemMinted)
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
func (it *PerTokenBaseItemMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseItemMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseItemMinted represents a ItemMinted event raised by the PerTokenBase contract.
type PerTokenBaseItemMinted struct {
	To       common.Address
	Nonce    *big.Int
	MintType string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemMinted is a free log retrieval operation binding the contract event 0x9aad673c07866c8b5f7049a3b0a7e8091c80c3685b33244faf2bb621f050fdd7.
//
// Solidity: event ItemMinted(address indexed to, uint256 indexed nonce, string mintType)
func (_PerTokenBase *PerTokenBaseFilterer) FilterItemMinted(opts *bind.FilterOpts, to []common.Address, nonce []*big.Int) (*PerTokenBaseItemMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "ItemMinted", toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseItemMintedIterator{contract: _PerTokenBase.contract, event: "ItemMinted", logs: logs, sub: sub}, nil
}

// WatchItemMinted is a free log subscription operation binding the contract event 0x9aad673c07866c8b5f7049a3b0a7e8091c80c3685b33244faf2bb621f050fdd7.
//
// Solidity: event ItemMinted(address indexed to, uint256 indexed nonce, string mintType)
func (_PerTokenBase *PerTokenBaseFilterer) WatchItemMinted(opts *bind.WatchOpts, sink chan<- *PerTokenBaseItemMinted, to []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "ItemMinted", toRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseItemMinted)
				if err := _PerTokenBase.contract.UnpackLog(event, "ItemMinted", log); err != nil {
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

// ParseItemMinted is a log parse operation binding the contract event 0x9aad673c07866c8b5f7049a3b0a7e8091c80c3685b33244faf2bb621f050fdd7.
//
// Solidity: event ItemMinted(address indexed to, uint256 indexed nonce, string mintType)
func (_PerTokenBase *PerTokenBaseFilterer) ParseItemMinted(log types.Log) (*PerTokenBaseItemMinted, error) {
	event := new(PerTokenBaseItemMinted)
	if err := _PerTokenBase.contract.UnpackLog(event, "ItemMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerTokenBaseRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PerTokenBase contract.
type PerTokenBaseRoleAdminChangedIterator struct {
	Event *PerTokenBaseRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseRoleAdminChanged)
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
		it.Event = new(PerTokenBaseRoleAdminChanged)
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
func (it *PerTokenBaseRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseRoleAdminChanged represents a RoleAdminChanged event raised by the PerTokenBase contract.
type PerTokenBaseRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PerTokenBase *PerTokenBaseFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PerTokenBaseRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseRoleAdminChangedIterator{contract: _PerTokenBase.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PerTokenBase *PerTokenBaseFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PerTokenBaseRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseRoleAdminChanged)
				if err := _PerTokenBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PerTokenBase *PerTokenBaseFilterer) ParseRoleAdminChanged(log types.Log) (*PerTokenBaseRoleAdminChanged, error) {
	event := new(PerTokenBaseRoleAdminChanged)
	if err := _PerTokenBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerTokenBaseRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PerTokenBase contract.
type PerTokenBaseRoleGrantedIterator struct {
	Event *PerTokenBaseRoleGranted // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseRoleGranted)
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
		it.Event = new(PerTokenBaseRoleGranted)
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
func (it *PerTokenBaseRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseRoleGranted represents a RoleGranted event raised by the PerTokenBase contract.
type PerTokenBaseRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PerTokenBase *PerTokenBaseFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PerTokenBaseRoleGrantedIterator, error) {

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

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseRoleGrantedIterator{contract: _PerTokenBase.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PerTokenBase *PerTokenBaseFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PerTokenBaseRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseRoleGranted)
				if err := _PerTokenBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PerTokenBase *PerTokenBaseFilterer) ParseRoleGranted(log types.Log) (*PerTokenBaseRoleGranted, error) {
	event := new(PerTokenBaseRoleGranted)
	if err := _PerTokenBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerTokenBaseRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PerTokenBase contract.
type PerTokenBaseRoleRevokedIterator struct {
	Event *PerTokenBaseRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PerTokenBaseRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerTokenBaseRoleRevoked)
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
		it.Event = new(PerTokenBaseRoleRevoked)
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
func (it *PerTokenBaseRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerTokenBaseRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerTokenBaseRoleRevoked represents a RoleRevoked event raised by the PerTokenBase contract.
type PerTokenBaseRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PerTokenBase *PerTokenBaseFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PerTokenBaseRoleRevokedIterator, error) {

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

	logs, sub, err := _PerTokenBase.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PerTokenBaseRoleRevokedIterator{contract: _PerTokenBase.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PerTokenBase *PerTokenBaseFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PerTokenBaseRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PerTokenBase.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerTokenBaseRoleRevoked)
				if err := _PerTokenBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PerTokenBase *PerTokenBaseFilterer) ParseRoleRevoked(log types.Log) (*PerTokenBaseRoleRevoked, error) {
	event := new(PerTokenBaseRoleRevoked)
	if err := _PerTokenBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
