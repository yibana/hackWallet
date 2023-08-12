package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yibana/hackWallet/internal/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"testing"
	"time"
)

var Wallet *hackWallet.HackWallet

func init() {
	var err error
	AnvilFork := false
	Wallet, err = hackWallet.NewHackWallet(configs.WSS_RPC_URL, AnvilFork)
	if err != nil {
		panic(err)
	}
}

func TestPrintAccountsBalance(t *testing.T) {
	for i, account := range Wallet.Accounts {
		balance, _ := account.GetBalance()
		hackWallet.DebugLog(fmt.Sprintf("[%d]acc:%s balance:%f\n", i+1, account.Address.String(), hackWallet.ConvertWei2Eth(balance)))
	}
}

func TestGetBlockHeader(t *testing.T) {
	for {
		blockHeader, _ := Wallet.GetBlockHeader(nil)
		next_baseFee, _ := Wallet.GetNextBlockBaseFee()
		fmt.Printf("time:%d blockNumber:%d baseFee:%d next_baseFee:%d\n", time.Now().Unix(),
			blockHeader.Number.Uint64(), blockHeader.BaseFee.Uint64(), next_baseFee)
		time.Sleep(time.Second)
	}
}

func TestSubscribe_alchemy_pendingTransactions(t *testing.T) {
	Wallet.SubscribeAlchemy_Pendingtransactions(hackWallet.Subscribe_alchemy_pendingTransactions_params{
		ToAddress:  []string{"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"},
		HashesOnly: false,
	},
		func(tx *types.Transaction) {
			json, err := tx.MarshalJSON()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		})
}

func TestSubscribeLogs(t *testing.T) {
	Wallet.SubscribeLogs([]string{"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"}, []string{}, func(tx *types.Log) {
		json, err := tx.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	})
}

func TestSubscribe_alchemy_newBlocks(t *testing.T) {
	Wallet.SubscribeHeader(func(header *types.Header) {
		json, err := header.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	})
}
