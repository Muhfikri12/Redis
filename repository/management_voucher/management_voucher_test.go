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

func TestUpdateVoucher(t *testing.T) {
	db, mock := setupTestDB()
	defer func() { _ = mock.ExpectationsWereMet() }()

	log := *zap.NewNop()
	voucherRepo := managementvoucher.NewManagementVoucherRepo(db, &log)

	t.Run("Successfully update a voucher", func(t *testing.T) {
		voucherID := 1
		voucher := &models.Voucher{
			VoucherName:     "Promo Updated",
			VoucherCode:     "UPDATED2024",
			VoucherType:     "e-commerce",
			PointsRequired:  10,
			Description:     "Updated discount",
			VoucherCategory: "discount",
			DiscountValue:   15,
			MinimumPurchase: 250000,
			Quota:           50,
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "vouchers" SET`).
			WithArgs(
				voucher.VoucherName,
				voucher.VoucherCode,
				voucher.VoucherType,
				voucher.PointsRequired,
				voucher.Description,
				voucher.VoucherCategory,
				voucher.DiscountValue,
				voucher.MinimumPurchase,
				voucher.Quota,
				sqlmock.AnyArg(), // updated_at
				voucherID,
			).
			WillReturnResult(sqlmock.NewResult(1, int64(voucherID)))
		mock.ExpectCommit()

		err := voucherRepo.UpdateVoucher(voucher, voucherID)
		assert.NoError(t, err)
	})

	t.Run("Failed to update due to no matching record", func(t *testing.T) {
		voucherID := 2
		voucher := &models.Voucher{
			VoucherName: "Promo Not Found",
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "vouchers" SET`).
			WithArgs(voucher.VoucherName, sqlmock.AnyArg(), voucherID).
			WillReturnResult(sqlmock.NewResult(0, 0)) // No rows affected
		mock.ExpectCommit()

		err := voucherRepo.UpdateVoucher(voucher, voucherID)
		assert.Error(t, err)
		assert.EqualError(t, err, "no record found with shipping_id 2")
	})

	t.Run("Failed to update due to database error", func(t *testing.T) {
		voucherID := 3
		voucher := &models.Voucher{
			VoucherName: "Promo Error",
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "vouchers" SET`).
			WithArgs(voucher.VoucherName, sqlmock.AnyArg(), voucherID).
			WillReturnError(fmt.Errorf("database error"))
		mock.ExpectRollback()

		err := voucherRepo.UpdateVoucher(voucher, voucherID)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}

func TestShowRedeemPoints(t *testing.T) {
	db, mock := setupTestDB()
	defer func() { _ = mock.ExpectationsWereMet() }()

	log := *zap.NewNop()
	voucherRepo := managementvoucher.NewManagementVoucherRepo(db, &log)

	t.Run("Successfully show redeem points", func(t *testing.T) {

		mockRows := sqlmock.NewRows([]string{"voucher_name", "discount_value", "points_required"}).
			AddRow("Promo A", 20.0, 50).
			AddRow("Promo B", 15.0, 30)

		mock.ExpectQuery(`SELECT v.voucher_name, v.discount_value, v.points_required FROM vouchers as v WHERE`).
			WithArgs("redeem points").
			WillReturnRows(mockRows)

		result, err := voucherRepo.ShowRedeemPoints()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, *result, 2)

		assert.Equal(t, "Promo A", (*result)[0].VoucherName)
		assert.Equal(t, 20.0, (*result)[0].DiscountValue)
		assert.Equal(t, 50, (*result)[0].PointsRequired)

		assert.Equal(t, "Promo B", (*result)[1].VoucherName)
		assert.Equal(t, 15.0, (*result)[1].DiscountValue)
		assert.Equal(t, 30, (*result)[1].PointsRequired)
	})

	t.Run("Failed to show redeem points due to database error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT v.voucher_name, v.discount_value, v.points_required FROM vouchers as v WHERE`).
			WithArgs("redeem points").
			WillReturnError(fmt.Errorf("database error"))

		result, err := voucherRepo.ShowRedeemPoints()
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")
	})
}
