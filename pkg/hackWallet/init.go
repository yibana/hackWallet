package hackWallet

import (
	"github.com/yibana/hackWallet/configs"
)

var ErrLog = configs.LOGGER.Error
var InfoLog = configs.LOGGER.Info
var DebugLog = configs.LOGGER.Debug
var WarnLog = configs.LOGGER.Warn
var FatalLog = configs.LOGGER.Fatal
