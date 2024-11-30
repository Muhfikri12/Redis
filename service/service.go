package service

import (
	"voucher_system/repository"

	"go.uber.org/zap"
)

type Service struct {
	Manage  ManageVoucherService
	Voucher VoucherService
	History HistoryService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Manage:  NewManagementVoucherService(repo, log),
		Voucher: NewVoucherService(repo, log),
		History: NewHistoryService(repo, log),
	}
}
