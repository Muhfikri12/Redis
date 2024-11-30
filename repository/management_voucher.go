package repository

import (
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ManagementVoucherInterface interface {
	CreateVoucher(voucher *models.Voucher) error
}

type ManagementVoucherRepo struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewManagementVoucherRepo(db *gorm.DB, log *zap.Logger) ManagementVoucherInterface {
	return &ManagementVoucherRepo{DB: db, log: log}
}

func (m *ManagementVoucherRepo) CreateVoucher(voucher *models.Voucher) error {
	err := m.DB.Create(voucher).Error
	if err != nil {
		m.log.Error("Error from repo creating voucher:", zap.Error(err))
		return err
	}

	return nil
}
