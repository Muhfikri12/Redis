package models

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID              uint      `gorm:"primaryKey"`
	VoucherName     string    `gorm:"size:255;not null"`
	VoucherCode     string    `gorm:"size:50;unique;not null"`
	VoucherType     string    `gorm:"type:ENUM('e-commerce', 'redeem points');not null"`
	PointsRequired  int       `gorm:"default:0"`
	Description     string    `gorm:"not null"`
	VoucherCategory string    `gorm:"type:ENUM('Free Shipping', 'Discount');not null"`
	DiscountValue   float64   `gorm:"type:decimal(10,2);not null"`
	MinimumPurchase float64   `gorm:"type:decimal(10,2);default:0"`
	PaymentMethods  []string  `gorm:"type:text[];default:array[]::text[]"`
	StartDate       time.Time `gorm:"not null"`
	EndDate         time.Time `gorm:"not null"`
	ApplicableAreas []string  `gorm:"type:text[];default:array[]::text[]"`
	Quota           int       `gorm:"default:0"`
	Status          bool      `gorm:"default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (v *Voucher) BeforeSave(tx *gorm.DB) (err error) {
	currentDate := time.Now()
	v.Status = currentDate.After(v.StartDate) && currentDate.Before(v.EndDate)
	return
}

type Redeem struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	VoucherID  uint      `gorm:"constraint:OnDelete:SET NULL"`
	RedeemDate time.Time `gorm:"default:CURRENT_DATE"`
}

type History struct {
	ID                uint      `gorm:"primaryKey"`
	UserID            uint      `gorm:"not null"`
	VoucherID         uint      `gorm:"constraint:OnDelete:SET NULL"`
	UsageDate         time.Time `gorm:"default:CURRENT_DATE"`
	TransactionAmount float64   `gorm:"type:decimal(10,2);not null"`
	BenefitValue      float64   `gorm:"type:decimal(10,2);not null"`
}
