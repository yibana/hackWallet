package configs

import (
	"github.com/joho/godotenv"
	"github.com/yibana/hackWallet/internal/logger"
	"os"
	"path"
)

var (
	ROOT = ""
	// load .env file first
	_ = loadEnv()

	PROXY        = os.Getenv("PROXY")
	HTTP_RPC_URL = os.Getenv("HTTP_RPC_URL")
	WSS_RPC_URL  = os.Getenv("WSS_RPC_URL")

	LOG_DIR   = GetEnvWithDefault("LOG_DIR", "./logs")
	LOG_FILE  = GetEnvWithDefault("LOG_FILE", "default")
	LOG_LEVEL = GetEnvWithDefault("LOG_LEVEL", "debug")

	LOGGER = logger.NewLogger(path.Join(ROOT, LOG_DIR), LOG_FILE, LOG_LEVEL)
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadEnv() error {
	var errs []error
	for _, f := range []string{"", "../", "../../"} {
		err := godotenv.Load(path.Join(f, ".env"))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		ROOT = f
		return nil
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
