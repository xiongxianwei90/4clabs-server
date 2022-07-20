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
)

// IComicWorksComicWorks is an auto generated low-level Go binding around an user-defined struct.
type IComicWorksComicWorks struct {
	Name            string
	ImageURIs       []string
	Limit           *big.Int
	Price           *big.Int
	Author          common.Address
	BasedOnContract common.Address
	BasedOnTokenId  *big.Int
}

// ForClabsMetaData contains all meta data concerning the ForClabs contract.
var ForClabsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"imageURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"basedOnContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIComicWorks.ComicWorks\",\"name\":\"comicWorks\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"}],\"name\":\"ComicWorksCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"}],\"name\":\"ComicWorksLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"}],\"name\":\"ComicWorksPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"imageURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"basedOnContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIComicWorks.ComicWorks\",\"name\":\"comicWorks\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ComicWorksSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ComicWorksTokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"ContractAndTokenAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"authorizeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"authorizeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"basedOnContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"basedOnTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"comicWorks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"imageURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"basedOnContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIComicWorks.ComicWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"comicWorksCountCreatedByOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"comicWorksIdByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"comicWorksIdOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"comicWorksIdWithTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"basedOnContract_\",\"type\":\"address\"}],\"name\":\"comicWorksTokenIdsBasedOnContract\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"basedOnContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId_\",\"type\":\"uint256\"}],\"name\":\"comicWorksTokenIdsBasedOnContractAndTokenId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"}],\"name\":\"comicWorksWithWorksId\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"imageURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"basedOnContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIComicWorks.ComicWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"contractAndTokenAuthorizedBy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imageURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imageURIs\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isContractAndTokenAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"isContractAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"imageURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"basedOnContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basedOnTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIComicWorks.ComicWorks\",\"name\":\"works\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"setLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"worksId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalComicWorksCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ForClabsABI is the input ABI used to generate the binding from.
// Deprecated: Use ForClabsMetaData.ABI instead.
var ForClabsABI = ForClabsMetaData.ABI

// ForClabs is an auto generated Go binding around an Ethereum contract.
type ForClabs struct {
	ForClabsCaller     // Read-only binding to the contract
	ForClabsTransactor // Write-only binding to the contract
	ForClabsFilterer   // Log filterer for contract events
}

// ForClabsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForClabsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForClabsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForClabsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForClabsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForClabsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForClabsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForClabsSession struct {
	Contract     *ForClabs         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForClabsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForClabsCallerSession struct {
	Contract *ForClabsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ForClabsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForClabsTransactorSession struct {
	Contract     *ForClabsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ForClabsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForClabsRaw struct {
	Contract *ForClabs // Generic contract binding to access the raw methods on
}

// ForClabsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForClabsCallerRaw struct {
	Contract *ForClabsCaller // Generic read-only contract binding to access the raw methods on
}

// ForClabsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForClabsTransactorRaw struct {
	Contract *ForClabsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForClabs creates a new instance of ForClabs, bound to a specific deployed contract.
func NewForClabs(address common.Address, backend bind.ContractBackend) (*ForClabs, error) {
	contract, err := bindForClabs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForClabs{ForClabsCaller: ForClabsCaller{contract: contract}, ForClabsTransactor: ForClabsTransactor{contract: contract}, ForClabsFilterer: ForClabsFilterer{contract: contract}}, nil
}

// NewForClabsCaller creates a new read-only instance of ForClabs, bound to a specific deployed contract.
func NewForClabsCaller(address common.Address, caller bind.ContractCaller) (*ForClabsCaller, error) {
	contract, err := bindForClabs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForClabsCaller{contract: contract}, nil
}

// NewForClabsTransactor creates a new write-only instance of ForClabs, bound to a specific deployed contract.
func NewForClabsTransactor(address common.Address, transactor bind.ContractTransactor) (*ForClabsTransactor, error) {
	contract, err := bindForClabs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForClabsTransactor{contract: contract}, nil
}

// NewForClabsFilterer creates a new log filterer instance of ForClabs, bound to a specific deployed contract.
func NewForClabsFilterer(address common.Address, filterer bind.ContractFilterer) (*ForClabsFilterer, error) {
	contract, err := bindForClabs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForClabsFilterer{contract: contract}, nil
}

// bindForClabs binds a generic wrapper to an already deployed contract.
func bindForClabs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForClabsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForClabs *ForClabsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForClabs.Contract.ForClabsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForClabs *ForClabsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForClabs.Contract.ForClabsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForClabs *ForClabsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForClabs.Contract.ForClabsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForClabs *ForClabsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForClabs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForClabs *ForClabsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForClabs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForClabs *ForClabsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForClabs.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ForClabs *ForClabsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ForClabs *ForClabsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ForClabs.Contract.BalanceOf(&_ForClabs.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ForClabs.Contract.BalanceOf(&_ForClabs.CallOpts, owner)
}

// BasedOnContract is a free data retrieval call binding the contract method 0x8273e81c.
//
// Solidity: function basedOnContract(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCaller) BasedOnContract(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "basedOnContract", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasedOnContract is a free data retrieval call binding the contract method 0x8273e81c.
//
// Solidity: function basedOnContract(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsSession) BasedOnContract(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.BasedOnContract(&_ForClabs.CallOpts, tokenId)
}

// BasedOnContract is a free data retrieval call binding the contract method 0x8273e81c.
//
// Solidity: function basedOnContract(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCallerSession) BasedOnContract(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.BasedOnContract(&_ForClabs.CallOpts, tokenId)
}

// BasedOnTokenId is a free data retrieval call binding the contract method 0xdeb8a02c.
//
// Solidity: function basedOnTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCaller) BasedOnTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "basedOnTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasedOnTokenId is a free data retrieval call binding the contract method 0xdeb8a02c.
//
// Solidity: function basedOnTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsSession) BasedOnTokenId(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.BasedOnTokenId(&_ForClabs.CallOpts, tokenId)
}

// BasedOnTokenId is a free data retrieval call binding the contract method 0xdeb8a02c.
//
// Solidity: function basedOnTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) BasedOnTokenId(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.BasedOnTokenId(&_ForClabs.CallOpts, tokenId)
}

// ComicWorks is a free data retrieval call binding the contract method 0x11393a1a.
//
// Solidity: function comicWorks(uint256 tokenId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsCaller) ComicWorks(opts *bind.CallOpts, tokenId *big.Int) (IComicWorksComicWorks, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorks", tokenId)

	if err != nil {
		return *new(IComicWorksComicWorks), err
	}

	out0 := *abi.ConvertType(out[0], new(IComicWorksComicWorks)).(*IComicWorksComicWorks)

	return out0, err

}

// ComicWorks is a free data retrieval call binding the contract method 0x11393a1a.
//
// Solidity: function comicWorks(uint256 tokenId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsSession) ComicWorks(tokenId *big.Int) (IComicWorksComicWorks, error) {
	return _ForClabs.Contract.ComicWorks(&_ForClabs.CallOpts, tokenId)
}

// ComicWorks is a free data retrieval call binding the contract method 0x11393a1a.
//
// Solidity: function comicWorks(uint256 tokenId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsCallerSession) ComicWorks(tokenId *big.Int) (IComicWorksComicWorks, error) {
	return _ForClabs.Contract.ComicWorks(&_ForClabs.CallOpts, tokenId)
}

// ComicWorksCountCreatedByOwner is a free data retrieval call binding the contract method 0xb0b34681.
//
// Solidity: function comicWorksCountCreatedByOwner(address owner) view returns(uint256)
func (_ForClabs *ForClabsCaller) ComicWorksCountCreatedByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksCountCreatedByOwner", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComicWorksCountCreatedByOwner is a free data retrieval call binding the contract method 0xb0b34681.
//
// Solidity: function comicWorksCountCreatedByOwner(address owner) view returns(uint256)
func (_ForClabs *ForClabsSession) ComicWorksCountCreatedByOwner(owner common.Address) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksCountCreatedByOwner(&_ForClabs.CallOpts, owner)
}

// ComicWorksCountCreatedByOwner is a free data retrieval call binding the contract method 0xb0b34681.
//
// Solidity: function comicWorksCountCreatedByOwner(address owner) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) ComicWorksCountCreatedByOwner(owner common.Address) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksCountCreatedByOwner(&_ForClabs.CallOpts, owner)
}

// ComicWorksIdByIndex is a free data retrieval call binding the contract method 0x16cb2c01.
//
// Solidity: function comicWorksIdByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCaller) ComicWorksIdByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksIdByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComicWorksIdByIndex is a free data retrieval call binding the contract method 0x16cb2c01.
//
// Solidity: function comicWorksIdByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsSession) ComicWorksIdByIndex(index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdByIndex(&_ForClabs.CallOpts, index)
}

// ComicWorksIdByIndex is a free data retrieval call binding the contract method 0x16cb2c01.
//
// Solidity: function comicWorksIdByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) ComicWorksIdByIndex(index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdByIndex(&_ForClabs.CallOpts, index)
}

// ComicWorksIdOfOwnerByIndex is a free data retrieval call binding the contract method 0x80a3dc18.
//
// Solidity: function comicWorksIdOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCaller) ComicWorksIdOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksIdOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComicWorksIdOfOwnerByIndex is a free data retrieval call binding the contract method 0x80a3dc18.
//
// Solidity: function comicWorksIdOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsSession) ComicWorksIdOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdOfOwnerByIndex(&_ForClabs.CallOpts, owner, index)
}

// ComicWorksIdOfOwnerByIndex is a free data retrieval call binding the contract method 0x80a3dc18.
//
// Solidity: function comicWorksIdOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) ComicWorksIdOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdOfOwnerByIndex(&_ForClabs.CallOpts, owner, index)
}

// ComicWorksIdWithTokenId is a free data retrieval call binding the contract method 0xc38344aa.
//
// Solidity: function comicWorksIdWithTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCaller) ComicWorksIdWithTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksIdWithTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComicWorksIdWithTokenId is a free data retrieval call binding the contract method 0xc38344aa.
//
// Solidity: function comicWorksIdWithTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsSession) ComicWorksIdWithTokenId(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdWithTokenId(&_ForClabs.CallOpts, tokenId)
}

// ComicWorksIdWithTokenId is a free data retrieval call binding the contract method 0xc38344aa.
//
// Solidity: function comicWorksIdWithTokenId(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) ComicWorksIdWithTokenId(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.ComicWorksIdWithTokenId(&_ForClabs.CallOpts, tokenId)
}

// ComicWorksTokenIdsBasedOnContract is a free data retrieval call binding the contract method 0xe30545ea.
//
// Solidity: function comicWorksTokenIdsBasedOnContract(address basedOnContract_) view returns(uint256[])
func (_ForClabs *ForClabsCaller) ComicWorksTokenIdsBasedOnContract(opts *bind.CallOpts, basedOnContract_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksTokenIdsBasedOnContract", basedOnContract_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComicWorksTokenIdsBasedOnContract is a free data retrieval call binding the contract method 0xe30545ea.
//
// Solidity: function comicWorksTokenIdsBasedOnContract(address basedOnContract_) view returns(uint256[])
func (_ForClabs *ForClabsSession) ComicWorksTokenIdsBasedOnContract(basedOnContract_ common.Address) ([]*big.Int, error) {
	return _ForClabs.Contract.ComicWorksTokenIdsBasedOnContract(&_ForClabs.CallOpts, basedOnContract_)
}

// ComicWorksTokenIdsBasedOnContract is a free data retrieval call binding the contract method 0xe30545ea.
//
// Solidity: function comicWorksTokenIdsBasedOnContract(address basedOnContract_) view returns(uint256[])
func (_ForClabs *ForClabsCallerSession) ComicWorksTokenIdsBasedOnContract(basedOnContract_ common.Address) ([]*big.Int, error) {
	return _ForClabs.Contract.ComicWorksTokenIdsBasedOnContract(&_ForClabs.CallOpts, basedOnContract_)
}

// ComicWorksTokenIdsBasedOnContractAndTokenId is a free data retrieval call binding the contract method 0x4fcfa474.
//
// Solidity: function comicWorksTokenIdsBasedOnContractAndTokenId(address basedOnContract_, uint256 basedOnTokenId_) view returns(uint256[])
func (_ForClabs *ForClabsCaller) ComicWorksTokenIdsBasedOnContractAndTokenId(opts *bind.CallOpts, basedOnContract_ common.Address, basedOnTokenId_ *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksTokenIdsBasedOnContractAndTokenId", basedOnContract_, basedOnTokenId_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ComicWorksTokenIdsBasedOnContractAndTokenId is a free data retrieval call binding the contract method 0x4fcfa474.
//
// Solidity: function comicWorksTokenIdsBasedOnContractAndTokenId(address basedOnContract_, uint256 basedOnTokenId_) view returns(uint256[])
func (_ForClabs *ForClabsSession) ComicWorksTokenIdsBasedOnContractAndTokenId(basedOnContract_ common.Address, basedOnTokenId_ *big.Int) ([]*big.Int, error) {
	return _ForClabs.Contract.ComicWorksTokenIdsBasedOnContractAndTokenId(&_ForClabs.CallOpts, basedOnContract_, basedOnTokenId_)
}

// ComicWorksTokenIdsBasedOnContractAndTokenId is a free data retrieval call binding the contract method 0x4fcfa474.
//
// Solidity: function comicWorksTokenIdsBasedOnContractAndTokenId(address basedOnContract_, uint256 basedOnTokenId_) view returns(uint256[])
func (_ForClabs *ForClabsCallerSession) ComicWorksTokenIdsBasedOnContractAndTokenId(basedOnContract_ common.Address, basedOnTokenId_ *big.Int) ([]*big.Int, error) {
	return _ForClabs.Contract.ComicWorksTokenIdsBasedOnContractAndTokenId(&_ForClabs.CallOpts, basedOnContract_, basedOnTokenId_)
}

// ComicWorksWithWorksId is a free data retrieval call binding the contract method 0x0c6e3b9c.
//
// Solidity: function comicWorksWithWorksId(uint256 worksId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsCaller) ComicWorksWithWorksId(opts *bind.CallOpts, worksId *big.Int) (IComicWorksComicWorks, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "comicWorksWithWorksId", worksId)

	if err != nil {
		return *new(IComicWorksComicWorks), err
	}

	out0 := *abi.ConvertType(out[0], new(IComicWorksComicWorks)).(*IComicWorksComicWorks)

	return out0, err

}

// ComicWorksWithWorksId is a free data retrieval call binding the contract method 0x0c6e3b9c.
//
// Solidity: function comicWorksWithWorksId(uint256 worksId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsSession) ComicWorksWithWorksId(worksId *big.Int) (IComicWorksComicWorks, error) {
	return _ForClabs.Contract.ComicWorksWithWorksId(&_ForClabs.CallOpts, worksId)
}

// ComicWorksWithWorksId is a free data retrieval call binding the contract method 0x0c6e3b9c.
//
// Solidity: function comicWorksWithWorksId(uint256 worksId) view returns((string,string[],uint256,uint256,address,address,uint256))
func (_ForClabs *ForClabsCallerSession) ComicWorksWithWorksId(worksId *big.Int) (IComicWorksComicWorks, error) {
	return _ForClabs.Contract.ComicWorksWithWorksId(&_ForClabs.CallOpts, worksId)
}

// ContractAndTokenAuthorizedBy is a free data retrieval call binding the contract method 0x378e5b09.
//
// Solidity: function contractAndTokenAuthorizedBy(address contractAddress, uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCaller) ContractAndTokenAuthorizedBy(opts *bind.CallOpts, contractAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "contractAndTokenAuthorizedBy", contractAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAndTokenAuthorizedBy is a free data retrieval call binding the contract method 0x378e5b09.
//
// Solidity: function contractAndTokenAuthorizedBy(address contractAddress, uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsSession) ContractAndTokenAuthorizedBy(contractAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.ContractAndTokenAuthorizedBy(&_ForClabs.CallOpts, contractAddress, tokenId)
}

// ContractAndTokenAuthorizedBy is a free data retrieval call binding the contract method 0x378e5b09.
//
// Solidity: function contractAndTokenAuthorizedBy(address contractAddress, uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCallerSession) ContractAndTokenAuthorizedBy(contractAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.ContractAndTokenAuthorizedBy(&_ForClabs.CallOpts, contractAddress, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.GetApproved(&_ForClabs.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.GetApproved(&_ForClabs.CallOpts, tokenId)
}

// ImageURI is a free data retrieval call binding the contract method 0x8f742d16.
//
// Solidity: function imageURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCaller) ImageURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "imageURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageURI is a free data retrieval call binding the contract method 0x8f742d16.
//
// Solidity: function imageURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsSession) ImageURI(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.ImageURI(&_ForClabs.CallOpts, tokenId)
}

// ImageURI is a free data retrieval call binding the contract method 0x8f742d16.
//
// Solidity: function imageURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCallerSession) ImageURI(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.ImageURI(&_ForClabs.CallOpts, tokenId)
}

// ImageURIs is a free data retrieval call binding the contract method 0xd67d9246.
//
// Solidity: function imageURIs(uint256 tokenId) view returns(string[])
func (_ForClabs *ForClabsCaller) ImageURIs(opts *bind.CallOpts, tokenId *big.Int) ([]string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "imageURIs", tokenId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ImageURIs is a free data retrieval call binding the contract method 0xd67d9246.
//
// Solidity: function imageURIs(uint256 tokenId) view returns(string[])
func (_ForClabs *ForClabsSession) ImageURIs(tokenId *big.Int) ([]string, error) {
	return _ForClabs.Contract.ImageURIs(&_ForClabs.CallOpts, tokenId)
}

// ImageURIs is a free data retrieval call binding the contract method 0xd67d9246.
//
// Solidity: function imageURIs(uint256 tokenId) view returns(string[])
func (_ForClabs *ForClabsCallerSession) ImageURIs(tokenId *big.Int) ([]string, error) {
	return _ForClabs.Contract.ImageURIs(&_ForClabs.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ForClabs *ForClabsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ForClabs *ForClabsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ForClabs.Contract.IsApprovedForAll(&_ForClabs.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ForClabs *ForClabsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ForClabs.Contract.IsApprovedForAll(&_ForClabs.CallOpts, owner, operator)
}

// IsContractAndTokenAuthorized is a free data retrieval call binding the contract method 0x4f8a0453.
//
// Solidity: function isContractAndTokenAuthorized(address contractAddress, uint256 tokenId) view returns(bool)
func (_ForClabs *ForClabsCaller) IsContractAndTokenAuthorized(opts *bind.CallOpts, contractAddress common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "isContractAndTokenAuthorized", contractAddress, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractAndTokenAuthorized is a free data retrieval call binding the contract method 0x4f8a0453.
//
// Solidity: function isContractAndTokenAuthorized(address contractAddress, uint256 tokenId) view returns(bool)
func (_ForClabs *ForClabsSession) IsContractAndTokenAuthorized(contractAddress common.Address, tokenId *big.Int) (bool, error) {
	return _ForClabs.Contract.IsContractAndTokenAuthorized(&_ForClabs.CallOpts, contractAddress, tokenId)
}

// IsContractAndTokenAuthorized is a free data retrieval call binding the contract method 0x4f8a0453.
//
// Solidity: function isContractAndTokenAuthorized(address contractAddress, uint256 tokenId) view returns(bool)
func (_ForClabs *ForClabsCallerSession) IsContractAndTokenAuthorized(contractAddress common.Address, tokenId *big.Int) (bool, error) {
	return _ForClabs.Contract.IsContractAndTokenAuthorized(&_ForClabs.CallOpts, contractAddress, tokenId)
}

// IsContractAuthorized is a free data retrieval call binding the contract method 0x103aeda7.
//
// Solidity: function isContractAuthorized(address contractAddress) view returns(bool)
func (_ForClabs *ForClabsCaller) IsContractAuthorized(opts *bind.CallOpts, contractAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "isContractAuthorized", contractAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractAuthorized is a free data retrieval call binding the contract method 0x103aeda7.
//
// Solidity: function isContractAuthorized(address contractAddress) view returns(bool)
func (_ForClabs *ForClabsSession) IsContractAuthorized(contractAddress common.Address) (bool, error) {
	return _ForClabs.Contract.IsContractAuthorized(&_ForClabs.CallOpts, contractAddress)
}

// IsContractAuthorized is a free data retrieval call binding the contract method 0x103aeda7.
//
// Solidity: function isContractAuthorized(address contractAddress) view returns(bool)
func (_ForClabs *ForClabsCallerSession) IsContractAuthorized(contractAddress common.Address) (bool, error) {
	return _ForClabs.Contract.IsContractAuthorized(&_ForClabs.CallOpts, contractAddress)
}

// Limit is a free data retrieval call binding the contract method 0x243bcfcb.
//
// Solidity: function limit(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCaller) Limit(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "limit", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0x243bcfcb.
//
// Solidity: function limit(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsSession) Limit(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.Limit(&_ForClabs.CallOpts, tokenId)
}

// Limit is a free data retrieval call binding the contract method 0x243bcfcb.
//
// Solidity: function limit(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) Limit(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.Limit(&_ForClabs.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCaller) Name(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "name", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsSession) Name(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.Name(&_ForClabs.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCallerSession) Name(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.Name(&_ForClabs.CallOpts, tokenId)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ForClabs *ForClabsCaller) Name0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "name0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ForClabs *ForClabsSession) Name0() (string, error) {
	return _ForClabs.Contract.Name0(&_ForClabs.CallOpts)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ForClabs *ForClabsCallerSession) Name0() (string, error) {
	return _ForClabs.Contract.Name0(&_ForClabs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForClabs *ForClabsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForClabs *ForClabsSession) Owner() (common.Address, error) {
	return _ForClabs.Contract.Owner(&_ForClabs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForClabs *ForClabsCallerSession) Owner() (common.Address, error) {
	return _ForClabs.Contract.Owner(&_ForClabs.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.OwnerOf(&_ForClabs.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ForClabs *ForClabsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ForClabs.Contract.OwnerOf(&_ForClabs.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCaller) Price(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "price", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsSession) Price(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.Price(&_ForClabs.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0x26a49e37.
//
// Solidity: function price(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) Price(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.Price(&_ForClabs.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ForClabs *ForClabsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ForClabs *ForClabsSession) ProxiableUUID() ([32]byte, error) {
	return _ForClabs.Contract.ProxiableUUID(&_ForClabs.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ForClabs *ForClabsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ForClabs.Contract.ProxiableUUID(&_ForClabs.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ForClabs *ForClabsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ForClabs *ForClabsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ForClabs.Contract.SupportsInterface(&_ForClabs.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ForClabs *ForClabsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ForClabs.Contract.SupportsInterface(&_ForClabs.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ForClabs *ForClabsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ForClabs *ForClabsSession) Symbol() (string, error) {
	return _ForClabs.Contract.Symbol(&_ForClabs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ForClabs *ForClabsCallerSession) Symbol() (string, error) {
	return _ForClabs.Contract.Symbol(&_ForClabs.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TokenByIndex(&_ForClabs.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TokenByIndex(&_ForClabs.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TokenOfOwnerByIndex(&_ForClabs.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TokenOfOwnerByIndex(&_ForClabs.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.TokenURI(&_ForClabs.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ForClabs *ForClabsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ForClabs.Contract.TokenURI(&_ForClabs.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_ForClabs *ForClabsCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_ForClabs *ForClabsSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _ForClabs.Contract.TokensOfOwner(&_ForClabs.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_ForClabs *ForClabsCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _ForClabs.Contract.TokensOfOwner(&_ForClabs.CallOpts, _owner)
}

// TotalComicWorksCount is a free data retrieval call binding the contract method 0xdbf6e1c7.
//
// Solidity: function totalComicWorksCount() view returns(uint256)
func (_ForClabs *ForClabsCaller) TotalComicWorksCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "totalComicWorksCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalComicWorksCount is a free data retrieval call binding the contract method 0xdbf6e1c7.
//
// Solidity: function totalComicWorksCount() view returns(uint256)
func (_ForClabs *ForClabsSession) TotalComicWorksCount() (*big.Int, error) {
	return _ForClabs.Contract.TotalComicWorksCount(&_ForClabs.CallOpts)
}

// TotalComicWorksCount is a free data retrieval call binding the contract method 0xdbf6e1c7.
//
// Solidity: function totalComicWorksCount() view returns(uint256)
func (_ForClabs *ForClabsCallerSession) TotalComicWorksCount() (*big.Int, error) {
	return _ForClabs.Contract.TotalComicWorksCount(&_ForClabs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ForClabs *ForClabsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ForClabs *ForClabsSession) TotalSupply() (*big.Int, error) {
	return _ForClabs.Contract.TotalSupply(&_ForClabs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ForClabs *ForClabsCallerSession) TotalSupply() (*big.Int, error) {
	return _ForClabs.Contract.TotalSupply(&_ForClabs.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCaller) TotalSupply0(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "totalSupply0", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsSession) TotalSupply0(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TotalSupply0(&_ForClabs.CallOpts, tokenId)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ForClabs *ForClabsCallerSession) TotalSupply0(tokenId *big.Int) (*big.Int, error) {
	return _ForClabs.Contract.TotalSupply0(&_ForClabs.CallOpts, tokenId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ForClabs *ForClabsCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ForClabs.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ForClabs *ForClabsSession) Version() (*big.Int, error) {
	return _ForClabs.Contract.Version(&_ForClabs.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ForClabs *ForClabsCallerSession) Version() (*big.Int, error) {
	return _ForClabs.Contract.Version(&_ForClabs.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.Approve(&_ForClabs.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.Approve(&_ForClabs.TransactOpts, to, tokenId)
}

// AuthorizeToken is a paid mutator transaction binding the contract method 0xdaac4e20.
//
// Solidity: function authorizeToken(address contractAddress, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactor) AuthorizeToken(opts *bind.TransactOpts, contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "authorizeToken", contractAddress, tokenId)
}

// AuthorizeToken is a paid mutator transaction binding the contract method 0xdaac4e20.
//
// Solidity: function authorizeToken(address contractAddress, uint256 tokenId) returns()
func (_ForClabs *ForClabsSession) AuthorizeToken(contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.AuthorizeToken(&_ForClabs.TransactOpts, contractAddress, tokenId)
}

// AuthorizeToken is a paid mutator transaction binding the contract method 0xdaac4e20.
//
// Solidity: function authorizeToken(address contractAddress, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactorSession) AuthorizeToken(contractAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.AuthorizeToken(&_ForClabs.TransactOpts, contractAddress, tokenId)
}

// AuthorizeTokens is a paid mutator transaction binding the contract method 0xc8ae0c74.
//
// Solidity: function authorizeTokens(address[] contractAddresses, uint256[] tokenIds) returns()
func (_ForClabs *ForClabsTransactor) AuthorizeTokens(opts *bind.TransactOpts, contractAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "authorizeTokens", contractAddresses, tokenIds)
}

// AuthorizeTokens is a paid mutator transaction binding the contract method 0xc8ae0c74.
//
// Solidity: function authorizeTokens(address[] contractAddresses, uint256[] tokenIds) returns()
func (_ForClabs *ForClabsSession) AuthorizeTokens(contractAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.AuthorizeTokens(&_ForClabs.TransactOpts, contractAddresses, tokenIds)
}

// AuthorizeTokens is a paid mutator transaction binding the contract method 0xc8ae0c74.
//
// Solidity: function authorizeTokens(address[] contractAddresses, uint256[] tokenIds) returns()
func (_ForClabs *ForClabsTransactorSession) AuthorizeTokens(contractAddresses []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.AuthorizeTokens(&_ForClabs.TransactOpts, contractAddresses, tokenIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signerAddress) returns()
func (_ForClabs *ForClabsTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, signerAddress common.Address) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "initialize", name, symbol, signerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signerAddress) returns()
func (_ForClabs *ForClabsSession) Initialize(name string, symbol string, signerAddress common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.Initialize(&_ForClabs.TransactOpts, name, symbol, signerAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address signerAddress) returns()
func (_ForClabs *ForClabsTransactorSession) Initialize(name string, symbol string, signerAddress common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.Initialize(&_ForClabs.TransactOpts, name, symbol, signerAddress)
}

// Mint is a paid mutator transaction binding the contract method 0xd3efdb4e.
//
// Solidity: function mint((string,string[],uint256,uint256,address,address,uint256) works, bytes sign) returns()
func (_ForClabs *ForClabsTransactor) Mint(opts *bind.TransactOpts, works IComicWorksComicWorks, sign []byte) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "mint", works, sign)
}

// Mint is a paid mutator transaction binding the contract method 0xd3efdb4e.
//
// Solidity: function mint((string,string[],uint256,uint256,address,address,uint256) works, bytes sign) returns()
func (_ForClabs *ForClabsSession) Mint(works IComicWorksComicWorks, sign []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.Mint(&_ForClabs.TransactOpts, works, sign)
}

// Mint is a paid mutator transaction binding the contract method 0xd3efdb4e.
//
// Solidity: function mint((string,string[],uint256,uint256,address,address,uint256) works, bytes sign) returns()
func (_ForClabs *ForClabsTransactorSession) Mint(works IComicWorksComicWorks, sign []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.Mint(&_ForClabs.TransactOpts, works, sign)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 tokenId) payable returns()
func (_ForClabs *ForClabsTransactor) Purchase(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "purchase", tokenId)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 tokenId) payable returns()
func (_ForClabs *ForClabsSession) Purchase(tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.Purchase(&_ForClabs.TransactOpts, tokenId)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 tokenId) payable returns()
func (_ForClabs *ForClabsTransactorSession) Purchase(tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.Purchase(&_ForClabs.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForClabs *ForClabsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForClabs *ForClabsSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForClabs.Contract.RenounceOwnership(&_ForClabs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForClabs *ForClabsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForClabs.Contract.RenounceOwnership(&_ForClabs.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SafeTransferFrom(&_ForClabs.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SafeTransferFrom(&_ForClabs.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ForClabs *ForClabsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ForClabs *ForClabsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.SafeTransferFrom0(&_ForClabs.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ForClabs *ForClabsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.SafeTransferFrom0(&_ForClabs.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ForClabs *ForClabsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ForClabs *ForClabsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ForClabs.Contract.SetApprovalForAll(&_ForClabs.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ForClabs *ForClabsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ForClabs.Contract.SetApprovalForAll(&_ForClabs.TransactOpts, operator, approved)
}

// SetLimit is a paid mutator transaction binding the contract method 0x207add91.
//
// Solidity: function setLimit(uint256 worksId, uint256 limit_) returns()
func (_ForClabs *ForClabsTransactor) SetLimit(opts *bind.TransactOpts, worksId *big.Int, limit_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "setLimit", worksId, limit_)
}

// SetLimit is a paid mutator transaction binding the contract method 0x207add91.
//
// Solidity: function setLimit(uint256 worksId, uint256 limit_) returns()
func (_ForClabs *ForClabsSession) SetLimit(worksId *big.Int, limit_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SetLimit(&_ForClabs.TransactOpts, worksId, limit_)
}

// SetLimit is a paid mutator transaction binding the contract method 0x207add91.
//
// Solidity: function setLimit(uint256 worksId, uint256 limit_) returns()
func (_ForClabs *ForClabsTransactorSession) SetLimit(worksId *big.Int, limit_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SetLimit(&_ForClabs.TransactOpts, worksId, limit_)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 worksId, uint256 price_) returns()
func (_ForClabs *ForClabsTransactor) SetPrice(opts *bind.TransactOpts, worksId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "setPrice", worksId, price_)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 worksId, uint256 price_) returns()
func (_ForClabs *ForClabsSession) SetPrice(worksId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SetPrice(&_ForClabs.TransactOpts, worksId, price_)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 worksId, uint256 price_) returns()
func (_ForClabs *ForClabsTransactorSession) SetPrice(worksId *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.SetPrice(&_ForClabs.TransactOpts, worksId, price_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.TransferFrom(&_ForClabs.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ForClabs *ForClabsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ForClabs.Contract.TransferFrom(&_ForClabs.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForClabs *ForClabsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForClabs *ForClabsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.TransferOwnership(&_ForClabs.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForClabs *ForClabsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.TransferOwnership(&_ForClabs.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ForClabs *ForClabsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ForClabs *ForClabsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.UpgradeTo(&_ForClabs.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ForClabs *ForClabsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ForClabs.Contract.UpgradeTo(&_ForClabs.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ForClabs *ForClabsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ForClabs *ForClabsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.UpgradeToAndCall(&_ForClabs.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ForClabs *ForClabsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ForClabs.Contract.UpgradeToAndCall(&_ForClabs.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ForClabs *ForClabsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForClabs.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ForClabs *ForClabsSession) Withdraw() (*types.Transaction, error) {
	return _ForClabs.Contract.Withdraw(&_ForClabs.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ForClabs *ForClabsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ForClabs.Contract.Withdraw(&_ForClabs.TransactOpts)
}

// ForClabsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ForClabs contract.
type ForClabsAdminChangedIterator struct {
	Event *ForClabsAdminChanged // Event containing the contract specifics and raw log

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
func (it *ForClabsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsAdminChanged)
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
		it.Event = new(ForClabsAdminChanged)
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
func (it *ForClabsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsAdminChanged represents a AdminChanged event raised by the ForClabs contract.
type ForClabsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ForClabs *ForClabsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ForClabsAdminChangedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ForClabsAdminChangedIterator{contract: _ForClabs.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ForClabs *ForClabsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ForClabsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsAdminChanged)
				if err := _ForClabs.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ForClabs *ForClabsFilterer) ParseAdminChanged(log types.Log) (*ForClabsAdminChanged, error) {
	event := new(ForClabsAdminChanged)
	if err := _ForClabs.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ForClabs contract.
type ForClabsApprovalIterator struct {
	Event *ForClabsApproval // Event containing the contract specifics and raw log

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
func (it *ForClabsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsApproval)
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
		it.Event = new(ForClabsApproval)
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
func (it *ForClabsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsApproval represents a Approval event raised by the ForClabs contract.
type ForClabsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ForClabsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsApprovalIterator{contract: _ForClabs.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ForClabsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsApproval)
				if err := _ForClabs.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) ParseApproval(log types.Log) (*ForClabsApproval, error) {
	event := new(ForClabsApproval)
	if err := _ForClabs.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ForClabs contract.
type ForClabsApprovalForAllIterator struct {
	Event *ForClabsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ForClabsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsApprovalForAll)
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
		it.Event = new(ForClabsApprovalForAll)
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
func (it *ForClabsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsApprovalForAll represents a ApprovalForAll event raised by the ForClabs contract.
type ForClabsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ForClabs *ForClabsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ForClabsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsApprovalForAllIterator{contract: _ForClabs.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ForClabs *ForClabsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ForClabsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsApprovalForAll)
				if err := _ForClabs.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ForClabs *ForClabsFilterer) ParseApprovalForAll(log types.Log) (*ForClabsApprovalForAll, error) {
	event := new(ForClabsApprovalForAll)
	if err := _ForClabs.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ForClabs contract.
type ForClabsBeaconUpgradedIterator struct {
	Event *ForClabsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ForClabsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsBeaconUpgraded)
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
		it.Event = new(ForClabsBeaconUpgraded)
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
func (it *ForClabsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsBeaconUpgraded represents a BeaconUpgraded event raised by the ForClabs contract.
type ForClabsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ForClabs *ForClabsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ForClabsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsBeaconUpgradedIterator{contract: _ForClabs.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ForClabs *ForClabsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ForClabsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsBeaconUpgraded)
				if err := _ForClabs.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ForClabs *ForClabsFilterer) ParseBeaconUpgraded(log types.Log) (*ForClabsBeaconUpgraded, error) {
	event := new(ForClabsBeaconUpgraded)
	if err := _ForClabs.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsComicWorksCreatedIterator is returned from FilterComicWorksCreated and is used to iterate over the raw logs and unpacked data for ComicWorksCreated events raised by the ForClabs contract.
type ForClabsComicWorksCreatedIterator struct {
	Event *ForClabsComicWorksCreated // Event containing the contract specifics and raw log

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
func (it *ForClabsComicWorksCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsComicWorksCreated)
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
		it.Event = new(ForClabsComicWorksCreated)
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
func (it *ForClabsComicWorksCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsComicWorksCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsComicWorksCreated represents a ComicWorksCreated event raised by the ForClabs contract.
type ForClabsComicWorksCreated struct {
	ComicWorks IComicWorksComicWorks
	WorksId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComicWorksCreated is a free log retrieval operation binding the contract event 0xf825e5cf5c4b9759e7e09466d7ea1edd09b5426ba962ba6a87813c648f3f7aab.
//
// Solidity: event ComicWorksCreated((string,string[],uint256,uint256,address,address,uint256) comicWorks, uint256 worksId)
func (_ForClabs *ForClabsFilterer) FilterComicWorksCreated(opts *bind.FilterOpts) (*ForClabsComicWorksCreatedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ComicWorksCreated")
	if err != nil {
		return nil, err
	}
	return &ForClabsComicWorksCreatedIterator{contract: _ForClabs.contract, event: "ComicWorksCreated", logs: logs, sub: sub}, nil
}

// WatchComicWorksCreated is a free log subscription operation binding the contract event 0xf825e5cf5c4b9759e7e09466d7ea1edd09b5426ba962ba6a87813c648f3f7aab.
//
// Solidity: event ComicWorksCreated((string,string[],uint256,uint256,address,address,uint256) comicWorks, uint256 worksId)
func (_ForClabs *ForClabsFilterer) WatchComicWorksCreated(opts *bind.WatchOpts, sink chan<- *ForClabsComicWorksCreated) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ComicWorksCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsComicWorksCreated)
				if err := _ForClabs.contract.UnpackLog(event, "ComicWorksCreated", log); err != nil {
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

// ParseComicWorksCreated is a log parse operation binding the contract event 0xf825e5cf5c4b9759e7e09466d7ea1edd09b5426ba962ba6a87813c648f3f7aab.
//
// Solidity: event ComicWorksCreated((string,string[],uint256,uint256,address,address,uint256) comicWorks, uint256 worksId)
func (_ForClabs *ForClabsFilterer) ParseComicWorksCreated(log types.Log) (*ForClabsComicWorksCreated, error) {
	event := new(ForClabsComicWorksCreated)
	if err := _ForClabs.contract.UnpackLog(event, "ComicWorksCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsComicWorksLimitChangedIterator is returned from FilterComicWorksLimitChanged and is used to iterate over the raw logs and unpacked data for ComicWorksLimitChanged events raised by the ForClabs contract.
type ForClabsComicWorksLimitChangedIterator struct {
	Event *ForClabsComicWorksLimitChanged // Event containing the contract specifics and raw log

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
func (it *ForClabsComicWorksLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsComicWorksLimitChanged)
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
		it.Event = new(ForClabsComicWorksLimitChanged)
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
func (it *ForClabsComicWorksLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsComicWorksLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsComicWorksLimitChanged represents a ComicWorksLimitChanged event raised by the ForClabs contract.
type ForClabsComicWorksLimitChanged struct {
	OldLimit *big.Int
	NewLimit *big.Int
	WorksId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterComicWorksLimitChanged is a free log retrieval operation binding the contract event 0x54e51332c541640bc4e02d8a2e4022acebe00cf7575de0d33e802a2d8ebb7866.
//
// Solidity: event ComicWorksLimitChanged(uint256 oldLimit, uint256 newLimit, uint256 worksId)
func (_ForClabs *ForClabsFilterer) FilterComicWorksLimitChanged(opts *bind.FilterOpts) (*ForClabsComicWorksLimitChangedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ComicWorksLimitChanged")
	if err != nil {
		return nil, err
	}
	return &ForClabsComicWorksLimitChangedIterator{contract: _ForClabs.contract, event: "ComicWorksLimitChanged", logs: logs, sub: sub}, nil
}

// WatchComicWorksLimitChanged is a free log subscription operation binding the contract event 0x54e51332c541640bc4e02d8a2e4022acebe00cf7575de0d33e802a2d8ebb7866.
//
// Solidity: event ComicWorksLimitChanged(uint256 oldLimit, uint256 newLimit, uint256 worksId)
func (_ForClabs *ForClabsFilterer) WatchComicWorksLimitChanged(opts *bind.WatchOpts, sink chan<- *ForClabsComicWorksLimitChanged) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ComicWorksLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsComicWorksLimitChanged)
				if err := _ForClabs.contract.UnpackLog(event, "ComicWorksLimitChanged", log); err != nil {
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

// ParseComicWorksLimitChanged is a log parse operation binding the contract event 0x54e51332c541640bc4e02d8a2e4022acebe00cf7575de0d33e802a2d8ebb7866.
//
// Solidity: event ComicWorksLimitChanged(uint256 oldLimit, uint256 newLimit, uint256 worksId)
func (_ForClabs *ForClabsFilterer) ParseComicWorksLimitChanged(log types.Log) (*ForClabsComicWorksLimitChanged, error) {
	event := new(ForClabsComicWorksLimitChanged)
	if err := _ForClabs.contract.UnpackLog(event, "ComicWorksLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsComicWorksPriceChangedIterator is returned from FilterComicWorksPriceChanged and is used to iterate over the raw logs and unpacked data for ComicWorksPriceChanged events raised by the ForClabs contract.
type ForClabsComicWorksPriceChangedIterator struct {
	Event *ForClabsComicWorksPriceChanged // Event containing the contract specifics and raw log

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
func (it *ForClabsComicWorksPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsComicWorksPriceChanged)
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
		it.Event = new(ForClabsComicWorksPriceChanged)
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
func (it *ForClabsComicWorksPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsComicWorksPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsComicWorksPriceChanged represents a ComicWorksPriceChanged event raised by the ForClabs contract.
type ForClabsComicWorksPriceChanged struct {
	OldPrice *big.Int
	NewPrice *big.Int
	WorksId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterComicWorksPriceChanged is a free log retrieval operation binding the contract event 0x28ed94e3c8c243f8ffd4e1c1e4b44b9c27c89996748ad5e101dec4377da688dc.
//
// Solidity: event ComicWorksPriceChanged(uint256 oldPrice, uint256 newPrice, uint256 worksId)
func (_ForClabs *ForClabsFilterer) FilterComicWorksPriceChanged(opts *bind.FilterOpts) (*ForClabsComicWorksPriceChangedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ComicWorksPriceChanged")
	if err != nil {
		return nil, err
	}
	return &ForClabsComicWorksPriceChangedIterator{contract: _ForClabs.contract, event: "ComicWorksPriceChanged", logs: logs, sub: sub}, nil
}

// WatchComicWorksPriceChanged is a free log subscription operation binding the contract event 0x28ed94e3c8c243f8ffd4e1c1e4b44b9c27c89996748ad5e101dec4377da688dc.
//
// Solidity: event ComicWorksPriceChanged(uint256 oldPrice, uint256 newPrice, uint256 worksId)
func (_ForClabs *ForClabsFilterer) WatchComicWorksPriceChanged(opts *bind.WatchOpts, sink chan<- *ForClabsComicWorksPriceChanged) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ComicWorksPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsComicWorksPriceChanged)
				if err := _ForClabs.contract.UnpackLog(event, "ComicWorksPriceChanged", log); err != nil {
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

// ParseComicWorksPriceChanged is a log parse operation binding the contract event 0x28ed94e3c8c243f8ffd4e1c1e4b44b9c27c89996748ad5e101dec4377da688dc.
//
// Solidity: event ComicWorksPriceChanged(uint256 oldPrice, uint256 newPrice, uint256 worksId)
func (_ForClabs *ForClabsFilterer) ParseComicWorksPriceChanged(log types.Log) (*ForClabsComicWorksPriceChanged, error) {
	event := new(ForClabsComicWorksPriceChanged)
	if err := _ForClabs.contract.UnpackLog(event, "ComicWorksPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsComicWorksSoldIterator is returned from FilterComicWorksSold and is used to iterate over the raw logs and unpacked data for ComicWorksSold events raised by the ForClabs contract.
type ForClabsComicWorksSoldIterator struct {
	Event *ForClabsComicWorksSold // Event containing the contract specifics and raw log

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
func (it *ForClabsComicWorksSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsComicWorksSold)
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
		it.Event = new(ForClabsComicWorksSold)
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
func (it *ForClabsComicWorksSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsComicWorksSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsComicWorksSold represents a ComicWorksSold event raised by the ForClabs contract.
type ForClabsComicWorksSold struct {
	ComicWorks IComicWorksComicWorks
	From       common.Address
	To         common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComicWorksSold is a free log retrieval operation binding the contract event 0xc4d54a84e9208251771379c631e6c9874eff62261676fb0a0f1423a6e129f5ba.
//
// Solidity: event ComicWorksSold((string,string[],uint256,uint256,address,address,uint256) comicWorks, address from, address to, uint256 tokenId)
func (_ForClabs *ForClabsFilterer) FilterComicWorksSold(opts *bind.FilterOpts) (*ForClabsComicWorksSoldIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ComicWorksSold")
	if err != nil {
		return nil, err
	}
	return &ForClabsComicWorksSoldIterator{contract: _ForClabs.contract, event: "ComicWorksSold", logs: logs, sub: sub}, nil
}

// WatchComicWorksSold is a free log subscription operation binding the contract event 0xc4d54a84e9208251771379c631e6c9874eff62261676fb0a0f1423a6e129f5ba.
//
// Solidity: event ComicWorksSold((string,string[],uint256,uint256,address,address,uint256) comicWorks, address from, address to, uint256 tokenId)
func (_ForClabs *ForClabsFilterer) WatchComicWorksSold(opts *bind.WatchOpts, sink chan<- *ForClabsComicWorksSold) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ComicWorksSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsComicWorksSold)
				if err := _ForClabs.contract.UnpackLog(event, "ComicWorksSold", log); err != nil {
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

// ParseComicWorksSold is a log parse operation binding the contract event 0xc4d54a84e9208251771379c631e6c9874eff62261676fb0a0f1423a6e129f5ba.
//
// Solidity: event ComicWorksSold((string,string[],uint256,uint256,address,address,uint256) comicWorks, address from, address to, uint256 tokenId)
func (_ForClabs *ForClabsFilterer) ParseComicWorksSold(log types.Log) (*ForClabsComicWorksSold, error) {
	event := new(ForClabsComicWorksSold)
	if err := _ForClabs.contract.UnpackLog(event, "ComicWorksSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsComicWorksTokenCreatedIterator is returned from FilterComicWorksTokenCreated and is used to iterate over the raw logs and unpacked data for ComicWorksTokenCreated events raised by the ForClabs contract.
type ForClabsComicWorksTokenCreatedIterator struct {
	Event *ForClabsComicWorksTokenCreated // Event containing the contract specifics and raw log

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
func (it *ForClabsComicWorksTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsComicWorksTokenCreated)
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
		it.Event = new(ForClabsComicWorksTokenCreated)
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
func (it *ForClabsComicWorksTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsComicWorksTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsComicWorksTokenCreated represents a ComicWorksTokenCreated event raised by the ForClabs contract.
type ForClabsComicWorksTokenCreated struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComicWorksTokenCreated is a free log retrieval operation binding the contract event 0x5ce25ca7986a96a12ff472fec6b04b94444f46c5c6df7de7d24bd259b6aa6298.
//
// Solidity: event ComicWorksTokenCreated(uint256 tokenId)
func (_ForClabs *ForClabsFilterer) FilterComicWorksTokenCreated(opts *bind.FilterOpts) (*ForClabsComicWorksTokenCreatedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ComicWorksTokenCreated")
	if err != nil {
		return nil, err
	}
	return &ForClabsComicWorksTokenCreatedIterator{contract: _ForClabs.contract, event: "ComicWorksTokenCreated", logs: logs, sub: sub}, nil
}

// WatchComicWorksTokenCreated is a free log subscription operation binding the contract event 0x5ce25ca7986a96a12ff472fec6b04b94444f46c5c6df7de7d24bd259b6aa6298.
//
// Solidity: event ComicWorksTokenCreated(uint256 tokenId)
func (_ForClabs *ForClabsFilterer) WatchComicWorksTokenCreated(opts *bind.WatchOpts, sink chan<- *ForClabsComicWorksTokenCreated) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ComicWorksTokenCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsComicWorksTokenCreated)
				if err := _ForClabs.contract.UnpackLog(event, "ComicWorksTokenCreated", log); err != nil {
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

// ParseComicWorksTokenCreated is a log parse operation binding the contract event 0x5ce25ca7986a96a12ff472fec6b04b94444f46c5c6df7de7d24bd259b6aa6298.
//
// Solidity: event ComicWorksTokenCreated(uint256 tokenId)
func (_ForClabs *ForClabsFilterer) ParseComicWorksTokenCreated(log types.Log) (*ForClabsComicWorksTokenCreated, error) {
	event := new(ForClabsComicWorksTokenCreated)
	if err := _ForClabs.contract.UnpackLog(event, "ComicWorksTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsContractAndTokenAuthorizedIterator is returned from FilterContractAndTokenAuthorized and is used to iterate over the raw logs and unpacked data for ContractAndTokenAuthorized events raised by the ForClabs contract.
type ForClabsContractAndTokenAuthorizedIterator struct {
	Event *ForClabsContractAndTokenAuthorized // Event containing the contract specifics and raw log

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
func (it *ForClabsContractAndTokenAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsContractAndTokenAuthorized)
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
		it.Event = new(ForClabsContractAndTokenAuthorized)
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
func (it *ForClabsContractAndTokenAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsContractAndTokenAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsContractAndTokenAuthorized represents a ContractAndTokenAuthorized event raised by the ForClabs contract.
type ForClabsContractAndTokenAuthorized struct {
	ContractAddress common.Address
	TokenId         *big.Int
	Holder          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractAndTokenAuthorized is a free log retrieval operation binding the contract event 0xc2dafeaf68847499134821649e130889b67ab4d7d86b6feed2435c354cae2280.
//
// Solidity: event ContractAndTokenAuthorized(address contractAddress, uint256 tokenId, address holder)
func (_ForClabs *ForClabsFilterer) FilterContractAndTokenAuthorized(opts *bind.FilterOpts) (*ForClabsContractAndTokenAuthorizedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "ContractAndTokenAuthorized")
	if err != nil {
		return nil, err
	}
	return &ForClabsContractAndTokenAuthorizedIterator{contract: _ForClabs.contract, event: "ContractAndTokenAuthorized", logs: logs, sub: sub}, nil
}

// WatchContractAndTokenAuthorized is a free log subscription operation binding the contract event 0xc2dafeaf68847499134821649e130889b67ab4d7d86b6feed2435c354cae2280.
//
// Solidity: event ContractAndTokenAuthorized(address contractAddress, uint256 tokenId, address holder)
func (_ForClabs *ForClabsFilterer) WatchContractAndTokenAuthorized(opts *bind.WatchOpts, sink chan<- *ForClabsContractAndTokenAuthorized) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "ContractAndTokenAuthorized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsContractAndTokenAuthorized)
				if err := _ForClabs.contract.UnpackLog(event, "ContractAndTokenAuthorized", log); err != nil {
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

// ParseContractAndTokenAuthorized is a log parse operation binding the contract event 0xc2dafeaf68847499134821649e130889b67ab4d7d86b6feed2435c354cae2280.
//
// Solidity: event ContractAndTokenAuthorized(address contractAddress, uint256 tokenId, address holder)
func (_ForClabs *ForClabsFilterer) ParseContractAndTokenAuthorized(log types.Log) (*ForClabsContractAndTokenAuthorized, error) {
	event := new(ForClabsContractAndTokenAuthorized)
	if err := _ForClabs.contract.UnpackLog(event, "ContractAndTokenAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ForClabs contract.
type ForClabsInitializedIterator struct {
	Event *ForClabsInitialized // Event containing the contract specifics and raw log

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
func (it *ForClabsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsInitialized)
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
		it.Event = new(ForClabsInitialized)
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
func (it *ForClabsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsInitialized represents a Initialized event raised by the ForClabs contract.
type ForClabsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ForClabs *ForClabsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ForClabsInitializedIterator, error) {

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ForClabsInitializedIterator{contract: _ForClabs.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ForClabs *ForClabsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ForClabsInitialized) (event.Subscription, error) {

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsInitialized)
				if err := _ForClabs.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ForClabs *ForClabsFilterer) ParseInitialized(log types.Log) (*ForClabsInitialized, error) {
	event := new(ForClabsInitialized)
	if err := _ForClabs.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ForClabs contract.
type ForClabsOwnershipTransferredIterator struct {
	Event *ForClabsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForClabsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsOwnershipTransferred)
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
		it.Event = new(ForClabsOwnershipTransferred)
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
func (it *ForClabsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsOwnershipTransferred represents a OwnershipTransferred event raised by the ForClabs contract.
type ForClabsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForClabs *ForClabsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ForClabsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsOwnershipTransferredIterator{contract: _ForClabs.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForClabs *ForClabsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForClabsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsOwnershipTransferred)
				if err := _ForClabs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForClabs *ForClabsFilterer) ParseOwnershipTransferred(log types.Log) (*ForClabsOwnershipTransferred, error) {
	event := new(ForClabsOwnershipTransferred)
	if err := _ForClabs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ForClabs contract.
type ForClabsTransferIterator struct {
	Event *ForClabsTransfer // Event containing the contract specifics and raw log

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
func (it *ForClabsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsTransfer)
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
		it.Event = new(ForClabsTransfer)
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
func (it *ForClabsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsTransfer represents a Transfer event raised by the ForClabs contract.
type ForClabsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ForClabsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsTransferIterator{contract: _ForClabs.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ForClabsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsTransfer)
				if err := _ForClabs.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ForClabs *ForClabsFilterer) ParseTransfer(log types.Log) (*ForClabsTransfer, error) {
	event := new(ForClabsTransfer)
	if err := _ForClabs.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForClabsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ForClabs contract.
type ForClabsUpgradedIterator struct {
	Event *ForClabsUpgraded // Event containing the contract specifics and raw log

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
func (it *ForClabsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForClabsUpgraded)
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
		it.Event = new(ForClabsUpgraded)
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
func (it *ForClabsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForClabsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForClabsUpgraded represents a Upgraded event raised by the ForClabs contract.
type ForClabsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ForClabs *ForClabsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ForClabsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ForClabs.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ForClabsUpgradedIterator{contract: _ForClabs.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ForClabs *ForClabsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ForClabsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ForClabs.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForClabsUpgraded)
				if err := _ForClabs.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ForClabs *ForClabsFilterer) ParseUpgraded(log types.Log) (*ForClabsUpgraded, error) {
	event := new(ForClabsUpgraded)
	if err := _ForClabs.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
