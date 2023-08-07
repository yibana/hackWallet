package hackWallet

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"time"
)

type HackWallet struct {
	Accounts             []*Account
	ProviderURL          string
	RPCClient            *ethclient.Client
	lastBlockHeader      *types.Header
	lastBlockHeader_time uint64
	chainID              *big.Int
}

func NewHackWallet(rpcUrl string, AnvilFork bool) (*HackWallet, error) {
	var accounts []*Account
	var err error
	var privateKeys []string
	if AnvilFork {
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
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &HackWallet{
		Accounts:    accounts,
		RPCClient:   client,
		ProviderURL: rpcUrl,
		chainID:     chainID,
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
	nowTime := uint64(time.Now().Unix())
	if hw.lastBlockHeader != nil {
		if nowTime > hw.lastBlockHeader.Time && nowTime <= hw.lastBlockHeader.Time+11 {
			return hw.lastBlockHeader, nil
		}
		if hw.lastBlockHeader_time == nowTime { // 1秒保护
			return hw.lastBlockHeader, nil
		}
	}
	hw.lastBlockHeader, err = hw.RPCClient.HeaderByNumber(context.Background(), blockNumber)
	hw.lastBlockHeader_time = nowTime
	if err != nil {
		return nil, err
	}
	DebugLog(fmt.Sprintf("GetBlockHeader finshed, blockNumber:%d, time:%d", hw.lastBlockHeader.Number.Uint64(), hw.lastBlockHeader.Time))
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

func (hw *HackWallet) GetChainID() *big.Int {
	return hw.chainID
}

// select  account
func (hw *HackWallet) SelectAccount(address common.Address) (*Account, error) {
	for _, account := range hw.Accounts {
		if account.GetAddress() == address {
			return account, nil
		}
	}
	return nil, errors.New("account not found in wallet")
}

// BuildTransaction from,to,value,gas,gasPrice,nonce
func (hw *HackWallet) BuildTransaction(
	fromAcc *Account, to common.Address, data []byte,
	value *big.Int, gasLimit, nonce uint64,
	GasTipCap *big.Int, // 小费
) (*types.Transaction, error) {

	baseFee, err := hw.GetNextBlockBaseFee()
	if err != nil {
		return nil, err
	}
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap) // 计算最大支出手续费（矿工费）
	return fromAcc.BuildTransaction(&to, value, data, nonce, gasLimit, hw.chainID,
		GasFeeCap, GasTipCap)
}

func (hw *HackWallet) GetTokenBalance(fromAcc *Account, TokenAddress common.Address) (*big.Int, error) {
	return fromAcc.GetTokenBalance(TokenAddress)
}

// WaitForTx
func (hw *HackWallet) WaitForTx(fromAcc *Account, tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {
	return fromAcc.WaitForTx(tx, maxWaitSeconds)
}
