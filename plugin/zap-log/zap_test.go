package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestZapPrint(t *testing.T) {
	logDev, _ := zap.NewDevelopment()
	logDev.Debug("this is debug message")
	logDev.Info("this is info message")
	logDev.Info("this is info message with fileds", zap.Int("age", 24), zap.String("agender", "man"))
	logDev.Warn("this is warn message")
	logDev.Error("this is error message")

	logPro, _ := zap.NewProduction()
	logPro.Debug("this is debug message")
	logPro.Info("this is info message")
	logPro.Info("this is info message with fileds", zap.Int("age", 24), zap.String("agender", "man"))
	logPro.Warn("this is warn message")
	logPro.Error("this is error message")
}

func TestZapToFile(t *testing.T) {
	// 日志配置
	encoderConfig  := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder 	// 在日志文件中使用大写字母记录日志级别
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	file , _ := os.Create("D://zap.log")
	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	sugarLogger.Info("this is info message")
	sugarLogger.Infof("this is %s, %d", "aaa", 1234)
	sugarLogger.Error("this is error message")
	sugarLogger.Info("this is info message")
}
