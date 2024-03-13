package middleware

import (
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
}

func NewMiddleware() Middleware {
	return Middleware{}
}

func (m *Middleware) RegistMiddlware(e *echo.Echo) error {
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORS())
	return nil
}
