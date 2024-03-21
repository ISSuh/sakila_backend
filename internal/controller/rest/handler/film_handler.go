package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/service"
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
		id, _ := strconv.Atoi(idStr)
		store, err := h.film.FilmById(id)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, store)
	}
}

func (h *FilmHandler) Films() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Films]")

		var page common.Pagnation
		if err := c.Bind(&page); err != nil {
			return
		}

		fmt.Println(common.Dump(page))

		if !page.IsValidate() {
			return
		}

		filmsPages, err := h.film.Films(&page)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, filmsPages)
	}
}
