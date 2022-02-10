// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20Holdable

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

// Erc20HoldableMetaData contains all meta data concerning the Erc20Holdable contract.
var Erc20HoldableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HoldCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HoldCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HoldExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"cancelHold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"executeHold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"hold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"holdFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"holds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600d81526020017f4552433230486f6c6461626c65000000000000000000000000000000000000008152506040518060400160405280600681526020017f4552433230480000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620002e5565b508060049080519060200190620000af929190620002e5565b50505033600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000153600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166200012a6200015960201b60201c565b600a6200013891906200052f565b61271062000147919062000580565b6200016260201b60201c565b62000754565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620001d5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001cc9062000642565b60405180910390fd5b620001e960008383620002db60201b60201c565b8060026000828254620001fd919062000664565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825462000254919062000664565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002bb9190620006d2565b60405180910390a3620002d760008383620002e060201b60201c565b5050565b505050565b505050565b828054620002f3906200071e565b90600052602060002090601f01602090048101928262000317576000855562000363565b82601f106200033257805160ff191683800117855562000363565b8280016001018555821562000363579182015b828111156200036257825182559160200191906001019062000345565b5b50905062000372919062000376565b5090565b5b808211156200039157600081600090555060010162000377565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156200042357808604811115620003fb57620003fa62000395565b5b60018516156200040b5780820291505b80810290506200041b85620003c4565b9450620003db565b94509492505050565b6000826200043e576001905062000511565b816200044e576000905062000511565b81600181146200046757600281146200047257620004a8565b600191505062000511565b60ff84111562000487576200048662000395565b5b8360020a915084821115620004a157620004a062000395565b5b5062000511565b5060208310610133831016604e8410600b8410161715620004e25782820a905083811115620004dc57620004db62000395565b5b62000511565b620004f18484846001620003d1565b925090508184048111156200050b576200050a62000395565b5b81810290505b9392505050565b6000819050919050565b600060ff82169050919050565b60006200053c8262000518565b9150620005498362000522565b9250620005787fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200042c565b905092915050565b60006200058d8262000518565b91506200059a8362000518565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620005d657620005d562000395565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200062a601f83620005e1565b91506200063782620005f2565b602082019050919050565b600060208201905081810360008301526200065d816200061b565b9050919050565b6000620006718262000518565b91506200067e8362000518565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620006b657620006b562000395565b5b828201905092915050565b620006cc8162000518565b82525050565b6000602082019050620006e96000830184620006c1565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200073757607f821691505b602082108114156200074e576200074d620006ef565b5b50919050565b6122cf80620007646000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806339509351116100a257806395d89b411161007157806395d89b4114610331578063a457c2d71461034f578063a9059cbb1461037f578063b71b0582146103af578063dd62ed3e146103e157610116565b806339509351146102715780636e3ebe6b146102a157806370369869146102d157806370a082311461030157610116565b8063226aa68a116100e9578063226aa68a146101b757806323b872dd146101e75780632e0f2625146102175780632ff2e9dc14610235578063313ce5671461025357610116565b806306fdde031461011b578063095ea7b31461013957806312c383641461016957806318160ddd14610199575b600080fd5b610123610411565b604051610130919061161d565b60405180910390f35b610153600480360381019061014e91906116d8565b6104a3565b6040516101609190611733565b60405180910390f35b610183600480360381019061017e919061174e565b6104c1565b6040516101909190611733565b60405180910390f35b6101a1610709565b6040516101ae919061178a565b60405180910390f35b6101d160048036038101906101cc91906117a5565b610713565b6040516101de919061178a565b60405180910390f35b61020160048036038101906101fc91906117f8565b610776565b60405161020e9190611733565b60405180910390f35b61021f61086e565b60405161022c9190611867565b60405180910390f35b61023d610873565b60405161024a919061178a565b60405180910390f35b61025b610894565b6040516102689190611867565b60405180910390f35b61028b600480360381019061028691906116d8565b61089d565b6040516102989190611733565b60405180910390f35b6102bb60048036038101906102b69190611882565b610949565b6040516102c8919061178a565b60405180910390f35b6102eb60048036038101906102e6919061174e565b610ac8565b6040516102f89190611733565b60405180910390f35b61031b600480360381019061031691906118e9565b610caf565b604051610328919061178a565b60405180910390f35b610339610cf7565b604051610346919061161d565b60405180910390f35b610369600480360381019061036491906116d8565b610d89565b6040516103769190611733565b60405180910390f35b610399600480360381019061039491906116d8565b610e74565b6040516103a69190611733565b60405180910390f35b6103c960048036038101906103c4919061174e565b610e92565b6040516103d893929190611925565b60405180910390f35b6103fb60048036038101906103f6919061195c565b610efc565b604051610408919061178a565b60405180910390f35b606060038054610420906119cb565b80601f016020809104026020016040519081016040528092919081815260200182805461044c906119cb565b80156104995780601f1061046e57610100808354040283529160200191610499565b820191906000526020600020905b81548152906001019060200180831161047c57829003601f168201915b5050505050905090565b60006104b76104b0610f83565b8484610f8b565b6001905092915050565b6000806008600084815260200190815260200160002090506000816002015411610520576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051790611a49565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610561610f83565b73ffffffffffffffffffffffffffffffffffffffff1614806105d95750610586610f83565b73ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610618576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060f90611adb565b60405180910390fd5b61066d600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360020154611156565b610676836113d7565b7f25bfaf5a77b3c6797250278df23201f88596159d87d6e1af0996ecec716475c9838260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600201546040516106f79493929190611afb565b60405180910390a16001915050919050565b6000600254905090565b6000610728610720610f83565b848487611446565b7f13b8bad3fde196ea364df8943f997584558f925555da397e2f3822e2436aa35984610752610f83565b85856040516107649493929190611afb565b60405180910390a18390509392505050565b6000610783848484611156565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006107ce610f83565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561084e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084590611bb2565b60405180910390fd5b6108628561085a610f83565b858403610f8b565b60019150509392505050565b601281565b601260ff16600a6108849190611d34565b6127106108919190611d7f565b81565b60006012905090565b600061093f6108aa610f83565b8484600160006108b8610f83565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461093a9190611dd9565b610f8b565b6001905092915050565b600080600860008781526020019081526020016000206002015411156109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b90611e7b565b60405180910390fd5b6000600860008781526020019081526020016000209050848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828160020181905550610a7785600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685610776565b507f13b8bad3fde196ea364df8943f997584558f925555da397e2f3822e2436aa35986610aa2610f83565b8686604051610ab49493929190611afb565b60405180910390a185915050949350505050565b600080600860008481526020019081526020016000209050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b21610f83565b73ffffffffffffffffffffffffffffffffffffffff1614610b77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6e90611f0d565b60405180910390fd5b6000816002015411610bbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb590611a49565b60405180910390fd5b610c13600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360020154611156565b7fa0d9f06531460a4718398a225fe1fcd1c16b23ffcb2966d00441e27f14d34c37838260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460020154604051610c949493929190611afb565b60405180910390a1610ca5836113d7565b6001915050919050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060048054610d06906119cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610d32906119cb565b8015610d7f5780601f10610d5457610100808354040283529160200191610d7f565b820191906000526020600020905b815481529060010190602001808311610d6257829003601f168201915b5050505050905090565b60008060016000610d98610f83565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610e55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4c90611f9f565b60405180910390fd5b610e69610e60610f83565b85858403610f8b565b600191505092915050565b6000610e88610e81610f83565b8484611156565b6001905092915050565b60086020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905083565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ffb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff290612031565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561106b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611062906120c3565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051611149919061178a565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156111c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bd90612155565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611236576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122d906121e7565b60405180910390fd5b61124183838361157a565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156112c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112be90612279565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461135a9190611dd9565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113be919061178a565b60405180910390a36113d184848461157f565b50505050565b60086000828152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160009055505050565b6000600860008381526020019081526020016000206002015411156114a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149790611e7b565b60405180910390fd5b6114cd84600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684611156565b6000600860008381526020019081526020016000209050848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508281600201819055505050505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115be5780820151818401526020810190506115a3565b838111156115cd576000848401525b50505050565b6000601f19601f8301169050919050565b60006115ef82611584565b6115f9818561158f565b93506116098185602086016115a0565b611612816115d3565b840191505092915050565b6000602082019050818103600083015261163781846115e4565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061166f82611644565b9050919050565b61167f81611664565b811461168a57600080fd5b50565b60008135905061169c81611676565b92915050565b6000819050919050565b6116b5816116a2565b81146116c057600080fd5b50565b6000813590506116d2816116ac565b92915050565b600080604083850312156116ef576116ee61163f565b5b60006116fd8582860161168d565b925050602061170e858286016116c3565b9150509250929050565b60008115159050919050565b61172d81611718565b82525050565b60006020820190506117486000830184611724565b92915050565b6000602082840312156117645761176361163f565b5b6000611772848285016116c3565b91505092915050565b611784816116a2565b82525050565b600060208201905061179f600083018461177b565b92915050565b6000806000606084860312156117be576117bd61163f565b5b60006117cc868287016116c3565b93505060206117dd8682870161168d565b92505060406117ee868287016116c3565b9150509250925092565b6000806000606084860312156118115761181061163f565b5b600061181f8682870161168d565b93505060206118308682870161168d565b9250506040611841868287016116c3565b9150509250925092565b600060ff82169050919050565b6118618161184b565b82525050565b600060208201905061187c6000830184611858565b92915050565b6000806000806080858703121561189c5761189b61163f565b5b60006118aa878288016116c3565b94505060206118bb8782880161168d565b93505060406118cc8782880161168d565b92505060606118dd878288016116c3565b91505092959194509250565b6000602082840312156118ff576118fe61163f565b5b600061190d8482850161168d565b91505092915050565b61191f81611664565b82525050565b600060608201905061193a6000830186611916565b6119476020830185611916565b611954604083018461177b565b949350505050565b600080604083850312156119735761197261163f565b5b60006119818582860161168d565b92505060206119928582860161168d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806119e357607f821691505b602082108114156119f7576119f661199c565b5b50919050565b7f486f6c6420646f6573206e6f7420657869737400000000000000000000000000600082015250565b6000611a3360138361158f565b9150611a3e826119fd565b602082019050919050565b60006020820190508181036000830152611a6281611a26565b9050919050565b7f4f6e6c7920636f6e7472616374206f776e6572206f7220686f6c64206372656160008201527f746f722063616e20657865637574652074686520686f6c642e00000000000000602082015250565b6000611ac560398361158f565b9150611ad082611a69565b604082019050919050565b60006020820190508181036000830152611af481611ab8565b9050919050565b6000608082019050611b10600083018761177b565b611b1d6020830186611916565b611b2a6040830185611916565b611b37606083018461177b565b95945050505050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b6000611b9c60288361158f565b9150611ba782611b40565b604082019050919050565b60006020820190508181036000830152611bcb81611b8f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b6001851115611c5857808604811115611c3457611c33611bd2565b5b6001851615611c435780820291505b8081029050611c5185611c01565b9450611c18565b94509492505050565b600082611c715760019050611d2d565b81611c7f5760009050611d2d565b8160018114611c955760028114611c9f57611cce565b6001915050611d2d565b60ff841115611cb157611cb0611bd2565b5b8360020a915084821115611cc857611cc7611bd2565b5b50611d2d565b5060208310610133831016604e8410600b8410161715611d035782820a905083811115611cfe57611cfd611bd2565b5b611d2d565b611d108484846001611c0e565b92509050818404811115611d2757611d26611bd2565b5b81810290505b9392505050565b6000611d3f826116a2565b9150611d4a836116a2565b9250611d777fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611c61565b905092915050565b6000611d8a826116a2565b9150611d95836116a2565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611dce57611dcd611bd2565b5b828202905092915050565b6000611de4826116a2565b9150611def836116a2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611e2457611e23611bd2565b5b828201905092915050565b7f486f6c64206964206d75737420626520756e6971756500000000000000000000600082015250565b6000611e6560168361158f565b9150611e7082611e2f565b602082019050919050565b60006020820190508181036000830152611e9481611e58565b9050919050565b7f4f6e6c7920636f6e7472616374206f776e65722063616e20726576657274207460008201527f686520686f6c642e000000000000000000000000000000000000000000000000602082015250565b6000611ef760288361158f565b9150611f0282611e9b565b604082019050919050565b60006020820190508181036000830152611f2681611eea565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611f8960258361158f565b9150611f9482611f2d565b604082019050919050565b60006020820190508181036000830152611fb881611f7c565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061201b60248361158f565b915061202682611fbf565b604082019050919050565b6000602082019050818103600083015261204a8161200e565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006120ad60228361158f565b91506120b882612051565b604082019050919050565b600060208201905081810360008301526120dc816120a0565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061213f60258361158f565b915061214a826120e3565b604082019050919050565b6000602082019050818103600083015261216e81612132565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006121d160238361158f565b91506121dc82612175565b604082019050919050565b60006020820190508181036000830152612200816121c4565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b600061226360268361158f565b915061226e82612207565b604082019050919050565b6000602082019050818103600083015261229281612256565b905091905056fea264697066735822122039acf5d35b28ae9689185c3f6598c357654ea3aa0a8a7f486e930406764e8d6764736f6c634300080b0033",
}

// Erc20HoldableABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20HoldableMetaData.ABI instead.
var Erc20HoldableABI = Erc20HoldableMetaData.ABI

// Erc20HoldableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20HoldableMetaData.Bin instead.
var Erc20HoldableBin = Erc20HoldableMetaData.Bin

// DeployErc20Holdable deploys a new Ethereum contract, binding an instance of Erc20Holdable to it.
func DeployErc20Holdable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Holdable, error) {
	parsed, err := Erc20HoldableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20HoldableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Holdable{Erc20HoldableCaller: Erc20HoldableCaller{contract: contract}, Erc20HoldableTransactor: Erc20HoldableTransactor{contract: contract}, Erc20HoldableFilterer: Erc20HoldableFilterer{contract: contract}}, nil
}

// Erc20Holdable is an auto generated Go binding around an Ethereum contract.
type Erc20Holdable struct {
	Erc20HoldableCaller     // Read-only binding to the contract
	Erc20HoldableTransactor // Write-only binding to the contract
	Erc20HoldableFilterer   // Log filterer for contract events
}

// Erc20HoldableCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20HoldableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HoldableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20HoldableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HoldableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20HoldableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HoldableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20HoldableSession struct {
	Contract     *Erc20Holdable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20HoldableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20HoldableCallerSession struct {
	Contract *Erc20HoldableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Erc20HoldableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20HoldableTransactorSession struct {
	Contract     *Erc20HoldableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Erc20HoldableRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20HoldableRaw struct {
	Contract *Erc20Holdable // Generic contract binding to access the raw methods on
}

// Erc20HoldableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20HoldableCallerRaw struct {
	Contract *Erc20HoldableCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20HoldableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20HoldableTransactorRaw struct {
	Contract *Erc20HoldableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Holdable creates a new instance of Erc20Holdable, bound to a specific deployed contract.
func NewErc20Holdable(address common.Address, backend bind.ContractBackend) (*Erc20Holdable, error) {
	contract, err := bindErc20Holdable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Holdable{Erc20HoldableCaller: Erc20HoldableCaller{contract: contract}, Erc20HoldableTransactor: Erc20HoldableTransactor{contract: contract}, Erc20HoldableFilterer: Erc20HoldableFilterer{contract: contract}}, nil
}

// NewErc20HoldableCaller creates a new read-only instance of Erc20Holdable, bound to a specific deployed contract.
func NewErc20HoldableCaller(address common.Address, caller bind.ContractCaller) (*Erc20HoldableCaller, error) {
	contract, err := bindErc20Holdable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableCaller{contract: contract}, nil
}

// NewErc20HoldableTransactor creates a new write-only instance of Erc20Holdable, bound to a specific deployed contract.
func NewErc20HoldableTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20HoldableTransactor, error) {
	contract, err := bindErc20Holdable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableTransactor{contract: contract}, nil
}

// NewErc20HoldableFilterer creates a new log filterer instance of Erc20Holdable, bound to a specific deployed contract.
func NewErc20HoldableFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20HoldableFilterer, error) {
	contract, err := bindErc20Holdable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableFilterer{contract: contract}, nil
}

// bindErc20Holdable binds a generic wrapper to an already deployed contract.
func bindErc20Holdable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20HoldableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Holdable *Erc20HoldableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Holdable.Contract.Erc20HoldableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Holdable *Erc20HoldableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Erc20HoldableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Holdable *Erc20HoldableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Erc20HoldableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Holdable *Erc20HoldableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Holdable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Holdable *Erc20HoldableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Holdable *Erc20HoldableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.contract.Transact(opts, method, params...)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableSession) DECIMALS() (uint8, error) {
	return _Erc20Holdable.Contract.DECIMALS(&_Erc20Holdable.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableCallerSession) DECIMALS() (uint8, error) {
	return _Erc20Holdable.Contract.DECIMALS(&_Erc20Holdable.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) INITIALSUPPLY() (*big.Int, error) {
	return _Erc20Holdable.Contract.INITIALSUPPLY(&_Erc20Holdable.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _Erc20Holdable.Contract.INITIALSUPPLY(&_Erc20Holdable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Holdable.Contract.Allowance(&_Erc20Holdable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Holdable.Contract.Allowance(&_Erc20Holdable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20Holdable.Contract.BalanceOf(&_Erc20Holdable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20Holdable.Contract.BalanceOf(&_Erc20Holdable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableSession) Decimals() (uint8, error) {
	return _Erc20Holdable.Contract.Decimals(&_Erc20Holdable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Holdable *Erc20HoldableCallerSession) Decimals() (uint8, error) {
	return _Erc20Holdable.Contract.Decimals(&_Erc20Holdable.CallOpts)
}

// Holds is a free data retrieval call binding the contract method 0xb71b0582.
//
// Solidity: function holds(uint256 ) view returns(address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableCaller) Holds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
}, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "holds", arg0)

	outstruct := new(struct {
		Sender    common.Address
		Recipient common.Address
		Amount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Holds is a free data retrieval call binding the contract method 0xb71b0582.
//
// Solidity: function holds(uint256 ) view returns(address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableSession) Holds(arg0 *big.Int) (struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
}, error) {
	return _Erc20Holdable.Contract.Holds(&_Erc20Holdable.CallOpts, arg0)
}

// Holds is a free data retrieval call binding the contract method 0xb71b0582.
//
// Solidity: function holds(uint256 ) view returns(address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableCallerSession) Holds(arg0 *big.Int) (struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
}, error) {
	return _Erc20Holdable.Contract.Holds(&_Erc20Holdable.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20Holdable *Erc20HoldableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20Holdable *Erc20HoldableSession) Name() (string, error) {
	return _Erc20Holdable.Contract.Name(&_Erc20Holdable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20Holdable *Erc20HoldableCallerSession) Name() (string, error) {
	return _Erc20Holdable.Contract.Name(&_Erc20Holdable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20Holdable *Erc20HoldableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20Holdable *Erc20HoldableSession) Symbol() (string, error) {
	return _Erc20Holdable.Contract.Symbol(&_Erc20Holdable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20Holdable *Erc20HoldableCallerSession) Symbol() (string, error) {
	return _Erc20Holdable.Contract.Symbol(&_Erc20Holdable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Holdable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) TotalSupply() (*big.Int, error) {
	return _Erc20Holdable.Contract.TotalSupply(&_Erc20Holdable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Holdable *Erc20HoldableCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20Holdable.Contract.TotalSupply(&_Erc20Holdable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Approve(&_Erc20Holdable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Approve(&_Erc20Holdable.TransactOpts, spender, amount)
}

// CancelHold is a paid mutator transaction binding the contract method 0x70369869.
//
// Solidity: function cancelHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) CancelHold(opts *bind.TransactOpts, holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "cancelHold", holdId)
}

// CancelHold is a paid mutator transaction binding the contract method 0x70369869.
//
// Solidity: function cancelHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) CancelHold(holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.CancelHold(&_Erc20Holdable.TransactOpts, holdId)
}

// CancelHold is a paid mutator transaction binding the contract method 0x70369869.
//
// Solidity: function cancelHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) CancelHold(holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.CancelHold(&_Erc20Holdable.TransactOpts, holdId)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.DecreaseAllowance(&_Erc20Holdable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.DecreaseAllowance(&_Erc20Holdable.TransactOpts, spender, subtractedValue)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) ExecuteHold(opts *bind.TransactOpts, holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "executeHold", holdId)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) ExecuteHold(holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.ExecuteHold(&_Erc20Holdable.TransactOpts, holdId)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) ExecuteHold(holdId *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.ExecuteHold(&_Erc20Holdable.TransactOpts, holdId)
}

// Hold is a paid mutator transaction binding the contract method 0x226aa68a.
//
// Solidity: function hold(uint256 holdId, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableTransactor) Hold(opts *bind.TransactOpts, holdId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "hold", holdId, recipient, amount)
}

// Hold is a paid mutator transaction binding the contract method 0x226aa68a.
//
// Solidity: function hold(uint256 holdId, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) Hold(holdId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Hold(&_Erc20Holdable.TransactOpts, holdId, recipient, amount)
}

// Hold is a paid mutator transaction binding the contract method 0x226aa68a.
//
// Solidity: function hold(uint256 holdId, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableTransactorSession) Hold(holdId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Hold(&_Erc20Holdable.TransactOpts, holdId, recipient, amount)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x6e3ebe6b.
//
// Solidity: function holdFrom(uint256 holdId, address sender, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableTransactor) HoldFrom(opts *bind.TransactOpts, holdId *big.Int, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "holdFrom", holdId, sender, recipient, amount)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x6e3ebe6b.
//
// Solidity: function holdFrom(uint256 holdId, address sender, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableSession) HoldFrom(holdId *big.Int, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.HoldFrom(&_Erc20Holdable.TransactOpts, holdId, sender, recipient, amount)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x6e3ebe6b.
//
// Solidity: function holdFrom(uint256 holdId, address sender, address recipient, uint256 amount) returns(uint256)
func (_Erc20Holdable *Erc20HoldableTransactorSession) HoldFrom(holdId *big.Int, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.HoldFrom(&_Erc20Holdable.TransactOpts, holdId, sender, recipient, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.IncreaseAllowance(&_Erc20Holdable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.IncreaseAllowance(&_Erc20Holdable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Transfer(&_Erc20Holdable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.Transfer(&_Erc20Holdable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.TransferFrom(&_Erc20Holdable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20Holdable *Erc20HoldableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20Holdable.Contract.TransferFrom(&_Erc20Holdable.TransactOpts, sender, recipient, amount)
}

// Erc20HoldableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20Holdable contract.
type Erc20HoldableApprovalIterator struct {
	Event *Erc20HoldableApproval // Event containing the contract specifics and raw log

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
func (it *Erc20HoldableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20HoldableApproval)
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
		it.Event = new(Erc20HoldableApproval)
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
func (it *Erc20HoldableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20HoldableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20HoldableApproval represents a Approval event raised by the Erc20Holdable contract.
type Erc20HoldableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20HoldableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Holdable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableApprovalIterator{contract: _Erc20Holdable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20HoldableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Holdable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20HoldableApproval)
				if err := _Erc20Holdable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) ParseApproval(log types.Log) (*Erc20HoldableApproval, error) {
	event := new(Erc20HoldableApproval)
	if err := _Erc20Holdable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20HoldableHoldCanceledIterator is returned from FilterHoldCanceled and is used to iterate over the raw logs and unpacked data for HoldCanceled events raised by the Erc20Holdable contract.
type Erc20HoldableHoldCanceledIterator struct {
	Event *Erc20HoldableHoldCanceled // Event containing the contract specifics and raw log

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
func (it *Erc20HoldableHoldCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20HoldableHoldCanceled)
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
		it.Event = new(Erc20HoldableHoldCanceled)
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
func (it *Erc20HoldableHoldCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20HoldableHoldCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20HoldableHoldCanceled represents a HoldCanceled event raised by the Erc20Holdable contract.
type Erc20HoldableHoldCanceled struct {
	HoldId    *big.Int
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHoldCanceled is a free log retrieval operation binding the contract event 0xa0d9f06531460a4718398a225fe1fcd1c16b23ffcb2966d00441e27f14d34c37.
//
// Solidity: event HoldCanceled(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) FilterHoldCanceled(opts *bind.FilterOpts) (*Erc20HoldableHoldCanceledIterator, error) {

	logs, sub, err := _Erc20Holdable.contract.FilterLogs(opts, "HoldCanceled")
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableHoldCanceledIterator{contract: _Erc20Holdable.contract, event: "HoldCanceled", logs: logs, sub: sub}, nil
}

// WatchHoldCanceled is a free log subscription operation binding the contract event 0xa0d9f06531460a4718398a225fe1fcd1c16b23ffcb2966d00441e27f14d34c37.
//
// Solidity: event HoldCanceled(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) WatchHoldCanceled(opts *bind.WatchOpts, sink chan<- *Erc20HoldableHoldCanceled) (event.Subscription, error) {

	logs, sub, err := _Erc20Holdable.contract.WatchLogs(opts, "HoldCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20HoldableHoldCanceled)
				if err := _Erc20Holdable.contract.UnpackLog(event, "HoldCanceled", log); err != nil {
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

// ParseHoldCanceled is a log parse operation binding the contract event 0xa0d9f06531460a4718398a225fe1fcd1c16b23ffcb2966d00441e27f14d34c37.
//
// Solidity: event HoldCanceled(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) ParseHoldCanceled(log types.Log) (*Erc20HoldableHoldCanceled, error) {
	event := new(Erc20HoldableHoldCanceled)
	if err := _Erc20Holdable.contract.UnpackLog(event, "HoldCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20HoldableHoldCreatedIterator is returned from FilterHoldCreated and is used to iterate over the raw logs and unpacked data for HoldCreated events raised by the Erc20Holdable contract.
type Erc20HoldableHoldCreatedIterator struct {
	Event *Erc20HoldableHoldCreated // Event containing the contract specifics and raw log

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
func (it *Erc20HoldableHoldCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20HoldableHoldCreated)
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
		it.Event = new(Erc20HoldableHoldCreated)
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
func (it *Erc20HoldableHoldCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20HoldableHoldCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20HoldableHoldCreated represents a HoldCreated event raised by the Erc20Holdable contract.
type Erc20HoldableHoldCreated struct {
	HoldId    *big.Int
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHoldCreated is a free log retrieval operation binding the contract event 0x13b8bad3fde196ea364df8943f997584558f925555da397e2f3822e2436aa359.
//
// Solidity: event HoldCreated(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) FilterHoldCreated(opts *bind.FilterOpts) (*Erc20HoldableHoldCreatedIterator, error) {

	logs, sub, err := _Erc20Holdable.contract.FilterLogs(opts, "HoldCreated")
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableHoldCreatedIterator{contract: _Erc20Holdable.contract, event: "HoldCreated", logs: logs, sub: sub}, nil
}

// WatchHoldCreated is a free log subscription operation binding the contract event 0x13b8bad3fde196ea364df8943f997584558f925555da397e2f3822e2436aa359.
//
// Solidity: event HoldCreated(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) WatchHoldCreated(opts *bind.WatchOpts, sink chan<- *Erc20HoldableHoldCreated) (event.Subscription, error) {

	logs, sub, err := _Erc20Holdable.contract.WatchLogs(opts, "HoldCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20HoldableHoldCreated)
				if err := _Erc20Holdable.contract.UnpackLog(event, "HoldCreated", log); err != nil {
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

// ParseHoldCreated is a log parse operation binding the contract event 0x13b8bad3fde196ea364df8943f997584558f925555da397e2f3822e2436aa359.
//
// Solidity: event HoldCreated(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) ParseHoldCreated(log types.Log) (*Erc20HoldableHoldCreated, error) {
	event := new(Erc20HoldableHoldCreated)
	if err := _Erc20Holdable.contract.UnpackLog(event, "HoldCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20HoldableHoldExecutedIterator is returned from FilterHoldExecuted and is used to iterate over the raw logs and unpacked data for HoldExecuted events raised by the Erc20Holdable contract.
type Erc20HoldableHoldExecutedIterator struct {
	Event *Erc20HoldableHoldExecuted // Event containing the contract specifics and raw log

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
func (it *Erc20HoldableHoldExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20HoldableHoldExecuted)
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
		it.Event = new(Erc20HoldableHoldExecuted)
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
func (it *Erc20HoldableHoldExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20HoldableHoldExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20HoldableHoldExecuted represents a HoldExecuted event raised by the Erc20Holdable contract.
type Erc20HoldableHoldExecuted struct {
	HoldId    *big.Int
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHoldExecuted is a free log retrieval operation binding the contract event 0x25bfaf5a77b3c6797250278df23201f88596159d87d6e1af0996ecec716475c9.
//
// Solidity: event HoldExecuted(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) FilterHoldExecuted(opts *bind.FilterOpts) (*Erc20HoldableHoldExecutedIterator, error) {

	logs, sub, err := _Erc20Holdable.contract.FilterLogs(opts, "HoldExecuted")
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableHoldExecutedIterator{contract: _Erc20Holdable.contract, event: "HoldExecuted", logs: logs, sub: sub}, nil
}

// WatchHoldExecuted is a free log subscription operation binding the contract event 0x25bfaf5a77b3c6797250278df23201f88596159d87d6e1af0996ecec716475c9.
//
// Solidity: event HoldExecuted(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) WatchHoldExecuted(opts *bind.WatchOpts, sink chan<- *Erc20HoldableHoldExecuted) (event.Subscription, error) {

	logs, sub, err := _Erc20Holdable.contract.WatchLogs(opts, "HoldExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20HoldableHoldExecuted)
				if err := _Erc20Holdable.contract.UnpackLog(event, "HoldExecuted", log); err != nil {
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

// ParseHoldExecuted is a log parse operation binding the contract event 0x25bfaf5a77b3c6797250278df23201f88596159d87d6e1af0996ecec716475c9.
//
// Solidity: event HoldExecuted(uint256 holdId, address sender, address recipient, uint256 amount)
func (_Erc20Holdable *Erc20HoldableFilterer) ParseHoldExecuted(log types.Log) (*Erc20HoldableHoldExecuted, error) {
	event := new(Erc20HoldableHoldExecuted)
	if err := _Erc20Holdable.contract.UnpackLog(event, "HoldExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20HoldableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20Holdable contract.
type Erc20HoldableTransferIterator struct {
	Event *Erc20HoldableTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20HoldableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20HoldableTransfer)
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
		it.Event = new(Erc20HoldableTransfer)
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
func (it *Erc20HoldableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20HoldableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20HoldableTransfer represents a Transfer event raised by the Erc20Holdable contract.
type Erc20HoldableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20HoldableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Holdable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20HoldableTransferIterator{contract: _Erc20Holdable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20HoldableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Holdable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20HoldableTransfer)
				if err := _Erc20Holdable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Holdable *Erc20HoldableFilterer) ParseTransfer(log types.Log) (*Erc20HoldableTransfer, error) {
	event := new(Erc20HoldableTransfer)
	if err := _Erc20Holdable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
