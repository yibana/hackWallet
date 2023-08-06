package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger(logDir string, fn string, LogLevel string) *zap.Logger {
	// Create log directory
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// Define log level
	logLevel := zap.NewAtomicLevel()
	level, err := zapcore.ParseLevel(LogLevel)
	if err != nil {
		logLevel.SetLevel(zapcore.InfoLevel)
	} else {
		logLevel.SetLevel(level)
	}

	// Define log encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
			Filename:   logDir + "/" + fn + ".log",
			MaxSize:    10, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			LocalTime:  true,
		})),
		logLevel,
	)
	return zap.New(core, zap.AddCaller())
}
