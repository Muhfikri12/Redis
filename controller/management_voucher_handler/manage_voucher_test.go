package managementvoucherhandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	managementvoucherhandler "voucher_system/controller/management_voucher_handler"
	"voucher_system/models"
	"voucher_system/service"
	managementvoucherservice "voucher_system/service/management_voucher_service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

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

func TestCreateVoucher(t *testing.T) {
	// Setup logger
	log := *zap.NewNop()

	t.Run("Successfully create voucher", func(t *testing.T) {
		// Mock Service
		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
		service := service.Service{
			Manage: mockService,
		}
		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)

		// Router and Endpoint
		r := gin.Default()
		r.POST("/vouchers", handler.CreateVoucher)

		// Mock Data
		mockVoucher := models.Voucher{
			VoucherName:     "Test Voucher",
			VoucherCode:     "TEST123",
			VoucherType:     "e-commerce",
			VoucherCategory: "Discount",
			DiscountValue:   10.0,
			MinimumPurchase: 100.0,
			PaymentMethods:  []string{"Credit Card"},
			StartDate:       time.Now().Round(0),
			EndDate:         time.Now().AddDate(0, 1, 0).Round(0),
			ApplicableAreas: []string{"Jawa"},
			Quota:           50,
		}

		// Mock Response
		mockService.On("CreateVoucher", &mockVoucher).Return(nil)

		// Create Request Body
		body, _ := json.Marshal(mockVoucher)

		// Perform Request
		req := httptest.NewRequest(http.MethodPost, "/vouchers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Call the Handler
		r.ServeHTTP(w, req)

		// Assert the Response
		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertCalled(t, "CreateVoucher", &mockVoucher)

		// Assert the JSON Response Body
		expectedResponse := map[string]interface{}{
			"status":  true,
			"message": "Created succesfully",
			"data": map[string]interface{}{
				"voucher_name":     "Test Voucher",
				"voucher_code":     "TEST123",
				"voucher_type":     "e-commerce",
				"voucher_category": "Discount",
				"discount_value":   10.0,
				"minimum_purchase": 100.0,
				"payment_methods":  []interface{}{"Credit Card"},
				"start_date":       mockVoucher.StartDate.Format(time.RFC3339Nano),
				"end_date":         mockVoucher.EndDate.Format(time.RFC3339Nano),
				"applicable_areas": []interface{}{"Jawa"},
				"quota":            float64(mockVoucher.Quota),
				"created_at":       "0001-01-01T00:00:00Z",
				"updated_at":       "0001-01-01T00:00:00Z",
			},
		}

		var actualResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, actualResponse)
	})

	t.Run("Fail to create voucher due to service error", func(t *testing.T) {
		// Mock Service
		mockService := &managementvoucherservice.ManagementVoucherServiceMock{}
		service := service.Service{
			Manage: mockService,
		}
		handler := managementvoucherhandler.NewManagementVoucherHanlder(service, &log)

		// Router and Endpoint
		r := gin.Default()
		r.POST("/vouchers", handler.CreateVoucher)

		// Mock Data
		mockVoucher := models.Voucher{
			VoucherName:     "Test Voucher",
			VoucherCode:     "TEST123",
			VoucherType:     "e-commerce",
			VoucherCategory: "Discount",
			DiscountValue:   10.0,
			MinimumPurchase: 100.0,
			PaymentMethods:  []string{"Credit Card"},
			StartDate:       time.Now().Round(0),
			EndDate:         time.Now().AddDate(0, 1, 0).Round(0),
			ApplicableAreas: []string{"Jawa"},
			Quota:           50,
		}

		// Mock Response
		mockService.On("CreateVoucher", &mockVoucher).Return(fmt.Errorf("failed to create voucher"))

		// Create Request Body
		body, _ := json.Marshal(mockVoucher)

		// Perform Request
		req := httptest.NewRequest(http.MethodPost, "/vouchers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Call the Handler
		r.ServeHTTP(w, req)

		// Assert the Response
		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertCalled(t, "CreateVoucher", &mockVoucher)

		// Assert the JSON Response Body
		expectedResponse := `{"error_msg":"FAILED", "message":"Failed to create Voucher", "status":false}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
