package test

import (
	"fmt"
	"github.com/yibana/hackWallet/configs"
	WETH9Interface "github.com/yibana/hackWallet/contracts/WETH9"
	"github.com/yibana/hackWallet/pkg/hackWallet"
	"math/big"
	"testing"
	"time"
)

var Wallet *hackWallet.HackWallet

func init() {
	var err error
	AnvilFork := true
	Wallet, err = hackWallet.NewHackWallet(configs.HTTP_RPC_URL, AnvilFork)
	if err != nil {
		panic(err)
	}
}

func TestPrintAccountsBalance(t *testing.T) {
	for i, account := range Wallet.Accounts {
		balance, _ := account.GetBalance()
		fmt.Printf("[%d]acc:%s balance:%f\n", i+1, account.Address.String(), hackWallet.ConvertWei2Eth(balance))
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

func TestBuildTransaction(t *testing.T) {
	abi, _ := WETH9Interface.WETH9InterfaceMetaData.GetAbi()

	fromAcc := Wallet.Accounts[0]

	// pack deposit 0.1eth
	depositData, err := abi.Pack("deposit")
	if err != nil {
		panic(err)
	}
	transaction, err := Wallet.BuildTransaction(fromAcc, &hackWallet.WETH9, depositData,
		hackWallet.ConvertETHToBigInt(0.1), 210000, 0, big.NewInt(1e8),
	)
	if err != nil {
		panic(err)
	}

	err = fromAcc.SendTransaction(transaction)
	if err != nil {
		panic(err)
	}

	//check transaction
	receipt, err := fromAcc.WaitForTx(transaction, 12*3)
	if err != nil {
		panic(err)
	}
	if receipt.Status == 0 {
		panic("deposit failed")
	}

	fmt.Println("deposit success")

	// get weth balance
	balance, err := fromAcc.GetTokenBalance(hackWallet.WETH9)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance:", hackWallet.ConvertWei2Eth(balance))

}
