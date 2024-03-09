package app

import (
	"github.com/ISSuh/msago-sample/internal/config"
	"github.com/ISSuh/msago-sample/internal/controller/api/handler"
	"github.com/ISSuh/msago-sample/internal/controller/api/middleware"
	"github.com/ISSuh/msago-sample/internal/controller/api/router"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/labstack/echo/v4"
)

type Application struct {
	config config.Config
	log    logger.Logger

	h handler.Handler
	m middleware.Middleware
	r router.Router

	echo *echo.Echo
}

func NewApplication(l logger.Logger, configPath string) (*Application, error) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	a := &Application{
		config: config,
		log:    l,
		echo:   echo.New(),
	}

	a.initRouter()

	return a, nil
}

func (a *Application) Init() error {
	var err error

	if err = a.initRepository(); err != nil {
		return err
	}

	if err = a.initService(); err != nil {
		return err
	}

	if err = a.initRouter(); err != nil {
		return err
	}

	return nil
}

func (a *Application) Start() {
	a.echo.Start(":33669")
}

func (a *Application) initRouter() error {
	v1 := a.echo.Group("v1")

	{
		v1Item := v1.Group("/item")
		v1Item.GET(":itemId", func(echoCtx echo.Context) error {
			return nil
		})

	}

	return nil
}

func (a *Application) initMiddleware() error {
	return a.m.RegistMiddlware(a.echo)
}

func (a *Application) initService() error {
	return nil
}

func (a *Application) initRepository() error {
	return nil
}
