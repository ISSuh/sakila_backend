package log

import (
	"github.com/ISSuh/monolith-sakila/internal/logger"
)

func NewLogger() logger.Logger {
	l := NewZapLogger()
	return l
}
