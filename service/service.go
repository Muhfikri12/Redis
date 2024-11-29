package service

import (
	"voucher_system/repository"

	"go.uber.org/zap"
)

type Service struct {
	// User UserService
	Voucher VoucherService

}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Voucher: NewVoucherService(repo, log),
	}
}
