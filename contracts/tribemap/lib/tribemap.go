// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tribemaplib

import (
	"math/big"
	"strings"

	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi"
	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi/bind"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/core/types"
)

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405260018054600160a060020a031916905534801561002057600080fd5b5060008054600160a060020a031916331790556101e9806100426000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046379ba509781146100625780638da5cb5b1461006c578063a6f9dae114610090575b600080fd5b61006a6100b6565b005b61007461014d565b60408051600160a060020a039092168252519081900360200190f35b61006a600480360360208110156100a657600080fd5b5035600160a060020a031661015c565b600154600160a060020a031633146100cd57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600054600160a060020a0316331461017357600080fd5b600054600160a060020a038281169116141561018e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a7230582046f053756e5db33b9acf5cf97e4803a19cf70ff97aa91ef2efbff6f5e9113ce60029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned                  // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller            // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Owned *OwnedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Owned.Contract.AcceptOwnership(&_Owned.TransactOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Owned *OwnedTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.ChangeOwner(&_Owned.TransactOpts, _newOwner)
}

// TribemapABI is the input ABI used to generate the binding from.
const TribemapABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"delAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"existAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"makeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"makeCoinLog\",\"outputs\":[{\"name\":\"numList\",\"type\":\"uint256[]\"},{\"name\":\"coinList\",\"type\":\"uint256[]\"},{\"name\":\"toList\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"addAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// TribemapBin is the compiled bytecode used for deploying new contracts.
const TribemapBin = `0x608060405260018054600160a060020a031990811690915560008054909116331790556107ae806100316000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c0100000000000000000000000000000000000000000000000000000000900480639dfad1e8116100785780639dfad1e8146101b3578063a6f9dae1146101df578063b3ff4b0614610205578063ba4c18db146102eb576100a5565b80635022edf8146100aa57806379ba50971461014f57806380b7069d146101575780638da5cb5b1461018f575b600080fd5b61014d600480360360208110156100c057600080fd5b8101906020810181356401000000008111156100db57600080fd5b8201836020820111156100ed57600080fd5b8035906020019184602083028401116401000000008311171561010f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610390945050505050565b005b61014d6103f8565b61017d6004803603602081101561016d57600080fd5b5035600160a060020a031661048f565b60408051918252519081900360200190f35b6101976104aa565b60408051600160a060020a039092168252519081900360200190f35b61014d600480360360408110156101c957600080fd5b50600160a060020a0381351690602001356104b9565b61014d600480360360208110156101f557600080fd5b5035600160a060020a031661054d565b61020d6105ae565b60405180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561025557818101518382015260200161023d565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561029457818101518382015260200161027c565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156102d35781810151838201526020016102bb565b50505050905001965050505050505060405180910390f35b61014d6004803603604081101561030157600080fd5b81019060208101813564010000000081111561031c57600080fd5b82018360208201111561032e57600080fd5b8035906020019184602083028401116401000000008311171561035057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550509135925061071a915050565b600054600160a060020a031633146103a757600080fd5b805160005b818110156103f35760006002600085848151811015156103c857fe5b6020908102909101810151600160a060020a03168252810191909152604001600020556001016103ac565b505050565b600154600160a060020a0316331461040f57600080fd5b60005460015460408051600160a060020a03938416815292909116602083015280517f343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a9281900390910190a1600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600160a060020a031660009081526002602052604090205490565b600054600160a060020a031681565b600054600160a060020a031633146104d057600080fd5b43600081815260036020526040812060028101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03969096169590951790945591835560019283018190556004805493840181559091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90910155565b600054600160a060020a0316331461056457600080fd5b600054600160a060020a038281169116141561057f57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60608060606000600480549050111561071557600480548060200260200160405190810160405280929190818152602001828054801561060d57602002820191906000526020600020905b8154815260200190600101908083116105f9575b50505050509250600480549050604051908082528060200260200182016040528015610643578160200160208202803883390190505b506004546040805182815260208084028201019091529193508015610672578160200160208202803883390190505b50905060005b60045481101561071357600060048281548110151561069357fe5b600091825260208083209091015480835260039091526040909120600201548451919250600160a060020a0316908490849081106106cd57fe5b600160a060020a039092166020928302909101820152600082815260039091526040902054845185908490811061070057fe5b6020908102909101015250600101610678565b505b909192565b600054600160a060020a0316331461073157600080fd5b815160005b8181101561077c578260026000868481518110151561075157fe5b6020908102909101810151600160a060020a0316825281019190915260400160002055600101610736565b5050505056fea165627a7a72305820a7d21e2ab82c98a9cea36b8d75a5453c5f3935f062293417038ae87ccf16b0c00029`

// DeployTribemap deploys a new Ethereum contract, binding an instance of Tribemap to it.
func DeployTribemap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tribemap, error) {
	parsed, err := abi.JSON(strings.NewReader(TribemapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TribemapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tribemap{TribemapCaller: TribemapCaller{contract: contract}, TribemapTransactor: TribemapTransactor{contract: contract}}, nil
}

// Tribemap is an auto generated Go binding around an Ethereum contract.
type Tribemap struct {
	TribemapCaller     // Read-only binding to the contract
	TribemapTransactor // Write-only binding to the contract
}

// TribemapCaller is an auto generated read-only Go binding around an Ethereum contract.
type TribemapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribemapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TribemapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TribemapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TribemapSession struct {
	Contract     *Tribemap               // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TribemapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TribemapCallerSession struct {
	Contract *TribemapCaller         // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// TribemapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TribemapTransactorSession struct {
	Contract     *TribemapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TribemapRaw is an auto generated low-level Go binding around an Ethereum contract.
type TribemapRaw struct {
	Contract *Tribemap // Generic contract binding to access the raw methods on
}

// TribemapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TribemapCallerRaw struct {
	Contract *TribemapCaller // Generic read-only contract binding to access the raw methods on
}

// TribemapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TribemapTransactorRaw struct {
	Contract *TribemapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTribemap creates a new instance of Tribemap, bound to a specific deployed contract.
func NewTribemap(address common.Address, backend bind.ContractBackend) (*Tribemap, error) {
	contract, err := bindTribemap(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tribemap{TribemapCaller: TribemapCaller{contract: contract}, TribemapTransactor: TribemapTransactor{contract: contract}}, nil
}

// NewTribemapCaller creates a new read-only instance of Tribemap, bound to a specific deployed contract.
func NewTribemapCaller(address common.Address, caller bind.ContractCaller) (*TribemapCaller, error) {
	contract, err := bindTribemap(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TribemapCaller{contract: contract}, nil
}

// NewTribemapTransactor creates a new write-only instance of Tribemap, bound to a specific deployed contract.
func NewTribemapTransactor(address common.Address, transactor bind.ContractTransactor) (*TribemapTransactor, error) {
	contract, err := bindTribemap(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TribemapTransactor{contract: contract}, nil
}

// bindTribemap binds a generic wrapper to an already deployed contract.
func bindTribemap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TribemapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tribemap *TribemapRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Tribemap.Contract.TribemapCaller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tribemap *TribemapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tribemap.Contract.TribemapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tribemap *TribemapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tribemap.Contract.TribemapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tribemap *TribemapCallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Tribemap.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tribemap *TribemapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tribemap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tribemap *TribemapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tribemap.Contract.contract.Transact(opts, method, params...)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_Tribemap *TribemapCaller) ExistAddress(opts *bind.CallOptsWithNumber, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tribemap.contract.CallWithNumber(opts, out, "existAddress", _owner)
	return *ret0, err
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_Tribemap *TribemapSession) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _Tribemap.Contract.ExistAddress(&_Tribemap.CallOpts, _owner)
}

// ExistAddress is a free data retrieval call binding the contract method 0x80b7069d.
//
// Solidity: function existAddress(_owner address) constant returns(uint256)
func (_Tribemap *TribemapCallerSession) ExistAddress(_owner common.Address) (*big.Int, error) {
	return _Tribemap.Contract.ExistAddress(&_Tribemap.CallOpts, _owner)
}

// MakeCoinLog is a free data retrieval call binding the contract method 0xb3ff4b06.
//
// Solidity: function makeCoinLog() constant returns(numList uint256[], coinList uint256[], toList address[])
func (_Tribemap *TribemapCaller) MakeCoinLog(opts *bind.CallOptsWithNumber) (struct {
	NumList  []*big.Int
	CoinList []*big.Int
	ToList   []common.Address
}, error) {
	ret := new(struct {
		NumList  []*big.Int
		CoinList []*big.Int
		ToList   []common.Address
	})
	out := ret
	err := _Tribemap.contract.CallWithNumber(opts, out, "makeCoinLog")
	return *ret, err
}

// MakeCoinLog is a free data retrieval call binding the contract method 0xb3ff4b06.
//
// Solidity: function makeCoinLog() constant returns(numList uint256[], coinList uint256[], toList address[])
func (_Tribemap *TribemapSession) MakeCoinLog() (struct {
	NumList  []*big.Int
	CoinList []*big.Int
	ToList   []common.Address
}, error) {
	return _Tribemap.Contract.MakeCoinLog(&_Tribemap.CallOpts)
}

// MakeCoinLog is a free data retrieval call binding the contract method 0xb3ff4b06.
//
// Solidity: function makeCoinLog() constant returns(numList uint256[], coinList uint256[], toList address[])
func (_Tribemap *TribemapCallerSession) MakeCoinLog() (struct {
	NumList  []*big.Int
	CoinList []*big.Int
	ToList   []common.Address
}, error) {
	return _Tribemap.Contract.MakeCoinLog(&_Tribemap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tribemap *TribemapCaller) Owner(opts *bind.CallOptsWithNumber) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tribemap.contract.CallWithNumber(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tribemap *TribemapSession) Owner() (common.Address, error) {
	return _Tribemap.Contract.Owner(&_Tribemap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tribemap *TribemapCallerSession) Owner() (common.Address, error) {
	return _Tribemap.Contract.Owner(&_Tribemap.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tribemap *TribemapTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tribemap.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tribemap *TribemapSession) AcceptOwnership() (*types.Transaction, error) {
	return _Tribemap.Contract.AcceptOwnership(&_Tribemap.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Tribemap *TribemapTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Tribemap.Contract.AcceptOwnership(&_Tribemap.TransactOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_Tribemap *TribemapTransactor) AddAddress(opts *bind.TransactOpts, _owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _Tribemap.contract.Transact(opts, "addAddress", _owners, version)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_Tribemap *TribemapSession) AddAddress(_owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _Tribemap.Contract.AddAddress(&_Tribemap.TransactOpts, _owners, version)
}

// AddAddress is a paid mutator transaction binding the contract method 0xba4c18db.
//
// Solidity: function addAddress(_owners address[], version uint256) returns()
func (_Tribemap *TribemapTransactorSession) AddAddress(_owners []common.Address, version *big.Int) (*types.Transaction, error) {
	return _Tribemap.Contract.AddAddress(&_Tribemap.TransactOpts, _owners, version)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Tribemap *TribemapTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Tribemap.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Tribemap *TribemapSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Tribemap.Contract.ChangeOwner(&_Tribemap.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Tribemap *TribemapTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Tribemap.Contract.ChangeOwner(&_Tribemap.TransactOpts, _newOwner)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_Tribemap *TribemapTransactor) DelAddress(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _Tribemap.contract.Transact(opts, "delAddress", _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_Tribemap *TribemapSession) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _Tribemap.Contract.DelAddress(&_Tribemap.TransactOpts, _owners)
}

// DelAddress is a paid mutator transaction binding the contract method 0x5022edf8.
//
// Solidity: function delAddress(_owners address[]) returns()
func (_Tribemap *TribemapTransactorSession) DelAddress(_owners []common.Address) (*types.Transaction, error) {
	return _Tribemap.Contract.DelAddress(&_Tribemap.TransactOpts, _owners)
}

// MakeCoin is a paid mutator transaction binding the contract method 0x9dfad1e8.
//
// Solidity: function makeCoin(to address, val uint256) returns()
func (_Tribemap *TribemapTransactor) MakeCoin(opts *bind.TransactOpts, to common.Address, val *big.Int) (*types.Transaction, error) {
	return _Tribemap.contract.Transact(opts, "makeCoin", to, val)
}

// MakeCoin is a paid mutator transaction binding the contract method 0x9dfad1e8.
//
// Solidity: function makeCoin(to address, val uint256) returns()
func (_Tribemap *TribemapSession) MakeCoin(to common.Address, val *big.Int) (*types.Transaction, error) {
	return _Tribemap.Contract.MakeCoin(&_Tribemap.TransactOpts, to, val)
}

// MakeCoin is a paid mutator transaction binding the contract method 0x9dfad1e8.
//
// Solidity: function makeCoin(to address, val uint256) returns()
func (_Tribemap *TribemapTransactorSession) MakeCoin(to common.Address, val *big.Int) (*types.Transaction, error) {
	return _Tribemap.Contract.MakeCoin(&_Tribemap.TransactOpts, to, val)
}
