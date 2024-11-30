package service

import (
	"voucher_system/repository"

	"go.uber.org/zap"
)

type Service struct {
	// User UserService
	Manage ManageVoucherService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		// User: NewUserService(repo.User),
		Manage: NewManagementVoucherService(repo, log),
	}
}
