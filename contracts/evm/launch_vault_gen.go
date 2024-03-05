// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

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

// LaunchVaultMetaData contains all meta data concerning the LaunchVault contract.
var LaunchVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDeposit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_DEPLOYMENT_ERR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ASSET_REMOVAL_ERR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_LIQUIDITY_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNAUTHORIZED_ERR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNREGISTERED_ASSET_ERR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_TOKEN_GRANULARITY_ERR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deployAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"disableLPTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"enableLPTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractBlackwingVaultToken\",\"name\":\"vaultToken\",\"type\":\"address\"},{\"internalType\":\"contractIDeployer\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"registerAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractIDeployer\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"updateDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vaultTokensToBurn\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LaunchVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use LaunchVaultMetaData.ABI instead.
var LaunchVaultABI = LaunchVaultMetaData.ABI

// LaunchVault is an auto generated Go binding around an Ethereum contract.
type LaunchVault struct {
	LaunchVaultCaller     // Read-only binding to the contract
	LaunchVaultTransactor // Write-only binding to the contract
	LaunchVaultFilterer   // Log filterer for contract events
}

// LaunchVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type LaunchVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LaunchVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LaunchVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LaunchVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LaunchVaultSession struct {
	Contract     *LaunchVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LaunchVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LaunchVaultCallerSession struct {
	Contract *LaunchVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LaunchVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LaunchVaultTransactorSession struct {
	Contract     *LaunchVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LaunchVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type LaunchVaultRaw struct {
	Contract *LaunchVault // Generic contract binding to access the raw methods on
}

// LaunchVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LaunchVaultCallerRaw struct {
	Contract *LaunchVaultCaller // Generic read-only contract binding to access the raw methods on
}

// LaunchVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LaunchVaultTransactorRaw struct {
	Contract *LaunchVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLaunchVault creates a new instance of LaunchVault, bound to a specific deployed contract.
func NewLaunchVault(address common.Address, backend bind.ContractBackend) (*LaunchVault, error) {
	contract, err := bindLaunchVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LaunchVault{LaunchVaultCaller: LaunchVaultCaller{contract: contract}, LaunchVaultTransactor: LaunchVaultTransactor{contract: contract}, LaunchVaultFilterer: LaunchVaultFilterer{contract: contract}}, nil
}

// NewLaunchVaultCaller creates a new read-only instance of LaunchVault, bound to a specific deployed contract.
func NewLaunchVaultCaller(address common.Address, caller bind.ContractCaller) (*LaunchVaultCaller, error) {
	contract, err := bindLaunchVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultCaller{contract: contract}, nil
}

// NewLaunchVaultTransactor creates a new write-only instance of LaunchVault, bound to a specific deployed contract.
func NewLaunchVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*LaunchVaultTransactor, error) {
	contract, err := bindLaunchVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultTransactor{contract: contract}, nil
}

// NewLaunchVaultFilterer creates a new log filterer instance of LaunchVault, bound to a specific deployed contract.
func NewLaunchVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*LaunchVaultFilterer, error) {
	contract, err := bindLaunchVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultFilterer{contract: contract}, nil
}

// bindLaunchVault binds a generic wrapper to an already deployed contract.
func bindLaunchVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LaunchVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaunchVault *LaunchVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaunchVault.Contract.LaunchVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaunchVault *LaunchVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaunchVault.Contract.LaunchVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaunchVault *LaunchVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaunchVault.Contract.LaunchVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LaunchVault *LaunchVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LaunchVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LaunchVault *LaunchVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaunchVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LaunchVault *LaunchVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LaunchVault.Contract.contract.Transact(opts, method, params...)
}

// ASSETDEPLOYMENTERR is a free data retrieval call binding the contract method 0xe9af7c24.
//
// Solidity: function ASSET_DEPLOYMENT_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCaller) ASSETDEPLOYMENTERR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "ASSET_DEPLOYMENT_ERR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ASSETDEPLOYMENTERR is a free data retrieval call binding the contract method 0xe9af7c24.
//
// Solidity: function ASSET_DEPLOYMENT_ERR() view returns(string)
func (_LaunchVault *LaunchVaultSession) ASSETDEPLOYMENTERR() (string, error) {
	return _LaunchVault.Contract.ASSETDEPLOYMENTERR(&_LaunchVault.CallOpts)
}

// ASSETDEPLOYMENTERR is a free data retrieval call binding the contract method 0xe9af7c24.
//
// Solidity: function ASSET_DEPLOYMENT_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCallerSession) ASSETDEPLOYMENTERR() (string, error) {
	return _LaunchVault.Contract.ASSETDEPLOYMENTERR(&_LaunchVault.CallOpts)
}

// ASSETREMOVALERR is a free data retrieval call binding the contract method 0xd80be039.
//
// Solidity: function ASSET_REMOVAL_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCaller) ASSETREMOVALERR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "ASSET_REMOVAL_ERR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ASSETREMOVALERR is a free data retrieval call binding the contract method 0xd80be039.
//
// Solidity: function ASSET_REMOVAL_ERR() view returns(string)
func (_LaunchVault *LaunchVaultSession) ASSETREMOVALERR() (string, error) {
	return _LaunchVault.Contract.ASSETREMOVALERR(&_LaunchVault.CallOpts)
}

// ASSETREMOVALERR is a free data retrieval call binding the contract method 0xd80be039.
//
// Solidity: function ASSET_REMOVAL_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCallerSession) ASSETREMOVALERR() (string, error) {
	return _LaunchVault.Contract.ASSETREMOVALERR(&_LaunchVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LaunchVault.Contract.DEFAULTADMINROLE(&_LaunchVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LaunchVault.Contract.DEFAULTADMINROLE(&_LaunchVault.CallOpts)
}

// INITIALLIQUIDITYMULTIPLIER is a free data retrieval call binding the contract method 0x5e09e7af.
//
// Solidity: function INITIAL_LIQUIDITY_MULTIPLIER() view returns(uint256)
func (_LaunchVault *LaunchVaultCaller) INITIALLIQUIDITYMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "INITIAL_LIQUIDITY_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALLIQUIDITYMULTIPLIER is a free data retrieval call binding the contract method 0x5e09e7af.
//
// Solidity: function INITIAL_LIQUIDITY_MULTIPLIER() view returns(uint256)
func (_LaunchVault *LaunchVaultSession) INITIALLIQUIDITYMULTIPLIER() (*big.Int, error) {
	return _LaunchVault.Contract.INITIALLIQUIDITYMULTIPLIER(&_LaunchVault.CallOpts)
}

// INITIALLIQUIDITYMULTIPLIER is a free data retrieval call binding the contract method 0x5e09e7af.
//
// Solidity: function INITIAL_LIQUIDITY_MULTIPLIER() view returns(uint256)
func (_LaunchVault *LaunchVaultCallerSession) INITIALLIQUIDITYMULTIPLIER() (*big.Int, error) {
	return _LaunchVault.Contract.INITIALLIQUIDITYMULTIPLIER(&_LaunchVault.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultSession) OWNERROLE() ([32]byte, error) {
	return _LaunchVault.Contract.OWNERROLE(&_LaunchVault.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_LaunchVault *LaunchVaultCallerSession) OWNERROLE() ([32]byte, error) {
	return _LaunchVault.Contract.OWNERROLE(&_LaunchVault.CallOpts)
}

// UNAUTHORIZEDERR is a free data retrieval call binding the contract method 0xaf4f19a6.
//
// Solidity: function UNAUTHORIZED_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCaller) UNAUTHORIZEDERR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "UNAUTHORIZED_ERR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UNAUTHORIZEDERR is a free data retrieval call binding the contract method 0xaf4f19a6.
//
// Solidity: function UNAUTHORIZED_ERR() view returns(string)
func (_LaunchVault *LaunchVaultSession) UNAUTHORIZEDERR() (string, error) {
	return _LaunchVault.Contract.UNAUTHORIZEDERR(&_LaunchVault.CallOpts)
}

// UNAUTHORIZEDERR is a free data retrieval call binding the contract method 0xaf4f19a6.
//
// Solidity: function UNAUTHORIZED_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCallerSession) UNAUTHORIZEDERR() (string, error) {
	return _LaunchVault.Contract.UNAUTHORIZEDERR(&_LaunchVault.CallOpts)
}

// UNREGISTEREDASSETERR is a free data retrieval call binding the contract method 0xdf11653d.
//
// Solidity: function UNREGISTERED_ASSET_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCaller) UNREGISTEREDASSETERR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "UNREGISTERED_ASSET_ERR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UNREGISTEREDASSETERR is a free data retrieval call binding the contract method 0xdf11653d.
//
// Solidity: function UNREGISTERED_ASSET_ERR() view returns(string)
func (_LaunchVault *LaunchVaultSession) UNREGISTEREDASSETERR() (string, error) {
	return _LaunchVault.Contract.UNREGISTEREDASSETERR(&_LaunchVault.CallOpts)
}

// UNREGISTEREDASSETERR is a free data retrieval call binding the contract method 0xdf11653d.
//
// Solidity: function UNREGISTERED_ASSET_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCallerSession) UNREGISTEREDASSETERR() (string, error) {
	return _LaunchVault.Contract.UNREGISTEREDASSETERR(&_LaunchVault.CallOpts)
}

// VAULTTOKENGRANULARITYERR is a free data retrieval call binding the contract method 0xc74b7de1.
//
// Solidity: function VAULT_TOKEN_GRANULARITY_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCaller) VAULTTOKENGRANULARITYERR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "VAULT_TOKEN_GRANULARITY_ERR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VAULTTOKENGRANULARITYERR is a free data retrieval call binding the contract method 0xc74b7de1.
//
// Solidity: function VAULT_TOKEN_GRANULARITY_ERR() view returns(string)
func (_LaunchVault *LaunchVaultSession) VAULTTOKENGRANULARITYERR() (string, error) {
	return _LaunchVault.Contract.VAULTTOKENGRANULARITYERR(&_LaunchVault.CallOpts)
}

// VAULTTOKENGRANULARITYERR is a free data retrieval call binding the contract method 0xc74b7de1.
//
// Solidity: function VAULT_TOKEN_GRANULARITY_ERR() view returns(string)
func (_LaunchVault *LaunchVaultCallerSession) VAULTTOKENGRANULARITYERR() (string, error) {
	return _LaunchVault.Contract.VAULTTOKENGRANULARITYERR(&_LaunchVault.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address asset, address user) view returns(uint256)
func (_LaunchVault *LaunchVaultCaller) Balance(opts *bind.CallOpts, asset common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "balance", asset, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address asset, address user) view returns(uint256)
func (_LaunchVault *LaunchVaultSession) Balance(asset common.Address, user common.Address) (*big.Int, error) {
	return _LaunchVault.Contract.Balance(&_LaunchVault.CallOpts, asset, user)
}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address asset, address user) view returns(uint256)
func (_LaunchVault *LaunchVaultCallerSession) Balance(asset common.Address, user common.Address) (*big.Int, error) {
	return _LaunchVault.Contract.Balance(&_LaunchVault.CallOpts, asset, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LaunchVault *LaunchVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LaunchVault *LaunchVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LaunchVault.Contract.GetRoleAdmin(&_LaunchVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LaunchVault *LaunchVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LaunchVault.Contract.GetRoleAdmin(&_LaunchVault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LaunchVault *LaunchVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LaunchVault *LaunchVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LaunchVault.Contract.HasRole(&_LaunchVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LaunchVault *LaunchVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LaunchVault.Contract.HasRole(&_LaunchVault.CallOpts, role, account)
}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address user) view returns(bool)
func (_LaunchVault *LaunchVaultCaller) HasWithdrawn(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "hasWithdrawn", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address user) view returns(bool)
func (_LaunchVault *LaunchVaultSession) HasWithdrawn(user common.Address) (bool, error) {
	return _LaunchVault.Contract.HasWithdrawn(&_LaunchVault.CallOpts, user)
}

// HasWithdrawn is a free data retrieval call binding the contract method 0x5e2c19db.
//
// Solidity: function hasWithdrawn(address user) view returns(bool)
func (_LaunchVault *LaunchVaultCallerSession) HasWithdrawn(user common.Address) (bool, error) {
	return _LaunchVault.Contract.HasWithdrawn(&_LaunchVault.CallOpts, user)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LaunchVault *LaunchVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LaunchVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LaunchVault *LaunchVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LaunchVault.Contract.SupportsInterface(&_LaunchVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LaunchVault *LaunchVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LaunchVault.Contract.SupportsInterface(&_LaunchVault.CallOpts, interfaceId)
}

// DeployAssets is a paid mutator transaction binding the contract method 0x891697ca.
//
// Solidity: function deployAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactor) DeployAssets(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "deployAssets", asset, amount)
}

// DeployAssets is a paid mutator transaction binding the contract method 0x891697ca.
//
// Solidity: function deployAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultSession) DeployAssets(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.DeployAssets(&_LaunchVault.TransactOpts, asset, amount)
}

// DeployAssets is a paid mutator transaction binding the contract method 0x891697ca.
//
// Solidity: function deployAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactorSession) DeployAssets(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.DeployAssets(&_LaunchVault.TransactOpts, asset, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "deposit", asset, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultSession) Deposit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.Deposit(&_LaunchVault.TransactOpts, asset, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactorSession) Deposit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.Deposit(&_LaunchVault.TransactOpts, asset, amount)
}

// DisableLPTransfers is a paid mutator transaction binding the contract method 0x2859b553.
//
// Solidity: function disableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultTransactor) DisableLPTransfers(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "disableLPTransfers", asset)
}

// DisableLPTransfers is a paid mutator transaction binding the contract method 0x2859b553.
//
// Solidity: function disableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultSession) DisableLPTransfers(asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.DisableLPTransfers(&_LaunchVault.TransactOpts, asset)
}

// DisableLPTransfers is a paid mutator transaction binding the contract method 0x2859b553.
//
// Solidity: function disableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultTransactorSession) DisableLPTransfers(asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.DisableLPTransfers(&_LaunchVault.TransactOpts, asset)
}

// EnableLPTransfers is a paid mutator transaction binding the contract method 0x4ada47c0.
//
// Solidity: function enableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultTransactor) EnableLPTransfers(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "enableLPTransfers", asset)
}

// EnableLPTransfers is a paid mutator transaction binding the contract method 0x4ada47c0.
//
// Solidity: function enableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultSession) EnableLPTransfers(asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.EnableLPTransfers(&_LaunchVault.TransactOpts, asset)
}

// EnableLPTransfers is a paid mutator transaction binding the contract method 0x4ada47c0.
//
// Solidity: function enableLPTransfers(address asset) returns()
func (_LaunchVault *LaunchVaultTransactorSession) EnableLPTransfers(asset common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.EnableLPTransfers(&_LaunchVault.TransactOpts, asset)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.GrantRole(&_LaunchVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.GrantRole(&_LaunchVault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LaunchVault *LaunchVaultTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LaunchVault *LaunchVaultSession) Initialize() (*types.Transaction, error) {
	return _LaunchVault.Contract.Initialize(&_LaunchVault.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_LaunchVault *LaunchVaultTransactorSession) Initialize() (*types.Transaction, error) {
	return _LaunchVault.Contract.Initialize(&_LaunchVault.TransactOpts)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x7f823d86.
//
// Solidity: function registerAsset(address asset, address vaultToken, address deployer) returns()
func (_LaunchVault *LaunchVaultTransactor) RegisterAsset(opts *bind.TransactOpts, asset common.Address, vaultToken common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "registerAsset", asset, vaultToken, deployer)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x7f823d86.
//
// Solidity: function registerAsset(address asset, address vaultToken, address deployer) returns()
func (_LaunchVault *LaunchVaultSession) RegisterAsset(asset common.Address, vaultToken common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RegisterAsset(&_LaunchVault.TransactOpts, asset, vaultToken, deployer)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x7f823d86.
//
// Solidity: function registerAsset(address asset, address vaultToken, address deployer) returns()
func (_LaunchVault *LaunchVaultTransactorSession) RegisterAsset(asset common.Address, vaultToken common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RegisterAsset(&_LaunchVault.TransactOpts, asset, vaultToken, deployer)
}

// RemoveAssets is a paid mutator transaction binding the contract method 0x2c474ed6.
//
// Solidity: function removeAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactor) RemoveAssets(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "removeAssets", asset, amount)
}

// RemoveAssets is a paid mutator transaction binding the contract method 0x2c474ed6.
//
// Solidity: function removeAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultSession) RemoveAssets(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.RemoveAssets(&_LaunchVault.TransactOpts, asset, amount)
}

// RemoveAssets is a paid mutator transaction binding the contract method 0x2c474ed6.
//
// Solidity: function removeAssets(address asset, uint256 amount) returns()
func (_LaunchVault *LaunchVaultTransactorSession) RemoveAssets(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.RemoveAssets(&_LaunchVault.TransactOpts, asset, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LaunchVault *LaunchVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LaunchVault *LaunchVaultSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RenounceRole(&_LaunchVault.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LaunchVault *LaunchVaultTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RenounceRole(&_LaunchVault.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RevokeRole(&_LaunchVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LaunchVault *LaunchVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.RevokeRole(&_LaunchVault.TransactOpts, role, account)
}

// UpdateDeployer is a paid mutator transaction binding the contract method 0x3f42015f.
//
// Solidity: function updateDeployer(address asset, address deployer) returns()
func (_LaunchVault *LaunchVaultTransactor) UpdateDeployer(opts *bind.TransactOpts, asset common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "updateDeployer", asset, deployer)
}

// UpdateDeployer is a paid mutator transaction binding the contract method 0x3f42015f.
//
// Solidity: function updateDeployer(address asset, address deployer) returns()
func (_LaunchVault *LaunchVaultSession) UpdateDeployer(asset common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.UpdateDeployer(&_LaunchVault.TransactOpts, asset, deployer)
}

// UpdateDeployer is a paid mutator transaction binding the contract method 0x3f42015f.
//
// Solidity: function updateDeployer(address asset, address deployer) returns()
func (_LaunchVault *LaunchVaultTransactorSession) UpdateDeployer(asset common.Address, deployer common.Address) (*types.Transaction, error) {
	return _LaunchVault.Contract.UpdateDeployer(&_LaunchVault.TransactOpts, asset, deployer)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 vaultTokensToBurn) returns()
func (_LaunchVault *LaunchVaultTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, vaultTokensToBurn *big.Int) (*types.Transaction, error) {
	return _LaunchVault.contract.Transact(opts, "withdraw", asset, vaultTokensToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 vaultTokensToBurn) returns()
func (_LaunchVault *LaunchVaultSession) Withdraw(asset common.Address, vaultTokensToBurn *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.Withdraw(&_LaunchVault.TransactOpts, asset, vaultTokensToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 vaultTokensToBurn) returns()
func (_LaunchVault *LaunchVaultTransactorSession) Withdraw(asset common.Address, vaultTokensToBurn *big.Int) (*types.Transaction, error) {
	return _LaunchVault.Contract.Withdraw(&_LaunchVault.TransactOpts, asset, vaultTokensToBurn)
}

// LaunchVaultBalanceChangeIterator is returned from FilterBalanceChange and is used to iterate over the raw logs and unpacked data for BalanceChange events raised by the LaunchVault contract.
type LaunchVaultBalanceChangeIterator struct {
	Event *LaunchVaultBalanceChange // Event containing the contract specifics and raw log

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
func (it *LaunchVaultBalanceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchVaultBalanceChange)
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
		it.Event = new(LaunchVaultBalanceChange)
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
func (it *LaunchVaultBalanceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchVaultBalanceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchVaultBalanceChange represents a BalanceChange event raised by the LaunchVault contract.
type LaunchVaultBalanceChange struct {
	IsDeposit bool
	Asset     common.Address
	User      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBalanceChange is a free log retrieval operation binding the contract event 0xd34c461d783a9885b01b2fa71160857f48963a0ed4514e902adbda302ea80100.
//
// Solidity: event BalanceChange(bool isDeposit, address asset, address user, uint256 amount)
func (_LaunchVault *LaunchVaultFilterer) FilterBalanceChange(opts *bind.FilterOpts) (*LaunchVaultBalanceChangeIterator, error) {

	logs, sub, err := _LaunchVault.contract.FilterLogs(opts, "BalanceChange")
	if err != nil {
		return nil, err
	}
	return &LaunchVaultBalanceChangeIterator{contract: _LaunchVault.contract, event: "BalanceChange", logs: logs, sub: sub}, nil
}

// WatchBalanceChange is a free log subscription operation binding the contract event 0xd34c461d783a9885b01b2fa71160857f48963a0ed4514e902adbda302ea80100.
//
// Solidity: event BalanceChange(bool isDeposit, address asset, address user, uint256 amount)
func (_LaunchVault *LaunchVaultFilterer) WatchBalanceChange(opts *bind.WatchOpts, sink chan<- *LaunchVaultBalanceChange) (event.Subscription, error) {

	logs, sub, err := _LaunchVault.contract.WatchLogs(opts, "BalanceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchVaultBalanceChange)
				if err := _LaunchVault.contract.UnpackLog(event, "BalanceChange", log); err != nil {
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

// ParseBalanceChange is a log parse operation binding the contract event 0xd34c461d783a9885b01b2fa71160857f48963a0ed4514e902adbda302ea80100.
//
// Solidity: event BalanceChange(bool isDeposit, address asset, address user, uint256 amount)
func (_LaunchVault *LaunchVaultFilterer) ParseBalanceChange(log types.Log) (*LaunchVaultBalanceChange, error) {
	event := new(LaunchVaultBalanceChange)
	if err := _LaunchVault.contract.UnpackLog(event, "BalanceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LaunchVault contract.
type LaunchVaultInitializedIterator struct {
	Event *LaunchVaultInitialized // Event containing the contract specifics and raw log

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
func (it *LaunchVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchVaultInitialized)
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
		it.Event = new(LaunchVaultInitialized)
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
func (it *LaunchVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchVaultInitialized represents a Initialized event raised by the LaunchVault contract.
type LaunchVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LaunchVault *LaunchVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*LaunchVaultInitializedIterator, error) {

	logs, sub, err := _LaunchVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LaunchVaultInitializedIterator{contract: _LaunchVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LaunchVault *LaunchVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LaunchVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _LaunchVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchVaultInitialized)
				if err := _LaunchVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LaunchVault *LaunchVaultFilterer) ParseInitialized(log types.Log) (*LaunchVaultInitialized, error) {
	event := new(LaunchVaultInitialized)
	if err := _LaunchVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LaunchVault contract.
type LaunchVaultRoleAdminChangedIterator struct {
	Event *LaunchVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LaunchVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchVaultRoleAdminChanged)
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
		it.Event = new(LaunchVaultRoleAdminChanged)
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
func (it *LaunchVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchVaultRoleAdminChanged represents a RoleAdminChanged event raised by the LaunchVault contract.
type LaunchVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LaunchVault *LaunchVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LaunchVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LaunchVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultRoleAdminChangedIterator{contract: _LaunchVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LaunchVault *LaunchVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LaunchVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LaunchVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchVaultRoleAdminChanged)
				if err := _LaunchVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LaunchVault *LaunchVaultFilterer) ParseRoleAdminChanged(log types.Log) (*LaunchVaultRoleAdminChanged, error) {
	event := new(LaunchVaultRoleAdminChanged)
	if err := _LaunchVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LaunchVault contract.
type LaunchVaultRoleGrantedIterator struct {
	Event *LaunchVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *LaunchVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchVaultRoleGranted)
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
		it.Event = new(LaunchVaultRoleGranted)
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
func (it *LaunchVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchVaultRoleGranted represents a RoleGranted event raised by the LaunchVault contract.
type LaunchVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LaunchVault *LaunchVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LaunchVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _LaunchVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultRoleGrantedIterator{contract: _LaunchVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LaunchVault *LaunchVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LaunchVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LaunchVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchVaultRoleGranted)
				if err := _LaunchVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LaunchVault *LaunchVaultFilterer) ParseRoleGranted(log types.Log) (*LaunchVaultRoleGranted, error) {
	event := new(LaunchVaultRoleGranted)
	if err := _LaunchVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LaunchVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LaunchVault contract.
type LaunchVaultRoleRevokedIterator struct {
	Event *LaunchVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LaunchVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LaunchVaultRoleRevoked)
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
		it.Event = new(LaunchVaultRoleRevoked)
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
func (it *LaunchVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LaunchVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LaunchVaultRoleRevoked represents a RoleRevoked event raised by the LaunchVault contract.
type LaunchVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LaunchVault *LaunchVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LaunchVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _LaunchVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LaunchVaultRoleRevokedIterator{contract: _LaunchVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LaunchVault *LaunchVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LaunchVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LaunchVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LaunchVaultRoleRevoked)
				if err := _LaunchVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LaunchVault *LaunchVaultFilterer) ParseRoleRevoked(log types.Log) (*LaunchVaultRoleRevoked, error) {
	event := new(LaunchVaultRoleRevoked)
	if err := _LaunchVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
