package hackWallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ERC20Interface "github.com/yibana/hackWallet/contracts/ERC20"
	"math/big"
	"strings"
	"time"
)

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
	ethclient  *ethclient.Client
}

func NewAccount(PrivateKey string, ethclient *ethclient.Client) (*Account, error) {
	if strings.HasPrefix(PrivateKey, "0x") {
		PrivateKey = PrivateKey[2:]
	}
	toECDSA, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(toECDSA.PublicKey)
	return &Account{
		PrivateKey: toECDSA,
		Address:    &address,
		ethclient:  ethclient,
	}, nil
}

// rand create a private key
func RandAccount(ethclient *ethclient.Client) (*Account, error) {
	PrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(PrivateKey.PublicKey)
	return &Account{
		PrivateKey: PrivateKey,
		Address:    &address,
		ethclient:  ethclient,
	}, nil
}

func (a *Account) GetAddress() common.Address {
	return *a.Address
}

// get the balance of the account
func (a *Account) GetBalance() (*big.Int, error) {
	return a.ethclient.BalanceAt(context.Background(), *a.Address, nil)
}

// BuildTransaction
func (a *Account) BuildTransaction(
	to *common.Address,
	value *big.Int, data []byte,
	nonce, gasLimit uint64,
	chainID, GasFeeCap, GasTipCap *big.Int) (*types.Transaction, error) {
	var err error
	if nonce == 0 {
		nonce, err = a.ethclient.PendingNonceAt(context.Background(), *a.Address)
		if err != nil {
			return nil, err
		}
	}

	// eip155
	tx := types.NewTx(&types.DynamicFeeTx{Nonce: nonce, To: to,
		Value: value, Gas: gasLimit, Data: data, GasFeeCap: GasFeeCap, GasTipCap: GasTipCap})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), a.PrivateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func (a *Account) SendTransaction(tx *types.Transaction) error {
	return a.ethclient.SendTransaction(context.Background(), tx)
}

func (a *Account) GetTokenBalance(erc20 common.Address) (*big.Int, error) {
	erc20Interface, err := ERC20Interface.NewERC20Interface(erc20, a.ethclient)
	if err != nil {
		return nil, err
	}
	return erc20Interface.BalanceOf(nil, *a.Address)
}

// WaitForTx
func (a *Account) WaitForTx(tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {

	DebugLog("Waiting for tx to complete...")
	loops := uint(0)
	for {
		_, isPending, err := a.ethclient.TransactionByHash(context.Background(), tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return nil, err
		}

		if !isPending {
			// It's not pending, so check if it's been mined
			receipt, err := a.ethclient.TransactionReceipt(context.Background(), tx.Hash())
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
