package hackWallet

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"time"
)

type Subscribe_alchemy_pendingTransactions_params struct {
	FromAddress []string `json:"fromAddress"`
	ToAddress   []string `json:"toAddress"`
	HashesOnly  bool     `json:"hashesOnly"` // if true, only return hashes of transactions, not the full transactions
}

// https://docs.alchemy.com/reference/alchemy-pendingtransactions
func (hw *HackWallet) SubscribeAlchemy_Pendingtransactions(param Subscribe_alchemy_pendingTransactions_params, callback func(tx *types.Transaction)) {
	hw.Subscribe(func(data interface{}) {
		// convertToType(data, &types.Transaction)
		var tx types.Transaction
		marshal, _ := json.Marshal(data)
		json.Unmarshal(marshal, &tx)
		go callback(&tx)
	}, "eth", "alchemy_pendingTransactions", param)
}

// Subscribe logs
func (hw *HackWallet) SubscribeLogs(address []string, topics []string, callback func(tx *types.Log)) {
	hw.Subscribe(func(data interface{}) {
		var log types.Log
		marshal, _ := json.Marshal(data)
		json.Unmarshal(marshal, &log)
		go callback(&log)
	}, "eth", "logs", address, topics)
}

func (hw *HackWallet) Subscribe(callback func(data interface{}), namespace string, args ...interface{}) {
	var channel interface{}
	for {
		channel = make(chan interface{})
		sub, err := hw.rpc.Subscribe(context.Background(),
			namespace, channel, args...,
		)
		if err == nil {
			isexit := false
			for !isexit {
				select {
				case v := <-channel.(chan interface{}):
					go callback(&v)
				case err = <-sub.Err():
					ErrLog(fmt.Sprintf("Subscribe_%s", args[0]), zap.Error(err))
					isexit = true
				}
			}
			sub.Unsubscribe()
		} else {
			ErrLog(fmt.Sprintf("Subscribe_%s", args[0]), zap.Error(err))
		}
		time.Sleep(time.Second * 5)
	}

}
