package helpers

import "go.uber.org/zap"

func GetSugar() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()

	return logger.Sugar()
}
