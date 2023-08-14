package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yibana/hackWallet/internal/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"math/big"
)

func main() {
	wallet, err := hackWallet.NewHackWallet(configs.HTTP_RPC_URL_Goeli, false)
	if err != nil {
		panic(err)
	}

	myacc1, err := wallet.LoadAccount(configs.MYPK1)
	if err != nil {
		panic(err)
	}

	gasTipCap := big.NewInt(1e8) // 设置默认小费
	chainId := wallet.GetChainID()
	weth := hackWallet.TokenMap[chainId.Uint64()]["WETH"]

	transactions, err := wallet.BuildBatchTxn(myacc1, gasTipCap,
		func(txp *hackWallet.TxBaseBuild) (*types.Transaction, error) {
			return txp.WethDeposit(hackWallet.ConvertETHToBigInt(0.00196))
		},
		func(txp *hackWallet.TxBaseBuild) (*types.Transaction, error) {
			return txp.WethWithdraw(hackWallet.ConvertETHToBigInt(0.0019))
		},
	)
	if err != nil {
		panic(err)
	}

	fbPrvKey, _ := wallet.AddRandAccount()

	err = wallet.InitFlashbot(fbPrvKey.PrivateKey, false)

	if err != nil {
		panic(err)
	}

	response, err := wallet.CallBundle(transactions, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	// 如果是goerli网络,因为可以打包的验证器极少,所以需要多尝试几次
	err = wallet.SendBundleFrom(transactions, 0, 2)

	if err != nil {
		panic(err)
	}

	balance, err := wallet.GetTokenBalance(myacc1, weth.Address)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WETH balance: %f\n", hackWallet.ConvertWei2Eth(balance))

}
