package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ISSuh/sakila_backend/internal/controller/rest/response"
	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	gin.ResponseWriter

	body *bytes.Buffer
}

func (w *ResponseWrapper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := &ResponseWrapper{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}

		c.Writer = writer

		c.Next()

		c.Writer.Header().Set("Content-Type", "application/json")
		status := c.Writer.Status()
		res := &response.AppResponse{
			Status: status,
		}

		if http.StatusOK <= status && status <= http.StatusIMUsed {
			body, _ := c.Get(response.BodyKey)
			res.Message = body
		} else {
			body, _ := c.Get(response.ErrorMessageKey)
			res.Error = body
		}

		responseByte, _ := json.Marshal(res)
		response := string(responseByte)
		writer.ResponseWriter.WriteString(response)
		writer.body.Reset()
	}
}
