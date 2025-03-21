package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.SugaredLogger

func InitLogger() {
	// 设置日志输出格式为 JSON
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 设置输出位置（文件写入器）
	writer := getLogWriter()

	// 设置日志等级
	level := zapcore.InfoLevel

	// 创建 core
	core := zapcore.NewCore(encoder, writer, level)

	// 创建 logger
	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	// 自动创建 logs 目录（如果不存在）
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", 0755)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
