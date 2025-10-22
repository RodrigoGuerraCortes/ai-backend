package logger

import (
	"github.com/RodrigoGuerraCortes/ai-backend/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {

	config.LoadEnv()
	env := config.GetEnv("ENV")
	logLevel := config.GetEnv("LOG_LEVEL")

	// Default level = info
	level := zapcore.InfoLevel
	if logLevel != "" {
		_ = level.Set(logLevel)
	}

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
