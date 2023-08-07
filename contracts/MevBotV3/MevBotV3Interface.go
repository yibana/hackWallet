// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MevBotV3Interface

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

// MevBotV3InterfaceMetaData contains all meta data concerning the MevBotV3Interface contract.
var MevBotV3InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegatecall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"}],\"name\":\"SendMulti_assembly_yul\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"targetwethBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bribe\",\"type\":\"uint256\"}],\"name\":\"SendMultiCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bribe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlockNumber\",\"type\":\"uint256\"}],\"name\":\"SendMultiNoCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetwethBalance\",\"type\":\"uint256\"}],\"name\":\"targetPay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MevBotV3InterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MevBotV3InterfaceMetaData.ABI instead.
var MevBotV3InterfaceABI = MevBotV3InterfaceMetaData.ABI

// MevBotV3Interface is an auto generated Go binding around an Ethereum contract.
type MevBotV3Interface struct {
	MevBotV3InterfaceCaller     // Read-only binding to the contract
	MevBotV3InterfaceTransactor // Write-only binding to the contract
	MevBotV3InterfaceFilterer   // Log filterer for contract events
}

// MevBotV3InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MevBotV3InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevBotV3InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MevBotV3InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevBotV3InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MevBotV3InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevBotV3InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MevBotV3InterfaceSession struct {
	Contract     *MevBotV3Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MevBotV3InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MevBotV3InterfaceCallerSession struct {
	Contract *MevBotV3InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MevBotV3InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MevBotV3InterfaceTransactorSession struct {
	Contract     *MevBotV3InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MevBotV3InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MevBotV3InterfaceRaw struct {
	Contract *MevBotV3Interface // Generic contract binding to access the raw methods on
}

// MevBotV3InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MevBotV3InterfaceCallerRaw struct {
	Contract *MevBotV3InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// MevBotV3InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MevBotV3InterfaceTransactorRaw struct {
	Contract *MevBotV3InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMevBotV3Interface creates a new instance of MevBotV3Interface, bound to a specific deployed contract.
func NewMevBotV3Interface(address common.Address, backend bind.ContractBackend) (*MevBotV3Interface, error) {
	contract, err := bindMevBotV3Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MevBotV3Interface{MevBotV3InterfaceCaller: MevBotV3InterfaceCaller{contract: contract}, MevBotV3InterfaceTransactor: MevBotV3InterfaceTransactor{contract: contract}, MevBotV3InterfaceFilterer: MevBotV3InterfaceFilterer{contract: contract}}, nil
}

// NewMevBotV3InterfaceCaller creates a new read-only instance of MevBotV3Interface, bound to a specific deployed contract.
func NewMevBotV3InterfaceCaller(address common.Address, caller bind.ContractCaller) (*MevBotV3InterfaceCaller, error) {
	contract, err := bindMevBotV3Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MevBotV3InterfaceCaller{contract: contract}, nil
}

// NewMevBotV3InterfaceTransactor creates a new write-only instance of MevBotV3Interface, bound to a specific deployed contract.
func NewMevBotV3InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MevBotV3InterfaceTransactor, error) {
	contract, err := bindMevBotV3Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MevBotV3InterfaceTransactor{contract: contract}, nil
}

// NewMevBotV3InterfaceFilterer creates a new log filterer instance of MevBotV3Interface, bound to a specific deployed contract.
func NewMevBotV3InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MevBotV3InterfaceFilterer, error) {
	contract, err := bindMevBotV3Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MevBotV3InterfaceFilterer{contract: contract}, nil
}

// bindMevBotV3Interface binds a generic wrapper to an already deployed contract.
func bindMevBotV3Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MevBotV3InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MevBotV3Interface *MevBotV3InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MevBotV3Interface.Contract.MevBotV3InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MevBotV3Interface *MevBotV3InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.MevBotV3InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MevBotV3Interface *MevBotV3InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.MevBotV3InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MevBotV3Interface *MevBotV3InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MevBotV3Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MevBotV3Interface *MevBotV3InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MevBotV3Interface *MevBotV3InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.contract.Transact(opts, method, params...)
}

// SendMultiCheck is a paid mutator transaction binding the contract method 0xd201b711.
//
// Solidity: function SendMultiCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 targetwethBalance, uint256 bribe) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) SendMultiCheck(opts *bind.TransactOpts, _targets []common.Address, _payloads [][]byte, values []*big.Int, targetwethBalance *big.Int, bribe *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "SendMultiCheck", _targets, _payloads, values, targetwethBalance, bribe)
}

// SendMultiCheck is a paid mutator transaction binding the contract method 0xd201b711.
//
// Solidity: function SendMultiCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 targetwethBalance, uint256 bribe) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) SendMultiCheck(_targets []common.Address, _payloads [][]byte, values []*big.Int, targetwethBalance *big.Int, bribe *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiCheck(&_MevBotV3Interface.TransactOpts, _targets, _payloads, values, targetwethBalance, bribe)
}

// SendMultiCheck is a paid mutator transaction binding the contract method 0xd201b711.
//
// Solidity: function SendMultiCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 targetwethBalance, uint256 bribe) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) SendMultiCheck(_targets []common.Address, _payloads [][]byte, values []*big.Int, targetwethBalance *big.Int, bribe *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiCheck(&_MevBotV3Interface.TransactOpts, _targets, _payloads, values, targetwethBalance, bribe)
}

// SendMultiNoCheck is a paid mutator transaction binding the contract method 0xce8873a8.
//
// Solidity: function SendMultiNoCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 bribe, uint256 endBlockNumber) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) SendMultiNoCheck(opts *bind.TransactOpts, _targets []common.Address, _payloads [][]byte, values []*big.Int, bribe *big.Int, endBlockNumber *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "SendMultiNoCheck", _targets, _payloads, values, bribe, endBlockNumber)
}

// SendMultiNoCheck is a paid mutator transaction binding the contract method 0xce8873a8.
//
// Solidity: function SendMultiNoCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 bribe, uint256 endBlockNumber) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) SendMultiNoCheck(_targets []common.Address, _payloads [][]byte, values []*big.Int, bribe *big.Int, endBlockNumber *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiNoCheck(&_MevBotV3Interface.TransactOpts, _targets, _payloads, values, bribe, endBlockNumber)
}

// SendMultiNoCheck is a paid mutator transaction binding the contract method 0xce8873a8.
//
// Solidity: function SendMultiNoCheck(address[] _targets, bytes[] _payloads, uint256[] values, uint256 bribe, uint256 endBlockNumber) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) SendMultiNoCheck(_targets []common.Address, _payloads [][]byte, values []*big.Int, bribe *big.Int, endBlockNumber *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiNoCheck(&_MevBotV3Interface.TransactOpts, _targets, _payloads, values, bribe, endBlockNumber)
}

// SendMultiAssemblyYul is a paid mutator transaction binding the contract method 0x0449a7a9.
//
// Solidity: function SendMulti_assembly_yul(address[] _targets, uint256[] values, bytes[] _payloads) returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) SendMultiAssemblyYul(opts *bind.TransactOpts, _targets []common.Address, values []*big.Int, _payloads [][]byte) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "SendMulti_assembly_yul", _targets, values, _payloads)
}

// SendMultiAssemblyYul is a paid mutator transaction binding the contract method 0x0449a7a9.
//
// Solidity: function SendMulti_assembly_yul(address[] _targets, uint256[] values, bytes[] _payloads) returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) SendMultiAssemblyYul(_targets []common.Address, values []*big.Int, _payloads [][]byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiAssemblyYul(&_MevBotV3Interface.TransactOpts, _targets, values, _payloads)
}

// SendMultiAssemblyYul is a paid mutator transaction binding the contract method 0x0449a7a9.
//
// Solidity: function SendMulti_assembly_yul(address[] _targets, uint256[] values, bytes[] _payloads) returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) SendMultiAssemblyYul(_targets []common.Address, values []*big.Int, _payloads [][]byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.SendMultiAssemblyYul(&_MevBotV3Interface.TransactOpts, _targets, values, _payloads)
}

// Delegatecall is a paid mutator transaction binding the contract method 0x6e0aacf7.
//
// Solidity: function delegatecall(address addr, bytes data) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) Delegatecall(opts *bind.TransactOpts, addr common.Address, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "delegatecall", addr, data)
}

// Delegatecall is a paid mutator transaction binding the contract method 0x6e0aacf7.
//
// Solidity: function delegatecall(address addr, bytes data) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) Delegatecall(addr common.Address, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Delegatecall(&_MevBotV3Interface.TransactOpts, addr, data)
}

// Delegatecall is a paid mutator transaction binding the contract method 0x6e0aacf7.
//
// Solidity: function delegatecall(address addr, bytes data) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) Delegatecall(addr common.Address, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Delegatecall(&_MevBotV3Interface.TransactOpts, addr, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC1155BatchReceived(&_MevBotV3Interface.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC1155BatchReceived(&_MevBotV3Interface.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC1155Received(&_MevBotV3Interface.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC1155Received(&_MevBotV3Interface.TransactOpts, operator, from, id, value, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC721Received(&_MevBotV3Interface.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.OnERC721Received(&_MevBotV3Interface.TransactOpts, operator, from, tokenId, data)
}

// TargetPay is a paid mutator transaction binding the contract method 0xf5475936.
//
// Solidity: function targetPay(uint256 targetwethBalance) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) TargetPay(opts *bind.TransactOpts, targetwethBalance *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "targetPay", targetwethBalance)
}

// TargetPay is a paid mutator transaction binding the contract method 0xf5475936.
//
// Solidity: function targetPay(uint256 targetwethBalance) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) TargetPay(targetwethBalance *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.TargetPay(&_MevBotV3Interface.TransactOpts, targetwethBalance)
}

// TargetPay is a paid mutator transaction binding the contract method 0xf5475936.
//
// Solidity: function targetPay(uint256 targetwethBalance) payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) TargetPay(targetwethBalance *big.Int) (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.TargetPay(&_MevBotV3Interface.TransactOpts, targetwethBalance)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) Withdraw() (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Withdraw(&_MevBotV3Interface.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Withdraw(&_MevBotV3Interface.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MevBotV3Interface.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceSession) Receive() (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Receive(&_MevBotV3Interface.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MevBotV3Interface *MevBotV3InterfaceTransactorSession) Receive() (*types.Transaction, error) {
	return _MevBotV3Interface.Contract.Receive(&_MevBotV3Interface.TransactOpts)
}
