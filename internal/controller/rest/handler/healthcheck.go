package handler

import (
	gohttp "net/http"

	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	log logger.Logger
}

func NewHealthcheck(l logger.Logger) *Healthcheck {
	return &Healthcheck{
		log: l,
	}
}

func (h *Healthcheck) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Healthcheck]")

		res := common.NewResponseOK()
		c.JSON(gohttp.StatusOK, res)
	}
}
