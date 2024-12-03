package controller

import (
	"net/http"
	"strconv"
	"voucher_system/helper"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Get Redeem History endpoint
// @Summary Get Redeem History
// @Description Feature Get Redeem History
// @Tags Get Redeem History
// @Accept json
// @Produce json
// @Success 200 {object} helper.HTTPResponse "Success response"
// @Failure 400 {object} helper.HTTPResponse "Bad request error"
// @Failure 500 {object} helper.HTTPResponse "Internal server error"
// @Security token
// @Security id_key
// @Router /vouchers/redeem-history/:user_id [get]
func (c *VoucherController) GetRedeemHistoryByUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		helper.ResponseError(ctx, "Invalid user ID", err.Error(), http.StatusBadRequest)
		return
	}
	redeems, err := c.service.History.GetRedeemHistoryByUser(userID)
	if err != nil {
		c.log.Error("Error fetching redeem history", zap.Error(err))
		helper.ResponseError(ctx, "Failed to fetch history", err.Error(), http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(ctx, gin.H{"redeem_history": redeems}, "Redeem history fetched successfully")
}

// Get Usage Users History endpoint
// @Summary Get Usage Users History
// @Description Feature Get Usage Users History
// @Tags Get Usage Users History
// @Accept json
// @Produce json
// @Success 200 {object} helper.HTTPResponse "Success response with voucher data"
// @Failure 400 {object} helper.HTTPResponse "Bad request error"
// @Failure 500 {object} helper.HTTPResponse "Internal server error"
// @Security token
// @Security id_key
// @Router /vouchers/usage-history/:user_id [get]
func (c *VoucherController) GetUsageHistoryByUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		helper.ResponseError(ctx, "Invalid user ID", err.Error(), http.StatusBadRequest)
		return
	}
	histories, err := c.service.History.GetUsageHistoryByUser(userID)
	if err != nil {
		c.log.Error("Error fetching usage history", zap.Error(err))
		helper.ResponseError(ctx, "Failed to fetch usage history", err.Error(), http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(ctx, gin.H{"usage_history": histories}, "Usage history fetched successfully")
}

// Get User By Voucher Code endpoint
// @Summary Get User By Voucher Code
// @Description Feature Get User By Voucher Code
// @Tags Get User By Voucher Code
// @Accept json
// @Produce json
// @Success 200 {object} helper.HTTPResponse "Success response"
// @Failure 400 {object} helper.HTTPResponse "Bad request error"
// @Failure 500 {object} helper.HTTPResponse "Internal server error"
// @Security token
// @Security id_key
// @Router /vouchers/users-by-voucher/:voucher_code [get]
func (c *VoucherController) GetUsersByVoucherCode(ctx *gin.Context) {
	voucherCode := ctx.Param("voucher_code")
	redeems, err := c.service.History.GetUsersByVoucherCode(voucherCode)
	if err != nil {
		c.log.Error("Error fetching users by voucher", zap.Error(err))
		helper.ResponseError(ctx, "Failed to fetch users", err.Error(), http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(ctx, gin.H{"users": redeems}, "Users fetched successfully")
}
