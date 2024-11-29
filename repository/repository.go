package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Voucher VoucherRepository
	Redeem  RedeemRepository
	History HistoryRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Voucher: NewVoucherRepository(db, log),
		Redeem:  NewRedeemRepository(db, log),
		History: NewHistoryRepository(db, log),
	}
}
