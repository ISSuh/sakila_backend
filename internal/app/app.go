package app

import (
	"strconv"

	"github.com/ISSuh/msago-sample/internal/controller/rest/middleware"
	"github.com/ISSuh/msago-sample/internal/controller/rest/router"
	"github.com/ISSuh/msago-sample/internal/factory"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/config"

	"github.com/labstack/echo/v4"
)

type Application struct {
	config *config.Config
	log    logger.Logger

	s *factory.Services
	r *factory.Repositories
	h *factory.Handlers
	m *middleware.Middleware

	e *echo.Echo
}

func NewApplication(l logger.Logger, configPath string) (*Application, error) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	a := &Application{
		config: config,
		log:    l,
		e:      echo.New(),
	}
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

	if err = a.initHandler(); err != nil {
		return err
	}

	if err = a.initMiddleware(); err != nil {
		return err
	}

	if err = router.Route(a.e, a.m, a.h); err != nil {
		return err
	}

	return nil
}

func (a *Application) Start() error {
	port := strconv.Itoa(a.config.Server.Port)
	return a.e.Start(":" + port)
}

func (a *Application) initRepository() error {
	a.log.Infof("init repository")

	var err error
	var repositories *factory.Repositories

	if repositories, err = factory.NewRepositories(a.log); err != nil {
		return err
	}

	a.r = repositories
	return nil
}

func (a *Application) initService() error {
	a.log.Infof("init service")

	var err error
	var services *factory.Services

	if services, err = factory.NewServices(a.log, a.r); err != nil {
		return err
	}

	a.s = services
	return nil
}

func (a *Application) initHandler() error {
	a.log.Infof("init handler")

	var err error
	var handlers *factory.Handlers

	if handlers, err = factory.NewHandlers(a.log, a.s); err != nil {
		return err
	}

	a.h = handlers
	return nil
}

func (a *Application) initMiddleware() error {
	a.log.Infof("init middleware")
	return a.m.RegistMiddlware(a.e)
}
