package app

import (
	"github.com/ISSuh/msago-sample/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	config config.Config

	echo *echo.Echo
}

func NewApplication(path string) (*Application, error) {
	config, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	a := &Application{
		config: config,
		echo:   echo.New(),
	}

	a.initRouter()

	return a, nil
}

func (a *Application) Start() {
}

func (a *Application) initRouter() {
	v1 := a.echo.Group(
		"v1",
		middleware.CORS(),
	)

	{
		v1Item := v1.Group("/item")
		v1Item.GET(":itemId", func(echoCtx echo.Context) error {
			return nil
		})

	}
}

func (a *Application) initService() {

}

func (a *Application) initRepository() {

}
