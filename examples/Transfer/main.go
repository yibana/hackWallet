package main

import (
	"fmt"
	"github.com/yibana/hackWallet/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
)

func main() {
	wallet, err := hackWallet.NewHackWallet(configs.HTTP_RPC_URL, true)
	if err != nil {
		panic(err)
	}
	myAccount, err := wallet.AddRandAccount()
	if err != nil {
		panic(err)
	}

	for i, account := range wallet.Accounts {
		balance, _ := account.GetBalance()
		fmt.Printf("[%d]acc:%s balance:%f\n", i+1,
			account.Address.String(), hackWallet.ConvertWei2Eth(balance))
	}

	transfer, err := wallet.Transfer(wallet.Accounts[0], myAccount.Address, hackWallet.ConvertETHToBigInt(1))
	if err != nil {
		panic(err)
	}
	receipt, err := wallet.WaitForTx(transfer, 12*3)
	if err != nil {
		panic(err)
	}

	if receipt.Status == 0 {
		panic("transfer failed")
	}

	balance, _ := wallet.Accounts[0].GetBalance()
	fmt.Printf("Accounts[0]:%s balance:%f\n", wallet.Accounts[0].Address.String(), hackWallet.ConvertWei2Eth(balance))

	balance, _ = myAccount.GetBalance()
	fmt.Printf("myAccount:%s balance:%f\n", myAccount.Address.String(), hackWallet.ConvertWei2Eth(balance))

}
