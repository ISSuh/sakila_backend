package handler

import (
	"net/http"
	"strconv"

	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/service"
	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	log     logger.Logger
	service service.StoreService
}

func NewStoreHandler(l logger.Logger, s service.StoreService) *StoreHandler {
	return &StoreHandler{
		log:     l,
		service: s,
	}
}

func (s *StoreHandler) StoreById() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		store, err := s.service.StoreById(id)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, store)
	}
}

func (s *StoreHandler) StoreAddressById() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreAddressById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		store, err := s.service.StoreAddressById(id)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, store)
	}
}

func (s *StoreHandler) StoreOnCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreOnCountry] id : %s", c.Param("countryId"))

		countryIdStr := c.Param("countryId")
		countryId, _ := strconv.Atoi(countryIdStr)
		stores, err := s.service.StoresOnCountry(countryId)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, stores)
	}
}
