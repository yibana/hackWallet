package hackWallet

import "github.com/ethereum/go-ethereum/common"

type TokenInfo struct {
	ChainId uint64
	Name    string
	Address common.Address
}

type TokenInfoMap map[uint64]map[string]TokenInfo

func (tm TokenInfoMap) AddTokenInfo(chainId uint64, tokenInfo TokenInfo) {
	if tm[chainId] == nil {
		tm[chainId] = make(map[string]TokenInfo)
	}
	tm[chainId][tokenInfo.Name] = tokenInfo
}

func (tm TokenInfoMap) GetTokenInfo(chainId uint64, name string) TokenInfo {
	return tm[chainId][name]
}

var TokenMap = TokenInfoMap{
	1: {
		"WETH": TokenInfo{
			ChainId: 1,
			Name:    "WETH",
			Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		},
	},
	5: {
		"WETH": TokenInfo{
			ChainId: 5,
			Name:    "WETH",
			Address: common.HexToAddress("0x0Bb7509324cE409F7bbC4b701f932eAca9736AB7 "),
		},
	},
}
