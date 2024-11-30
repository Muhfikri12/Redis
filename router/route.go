package router

import (
	"voucher_system/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	voucherRoutes := r.Group("/vouchers")
	{
		voucherRoutes.POST("/", ctx.Ctl.Manage.CreateVoucher)
		voucherRoutes.DELETE("/:id", ctx.Ctl.Manage.SoftDeleteVoucher)
		voucherRoutes.PUT("/:id", ctx.Ctl.Manage.UpdateVoucher)
		voucherRoutes.GET("/reedem-points", ctx.Ctl.Manage.ShowRedeemPoints)
		voucherRoutes.GET("/", ctx.Ctl.Manage.GetVouchersByQueryParams)
		voucherRoutes.POST("/redeem", ctx.Ctl.Manage.CreateRedeemVoucher)

	}
	return r
}
