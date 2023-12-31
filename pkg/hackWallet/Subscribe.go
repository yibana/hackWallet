package hackWallet

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
func (hw *HackWallet) SubscribeAlchemy_Pendingtransactions(ctx context.Context, param Subscribe_alchemy_pendingTransactions_params, callback func(tx *types.Transaction)) error {
	return hw.Subscribe(ctx, func(data interface{}) {
		// convertToType(data, &types.Transaction)
		var tx types.Transaction
		marshal, _ := json.Marshal(data)
		json.Unmarshal(marshal, &tx)
		go callback(&tx)
	}, "eth", "alchemy_pendingTransactions", param)
}

func (hw *HackWallet) SubscribePendingTransactions(ctx context.Context, callback func(txHash common.Hash)) error {
	return hw.Subscribe(ctx, func(data interface{}) {
		// convertToType(data, &types.Transaction)
		var txHash common.Hash
		marshal, _ := json.Marshal(data)
		json.Unmarshal(marshal, &txHash)
		go callback(txHash)
	}, "eth", "newPendingTransactions")
}

// Subscribe newHeads
func (hw *HackWallet) SubscribeHeader(ctx context.Context, callback func(header *types.Header)) error {
	return hw.Subscribe(ctx, func(data interface{}) {
		var header types.Header
		marshal, err := json.Marshal(data)
		if err != nil {
			ErrLog("SubscribeHeader", zap.Error(err))
			return
		}
		err = json.Unmarshal(marshal, &header)
		if err != nil {
			ErrLog("SubscribeHeader", zap.Error(err))
			return
		}
		if hw.lastBlockHeader == nil || header.Number.Cmp(hw.lastBlockHeader.Number) > 0 {
			hw.lastBlockHeader = &header
			hw.lastBlockHeader_time = uint64(time.Now().Unix())
		}

		go callback(&header)
	}, "eth", "newHeads")
}

type SubscribeLogsParams struct {
	Address []string   `json:"address,omitempty"`
	Topics  [][]string `json:"topics,omitempty"`
}

// Subscribe logs
func (hw *HackWallet) SubscribeLogs(ctx context.Context, params SubscribeLogsParams, callback func(tx *types.Log)) error {

	return hw.Subscribe(ctx, func(data interface{}) {
		var log types.Log
		marshal, _ := json.Marshal(data)
		json.Unmarshal(marshal, &log)
		go callback(&log)
	}, "eth", "logs", params)
}

func (hw *HackWallet) Subscribe(ctx context.Context, callback func(data interface{}), namespace string, args ...interface{}) error {
	var channel interface{}
	channel = make(chan interface{})
	var err_cn int
	var ctxDone bool
	if ctx == nil {
		ctx = context.Background()
	}
	for !ctxDone {
		sub, err := hw.rpc.Subscribe(ctx,
			namespace, channel, args...,
		)
		if err != nil {
			ErrLog(fmt.Sprintf("[%d]Subscribe_%s", err_cn, args[0]), zap.Error(err))
			err_cn++
			if err_cn > 5 {
				return err
			}
			time.Sleep(time.Second * 5)
			continue
		}
		err_cn = 0

		loop := func() {
			for {
				select {
				case v := <-channel.(chan interface{}):
					go callback(&v)
				case err = <-sub.Err():
					ErrLog(fmt.Sprintf("[%d]Subscribe_%s", err_cn, args[0]), zap.Error(err))
					return
				case <-ctx.Done():
					ctxDone = true
					return
				}
			}
		}
		loop()
		sub.Unsubscribe()

	}
	return nil
}
