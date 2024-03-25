package handler

import (
	"net/http"
	"strconv"

	"github.com/ISSuh/sakila_backend/internal/controller/rest/response"
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/service"
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
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		store, err := s.service.StoreById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, store)
	}
}

func (s *StoreHandler) StoreAddressById() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreAddressById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		store, err := s.service.StoreAddressById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, store)
	}
}

func (s *StoreHandler) StoreOnCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreOnCountry] id : %s", c.Param("countryId"))

		countryIdStr := c.Param("countryId")
		countryId, err := strconv.Atoi(countryIdStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		stores, err := s.service.StoresOnCountry(countryId)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, stores)
	}
}
