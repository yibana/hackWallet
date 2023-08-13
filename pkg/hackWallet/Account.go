package hackWallet

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
	ethclient  *ethclient.Client
}

func NewAccount(PrivateKey string, ethclient *ethclient.Client) (*Account, error) {
	if strings.HasPrefix(PrivateKey, "0x") {
		PrivateKey = PrivateKey[2:]
	}
	toECDSA, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(toECDSA.PublicKey)
	return &Account{
		PrivateKey: toECDSA,
		Address:    &address,
		ethclient:  ethclient,
	}, nil
}

// rand create a private key
func RandAccount(ethclient *ethclient.Client) (*Account, error) {
	PrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(PrivateKey.PublicKey)
	return &Account{
		PrivateKey: PrivateKey,
		Address:    &address,
		ethclient:  ethclient,
	}, nil
}

func (a *Account) GetAddress() common.Address {
	return *a.Address
}

// get the balance of the account
func (a *Account) GetBalance() (*big.Int, error) {
	return GetBalance(a.ethclient, *a.Address)
}

// BuildTransaction
func (a *Account) BuildTransaction(
	to *common.Address,
	value *big.Int, data []byte,
	nonce, gasLimit uint64,
	chainID, GasFeeCap, GasTipCap *big.Int) (*types.Transaction, error) {
	return BuildTransaction(a.ethclient, a.PrivateKey, *a.Address, to, value, data, nonce, gasLimit, chainID, GasFeeCap, GasTipCap)
}

func (a *Account) SendTransaction(tx *types.Transaction) error {
	return SendTransaction(a.ethclient, tx)
}

func (a *Account) GetTokenBalance(erc20 common.Address) (*big.Int, error) {
	return GetTokenBalance(a.ethclient, erc20, *a.Address)
}

// WaitForTx
func (a *Account) WaitForTx(tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {
	return WaitForTx(a.ethclient, tx, maxWaitSeconds)
}
func (a *Account) Build_Pack_Data(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	value *big.Int, to *common.Address, gasLimit uint64,
	data []byte) (*types.Transaction, error) {
	var log_fields []zap.Field
	log_fields = append(log_fields, zap.Uint64("nonce", nonce))
	log_fields = append(log_fields, zap.Uint64("gasLimit", gasLimit))
	log_fields = append(log_fields, zap.String("to", to.Hex()))
	log_fields = append(log_fields, zap.String("value", value.String()))
	log_fields = append(log_fields, zap.String("gasTipCap", GasTipCap.String()))
	log_fields = append(log_fields, zap.String("chainID", chainID.String()))
	log_fields = append(log_fields, zap.String("baseFee", baseFee.String()))
	defer func() {
		DebugLog("Build_Pack_Data", log_fields...)
	}()
	log_fields = append(log_fields, zap.String("data", hexutil.Encode(data)))
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap)
	log_fields = append(log_fields, zap.String("gasFeeCap", GasFeeCap.String()))

	transaction, err := a.BuildTransaction(to, value, data, nonce, gasLimit, chainID, GasFeeCap, GasTipCap)
	if err != nil {
		return nil, err
	}
	log_fields = append(log_fields, zap.String("transaction", transaction.Hash().Hex()))
	return transaction, nil
}
func (a *Account) Build_Pack(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	value *big.Int, to *common.Address, gasLimit uint64,
	_abi *abi.ABI, name string, args ...interface{}) (*types.Transaction, error) {
	var log_fields []zap.Field
	log_fields = append(log_fields, zap.Uint64("nonce", nonce))
	log_fields = append(log_fields, zap.Uint64("gasLimit", gasLimit))
	log_fields = append(log_fields, zap.String("to", to.Hex()))
	log_fields = append(log_fields, zap.String("value", value.String()))
	log_fields = append(log_fields, zap.String("gasTipCap", GasTipCap.String()))
	log_fields = append(log_fields, zap.String("chainID", chainID.String()))
	log_fields = append(log_fields, zap.String("baseFee", baseFee.String()))
	log_fields = append(log_fields, zap.String("name", name))
	log_fields = append(log_fields, zap.Any("args", args))
	defer func() {
		DebugLog("Build_Pack", log_fields...)
	}()
	data, err := _abi.Pack(name, args...)
	if err != nil {
		return nil, err
	}
	log_fields = append(log_fields, zap.String("data", hexutil.Encode(data)))
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap)
	log_fields = append(log_fields, zap.String("gasFeeCap", GasFeeCap.String()))

	transaction, err := a.BuildTransaction(to, value, data, nonce, gasLimit, chainID, GasFeeCap, GasTipCap)
	if err != nil {
		return nil, err
	}
	log_fields = append(log_fields, zap.String("transaction", transaction.Hash().Hex()))
	return transaction, nil
}

// Build_WETH_deposit
func (a *Account) Build_WETH_deposit(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	amount *big.Int) (*types.Transaction, error) {
	weth := TokenMap[chainID.Uint64()]["WETH"]
	return a.Build_Pack(
		baseFee, nonce, chainID, GasTipCap,
		amount,
		&weth.Address, DefaultWETHDepositGas, WETH_ABI,
		"deposit",
	)
}

func (a *Account) Build_WETH_withdraw(
	baseFee *big.Int, nonce uint64,
	chainID, GasTipCap *big.Int,
	amount *big.Int) (*types.Transaction, error) {
	weth := TokenMap[chainID.Uint64()]["WETH"]
	return a.Build_Pack(
		baseFee, nonce, chainID, GasTipCap,
		big.NewInt(0),
		&weth.Address, DefaultWETHWithdrawGas, WETH_ABI,
		"withdraw", amount,
	)
}
