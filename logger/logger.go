package logger

import (
	"gin-learn-notes/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var Log *zap.SugaredLogger

func InitLogger() {
	logConf := config.Conf.Logger

	// 设置日志输出格式为 JSON
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 设置输出位置（文件写入器）
	writeSyncer := getLogWriter(logConf.File)

	// 设置日志等级
	level := getLogLevel(logConf.Level)

	// 创建 core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 创建 logger
	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}

// 日志等级转换
func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getLogWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	// 同时写入日志文件 + 控制台（开发环境）
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberJackLogger),
		zapcore.AddSync(os.Stdout),
	)
}
