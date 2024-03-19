package common

import (
	"io"
	gohttp "net/http"

	"github.com/labstack/echo/v4"
)

type RequestHeader map[string]string

type AppRequest struct {
	Header gohttp.Header
	Cookie []*gohttp.Cookie
	Body   string
}

func NewAppRequest(e echo.Context) (*AppRequest, error) {
	h := e.Request().Header
	c := e.Request().Cookies()

	buffer, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return nil, err
	}

	return &AppRequest{
		Header: h,
		Cookie: c,
		Body:   string(buffer),
	}, nil
}
