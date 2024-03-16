package middleware

import (
	"context"
	"fmt"
	gohttp "net/http"

	"github.com/ISSuh/msago-sample/internal/common"
	"github.com/ISSuh/msago-sample/pkg/http"

	"github.com/labstack/echo/v4"
)

func WrapContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("[WrapContext]")
		var err error
		req, err := http.NewAppRequest(c)
		if err != nil {
			res := http.NewResponseErrWithBody(gohttp.StatusBadRequest, common.InvalidReqeustMessage)
			return c.JSON(gohttp.StatusBadRequest, res)
		}

		resp := http.NewResponseOK()

		appContext := &common.AppContext{c, context.Background()}
		appContext.Ctx = context.WithValue(appContext.Ctx, common.ReqeustContextKey, req)
		appContext.Ctx = context.WithValue(appContext.Ctx, common.ResponseContextKey, resp)

		err = next(appContext)
		if err != nil {
			res := http.NewResponseErrWithBody(gohttp.StatusBadRequest, common.InvalidReqeustMessage)
			return c.JSON(gohttp.StatusBadRequest, res)
		}
		return nil
	}
}
