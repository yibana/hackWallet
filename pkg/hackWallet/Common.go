package hackWallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ERC20Interface "github.com/yibana/hackWallet/contracts/ERC20"
	ERC721Interface "github.com/yibana/hackWallet/contracts/ERC721"
	WETH9Interface "github.com/yibana/hackWallet/contracts/WETH9"
	"math/big"
	"time"
)

var WETH_ABI, _ = WETH9Interface.WETH9InterfaceMetaData.GetAbi()
var ERC20_ABI, _ = ERC20Interface.ERC20InterfaceMetaData.GetAbi()
var ERC721_ABI, _ = ERC721Interface.ERC721InterfaceMetaData.GetAbi()

const (
	DefaultGasTipCap               = 1e8 //0.1 Gwei
	DefaultWETHDepositGas   uint64 = 60000
	DefaultWETHWithdrawGas  uint64 = 50000
	DefaultERC721ApproveGas uint64 = 60000
)

// WaitForTx
func WaitForTx(ethclient *ethclient.Client, tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {

	DebugLog("Waiting for tx to complete...")
	loops := uint(0)
	for {
		_, isPending, err := ethclient.TransactionByHash(context.Background(), tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return nil, err
		}

		if !isPending {
			// It's not pending, so check if it's been mined
			receipt, err := ethclient.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil && err != ethereum.NotFound {
				return nil, err
			}
			if receipt != nil {
				DebugLog("tx complete. it may take a few minutes to appear in etherscan.")
				DebugLog(fmt.Sprintf("https://etherscan.io/tx/%s", tx.Hash().Hex()))
				return receipt, nil
			}
		}

		time.Sleep(1 * time.Second)

		loops = loops + 1
		if loops > maxWaitSeconds {
			DebugLog(fmt.Sprintf("timed out after %d seconds. check manually here:", maxWaitSeconds))
			DebugLog(fmt.Sprintf("https://etherscan.io/tx/%s", tx.Hash().Hex()))
			return nil, errors.New("等待tx超时，上链失败！")
		}
	}
}

func SimulationActionFrom(ethclient *ethclient.Client, from, to common.Address, value *big.Int, data []byte) (gas uint64, err error) {
	return ethclient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:       from,
		To:         &to,
		Gas:        0,
		GasPrice:   nil,
		GasFeeCap:  nil,
		GasTipCap:  nil,
		Value:      value,
		Data:       data,
		AccessList: nil,
	})
}

func GetTokenBalance(ethclient *ethclient.Client, from common.Address, erc20 common.Address) (*big.Int, error) {
	erc20Interface, err := ERC20Interface.NewERC20Interface(erc20, ethclient)
	if err != nil {
		return nil, err
	}
	return erc20Interface.BalanceOf(nil, from)
}

func GetBalance(ethclient *ethclient.Client, from common.Address) (*big.Int, error) {
	return ethclient.BalanceAt(context.Background(), from, nil)
}

func SendTransaction(ethclient *ethclient.Client, tx *types.Transaction) error {
	return ethclient.SendTransaction(context.Background(), tx)
}

func BuildTransaction(
	ethclient *ethclient.Client,
	PrivateKey *ecdsa.PrivateKey,
	from common.Address,
	to *common.Address,
	value *big.Int, data []byte,
	nonce, gasLimit uint64,
	chainID, GasFeeCap, GasTipCap *big.Int) (*types.Transaction, error) {
	var err error
	if nonce == 0 {
		nonce, err = ethclient.PendingNonceAt(context.Background(), from)
		if err != nil {
			return nil, err
		}
	}

	// eip155
	tx := types.NewTx(&types.DynamicFeeTx{Nonce: nonce, To: to,
		Value: value, Gas: gasLimit, Data: data, GasFeeCap: GasFeeCap, GasTipCap: GasTipCap})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), PrivateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
