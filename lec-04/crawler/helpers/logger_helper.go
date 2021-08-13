package helpers

import "go.uber.org/zap"

func GetSugar() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	return logger.Sugar()
}
