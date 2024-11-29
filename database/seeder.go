package database

import (

	"log"
	"time"
	"voucher_system/models"


	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf("Transaction rolled back due to panic: %v", r)
		}
	}()

	var count int64
	if err := tx.Model(&models.Voucher{}).Count(&count).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Error checking shipping data: %v", err)
		return
	}

	if count > 0 {
		tx.Rollback()
		log.Println("Seeding skipped, data already exists.")
		return
	}
	vouchers := []models.Voucher{
		{
			VoucherName:     "10% Discount",
			VoucherCode:     "DISCOUNT10",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "10% off for purchases above $100",
			VoucherCategory: "Discount",
			DiscountValue:   10.0,
			MinimumPurchase: 100.0,
			PaymentMethods:  []string{"Credit Card", "PayPal"},
			StartDate:       time.Now().AddDate(0, 0, -5), // StartDate 5 days ago
			EndDate:         time.Now().AddDate(0, 0, -1), // EndDate 1 day ago
			ApplicableAreas: []string{"US", "Canada"},
			Quota:           100,
		},
		{
			VoucherName:     "Free Shipping",
			VoucherCode:     "FREESHIP50",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Free shipping for orders above $50",
			VoucherCategory: "Free Shipping",
			DiscountValue:   0.0,
			MinimumPurchase: 50.0,
			PaymentMethods:  []string{"All"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 2, 0), // 2 months valid
			ApplicableAreas: []string{"Worldwide"},
			Quota:           200,
		},
		{
			VoucherName:     "Redeem 500 Points",
			VoucherCode:     "POINTS500",
			VoucherType:     "redeem points",
			PointsRequired:  500,
			Description:     "Redeem 500 points for a $20 discount",
			VoucherCategory: "Discount",
			DiscountValue:   20.0,
			MinimumPurchase: 0.0,
			PaymentMethods:  []string{"Credit Card"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 3, 0), // 3 months valid
			ApplicableAreas: []string{"US"},
			Quota:           150,
		},
		{
			VoucherName:     "5% Discount",
			VoucherCode:     "DISCOUNT5",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "5% discount on all purchases",
			VoucherCategory: "Discount",
			DiscountValue:   5.0,
			MinimumPurchase: 0.0,
			PaymentMethods:  []string{"PayPal"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 1, 0), // 1 month valid
			ApplicableAreas: []string{"Europe"},
			Quota:           500,
		},
		{
			VoucherName:     "Black Friday Sale",
			VoucherCode:     "BLACKFRIDAY",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "20% off for Black Friday",
			VoucherCategory: "Discount",
			DiscountValue:   20.0,
			MinimumPurchase: 200.0,
			PaymentMethods:  []string{"Credit Card", "Bank Transfer"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 0, 7), // 1 week valid
			ApplicableAreas: []string{"Worldwide"},
			Quota:           300,
		},
		{
			VoucherName:     "Holiday Free Shipping",
			VoucherCode:     "HOLIDAYSHIP",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Free shipping during the holiday season",
			VoucherCategory: "Free Shipping",
			DiscountValue:   0.0,
			MinimumPurchase: 75.0,
			PaymentMethods:  []string{"All"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 1, 0), // 1 month valid
			ApplicableAreas: []string{"US", "Canada"},
			Quota:           400,
		},
		{
			VoucherName:     "Cyber Monday Special",
			VoucherCode:     "CYBERMON",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "15% off for Cyber Monday",
			VoucherCategory: "Discount",
			DiscountValue:   15.0,
			MinimumPurchase: 150.0,
			PaymentMethods:  []string{"Credit Card"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 0, 5), // 5 days valid
			ApplicableAreas: []string{"Worldwide"},
			Quota:           100,
		},
		{
			VoucherName:     "Student Discount",
			VoucherCode:     "STUDENT15",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "15% discount for students",
			VoucherCategory: "Discount",
			DiscountValue:   15.0,
			MinimumPurchase: 0.0,
			PaymentMethods:  []string{"Credit Card", "PayPal"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 2, 0), // 2 months valid
			ApplicableAreas: []string{"Europe"},
			Quota:           200,
		},
		{
			VoucherName:     "New Year Sale",
			VoucherCode:     "NEWYEAR50",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Flat $50 off for the New Year sale",
			VoucherCategory: "Discount",
			DiscountValue:   50.0,
			MinimumPurchase: 300.0,
			PaymentMethods:  []string{"All"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 1, 0), // 1 month valid
			ApplicableAreas: []string{"US"},
			Quota:           150,
		},
		{
			VoucherName:     "Valentine's Free Shipping",
			VoucherCode:     "VALSHIP",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Free shipping for Valentine's Day",
			VoucherCategory: "Free Shipping",
			DiscountValue:   0.0,
			MinimumPurchase: 100.0,
			PaymentMethods:  []string{"Credit Card", "PayPal"},
			StartDate:       time.Now(),
			EndDate:         time.Now().AddDate(0, 1, 14), // 1 month 14 days valid
			ApplicableAreas: []string{"Worldwide"},
			Quota:           300,
		},
	}


	for _, voucher := range vouchers {
		currentDate := time.Now()
		if currentDate.After(voucher.EndDate) {
			voucher.Status = false
		} else {
			voucher.Status = currentDate.After(voucher.StartDate) && currentDate.Before(voucher.EndDate)
		}

		if err := tx.Create(&voucher).Error; err != nil {
			log.Printf("Failed to insert voucher %s: %v", voucher.VoucherCode, err)
			tx.Rollback()
			return
		}
		log.Printf("Inserted voucher: %s", voucher.VoucherCode)
	}
	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Transaction commit failed: %v", err)
	}
	log.Println("Seeding completed successfully.")

}
