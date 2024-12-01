package managementvoucherhandler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	managementvoucherhandler "voucher_system/controller/management_voucher_handler"
	"voucher_system/service"
	managementvoucherservice "voucher_system/service/management_voucher_service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// func TestCreateVoucher(t *testing.T) {
// 	log := *zap.NewNop()
//
// 	mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 	service := service.Service{
// 		Manage: mockService,
// 	}
//
// 	handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 	r := gin.Default()
// 	r.POST("/voucher", handler.CreateVoucher)
// 	t.Run("Invalid JSON payload", func(t *testing.T) {
//
// 		r := gin.Default()
// 		r.POST("/voucher", handler.CreateVoucher)
// 		// Prepare invalid JSON payload (e.g., incomplete JSON)
// 		req := httptest.NewRequest(http.MethodPost, "/voucher", strings.NewReader(`{"voucher_name": 123}`))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		// Call handler to process the request
// 		r.ServeHTTP(w, req)
//
// 		// Assert response code is 500 (Internal Server Error)
// 		assert.Equal(t, http.StatusInternalServerError, w.Code) // Match with the actual response in handler
// 	})
//
// 	// Test case when CreateVoucher fails
// 	t.Run("Fail to create voucher", func(t *testing.T) {
//
// 		r := gin.Default()
// 		r.POST("/voucher", handler.CreateVoucher)
// 		voucher := &models.Voucher{
// 			VoucherName:     "Promo December",
// 			VoucherCode:     "DESC2024",
// 			VoucherType:     "e-commerce",
// 			PointsRequired:  0,
// 			Description:     "Get more discount on december",
// 			VoucherCategory: "discount",
// 			DiscountValue:   10,
// 			MinimumPurchase: 200000,
// 			PaymentMethods:  []string{"Credit Card", "PayPal"},
// 			StartDate:       time.Now().AddDate(0, 0, -5),
// 			EndDate:         time.Now().AddDate(0, 0, -1),
// 			ApplicableAreas: []string{"US", "Canada"},
// 			Quota:           100,
// 			Status:          false,
// 			CreatedAt:       time.Now().AddDate(0, 0, 1),
// 			UpdatedAt:       time.Now().AddDate(0, 0, -1),
// 		}
//
// 		// Set up mock for CreateVoucher to fail
// 		mockService.On("CreateVoucher", voucher).Return(fmt.Errorf("failed to create voucher"))
//
// 		// Create a request with valid payload
// 		req := httptest.NewRequest(http.MethodPost, "/voucher", strings.NewReader(`{
// 			"voucher_name": "Discount 10%",
// 			"quota": 100,
// 			"points_required": 50,
// 			"start_date": "2024-11-01",
// 			"end_date": "2024-12-01"
// 		}`))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		// Call handler to process the request
// 		r.ServeHTTP(w, req)
//
// 		// Assert response code is 400 (Bad Request) based on actual handler behavior
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})
//
// 	// Test case when CreateVoucher succeeds
// 	t.Run("Successfully create voucher", func(t *testing.T) {
//
// 		// Creating a new voucher object with all required fields
// 		voucher := &models.Voucher{
// 			VoucherName:     "Promo December",
// 			VoucherCode:     "DESC2024",
// 			VoucherType:     "e-commerce",
// 			PointsRequired:  0,
// 			Description:     "Get more discount on december",
// 			VoucherCategory: "discount",
// 			DiscountValue:   10,
// 			MinimumPurchase: 200000,
// 			PaymentMethods:  []string{"Credit Card", "PayPal"},
// 			StartDate:       time.Now().AddDate(0, 0, -5),
// 			EndDate:         time.Now().AddDate(0, 0, -1),
// 			ApplicableAreas: []string{"US", "Canada"},
// 			Quota:           100,
// 			Status:          false,
// 			CreatedAt:       time.Now().AddDate(0, 0, 1),
// 			UpdatedAt:       time.Now().AddDate(0, 0, -1),
// 		}
//
// 		// Set up mock for CreateVoucher to succeed
// 		mockService.On("CreateVoucher", voucher).Return(nil)
//
// 		// Create a request with valid payload (matching the fields in the Voucher struct)
// 		req := httptest.NewRequest(http.MethodPost, "/voucher", strings.NewReader(`{
//         "voucher_name": "Promo December",
//         "voucher_code": "DESC2024",
//         "voucher_type": "e-commerce",
//         "points_required": 0,
//         "description": "Get more discount on december",
//         "voucher_category": "discount",
//         "discount_value": 10,
//         "minimum_purchase": 200000,
//         "payment_methods": ["Credit Card", "PayPal"],
//         "start_date": "2024-11-01T00:00:00Z",
//         "end_date": "2024-11-30T23:59:59Z",
//         "applicable_areas": ["US", "Canada"],
//         "quota": 100
//     }`))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		// Call handler to process the request
// 		r.ServeHTTP(w, req)
//
// 		// Assert response code is 200 (OK)
// 		assert.Equal(t, http.StatusOK, w.Code)
//
// 		// Ensure CreateVoucher was called with the correct voucher
// 		mockService.AssertCalled(t, "CreateVoucher", voucher)
// 	})
//
// }

func TestSoftDeleteVoucher(t *testing.T) {
	// Create a new No-op logger and mock service for each test
	log := *zap.NewNop()

	t.Run("Successfully delete voucher", func(t *testing.T) {
		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
		service := service.Service{
			Manage: mockService,
		}
		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)

		r := gin.Default() // Always create a new router for each test
		r.DELETE("/voucher/:id", handler.SoftDeleteVoucher)

		voucherID := 123

		// Mock the service call to SoftDeleteVoucher to return no error (successful deletion)
		mockService.On("SoftDeleteVoucher", voucherID).Return(nil)

		// Create a request to delete the voucher with ID 123
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/voucher/%d", voucherID), nil)
		w := httptest.NewRecorder()

		// Call the handler
		r.ServeHTTP(w, req)

		// Assert the response code is 200 (OK)
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert that the service method was called with the correct voucher ID
		mockService.AssertCalled(t, "SoftDeleteVoucher", voucherID)

		// Check the response body
		expectedResponse := `{"status":true,"data":123,"message":"Deleted succesfully"}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	t.Run("Fail to delete voucher due to service error", func(t *testing.T) {
		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
		service := service.Service{
			Manage: mockService,
		}
		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)

		r := gin.Default() // Always create a new router for each test
		r.DELETE("/voucher/:id", handler.SoftDeleteVoucher)

		voucherID := 123

		// Mock the service call to SoftDeleteVoucher to return an error (deletion failed)
		mockService.On("SoftDeleteVoucher", voucherID).Return(fmt.Errorf("failed to delete voucher"))

		// Create a request to delete the voucher with ID 123
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/voucher/%d", voucherID), nil)
		w := httptest.NewRecorder()

		// Call the handler
		r.ServeHTTP(w, req)

		// Assert the response code is 500 (Internal Server Error)
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Assert that the service method was called with the correct voucher ID
		mockService.AssertCalled(t, "SoftDeleteVoucher", voucherID)

		// Check the response body
		expectedResponse := `{"error_msg":"FAILED", "message":"Failed to deleted Voucher", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}

// func TestUpdateVoucher(t *testing.T) {
//
// 	log := *zap.NewNop()
//
// 	t.Run("Successfully update voucher", func(t *testing.T) {
// 		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 		service := service.Service{
// 			Manage: mockService,
// 		}
// 		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 		r := gin.Default() // Always create a new router for each test
// 		r.PUT("/voucher/:id", handler.UpdateVoucher)
//
// 		voucherID := 123
// 		voucher := models.Voucher{
// 			VoucherName:    "Promo December",
// 			VoucherCode:    "DESC2024",
// 			VoucherType:    "e-commerce",
// 			PointsRequired: 0,
// 			Description:    "Get more discount in december",
// 			Quota:          100,
// 		}
//
// 		mockService.On("UpdateVoucher", &voucher, voucherID).Return(nil)
//
// 		reqBody := `{
// 			"voucher_name": "Promo December",
// 			"voucher_code": "DESC2024",
// 			"voucher_type": "e-commerce",
// 			"points_required": 0,
// 			"description": "Get more discount in december",
// 			"quota": 100
// 		}`
// 		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/voucher/%d", voucherID), strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		r.ServeHTTP(w, req)
//
// 		assert.Equal(t, http.StatusOK, w.Code)
//
// 		mockService.AssertCalled(t, "UpdateVoucher", &voucher, voucherID)
//
// 		expectedResponse := `{"status":true,"data":123,"message":"updated succesfully"}`
// 		assert.JSONEq(t, expectedResponse, w.Body.String())
// 	})
//
// 	t.Run("Fail to update voucher due to invalid payload", func(t *testing.T) {
// 		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 		service := service.Service{
// 			Manage: mockService,
// 		}
// 		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 		r := gin.Default()
// 		r.PUT("/voucher/:id", handler.UpdateVoucher)
//
// 		voucherID := 123
//
// 		reqBody := `{
// 			"voucher_name": "Promo December",
// 			"voucher_code": "DESC2024",
// 			"voucher_type": "e-commerce"
// 			// Missing other fields
// 		}`
// 		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/voucher/%d", voucherID), strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		// Call the handler
// 		r.ServeHTTP(w, req)
//
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
//
// 		// Assert the service method was not called due to invalid payload
// 		mockService.AssertNotCalled(t, "UpdateVoucher", mock.Anything, mock.Anything)
//
// 		// Check the response body
// 		expectedResponse := `{"error_msg":"INVALID", "message":"Invalid Payloadinvalid character '/' after object key:value pair", "status":false}`
// 		assert.JSONEq(t, expectedResponse, w.Body.String())
// 	})
//
// 	t.Run("Fail to update voucher due to service error", func(t *testing.T) {
// 		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 		service := service.Service{
// 			Manage: mockService,
// 		}
// 		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 		r := gin.Default() // Always create a new router for each test
// 		r.PUT("/voucher/:id", handler.UpdateVoucher)
//
// 		voucherID := 123
// 		voucher := models.Voucher{
// 			VoucherName:    "Promo December",
// 			VoucherCode:    "DESC2024",
// 			VoucherType:    "e-commerce",
// 			PointsRequired: 0,
// 			Description:    "Get more discount in december",
// 			Quota:          100,
// 		}
//
// 		// Mock the service call to UpdateVoucher to return an error (update failed)
// 		mockService.On("UpdateVoucher", &voucher, voucherID).Return(fmt.Errorf("failed to update voucher"))
//
// 		// Create a request to update the voucher with ID 123
// 		reqBody := `{
// 			"voucher_name": "Promo December",
// 			"voucher_code": "DESC2024",
// 			"voucher_type": "e-commerce",
// 			"points_required": 0,
// 			"description": "Get more discount in december",
// 			"quota": 100
// 		}`
// 		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/voucher/%d", voucherID), strings.NewReader(reqBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
//
// 		// Call the handler
// 		r.ServeHTTP(w, req)
//
// 		// Assert the response code is 500 (Internal Server Error)
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
// 		// Assert that the service method was called with the correct voucher ID and payload
// 		mockService.AssertCalled(t, "UpdateVoucher", &voucher, voucherID)
//
// 		// Check the response body
// 		expectedResponse := `{"error_msg":"FAILED", "message":"Failed to Updated Voucher", "status":false}`
// 		assert.JSONEq(t, expectedResponse, w.Body.String())
// 	})
// }

// func TestShowRedeemPoints(t *testing.T) {
// 	// Create a new No-op logger and mock service for each test
// 	log := *zap.NewNop()
//
// 	t.Run("Successfully retrieve redeem points", func(t *testing.T) {
// 		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 		service := service.Service{
// 			Manage: mockService,
// 		}
// 		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 		r := gin.Default() // Always create a new router for each test
// 		r.GET("/vouchers/reedem-points", handler.ShowRedeemPoints)
//
// 		redeemPoints := []managementvoucher.RedeemPoint{
// 			{VoucherName: "Promo December", PointsRequired: 100, DiscountValue: 10.0},
// 			{VoucherName: "Black Friday", PointsRequired: 200, DiscountValue: 20.0},
// 		}
//
// 		// Mock the service method to return redeem points
// 		mockService.On("ShowRedeemPoints").Return(&redeemPoints, nil)
//
// 		req := httptest.NewRequest(http.MethodGet, "/redeem-points", nil)
// 		w := httptest.NewRecorder()
//
// 		r.ServeHTTP(w, req)
//
// 		assert.Equal(t, http.StatusOK, w.Code)
//
// 		// Validate the response body matches the expected structure
// 		expectedResponse := `{
//         "status": true,
//         "data": [
//             {"voucher_name": "Promo December", "points_required": 100, "discount_value": 10},
//             {"voucher_name": "Black Friday", "points_required": 200, "discount_value": 20}
//         ],
//         "message": "Redeem points retrieved successfully"
//     }`
// 		assert.JSONEq(t, expectedResponse, w.Body.String())
//
// 		// Ensure the service method was called with no arguments
// 		mockService.AssertExpectations(t)
// 	})
//
// 	t.Run("Fail to retrieve redeem points due to service error", func(t *testing.T) {
// 		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
// 		service := service.Service{
// 			Manage: mockService,
// 		}
// 		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)
//
// 		r := gin.Default() // Always create a new router for each test
// 		r.GET("/redeem-points", handler.ShowRedeemPoints)
//
// 		// Mock the service call to ShowRedeemPoints to return an error
// 		mockService.On("ShowRedeemPoints").Return(nil, fmt.Errorf("failed to retrieve redeem points"))
//
// 		// Create a request to retrieve redeem points
// 		req := httptest.NewRequest(http.MethodGet, "/redeem-points", nil)
// 		w := httptest.NewRecorder()
//
// 		// Call the handler
// 		r.ServeHTTP(w, req)
//
// 		// Assert the response code is 404 (Not Found)
// 		assert.Equal(t, http.StatusNotFound, w.Code)
//
// 		// Assert that the service method was called to retrieve redeem points
// 		mockService.AssertCalled(t, "ShowRedeemPoints")
//
// 		// Check the response body
// 		expectedResponse := `{"error_msg":"NOT FOUND", "message":"Reedem Points List Not Found", "status":false}`
// 		assert.JSONEq(t, expectedResponse, w.Body.String())
// 	})
// }
