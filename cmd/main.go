package main

import (
	"github.com/yibana/hackWallet/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"time"
)

func main() {
	println(configs.PROXY)
	wallet, err := hackWallet.NewHackWallet(configs.HTTP_RPC_URL, true)
	if err != nil {
		panic(err)
	}
	wallet.AddRandAccount()
	wallet.AddRandAccount()
	wallet.AddRandAccount()

	for i, account := range wallet.Accounts {
		println(i, account.Address.String())
		balance, err := account.GetBalance()
		if err != nil {
			continue
		}
		println(hackWallet.BigIntToDecimal_18(balance))
	}

	for {
		println(wallet.GetBlockNumber())
		time.Sleep(time.Second * 12)
	}

}
