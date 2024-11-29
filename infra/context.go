package infra

import (
	"voucher_system/config"
	"voucher_system/controller"
	"voucher_system/database"
	"voucher_system/helper"
	"voucher_system/repository"
	"voucher_system/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Cfg config.Config
	DB  *gorm.DB
	Ctl controller.Controller
	Log *zap.Logger
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := helper.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := controller.NewController(service, log)

	return &ServiceContext{Cfg: config, DB: db, Ctl: *Ctl, Log: log}, nil
}
