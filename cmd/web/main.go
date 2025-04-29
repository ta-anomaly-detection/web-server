package main

import (
	"fmt"
	"github.com/ta-anomaly-detection/web-server/internal/config"
	"go.uber.org/zap"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewEcho(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Start(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal("Failed to start server: %v", zap.Error(err))
	}
}