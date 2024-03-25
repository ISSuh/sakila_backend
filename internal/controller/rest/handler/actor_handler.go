package handler

import (
	"net/http"
	"strconv"

	"github.com/ISSuh/sakila_backend/internal/controller/rest/response"
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ActorHandler struct {
	log     logger.Logger
	service service.ActorService
}

func NewActorHandler(l logger.Logger, s service.ActorService) *ActorHandler {
	return &ActorHandler{
		log:     l,
		service: s,
	}
}

func (h *ActorHandler) ActorById() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[ActorById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		actor, err := h.service.ActorById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, actor)
	}
}
