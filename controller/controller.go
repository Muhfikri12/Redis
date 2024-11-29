package controller

import (
	"voucher_system/service"

	"go.uber.org/zap"
)

type Controller struct {
	Voucher VoucherController
	

}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		Voucher: *NewVoucherController(service, logger),

	}
}
