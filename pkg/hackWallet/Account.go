package hackWallet

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
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

func (a *Account) GetAddress() *common.Address {
	return a.Address
}

// get the balance of the account
func (a *Account) GetBalance() (*big.Int, error) {
	return a.ethclient.BalanceAt(context.Background(), *a.Address, nil)
}
