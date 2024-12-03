package main

import (
	"log"
	"voucher_system/database"
	"voucher_system/infra"
	"voucher_system/router"

	_ "voucher_system/docs"
)

// @title Sistem Voucher
// @version 1.0
// @description This is API for system voucher.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey id_key
// @in header
// @name id_key
// @securityDefinitions.apikey token
// @in header
// @name token

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
