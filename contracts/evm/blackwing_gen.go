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

// SwapPayload is an auto generated low-level Go binding around an user-defined struct.
type SwapPayload struct {
	FromTokenAddress common.Address
	ToTokenAddress   common.Address
	Target           common.Address
	Value            *big.Int
	Data             []byte
}

// SwapResult is an auto generated low-level Go binding around an user-defined struct.
type SwapResult struct {
	FromTokenGiven  *big.Int
	ToTokenReceived *big.Int
}

// BlackwingMetaData contains all meta data concerning the Blackwing contract.
var BlackwingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"inputdata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"returndata\",\"type\":\"bytes\"}],\"name\":\"TransactionFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"duplicate\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenGiven\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenReceived\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSwapResult\",\"name\":\"result\",\"type\":\"tuple\"}],\"name\":\"OrderProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"isOrderProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"id\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSwapPayload\",\"name\":\"swapPayload\",\"type\":\"tuple\"}],\"name\":\"processOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenGiven\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenReceived\",\"type\":\"uint256\"}],\"internalType\":\"structSwapResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlackwingABI is the input ABI used to generate the binding from.
// Deprecated: Use BlackwingMetaData.ABI instead.
var BlackwingABI = BlackwingMetaData.ABI

// Blackwing is an auto generated Go binding around an Ethereum contract.
type Blackwing struct {
	BlackwingCaller     // Read-only binding to the contract
	BlackwingTransactor // Write-only binding to the contract
	BlackwingFilterer   // Log filterer for contract events
}

// BlackwingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlackwingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackwingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlackwingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackwingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlackwingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackwingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlackwingSession struct {
	Contract     *Blackwing        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlackwingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlackwingCallerSession struct {
	Contract *BlackwingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlackwingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlackwingTransactorSession struct {
	Contract     *BlackwingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlackwingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlackwingRaw struct {
	Contract *Blackwing // Generic contract binding to access the raw methods on
}

// BlackwingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlackwingCallerRaw struct {
	Contract *BlackwingCaller // Generic read-only contract binding to access the raw methods on
}

// BlackwingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlackwingTransactorRaw struct {
	Contract *BlackwingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlackwing creates a new instance of Blackwing, bound to a specific deployed contract.
func NewBlackwing(address common.Address, backend bind.ContractBackend) (*Blackwing, error) {
	contract, err := bindBlackwing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blackwing{BlackwingCaller: BlackwingCaller{contract: contract}, BlackwingTransactor: BlackwingTransactor{contract: contract}, BlackwingFilterer: BlackwingFilterer{contract: contract}}, nil
}

// NewBlackwingCaller creates a new read-only instance of Blackwing, bound to a specific deployed contract.
func NewBlackwingCaller(address common.Address, caller bind.ContractCaller) (*BlackwingCaller, error) {
	contract, err := bindBlackwing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlackwingCaller{contract: contract}, nil
}

// NewBlackwingTransactor creates a new write-only instance of Blackwing, bound to a specific deployed contract.
func NewBlackwingTransactor(address common.Address, transactor bind.ContractTransactor) (*BlackwingTransactor, error) {
	contract, err := bindBlackwing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlackwingTransactor{contract: contract}, nil
}

// NewBlackwingFilterer creates a new log filterer instance of Blackwing, bound to a specific deployed contract.
func NewBlackwingFilterer(address common.Address, filterer bind.ContractFilterer) (*BlackwingFilterer, error) {
	contract, err := bindBlackwing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlackwingFilterer{contract: contract}, nil
}

// bindBlackwing binds a generic wrapper to an already deployed contract.
func bindBlackwing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlackwingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blackwing *BlackwingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blackwing.Contract.BlackwingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blackwing *BlackwingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blackwing.Contract.BlackwingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blackwing *BlackwingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blackwing.Contract.BlackwingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blackwing *BlackwingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blackwing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blackwing *BlackwingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blackwing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blackwing *BlackwingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blackwing.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Blackwing.Contract.DEFAULTADMINROLE(&_Blackwing.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Blackwing.Contract.DEFAULTADMINROLE(&_Blackwing.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingSession) OWNERROLE() ([32]byte, error) {
	return _Blackwing.Contract.OWNERROLE(&_Blackwing.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Blackwing *BlackwingCallerSession) OWNERROLE() ([32]byte, error) {
	return _Blackwing.Contract.OWNERROLE(&_Blackwing.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blackwing *BlackwingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blackwing *BlackwingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Blackwing.Contract.GetRoleAdmin(&_Blackwing.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blackwing *BlackwingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Blackwing.Contract.GetRoleAdmin(&_Blackwing.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blackwing *BlackwingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blackwing *BlackwingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Blackwing.Contract.HasRole(&_Blackwing.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blackwing *BlackwingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Blackwing.Contract.HasRole(&_Blackwing.CallOpts, role, account)
}

// IsOrderProcessed is a free data retrieval call binding the contract method 0x06efeb11.
//
// Solidity: function isOrderProcessed(bytes16 id) view returns(bool)
func (_Blackwing *BlackwingCaller) IsOrderProcessed(opts *bind.CallOpts, id [16]byte) (bool, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "isOrderProcessed", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderProcessed is a free data retrieval call binding the contract method 0x06efeb11.
//
// Solidity: function isOrderProcessed(bytes16 id) view returns(bool)
func (_Blackwing *BlackwingSession) IsOrderProcessed(id [16]byte) (bool, error) {
	return _Blackwing.Contract.IsOrderProcessed(&_Blackwing.CallOpts, id)
}

// IsOrderProcessed is a free data retrieval call binding the contract method 0x06efeb11.
//
// Solidity: function isOrderProcessed(bytes16 id) view returns(bool)
func (_Blackwing *BlackwingCallerSession) IsOrderProcessed(id [16]byte) (bool, error) {
	return _Blackwing.Contract.IsOrderProcessed(&_Blackwing.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blackwing *BlackwingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Blackwing.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blackwing *BlackwingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Blackwing.Contract.SupportsInterface(&_Blackwing.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blackwing *BlackwingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Blackwing.Contract.SupportsInterface(&_Blackwing.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.GrantRole(&_Blackwing.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.GrantRole(&_Blackwing.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blackwing *BlackwingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blackwing.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blackwing *BlackwingSession) Initialize() (*types.Transaction, error) {
	return _Blackwing.Contract.Initialize(&_Blackwing.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Blackwing *BlackwingTransactorSession) Initialize() (*types.Transaction, error) {
	return _Blackwing.Contract.Initialize(&_Blackwing.TransactOpts)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0x564651b3.
//
// Solidity: function processOrder(bytes16 id, (address,address,address,uint256,bytes) swapPayload) returns(bool, (uint256,uint256))
func (_Blackwing *BlackwingTransactor) ProcessOrder(opts *bind.TransactOpts, id [16]byte, swapPayload SwapPayload) (*types.Transaction, error) {
	return _Blackwing.contract.Transact(opts, "processOrder", id, swapPayload)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0x564651b3.
//
// Solidity: function processOrder(bytes16 id, (address,address,address,uint256,bytes) swapPayload) returns(bool, (uint256,uint256))
func (_Blackwing *BlackwingSession) ProcessOrder(id [16]byte, swapPayload SwapPayload) (*types.Transaction, error) {
	return _Blackwing.Contract.ProcessOrder(&_Blackwing.TransactOpts, id, swapPayload)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0x564651b3.
//
// Solidity: function processOrder(bytes16 id, (address,address,address,uint256,bytes) swapPayload) returns(bool, (uint256,uint256))
func (_Blackwing *BlackwingTransactorSession) ProcessOrder(id [16]byte, swapPayload SwapPayload) (*types.Transaction, error) {
	return _Blackwing.Contract.ProcessOrder(&_Blackwing.TransactOpts, id, swapPayload)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blackwing *BlackwingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blackwing.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blackwing *BlackwingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.RenounceRole(&_Blackwing.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blackwing *BlackwingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.RenounceRole(&_Blackwing.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.RevokeRole(&_Blackwing.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blackwing *BlackwingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blackwing.Contract.RevokeRole(&_Blackwing.TransactOpts, role, account)
}

// BlackwingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Blackwing contract.
type BlackwingInitializedIterator struct {
	Event *BlackwingInitialized // Event containing the contract specifics and raw log

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
func (it *BlackwingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackwingInitialized)
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
		it.Event = new(BlackwingInitialized)
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
func (it *BlackwingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackwingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackwingInitialized represents a Initialized event raised by the Blackwing contract.
type BlackwingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Blackwing *BlackwingFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlackwingInitializedIterator, error) {

	logs, sub, err := _Blackwing.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlackwingInitializedIterator{contract: _Blackwing.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Blackwing *BlackwingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlackwingInitialized) (event.Subscription, error) {

	logs, sub, err := _Blackwing.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackwingInitialized)
				if err := _Blackwing.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Blackwing *BlackwingFilterer) ParseInitialized(log types.Log) (*BlackwingInitialized, error) {
	event := new(BlackwingInitialized)
	if err := _Blackwing.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackwingOrderProcessedIterator is returned from FilterOrderProcessed and is used to iterate over the raw logs and unpacked data for OrderProcessed events raised by the Blackwing contract.
type BlackwingOrderProcessedIterator struct {
	Event *BlackwingOrderProcessed // Event containing the contract specifics and raw log

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
func (it *BlackwingOrderProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackwingOrderProcessed)
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
		it.Event = new(BlackwingOrderProcessed)
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
func (it *BlackwingOrderProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackwingOrderProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackwingOrderProcessed represents a OrderProcessed event raised by the Blackwing contract.
type BlackwingOrderProcessed struct {
	Id        [16]byte
	Duplicate bool
	Result    SwapResult
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderProcessed is a free log retrieval operation binding the contract event 0x38c87bb5da85dbbeae691205e65468ae4a562085706bd1cc97fdc53a0a0778b3.
//
// Solidity: event OrderProcessed(bytes16 id, bool duplicate, (uint256,uint256) result)
func (_Blackwing *BlackwingFilterer) FilterOrderProcessed(opts *bind.FilterOpts) (*BlackwingOrderProcessedIterator, error) {

	logs, sub, err := _Blackwing.contract.FilterLogs(opts, "OrderProcessed")
	if err != nil {
		return nil, err
	}
	return &BlackwingOrderProcessedIterator{contract: _Blackwing.contract, event: "OrderProcessed", logs: logs, sub: sub}, nil
}

// WatchOrderProcessed is a free log subscription operation binding the contract event 0x38c87bb5da85dbbeae691205e65468ae4a562085706bd1cc97fdc53a0a0778b3.
//
// Solidity: event OrderProcessed(bytes16 id, bool duplicate, (uint256,uint256) result)
func (_Blackwing *BlackwingFilterer) WatchOrderProcessed(opts *bind.WatchOpts, sink chan<- *BlackwingOrderProcessed) (event.Subscription, error) {

	logs, sub, err := _Blackwing.contract.WatchLogs(opts, "OrderProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackwingOrderProcessed)
				if err := _Blackwing.contract.UnpackLog(event, "OrderProcessed", log); err != nil {
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

// ParseOrderProcessed is a log parse operation binding the contract event 0x38c87bb5da85dbbeae691205e65468ae4a562085706bd1cc97fdc53a0a0778b3.
//
// Solidity: event OrderProcessed(bytes16 id, bool duplicate, (uint256,uint256) result)
func (_Blackwing *BlackwingFilterer) ParseOrderProcessed(log types.Log) (*BlackwingOrderProcessed, error) {
	event := new(BlackwingOrderProcessed)
	if err := _Blackwing.contract.UnpackLog(event, "OrderProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackwingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Blackwing contract.
type BlackwingRoleAdminChangedIterator struct {
	Event *BlackwingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlackwingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackwingRoleAdminChanged)
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
		it.Event = new(BlackwingRoleAdminChanged)
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
func (it *BlackwingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackwingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackwingRoleAdminChanged represents a RoleAdminChanged event raised by the Blackwing contract.
type BlackwingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Blackwing *BlackwingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BlackwingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Blackwing.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BlackwingRoleAdminChangedIterator{contract: _Blackwing.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Blackwing *BlackwingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BlackwingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Blackwing.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackwingRoleAdminChanged)
				if err := _Blackwing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Blackwing *BlackwingFilterer) ParseRoleAdminChanged(log types.Log) (*BlackwingRoleAdminChanged, error) {
	event := new(BlackwingRoleAdminChanged)
	if err := _Blackwing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackwingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Blackwing contract.
type BlackwingRoleGrantedIterator struct {
	Event *BlackwingRoleGranted // Event containing the contract specifics and raw log

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
func (it *BlackwingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackwingRoleGranted)
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
		it.Event = new(BlackwingRoleGranted)
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
func (it *BlackwingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackwingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackwingRoleGranted represents a RoleGranted event raised by the Blackwing contract.
type BlackwingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blackwing *BlackwingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlackwingRoleGrantedIterator, error) {

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

	logs, sub, err := _Blackwing.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlackwingRoleGrantedIterator{contract: _Blackwing.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blackwing *BlackwingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BlackwingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Blackwing.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackwingRoleGranted)
				if err := _Blackwing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Blackwing *BlackwingFilterer) ParseRoleGranted(log types.Log) (*BlackwingRoleGranted, error) {
	event := new(BlackwingRoleGranted)
	if err := _Blackwing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackwingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Blackwing contract.
type BlackwingRoleRevokedIterator struct {
	Event *BlackwingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BlackwingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackwingRoleRevoked)
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
		it.Event = new(BlackwingRoleRevoked)
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
func (it *BlackwingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackwingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackwingRoleRevoked represents a RoleRevoked event raised by the Blackwing contract.
type BlackwingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blackwing *BlackwingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlackwingRoleRevokedIterator, error) {

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

	logs, sub, err := _Blackwing.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlackwingRoleRevokedIterator{contract: _Blackwing.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blackwing *BlackwingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BlackwingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Blackwing.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackwingRoleRevoked)
				if err := _Blackwing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Blackwing *BlackwingFilterer) ParseRoleRevoked(log types.Log) (*BlackwingRoleRevoked, error) {
	event := new(BlackwingRoleRevoked)
	if err := _Blackwing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
