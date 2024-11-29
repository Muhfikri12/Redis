package main

import (
	"log"
	"voucher_system/database"
	"voucher_system/infra"
	"voucher_system/router"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}
	log.Println("Starting migration...")
	database.Migrate(ctx.DB)
	log.Println("Migration completed successfully.")

	log.Println("Starting seeding...")
	database.SeedAll(ctx.DB)
	log.Println("Seeding completed successfully.")

	r := router.NewRoutes(*ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
