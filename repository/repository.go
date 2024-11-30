package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Manage  ManagementVoucherInterface
	Voucher VoucherRepository
	Redeem  RedeemRepository
	History HistoryRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Manage:  NewManagementVoucherRepo(db, log),
		Voucher: NewVoucherRepository(db, log),
		Redeem:  NewRedeemRepository(db, log),
		History: NewHistoryRepository(db, log),
	}
}
