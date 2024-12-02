package controller

import (
	managementvoucherhandler "voucher_system/controller/management_voucher_handler"
	"voucher_system/service"

	"go.uber.org/zap"
)

type Controller struct {
	Manage  managementvoucherhandler.ManageVoucherHandler
	Voucher VoucherController
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		Manage:  managementvoucherhandler.NewManagementVoucherHanlder(service, logger),
		Voucher: *NewVoucherController(service, logger),
	}
}
