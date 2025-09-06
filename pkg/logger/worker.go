package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
	wg        sync.WaitGroup
)

// Create inicializa o logger e inicia o worker assíncrono
func Create() {
	var err error

	cfg := zap.NewProductionConfig()

	zapLogger, err = cfg.Build()
	if err != nil {
		panic(fmt.Errorf("erro ao criar zap logger: %w", err))
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for entry := range logChan {
			switch entry.level {

			case zapcore.DebugLevel:
				zapLogger.Debug(entry.msg, entry.fields...)

			case zapcore.InfoLevel:
				zapLogger.Info(entry.msg, entry.fields...)

			case zapcore.WarnLevel:
				zapLogger.Warn(entry.msg, entry.fields...)

			case zap.ErrorLevel:
				zapLogger.Error(entry.msg, entry.fields...)

			case zapcore.PanicLevel:
				zapLogger.Panic(entry.msg, entry.fields...)

			case zapcore.FatalLevel:
				zapLogger.Fatal(entry.msg, entry.fields...)

			default:
				zapLogger.Info(entry.msg, entry.fields...)
			}
		}
	}()
}

// Sync garante segurança dos logs antes de fechar
func Sync() {
	if zapLogger != nil {
		_ = zapLogger.Sync()
	}

	close(logChan) // fecha o canal antes de esperar
	wg.Wait()
}
