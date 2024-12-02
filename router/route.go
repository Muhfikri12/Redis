package router

import (
	"voucher_system/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	router := r.Group("/vouchers")
	{
		router.POST("/create", ctx.Ctl.Manage.CreateVoucher)
		router.DELETE("/:id", ctx.Ctl.Manage.SoftDeleteVoucher)
		router.PUT("/:id", ctx.Ctl.Manage.UpdateVoucher)
		router.GET("/reedem-points", ctx.Ctl.Manage.ShowRedeemPoints)
		router.GET("/", ctx.Ctl.Manage.GetVouchersByQueryParams)
		router.POST("/redeem", ctx.Ctl.Manage.CreateRedeemVoucher)
		router.GET("/:user_id", ctx.Ctl.Voucher.FindVouchers)
		router.GET("/:user_id/validate", ctx.Ctl.Voucher.ValidateVoucher)
		router.POST("/", ctx.Ctl.Voucher.UseVoucher)
		router.GET("/redeem-history/:user_id", ctx.Ctl.Voucher.GetRedeemHistoryByUser)
		router.GET("/usage-history/:user_id", ctx.Ctl.Voucher.GetUsageHistoryByUser)
		router.GET("/users-by-voucher/:voucher_code", ctx.Ctl.Voucher.GetUsersByVoucherCode)

	}

	return r
}
