package infra

import (
	"voucher_system/config"
	"voucher_system/controller"
	"voucher_system/database"
	"voucher_system/helper"
	"voucher_system/middleware"
	"voucher_system/repository"
	"voucher_system/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Cfg        config.Config
	Config     config.Config
	DB         *gorm.DB
	Ctl        controller.Controller
	Log        *zap.Logger
	Cacher     database.Cacher
	Middleware middleware.Middleware
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	log, err := helper.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*5)

	repository := repository.NewRepository(db, log)

	service := service.NewService(repository, log)

	Ctl := controller.NewController(service, log, rdb)

	middleware := middleware.NewMiddleware(rdb)

	return &ServiceContext{Cfg: config, DB: db, Ctl: *Ctl, Log: log, Middleware: middleware}, nil
}
