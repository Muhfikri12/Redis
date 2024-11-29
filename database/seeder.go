package database

import (
	"fmt"
	"voucher_system/database/seeder"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {

	if err := seeder.SeedVouchers(db); err != nil {
		return fmt.Errorf("failed to seed Shipping data: %w", err)
	}
	return nil
}
