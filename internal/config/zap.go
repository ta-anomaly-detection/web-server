package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	return logger
}
