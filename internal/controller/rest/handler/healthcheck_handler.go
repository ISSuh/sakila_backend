package handler

import (
	gohttp "net/http"

	"github.com/ISSuh/sakila_backend/internal/common"
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	log logger.Logger
}

func NewHealthcheck(l logger.Logger) *HealthcheckHandler {
	return &HealthcheckHandler{
		log: l,
	}
}

func (h *HealthcheckHandler) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Healthcheck]")

		res := common.NewResponseOK()
		c.JSON(gohttp.StatusOK, res)
	}
}
