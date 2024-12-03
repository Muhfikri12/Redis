package repository

import (
	"voucher_system/repository/auth"
	managementvoucher "voucher_system/repository/management_voucher"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Manage  managementvoucher.ManagementVoucherInterface
	Voucher VoucherRepository
	Redeem  RedeemRepository
	History HistoryRepository
	Auth    auth.AuthRepoInterface
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Manage:  managementvoucher.NewManagementVoucherRepo(db, log),
		Voucher: NewVoucherRepository(db, log),
		Redeem:  NewRedeemRepository(db, log),
		History: NewHistoryRepository(db, log),
		Auth:    auth.NewManagementVoucherRepo(db, log),
	}
}
