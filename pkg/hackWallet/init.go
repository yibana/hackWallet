package hackWallet

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yibana/hackWallet/configs"
)

var ErrLog = configs.LOGGER.Error
var InfoLog = configs.LOGGER.Info
var DebugLog = configs.LOGGER.Debug
var WarnLog = configs.LOGGER.Warn
var FatalLog = configs.LOGGER.Fatal

var WETH9 = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
