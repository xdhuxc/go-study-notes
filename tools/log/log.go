package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	Debug         bool   `yaml:"debug"`
	App           string `yaml:"app"`
	DataLogPrefix string `yaml:"data_log_prefix"`
	Group         string `yaml:"group"`
	MaxSize       int    `yaml:"max_size"`
	MaxBackups    int    `yaml:"max_backups"`
	MaxAge        int    `yaml:"max_age"`
	Compress      bool   `yaml:"compress"`
	Stdout        bool   `yaml:"stdout"`
}

var (
	Logger *zap.SugaredLogger

	defaultApp           = "scmp"
	defaultGroup         = "sgt"
	defaultDataLogPrefix = "/data/logs/"
)

func InitLogger(cfg *Config) {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoWriter := getLogWriter("info", cfg)
	errorWriter := getLogWriter("error", cfg)

	encoder := getEncoder()

	var core zapcore.Core

	if cfg.Debug {
		debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		})
		debugWriter := getLogWriter("debug", cfg)

		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		)
	}

	sugarLogger := zap.New(core, zap.AddCaller())

	Logger = sugarLogger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(level string, cfg *Config) zapcore.WriteSyncer {
	if !cfg.Stdout {
		lumberJackLogger := &lumberjack.Logger{
			MaxSize:    100,          // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 10,           // 日志文件最多保存多少个备份
			MaxAge:     2,            // 文件最多保存多少天
			Compress:   cfg.Compress, // 是否压缩
		}
		// 此处适配奇葩metis 要求,全部收集到es
		if level == "error" {
			level = "common-error"
		} else {
			level = "digest_info"
		}

		fileName := defaultDataLogPrefix
		if cfg.DataLogPrefix != "" {
			fileName = cfg.DataLogPrefix
		} else {
			fileName = defaultDataLogPrefix
		}
		if cfg.Group != "" {
			fileName = fileName + cfg.Group + "_"
		} else {
			fileName = fileName + defaultGroup + "_"
		}
		if cfg.App != "" {
			fileName = fileName + cfg.App + "_"
		} else {
			fileName = fileName + defaultApp + "_"
		}
		fileName = fileName + level + ".log"

		lumberJackLogger.Filename = fileName

		if cfg.MaxSize > 0 {
			lumberJackLogger.MaxSize = cfg.MaxSize
		}
		if cfg.MaxBackups > 0 {
			lumberJackLogger.MaxBackups = cfg.MaxBackups
		}
		if cfg.MaxAge > 0 {
			lumberJackLogger.MaxAge = cfg.MaxAge
		}

		return zapcore.AddSync(lumberJackLogger)
	}

	return zapcore.AddSync(os.Stdout)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	Logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	Logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
