package controller

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
		helper.ResponseError(c, "INVALID", "Invalid Payload"+err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.service.Manage.CreateVoucher(&voucher)
	if err != nil {
		mh.log.Error("Failed to create", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to create Voucher", http.StatusInternalServerError)
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
