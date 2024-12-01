package managementvoucher_test

import (
	"fmt"
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
			StartDate:       time.Now().AddDate(0, 0, -5),
			EndDate:         time.Now().AddDate(0, 0, -1),
			ApplicableAreas: []string{"US", "Canada"},
			Quota:           100,
			Status:          false,
			CreatedAt:       time.Now().AddDate(0, 0, 1),
			UpdatedAt:       time.Now().AddDate(0, 0, -1),
		}

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "vouchers"`).
			WithArgs(
				voucher.VoucherName,
				voucher.VoucherCode,
				voucher.VoucherType,
				voucher.PointsRequired,
				voucher.Description,
				voucher.VoucherCategory,
				voucher.DiscountValue,
				voucher.MinimumPurchase,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				voucher.Quota,
				voucher.Status,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectCommit()
		err := customerRepo.CreateVoucher(voucher)
		assert.NoError(t, err)
		assert.Equal(t, 1, voucher.ID)
		assert.NotEmpty(t, voucher.VoucherName)
	})

	t.Run("Failed to create a voucher", func(t *testing.T) {
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
			StartDate:       time.Now().AddDate(0, 0, -5),
			EndDate:         time.Now().AddDate(0, 0, -1),
			ApplicableAreas: []string{"US", "Canada"},
			Quota:           100,
			Status:          false,
			CreatedAt:       time.Now().AddDate(0, 0, 1),
			UpdatedAt:       time.Now().AddDate(0, 0, -1),
		}

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "vouchers"`).
			WithArgs(
				voucher.VoucherName,
				voucher.VoucherCode,
				voucher.VoucherType,
				voucher.PointsRequired,
				voucher.Description,
				voucher.VoucherCategory,
				voucher.DiscountValue,
				voucher.MinimumPurchase,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				voucher.Quota,
				voucher.Status,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnError(fmt.Errorf("database error"))

		mock.ExpectRollback()
		err := customerRepo.CreateVoucher(voucher)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}

func TestSoftDeleteVoucher(t *testing.T) {
	db, mock := setupTestDB()
	defer func() { _ = mock.ExpectationsWereMet() }()

	log := *zap.NewNop()

	voucherRepo := managementvoucher.NewManagementVoucherRepo(db, &log)

	t.Run("Successfully soft delete a voucher", func(t *testing.T) {
		voucherID := 1

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "vouchers" SET "deleted_at"=`).
			WithArgs(sqlmock.AnyArg(), voucherID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := voucherRepo.SoftDeleteVoucher(voucherID)
		assert.NoError(t, err)
	})

	t.Run("Failed to soft delete a voucher", func(t *testing.T) {
		voucherID := 2

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "vouchers" SET "deleted_at"=`).
			WithArgs(sqlmock.AnyArg(), voucherID).
			WillReturnError(fmt.Errorf("database error"))
		mock.ExpectRollback()

		err := voucherRepo.SoftDeleteVoucher(voucherID)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}
