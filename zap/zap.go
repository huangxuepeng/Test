package main

import (
	"time"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./logger.log",
		"stderr",
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	su := logger.Sugar()
	defer su.Sync()
	su.Info("shibai",
		zap.String("url", "jjj"),
		zap.Int("attemp", 3),
		zap.Duration("backoff", time.Second),
	)
}
