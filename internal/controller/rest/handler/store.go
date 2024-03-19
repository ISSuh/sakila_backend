package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/service"
	"github.com/gin-gonic/gin"
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

func (s *Store) StoreById() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		store, err := s.service.StoreById(id)
		if err != nil {
			return
		}

		temp, _ := json.Marshal(store)
		fmt.Println(string(temp))

		c.JSON(http.StatusOK, store)
	}
}

func (s *Store) StoreAddressById() gin.HandlerFunc {
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

func (s *Store) StoreOnCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.log.Infof("[StoreOnCountry] id : %s", c.Param("id"))

		countryIdStr := c.Param("id")
		countryId, _ := strconv.Atoi(countryIdStr)
		stores, err := s.service.StoresOnCountry(countryId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, stores)
	}
}
