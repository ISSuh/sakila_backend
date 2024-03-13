package handler

import (
	gohttp "net/http"

	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/http"
	"github.com/labstack/echo/v4"
)

type Healthcheck struct {
	log logger.Logger
}

func NewHealthcheck(l logger.Logger) *Healthcheck {
	return &Healthcheck{
		log: l,
	}
}

func (h *Healthcheck) Check() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.log.Infof("[Healthcheck]")

		res := http.NewResponseOK()
		return c.JSON(gohttp.StatusOK, res)
	}
}
