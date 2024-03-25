package log

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
)

func NewLogger() logger.Logger {
	l := NewZapLogger()
	return l
}
