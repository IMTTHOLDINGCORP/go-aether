// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registerlib

import (
	"strings"

	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi"
	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi/bind"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/core/types"
)

// Register_0_0_1ABI is the input ABI used to generate the binding from.
const Register_0_0_1ABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"appendBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"verifyBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"append\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVsn\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getId\",\"outputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_cb\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"hasId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"rawdata\",\"type\":\"string\"}],\"name\":\"appendId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Register_0_0_1Bin is the compiled bytecode used for deploying new contracts.
const Register_0_0_1Bin = `0x606060405260408051908101604052600581527f302e302e310000000000000000000000000000000000000000000000000000006020820152600090805161004b92916020019061005c565b50341561005757600080fd5b6100f7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009d57805160ff19168380011785556100ca565b828001600101855582156100ca579182015b828111156100ca5782518255916020019190600101906100af565b506100d69291506100da565b5090565b6100f491905b808211156100d657600081556001016100e0565b90565b610af8806101066000396000f3006060604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166319e5bf3a81146100be57806320accbe6146100d357806329092d0e146101245780633c094c4c1461014357806363a9c3d7146101a85780638c25bc12146101c7578063ac04f5a714610218578063b6ce614314610237578063bee51f3b146102c1578063cb72361c14610399578063ebd746c9146103ea578063f896cf101461043b575b600080fd5b34156100c957600080fd5b6100d16104ce565b005b34156100de57600080fd5b6100d160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061051895505050505050565b341561012f57600080fd5b6100d1600160a060020a03600435166105af565b341561014e57600080fd5b61019460046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506105ed95505050505050565b604051901515815260200160405180910390f35b34156101b357600080fd5b610194600160a060020a0360043516610661565b34156101d257600080fd5b6100d160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061067f95505050505050565b341561022357600080fd5b6100d1600160a060020a036004351661070d565b341561024257600080fd5b61024a61074e565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561028657808201518382015260200161026e565b50505050905090810190601f1680156102b35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102cc57600080fd5b61031260046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107f695505050505050565b604051600160a060020a038216602082015260408082528190810184818151815260200191508051906020019080838360005b8381101561035d578082015183820152602001610345565b50505050905090810190601f16801561038a5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34156103a457600080fd5b61019460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061087995505050505050565b34156103f557600080fd5b6100d160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061090195505050505050565b341561044657600080fd5b6100d160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506109a295505050505050565b6001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a0390811691909117808355166000908152600260205260409020805460ff19169091179055565b6001543390600160a060020a0380831691161461053457600080fd5b60016003836040518082805190602001908083835b602083106105685780518252601f199092019160209182019101610549565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805460ff19169115159190911790555050565b6001543390600160a060020a038083169116146105cb57600080fd5b50600160a060020a03166000908152600260205260409020805460ff19169055565b60006003826040518082805190602001908083835b602083106106215780518252601f199092019160209182019101610602565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205460ff1690505b919050565b600160a060020a031660009081526002602052604090205460ff1690565b6001543390600160a060020a0380831691161461069b57600080fd5b6003826040518082805190602001908083835b602083106106cd5780518252601f1990920191602091820191016106ae565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805460ff191690555050565b6001543390600160a060020a0380831691161461072957600080fd5b50600160a060020a03166000908152600260205260409020805460ff19166001179055565b610756610aba565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107ec5780601f106107c1576101008083540402835291602001916107ec565b820191906000526020600020905b8154815290600101906020018083116107cf57829003601f168201915b5050505050905090565b6107fe610aba565b508060006004826040518082805190602001908083835b602083106108345780518252601f199092019160209182019101610815565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054600160a060020a03169050915091565b60006004826040518082805190602001908083835b602083106108ad5780518252601f19909201916020918201910161088e565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054600160a060020a031615156108f95750600061065c565b506001919050565b6001543390600160a060020a0380831691161461091d57600080fd5b6004826040518082805190602001908083835b6020831061094f5780518252601f199092019160209182019101610930565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805473ffffffffffffffffffffffffffffffffffffffff191690555050565b60008151116109ad57fe5b6004826040518082805190602001908083835b602083106109df5780518252601f1990920191602091820191016109c0565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054600160a060020a03161515610ab657336004836040518082805190602001908083835b60208310610a565780518252601f199092019160209182019101610a37565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555b5050565b602060405190810160405260008152905600a165627a7a72305820681bd32ce61ecbaa7f7dff71938658750bf04032deb1f13452a9089ca2db28c20029`

// DeployRegister_0_0_1 deploys a new Ethereum contract, binding an instance of Register_0_0_1 to it.
func DeployRegister_0_0_1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Register_0_0_1, error) {
	parsed, err := abi.JSON(strings.NewReader(Register_0_0_1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Register_0_0_1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Register_0_0_1{Register_0_0_1Caller: Register_0_0_1Caller{contract: contract}, Register_0_0_1Transactor: Register_0_0_1Transactor{contract: contract}}, nil
}

// Register_0_0_1 is an auto generated Go binding around an Ethereum contract.
type Register_0_0_1 struct {
	Register_0_0_1Caller     // Read-only binding to the contract
	Register_0_0_1Transactor // Write-only binding to the contract
}

// Register_0_0_1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Register_0_0_1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Register_0_0_1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Register_0_0_1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Register_0_0_1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Register_0_0_1Session struct {
	Contract     *Register_0_0_1         // Generic contract binding to set the session for
	CallOpts     bind.CallOptsWithNumber // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Register_0_0_1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Register_0_0_1CallerSession struct {
	Contract *Register_0_0_1Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOptsWithNumber // Call options to use throughout this session
}

// Register_0_0_1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Register_0_0_1TransactorSession struct {
	Contract     *Register_0_0_1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Register_0_0_1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Register_0_0_1Raw struct {
	Contract *Register_0_0_1 // Generic contract binding to access the raw methods on
}

// Register_0_0_1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Register_0_0_1CallerRaw struct {
	Contract *Register_0_0_1Caller // Generic read-only contract binding to access the raw methods on
}

// Register_0_0_1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Register_0_0_1TransactorRaw struct {
	Contract *Register_0_0_1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRegister_0_0_1 creates a new instance of Register_0_0_1, bound to a specific deployed contract.
func NewRegister_0_0_1(address common.Address, backend bind.ContractBackend) (*Register_0_0_1, error) {
	contract, err := bindRegister_0_0_1(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Register_0_0_1{Register_0_0_1Caller: Register_0_0_1Caller{contract: contract}, Register_0_0_1Transactor: Register_0_0_1Transactor{contract: contract}}, nil
}

// NewRegister_0_0_1Caller creates a new read-only instance of Register_0_0_1, bound to a specific deployed contract.
func NewRegister_0_0_1Caller(address common.Address, caller bind.ContractCaller) (*Register_0_0_1Caller, error) {
	contract, err := bindRegister_0_0_1(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &Register_0_0_1Caller{contract: contract}, nil
}

// NewRegister_0_0_1Transactor creates a new write-only instance of Register_0_0_1, bound to a specific deployed contract.
func NewRegister_0_0_1Transactor(address common.Address, transactor bind.ContractTransactor) (*Register_0_0_1Transactor, error) {
	contract, err := bindRegister_0_0_1(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &Register_0_0_1Transactor{contract: contract}, nil
}

// bindRegister_0_0_1 binds a generic wrapper to an already deployed contract.
func bindRegister_0_0_1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Register_0_0_1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register_0_0_1 *Register_0_0_1Raw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Register_0_0_1.Contract.Register_0_0_1Caller.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register_0_0_1 *Register_0_0_1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Register_0_0_1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register_0_0_1 *Register_0_0_1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Register_0_0_1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register_0_0_1 *Register_0_0_1CallerRaw) CallWithNumber(opts *bind.CallOptsWithNumber, result interface{}, method string, params ...interface{}) error {
	return _Register_0_0_1.Contract.contract.CallWithNumber(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register_0_0_1 *Register_0_0_1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register_0_0_1 *Register_0_0_1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.contract.Transact(opts, method, params...)
}

// GetId is a free data retrieval call binding the contract method 0xbee51f3b.
//
// Solidity: function getId(id string) constant returns(_id string, _cb address)
func (_Register_0_0_1 *Register_0_0_1Caller) GetId(opts *bind.CallOptsWithNumber, id string) (struct {
	Id string
	Cb common.Address
}, error) {
	ret := new(struct {
		Id string
		Cb common.Address
	})
	out := ret
	err := _Register_0_0_1.contract.CallWithNumber(opts, out, "getId", id)
	return *ret, err
}

// GetId is a free data retrieval call binding the contract method 0xbee51f3b.
//
// Solidity: function getId(id string) constant returns(_id string, _cb address)
func (_Register_0_0_1 *Register_0_0_1Session) GetId(id string) (struct {
	Id string
	Cb common.Address
}, error) {
	return _Register_0_0_1.Contract.GetId(&_Register_0_0_1.CallOpts, id)
}

// GetId is a free data retrieval call binding the contract method 0xbee51f3b.
//
// Solidity: function getId(id string) constant returns(_id string, _cb address)
func (_Register_0_0_1 *Register_0_0_1CallerSession) GetId(id string) (struct {
	Id string
	Cb common.Address
}, error) {
	return _Register_0_0_1.Contract.GetId(&_Register_0_0_1.CallOpts, id)
}

// GetVsn is a free data retrieval call binding the contract method 0xb6ce6143.
//
// Solidity: function getVsn() constant returns(string)
func (_Register_0_0_1 *Register_0_0_1Caller) GetVsn(opts *bind.CallOptsWithNumber) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Register_0_0_1.contract.CallWithNumber(opts, out, "getVsn")
	return *ret0, err
}

// GetVsn is a free data retrieval call binding the contract method 0xb6ce6143.
//
// Solidity: function getVsn() constant returns(string)
func (_Register_0_0_1 *Register_0_0_1Session) GetVsn() (string, error) {
	return _Register_0_0_1.Contract.GetVsn(&_Register_0_0_1.CallOpts)
}

// GetVsn is a free data retrieval call binding the contract method 0xb6ce6143.
//
// Solidity: function getVsn() constant returns(string)
func (_Register_0_0_1 *Register_0_0_1CallerSession) GetVsn() (string, error) {
	return _Register_0_0_1.Contract.GetVsn(&_Register_0_0_1.CallOpts)
}

// HasId is a free data retrieval call binding the contract method 0xcb72361c.
//
// Solidity: function hasId(id string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Caller) HasId(opts *bind.CallOptsWithNumber, id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Register_0_0_1.contract.CallWithNumber(opts, out, "hasId", id)
	return *ret0, err
}

// HasId is a free data retrieval call binding the contract method 0xcb72361c.
//
// Solidity: function hasId(id string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Session) HasId(id string) (bool, error) {
	return _Register_0_0_1.Contract.HasId(&_Register_0_0_1.CallOpts, id)
}

// HasId is a free data retrieval call binding the contract method 0xcb72361c.
//
// Solidity: function hasId(id string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1CallerSession) HasId(id string) (bool, error) {
	return _Register_0_0_1.Contract.HasId(&_Register_0_0_1.CallOpts, id)
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(addr address) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Caller) Verify(opts *bind.CallOptsWithNumber, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Register_0_0_1.contract.CallWithNumber(opts, out, "verify", addr)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(addr address) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Session) Verify(addr common.Address) (bool, error) {
	return _Register_0_0_1.Contract.Verify(&_Register_0_0_1.CallOpts, addr)
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(addr address) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1CallerSession) Verify(addr common.Address) (bool, error) {
	return _Register_0_0_1.Contract.Verify(&_Register_0_0_1.CallOpts, addr)
}

// VerifyBlackList is a free data retrieval call binding the contract method 0x3c094c4c.
//
// Solidity: function verifyBlackList(h string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Caller) VerifyBlackList(opts *bind.CallOptsWithNumber, h string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Register_0_0_1.contract.CallWithNumber(opts, out, "verifyBlackList", h)
	return *ret0, err
}

// VerifyBlackList is a free data retrieval call binding the contract method 0x3c094c4c.
//
// Solidity: function verifyBlackList(h string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1Session) VerifyBlackList(h string) (bool, error) {
	return _Register_0_0_1.Contract.VerifyBlackList(&_Register_0_0_1.CallOpts, h)
}

// VerifyBlackList is a free data retrieval call binding the contract method 0x3c094c4c.
//
// Solidity: function verifyBlackList(h string) constant returns(bool)
func (_Register_0_0_1 *Register_0_0_1CallerSession) VerifyBlackList(h string) (bool, error) {
	return _Register_0_0_1.Contract.VerifyBlackList(&_Register_0_0_1.CallOpts, h)
}

// Register is a paid mutator transaction binding the contract method 0x19e5bf3a.
//
// Solidity: function Register() returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "Register")
}

// Register is a paid mutator transaction binding the contract method 0x19e5bf3a.
//
// Solidity: function Register() returns()
func (_Register_0_0_1 *Register_0_0_1Session) Register() (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Register(&_Register_0_0_1.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x19e5bf3a.
//
// Solidity: function Register() returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) Register() (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Register(&_Register_0_0_1.TransactOpts)
}

// Append is a paid mutator transaction binding the contract method 0xac04f5a7.
//
// Solidity: function append(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) Append(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "append", addr)
}

// Append is a paid mutator transaction binding the contract method 0xac04f5a7.
//
// Solidity: function append(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1Session) Append(addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Append(&_Register_0_0_1.TransactOpts, addr)
}

// Append is a paid mutator transaction binding the contract method 0xac04f5a7.
//
// Solidity: function append(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) Append(addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Append(&_Register_0_0_1.TransactOpts, addr)
}

// AppendBlackList is a paid mutator transaction binding the contract method 0x20accbe6.
//
// Solidity: function appendBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) AppendBlackList(opts *bind.TransactOpts, h string) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "appendBlackList", h)
}

// AppendBlackList is a paid mutator transaction binding the contract method 0x20accbe6.
//
// Solidity: function appendBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1Session) AppendBlackList(h string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.AppendBlackList(&_Register_0_0_1.TransactOpts, h)
}

// AppendBlackList is a paid mutator transaction binding the contract method 0x20accbe6.
//
// Solidity: function appendBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) AppendBlackList(h string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.AppendBlackList(&_Register_0_0_1.TransactOpts, h)
}

// AppendId is a paid mutator transaction binding the contract method 0xf896cf10.
//
// Solidity: function appendId(id string, rawdata string) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) AppendId(opts *bind.TransactOpts, id string, rawdata string) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "appendId", id, rawdata)
}

// AppendId is a paid mutator transaction binding the contract method 0xf896cf10.
//
// Solidity: function appendId(id string, rawdata string) returns()
func (_Register_0_0_1 *Register_0_0_1Session) AppendId(id string, rawdata string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.AppendId(&_Register_0_0_1.TransactOpts, id, rawdata)
}

// AppendId is a paid mutator transaction binding the contract method 0xf896cf10.
//
// Solidity: function appendId(id string, rawdata string) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) AppendId(id string, rawdata string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.AppendId(&_Register_0_0_1.TransactOpts, id, rawdata)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) Remove(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "remove", addr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1Session) Remove(addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Remove(&_Register_0_0_1.TransactOpts, addr)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(addr address) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) Remove(addr common.Address) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.Remove(&_Register_0_0_1.TransactOpts, addr)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0x8c25bc12.
//
// Solidity: function removeBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) RemoveBlackList(opts *bind.TransactOpts, h string) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "removeBlackList", h)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0x8c25bc12.
//
// Solidity: function removeBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1Session) RemoveBlackList(h string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.RemoveBlackList(&_Register_0_0_1.TransactOpts, h)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0x8c25bc12.
//
// Solidity: function removeBlackList(h string) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) RemoveBlackList(h string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.RemoveBlackList(&_Register_0_0_1.TransactOpts, h)
}

// RemoveId is a paid mutator transaction binding the contract method 0xebd746c9.
//
// Solidity: function removeId(id string) returns()
func (_Register_0_0_1 *Register_0_0_1Transactor) RemoveId(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Register_0_0_1.contract.Transact(opts, "removeId", id)
}

// RemoveId is a paid mutator transaction binding the contract method 0xebd746c9.
//
// Solidity: function removeId(id string) returns()
func (_Register_0_0_1 *Register_0_0_1Session) RemoveId(id string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.RemoveId(&_Register_0_0_1.TransactOpts, id)
}

// RemoveId is a paid mutator transaction binding the contract method 0xebd746c9.
//
// Solidity: function removeId(id string) returns()
func (_Register_0_0_1 *Register_0_0_1TransactorSession) RemoveId(id string) (*types.Transaction, error) {
	return _Register_0_0_1.Contract.RemoveId(&_Register_0_0_1.TransactOpts, id)
}
