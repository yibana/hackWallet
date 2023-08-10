// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlurExchangeV2Interface

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

// Cancel is an auto generated low-level Go binding around an user-defined struct.
type Cancel struct {
	Hash   [32]byte
	Index  *big.Int
	Amount *big.Int
}

// Exchange is an auto generated low-level Go binding around an user-defined struct.
type Exchange struct {
	Index   *big.Int
	Proof   [][32]byte
	Listing Listing
	Taker   Taker
}

// FeeRate is an auto generated low-level Go binding around an user-defined struct.
type FeeRate struct {
	Recipient common.Address
	Rate      uint16
}

// Fees is an auto generated low-level Go binding around an user-defined struct.
type Fees struct {
	ProtocolFee FeeRate
	TakerFee    FeeRate
}

// Listing is an auto generated low-level Go binding around an user-defined struct.
type Listing struct {
	Index   *big.Int
	TokenId *big.Int
	Amount  *big.Int
	Price   *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader           common.Address
	Collection       common.Address
	ListingsRoot     [32]byte
	NumberOfListings *big.Int
	ExpirationTime   *big.Int
	AssetType        uint8
	MakerFee         FeeRate
	Salt             *big.Int
}

// TakeAsk is an auto generated low-level Go binding around an user-defined struct.
type TakeAsk struct {
	Orders         []Order
	Exchanges      []Exchange
	TakerFee       FeeRate
	Signatures     []byte
	TokenRecipient common.Address
}

// TakeAskSingle is an auto generated low-level Go binding around an user-defined struct.
type TakeAskSingle struct {
	Order          Order
	Exchange       Exchange
	TakerFee       FeeRate
	Signature      []byte
	TokenRecipient common.Address
}

// TakeBid is an auto generated low-level Go binding around an user-defined struct.
type TakeBid struct {
	Orders     []Order
	Exchanges  []Exchange
	TakerFee   FeeRate
	Signatures []byte
}

// TakeBidSingle is an auto generated low-level Go binding around an user-defined struct.
type TakeBidSingle struct {
	Order     Order
	Exchange  Exchange
	TakerFee  FeeRate
	Signature []byte
}

// Taker is an auto generated low-level Go binding around an user-defined struct.
type Taker struct {
	TokenId *big.Int
	Amount  *big.Int
}

// Transfer is an auto generated low-level Go binding around an user-defined struct.
type Transfer struct {
	Trader     common.Address
	Id         *big.Int
	Amount     *big.Int
	Collection common.Address
	AssetType  uint8
}

// BlurExchangeV2InterfaceMetaData contains all meta data concerning the BlurExchangeV2Interface contract.
var BlurExchangeV2InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredOracleSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOracleSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolWithdrawFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedOracle\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CancelTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structTransfer\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structFees\",\"name\":\"fees\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenIdListingIndexTrader\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionPriceSide\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerFeeRecipientRate\",\"type\":\"uint256\"}],\"name\":\"Execution721MakerFeePacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenIdListingIndexTrader\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionPriceSide\",\"type\":\"uint256\"}],\"name\":\"Execution721Packed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenIdListingIndexTrader\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionPriceSide\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerFeeRecipientRate\",\"type\":\"uint256\"}],\"name\":\"Execution721TakerFeePacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockRange\",\"type\":\"uint256\"}],\"name\":\"NewBlockRange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"NewProtocolFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"SetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"amountTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCancel[]\",\"name\":\"cancels\",\"type\":\"tuple[]\"}],\"name\":\"cancelTrades\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"hashListing\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange[]\",\"name\":\"exchanges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAsk\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"hashTakeAsk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange\",\"name\":\"exchange\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAskSingle\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"hashTakeAskSingle\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange[]\",\"name\":\"exchanges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"internalType\":\"structTakeBid\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"hashTakeBid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange\",\"name\":\"exchange\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTakeBidSingle\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"hashTakeBidSingle\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange[]\",\"name\":\"exchanges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAsk\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"}],\"name\":\"takeAsk\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange[]\",\"name\":\"exchanges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAsk\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"takeAskPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange\",\"name\":\"exchange\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAskSingle\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"}],\"name\":\"takeAskSingle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange\",\"name\":\"exchange\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tokenRecipient\",\"type\":\"address\"}],\"internalType\":\"structTakeAskSingle\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"takeAskSinglePool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange[]\",\"name\":\"exchanges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"internalType\":\"structTakeBid\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"}],\"name\":\"takeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"listingsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfListings\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"enumAssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"makerFee\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structListing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTaker\",\"name\":\"taker\",\"type\":\"tuple\"}],\"internalType\":\"structExchange\",\"name\":\"exchange\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structFeeRate\",\"name\":\"takerFee\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTakeBidSingle\",\"name\":\"inputs\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"oracleSignature\",\"type\":\"bytes\"}],\"name\":\"takeBidSingle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyDomain\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BlurExchangeV2InterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use BlurExchangeV2InterfaceMetaData.ABI instead.
var BlurExchangeV2InterfaceABI = BlurExchangeV2InterfaceMetaData.ABI

// BlurExchangeV2Interface is an auto generated Go binding around an Ethereum contract.
type BlurExchangeV2Interface struct {
	BlurExchangeV2InterfaceCaller     // Read-only binding to the contract
	BlurExchangeV2InterfaceTransactor // Write-only binding to the contract
	BlurExchangeV2InterfaceFilterer   // Log filterer for contract events
}

// BlurExchangeV2InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlurExchangeV2InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeV2InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlurExchangeV2InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeV2InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlurExchangeV2InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlurExchangeV2InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlurExchangeV2InterfaceSession struct {
	Contract     *BlurExchangeV2Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlurExchangeV2InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlurExchangeV2InterfaceCallerSession struct {
	Contract *BlurExchangeV2InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BlurExchangeV2InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlurExchangeV2InterfaceTransactorSession struct {
	Contract     *BlurExchangeV2InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BlurExchangeV2InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlurExchangeV2InterfaceRaw struct {
	Contract *BlurExchangeV2Interface // Generic contract binding to access the raw methods on
}

// BlurExchangeV2InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlurExchangeV2InterfaceCallerRaw struct {
	Contract *BlurExchangeV2InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// BlurExchangeV2InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlurExchangeV2InterfaceTransactorRaw struct {
	Contract *BlurExchangeV2InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlurExchangeV2Interface creates a new instance of BlurExchangeV2Interface, bound to a specific deployed contract.
func NewBlurExchangeV2Interface(address common.Address, backend bind.ContractBackend) (*BlurExchangeV2Interface, error) {
	contract, err := bindBlurExchangeV2Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2Interface{BlurExchangeV2InterfaceCaller: BlurExchangeV2InterfaceCaller{contract: contract}, BlurExchangeV2InterfaceTransactor: BlurExchangeV2InterfaceTransactor{contract: contract}, BlurExchangeV2InterfaceFilterer: BlurExchangeV2InterfaceFilterer{contract: contract}}, nil
}

// NewBlurExchangeV2InterfaceCaller creates a new read-only instance of BlurExchangeV2Interface, bound to a specific deployed contract.
func NewBlurExchangeV2InterfaceCaller(address common.Address, caller bind.ContractCaller) (*BlurExchangeV2InterfaceCaller, error) {
	contract, err := bindBlurExchangeV2Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceCaller{contract: contract}, nil
}

// NewBlurExchangeV2InterfaceTransactor creates a new write-only instance of BlurExchangeV2Interface, bound to a specific deployed contract.
func NewBlurExchangeV2InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BlurExchangeV2InterfaceTransactor, error) {
	contract, err := bindBlurExchangeV2Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceTransactor{contract: contract}, nil
}

// NewBlurExchangeV2InterfaceFilterer creates a new log filterer instance of BlurExchangeV2Interface, bound to a specific deployed contract.
func NewBlurExchangeV2InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BlurExchangeV2InterfaceFilterer, error) {
	contract, err := bindBlurExchangeV2Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceFilterer{contract: contract}, nil
}

// bindBlurExchangeV2Interface binds a generic wrapper to an already deployed contract.
func bindBlurExchangeV2Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlurExchangeV2InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchangeV2Interface.Contract.BlurExchangeV2InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.BlurExchangeV2InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.BlurExchangeV2InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlurExchangeV2Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.contract.Transact(opts, method, params...)
}

// AmountTaken is a free data retrieval call binding the contract method 0x3c2ab6bf.
//
// Solidity: function amountTaken(address , bytes32 , uint256 ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) AmountTaken(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "amountTaken", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountTaken is a free data retrieval call binding the contract method 0x3c2ab6bf.
//
// Solidity: function amountTaken(address , bytes32 , uint256 ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) AmountTaken(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.AmountTaken(&_BlurExchangeV2Interface.CallOpts, arg0, arg1, arg2)
}

// AmountTaken is a free data retrieval call binding the contract method 0x3c2ab6bf.
//
// Solidity: function amountTaken(address , bytes32 , uint256 ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) AmountTaken(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.AmountTaken(&_BlurExchangeV2Interface.CallOpts, arg0, arg1, arg2)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "blockRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) BlockRange() (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.BlockRange(&_BlurExchangeV2Interface.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) BlockRange() (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.BlockRange(&_BlurExchangeV2Interface.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Governor() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.Governor(&_BlurExchangeV2Interface.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) Governor() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.Governor(&_BlurExchangeV2Interface.CallOpts)
}

// HashListing is a free data retrieval call binding the contract method 0x87cc694d.
//
// Solidity: function hashListing((uint256,uint256,uint256,uint256) listing) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashListing(opts *bind.CallOpts, listing Listing) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashListing", listing)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashListing is a free data retrieval call binding the contract method 0x87cc694d.
//
// Solidity: function hashListing((uint256,uint256,uint256,uint256) listing) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashListing(listing Listing) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashListing(&_BlurExchangeV2Interface.CallOpts, listing)
}

// HashListing is a free data retrieval call binding the contract method 0x87cc694d.
//
// Solidity: function hashListing((uint256,uint256,uint256,uint256) listing) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashListing(listing Listing) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashListing(&_BlurExchangeV2Interface.CallOpts, listing)
}

// HashOrder is a free data retrieval call binding the contract method 0x51114ffa.
//
// Solidity: function hashOrder((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256) order, uint8 orderType) view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashOrder(opts *bind.CallOpts, order Order, orderType uint8) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashOrder", order, orderType)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x51114ffa.
//
// Solidity: function hashOrder((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256) order, uint8 orderType) view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashOrder(order Order, orderType uint8) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashOrder(&_BlurExchangeV2Interface.CallOpts, order, orderType)
}

// HashOrder is a free data retrieval call binding the contract method 0x51114ffa.
//
// Solidity: function hashOrder((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256) order, uint8 orderType) view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashOrder(order Order, orderType uint8) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashOrder(&_BlurExchangeV2Interface.CallOpts, order, orderType)
}

// HashTakeAsk is a free data retrieval call binding the contract method 0xcf6b0f52.
//
// Solidity: function hashTakeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashTakeAsk(opts *bind.CallOpts, inputs TakeAsk, _caller common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashTakeAsk", inputs, _caller)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTakeAsk is a free data retrieval call binding the contract method 0xcf6b0f52.
//
// Solidity: function hashTakeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashTakeAsk(inputs TakeAsk, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeAsk(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeAsk is a free data retrieval call binding the contract method 0xcf6b0f52.
//
// Solidity: function hashTakeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashTakeAsk(inputs TakeAsk, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeAsk(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeAskSingle is a free data retrieval call binding the contract method 0x579077b8.
//
// Solidity: function hashTakeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashTakeAskSingle(opts *bind.CallOpts, inputs TakeAskSingle, _caller common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashTakeAskSingle", inputs, _caller)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTakeAskSingle is a free data retrieval call binding the contract method 0x579077b8.
//
// Solidity: function hashTakeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashTakeAskSingle(inputs TakeAskSingle, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeAskSingle(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeAskSingle is a free data retrieval call binding the contract method 0x579077b8.
//
// Solidity: function hashTakeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashTakeAskSingle(inputs TakeAskSingle, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeAskSingle(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeBid is a free data retrieval call binding the contract method 0xabe3bb66.
//
// Solidity: function hashTakeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashTakeBid(opts *bind.CallOpts, inputs TakeBid, _caller common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashTakeBid", inputs, _caller)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTakeBid is a free data retrieval call binding the contract method 0xabe3bb66.
//
// Solidity: function hashTakeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashTakeBid(inputs TakeBid, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeBid(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeBid is a free data retrieval call binding the contract method 0xabe3bb66.
//
// Solidity: function hashTakeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashTakeBid(inputs TakeBid, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeBid(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeBidSingle is a free data retrieval call binding the contract method 0x91bea840.
//
// Solidity: function hashTakeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) HashTakeBidSingle(opts *bind.CallOpts, inputs TakeBidSingle, _caller common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "hashTakeBidSingle", inputs, _caller)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTakeBidSingle is a free data retrieval call binding the contract method 0x91bea840.
//
// Solidity: function hashTakeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) HashTakeBidSingle(inputs TakeBidSingle, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeBidSingle(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// HashTakeBidSingle is a free data retrieval call binding the contract method 0x91bea840.
//
// Solidity: function hashTakeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, address _caller) pure returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) HashTakeBidSingle(inputs TakeBidSingle, _caller common.Address) ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.HashTakeBidSingle(&_BlurExchangeV2Interface.CallOpts, inputs, _caller)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) Information(opts *bind.CallOpts) (struct {
	Version         string
	DomainSeparator [32]byte
}, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "information")

	outstruct := new(struct {
		Version         string
		DomainSeparator [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainSeparator = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Information() (struct {
	Version         string
	DomainSeparator [32]byte
}, error) {
	return _BlurExchangeV2Interface.Contract.Information(&_BlurExchangeV2Interface.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) Information() (struct {
	Version         string
	DomainSeparator [32]byte
}, error) {
	return _BlurExchangeV2Interface.Contract.Information(&_BlurExchangeV2Interface.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.Nonces(&_BlurExchangeV2Interface.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.Nonces(&_BlurExchangeV2Interface.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Oracles(arg0 common.Address) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.Oracles(&_BlurExchangeV2Interface.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint256)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) Oracles(arg0 common.Address) (*big.Int, error) {
	return _BlurExchangeV2Interface.Contract.Oracles(&_BlurExchangeV2Interface.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Owner() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.Owner(&_BlurExchangeV2Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) Owner() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.Owner(&_BlurExchangeV2Interface.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) PendingOwner() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.PendingOwner(&_BlurExchangeV2Interface.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) PendingOwner() (common.Address, error) {
	return _BlurExchangeV2Interface.Contract.PendingOwner(&_BlurExchangeV2Interface.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(address recipient, uint16 rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) ProtocolFee(opts *bind.CallOpts) (struct {
	Recipient common.Address
	Rate      uint16
}, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "protocolFee")

	outstruct := new(struct {
		Recipient common.Address
		Rate      uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Rate = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(address recipient, uint16 rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) ProtocolFee() (struct {
	Recipient common.Address
	Rate      uint16
}, error) {
	return _BlurExchangeV2Interface.Contract.ProtocolFee(&_BlurExchangeV2Interface.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(address recipient, uint16 rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) ProtocolFee() (struct {
	Recipient common.Address
	Rate      uint16
}, error) {
	return _BlurExchangeV2Interface.Contract.ProtocolFee(&_BlurExchangeV2Interface.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) ProxiableUUID() ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.ProxiableUUID(&_BlurExchangeV2Interface.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BlurExchangeV2Interface.Contract.ProxiableUUID(&_BlurExchangeV2Interface.CallOpts)
}

// VerifyDomain is a free data retrieval call binding the contract method 0x708ef9aa.
//
// Solidity: function verifyDomain() view returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCaller) VerifyDomain(opts *bind.CallOpts) error {
	var out []interface{}
	err := _BlurExchangeV2Interface.contract.Call(opts, &out, "verifyDomain")

	if err != nil {
		return err
	}

	return err

}

// VerifyDomain is a free data retrieval call binding the contract method 0x708ef9aa.
//
// Solidity: function verifyDomain() view returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) VerifyDomain() error {
	return _BlurExchangeV2Interface.Contract.VerifyDomain(&_BlurExchangeV2Interface.CallOpts)
}

// VerifyDomain is a free data retrieval call binding the contract method 0x708ef9aa.
//
// Solidity: function verifyDomain() view returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceCallerSession) VerifyDomain() error {
	return _BlurExchangeV2Interface.Contract.VerifyDomain(&_BlurExchangeV2Interface.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.AcceptOwnership(&_BlurExchangeV2Interface.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.AcceptOwnership(&_BlurExchangeV2Interface.TransactOpts)
}

// CancelTrades is a paid mutator transaction binding the contract method 0x3f8fc233.
//
// Solidity: function cancelTrades((bytes32,uint256,uint256)[] cancels) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) CancelTrades(opts *bind.TransactOpts, cancels []Cancel) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "cancelTrades", cancels)
}

// CancelTrades is a paid mutator transaction binding the contract method 0x3f8fc233.
//
// Solidity: function cancelTrades((bytes32,uint256,uint256)[] cancels) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) CancelTrades(cancels []Cancel) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.CancelTrades(&_BlurExchangeV2Interface.TransactOpts, cancels)
}

// CancelTrades is a paid mutator transaction binding the contract method 0x3f8fc233.
//
// Solidity: function cancelTrades((bytes32,uint256,uint256)[] cancels) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) CancelTrades(cancels []Cancel) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.CancelTrades(&_BlurExchangeV2Interface.TransactOpts, cancels)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) IncrementNonce() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.IncrementNonce(&_BlurExchangeV2Interface.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.IncrementNonce(&_BlurExchangeV2Interface.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Initialize() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.Initialize(&_BlurExchangeV2Interface.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) Initialize() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.Initialize(&_BlurExchangeV2Interface.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.RenounceOwnership(&_BlurExchangeV2Interface.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.RenounceOwnership(&_BlurExchangeV2Interface.TransactOpts)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetBlockRange(&_BlurExchangeV2Interface.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetBlockRange(&_BlurExchangeV2Interface.TransactOpts, _blockRange)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetGovernor(&_BlurExchangeV2Interface.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetGovernor(&_BlurExchangeV2Interface.TransactOpts, _governor)
}

// SetOracle is a paid mutator transaction binding the contract method 0x65360843.
//
// Solidity: function setOracle(address oracle, bool approved) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address, approved bool) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "setOracle", oracle, approved)
}

// SetOracle is a paid mutator transaction binding the contract method 0x65360843.
//
// Solidity: function setOracle(address oracle, bool approved) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) SetOracle(oracle common.Address, approved bool) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetOracle(&_BlurExchangeV2Interface.TransactOpts, oracle, approved)
}

// SetOracle is a paid mutator transaction binding the contract method 0x65360843.
//
// Solidity: function setOracle(address oracle, bool approved) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) SetOracle(oracle common.Address, approved bool) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetOracle(&_BlurExchangeV2Interface.TransactOpts, oracle, approved)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x3a16b768.
//
// Solidity: function setProtocolFee(address recipient, uint16 rate) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) SetProtocolFee(opts *bind.TransactOpts, recipient common.Address, rate uint16) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "setProtocolFee", recipient, rate)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x3a16b768.
//
// Solidity: function setProtocolFee(address recipient, uint16 rate) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) SetProtocolFee(recipient common.Address, rate uint16) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetProtocolFee(&_BlurExchangeV2Interface.TransactOpts, recipient, rate)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x3a16b768.
//
// Solidity: function setProtocolFee(address recipient, uint16 rate) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) SetProtocolFee(recipient common.Address, rate uint16) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.SetProtocolFee(&_BlurExchangeV2Interface.TransactOpts, recipient, rate)
}

// TakeAsk is a paid mutator transaction binding the contract method 0x3925c3c3.
//
// Solidity: function takeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeAsk(opts *bind.TransactOpts, inputs TakeAsk, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeAsk", inputs, oracleSignature)
}

// TakeAsk is a paid mutator transaction binding the contract method 0x3925c3c3.
//
// Solidity: function takeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeAsk(inputs TakeAsk, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAsk(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeAsk is a paid mutator transaction binding the contract method 0x3925c3c3.
//
// Solidity: function takeAsk(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeAsk(inputs TakeAsk, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAsk(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeAskPool is a paid mutator transaction binding the contract method 0x133ba9a6.
//
// Solidity: function takeAskPool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeAskPool(opts *bind.TransactOpts, inputs TakeAsk, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeAskPool", inputs, oracleSignature, amountToWithdraw)
}

// TakeAskPool is a paid mutator transaction binding the contract method 0x133ba9a6.
//
// Solidity: function takeAskPool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeAskPool(inputs TakeAsk, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskPool(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature, amountToWithdraw)
}

// TakeAskPool is a paid mutator transaction binding the contract method 0x133ba9a6.
//
// Solidity: function takeAskPool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeAskPool(inputs TakeAsk, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskPool(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature, amountToWithdraw)
}

// TakeAskSingle is a paid mutator transaction binding the contract method 0x70bce2d6.
//
// Solidity: function takeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeAskSingle(opts *bind.TransactOpts, inputs TakeAskSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeAskSingle", inputs, oracleSignature)
}

// TakeAskSingle is a paid mutator transaction binding the contract method 0x70bce2d6.
//
// Solidity: function takeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeAskSingle(inputs TakeAskSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskSingle(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeAskSingle is a paid mutator transaction binding the contract method 0x70bce2d6.
//
// Solidity: function takeAskSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeAskSingle(inputs TakeAskSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskSingle(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeAskSinglePool is a paid mutator transaction binding the contract method 0x336d8206.
//
// Solidity: function takeAskSinglePool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeAskSinglePool(opts *bind.TransactOpts, inputs TakeAskSingle, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeAskSinglePool", inputs, oracleSignature, amountToWithdraw)
}

// TakeAskSinglePool is a paid mutator transaction binding the contract method 0x336d8206.
//
// Solidity: function takeAskSinglePool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeAskSinglePool(inputs TakeAskSingle, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskSinglePool(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature, amountToWithdraw)
}

// TakeAskSinglePool is a paid mutator transaction binding the contract method 0x336d8206.
//
// Solidity: function takeAskSinglePool(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes,address) inputs, bytes oracleSignature, uint256 amountToWithdraw) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeAskSinglePool(inputs TakeAskSingle, oracleSignature []byte, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeAskSinglePool(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature, amountToWithdraw)
}

// TakeBid is a paid mutator transaction binding the contract method 0x7034d120.
//
// Solidity: function takeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeBid(opts *bind.TransactOpts, inputs TakeBid, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeBid", inputs, oracleSignature)
}

// TakeBid is a paid mutator transaction binding the contract method 0x7034d120.
//
// Solidity: function takeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeBid(inputs TakeBid, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeBid(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeBid is a paid mutator transaction binding the contract method 0x7034d120.
//
// Solidity: function takeBid(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256)[],(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256))[],(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeBid(inputs TakeBid, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeBid(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeBidSingle is a paid mutator transaction binding the contract method 0xda815cb5.
//
// Solidity: function takeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TakeBidSingle(opts *bind.TransactOpts, inputs TakeBidSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "takeBidSingle", inputs, oracleSignature)
}

// TakeBidSingle is a paid mutator transaction binding the contract method 0xda815cb5.
//
// Solidity: function takeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TakeBidSingle(inputs TakeBidSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeBidSingle(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TakeBidSingle is a paid mutator transaction binding the contract method 0xda815cb5.
//
// Solidity: function takeBidSingle(((address,address,bytes32,uint256,uint256,uint8,(address,uint16),uint256),(uint256,bytes32[],(uint256,uint256,uint256,uint256),(uint256,uint256)),(address,uint16),bytes) inputs, bytes oracleSignature) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TakeBidSingle(inputs TakeBidSingle, oracleSignature []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TakeBidSingle(&_BlurExchangeV2Interface.TransactOpts, inputs, oracleSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TransferOwnership(&_BlurExchangeV2Interface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.TransferOwnership(&_BlurExchangeV2Interface.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.UpgradeTo(&_BlurExchangeV2Interface.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.UpgradeTo(&_BlurExchangeV2Interface.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.UpgradeToAndCall(&_BlurExchangeV2Interface.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.UpgradeToAndCall(&_BlurExchangeV2Interface.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlurExchangeV2Interface.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceSession) Receive() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.Receive(&_BlurExchangeV2Interface.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceTransactorSession) Receive() (*types.Transaction, error) {
	return _BlurExchangeV2Interface.Contract.Receive(&_BlurExchangeV2Interface.TransactOpts)
}

// BlurExchangeV2InterfaceAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceAdminChangedIterator struct {
	Event *BlurExchangeV2InterfaceAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceAdminChanged)
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
		it.Event = new(BlurExchangeV2InterfaceAdminChanged)
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
func (it *BlurExchangeV2InterfaceAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceAdminChanged represents a AdminChanged event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceAdminChangedIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceAdminChangedIterator{contract: _BlurExchangeV2Interface.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceAdminChanged)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseAdminChanged(log types.Log) (*BlurExchangeV2InterfaceAdminChanged, error) {
	event := new(BlurExchangeV2InterfaceAdminChanged)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceBeaconUpgradedIterator struct {
	Event *BlurExchangeV2InterfaceBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceBeaconUpgraded)
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
		it.Event = new(BlurExchangeV2InterfaceBeaconUpgraded)
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
func (it *BlurExchangeV2InterfaceBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceBeaconUpgraded represents a BeaconUpgraded event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BlurExchangeV2InterfaceBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceBeaconUpgradedIterator{contract: _BlurExchangeV2Interface.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceBeaconUpgraded)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseBeaconUpgraded(log types.Log) (*BlurExchangeV2InterfaceBeaconUpgraded, error) {
	event := new(BlurExchangeV2InterfaceBeaconUpgraded)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceCancelTradeIterator is returned from FilterCancelTrade and is used to iterate over the raw logs and unpacked data for CancelTrade events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceCancelTradeIterator struct {
	Event *BlurExchangeV2InterfaceCancelTrade // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceCancelTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceCancelTrade)
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
		it.Event = new(BlurExchangeV2InterfaceCancelTrade)
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
func (it *BlurExchangeV2InterfaceCancelTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceCancelTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceCancelTrade represents a CancelTrade event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceCancelTrade struct {
	User   common.Address
	Hash   [32]byte
	Index  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelTrade is a free log retrieval operation binding the contract event 0xf4092a7c54e135dc5f273d6675327b7b7838392537d2f7b63f7acbec8c7cd296.
//
// Solidity: event CancelTrade(address indexed user, bytes32 hash, uint256 index, uint256 amount)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterCancelTrade(opts *bind.FilterOpts, user []common.Address) (*BlurExchangeV2InterfaceCancelTradeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "CancelTrade", userRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceCancelTradeIterator{contract: _BlurExchangeV2Interface.contract, event: "CancelTrade", logs: logs, sub: sub}, nil
}

// WatchCancelTrade is a free log subscription operation binding the contract event 0xf4092a7c54e135dc5f273d6675327b7b7838392537d2f7b63f7acbec8c7cd296.
//
// Solidity: event CancelTrade(address indexed user, bytes32 hash, uint256 index, uint256 amount)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchCancelTrade(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceCancelTrade, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "CancelTrade", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceCancelTrade)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "CancelTrade", log); err != nil {
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

// ParseCancelTrade is a log parse operation binding the contract event 0xf4092a7c54e135dc5f273d6675327b7b7838392537d2f7b63f7acbec8c7cd296.
//
// Solidity: event CancelTrade(address indexed user, bytes32 hash, uint256 index, uint256 amount)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseCancelTrade(log types.Log) (*BlurExchangeV2InterfaceCancelTrade, error) {
	event := new(BlurExchangeV2InterfaceCancelTrade)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "CancelTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecutionIterator struct {
	Event *BlurExchangeV2InterfaceExecution // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceExecution)
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
		it.Event = new(BlurExchangeV2InterfaceExecution)
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
func (it *BlurExchangeV2InterfaceExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceExecution represents a Execution event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution struct {
	Transfer     Transfer
	OrderHash    [32]byte
	ListingIndex *big.Int
	Price        *big.Int
	MakerFee     FeeRate
	Fees         Fees
	OrderType    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0xf2f66294df6fae7ac681cbe2f6d91c6904485929679dce263e8f6539b7d5c559.
//
// Solidity: event Execution((address,uint256,uint256,address,uint8) transfer, bytes32 orderHash, uint256 listingIndex, uint256 price, (address,uint16) makerFee, ((address,uint16),(address,uint16)) fees, uint8 orderType)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterExecution(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceExecutionIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Execution")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceExecutionIterator{contract: _BlurExchangeV2Interface.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0xf2f66294df6fae7ac681cbe2f6d91c6904485929679dce263e8f6539b7d5c559.
//
// Solidity: event Execution((address,uint256,uint256,address,uint8) transfer, bytes32 orderHash, uint256 listingIndex, uint256 price, (address,uint16) makerFee, ((address,uint16),(address,uint16)) fees, uint8 orderType)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceExecution) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Execution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceExecution)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution", log); err != nil {
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

// ParseExecution is a log parse operation binding the contract event 0xf2f66294df6fae7ac681cbe2f6d91c6904485929679dce263e8f6539b7d5c559.
//
// Solidity: event Execution((address,uint256,uint256,address,uint8) transfer, bytes32 orderHash, uint256 listingIndex, uint256 price, (address,uint16) makerFee, ((address,uint16),(address,uint16)) fees, uint8 orderType)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseExecution(log types.Log) (*BlurExchangeV2InterfaceExecution, error) {
	event := new(BlurExchangeV2InterfaceExecution)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceExecution721MakerFeePackedIterator is returned from FilterExecution721MakerFeePacked and is used to iterate over the raw logs and unpacked data for Execution721MakerFeePacked events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721MakerFeePackedIterator struct {
	Event *BlurExchangeV2InterfaceExecution721MakerFeePacked // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceExecution721MakerFeePackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceExecution721MakerFeePacked)
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
		it.Event = new(BlurExchangeV2InterfaceExecution721MakerFeePacked)
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
func (it *BlurExchangeV2InterfaceExecution721MakerFeePackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceExecution721MakerFeePackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceExecution721MakerFeePacked represents a Execution721MakerFeePacked event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721MakerFeePacked struct {
	OrderHash                 [32]byte
	TokenIdListingIndexTrader *big.Int
	CollectionPriceSide       *big.Int
	MakerFeeRecipientRate     *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterExecution721MakerFeePacked is a free log retrieval operation binding the contract event 0x7dc5c0699ac8dd5250cbe368a2fc3b4a2daadb120ad07f6cccea29f83482686e.
//
// Solidity: event Execution721MakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 makerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterExecution721MakerFeePacked(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceExecution721MakerFeePackedIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Execution721MakerFeePacked")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceExecution721MakerFeePackedIterator{contract: _BlurExchangeV2Interface.contract, event: "Execution721MakerFeePacked", logs: logs, sub: sub}, nil
}

// WatchExecution721MakerFeePacked is a free log subscription operation binding the contract event 0x7dc5c0699ac8dd5250cbe368a2fc3b4a2daadb120ad07f6cccea29f83482686e.
//
// Solidity: event Execution721MakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 makerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchExecution721MakerFeePacked(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceExecution721MakerFeePacked) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Execution721MakerFeePacked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceExecution721MakerFeePacked)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721MakerFeePacked", log); err != nil {
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

// ParseExecution721MakerFeePacked is a log parse operation binding the contract event 0x7dc5c0699ac8dd5250cbe368a2fc3b4a2daadb120ad07f6cccea29f83482686e.
//
// Solidity: event Execution721MakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 makerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseExecution721MakerFeePacked(log types.Log) (*BlurExchangeV2InterfaceExecution721MakerFeePacked, error) {
	event := new(BlurExchangeV2InterfaceExecution721MakerFeePacked)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721MakerFeePacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceExecution721PackedIterator is returned from FilterExecution721Packed and is used to iterate over the raw logs and unpacked data for Execution721Packed events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721PackedIterator struct {
	Event *BlurExchangeV2InterfaceExecution721Packed // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceExecution721PackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceExecution721Packed)
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
		it.Event = new(BlurExchangeV2InterfaceExecution721Packed)
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
func (it *BlurExchangeV2InterfaceExecution721PackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceExecution721PackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceExecution721Packed represents a Execution721Packed event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721Packed struct {
	OrderHash                 [32]byte
	TokenIdListingIndexTrader *big.Int
	CollectionPriceSide       *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterExecution721Packed is a free log retrieval operation binding the contract event 0x1d5e12b51dee5e4d34434576c3fb99714a85f57b0fd546ada4b0bddd736d12b2.
//
// Solidity: event Execution721Packed(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterExecution721Packed(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceExecution721PackedIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Execution721Packed")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceExecution721PackedIterator{contract: _BlurExchangeV2Interface.contract, event: "Execution721Packed", logs: logs, sub: sub}, nil
}

// WatchExecution721Packed is a free log subscription operation binding the contract event 0x1d5e12b51dee5e4d34434576c3fb99714a85f57b0fd546ada4b0bddd736d12b2.
//
// Solidity: event Execution721Packed(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchExecution721Packed(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceExecution721Packed) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Execution721Packed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceExecution721Packed)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721Packed", log); err != nil {
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

// ParseExecution721Packed is a log parse operation binding the contract event 0x1d5e12b51dee5e4d34434576c3fb99714a85f57b0fd546ada4b0bddd736d12b2.
//
// Solidity: event Execution721Packed(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseExecution721Packed(log types.Log) (*BlurExchangeV2InterfaceExecution721Packed, error) {
	event := new(BlurExchangeV2InterfaceExecution721Packed)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721Packed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceExecution721TakerFeePackedIterator is returned from FilterExecution721TakerFeePacked and is used to iterate over the raw logs and unpacked data for Execution721TakerFeePacked events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721TakerFeePackedIterator struct {
	Event *BlurExchangeV2InterfaceExecution721TakerFeePacked // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceExecution721TakerFeePackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceExecution721TakerFeePacked)
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
		it.Event = new(BlurExchangeV2InterfaceExecution721TakerFeePacked)
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
func (it *BlurExchangeV2InterfaceExecution721TakerFeePackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceExecution721TakerFeePackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceExecution721TakerFeePacked represents a Execution721TakerFeePacked event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceExecution721TakerFeePacked struct {
	OrderHash                 [32]byte
	TokenIdListingIndexTrader *big.Int
	CollectionPriceSide       *big.Int
	TakerFeeRecipientRate     *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterExecution721TakerFeePacked is a free log retrieval operation binding the contract event 0x0fcf17fac114131b10f37b183c6a60f905911e52802caeeb3e6ea210398b81ab.
//
// Solidity: event Execution721TakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 takerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterExecution721TakerFeePacked(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceExecution721TakerFeePackedIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Execution721TakerFeePacked")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceExecution721TakerFeePackedIterator{contract: _BlurExchangeV2Interface.contract, event: "Execution721TakerFeePacked", logs: logs, sub: sub}, nil
}

// WatchExecution721TakerFeePacked is a free log subscription operation binding the contract event 0x0fcf17fac114131b10f37b183c6a60f905911e52802caeeb3e6ea210398b81ab.
//
// Solidity: event Execution721TakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 takerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchExecution721TakerFeePacked(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceExecution721TakerFeePacked) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Execution721TakerFeePacked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceExecution721TakerFeePacked)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721TakerFeePacked", log); err != nil {
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

// ParseExecution721TakerFeePacked is a log parse operation binding the contract event 0x0fcf17fac114131b10f37b183c6a60f905911e52802caeeb3e6ea210398b81ab.
//
// Solidity: event Execution721TakerFeePacked(bytes32 orderHash, uint256 tokenIdListingIndexTrader, uint256 collectionPriceSide, uint256 takerFeeRecipientRate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseExecution721TakerFeePacked(log types.Log) (*BlurExchangeV2InterfaceExecution721TakerFeePacked, error) {
	event := new(BlurExchangeV2InterfaceExecution721TakerFeePacked)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Execution721TakerFeePacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceInitializedIterator struct {
	Event *BlurExchangeV2InterfaceInitialized // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceInitialized)
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
		it.Event = new(BlurExchangeV2InterfaceInitialized)
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
func (it *BlurExchangeV2InterfaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceInitialized represents a Initialized event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceInitializedIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceInitializedIterator{contract: _BlurExchangeV2Interface.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceInitialized) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceInitialized)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseInitialized(log types.Log) (*BlurExchangeV2InterfaceInitialized, error) {
	event := new(BlurExchangeV2InterfaceInitialized)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceNewBlockRangeIterator is returned from FilterNewBlockRange and is used to iterate over the raw logs and unpacked data for NewBlockRange events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewBlockRangeIterator struct {
	Event *BlurExchangeV2InterfaceNewBlockRange // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceNewBlockRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceNewBlockRange)
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
		it.Event = new(BlurExchangeV2InterfaceNewBlockRange)
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
func (it *BlurExchangeV2InterfaceNewBlockRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceNewBlockRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceNewBlockRange represents a NewBlockRange event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewBlockRange struct {
	BlockRange *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBlockRange is a free log retrieval operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterNewBlockRange(opts *bind.FilterOpts) (*BlurExchangeV2InterfaceNewBlockRangeIterator, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceNewBlockRangeIterator{contract: _BlurExchangeV2Interface.contract, event: "NewBlockRange", logs: logs, sub: sub}, nil
}

// WatchNewBlockRange is a free log subscription operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchNewBlockRange(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceNewBlockRange) (event.Subscription, error) {

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceNewBlockRange)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
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

// ParseNewBlockRange is a log parse operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseNewBlockRange(log types.Log) (*BlurExchangeV2InterfaceNewBlockRange, error) {
	event := new(BlurExchangeV2InterfaceNewBlockRange)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewGovernorIterator struct {
	Event *BlurExchangeV2InterfaceNewGovernor // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceNewGovernor)
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
		it.Event = new(BlurExchangeV2InterfaceNewGovernor)
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
func (it *BlurExchangeV2InterfaceNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceNewGovernor represents a NewGovernor event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewGovernor struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address indexed governor)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterNewGovernor(opts *bind.FilterOpts, governor []common.Address) (*BlurExchangeV2InterfaceNewGovernorIterator, error) {

	var governorRule []interface{}
	for _, governorItem := range governor {
		governorRule = append(governorRule, governorItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "NewGovernor", governorRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceNewGovernorIterator{contract: _BlurExchangeV2Interface.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address indexed governor)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceNewGovernor, governor []common.Address) (event.Subscription, error) {

	var governorRule []interface{}
	for _, governorItem := range governor {
		governorRule = append(governorRule, governorItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "NewGovernor", governorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceNewGovernor)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address indexed governor)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseNewGovernor(log types.Log) (*BlurExchangeV2InterfaceNewGovernor, error) {
	event := new(BlurExchangeV2InterfaceNewGovernor)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceNewProtocolFeeIterator is returned from FilterNewProtocolFee and is used to iterate over the raw logs and unpacked data for NewProtocolFee events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewProtocolFeeIterator struct {
	Event *BlurExchangeV2InterfaceNewProtocolFee // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceNewProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceNewProtocolFee)
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
		it.Event = new(BlurExchangeV2InterfaceNewProtocolFee)
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
func (it *BlurExchangeV2InterfaceNewProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceNewProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceNewProtocolFee represents a NewProtocolFee event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNewProtocolFee struct {
	Recipient common.Address
	Rate      uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFee is a free log retrieval operation binding the contract event 0x1d9e390a0f55a4e3251a1de541b3da1cb012fd85d9b4f0098bcffb70c4f4032d.
//
// Solidity: event NewProtocolFee(address indexed recipient, uint16 indexed rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterNewProtocolFee(opts *bind.FilterOpts, recipient []common.Address, rate []uint16) (*BlurExchangeV2InterfaceNewProtocolFeeIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "NewProtocolFee", recipientRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceNewProtocolFeeIterator{contract: _BlurExchangeV2Interface.contract, event: "NewProtocolFee", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFee is a free log subscription operation binding the contract event 0x1d9e390a0f55a4e3251a1de541b3da1cb012fd85d9b4f0098bcffb70c4f4032d.
//
// Solidity: event NewProtocolFee(address indexed recipient, uint16 indexed rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchNewProtocolFee(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceNewProtocolFee, recipient []common.Address, rate []uint16) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "NewProtocolFee", recipientRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceNewProtocolFee)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewProtocolFee", log); err != nil {
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

// ParseNewProtocolFee is a log parse operation binding the contract event 0x1d9e390a0f55a4e3251a1de541b3da1cb012fd85d9b4f0098bcffb70c4f4032d.
//
// Solidity: event NewProtocolFee(address indexed recipient, uint16 indexed rate)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseNewProtocolFee(log types.Log) (*BlurExchangeV2InterfaceNewProtocolFee, error) {
	event := new(BlurExchangeV2InterfaceNewProtocolFee)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NewProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNonceIncrementedIterator struct {
	Event *BlurExchangeV2InterfaceNonceIncremented // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceNonceIncremented)
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
		it.Event = new(BlurExchangeV2InterfaceNonceIncremented)
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
func (it *BlurExchangeV2InterfaceNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceNonceIncremented represents a NonceIncremented event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceNonceIncremented struct {
	User     common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterNonceIncremented(opts *bind.FilterOpts, user []common.Address) (*BlurExchangeV2InterfaceNonceIncrementedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "NonceIncremented", userRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceNonceIncrementedIterator{contract: _BlurExchangeV2Interface.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceNonceIncremented, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "NonceIncremented", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceNonceIncremented)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseNonceIncremented(log types.Log) (*BlurExchangeV2InterfaceNonceIncremented, error) {
	event := new(BlurExchangeV2InterfaceNonceIncremented)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceOwnershipTransferStartedIterator struct {
	Event *BlurExchangeV2InterfaceOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceOwnershipTransferStarted)
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
		it.Event = new(BlurExchangeV2InterfaceOwnershipTransferStarted)
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
func (it *BlurExchangeV2InterfaceOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlurExchangeV2InterfaceOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceOwnershipTransferStartedIterator{contract: _BlurExchangeV2Interface.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceOwnershipTransferStarted)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseOwnershipTransferStarted(log types.Log) (*BlurExchangeV2InterfaceOwnershipTransferStarted, error) {
	event := new(BlurExchangeV2InterfaceOwnershipTransferStarted)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceOwnershipTransferredIterator struct {
	Event *BlurExchangeV2InterfaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceOwnershipTransferred)
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
		it.Event = new(BlurExchangeV2InterfaceOwnershipTransferred)
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
func (it *BlurExchangeV2InterfaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceOwnershipTransferred represents a OwnershipTransferred event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlurExchangeV2InterfaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceOwnershipTransferredIterator{contract: _BlurExchangeV2Interface.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceOwnershipTransferred)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseOwnershipTransferred(log types.Log) (*BlurExchangeV2InterfaceOwnershipTransferred, error) {
	event := new(BlurExchangeV2InterfaceOwnershipTransferred)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceSetOracleIterator is returned from FilterSetOracle and is used to iterate over the raw logs and unpacked data for SetOracle events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceSetOracleIterator struct {
	Event *BlurExchangeV2InterfaceSetOracle // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceSetOracle)
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
		it.Event = new(BlurExchangeV2InterfaceSetOracle)
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
func (it *BlurExchangeV2InterfaceSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceSetOracle represents a SetOracle event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceSetOracle struct {
	User     common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOracle is a free log retrieval operation binding the contract event 0xcc852792b7afae13c99037685c90dd3be44073d4bc32aa8c1b72fd07a2ac537c.
//
// Solidity: event SetOracle(address indexed user, bool approved)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterSetOracle(opts *bind.FilterOpts, user []common.Address) (*BlurExchangeV2InterfaceSetOracleIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "SetOracle", userRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceSetOracleIterator{contract: _BlurExchangeV2Interface.contract, event: "SetOracle", logs: logs, sub: sub}, nil
}

// WatchSetOracle is a free log subscription operation binding the contract event 0xcc852792b7afae13c99037685c90dd3be44073d4bc32aa8c1b72fd07a2ac537c.
//
// Solidity: event SetOracle(address indexed user, bool approved)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchSetOracle(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceSetOracle, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "SetOracle", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceSetOracle)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "SetOracle", log); err != nil {
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

// ParseSetOracle is a log parse operation binding the contract event 0xcc852792b7afae13c99037685c90dd3be44073d4bc32aa8c1b72fd07a2ac537c.
//
// Solidity: event SetOracle(address indexed user, bool approved)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseSetOracle(log types.Log) (*BlurExchangeV2InterfaceSetOracle, error) {
	event := new(BlurExchangeV2InterfaceSetOracle)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "SetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlurExchangeV2InterfaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceUpgradedIterator struct {
	Event *BlurExchangeV2InterfaceUpgraded // Event containing the contract specifics and raw log

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
func (it *BlurExchangeV2InterfaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlurExchangeV2InterfaceUpgraded)
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
		it.Event = new(BlurExchangeV2InterfaceUpgraded)
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
func (it *BlurExchangeV2InterfaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlurExchangeV2InterfaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlurExchangeV2InterfaceUpgraded represents a Upgraded event raised by the BlurExchangeV2Interface contract.
type BlurExchangeV2InterfaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BlurExchangeV2InterfaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BlurExchangeV2InterfaceUpgradedIterator{contract: _BlurExchangeV2Interface.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BlurExchangeV2InterfaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BlurExchangeV2Interface.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlurExchangeV2InterfaceUpgraded)
				if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BlurExchangeV2Interface *BlurExchangeV2InterfaceFilterer) ParseUpgraded(log types.Log) (*BlurExchangeV2InterfaceUpgraded, error) {
	event := new(BlurExchangeV2InterfaceUpgraded)
	if err := _BlurExchangeV2Interface.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
