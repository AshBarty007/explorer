// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ProofRecord is an auto generated low-level Go binding around an user-defined struct.
type ProofRecord struct {
	Timestamp *big.Int
	TypeName  string
	IpfsHash  string
}

// EthClientMetaData contains all meta data concerning the EthClient contract.
var EthClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRecordsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"internalType\":\"structProof.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506111668061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806306661abd1461004e57806357ac40c31461006c5780639507d39a1461009c578063a2387ea3146100cc575b5f80fd5b6100566100e8565b60405161006391906106f5565b60405180910390f35b61008660048036038101906100819190610770565b6100f3565b6040516100939190610929565b60405180910390f35b6100b660048036038101906100b19190610973565b6102b5565b6040516100c391906109e6565b60405180910390f35b6100e660048036038101906100e19190610a67565b6103a6565b005b5f8080549050905090565b606060015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156102aa578382905f5260205f2090600302016040518060600160405290815f820154815260200160018201805461018b90610b12565b80601f01602080910402602001604051908101604052809291908181526020018280546101b790610b12565b80156102025780601f106101d957610100808354040283529160200191610202565b820191905f5260205f20905b8154815290600101906020018083116101e557829003601f168201915b5050505050815260200160028201805461021b90610b12565b80601f016020809104026020016040519081016040528092919081815260200182805461024790610b12565b80156102925780601f1061026957610100808354040283529160200191610292565b820191905f5260205f20905b81548152906001019060200180831161027557829003601f168201915b50505050508152505081526020019060010190610151565b505050509050919050565b60605f8054905082106102fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f490610b8c565b60405180910390fd5b5f82815481106103105761030f610baa565b5b905f5260205f2001805461032390610b12565b80601f016020809104026020016040519081016040528092919081815260200182805461034f90610b12565b801561039a5780601f106103715761010080835404028352916020019161039a565b820191905f5260205f20905b81548152906001019060200180831161037d57829003601f168201915b50505050509050919050565b5f82829050116103eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e290610c21565b60405180910390fd5b6103f58484610618565b610434576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042b90610c89565b60405180910390fd5b5f828290918060018154018082558091505060019003905f5260205f20015f90919290919290919290919250918261046d929190610e7b565b505f604051806060016040528042815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815250905060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015560208201518160010190816105939190610f48565b5060408201518160020190816105a99190610f48565b5050503373ffffffffffffffffffffffffffffffffffffffff1660015f805490506105d49190611044565b7ff12893fae46463a091829bbfbed5e8967fb04f4bd11d26388620b56b393c905c8585898960405161060994939291906110b1565b60405180910390a35050505050565b5f7fe91e69686af56c584e188c4128cbaf29402b5eceaf4d71f2e4497f57e80dea82838360405161064a929190611118565b6040518091039020148061069457507f6a95564804bd5a2a5beeb586b12acc1871828a673011454654642a7d945d5b26838360405161068a929190611118565b6040518091039020145b806106d557507f7477535acdef313b25d16b4871e7023fac62af68d6312bbdbdb96203a4710dc383836040516106cb929190611118565b6040518091039020145b905092915050565b5f819050919050565b6106ef816106dd565b82525050565b5f6020820190506107085f8301846106e6565b92915050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61073f82610716565b9050919050565b61074f81610735565b8114610759575f80fd5b50565b5f8135905061076a81610746565b92915050565b5f602082840312156107855761078461070e565b5b5f6107928482850161075c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6107cd816106dd565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610815826107d3565b61081f81856107dd565b935061082f8185602086016107ed565b610838816107fb565b840191505092915050565b5f606083015f8301516108585f8601826107c4565b5060208301518482036020860152610870828261080b565b9150506040830151848203604086015261088a828261080b565b9150508091505092915050565b5f6108a28383610843565b905092915050565b5f602082019050919050565b5f6108c08261079b565b6108ca81856107a5565b9350836020820285016108dc856107b5565b805f5b8581101561091757848403895281516108f88582610897565b9450610903836108aa565b925060208a019950506001810190506108df565b50829750879550505050505092915050565b5f6020820190508181035f83015261094181846108b6565b905092915050565b610952816106dd565b811461095c575f80fd5b50565b5f8135905061096d81610949565b92915050565b5f602082840312156109885761098761070e565b5b5f6109958482850161095f565b91505092915050565b5f82825260208201905092915050565b5f6109b8826107d3565b6109c2818561099e565b93506109d28185602086016107ed565b6109db816107fb565b840191505092915050565b5f6020820190508181035f8301526109fe81846109ae565b905092915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610a2757610a26610a06565b5b8235905067ffffffffffffffff811115610a4457610a43610a0a565b5b602083019150836001820283011115610a6057610a5f610a0e565b5b9250929050565b5f805f8060408587031215610a7f57610a7e61070e565b5b5f85013567ffffffffffffffff811115610a9c57610a9b610712565b5b610aa887828801610a12565b9450945050602085013567ffffffffffffffff811115610acb57610aca610712565b5b610ad787828801610a12565b925092505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610b2957607f821691505b602082108103610b3c57610b3b610ae5565b5b50919050565b7f496e76616c696420696e646578000000000000000000000000000000000000005f82015250565b5f610b76600d8361099e565b9150610b8182610b42565b602082019050919050565b5f6020820190508181035f830152610ba381610b6a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f456d7074792068617368000000000000000000000000000000000000000000005f82015250565b5f610c0b600a8361099e565b9150610c1682610bd7565b602082019050919050565b5f6020820190508181035f830152610c3881610bff565b9050919050565b7f496e76616c69642074797065206e616d650000000000000000000000000000005f82015250565b5f610c7360118361099e565b9150610c7e82610c3f565b602082019050919050565b5f6020820190508181035f830152610ca081610c67565b9050919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610d3a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610cff565b610d448683610cff565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610d7f610d7a610d75846106dd565b610d5c565b6106dd565b9050919050565b5f819050919050565b610d9883610d65565b610dac610da482610d86565b848454610d0b565b825550505050565b5f90565b610dc0610db4565b610dcb818484610d8f565b505050565b5b81811015610dee57610de35f82610db8565b600181019050610dd1565b5050565b601f821115610e3357610e0481610cde565b610e0d84610cf0565b81016020851015610e1c578190505b610e30610e2885610cf0565b830182610dd0565b50505b505050565b5f82821c905092915050565b5f610e535f1984600802610e38565b1980831691505092915050565b5f610e6b8383610e44565b9150826002028217905092915050565b610e858383610ca7565b67ffffffffffffffff811115610e9e57610e9d610cb1565b5b610ea88254610b12565b610eb3828285610df2565b5f601f831160018114610ee0575f8415610ece578287013590505b610ed88582610e60565b865550610f3f565b601f198416610eee86610cde565b5f5b82811015610f1557848901358255600182019150602085019450602081019050610ef0565b86831015610f325784890135610f2e601f891682610e44565b8355505b6001600288020188555050505b50505050505050565b610f51826107d3565b67ffffffffffffffff811115610f6a57610f69610cb1565b5b610f748254610b12565b610f7f828285610df2565b5f60209050601f831160018114610fb0575f8415610f9e578287015190505b610fa88582610e60565b86555061100f565b601f198416610fbe86610cde565b5f5b82811015610fe557848901518255600182019150602085019450602081019050610fc0565b868310156110025784890151610ffe601f891682610e44565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61104e826106dd565b9150611059836106dd565b925082820390508181111561107157611070611017565b5b92915050565b828183375f83830152505050565b5f611090838561099e565b935061109d838584611077565b6110a6836107fb565b840190509392505050565b5f6040820190508181035f8301526110ca818688611085565b905081810360208301526110df818486611085565b905095945050505050565b5f81905092915050565b5f6110ff83856110ea565b935061110c838584611077565b82840190509392505050565b5f6111248284866110f4565b9150819050939250505056fea26469706673582212201d2e0134d6331cf1e15668c7ff3b86a09ec76d3be7f5b5cfcceecdfd4d8a849e64736f6c634300081a0033",
}

// EthClientABI is the input ABI used to generate the binding from.
// Deprecated: Use EthClientMetaData.ABI instead.
var EthClientABI = EthClientMetaData.ABI

// EthClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthClientMetaData.Bin instead.
var EthClientBin = EthClientMetaData.Bin

// DeployEthClient deploys a new Ethereum contract, binding an instance of EthClient to it.
func DeployEthClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthClient, error) {
	parsed, err := EthClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthClient{EthClientCaller: EthClientCaller{contract: contract}, EthClientTransactor: EthClientTransactor{contract: contract}, EthClientFilterer: EthClientFilterer{contract: contract}}, nil
}

// EthClient is an auto generated Go binding around an Ethereum contract.
type EthClient struct {
	EthClientCaller     // Read-only binding to the contract
	EthClientTransactor // Write-only binding to the contract
	EthClientFilterer   // Log filterer for contract events
}

// EthClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthClientSession struct {
	Contract     *EthClient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthClientCallerSession struct {
	Contract *EthClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthClientTransactorSession struct {
	Contract     *EthClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthClientRaw struct {
	Contract *EthClient // Generic contract binding to access the raw methods on
}

// EthClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthClientCallerRaw struct {
	Contract *EthClientCaller // Generic read-only contract binding to access the raw methods on
}

// EthClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthClientTransactorRaw struct {
	Contract *EthClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthClient creates a new instance of EthClient, bound to a specific deployed contract.
func NewEthClient(address common.Address, backend bind.ContractBackend) (*EthClient, error) {
	contract, err := bindEthClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthClient{EthClientCaller: EthClientCaller{contract: contract}, EthClientTransactor: EthClientTransactor{contract: contract}, EthClientFilterer: EthClientFilterer{contract: contract}}, nil
}

// NewEthClientCaller creates a new read-only instance of EthClient, bound to a specific deployed contract.
func NewEthClientCaller(address common.Address, caller bind.ContractCaller) (*EthClientCaller, error) {
	contract, err := bindEthClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthClientCaller{contract: contract}, nil
}

// NewEthClientTransactor creates a new write-only instance of EthClient, bound to a specific deployed contract.
func NewEthClientTransactor(address common.Address, transactor bind.ContractTransactor) (*EthClientTransactor, error) {
	contract, err := bindEthClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthClientTransactor{contract: contract}, nil
}

// NewEthClientFilterer creates a new log filterer instance of EthClient, bound to a specific deployed contract.
func NewEthClientFilterer(address common.Address, filterer bind.ContractFilterer) (*EthClientFilterer, error) {
	contract, err := bindEthClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthClientFilterer{contract: contract}, nil
}

// bindEthClient binds a generic wrapper to an already deployed contract.
func bindEthClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClient *EthClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClient.Contract.EthClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClient *EthClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClient.Contract.EthClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClient *EthClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClient.Contract.EthClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClient *EthClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClient *EthClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClient *EthClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClient.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_EthClient *EthClientCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_EthClient *EthClientSession) Count() (*big.Int, error) {
	return _EthClient.Contract.Count(&_EthClient.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_EthClient *EthClientCallerSession) Count() (*big.Int, error) {
	return _EthClient.Contract.Count(&_EthClient.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 index) view returns(string)
func (_EthClient *EthClientCaller) Get(opts *bind.CallOpts, index *big.Int) (string, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "get", index)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 index) view returns(string)
func (_EthClient *EthClientSession) Get(index *big.Int) (string, error) {
	return _EthClient.Contract.Get(&_EthClient.CallOpts, index)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 index) view returns(string)
func (_EthClient *EthClientCallerSession) Get(index *big.Int) (string, error) {
	return _EthClient.Contract.Get(&_EthClient.CallOpts, index)
}

// GetRecordsByUser is a free data retrieval call binding the contract method 0x57ac40c3.
//
// Solidity: function getRecordsByUser(address user) view returns((uint256,string,string)[])
func (_EthClient *EthClientCaller) GetRecordsByUser(opts *bind.CallOpts, user common.Address) ([]ProofRecord, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "getRecordsByUser", user)

	if err != nil {
		return *new([]ProofRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]ProofRecord)).(*[]ProofRecord)

	return out0, err

}

// GetRecordsByUser is a free data retrieval call binding the contract method 0x57ac40c3.
//
// Solidity: function getRecordsByUser(address user) view returns((uint256,string,string)[])
func (_EthClient *EthClientSession) GetRecordsByUser(user common.Address) ([]ProofRecord, error) {
	return _EthClient.Contract.GetRecordsByUser(&_EthClient.CallOpts, user)
}

// GetRecordsByUser is a free data retrieval call binding the contract method 0x57ac40c3.
//
// Solidity: function getRecordsByUser(address user) view returns((uint256,string,string)[])
func (_EthClient *EthClientCallerSession) GetRecordsByUser(user common.Address) ([]ProofRecord, error) {
	return _EthClient.Contract.GetRecordsByUser(&_EthClient.CallOpts, user)
}

// Submit is a paid mutator transaction binding the contract method 0xa2387ea3.
//
// Solidity: function submit(string typeName, string ipfsHash) returns()
func (_EthClient *EthClientTransactor) Submit(opts *bind.TransactOpts, typeName string, ipfsHash string) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "submit", typeName, ipfsHash)
}

// Submit is a paid mutator transaction binding the contract method 0xa2387ea3.
//
// Solidity: function submit(string typeName, string ipfsHash) returns()
func (_EthClient *EthClientSession) Submit(typeName string, ipfsHash string) (*types.Transaction, error) {
	return _EthClient.Contract.Submit(&_EthClient.TransactOpts, typeName, ipfsHash)
}

// Submit is a paid mutator transaction binding the contract method 0xa2387ea3.
//
// Solidity: function submit(string typeName, string ipfsHash) returns()
func (_EthClient *EthClientTransactorSession) Submit(typeName string, ipfsHash string) (*types.Transaction, error) {
	return _EthClient.Contract.Submit(&_EthClient.TransactOpts, typeName, ipfsHash)
}

// EthClientSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the EthClient contract.
type EthClientSubmittedIterator struct {
	Event *EthClientSubmitted // Event containing the contract specifics and raw log

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
func (it *EthClientSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClientSubmitted)
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
		it.Event = new(EthClientSubmitted)
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
func (it *EthClientSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClientSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClientSubmitted represents a Submitted event raised by the EthClient contract.
type EthClientSubmitted struct {
	IpfsHash string
	Index    *big.Int
	User     common.Address
	TypeName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0xf12893fae46463a091829bbfbed5e8967fb04f4bd11d26388620b56b393c905c.
//
// Solidity: event Submitted(string ipfsHash, uint256 indexed index, address indexed user, string typeName)
func (_EthClient *EthClientFilterer) FilterSubmitted(opts *bind.FilterOpts, index []*big.Int, user []common.Address) (*EthClientSubmittedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EthClient.contract.FilterLogs(opts, "Submitted", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &EthClientSubmittedIterator{contract: _EthClient.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0xf12893fae46463a091829bbfbed5e8967fb04f4bd11d26388620b56b393c905c.
//
// Solidity: event Submitted(string ipfsHash, uint256 indexed index, address indexed user, string typeName)
func (_EthClient *EthClientFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *EthClientSubmitted, index []*big.Int, user []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _EthClient.contract.WatchLogs(opts, "Submitted", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClientSubmitted)
				if err := _EthClient.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0xf12893fae46463a091829bbfbed5e8967fb04f4bd11d26388620b56b393c905c.
//
// Solidity: event Submitted(string ipfsHash, uint256 indexed index, address indexed user, string typeName)
func (_EthClient *EthClientFilterer) ParseSubmitted(log types.Log) (*EthClientSubmitted, error) {
	event := new(EthClientSubmitted)
	if err := _EthClient.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
