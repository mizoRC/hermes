package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func Init() {
	var config zap.Config
	if os.Getenv("LOG_LEVEL") == "debug" {
		config = zap.NewDevelopmentConfig()
		// config = zap.NewProductionConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	/*logger.Info("Hello!")
	  logger.Warn("Hello!")
	  logger.Info("failed to fetch URL",
	  // Structured context as loosely typed key-value pairs.
	  zap.String("url", "www.test.org"),
	  zap.Int("attempt", 3),
	  zap.Duration("backoff", time.Second),
	  )
	  logger.Error("Hello!")*/

	Logger = logger
}
