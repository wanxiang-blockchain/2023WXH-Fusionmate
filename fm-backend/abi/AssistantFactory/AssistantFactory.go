// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AssistantFactory

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

// AssistantFactoryMetaData contains all meta data concerning the AssistantFactory contract.
var AssistantFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fmToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountImpl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"astBot\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NewAssistantCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assistantMaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assistantsMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"createAssistant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fmToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"harvestVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AssistantFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AssistantFactoryMetaData.ABI instead.
var AssistantFactoryABI = AssistantFactoryMetaData.ABI

// AssistantFactory is an auto generated Go binding around an Ethereum contract.
type AssistantFactory struct {
	AssistantFactoryCaller     // Read-only binding to the contract
	AssistantFactoryTransactor // Write-only binding to the contract
	AssistantFactoryFilterer   // Log filterer for contract events
}

// AssistantFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssistantFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssistantFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssistantFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssistantFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssistantFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssistantFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssistantFactorySession struct {
	Contract     *AssistantFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssistantFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssistantFactoryCallerSession struct {
	Contract *AssistantFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AssistantFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssistantFactoryTransactorSession struct {
	Contract     *AssistantFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AssistantFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssistantFactoryRaw struct {
	Contract *AssistantFactory // Generic contract binding to access the raw methods on
}

// AssistantFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssistantFactoryCallerRaw struct {
	Contract *AssistantFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AssistantFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssistantFactoryTransactorRaw struct {
	Contract *AssistantFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssistantFactory creates a new instance of AssistantFactory, bound to a specific deployed contract.
func NewAssistantFactory(address common.Address, backend bind.ContractBackend) (*AssistantFactory, error) {
	contract, err := bindAssistantFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssistantFactory{AssistantFactoryCaller: AssistantFactoryCaller{contract: contract}, AssistantFactoryTransactor: AssistantFactoryTransactor{contract: contract}, AssistantFactoryFilterer: AssistantFactoryFilterer{contract: contract}}, nil
}

// NewAssistantFactoryCaller creates a new read-only instance of AssistantFactory, bound to a specific deployed contract.
func NewAssistantFactoryCaller(address common.Address, caller bind.ContractCaller) (*AssistantFactoryCaller, error) {
	contract, err := bindAssistantFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssistantFactoryCaller{contract: contract}, nil
}

// NewAssistantFactoryTransactor creates a new write-only instance of AssistantFactory, bound to a specific deployed contract.
func NewAssistantFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AssistantFactoryTransactor, error) {
	contract, err := bindAssistantFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssistantFactoryTransactor{contract: contract}, nil
}

// NewAssistantFactoryFilterer creates a new log filterer instance of AssistantFactory, bound to a specific deployed contract.
func NewAssistantFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AssistantFactoryFilterer, error) {
	contract, err := bindAssistantFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssistantFactoryFilterer{contract: contract}, nil
}

// bindAssistantFactory binds a generic wrapper to an already deployed contract.
func bindAssistantFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssistantFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssistantFactory *AssistantFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssistantFactory.Contract.AssistantFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssistantFactory *AssistantFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssistantFactory.Contract.AssistantFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssistantFactory *AssistantFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssistantFactory.Contract.AssistantFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssistantFactory *AssistantFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssistantFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssistantFactory *AssistantFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssistantFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssistantFactory *AssistantFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssistantFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImpl is a free data retrieval call binding the contract method 0x8804eb9f.
//
// Solidity: function accountImpl() view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) AccountImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "accountImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImpl is a free data retrieval call binding the contract method 0x8804eb9f.
//
// Solidity: function accountImpl() view returns(address)
func (_AssistantFactory *AssistantFactorySession) AccountImpl() (common.Address, error) {
	return _AssistantFactory.Contract.AccountImpl(&_AssistantFactory.CallOpts)
}

// AccountImpl is a free data retrieval call binding the contract method 0x8804eb9f.
//
// Solidity: function accountImpl() view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) AccountImpl() (common.Address, error) {
	return _AssistantFactory.Contract.AccountImpl(&_AssistantFactory.CallOpts)
}

// AssistantMaker is a free data retrieval call binding the contract method 0xb3bd770e.
//
// Solidity: function assistantMaker(address ) view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) AssistantMaker(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "assistantMaker", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssistantMaker is a free data retrieval call binding the contract method 0xb3bd770e.
//
// Solidity: function assistantMaker(address ) view returns(address)
func (_AssistantFactory *AssistantFactorySession) AssistantMaker(arg0 common.Address) (common.Address, error) {
	return _AssistantFactory.Contract.AssistantMaker(&_AssistantFactory.CallOpts, arg0)
}

// AssistantMaker is a free data retrieval call binding the contract method 0xb3bd770e.
//
// Solidity: function assistantMaker(address ) view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) AssistantMaker(arg0 common.Address) (common.Address, error) {
	return _AssistantFactory.Contract.AssistantMaker(&_AssistantFactory.CallOpts, arg0)
}

// AssistantsMap is a free data retrieval call binding the contract method 0xf42889fe.
//
// Solidity: function assistantsMap(uint256 ) view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) AssistantsMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "assistantsMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssistantsMap is a free data retrieval call binding the contract method 0xf42889fe.
//
// Solidity: function assistantsMap(uint256 ) view returns(address)
func (_AssistantFactory *AssistantFactorySession) AssistantsMap(arg0 *big.Int) (common.Address, error) {
	return _AssistantFactory.Contract.AssistantsMap(&_AssistantFactory.CallOpts, arg0)
}

// AssistantsMap is a free data retrieval call binding the contract method 0xf42889fe.
//
// Solidity: function assistantsMap(uint256 ) view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) AssistantsMap(arg0 *big.Int) (common.Address, error) {
	return _AssistantFactory.Contract.AssistantsMap(&_AssistantFactory.CallOpts, arg0)
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) Backend(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "backend")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_AssistantFactory *AssistantFactorySession) Backend() (common.Address, error) {
	return _AssistantFactory.Contract.Backend(&_AssistantFactory.CallOpts)
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) Backend() (common.Address, error) {
	return _AssistantFactory.Contract.Backend(&_AssistantFactory.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AssistantFactory *AssistantFactoryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AssistantFactory *AssistantFactorySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AssistantFactory.Contract.Eip712Domain(&_AssistantFactory.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AssistantFactory *AssistantFactoryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AssistantFactory.Contract.Eip712Domain(&_AssistantFactory.CallOpts)
}

// FmToken is a free data retrieval call binding the contract method 0x69e34381.
//
// Solidity: function fmToken() view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) FmToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "fmToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FmToken is a free data retrieval call binding the contract method 0x69e34381.
//
// Solidity: function fmToken() view returns(address)
func (_AssistantFactory *AssistantFactorySession) FmToken() (common.Address, error) {
	return _AssistantFactory.Contract.FmToken(&_AssistantFactory.CallOpts)
}

// FmToken is a free data retrieval call binding the contract method 0x69e34381.
//
// Solidity: function fmToken() view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) FmToken() (common.Address, error) {
	return _AssistantFactory.Contract.FmToken(&_AssistantFactory.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xce88b145.
//
// Solidity: function getAccount(uint256 _tokenId) view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) GetAccount(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "getAccount", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xce88b145.
//
// Solidity: function getAccount(uint256 _tokenId) view returns(address)
func (_AssistantFactory *AssistantFactorySession) GetAccount(_tokenId *big.Int) (common.Address, error) {
	return _AssistantFactory.Contract.GetAccount(&_AssistantFactory.CallOpts, _tokenId)
}

// GetAccount is a free data retrieval call binding the contract method 0xce88b145.
//
// Solidity: function getAccount(uint256 _tokenId) view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) GetAccount(_tokenId *big.Int) (common.Address, error) {
	return _AssistantFactory.Contract.GetAccount(&_AssistantFactory.CallOpts, _tokenId)
}

// HarvestVerify is a free data retrieval call binding the contract method 0x541a3ab7.
//
// Solidity: function harvestVerify(bytes _signature, uint256 _amount, uint256 _collectionId, uint256 _tokenId) view returns(bool)
func (_AssistantFactory *AssistantFactoryCaller) HarvestVerify(opts *bind.CallOpts, _signature []byte, _amount *big.Int, _collectionId *big.Int, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "harvestVerify", _signature, _amount, _collectionId, _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HarvestVerify is a free data retrieval call binding the contract method 0x541a3ab7.
//
// Solidity: function harvestVerify(bytes _signature, uint256 _amount, uint256 _collectionId, uint256 _tokenId) view returns(bool)
func (_AssistantFactory *AssistantFactorySession) HarvestVerify(_signature []byte, _amount *big.Int, _collectionId *big.Int, _tokenId *big.Int) (bool, error) {
	return _AssistantFactory.Contract.HarvestVerify(&_AssistantFactory.CallOpts, _signature, _amount, _collectionId, _tokenId)
}

// HarvestVerify is a free data retrieval call binding the contract method 0x541a3ab7.
//
// Solidity: function harvestVerify(bytes _signature, uint256 _amount, uint256 _collectionId, uint256 _tokenId) view returns(bool)
func (_AssistantFactory *AssistantFactoryCallerSession) HarvestVerify(_signature []byte, _amount *big.Int, _collectionId *big.Int, _tokenId *big.Int) (bool, error) {
	return _AssistantFactory.Contract.HarvestVerify(&_AssistantFactory.CallOpts, _signature, _amount, _collectionId, _tokenId)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AssistantFactory *AssistantFactoryCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AssistantFactory *AssistantFactorySession) Registry() (common.Address, error) {
	return _AssistantFactory.Contract.Registry(&_AssistantFactory.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AssistantFactory *AssistantFactoryCallerSession) Registry() (common.Address, error) {
	return _AssistantFactory.Contract.Registry(&_AssistantFactory.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xaa2f5095.
//
// Solidity: function verify(bytes _signature, string _name, string _symbol, string _baseURI, address makerAddress, uint256 collectionId, uint256 maxSupply, uint256 mintPrice) view returns(bool)
func (_AssistantFactory *AssistantFactoryCaller) Verify(opts *bind.CallOpts, _signature []byte, _name string, _symbol string, _baseURI string, makerAddress common.Address, collectionId *big.Int, maxSupply *big.Int, mintPrice *big.Int) (bool, error) {
	var out []interface{}
	err := _AssistantFactory.contract.Call(opts, &out, "verify", _signature, _name, _symbol, _baseURI, makerAddress, collectionId, maxSupply, mintPrice)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xaa2f5095.
//
// Solidity: function verify(bytes _signature, string _name, string _symbol, string _baseURI, address makerAddress, uint256 collectionId, uint256 maxSupply, uint256 mintPrice) view returns(bool)
func (_AssistantFactory *AssistantFactorySession) Verify(_signature []byte, _name string, _symbol string, _baseURI string, makerAddress common.Address, collectionId *big.Int, maxSupply *big.Int, mintPrice *big.Int) (bool, error) {
	return _AssistantFactory.Contract.Verify(&_AssistantFactory.CallOpts, _signature, _name, _symbol, _baseURI, makerAddress, collectionId, maxSupply, mintPrice)
}

// Verify is a free data retrieval call binding the contract method 0xaa2f5095.
//
// Solidity: function verify(bytes _signature, string _name, string _symbol, string _baseURI, address makerAddress, uint256 collectionId, uint256 maxSupply, uint256 mintPrice) view returns(bool)
func (_AssistantFactory *AssistantFactoryCallerSession) Verify(_signature []byte, _name string, _symbol string, _baseURI string, makerAddress common.Address, collectionId *big.Int, maxSupply *big.Int, mintPrice *big.Int) (bool, error) {
	return _AssistantFactory.Contract.Verify(&_AssistantFactory.CallOpts, _signature, _name, _symbol, _baseURI, makerAddress, collectionId, maxSupply, mintPrice)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xcab13915.
//
// Solidity: function createAccount(uint256 _tokenId) returns(address)
func (_AssistantFactory *AssistantFactoryTransactor) CreateAccount(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.contract.Transact(opts, "createAccount", _tokenId)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xcab13915.
//
// Solidity: function createAccount(uint256 _tokenId) returns(address)
func (_AssistantFactory *AssistantFactorySession) CreateAccount(_tokenId *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.Contract.CreateAccount(&_AssistantFactory.TransactOpts, _tokenId)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xcab13915.
//
// Solidity: function createAccount(uint256 _tokenId) returns(address)
func (_AssistantFactory *AssistantFactoryTransactorSession) CreateAccount(_tokenId *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.Contract.CreateAccount(&_AssistantFactory.TransactOpts, _tokenId)
}

// CreateAssistant is a paid mutator transaction binding the contract method 0x07a42ac5.
//
// Solidity: function createAssistant(bytes _signature, string _name, string _symbol, string _baseURI, uint256 _collectionId, uint256 _maxSupply, uint256 _mintPrice) returns()
func (_AssistantFactory *AssistantFactoryTransactor) CreateAssistant(opts *bind.TransactOpts, _signature []byte, _name string, _symbol string, _baseURI string, _collectionId *big.Int, _maxSupply *big.Int, _mintPrice *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.contract.Transact(opts, "createAssistant", _signature, _name, _symbol, _baseURI, _collectionId, _maxSupply, _mintPrice)
}

// CreateAssistant is a paid mutator transaction binding the contract method 0x07a42ac5.
//
// Solidity: function createAssistant(bytes _signature, string _name, string _symbol, string _baseURI, uint256 _collectionId, uint256 _maxSupply, uint256 _mintPrice) returns()
func (_AssistantFactory *AssistantFactorySession) CreateAssistant(_signature []byte, _name string, _symbol string, _baseURI string, _collectionId *big.Int, _maxSupply *big.Int, _mintPrice *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.Contract.CreateAssistant(&_AssistantFactory.TransactOpts, _signature, _name, _symbol, _baseURI, _collectionId, _maxSupply, _mintPrice)
}

// CreateAssistant is a paid mutator transaction binding the contract method 0x07a42ac5.
//
// Solidity: function createAssistant(bytes _signature, string _name, string _symbol, string _baseURI, uint256 _collectionId, uint256 _maxSupply, uint256 _mintPrice) returns()
func (_AssistantFactory *AssistantFactoryTransactorSession) CreateAssistant(_signature []byte, _name string, _symbol string, _baseURI string, _collectionId *big.Int, _maxSupply *big.Int, _mintPrice *big.Int) (*types.Transaction, error) {
	return _AssistantFactory.Contract.CreateAssistant(&_AssistantFactory.TransactOpts, _signature, _name, _symbol, _baseURI, _collectionId, _maxSupply, _mintPrice)
}

// AssistantFactoryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AssistantFactory contract.
type AssistantFactoryEIP712DomainChangedIterator struct {
	Event *AssistantFactoryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AssistantFactoryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssistantFactoryEIP712DomainChanged)
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
		it.Event = new(AssistantFactoryEIP712DomainChanged)
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
func (it *AssistantFactoryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssistantFactoryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssistantFactoryEIP712DomainChanged represents a EIP712DomainChanged event raised by the AssistantFactory contract.
type AssistantFactoryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AssistantFactory *AssistantFactoryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AssistantFactoryEIP712DomainChangedIterator, error) {

	logs, sub, err := _AssistantFactory.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AssistantFactoryEIP712DomainChangedIterator{contract: _AssistantFactory.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AssistantFactory *AssistantFactoryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AssistantFactoryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AssistantFactory.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssistantFactoryEIP712DomainChanged)
				if err := _AssistantFactory.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AssistantFactory *AssistantFactoryFilterer) ParseEIP712DomainChanged(log types.Log) (*AssistantFactoryEIP712DomainChanged, error) {
	event := new(AssistantFactoryEIP712DomainChanged)
	if err := _AssistantFactory.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssistantFactoryNewAssistantCreatedIterator is returned from FilterNewAssistantCreated and is used to iterate over the raw logs and unpacked data for NewAssistantCreated events raised by the AssistantFactory contract.
type AssistantFactoryNewAssistantCreatedIterator struct {
	Event *AssistantFactoryNewAssistantCreated // Event containing the contract specifics and raw log

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
func (it *AssistantFactoryNewAssistantCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssistantFactoryNewAssistantCreated)
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
		it.Event = new(AssistantFactoryNewAssistantCreated)
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
func (it *AssistantFactoryNewAssistantCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssistantFactoryNewAssistantCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssistantFactoryNewAssistantCreated represents a NewAssistantCreated event raised by the AssistantFactory contract.
type AssistantFactoryNewAssistantCreated struct {
	Maker  common.Address
	AstBot common.Address
	NftId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewAssistantCreated is a free log retrieval operation binding the contract event 0xcb781daaebdd4349dbc7e45f9497a25ab7cc731f5f2c74c6fa9100958683a607.
//
// Solidity: event NewAssistantCreated(address indexed maker, address indexed astBot, uint256 indexed nftId)
func (_AssistantFactory *AssistantFactoryFilterer) FilterNewAssistantCreated(opts *bind.FilterOpts, maker []common.Address, astBot []common.Address, nftId []*big.Int) (*AssistantFactoryNewAssistantCreatedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var astBotRule []interface{}
	for _, astBotItem := range astBot {
		astBotRule = append(astBotRule, astBotItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _AssistantFactory.contract.FilterLogs(opts, "NewAssistantCreated", makerRule, astBotRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return &AssistantFactoryNewAssistantCreatedIterator{contract: _AssistantFactory.contract, event: "NewAssistantCreated", logs: logs, sub: sub}, nil
}

// WatchNewAssistantCreated is a free log subscription operation binding the contract event 0xcb781daaebdd4349dbc7e45f9497a25ab7cc731f5f2c74c6fa9100958683a607.
//
// Solidity: event NewAssistantCreated(address indexed maker, address indexed astBot, uint256 indexed nftId)
func (_AssistantFactory *AssistantFactoryFilterer) WatchNewAssistantCreated(opts *bind.WatchOpts, sink chan<- *AssistantFactoryNewAssistantCreated, maker []common.Address, astBot []common.Address, nftId []*big.Int) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var astBotRule []interface{}
	for _, astBotItem := range astBot {
		astBotRule = append(astBotRule, astBotItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _AssistantFactory.contract.WatchLogs(opts, "NewAssistantCreated", makerRule, astBotRule, nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssistantFactoryNewAssistantCreated)
				if err := _AssistantFactory.contract.UnpackLog(event, "NewAssistantCreated", log); err != nil {
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

// ParseNewAssistantCreated is a log parse operation binding the contract event 0xcb781daaebdd4349dbc7e45f9497a25ab7cc731f5f2c74c6fa9100958683a607.
//
// Solidity: event NewAssistantCreated(address indexed maker, address indexed astBot, uint256 indexed nftId)
func (_AssistantFactory *AssistantFactoryFilterer) ParseNewAssistantCreated(log types.Log) (*AssistantFactoryNewAssistantCreated, error) {
	event := new(AssistantFactoryNewAssistantCreated)
	if err := _AssistantFactory.contract.UnpackLog(event, "NewAssistantCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
