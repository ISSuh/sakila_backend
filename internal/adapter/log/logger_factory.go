package log

import (
	"github.com/ISSuh/msago-sample/internal/logger"
)

func NewLogger() logger.Logger {
	l := NewZapLogger()
	return l
}
