package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



func SilentError(err error) zap.Field {
    return zap.Field{Key: "error", Type: zapcore.ErrorType, String: err.Error()}
}
