package seeder

import (
	"log"
	"time"
	"voucher_system/models"

	"gorm.io/gorm"
)

func SeedVouchers(db *gorm.DB) error {
	vouchers := []models.Voucher{
		{
			VoucherName:     "Free Shipping October",
			VoucherCode:     "FREEOCT",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Enjoy free shipping for all orders in October.",
			VoucherCategory: "Free Shipping",
			DiscountValue:   10000.00,
			MinimumPurchase: 50000.00,
			PaymentMethods:  []string{"Credit Card", "PayPal"},
			StartDate:       time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: []string{"Sumatera", "Jawa"},
			Quota:           20,
			Status:          true, // Disesuaikan dengan hook untuk perhitungan otomatis
		},
		{
			VoucherName:     "10% Discount November",
			VoucherCode:     "DISCNOV",
			VoucherType:     "redeem points",
			PointsRequired:  500,
			Description:     "Redeem this voucher to get 10% off in November.",
			VoucherCategory: "Discount",
			DiscountValue:   10.00,
			MinimumPurchase: 200000.00,
			PaymentMethods:  []string{"Credit Card", "Bank Transfer"},
			StartDate:       time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 11, 30, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: []string{"Bali"},
			Quota:           25,
			Status:          true, // Disesuaikan dengan hook untuk perhitungan otomatis
		},
	}

	for _, voucher := range vouchers {
		if err := db.Create(&voucher).Error; err != nil {
			log.Printf("Failed to seed voucher: %s, error: %v", voucher.VoucherName, err)
		} else {
			log.Printf("Seeded voucher: %s", voucher.VoucherName)
		}
	}

	return nil
}
