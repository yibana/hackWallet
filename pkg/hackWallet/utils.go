package hackWallet

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

func anvil_fork(rpc_url string) ([]string, error) {
	// 要执行的命令及其参数
	cmd := exec.Command("anvil", "--fork-url", rpc_url)
	var privateKey []string
	// 创建管道，捕获标准输出
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		ErrLog(fmt.Sprintf("无法创建标准输出管道：%s\n", err))
		return nil, err
	}

	// 开始命令
	err = cmd.Start()
	if err != nil {
		ErrLog(fmt.Sprintf("无法启动命令：%s\n", err))
		return nil, err
	}

	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		defer wait.Done()
		// 读取持续输出并处理
		scanner := bufio.NewScanner(stdoutPipe)
		start_extractPrivateKey := false
		for scanner.Scan() {
			outputLine := scanner.Text()
			fmt.Println(outputLine)
			if strings.Contains(outputLine, "Listening on 127.0.0.1:8545") {
				wait.Done()
			}
			if strings.Contains(outputLine, "Private Keys") {
				start_extractPrivateKey = true
			}
			if start_extractPrivateKey {
				if strings.HasPrefix(outputLine, "Wallet") {
					start_extractPrivateKey = false
				} else {
					match := extractPrivateKey(outputLine)
					if match != "" {
						privateKey = append(privateKey, match)
					}
				}
			}
		}

		// 检查是否有错误发生
		err = cmd.Wait()
		if err != nil {
			ErrLog(fmt.Sprintf("命令执行失败：%s\n", err))
		}
	}()
	wait.Wait()

	return privateKey, err
}

func extractPrivateKey(input string) string {
	// 定义一个匹配私钥的正则表达式
	re := regexp.MustCompile(`0x[0-9a-fA-F]{64}`)

	// 使用正则表达式查找匹配的字符串
	match := re.FindString(input)

	return match
}

func BigIntToDecimal_18(bigInt *big.Int) uint64 {
	var result uint64
	big_decimal := big.NewInt(0).Mul(big.NewInt(1e10), big.NewInt(1e8))
	result = bigInt.Div(bigInt, big_decimal).Uint64()
	return result
}

// 准确计算出下一个区块的basefee值
func GetNextBlockBaseFee(header *types.Header) (*big.Int, error) {
	// 获取最新区块

	//On the off chance that the last block was precisely half full, the Base Charge will stay unaltered.
	//	On the off chance that the last block was 100 percent full, the Base Charge will increment by the greatest 12.5% for the following block.
	//	On the off chance that the last block was over half full yet under 100 percent full, the Base Expense will increment by under 12.5%.
	//	Assuming the last block was 0% full - that is, vacant - the Base expense will diminish the most extreme 12.5% for the following block.
	//	In the event that the last block was over 0% full yet under half full, the Base Charge will diminish by under 12.5%
	// 计算当前区块的gas使用率
	currentBlockGasLimit := header.GasLimit
	currentBlockGasUsed := header.GasUsed
	currentBlockGasUsage := new(big.Int).Div(new(big.Int).Mul(big.NewInt(100*1000), big.NewInt(int64(currentBlockGasUsed))), big.NewInt(int64(currentBlockGasLimit)))

	// 计算下一个区块的basefee
	currentBaseFee := new(big.Int).SetBytes(header.BaseFee.Bytes())
	nextBaseFee := new(big.Int)
	if currentBlockGasUsage.Cmp(big.NewInt(50*1000)) < 0 { // < 50%
		p := new(big.Int).Sub(big.NewInt(50*1000), currentBlockGasUsage) //计算下降了多少
		p = new(big.Int).Mul(p, big.NewInt(100))
		p = new(big.Int).Div(p, big.NewInt(50)) //下降率
		//fmt.Println(float64(p.Int64()) / 1000)
		p = new(big.Int).Mul(big.NewInt(125), p)
		//fmt.Println(float64(p.Int64()) / 1000000)

		nextBaseFee = new(big.Int).Mul(currentBaseFee, big.NewInt(0).Sub(big.NewInt(1000*1000*100), p))
		nextBaseFee = new(big.Int).Div(nextBaseFee, big.NewInt(1000*1000*100))

	} else {
		p := new(big.Int).Sub(currentBlockGasUsage, big.NewInt(50*1000)) //计算上升了多少
		p = new(big.Int).Mul(p, big.NewInt(100))
		p = new(big.Int).Div(p, big.NewInt(50)) //上升率
		//fmt.Println(float64(p.Int64()) / 1000)
		p = new(big.Int).Mul(big.NewInt(125), p)
		//fmt.Println(float64(p.Int64()) / 1000000)
		nextBaseFee = new(big.Int).Mul(currentBaseFee, big.NewInt(0).Add(big.NewInt(1000*1000*100), p))
		nextBaseFee = new(big.Int).Div(nextBaseFee, big.NewInt(1000*1000*100))
	}

	nextBaseFee = big.NewInt(0).Add(nextBaseFee, big.NewInt(1e8)) // 加0.1gwei 防止不够
	return nextBaseFee, nil
}

func ToGwei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(1000000000))
}
func ToGweiFloat(wei *big.Int) float64 {
	return float64(new(big.Int).Div(wei, big.NewInt(10000)).Int64()) / 100000
}
