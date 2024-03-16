package common

import (
	"context"

	"github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
	Ctx context.Context
}

func ToAppContext(e echo.Context) *AppContext {
	return e.(*AppContext)
}

func (a *AppContext) Value(key contextKey) any {
	return a.Ctx.Value(key)
}
