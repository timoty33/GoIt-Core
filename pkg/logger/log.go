package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logEntry struct {
	level  zapcore.Level
	msg    string
	fields []zap.Field
}

// Log adiciona uma entrada no canal assíncrono
func Log(level zapcore.Level, msg string, fields ...zap.Field) {
	select {
	case logChan <- logEntry{level: level, msg: msg, fields: fields}:
	default:
		// se o canal estiver cheio, descartamos para não travar
	}
}
