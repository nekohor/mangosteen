package logger

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func InitLogger(level string, filePath string) {

	//// 打印错误级别的日志
	//highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl >= zapcore.ErrorLevel
	//})
	//// 打印所有级别的日志
	//lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl >= zapcore.DebugLevel
	//})

	var allCore []zapcore.Core

	hook := lumberjack.Logger{
		Filename: filePath,
		MaxSize:  20,
		MaxAge:   0,
		Compress: false,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var writers = []zapcore.WriteSyncer{}
	writers = append(writers, os.Stdout)
	writers = append(writers, zapcore.AddSync(&hook))
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getLevel(level))
	writerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		atomicLevel,
	)
	allCore = append(allCore, writerCore)

	//// High-priority output should also go to standard error, and low-priority
	//// output should also go to standard out.
	//consoleDebugging := zapcore.Lock(os.Stdout)
	//
	//// for human operators.
	//consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//
	//if level == "debug" {
	//	allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	//}

	core := zapcore.NewTee(allCore...)
	caller := zap.AddCaller()
	development := zap.Development()
	logger = zap.New(core, caller, development, zap.AddCallerSkip(1)).Sugar()
}

func getLevel(level string) (l zapcore.Level) {
	switch level {
	case "debug":
		l = zap.DebugLevel
	case "info":
		l = zap.InfoLevel
	case "warn":
		l = zap.WarnLevel
	case "error":
		l = zap.ErrorLevel
	case "panic":
		l = zap.PanicLevel
	case "fatal":
		l = zap.FatalLevel
	default:
		l = zap.InfoLevel
	}
	return
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Info(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
