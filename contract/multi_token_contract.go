// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MultiTokenMetaData contains all meta data concerning the MultiToken contract.
var MultiTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"createToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610ad28061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c8063040318521461004e57806304c2320b1461007e5780635b060530146100b05780638892ebce146100cc575b5f5ffd5b61006860048036038101906100639190610559565b6100fc565b60405161007591906105cb565b60405180910390f35b610098600480360381019061009391906105e4565b61015f565b6040516100a79392919061068b565b60405180910390f35b6100ca60048036038101906100c591906106f8565b6102a8565b005b6100e660048036038101906100e19190610559565b61037a565b6040516100f391906105cb565b60405180910390f35b5f60018360405161010d91906107ba565b90815260200160405180910390205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f018054610195906107fd565b80601f01602080910402602001604051908101604052809291908181526020018280546101c1906107fd565b801561020c5780601f106101e35761010080835404028352916020019161020c565b820191905f5260205f20905b8154815290600101906020018083116101ef57829003601f168201915b505050505090806001018054610221906107fd565b80601f016020809104026020016040519081016040528092919081815260200182805461024d906107fd565b80156102985780601f1061026f57610100808354040283529160200191610298565b820191905f5260205f20905b81548152906001019060200180831161027b57829003601f168201915b5050505050908060020154905083565b6040518060600160405280848152602001838152602001828152505f836040516102d291906107ba565b90815260200160405180910390205f820151815f0190816102f391906109cd565b50602082015181600101908161030991906109cd565b50604082015181600201559050508060018360405161032891906107ba565b90815260200160405180910390205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505050565b600182805160208101820180518482526020830160208501208183528095505050505050602052805f5260405f205f91509150505481565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610411826103cb565b810181811067ffffffffffffffff821117156104305761042f6103db565b5b80604052505050565b5f6104426103b2565b905061044e8282610408565b919050565b5f67ffffffffffffffff82111561046d5761046c6103db565b5b610476826103cb565b9050602081019050919050565b828183375f83830152505050565b5f6104a361049e84610453565b610439565b9050828152602081018484840111156104bf576104be6103c7565b5b6104ca848285610483565b509392505050565b5f82601f8301126104e6576104e56103c3565b5b81356104f6848260208601610491565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610528826104ff565b9050919050565b6105388161051e565b8114610542575f5ffd5b50565b5f813590506105538161052f565b92915050565b5f5f6040838503121561056f5761056e6103bb565b5b5f83013567ffffffffffffffff81111561058c5761058b6103bf565b5b610598858286016104d2565b92505060206105a985828601610545565b9150509250929050565b5f819050919050565b6105c5816105b3565b82525050565b5f6020820190506105de5f8301846105bc565b92915050565b5f602082840312156105f9576105f86103bb565b5b5f82013567ffffffffffffffff811115610616576106156103bf565b5b610622848285016104d2565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61065d8261062b565b6106678185610635565b9350610677818560208601610645565b610680816103cb565b840191505092915050565b5f6060820190508181035f8301526106a38186610653565b905081810360208301526106b78185610653565b90506106c660408301846105bc565b949350505050565b6106d7816105b3565b81146106e1575f5ffd5b50565b5f813590506106f2816106ce565b92915050565b5f5f5f6060848603121561070f5761070e6103bb565b5b5f84013567ffffffffffffffff81111561072c5761072b6103bf565b5b610738868287016104d2565b935050602084013567ffffffffffffffff811115610759576107586103bf565b5b610765868287016104d2565b9250506040610776868287016106e4565b9150509250925092565b5f81905092915050565b5f6107948261062b565b61079e8185610780565b93506107ae818560208601610645565b80840191505092915050565b5f6107c5828461078a565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061081457607f821691505b602082108103610827576108266107d0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261084e565b610893868361084e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6108ce6108c96108c4846105b3565b6108ab565b6105b3565b9050919050565b5f819050919050565b6108e7836108b4565b6108fb6108f3826108d5565b84845461085a565b825550505050565b5f5f905090565b610912610903565b61091d8184846108de565b505050565b5b81811015610940576109355f8261090a565b600181019050610923565b5050565b601f821115610985576109568161082d565b61095f8461083f565b8101602085101561096e578190505b61098261097a8561083f565b830182610922565b50505b505050565b5f82821c905092915050565b5f6109a55f198460080261098a565b1980831691505092915050565b5f6109bd8383610996565b9150826002028217905092915050565b6109d68261062b565b67ffffffffffffffff8111156109ef576109ee6103db565b5b6109f982546107fd565b610a04828285610944565b5f60209050601f831160018114610a35575f8415610a23578287015190505b610a2d85826109b2565b865550610a94565b601f198416610a438661082d565b5f5b82811015610a6a57848901518255600182019150602085019450602081019050610a45565b86831015610a875784890151610a83601f891682610996565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220757d78eea8c5425a37d404bd297d5a53e942a25ed449b1e50df69765acdf502564736f6c634300081c0033",
}

// MultiTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiTokenMetaData.ABI instead.
var MultiTokenABI = MultiTokenMetaData.ABI

// MultiTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiTokenMetaData.Bin instead.
var MultiTokenBin = MultiTokenMetaData.Bin

// DeployMultiToken deploys a new Ethereum contract, binding an instance of MultiToken to it.
func DeployMultiToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiToken, error) {
	parsed, err := MultiTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiToken{MultiTokenCaller: MultiTokenCaller{contract: contract}, MultiTokenTransactor: MultiTokenTransactor{contract: contract}, MultiTokenFilterer: MultiTokenFilterer{contract: contract}}, nil
}

// MultiToken is an auto generated Go binding around an Ethereum contract.
type MultiToken struct {
	MultiTokenCaller     // Read-only binding to the contract
	MultiTokenTransactor // Write-only binding to the contract
	MultiTokenFilterer   // Log filterer for contract events
}

// MultiTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiTokenSession struct {
	Contract     *MultiToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiTokenCallerSession struct {
	Contract *MultiTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MultiTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiTokenTransactorSession struct {
	Contract     *MultiTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MultiTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiTokenRaw struct {
	Contract *MultiToken // Generic contract binding to access the raw methods on
}

// MultiTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiTokenCallerRaw struct {
	Contract *MultiTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MultiTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiTokenTransactorRaw struct {
	Contract *MultiTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiToken creates a new instance of MultiToken, bound to a specific deployed contract.
func NewMultiToken(address common.Address, backend bind.ContractBackend) (*MultiToken, error) {
	contract, err := bindMultiToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiToken{MultiTokenCaller: MultiTokenCaller{contract: contract}, MultiTokenTransactor: MultiTokenTransactor{contract: contract}, MultiTokenFilterer: MultiTokenFilterer{contract: contract}}, nil
}

// NewMultiTokenCaller creates a new read-only instance of MultiToken, bound to a specific deployed contract.
func NewMultiTokenCaller(address common.Address, caller bind.ContractCaller) (*MultiTokenCaller, error) {
	contract, err := bindMultiToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenCaller{contract: contract}, nil
}

// NewMultiTokenTransactor creates a new write-only instance of MultiToken, bound to a specific deployed contract.
func NewMultiTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTokenTransactor, error) {
	contract, err := bindMultiToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTokenTransactor{contract: contract}, nil
}

// NewMultiTokenFilterer creates a new log filterer instance of MultiToken, bound to a specific deployed contract.
func NewMultiTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiTokenFilterer, error) {
	contract, err := bindMultiToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiTokenFilterer{contract: contract}, nil
}

// bindMultiToken binds a generic wrapper to an already deployed contract.
func bindMultiToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiToken *MultiTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiToken.Contract.MultiTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiToken *MultiTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiToken.Contract.MultiTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiToken *MultiTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiToken.Contract.MultiTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiToken *MultiTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiToken *MultiTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiToken *MultiTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x04031852.
//
// Solidity: function balanceOf(string symbol, address account) view returns(uint256)
func (_MultiToken *MultiTokenCaller) BalanceOf(opts *bind.CallOpts, symbol string, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiToken.contract.Call(opts, &out, "balanceOf", symbol, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x04031852.
//
// Solidity: function balanceOf(string symbol, address account) view returns(uint256)
func (_MultiToken *MultiTokenSession) BalanceOf(symbol string, account common.Address) (*big.Int, error) {
	return _MultiToken.Contract.BalanceOf(&_MultiToken.CallOpts, symbol, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x04031852.
//
// Solidity: function balanceOf(string symbol, address account) view returns(uint256)
func (_MultiToken *MultiTokenCallerSession) BalanceOf(symbol string, account common.Address) (*big.Int, error) {
	return _MultiToken.Contract.BalanceOf(&_MultiToken.CallOpts, symbol, account)
}

// Balances is a free data retrieval call binding the contract method 0x8892ebce.
//
// Solidity: function balances(string , address ) view returns(uint256)
func (_MultiToken *MultiTokenCaller) Balances(opts *bind.CallOpts, arg0 string, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiToken.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x8892ebce.
//
// Solidity: function balances(string , address ) view returns(uint256)
func (_MultiToken *MultiTokenSession) Balances(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _MultiToken.Contract.Balances(&_MultiToken.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x8892ebce.
//
// Solidity: function balances(string , address ) view returns(uint256)
func (_MultiToken *MultiTokenCallerSession) Balances(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _MultiToken.Contract.Balances(&_MultiToken.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(string name, string symbol, uint256 totalSupply)
func (_MultiToken *MultiTokenCaller) Tokens(opts *bind.CallOpts, arg0 string) (struct {
	Name        string
	Symbol      string
	TotalSupply *big.Int
}, error) {
	var out []interface{}
	err := _MultiToken.contract.Call(opts, &out, "tokens", arg0)

	outstruct := new(struct {
		Name        string
		Symbol      string
		TotalSupply *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TotalSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(string name, string symbol, uint256 totalSupply)
func (_MultiToken *MultiTokenSession) Tokens(arg0 string) (struct {
	Name        string
	Symbol      string
	TotalSupply *big.Int
}, error) {
	return _MultiToken.Contract.Tokens(&_MultiToken.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(string name, string symbol, uint256 totalSupply)
func (_MultiToken *MultiTokenCallerSession) Tokens(arg0 string) (struct {
	Name        string
	Symbol      string
	TotalSupply *big.Int
}, error) {
	return _MultiToken.Contract.Tokens(&_MultiToken.CallOpts, arg0)
}

// CreateToken is a paid mutator transaction binding the contract method 0x5b060530.
//
// Solidity: function createToken(string name, string symbol, uint256 totalSupply) returns()
func (_MultiToken *MultiTokenTransactor) CreateToken(opts *bind.TransactOpts, name string, symbol string, totalSupply *big.Int) (*types.Transaction, error) {
	return _MultiToken.contract.Transact(opts, "createToken", name, symbol, totalSupply)
}

// CreateToken is a paid mutator transaction binding the contract method 0x5b060530.
//
// Solidity: function createToken(string name, string symbol, uint256 totalSupply) returns()
func (_MultiToken *MultiTokenSession) CreateToken(name string, symbol string, totalSupply *big.Int) (*types.Transaction, error) {
	return _MultiToken.Contract.CreateToken(&_MultiToken.TransactOpts, name, symbol, totalSupply)
}

// CreateToken is a paid mutator transaction binding the contract method 0x5b060530.
//
// Solidity: function createToken(string name, string symbol, uint256 totalSupply) returns()
func (_MultiToken *MultiTokenTransactorSession) CreateToken(name string, symbol string, totalSupply *big.Int) (*types.Transaction, error) {
	return _MultiToken.Contract.CreateToken(&_MultiToken.TransactOpts, name, symbol, totalSupply)
}
