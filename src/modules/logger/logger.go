package logger

import (
	"encoding/json"
	"github.com/hramov/tg-bot-admin/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var Instance *zap.SugaredLogger

func New() error {
	var cfg zap.Config
	if err := json.Unmarshal(config.LoggerConfig, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.OutputPaths = []string{"data/logs/logs.log", "stdout"}

	logger := zap.Must(cfg.Build())
	Instance = logger.Sugar()
	return nil
}

func NewTest() error {
	var cfg zap.Config
	if err := json.Unmarshal(config.LoggerConfig, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.OutputPaths = []string{"stdout"}
	logger := zap.Must(cfg.Build())
	Instance = logger.Sugar()
	return nil
}

func Destroy() {
	if err := Instance.Sync(); err != nil {
		log.Printf("cannot sync logger: %v", err.Error())
	}
	Instance = nil
}
