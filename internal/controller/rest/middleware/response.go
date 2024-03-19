package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ISSuh/msago-sample/internal/common"
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

		bodyStr := writer.body.String()
		status := c.Writer.Status()

		res := common.NewRespons()
		res.Status = status
		res.Message = bodyStr
		if status != http.StatusOK {
			res.Error = "Error!!"
		}

		responseByte, _ := json.Marshal(res)
		responseStr := string(responseByte)
		writer.ResponseWriter.WriteString(responseStr)
		writer.body.Reset()

		fmt.Printf("[ResponseMiddleware] end\n")
	}
}
