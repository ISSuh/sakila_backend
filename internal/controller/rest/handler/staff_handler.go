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
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendResponseError(c, http.StatusBadRequest, 500, err)
			return
		}

		staff, err := h.service.StaffById(id)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, staff)
	}
}

func (h *StaffHandler) Staffs() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.log.Infof("[Staffs]")

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

		staff, err := h.service.Staffs(page)
		if err != nil {
			response.SendResponseError(c, http.StatusInternalServerError, 2, err)
			return
		}

		response.SendResponseOk(c, staff)
	}
}
