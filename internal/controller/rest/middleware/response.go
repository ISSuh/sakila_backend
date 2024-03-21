package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ISSuh/monolith-sakila/internal/common"
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
		fmt.Printf("[ResponseMiddleware] start\n")

		writer := &ResponseWrapper{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}

		c.Writer = writer
		c.Next()

		body, _ := c.Get("body")

		// bodyStr := writer.body.String()
		status := c.Writer.Status()

		res := common.NewRespons()
		res.Status = status
		res.Message = body
		if status != http.StatusOK {
			res.Error = "Error!!"
		}

		fmt.Println(res)

		responseByte, _ := json.Marshal(res)
		responseStr := string(responseByte)
		writer.ResponseWriter.WriteString(responseStr)
		writer.body.Reset()

		fmt.Printf("[ResponseMiddleware] end\n")
	}
}
