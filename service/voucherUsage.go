package service

import (
	"errors"
	"fmt"
	"time"
	"voucher_system/models"
	"voucher_system/repository"

	"go.uber.org/zap"
)

type VoucherService interface {
	FindVouchers(userID int, voucherType string) ([]*models.Voucher, error)
	ValidateVoucher(voucherCode string, transactionAmount float64, shippingAmount float64, area string, paymentMethod string) (*models.Voucher, float64, error)
	UseVoucher(userID int, voucherCode string, transactionAmount float64, paymentMethod string, area string) error
}

type voucherService struct {
	repo repository.Repository
	log *zap.Logger

}

func NewVoucherService(repo repository.Repository, log *zap.Logger) VoucherService {
	return &voucherService{
		repo: repo,
		log: log,
	}
}

func (s *voucherService) FindVouchers(userID int, voucherType string) ([]*models.Voucher, error) {
	vouchers, err := s.repo.Voucher.FindAll(userID, voucherType)
	if err != nil {
		return nil, err
	}
	if len(vouchers) == 0 {
		return nil, errors.New("no vouchers available")
	}
	return vouchers, nil
}

func (s *voucherService) ValidateVoucher(voucherCode string, transactionAmount float64, shippingAmount float64, area string, paymentMethod string) (*models.Voucher, float64, error) {
	currentDate := time.Now()
	voucher, err := s.repo.Voucher.FindValidVoucher(voucherCode, area, transactionAmount, paymentMethod, currentDate)
	if err != nil {
		return nil, 0, err
	}

	// Calculate the benefit value (discount value)
	benefitValue := voucher.DiscountValue
	if voucher.VoucherCategory == "Free Shipping" {
		benefitValue = shippingAmount
	}

	return voucher, benefitValue, nil
}

func (s *voucherService) UseVoucher(userID int, voucherCode string, transactionAmount float64, paymentMethod string, area string) error {
	// Validate the voucher
	voucher, benefitValue, err := s.ValidateVoucher(voucherCode, transactionAmount, 0, area, paymentMethod)
	if err != nil {
		return err
	}

	// Save voucher usage history
	history := &models.History{
		UserID:            userID,
		VoucherID:         voucher.ID,
		TransactionAmount: transactionAmount,
		BenefitValue:      benefitValue,
		UsageDate:         time.Now(),
	}

	err = s.repo.History.CreateHistory(history)
	if err != nil {
		return err
	}

	// Update voucher quota
	newQuota := voucher.Quota - 1
	if newQuota < 0 {
		return fmt.Errorf("voucher quota exceeded")
	}

	err = s.repo.Voucher.UpdateVoucherQuota(voucher.ID, newQuota)
	if err != nil {
		return err
	}

	// Log the voucher redemption
	redeem := &models.Redeem{
		UserID:     userID,
		VoucherID:  voucher.ID,
		RedeemDate: time.Now(),
	}
	return s.repo.Redeem.CreateRedeem(redeem)
}
