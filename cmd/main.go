package main

import (
	"medtest/config"
	"medtest/internal/app"
	"medtest/pkg/logger"
)

func main() {
	Logger := logger.NewLogger()
	Config := config.NewConfig()
	if Config == nil {
		Logger.Info("Bad config")
	}
	App, err := app.NewApp(Logger, *Config)
	if err != nil {
		Logger.Debug(err)
		return
	}
	App.Run()
}
