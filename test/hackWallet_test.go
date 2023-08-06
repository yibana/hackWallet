package test

import (
	"fmt"
	"github.com/yibana/hackWallet/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"testing"
	"time"
)

var Wallet *hackWallet.HackWallet

func init() {
	var err error
	Wallet, err = hackWallet.NewHackWallet(configs.HTTP_RPC_URL, false)
	if err != nil {
		panic(err)
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
