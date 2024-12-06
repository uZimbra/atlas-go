package config

import "go.uber.org/zap"

var Zap *zap.SugaredLogger

func ConfigureLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	Zap = logger.Sugar()
}
