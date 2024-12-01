package managementvoucher_test

import (
	"testing"
	"time"
	"voucher_system/models"
	managementvoucher "voucher_system/repository/management_voucher"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	return gormDB, mock
}

func TestCreateVoucher(t *testing.T) {
	db, mock := setupTestDB()
	defer func() { _ = mock.ExpectationsWereMet() }()

	log := *zap.NewNop()

	customerRepo := managementvoucher.NewManagementVoucherRepo(db, &log)

	t.Run("Succesfully create a voucher", func(t *testing.T) {

		voucher := &models.Voucher{
			VoucherName:     "Promo December",
			VoucherCode:     "DESC2024",
			VoucherType:     "e-commerce",
			PointsRequired:  0,
			Description:     "Get more discount on december",
			VoucherCategory: "discount",
			DiscountValue:   10,
			MinimumPurchase: 200000,
			PaymentMethods:  []string{"Credit Card", "PayPal"},
			StartDate:       time.Now().AddDate(0, 0, -5), // StartDate 5 days ago
			EndDate:         time.Now().AddDate(0, 0, -1), // EndDate 1 day ago
			ApplicableAreas: []string{"US", "Canada"},
			Quota:           100,
			Status:          false,
			CreatedAt:       time.Now().AddDate(0, 0, 1),
			UpdatedAt:       time.Now().AddDate(0, 0, -1),
		}

		mock.ExpectBegin()
		mock.ExpectQuery(`(?i)INSERT INTO "vouchers"`).
			WithArgs(
				voucher.VoucherName,     // Argumen 1
				voucher.VoucherCode,     // Argumen 2
				voucher.VoucherType,     // Argumen 3
				voucher.PointsRequired,  // Argumen 4
				voucher.Description,     // Argumen 5
				voucher.VoucherCategory, // Argumen 6
				voucher.DiscountValue,   // Argumen 7
				voucher.MinimumPurchase, // Argumen 8
				sqlmock.AnyArg(),        // Argumen 9: PaymentMethods (JSON serialized)
				sqlmock.AnyArg(),        // Argumen 10: StartDate (time.Time formatted)
				sqlmock.AnyArg(),        // Argumen 11: EndDate (time.Time formatted)
				sqlmock.AnyArg(),        // Argumen 12: ApplicableAreas (JSON serialized)
				voucher.Quota,           // Argumen 13
				voucher.Status,          // Argumen 14
				sqlmock.AnyArg(),        // Argumen 15: CreatedAt (time.Time formatted)
				sqlmock.AnyArg(),        // Argumen 16: UpdatedAt (time.Time formatted)
				sqlmock.AnyArg(),        // Argumen 17: DeletedAt (NULL)
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectCommit()
		err := customerRepo.CreateVoucher(voucher)
		assert.NoError(t, err)
		assert.Equal(t, 1, voucher.ID)
		assert.NotEmpty(t, voucher.VoucherName)
	})

}
