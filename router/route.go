package router

import (
	"voucher_system/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	voucherRoutes := r.Group("/vouchers")
	{
		voucherRoutes.GET("/:user_id", ctx.Ctl.Voucher.GetVoucher)
		
	}

	return r
}
