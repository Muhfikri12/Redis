package controller

import (
	"voucher_system/service"

	"go.uber.org/zap"
)

type Controller struct {
	Manage  ManageVoucherHandler
	Voucher VoucherController
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		Manage:  NewManagementVoucherHanlder(service, logger),
		Voucher: *NewVoucherController(service, logger),
	}
}
