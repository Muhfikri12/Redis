package managementvoucherhandler

import (
	"net/http"
	"strconv"
	"voucher_system/helper"
	"voucher_system/models"
	"voucher_system/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageVoucherHandler interface {
	CreateVoucher(c *gin.Context)
	SoftDeleteVoucher(c *gin.Context)
	UpdateVoucher(c *gin.Context)
	ShowRedeemPoints(c *gin.Context)
	GetVouchersByQueryParams(c *gin.Context)
	CreateRedeemVoucher(c *gin.Context)
}

type ManagementVoucherHandler struct {
	service service.Service
	log     *zap.Logger
}

func NewManagementVoucherHanlder(service service.Service, log *zap.Logger) ManageVoucherHandler {
	return &ManagementVoucherHandler{service: service, log: log}
}

func (mh *ManagementVoucherHandler) CreateVoucher(c *gin.Context) {

	voucher := models.Voucher{}

	err := c.ShouldBindJSON(&voucher)
	if err != nil {
		mh.log.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid Payload"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = mh.service.Manage.CreateVoucher(&voucher)
	if err != nil {
		mh.log.Error("Failed to create", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to create Voucher", http.StatusBadRequest)
		return
	}

	mh.log.Info("Create Voucher successfully")
	helper.ResponseOK(c, voucher, "Created succesfully")
}

func (mh *ManagementVoucherHandler) SoftDeleteVoucher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := mh.service.Manage.SoftDeleteVoucher(id)
	if err != nil {
		mh.log.Error("Failed to Deleted", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to deleted Voucher", http.StatusInternalServerError)
		return
	}

	mh.log.Info("Deleted Voucher successfully")
	helper.ResponseOK(c, id, "Deleted succesfully")
}

func (mh *ManagementVoucherHandler) UpdateVoucher(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	voucher := models.Voucher{}

	if err := c.ShouldBindJSON(&voucher); err != nil {
		mh.log.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid Payload"+err.Error(), http.StatusBadRequest)
		return
	}

	err := mh.service.Manage.UpdateVoucher(&voucher, id)
	if err != nil {
		mh.log.Error("Failed to Updated Voucher", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to Updated Voucher", http.StatusInternalServerError)
		return
	}

	mh.log.Info("Updated Voucher successfully")
	helper.ResponseOK(c, id, "updated succesfully")
}

func (mh *ManagementVoucherHandler) ShowRedeemPoints(c *gin.Context) {

	voucher, err := mh.service.Manage.ShowRedeemPoints()
	if err != nil {
		mh.log.Error("Failed to Get Reedem Points List", zap.Error(err))
		helper.ResponseError(c, "NOT FOUND", "Reedem Points List Not Found", http.StatusNotFound)
		return
	}

	mh.log.Info("Redeem points retrieved successfully")
	helper.ResponseOK(c, voucher, "Redeem points retrieved successfully")

}

func (mh *ManagementVoucherHandler) GetVouchersByQueryParams(c *gin.Context) {

	status := c.Query("status")
	area := c.Query("area")
	voucher_type := c.Query("voucher_type")

	voucher, err := mh.service.Manage.GetVouchersByQueryParams(status, area, voucher_type)
	if err != nil {
		mh.log.Error("Failed to Get Voucher List", zap.Error(err))
		helper.ResponseError(c, "NOT FOUND", "Voucher Not Found", http.StatusNotFound)
		return
	}

	mh.log.Info("Voucher retrieved successfully")
	helper.ResponseOK(c, voucher, "Voucher retrieved successfully")

}

func (mh *ManagementVoucherHandler) CreateRedeemVoucher(c *gin.Context) {
	var payload struct {
		VoucherID int `json:"voucher_id" binding:"required"`
		UserID    int `json:"user_id" binding:"required"`
		Points    int `json:"points" binding:"required"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		mh.log.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid Payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	redeem := models.Redeem{
		VoucherID: payload.VoucherID,
		UserID:    payload.UserID,
	}

	err = mh.service.Manage.CreateRedeemVoucher(&redeem, payload.Points)
	if err != nil {
		mh.log.Error("Failed to create redeem voucher", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to create redeem voucher: "+err.Error(), http.StatusInternalServerError)
		return
	}

	mh.log.Info("Create Redeem Voucher successfully")
	helper.ResponseOK(c, redeem, "Created successfully")
}
