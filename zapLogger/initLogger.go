package zapLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Configuration struct {
	Log
	LogPath 	string
}
type Log struct {
	Path 	string
}

func (c *Configuration)InitLogger(level string) *zap.Logger  {
	hook := lumberjack.Logger{
		Filename: 	c.LogPath,
		MaxSize: 	1024,	//megabytes
		MaxBackups: 3,
		MaxAge: 	7,		//days
		Compress: 	true,
	}

	w := zapcore.AddSync(&hook)

	var l zapcore.Level
	switch level {
	case "debug":
		l = zap.DebugLevel
	case "info":
		l = zap.InfoLevel
	case "error":
		l = zap.ErrorLevel
	default:
		l = zap.InfoLevel
	}

	//encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		// StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, l)

	logger := zap.New(core)
	logger.Info("Default logger init SUCCESS!")

	return logger
}

func (c *Configuration)InitLogger2(level string) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: 	c.LogPath,
		MaxSize: 	500,
		MaxBackups: 3,
		MaxAge: 	28,
	})

	encoderConf := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf), w, zap.DebugLevel)
	logger := zap.New(core)

	return logger
}
