package hackWallet

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/yibana/hackWallet/pkg/flashbot"
	"go.uber.org/zap"
	"math/big"
	"sync"
	"time"
)

type HackWallet struct {
	Accounts             []*Account
	ProviderURL          string
	RPCClient            *ethclient.Client
	rpc                  *rpc.Client
	lastBlockHeader      *types.Header
	lastBlockHeader_time uint64
	chainID              *big.Int
	defaultGasTipCap     *big.Int
	fbs                  []flashbot.Flashboter
}

func NewHackWallet(rpcUrl string, AnvilFork bool) (*HackWallet, error) {
	var accounts []*Account
	var err error
	var privateKeys []string
	if AnvilFork {
		privateKeys, err = anvil_fork(rpcUrl)
		if err != nil {
			return nil, err
		}
		rpcUrl = "http://127.0.0.1:8545"
	}
	rpc, err := rpc.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(rpc)
	for _, privateKey := range privateKeys {
		if acc, err := NewAccount(privateKey, client); err == nil {
			accounts = append(accounts, acc)
		}
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &HackWallet{
		Accounts:         accounts,
		RPCClient:        client,
		ProviderURL:      rpcUrl,
		chainID:          chainID,
		defaultGasTipCap: big.NewInt(DefaultGasTipCap),
		rpc:              rpc,
	}, nil
}

// init default flashbot
func (hw *HackWallet) InitFlashbot(prvKey *ecdsa.PrivateKey, additional bool) error {
	rpcList := []string{"https://rpc.lightspeedbuilder.info",
		"https://rpc.beaverbuild.org", "https://builder0x69.io", "https://eth-builder.com",
		"https://buildai.net", "https://builder.gmbit.co/rpc", "https://rsync-builder.xyz",
		"https://rpc.lightspeedbuilder.info", "https://rpc.payload.de",
		"https://api.blocknative.com/v1/auction", "https://rpc.titanbuilder.xyz", "https://rpc.nfactorial.xyz"}

	var additionals []*flashbot.Api
	if additional {
		for _, s := range rpcList {
			additionals = append(additionals, &flashbot.Api{
				URL:                s,
				SupportsSimulation: false,
			})
		}
	}
	if len(hw.fbs) == 0 {
		flashboters, err := flashbot.NewAll(hw.chainID.Int64(), prvKey, additionals...)
		if err != nil {
			return err
		}
		hw.fbs = flashboters
	}
	return nil
}

// set defaultGasTipCap
func (hw *HackWallet) SetDefaultGasTipCap(gasTipCap int64) {
	hw.defaultGasTipCap = big.NewInt(gasTipCap)
}

func (hw *HackWallet) AddAccount(account *Account) {
	hw.Accounts = append(hw.Accounts, account)
}

// load account
func (hw *HackWallet) LoadAccount(privateKey string) (*Account, error) {
	account, err := NewAccount(privateKey, hw.RPCClient)
	if err != nil {
		return nil, err
	}
	hw.AddAccount(account)
	return account, nil
}

// add rand account
func (hw *HackWallet) AddRandAccount() (*Account, error) {
	account, err := RandAccount(hw.RPCClient)
	if err != nil {
		return nil, err
	}
	hw.AddAccount(account)
	return account, nil
}

// get blocknumber
func (hw *HackWallet) GetBlockNumber() (uint64, error) {
	header, err := hw.GetBlockHeader(nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

// get block header
func (hw *HackWallet) GetBlockHeader(blockNumber *big.Int) (*types.Header, error) {
	var err error
	nowTime := uint64(time.Now().Unix())
	if hw.lastBlockHeader != nil {
		if nowTime > hw.lastBlockHeader.Time && nowTime <= hw.lastBlockHeader.Time+11 {
			return hw.lastBlockHeader, nil
		}
		if hw.lastBlockHeader_time == nowTime { // 1秒保护
			return hw.lastBlockHeader, nil
		}
	}
	hw.lastBlockHeader, err = hw.RPCClient.HeaderByNumber(context.Background(), blockNumber)
	hw.lastBlockHeader_time = nowTime
	if err != nil {
		return nil, err
	}
	var log_fields []zap.Field
	log_fields = append(log_fields, zap.String("blockNumber", hw.lastBlockHeader.Number.String()))
	log_fields = append(log_fields, zap.Uint64("blockTime", hw.lastBlockHeader.Time))
	log_fields = append(log_fields, zap.String("BaseFee", ConvertToGwei(hw.lastBlockHeader.BaseFee).String()+"(Gwei)"))
	defer func() {
		DebugLog("GetBlockHeader", log_fields...)
	}()
	return hw.lastBlockHeader, nil
}

// GetNextBlockBaseFee
func (hw *HackWallet) GetNextBlockBaseFee() (*big.Int, error) {
	header, err := hw.GetBlockHeader(nil)
	if err != nil {
		return nil, err
	}
	return GetNextBlockBaseFee(header)
}

func (hw *HackWallet) GetChainID() *big.Int {
	return hw.chainID
}

// select  account
func (hw *HackWallet) SelectAccount(address common.Address) (*Account, error) {
	for _, account := range hw.Accounts {
		if account.GetAddress() == address {
			return account, nil
		}
	}
	return nil, errors.New("account not found in wallet")
}

type BuildBack func(from *Account, baseFee *big.Int, nonce uint64) (*types.Transaction, error)
type BuildTxnBack func(txb *TxBaseBuild) (*types.Transaction, error)

// BuildBatchTransaction
func (hw *HackWallet) BuildBatchTransaction(fromAcc *Account, bc ...BuildBack) ([]*types.Transaction, error) {
	var transactions []*types.Transaction
	var err error
	var nonce uint64
	var baseFee *big.Int

	baseFee, err = hw.GetNextBlockBaseFee()
	if err != nil {
		return nil, err
	}
	nonce, err = fromAcc.ethclient.PendingNonceAt(context.Background(), *fromAcc.Address)
	if err != nil {
		return nil, err
	}
	for _, b := range bc {
		transaction, err := b(fromAcc, big.NewInt(0).Set(baseFee), nonce)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
		nonce++
	}
	return transactions, nil
}

func (hw *HackWallet) BuildBatchTxn(fromAcc *Account, gasTipCap *big.Int, bc ...BuildTxnBack) ([]*types.Transaction, error) {
	var transactions []*types.Transaction
	var err error
	var nonce uint64
	var baseFee *big.Int
	chain := hw.chainID
	baseFee, err = hw.GetNextBlockBaseFee()
	if err != nil {
		return nil, err
	}
	nonce, err = fromAcc.ethclient.PendingNonceAt(context.Background(), *fromAcc.Address)
	if err != nil {
		return nil, err
	}
	for _, b := range bc {
		transaction, err := b(&TxBaseBuild{
			From:      fromAcc,
			BaseFee:   big.NewInt(0).Set(baseFee),
			Nonce:     nonce,
			ChainID:   chain,
			GasTipCap: big.NewInt(0).Set(gasTipCap),
		})
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
		nonce++
	}
	return transactions, nil
}

func (hw *HackWallet) SendBatchTransaction(transactions []*types.Transaction) error {
	for _, transaction := range transactions {
		if err := SendTransaction(hw.RPCClient, transaction); err != nil {
			return err
		}
	}
	return nil
}

func (hw *HackWallet) SendBundleFrom(transactions []*types.Transaction, blockNum uint64, attempts int64) error {
	if len(hw.fbs) == 0 {
		return errors.New("flashbot not init")
	}
	if attempts <= 0 {
		attempts = 1
	}
	var err error
	if blockNum <= 0 {
		blockNum, err = hw.GetBlockNumber()
		if err != nil {
			return err
		}
	}

	txs := []string{}
	for _, transaction := range transactions {
		data, err := transaction.MarshalBinary()
		if err != nil {
			return err
		}
		txs = append(txs, hexutil.Encode(data))
	}
	for i := int64(0); i < attempts; i++ {
		blockNum++
		if err := hw.SendBundle(txs, blockNum); err != nil {
			return err
		}
	}
	return nil
}

func (hw *HackWallet) SendBundle(txs []string, blockNum uint64) error {
	if len(hw.fbs) == 0 {
		return errors.New("flashbot not init")
	}
	ctx := context.Background()
	var err error
	var succ bool
	if blockNum <= 0 {
		blockNum, err = hw.GetBlockNumber()
		if err != nil {
			return err
		}
	}
	wait := sync.WaitGroup{}
	for _, fb := range hw.fbs {
		wait.Add(1)
		go func(gofb flashbot.Flashboter) {
			defer wait.Done()
			response, err := gofb.SendBundle(ctx, txs, blockNum)
			if err != nil {
				ErrLog("SendBundle", zap.Uint64("blockNum", blockNum),
					zap.String("rpc", gofb.Api().URL), zap.Error(err))
			} else {
				if response.Error.Code != 0 {
					ErrLog("SendBundle", zap.Uint64("blockNum", blockNum),
						zap.String("rpc", gofb.Api().URL), zap.Any("Error", response.Error))
				} else {
					InfoLog("SendBundle", zap.Uint64("blockNum", blockNum),
						zap.String("rpc", gofb.Api().URL), zap.Any("Result", response.Result))
					succ = true
				}
			}
		}(fb)
	}
	wait.Wait()

	if !succ {
		return errors.New("send bundle failed")
	}
	return nil
}

func (hw *HackWallet) CallBundle(transactions []*types.Transaction, blockNumState uint64) (*flashbot.Response, error) {
	if len(hw.fbs) == 0 {
		return nil, errors.New("flashbot not init")
	}

	txs := []string{}
	for _, transaction := range transactions {
		data, err := transaction.MarshalBinary()
		if err != nil {
			return nil, err
		}
		txs = append(txs, hexutil.Encode(data))
	}
	ctx := context.Background()
	var err error
	var log_fields []zap.Field
	var response *flashbot.Response
	fb := hw.fbs[0]

	log_fields = append(log_fields, zap.Uint64("blockNum", blockNumState))
	log_fields = append(log_fields, zap.String("rpc", fb.Api().URL))

	response, err = fb.CallBundle(ctx, txs, blockNumState)
	if err != nil {
		log_fields = append(log_fields, zap.Error(err))
		ErrLog("CallBundle", log_fields...)
		return nil, err
	} else {
		if response.Error.Code != 0 {
			log_fields = append(log_fields, zap.Any("Error", response.Error))
			ErrLog("CallBundle", log_fields...)
		} else {
			log_fields = append(log_fields, zap.Any("Result", response.Result))
			InfoLog("CallBundle", log_fields...)
		}
	}
	err = response.HasError()
	return response, err
}

// BuildTransaction From,to,value,gas,gasPrice,nonce
func (hw *HackWallet) BuildTransaction(
	fromAcc *Account, to *common.Address, data []byte,
	value *big.Int, gasLimit, nonce uint64,
	GasTipCap *big.Int, // 小费
) (*types.Transaction, error) {

	baseFee, err := hw.GetNextBlockBaseFee()
	if err != nil {
		return nil, err
	}
	GasFeeCap := big.NewInt(0).Add(baseFee, GasTipCap) // 计算最大支出手续费（矿工费）
	return fromAcc.BuildTransaction(to, value, data, nonce, gasLimit, hw.chainID,
		GasFeeCap, GasTipCap)
}

func (hw *HackWallet) GetTokenBalance(fromAcc *Account, TokenAddress common.Address) (*big.Int, error) {
	return GetTokenBalance(hw.RPCClient, fromAcc.GetAddress(), TokenAddress)
}

// WaitForTx
func (hw *HackWallet) WaitForTx(tx *types.Transaction, maxWaitSeconds uint) (*types.Receipt, error) {
	return WaitForTx(hw.RPCClient, tx, maxWaitSeconds)
}

// BuildTransfer
func (hw *HackWallet) BuildTransfer(fromAcc *Account, to *common.Address, value, gasTipCap *big.Int, nonce uint64) (*types.Transaction, error) {
	return hw.BuildTransaction(fromAcc, to, nil, value, 21000, nonce, gasTipCap)
}

// Transfer
func (hw *HackWallet) Transfer(fromAcc *Account, to *common.Address, value *big.Int) (*types.Transaction, error) {
	transaction, err := hw.BuildTransfer(fromAcc, to, value, hw.defaultGasTipCap, 0)
	if err != nil {
		return nil, err
	}
	return transaction, fromAcc.SendTransaction(transaction)
}

// SimulationActionFrom
func (hw *HackWallet) SimulationActionFrom(from, to common.Address, value *big.Int, data []byte) (gas uint64, err error) {
	return SimulationActionFrom(hw.RPCClient, from, to, value, data)
}
