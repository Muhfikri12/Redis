package router

import (
	"voucher_system/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.POST("/login", ctx.Ctl.Auth.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := r.Group("/vouchers")
	{
		router.Use(ctx.Middleware.Authentication())
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
