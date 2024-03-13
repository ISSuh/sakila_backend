package log

import (
	"github.com/ISSuh/msago-sample/internal/logger"
	"go.uber.org/zap"
)

type ZapLogger struct {
	zapLogger *zap.SugaredLogger
}

func NewZapLogger() logger.Logger {
	log, err := zap.NewDevelopment()
	if err != nil {
		return nil
	}

	sugar := log.Sugar()
	defer log.Sync()

	return &ZapLogger{
		zapLogger: sugar,
	}
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.zapLogger.Infof(format, args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.zapLogger.Warnf(format, args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.zapLogger.Errorf(format, args...)
}

func (l *ZapLogger) Fatalln(args ...interface{}) {
	l.zapLogger.Fatal(args...)
}

func (l *ZapLogger) WithFields(items logger.Fields) logger.Logger {
	field := make([]interface{}, 0)
	for index, value := range items {
		field = append(field, index)
		field = append(field, value)
	}

	logger := l.zapLogger.With(field...)
	return &ZapLogger{
		zapLogger: logger,
	}
}

func (l *ZapLogger) WithError(err error) logger.Logger {
	logger := l.zapLogger.With(err.Error())
	return &ZapLogger{
		zapLogger: logger,
	}
}
