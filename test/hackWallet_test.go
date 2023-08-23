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
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("cancel")
		cancelFunc()
	}()
	Wallet.SubscribeAlchemy_Pendingtransactions(ctx, hackWallet.Subscribe_alchemy_pendingTransactions_params{
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

	t.Log("end")

}

func TestSubscribePendingTransactions(t *testing.T) {
	Wallet.SubscribePendingTransactions(nil, func(txHash common.Hash) {
		t.Log(txHash.String())
	})
}

func TestSubscribeLogs(t *testing.T) {
	param := hackWallet.SubscribeLogsParams{
		Topics: [][]string{
			{
				//"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c",
				"0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118",
				"0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9",
			},
		},
		Address: []string{
			//"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"0x1F98431c8aD98523631AE4a59f267346ea31F984",
			"0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f",
		},
	}
	Wallet.SubscribeLogs(nil, param, func(tx *types.Log) {
		json, err := tx.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	})
}

func TestSubscribe_alchemy_newBlocks(t *testing.T) {
	Wallet.SubscribeHeader(nil, func(header *types.Header) {
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

	address_list = append(address_list, common.HexToAddress("0x08cf2C23CE70f566E97A4002F75F74C813Bc88Ad"))

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
