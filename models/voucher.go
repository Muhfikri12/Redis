package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID              int        `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherName     string     `gorm:"type:varchar(255);not null" json:"voucher_name"`
	VoucherCode     string     `gorm:"type:varchar(50);unique;not null" json:"voucher_code"`
	VoucherType     string     `gorm:"type:varchar(20);not null;check:voucher_type in ('e-commerce', 'redeem points')" json:"voucher_type"`
	PointsRequired  int        `gorm:"default:0" json:"points_required"`
	Description     string     `gorm:"type:text;not null" json:"description"`
	VoucherCategory string     `gorm:"type:varchar(20);not null;check:voucher_category in ('Free Shipping', 'Discount')" json:"voucher_category"`
	DiscountValue   float64    `gorm:"type:numeric(10,2);not null" json:"discount_value"`
	MinimumPurchase float64    `gorm:"type:numeric(10,2);default:0" json:"minimum_purchase"`
	PaymentMethods  []string   `gorm:"type:jsonb" json:"payment_methods"`
	StartDate       time.Time  `gorm:"type:timestamp with time zone;not null" json:"start_date"`
	EndDate         time.Time  `gorm:"type:timestamp with time zone;not null" json:"end_date"`
	ApplicableAreas []string   `gorm:"type:jsonb" json:"applicable_areas"`
	Quota           int        `gorm:"default:0" json:"quota"`
	Status          bool       `gorm:"type:boolean" json:"status"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (v *Voucher) BeforeSave(tx *gorm.DB) (err error) {
	currentDate := time.Now()
	v.Status = currentDate.After(v.StartDate) && currentDate.Before(v.EndDate)
	// Marshal PaymentMethods and ApplicableAreas to JSON before saving
	if len(v.PaymentMethods) > 0 {
		// Marshal to JSON to ensure proper formatting
		paymentMethodsJSON, err := json.Marshal(v.PaymentMethods)
		if err != nil {
			return err
		}
		v.PaymentMethods = nil // Clear the original array to use the marshaled value
		v.PaymentMethods = []string{string(paymentMethodsJSON)}
	}

	if len(v.ApplicableAreas) > 0 {
		// Marshal to JSON to ensure proper formatting
		applicableAreasJSON, err := json.Marshal(v.ApplicableAreas)
		if err != nil {
			return err
		}
		v.ApplicableAreas = nil // Clear the original array to use the marshaled value
		v.ApplicableAreas = []string{string(applicableAreasJSON)}
	}

	return nil
}

type Redeem struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int       `gorm:"not null" json:"user_id"`
	VoucherID  int       `gorm:"not null" json:"voucher_id"`
	RedeemDate time.Time `gorm:"default:current_date" json:"redeem_date"`
}

type History struct {
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            int       `gorm:"not null" json:"user_id"`
	VoucherID         int       `gorm:"not null" json:"voucher_id"`
	UsageDate         time.Time `gorm:"default:current_date" json:"usage_date"`
	TransactionAmount float64   `gorm:"type:numeric(10,2);not null" json:"transaction_amount"`
	BenefitValue      float64   `gorm:"type:numeric(10,2);not null" json:"benefit_value"`
}
