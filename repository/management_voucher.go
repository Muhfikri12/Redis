package repository

import (
	"fmt"
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ManagementVoucherInterface interface {
	CreateVoucher(voucher *models.Voucher) error
	SoftDeleteVoucher(voucherID int) error
	UpdateVoucher(voucher *models.Voucher, voucherID int) error
	ShowRedeemPoints() (*[]RedeemPoint, error)
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

func (m *ManagementVoucherRepo) SoftDeleteVoucher(voucherID int) error {

	err := m.DB.Delete(&models.Voucher{}, voucherID).Error
	if err != nil {
		m.log.Error("Error from repo soft deleting voucher:", zap.Error(err))
		return err
	}

	return nil
}

func (m *ManagementVoucherRepo) UpdateVoucher(voucher *models.Voucher, voucherID int) error {

	result := m.DB.Model(&voucher).
		Where("id = ?", voucherID).
		Updates(voucher)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with shipping_id %d", voucherID)
	}

	return nil
}

type RedeemPoint struct {
	VoucherName    string  `json:"voucher_name"`
	PointsRequired int     `json:"points_required"`
	DiscountValue  float64 `json:"discount_value"`
}

func (m *ManagementVoucherRepo) ShowRedeemPoints() (*[]RedeemPoint, error) {

	voucher := []RedeemPoint{}

	query := m.DB.
		Table("vouchers as v").
		Select(`v.voucher_name, v.discount_value, v.points_required`).
		Where("voucher_type = ?", "redeem points")

	err := query.Find(&voucher).Error
	if err != nil {
		return nil, err
	}

	return &voucher, nil
}
