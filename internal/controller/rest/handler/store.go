package handler

import (
	"net/http"
	"strconv"

	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/service"
	"github.com/labstack/echo/v4"
)

type Store struct {
	log     logger.Logger
	service service.StoreService
}

func NewStoreHandler(l logger.Logger, s service.StoreService) *Store {
	return &Store{
		log:     l,
		service: s,
	}
}

func (s *Store) StoreById() echo.HandlerFunc {
	return func(c echo.Context) error {
		s.log.Infof("[StoreById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		model, err := s.service.StoreById(id)
		if err != nil {
			return err
		}

		s.log.Infof("[StoreById] %+V", model)
		return c.String(http.StatusOK, idStr)
	}
}
