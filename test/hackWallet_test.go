package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yibana/hackWallet/internal/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"math/big"
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

func TestSubscribe_log1(t *testing.T) {
	const (
		SpotPriceUpdateFilter = "0xf06180fdbe95e5193df4dcd1352726b1f04cb58599ce58552cc952447af2ffbb"
		SwapNFTInPairFilter   = "0x3614eb567740a0ee3897c0e2b11ad6a5720d2e4438f9c8accf6c95c24af3a470"
		TokenWithdrawalFilter = "0x0e266e8f38544aa1480d73762386eb10df55b1b8453d935762e891c44b69a1e6"
		SwapNFTOutPairFilter  = "0xbc479dfc6cb9c1a9d880f987ee4b30fa43dd7f06aec121db685b67d587c93c93"
	)
	var address_list []common.Address
	//bytes, _ := ioutil.ReadFile("address_list")
	//json.Unmarshal(bytes, &address_list)

	address_list = append(address_list, common.HexToAddress("0x7c0BA1f5CF601834f084b23e7DC3f75e9695a2Dd"))

	logs, err := Wallet.RPCClient.FilterLogs(context.Background(),
		ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(17566372)),
			ToBlock:   big.NewInt(int64(17877861)),
			Addresses: address_list,
			Topics: [][]common.Hash{
				{
					common.HexToHash(SwapNFTOutPairFilter),
					common.HexToHash(SwapNFTInPairFilter),
					common.HexToHash(SpotPriceUpdateFilter),
					common.HexToHash(TokenWithdrawalFilter),
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	for _, log := range logs {
		fmt.Println(log)
	}
}
