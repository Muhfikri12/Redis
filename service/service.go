package service

import (
	"voucher_system/repository"
	authservice "voucher_system/service/auth_service"
	managementvoucherservice "voucher_system/service/management_voucher_service"

	"go.uber.org/zap"
)

type Service struct {
	Manage  managementvoucherservice.ManageVoucherService
	Voucher VoucherService
	History HistoryService
	Auth    authservice.AuthServiceInterface
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Manage:  managementvoucherservice.NewManagementVoucherService(repo, log),
		Voucher: NewVoucherService(repo, log),
		History: NewHistoryService(repo, log),
		Auth:    authservice.NewManagementVoucherService(repo, log),
	}
}
