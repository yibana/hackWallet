package hackWallet

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type HackWallet struct {
	Accounts        []*Account
	ProviderURL     string
	RPCClient       *ethclient.Client
	lastBlockHeader *types.Header
}

func NewHackWallet(rpcUrl string, is_anvil_fork bool) (*HackWallet, error) {
	var accounts []*Account
	var err error
	var privateKeys []string
	if is_anvil_fork {
		privateKeys, err = anvil_fork(rpcUrl)
		if err != nil {
			return nil, err
		}
		rpcUrl = "http://127.0.0.1:8545"
	}

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	for _, privateKey := range privateKeys {
		if acc, err := NewAccount(privateKey, client); err == nil {
			accounts = append(accounts, acc)
		}
	}

	return &HackWallet{
		Accounts:    accounts,
		RPCClient:   client,
		ProviderURL: rpcUrl,
	}, nil
}

func (hw *HackWallet) AddAccount(account *Account) {
	hw.Accounts = append(hw.Accounts, account)
}

// add rand account
func (hw *HackWallet) AddRandAccount() error {
	account, err := RandAccount(hw.RPCClient)
	if err != nil {
		return err
	}
	hw.AddAccount(account)
	return nil
}

// get blocknumber
func (hw *HackWallet) GetBlockNumber() (uint64, error) {
	header, err := hw.GetBlockHeader(nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// get block header
func (hw *HackWallet) GetBlockHeader(blockNumber *big.Int) (*types.Header, error) {
	var err error
	if hw.lastBlockHeader != nil {
		nowTime := uint64(time.Now().Unix())
		if nowTime > hw.lastBlockHeader.Time && nowTime <= hw.lastBlockHeader.Time+11 {
			return hw.lastBlockHeader, nil
		}
	}
	hw.lastBlockHeader, err = hw.RPCClient.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return hw.lastBlockHeader, nil
}

// GetNextBlockBaseFee
func (hw *HackWallet) GetNextBlockBaseFee() (*big.Int, error) {
	header, err := hw.GetBlockHeader(nil)
	if err != nil {
		return nil, err
	}
	return GetNextBlockBaseFee(header)
}
