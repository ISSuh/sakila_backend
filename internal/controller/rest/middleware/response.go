package middleware

import (
	"fmt"
	gohttp "net/http"

	"github.com/ISSuh/msago-sample/internal/common"
	"github.com/ISSuh/msago-sample/pkg/http"
	"github.com/labstack/echo/v4"
)

func ResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("[ResponseMiddleware]")
		context := common.ToAppContext(c)
		r := context.Value(common.ResponseContextKey)

		err := next(c)

		resp, _ := r.(*http.AppResponse)
		if err != nil {
			resp.Status = gohttp.StatusInternalServerError
			resp.Error = "error!"
		}
		return nil
	}
}
