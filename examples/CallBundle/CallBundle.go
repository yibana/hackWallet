package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/yibana/hackWallet/internal/configs"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"math/big"
)

func main() {
	wallet, err := hackWallet.NewHackWallet(configs.WSS_RPC_URL, false)
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
		})

	if err != nil {
		panic(err)
	}

	//fbPrvKey, _ := wallet.AddRandAccount()
	privateKey, _ := crypto.HexToECDSA(
		"2e19800fcbbf0abb7cf6d72ee7171f08943bc8e5c3568d1d7420e52136898154",
	)
	err = wallet.InitFlashbot(privateKey, false)

	if err != nil {
		panic(err)
	}

	response, err := wallet.CallBundle(transactions, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	// 如果是goerli网络,因为可以打包的验证器极少,所以需要多尝试几次
	//err = wallet.SendBundleFrom(transactions, 0, 2)

	if err != nil {
		panic(err)
	}

	balance, err := wallet.GetTokenBalance(myacc1, weth.Address)
	if err != nil {
		panic(err)
	}

	fmt.Printf("WETH balance: %f\n", hackWallet.ConvertWei2Eth(balance))

}
