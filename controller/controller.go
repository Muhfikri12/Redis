package controller

import (
	"voucher_system/service"

	"go.uber.org/zap"
)

type Controller struct {
	// User UserController
	Manage ManageVoucherHandler
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		// User: *NewUserController(service.User, logger),
		Manage: NewManagementVoucherHanlder(service, logger),
	}
}
