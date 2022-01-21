package main

import (
	"fmt"
	"github.com/natefinch/lumberjack" //用于日志切割
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger

func initLogger(logpath string, loglevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    1,    //多少MB切割
		MaxBackups: 30,   //备份数量
		MaxAge:     7,    // 保存几天
		Compress:   true, //是否压缩
	}
	w := zapcore.AddSync(&hook)
	// 设置日志级别
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level)
	logger := zap.New(core)
	logger.Info("DefaultLogger init success")
	return logger
}

func main() {
	ZapLogger = initLogger("D://all.log", "debug")
	for i := 0; i < 1000000; i++ {
		ZapLogger.Debug(fmt.Sprint("test log ", i))
	}
}
