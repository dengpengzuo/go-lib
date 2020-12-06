package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

type Level int8

const (
    FILE_MAX_SIZE = 512 * 1024 * 1024

    DebugLevel Level = iota - 1
    InfoLevel
    WarnLevel
    ErrorLevel
)

var logger *zap.SugaredLogger
var log_level = zapcore.InfoLevel

// 输出文件和控制台
func InitLogger(file string, level Level) {
    logger, log_level = newZapLogger(true, file, FILE_MAX_SIZE, level)
}

// 只输出文件
func InitFileLogger(file string, level Level) {
    logger, log_level = newZapLogger(false, file, FILE_MAX_SIZE, level)
}

// 只输出控制台
func InitConsoleLogger(level Level) {
    logger, log_level = newZapLogger(true, "", 0, level)
}

func newZapLogger(console bool, logpath string, logfileMax int32, loglevel Level) (*zap.SugaredLogger, zapcore.Level) {
    // 初始是必须写stdout
    var w zapcore.WriteSyncer
    if logpath != "" {
        // lumberjack 内部有锁
        hook := lumberjack.Logger{
            Filename:   logpath,         // ⽇志⽂件路径
            MaxSize:    int(logfileMax), // 1G
            MaxBackups: 10,              // 最多保留10个备份
            MaxAge:     7,               // days
            Compress:   true,            // 是否压缩
        }
        fw := zapcore.AddSync(&hook)
        if console {
            w = zapcore.NewMultiWriteSyncer(os.Stdout, fw)
        } else {
            w = fw
        }
    } else if console {
        w = zapcore.AddSync(os.Stdout)
    } else {
        w = zapcore.AddSync(os.Stderr)
    }
    var zaplevel zapcore.Level
    switch loglevel {
    case InfoLevel:
        zaplevel = zapcore.InfoLevel
    case WarnLevel:
        zaplevel = zapcore.WarnLevel
    case ErrorLevel:
        zaplevel = zapcore.ErrorLevel
    default:
        zaplevel = zapcore.DebugLevel
    }

    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoderConfig),
        w,
        zaplevel,
    )
    // 显示调用者和代码行数
    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    return logger.Sugar(), zaplevel
}

func Debugf(fmt string, args ...interface{}) {
    if log_level.Enabled(zapcore.DebugLevel) {
        logger.Debugf(fmt, args)
    }
}

func Infof(fmt string, args ...interface{}) {
    if log_level.Enabled(zapcore.InfoLevel) {
        logger.Infof(fmt, args)
    }
}

func Warnf(fmt string, args ...interface{}) {
    if log_level.Enabled(zapcore.WarnLevel) {
        logger.Warnf(fmt, args)
    }
}

func Errorf(fmt string, args ...interface{}) {
    if log_level.Enabled(zapcore.ErrorLevel) {
        logger.Errorf(fmt, args)
    }
}
