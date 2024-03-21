package handler

import (
	"net/http"
	"strconv"

	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/service"
	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	log     logger.Logger
	service service.StaffService
}

func NewStaffHandler(l logger.Logger, s service.StaffService) *StaffHandler {
	return &StaffHandler{
		log:     l,
		service: s,
	}
}

func (h *StaffHandler) StaffById() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[StaffById] id : %s", c.Param("id"))

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		staff, err := h.service.StaffById(id)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, staff)
	}
}

func (h *StaffHandler) Staffs() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Staffs]")

		var page common.Pagnation
		if err := c.Bind(&page); err != nil {
			return
		}

		if !page.IsValidate() {
			return
		}

		staff, err := h.service.Staffs(page)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, staff)
	}
}
