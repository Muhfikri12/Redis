package service

import (
	"voucher_system/models"
	"voucher_system/repository"

	"go.uber.org/zap"
)

type ManageVoucherService interface {
	CreateVoucher(voucher *models.Voucher) error
}

type ManagementVoucherservice struct {
	repo repository.Repository
	log  *zap.Logger
}

func NewManagementVoucherService(repo repository.Repository, log *zap.Logger) ManageVoucherService {
	return &ManagementVoucherservice{repo: repo, log: log}
}

func (ms *ManagementVoucherservice) CreateVoucher(voucher *models.Voucher) error {

	if err := ms.repo.Manage.CreateVoucher(voucher); err != nil {
		ms.log.Error("Error from service creating voucher: " + err.Error())
		return err
	}
	return nil
}
