package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func SetLogger(l *zap.Logger) {
	logger = l
	logger.Named("history")
}

func LogInOtherFile() {
	logger.Info("log in demo", zap.String("user", "cc"))
}
