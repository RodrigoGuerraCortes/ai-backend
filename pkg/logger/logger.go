package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitWith(env, levelStr string) {
	level := zapcore.InfoLevel
	_ = level.Set(levelStr)

	var cfg zap.Config

	if env == "development" {
		//  Pretty console logs for local debugging
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Level = zap.NewAtomicLevelAt(level)
	} else {
		//  JSON structured logs for production
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(level)
	}

	var err error
	Log, err = cfg.Build()
	if err != nil {
		panic("❌ Failed to initialize zap logger: " + err.Error())
	}

	Log.Info("Logger initialized",
		zap.String("env", env),
		zap.String("level", level.String()))
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
