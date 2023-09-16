// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC6551Registry

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

// ERC6551RegistryMetaData contains all meta data concerning the ERC6551Registry contract.
var ERC6551RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InitializationFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC6551RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC6551RegistryMetaData.ABI instead.
var ERC6551RegistryABI = ERC6551RegistryMetaData.ABI

// ERC6551Registry is an auto generated Go binding around an Ethereum contract.
type ERC6551Registry struct {
	ERC6551RegistryCaller     // Read-only binding to the contract
	ERC6551RegistryTransactor // Write-only binding to the contract
	ERC6551RegistryFilterer   // Log filterer for contract events
}

// ERC6551RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC6551RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC6551RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC6551RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC6551RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC6551RegistrySession struct {
	Contract     *ERC6551Registry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC6551RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC6551RegistryCallerSession struct {
	Contract *ERC6551RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC6551RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC6551RegistryTransactorSession struct {
	Contract     *ERC6551RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC6551RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC6551RegistryRaw struct {
	Contract *ERC6551Registry // Generic contract binding to access the raw methods on
}

// ERC6551RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC6551RegistryCallerRaw struct {
	Contract *ERC6551RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ERC6551RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC6551RegistryTransactorRaw struct {
	Contract *ERC6551RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC6551Registry creates a new instance of ERC6551Registry, bound to a specific deployed contract.
func NewERC6551Registry(address common.Address, backend bind.ContractBackend) (*ERC6551Registry, error) {
	contract, err := bindERC6551Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC6551Registry{ERC6551RegistryCaller: ERC6551RegistryCaller{contract: contract}, ERC6551RegistryTransactor: ERC6551RegistryTransactor{contract: contract}, ERC6551RegistryFilterer: ERC6551RegistryFilterer{contract: contract}}, nil
}

// NewERC6551RegistryCaller creates a new read-only instance of ERC6551Registry, bound to a specific deployed contract.
func NewERC6551RegistryCaller(address common.Address, caller bind.ContractCaller) (*ERC6551RegistryCaller, error) {
	contract, err := bindERC6551Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC6551RegistryCaller{contract: contract}, nil
}

// NewERC6551RegistryTransactor creates a new write-only instance of ERC6551Registry, bound to a specific deployed contract.
func NewERC6551RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC6551RegistryTransactor, error) {
	contract, err := bindERC6551Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC6551RegistryTransactor{contract: contract}, nil
}

// NewERC6551RegistryFilterer creates a new log filterer instance of ERC6551Registry, bound to a specific deployed contract.
func NewERC6551RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC6551RegistryFilterer, error) {
	contract, err := bindERC6551Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC6551RegistryFilterer{contract: contract}, nil
}

// bindERC6551Registry binds a generic wrapper to an already deployed contract.
func bindERC6551Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC6551RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC6551Registry *ERC6551RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC6551Registry.Contract.ERC6551RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC6551Registry *ERC6551RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.ERC6551RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC6551Registry *ERC6551RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.ERC6551RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC6551Registry *ERC6551RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC6551Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC6551Registry *ERC6551RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC6551Registry *ERC6551RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x5e9bc536.
//
// Solidity: function account(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt) view returns(address)
func (_ERC6551Registry *ERC6551RegistryCaller) Account(opts *bind.CallOpts, implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC6551Registry.contract.Call(opts, &out, "account", implementation, chainId, tokenContract, tokenId, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x5e9bc536.
//
// Solidity: function account(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt) view returns(address)
func (_ERC6551Registry *ERC6551RegistrySession) Account(implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int) (common.Address, error) {
	return _ERC6551Registry.Contract.Account(&_ERC6551Registry.CallOpts, implementation, chainId, tokenContract, tokenId, salt)
}

// Account is a free data retrieval call binding the contract method 0x5e9bc536.
//
// Solidity: function account(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt) view returns(address)
func (_ERC6551Registry *ERC6551RegistryCallerSession) Account(implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int) (common.Address, error) {
	return _ERC6551Registry.Contract.Account(&_ERC6551Registry.CallOpts, implementation, chainId, tokenContract, tokenId, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xda7323b3.
//
// Solidity: function createAccount(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt, bytes initData) returns(address)
func (_ERC6551Registry *ERC6551RegistryTransactor) CreateAccount(opts *bind.TransactOpts, implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int, initData []byte) (*types.Transaction, error) {
	return _ERC6551Registry.contract.Transact(opts, "createAccount", implementation, chainId, tokenContract, tokenId, salt, initData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xda7323b3.
//
// Solidity: function createAccount(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt, bytes initData) returns(address)
func (_ERC6551Registry *ERC6551RegistrySession) CreateAccount(implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int, initData []byte) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.CreateAccount(&_ERC6551Registry.TransactOpts, implementation, chainId, tokenContract, tokenId, salt, initData)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xda7323b3.
//
// Solidity: function createAccount(address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt, bytes initData) returns(address)
func (_ERC6551Registry *ERC6551RegistryTransactorSession) CreateAccount(implementation common.Address, chainId *big.Int, tokenContract common.Address, tokenId *big.Int, salt *big.Int, initData []byte) (*types.Transaction, error) {
	return _ERC6551Registry.Contract.CreateAccount(&_ERC6551Registry.TransactOpts, implementation, chainId, tokenContract, tokenId, salt, initData)
}

// ERC6551RegistryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the ERC6551Registry contract.
type ERC6551RegistryAccountCreatedIterator struct {
	Event *ERC6551RegistryAccountCreated // Event containing the contract specifics and raw log

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
func (it *ERC6551RegistryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC6551RegistryAccountCreated)
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
		it.Event = new(ERC6551RegistryAccountCreated)
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
func (it *ERC6551RegistryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC6551RegistryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC6551RegistryAccountCreated represents a AccountCreated event raised by the ERC6551Registry contract.
type ERC6551RegistryAccountCreated struct {
	Account        common.Address
	Implementation common.Address
	ChainId        *big.Int
	TokenContract  common.Address
	TokenId        *big.Int
	Salt           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x07fba7bba1191da7ee1155dcfa0030701c9c9a9cc34a93b991fc6fd0c9268d8f.
//
// Solidity: event AccountCreated(address account, address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt)
func (_ERC6551Registry *ERC6551RegistryFilterer) FilterAccountCreated(opts *bind.FilterOpts) (*ERC6551RegistryAccountCreatedIterator, error) {

	logs, sub, err := _ERC6551Registry.contract.FilterLogs(opts, "AccountCreated")
	if err != nil {
		return nil, err
	}
	return &ERC6551RegistryAccountCreatedIterator{contract: _ERC6551Registry.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x07fba7bba1191da7ee1155dcfa0030701c9c9a9cc34a93b991fc6fd0c9268d8f.
//
// Solidity: event AccountCreated(address account, address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt)
func (_ERC6551Registry *ERC6551RegistryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *ERC6551RegistryAccountCreated) (event.Subscription, error) {

	logs, sub, err := _ERC6551Registry.contract.WatchLogs(opts, "AccountCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC6551RegistryAccountCreated)
				if err := _ERC6551Registry.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x07fba7bba1191da7ee1155dcfa0030701c9c9a9cc34a93b991fc6fd0c9268d8f.
//
// Solidity: event AccountCreated(address account, address implementation, uint256 chainId, address tokenContract, uint256 tokenId, uint256 salt)
func (_ERC6551Registry *ERC6551RegistryFilterer) ParseAccountCreated(log types.Log) (*ERC6551RegistryAccountCreated, error) {
	event := new(ERC6551RegistryAccountCreated)
	if err := _ERC6551Registry.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
