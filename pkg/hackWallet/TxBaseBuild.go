package hackWallet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TxBaseBuild struct {
	From      *Account
	BaseFee   *big.Int
	Nonce     uint64
	ChainID   *big.Int
	GasTipCap *big.Int
}

func (txb *TxBaseBuild) Erc721Approve(
	nftaddr, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return txb.From.Build_ERC721_approve(
		txb.BaseFee, txb.Nonce, txb.ChainID, txb.GasTipCap,
		nftaddr, to, tokenId,
	)
}

func (txb *TxBaseBuild) WethWithdraw(
	amount *big.Int) (*types.Transaction, error) {
	return txb.From.Build_WETH_withdraw(
		txb.BaseFee, txb.Nonce, txb.ChainID, txb.GasTipCap,
		amount,
	)
}

func (txb *TxBaseBuild) WethDeposit(
	amount *big.Int) (*types.Transaction, error) {
	return txb.From.Build_WETH_deposit(
		txb.BaseFee, txb.Nonce, txb.ChainID, txb.GasTipCap, amount,
	)
}

func (txb *TxBaseBuild) Pack(
	value *big.Int, to *common.Address, gasLimit uint64,
	_abi *abi.ABI, name string, args ...interface{}) (*types.Transaction, error) {
	return txb.From.Build_Pack(
		txb.BaseFee, txb.Nonce, txb.ChainID, txb.GasTipCap,
		value, to, gasLimit,
		_abi, name, args...)
}

func (txb *TxBaseBuild) PackData(
	value *big.Int, to *common.Address, gasLimit uint64,
	data []byte) (*types.Transaction, error) {
	return txb.From.Build_Pack_Data(txb.BaseFee,
		txb.Nonce, txb.ChainID, txb.GasTipCap, value, to, gasLimit, data)
}
