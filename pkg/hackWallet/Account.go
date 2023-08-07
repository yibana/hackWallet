package hackWallet

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
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
	return GetBalance(a.ethclient, *a.Address)
}

// BuildTransaction
func (a *Account) BuildTransaction(
	to *common.Address,
	value *big.Int, data []byte,
	nonce, gasLimit uint64,
	chainID, GasFeeCap, GasTipCap *big.Int) (*types.Transaction, error) {
	return BuildTransaction(a.ethclient, a.PrivateKey, *a.Address, to, value, data, nonce, gasLimit, chainID, GasFeeCap, GasTipCap)
}

func (a *Account) SendTransaction(tx *types.Transaction) error {
	return SendTransaction(a.ethclient, tx)
}

func (a *Account) GetTokenBalance(erc20 common.Address) (*big.Int, error) {
	return GetTokenBalance(a.ethclient, erc20, *a.Address)
}

// WaitForTx
func (a *Account) WaitForTx(tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {
	return WaitForTx(a.ethclient, tx, maxWaitSeconds)
}

// Build_WETH_deposit
func (a *Account) Build_WETH_deposit(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	amount *big.Int) (*types.Transaction, error) {
	data, err := WETHABI.Pack("deposit")
	if err != nil {
		return nil, err
	}
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap)
	transaction, err := a.BuildTransaction(&WETH9, amount, data, nonce, DefaultWETHDepositGas, chainID, GasFeeCap, GasTipCap)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (a *Account) Build_WETH_withdraw(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	amount *big.Int) (*types.Transaction, error) {
	data, err := WETHABI.Pack("withdraw", amount)
	if err != nil {
		return nil, err
	}
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap)
	transaction, err := a.BuildTransaction(&WETH9, big.NewInt(0), data, nonce, DefaultWETHWithdrawGas, chainID, GasFeeCap, GasTipCap)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
