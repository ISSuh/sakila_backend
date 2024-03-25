package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BodyKey         string = "body"
	ErrorMessageKey string = "error_body"
)

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type AppResponse struct {
	Status  int `json:"status"`
	Message any `json:"message,omitempty"`
	Error   any `json:"error,omitempty"`
}

func SendResponseOk(c *gin.Context, body any) {
	c.Status(http.StatusOK)
	c.Set(BodyKey, body)
}

func SendResponseError(c *gin.Context, status int, code int, err error) {
	body := &ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
	c.Status(status)
	c.Set(ErrorMessageKey, body)
}
