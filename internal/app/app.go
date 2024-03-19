package app

import (
	"strconv"

	"github.com/ISSuh/msago-sample/internal/controller/rest/router"
	"github.com/ISSuh/msago-sample/internal/factory"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/config"
	"github.com/ISSuh/msago-sample/pkg/db"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *config.Config
	log    logger.Logger

	serice     *factory.Services
	repository *factory.Repositories
	handler    *factory.Handlers

	engine *gin.Engine
	db     *db.Database
}

func NewApplication(l logger.Logger, configPath string) (*Application, error) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	a := &Application{
		config: config,
		log:    l,
		engine: gin.Default(),
		db:     db.NewDatabase(config.Database),
	}
	return a, nil
}

func (a *Application) Init() error {
	var err error

	if err = a.initInfrastructure(); err != nil {
		return err
	}

	if err = a.initRepository(); err != nil {
		return err
	}

	if err = a.initService(); err != nil {
		return err
	}

	if err = a.initHandler(); err != nil {
		return err
	}

	if err = router.Route(a.engine, a.handler); err != nil {
		return err
	}

	return nil
}

func (a *Application) Start() error {
	port := strconv.Itoa(a.config.Server.Port)
	address := ":" + port
	return a.engine.Run(address)
}

func (a *Application) initInfrastructure() error {
	a.log.Infof("init infrastructure")

	var err error
	if err = a.db.Connect(); err != nil {
		return err
	}
	return nil
}

func (a *Application) initRepository() error {
	a.log.Infof("init repository")

	var err error
	var repositories *factory.Repositories

	if repositories, err = factory.NewRepositories(a.log, a.db); err != nil {
		return err
	}

	a.repository = repositories
	return nil
}

func (a *Application) initService() error {
	a.log.Infof("init service")

	var err error
	var services *factory.Services

	if services, err = factory.NewServices(a.log, a.repository); err != nil {
		return err
	}

	a.serice = services
	return nil
}

func (a *Application) initHandler() error {
	a.log.Infof("init handler")

	var err error
	var handlers *factory.Handlers

	if handlers, err = factory.NewHandlers(a.log, a.serice); err != nil {
		return err
	}

	a.handler = handlers
	return nil
}
