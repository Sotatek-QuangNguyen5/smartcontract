// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

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

// MyContractSubcriber is an auto generated low-level Go binding around an user-defined struct.
type MyContractSubcriber struct {
	Addr          common.Address
	IsCreator     bool
	IsSubciber    bool
	IsUnSubcirber bool
	Refund        bool
	Amount        *big.Int
}

// MyContractMetaData contains all meta data concerning the MyContract contract.
var MyContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"List\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCreator\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSubciber\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnSubcirber\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structMyContract.Subcriber[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Refund\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subcribe\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unSubcribe\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff02191690831515021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160156101000a81548160ff0219169083151502179055506000600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160166101000a81548160ff02191690831515021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506002600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900460ff168160000160146101000a81548160ff0219169083151502179055506000820160159054906101000a900460ff168160000160156101000a81548160ff0219169083151502179055506000820160169054906101000a900460ff168160000160166101000a81548160ff0219169083151502179055506000820160179054906101000a900460ff168160000160176101000a81548160ff02191690831515021790555060018201548160010155505061191f806104af6000396000f3fe60806040526004361061004a5760003560e01c806302d05d3f1461004f5780635d2686291461007a5780639645c393146100a5578063c18858fd146100d0578063c258ff74146100ee575b600080fd5b34801561005b57600080fd5b50610064610119565b60405161007191906110d8565b60405180910390f35b34801561008657600080fd5b5061008f61013d565b60405161009c9190611183565b60405180910390f35b3480156100b157600080fd5b506100ba61045a565b6040516100c79190611183565b60405180910390f35b6100d86106ee565b6040516100e59190611183565b60405180910390f35b3480156100fa57600080fd5b50610103610907565b6040516101109190611312565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160159054906101000a900460ff1615151480156101f5575060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160169054906101000a900460ff161515145b610234576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022b90611380565b60405180910390fd5b60001515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160179054906101000a900460ff161515146102ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c1906113ec565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff166002600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154610332919061146a565b60405161033e906114cc565b60006040518083038185875af1925050503d806000811461037b576040519150601f19603f3d011682016040523d82523d6000602084013e610380565b606091505b50509050806103c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bb9061152d565b60405180910390fd5b60018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160176101000a81548160ff0219169083151502179055506040518060400160405280601781526020017f526566756e64207375636365737366756c6c792021212100000000000000000081525091505090565b606060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff1615151480610511575060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160159054906101000a900460ff161515145b8015610570575060001515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160169054906101000a900460ff161515145b6105af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a690611599565b60405180910390fd5b60018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160166101000a81548160ff02191690831515021790555060005b6002805490508110156106b2576002818154811061062d5761062c6115b9565b5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361069f5761069a81610a32565b6106b2565b80806106aa906115e8565b91505061060c565b506040518060400160405280602081526020017f596f7520756e737562637269626572207375636365737366756c6c7920212121815250905090565b606060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160169054906101000a900460ff16151503610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d906116a2565b60405180910390fd5b60011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160159054906101000a900460ff161515148061083b575060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515145b1561087b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108729061170e565b60405180910390fd5b655af3107a40003410156108c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bb906117a0565b60405180910390fd5b6108cc610cd1565b6040518060400160405280601e81526020017f596f7520737562637269626572207375636365737366756c6c79202121210000815250905090565b60606002805480602002602001604051908101604052809291908181526020016000905b82821015610a2957838290600052602060002090600202016040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff161515151581526020016000820160159054906101000a900460ff161515151581526020016000820160169054906101000a900460ff161515151581526020016000820160179054906101000a900460ff161515151581526020016001820154815250508152602001906001019061092b565b50505050905090565b6002805490508110610a79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7090611832565b60405180910390fd5b60008190505b6001600280549050610a919190611852565b811015610c23576002600182610aa79190611886565b81548110610ab857610ab76115b9565b5b906000526020600020906002020160028281548110610ada57610ad96115b9565b5b90600052602060002090600202016000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900460ff168160000160146101000a81548160ff0219169083151502179055506000820160159054906101000a900460ff168160000160156101000a81548160ff0219169083151502179055506000820160169054906101000a900460ff168160000160166101000a81548160ff0219169083151502179055506000820160179054906101000a900460ff168160000160176101000a81548160ff021916908315150217905550600182015481600101559050508080610c1b906115e8565b915050610a7f565b506002805480610c3657610c356118ba565b5b6001900381819060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff02191690556000820160156101000a81549060ff02191690556000820160166101000a81549060ff02191690556000820160176101000a81549060ff021916905560018201600090555050905550565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff02191690831515021790555060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160156101000a81548160ff0219169083151502179055506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160166101000a81548160ff0219169083151502179055506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160176101000a81548160ff02191690831515021790555033600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506002600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900460ff168160000160146101000a81548160ff0219169083151502179055506000820160159054906101000a900460ff168160000160156101000a81548160ff0219169083151502179055506000820160169054906101000a900460ff168160000160166101000a81548160ff0219169083151502179055506000820160179054906101000a900460ff168160000160176101000a81548160ff021916908315150217905550600182015481600101555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110c282611097565b9050919050565b6110d2816110b7565b82525050565b60006020820190506110ed60008301846110c9565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561112d578082015181840152602081019050611112565b60008484015250505050565b6000601f19601f8301169050919050565b6000611155826110f3565b61115f81856110fe565b935061116f81856020860161110f565b61117881611139565b840191505092915050565b6000602082019050818103600083015261119d818461114a565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6111da816110b7565b82525050565b60008115159050919050565b6111f5816111e0565b82525050565b6000819050919050565b61120e816111fb565b82525050565b60c08201600082015161122a60008501826111d1565b50602082015161123d60208501826111ec565b50604082015161125060408501826111ec565b50606082015161126360608501826111ec565b50608082015161127660808501826111ec565b5060a082015161128960a0850182611205565b50505050565b600061129b8383611214565b60c08301905092915050565b6000602082019050919050565b60006112bf826111a5565b6112c981856111b0565b93506112d4836111c1565b8060005b838110156113055781516112ec888261128f565b97506112f7836112a7565b9250506001810190506112d8565b5085935050505092915050565b6000602082019050818103600083015261132c81846112b4565b905092915050565b7f596f752068617665206e6f7420756e7375627363726962656420212121000000600082015250565b600061136a601d836110fe565b915061137582611334565b602082019050919050565b600060208201905081810360008301526113998161135d565b9050919050565b7f596f7520676f7420796f7572206d6f6e6579206261636b202121210000000000600082015250565b60006113d6601b836110fe565b91506113e1826113a0565b602082019050919050565b60006020820190508181036000830152611405816113c9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611475826111fb565b9150611480836111fb565b9250826114905761148f61140c565b5b828204905092915050565b600081905092915050565b50565b60006114b660008361149b565b91506114c1826114a6565b600082019050919050565b60006114d7826114a9565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b60006115176014836110fe565b9150611522826114e1565b602082019050919050565b600060208201905081810360008301526115468161150a565b9050919050565b7f4e6f742061757468656e7469636174696f6e2021212100000000000000000000600082015250565b60006115836016836110fe565b915061158e8261154d565b602082019050919050565b600060208201905081810360008301526115b281611576565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006115f3826111fb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116255761162461143b565b5b600182019050919050565b7f596f75206861766520756e737562736372696265642c20796f752063616e6e6f60008201527f742073756273637269626520616761696e202121210000000000000000000000602082015250565b600061168c6035836110fe565b915061169782611630565b604082019050919050565b600060208201905081810360008301526116bb8161167f565b9050919050565b7f596f752061726520616c72656164792072656769737465726564202121210000600082015250565b60006116f8601e836110fe565b9150611703826116c2565b602082019050919050565b60006020820190508181036000830152611727816116eb565b9050919050565b7f596f75206e65656420746f20706179206174206c6561737420302e303030312060008201527f657468657220746f207369676e2074686520636f6e7472616374202121210000602082015250565b600061178a603e836110fe565b91506117958261172e565b604082019050919050565b600060208201905081810360008301526117b98161177d565b9050919050565b7f4572726f7220737562637269626572206973206e6f742073757276697665202160008201527f2121000000000000000000000000000000000000000000000000000000000000602082015250565b600061181c6022836110fe565b9150611827826117c0565b604082019050919050565b6000602082019050818103600083015261184b8161180f565b9050919050565b600061185d826111fb565b9150611868836111fb565b92508282039050818111156118805761187f61143b565b5b92915050565b6000611891826111fb565b915061189c836111fb565b92508282019050808211156118b4576118b361143b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212200f6bdc73c64a4b495ef8363cd58326e1c6c64fdf23099e9790d25b82954c464464736f6c63430008110033",
}

// MyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContractMetaData.ABI instead.
var MyContractABI = MyContractMetaData.ABI

// MyContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyContractMetaData.Bin instead.
var MyContractBin = MyContractMetaData.Bin

// DeployMyContract deploys a new Ethereum contract, binding an instance of MyContract to it.
func DeployMyContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyContract, error) {
	parsed, err := MyContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract     *MyContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() view returns((address,bool,bool,bool,bool,uint256)[])
func (_MyContract *MyContractCaller) List(opts *bind.CallOpts) ([]MyContractSubcriber, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "List")

	if err != nil {
		return *new([]MyContractSubcriber), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyContractSubcriber)).(*[]MyContractSubcriber)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() view returns((address,bool,bool,bool,bool,uint256)[])
func (_MyContract *MyContractSession) List() ([]MyContractSubcriber, error) {
	return _MyContract.Contract.List(&_MyContract.CallOpts)
}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() view returns((address,bool,bool,bool,bool,uint256)[])
func (_MyContract *MyContractCallerSession) List() ([]MyContractSubcriber, error) {
	return _MyContract.Contract.List(&_MyContract.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MyContract *MyContractCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MyContract *MyContractSession) Creator() (common.Address, error) {
	return _MyContract.Contract.Creator(&_MyContract.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_MyContract *MyContractCallerSession) Creator() (common.Address, error) {
	return _MyContract.Contract.Creator(&_MyContract.CallOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x5d268629.
//
// Solidity: function Refund() returns(string)
func (_MyContract *MyContractTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "Refund")
}

// Refund is a paid mutator transaction binding the contract method 0x5d268629.
//
// Solidity: function Refund() returns(string)
func (_MyContract *MyContractSession) Refund() (*types.Transaction, error) {
	return _MyContract.Contract.Refund(&_MyContract.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x5d268629.
//
// Solidity: function Refund() returns(string)
func (_MyContract *MyContractTransactorSession) Refund() (*types.Transaction, error) {
	return _MyContract.Contract.Refund(&_MyContract.TransactOpts)
}

// Subcribe is a paid mutator transaction binding the contract method 0xc18858fd.
//
// Solidity: function subcribe() payable returns(string)
func (_MyContract *MyContractTransactor) Subcribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "subcribe")
}

// Subcribe is a paid mutator transaction binding the contract method 0xc18858fd.
//
// Solidity: function subcribe() payable returns(string)
func (_MyContract *MyContractSession) Subcribe() (*types.Transaction, error) {
	return _MyContract.Contract.Subcribe(&_MyContract.TransactOpts)
}

// Subcribe is a paid mutator transaction binding the contract method 0xc18858fd.
//
// Solidity: function subcribe() payable returns(string)
func (_MyContract *MyContractTransactorSession) Subcribe() (*types.Transaction, error) {
	return _MyContract.Contract.Subcribe(&_MyContract.TransactOpts)
}

// UnSubcribe is a paid mutator transaction binding the contract method 0x9645c393.
//
// Solidity: function unSubcribe() returns(string)
func (_MyContract *MyContractTransactor) UnSubcribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "unSubcribe")
}

// UnSubcribe is a paid mutator transaction binding the contract method 0x9645c393.
//
// Solidity: function unSubcribe() returns(string)
func (_MyContract *MyContractSession) UnSubcribe() (*types.Transaction, error) {
	return _MyContract.Contract.UnSubcribe(&_MyContract.TransactOpts)
}

// UnSubcribe is a paid mutator transaction binding the contract method 0x9645c393.
//
// Solidity: function unSubcribe() returns(string)
func (_MyContract *MyContractTransactorSession) UnSubcribe() (*types.Transaction, error) {
	return _MyContract.Contract.UnSubcribe(&_MyContract.TransactOpts)
}
