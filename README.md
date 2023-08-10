# HackWallet
模拟以太坊钱包(无界面),方便科学家开发套利程序

### `contracts` 目录
将合约`xxx.abi`文件放到这个目录下的`abis`子目录
* 运行`scripts/abigen/abigen_generate_go.go` 即可自动生成合约接口文件

### `configs` 目录
* 自动加载`.evn` 文件到环境变量,并读取配置
* 会创建一个自动分割的日志对象,默认保存到`logs`目录

### `pkg` 目录
* `flashbot`
  * [https://github.com/cryptoriums/flashbot/blob/main/flashbot.go](https://github.com/cryptoriums/flashbot/blob/main/flashbot.go)
* `hackWallet`
  * 核心包

### `examples` 目录
* [BatchTransaction](examples/BatchTransaction/BatchTransaction.go) 批量执行交易
* [CallBundle](examples/CallBundle/CallBundle.go) 打包交易(flashbot)
* [Transfer](examples/Transfer/Transfer.go) 发送交易

### AnvilFork 选项
[anvil](https://github.com/foundry-rs/foundry) 运行本地分叉节点，这样就可以创建主网的分叉环境然后模拟交易,anvil会创建10个有10000ETH的虚拟钱包地址
```go
AnvilFork := true
Wallet, _ = hackWallet.NewHackWallet(configs.HTTP_RPC_URL, AnvilFork)
for i, account := range Wallet.Accounts {
	balance, _ := account.GetBalance()
	fmt.Printf("[%d]acc:%s balance:%f\n", 
		i+1, account.Address.String(), 
		hackWallet.ConvertWei2Eth(balance))
}
```
```shell
[1]acc:0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 balance:10000.000000
[2]acc:0x70997970C51812dc3A010C7d01b50e0d17dc79C8 balance:10000.000000
...
[10]acc:0xa0Ee7A142d267C1f36714E4a8F75612F20a79720 balance:10000.000000
```
