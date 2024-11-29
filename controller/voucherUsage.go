package controller

import (
	"net/http"
	"strconv"
	"voucher_system/helper"
	"voucher_system/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoucherController struct {
	service service.Service
	log     *zap.Logger
}

func NewVoucherController(service service.Service, log *zap.Logger) *VoucherController {
	return &VoucherController{service: service, log: log}
}

func (c *VoucherController) GetVoucher(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		helper.ResponseError(ctx, err.Error(), "Invalid user ID", http.StatusBadRequest)
		return
	}
	voucherType := ctx.Query("type")

	voucher, err := c.service.Voucher.FindVouchers(userID, voucherType)
	if err != nil {
		if err.Error() == "no vouchers available" {
			helper.ResponseError(ctx, err.Error(), "", http.StatusBadRequest)
			return
		}
		helper.ResponseError(ctx, err.Error(), "", http.StatusInternalServerError)
		return
	}

	result := gin.H{
		"voucher": voucher,
	}

	helper.ResponseOK(ctx, result, "")
}

func (c *VoucherController) ValidateVoucher(ctx *gin.Context) {
	var request struct {
		VoucherCode       string  `json:"voucher_code"`
		TransactionAmount float64 `json:"transaction_amount"`
		ShippingAmount    float64 `json:"shipping_amount"`
		Area              string  `json:"area"`
		PaymentMethod     string  `json:"payment_method"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		helper.ResponseError(ctx, err.Error(), "Invalid input", http.StatusBadRequest)
		return
	}

	voucher, benefit, err := c.service.Voucher.ValidateVoucher(request.VoucherCode, request.TransactionAmount, request.ShippingAmount, request.Area, request.PaymentMethod)
	if err != nil {
		helper.ResponseError(ctx, err.Error(), "", http.StatusBadRequest)
		return
	}

	result := gin.H{
		"voucher":       voucher,
		"benefit_value": benefit,
		"status":        "valid",
	}
	
	helper.ResponseOK(ctx, result, "")
}

func (c *VoucherController) UseVoucher(ctx *gin.Context) {
	var request struct {
		UserID            int     `json:"user_id"`
		VoucherCode       string  `json:"voucher_code"`
		TransactionAmount float64 `json:"transaction_amount"`
		PaymentMethod     string  `json:"payment_method"`
		Area              string  `json:"area"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := c.service.Voucher.UseVoucher(request.UserID, request.VoucherCode, request.TransactionAmount, request.PaymentMethod, request.Area)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "voucher used successfully",
	})
}
