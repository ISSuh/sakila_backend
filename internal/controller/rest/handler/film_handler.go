package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ISSuh/sakila_backend/internal/common"
	"github.com/ISSuh/sakila_backend/internal/controller/rest/response"
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/service"
	"github.com/gin-gonic/gin"
)

type FilmHandler struct {
	log  logger.Logger
	film service.FilmService
}

func NewFilmHandler(l logger.Logger, s service.FilmService) *FilmHandler {
	return &FilmHandler{
		log:  l,
		film: s,
	}
}

func (h *FilmHandler) FilmById() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[FilmById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		film, err := h.film.FilmById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		response.SendResponseOk(c, film)
	}
}

func (h *FilmHandler) Films() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Films]")

		var page common.Pagnation
		if err := c.Bind(&page); err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		if !page.IsValidate() {
			err := errors.New("validate fail about page parameter")
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		filmsPages, err := h.film.Films(&page)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 2, err)
			return
		}

		response.SendResponseOk(c, filmsPages)
	}
}
