package controller

import (
	authhandler "voucher_system/controller/auth_handler"
	managementvoucherhandler "voucher_system/controller/management_voucher_handler"
	"voucher_system/database"
	"voucher_system/service"

	"go.uber.org/zap"
)

type Controller struct {
	Manage  managementvoucherhandler.ManageVoucherHandler
	Voucher VoucherController
	Auth    authhandler.AuthHadler
}

func NewController(service service.Service, logger *zap.Logger, rdb database.Cacher) *Controller {

	return &Controller{
		Manage:  managementvoucherhandler.NewManagementVoucherHanlder(service, logger),
		Voucher: *NewVoucherController(service, logger),
		Auth:    authhandler.NewUserHandler(service, logger, rdb),
	}
}
