package repository

import (
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RedeemRepository interface {
	CreateRedeem(redeem *models.Redeem) error
	FindByUserAndVoucher(userID int, voucherID int) (*models.Redeem, error)
}

type redeemRepository struct {
	DB *gorm.DB
	log *zap.Logger
}

func NewRedeemRepository(db *gorm.DB, log *zap.Logger) RedeemRepository {
	return &redeemRepository{DB: db, log: log}
}

func (r *redeemRepository) CreateRedeem(redeem *models.Redeem) error {
	return r.DB.Create(redeem).Error
}

func (r *redeemRepository) FindByUserAndVoucher(userID int, voucherID int) (*models.Redeem, error) {
	var redeem models.Redeem
	err := r.DB.Where("user_id = ? AND voucher_id = ?", userID, voucherID).First(&redeem).Error
	if err != nil {
		return nil, err
	}
	return &redeem, nil
}
