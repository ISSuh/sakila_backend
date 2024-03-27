package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ISSuh/sakila_backend/internal/controller/rest/response"
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ActorForm struct {
	WithFilm bool `form:"with_film"`
}

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

		var actorForm ActorForm
		if err := c.Bind(&actorForm); err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		actor, err := h.service.ActorById(id, actorForm.WithFilm)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, actor)
	}
}

func (h *ActorHandler) FilmOnActorById() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[FilmOnActorById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		films, err := h.service.FilmOnActorById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, films)
	}
}

func (h *ActorHandler) ActorByLastName() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[ActorByLastName] lastName : %s", c.Param("lastName"))

		lastName := c.Param("lastName")
		if len(lastName) < 1 {
			err := errors.New("need parameter")
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		actors, err := h.service.ActorByLastName(lastName)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, actors)
	}
}
