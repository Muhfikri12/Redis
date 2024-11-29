package repository

import (
	"encoding/json"
	"time"
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	FindAll(userID int, voucherType string) ([]*models.Voucher, error)
	FindByVoucherCode(voucherCode string) (*models.Voucher, error)
	FindValidVoucher(voucherCode, area string, transactionAmount float64, paymentMethod string, currentDate time.Time) (*models.Voucher, error)
	UpdateVoucherQuota(voucherID int, quota int) error
}

type voucherRepository struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewVoucherRepository(db *gorm.DB, log *zap.Logger) VoucherRepository {
	return &voucherRepository{DB: db, log: log}
}

func (r *voucherRepository) FindAll(userID int, voucherType string) ([]*models.Voucher, error) {
	var rawVouchers []struct {
		models.Voucher
		RawPaymentMethods  []byte `gorm:"column:payment_methods"`
		RawApplicableAreas []byte `gorm:"column:applicable_areas"`
	}

	query := r.DB.
		Table("vouchers").
		Select(`vouchers.*, vouchers.payment_methods AS raw_payment_methods, vouchers.applicable_areas AS raw_applicable_areas`).
		Joins("JOIN redeems ON redeems.voucher_id = vouchers.id").
		Where("redeems.user_id = ? AND vouchers.status = ?", userID, true)

	if voucherType != "" {
		query = query.Where("vouchers.voucher_type = ?", voucherType)
	}

	err := query.Find(&rawVouchers).Error
	if err != nil {
		return nil, err
	}

	vouchers := make([]*models.Voucher, 0, len(rawVouchers))
	for _, rawVoucher := range rawVouchers {
		v := rawVoucher.Voucher

		if len(rawVoucher.RawPaymentMethods) > 0 {
			if err := json.Unmarshal(rawVoucher.RawPaymentMethods, &v.PaymentMethods); err != nil {
				return nil, err
			}
		}

		if len(rawVoucher.RawApplicableAreas) > 0 {
			if err := json.Unmarshal(rawVoucher.RawApplicableAreas, &v.ApplicableAreas); err != nil {
				return nil, err
			}
		}
		vouchers = append(vouchers, &v)

	}

	return vouchers, nil
}

func (r *voucherRepository) FindByVoucherCode(voucherCode string) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.DB.Where("voucher_code = ?", voucherCode).First(&voucher).Error
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) FindValidVoucher(voucherCode, area string, transactionAmount float64, paymentMethod string, currentDate time.Time) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.DB.Where("voucher_code = ? AND status = ? AND ? BETWEEN start_date AND end_date AND ? >= minimum_purchase AND ? = ANY(applicable_areas) AND ? = ANY(payment_methods) AND quota > 0",
		voucherCode, true, currentDate, transactionAmount, area, paymentMethod).
		First(&voucher).Error
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *voucherRepository) UpdateVoucherQuota(voucherID int, quota int) error {
	err := r.DB.Model(&models.Voucher{}).Where("id = ?", voucherID).Update("quota", quota).Error
	return err
}
