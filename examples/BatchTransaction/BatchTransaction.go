package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yibana/hackWallet/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"math/big"
)

func main() {
	wallet, err := hackWallet.NewHackWallet(configs.HTTP_RPC_URL, true)
	if err != nil {
		panic(err)
	}

	gasTipCap := big.NewInt(1e8) // 设置默认小费
	chainId := wallet.GetChainID()

	transactions, err := wallet.BuildBatchTransaction(wallet.Accounts[0],
		func(from *hackWallet.Account, baseFee *big.Int, nonce uint64) (*types.Transaction, error) {
			return from.Build_WETH_deposit(baseFee, nonce, chainId, gasTipCap, hackWallet.ConvertETHToBigInt(1.96))
		},
		func(from *hackWallet.Account, baseFee *big.Int, nonce uint64) (*types.Transaction, error) {
			return from.Build_WETH_withdraw(baseFee, nonce, chainId, gasTipCap, hackWallet.ConvertETHToBigInt(0.96))
		},
	)
	if err != nil {
		panic(err)
	}

	err = wallet.SendBatchTransaction(transactions)

	if err != nil {
		panic(err)
	}

	balance, err := wallet.GetTokenBalance(wallet.Accounts[0], hackWallet.WETH9)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WETH balance: %f\n", hackWallet.ConvertWei2Eth(balance))

}
