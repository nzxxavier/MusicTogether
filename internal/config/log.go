package config

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	logursLogger "github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type MusicTogetherLogger struct {
	logger *logursLogger.Logger
}

var log = &MusicTogetherLogger{
	logger: logursLogger.New(),
}

func init() {
	output := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   viper.GetString(LogFileName),
		MaxSize:    viper.GetInt(LogFileMaxSizeMB),
		MaxAge:     viper.GetInt(LogFileMaxAgeDays),
		MaxBackups: viper.GetInt(LogFileMaxBackups),
		LocalTime:  true,
		Compress:   viper.GetBool(LogCompress),
	})
	log.logger.SetOutput(output)
	log.logger.SetFormatter(&logursLogger.TextFormatter{
		FullTimestamp: true,
	})
}

func (logger *MusicTogetherLogger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return logger
}

func (logger *MusicTogetherLogger) Info(ctx context.Context, message string, args ...interface{}) {
	if len(args) != 0 {
		logger.logger.WithContext(ctx).Info(message, args)
	} else {
		logger.logger.WithContext(ctx).Info(message)
	}
}

func (logger *MusicTogetherLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	if len(args) != 0 {
		logger.logger.WithContext(ctx).Warn(message, args)
	} else {
		logger.logger.WithContext(ctx).Warn(message)
	}
}

func (logger *MusicTogetherLogger) Error(ctx context.Context, message string, args ...interface{}) {
	if len(args) != 0 {
		logger.logger.WithContext(ctx).Error(message, args)
	} else {
		logger.logger.WithContext(ctx).Error(message)
	}
}

func (logger *MusicTogetherLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logursLogger.Fields{}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		fields[logursLogger.ErrorKey] = err

	}
	logger.logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
}

func Log() *MusicTogetherLogger {
	return log
}
