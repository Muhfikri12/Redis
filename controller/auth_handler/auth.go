package authhandler

import (
	"net/http"
	"voucher_system/database"
	"voucher_system/helper"
	"voucher_system/models"
	"voucher_system/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHadler struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
}

func NewUserHandler(service service.Service, log *zap.Logger, rdb database.Cacher) AuthHadler {
	return AuthHadler{
		Service: service,
		Log:     log,
		Cacher:  rdb,
	}
}

// Login endpoint
// @Summary Login feature
// @Description Login first for all access.
// @Tags login
// @Accept  json
// @Produce  json
// @Param   payload body models.Login true "Login payload"
// @Success 200 {object} helper.HTTPResponse "successfully login"
// @Failure 404 {object} helper.HTTPResponse "User not found"
// @Failure 500 {object} helper.HTTPResponse "Internal server error"
// @Router  /login [post]
func (auth *AuthHadler) Login(c *gin.Context) {
	login := models.Login{}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		auth.Log.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid Payload"+err.Error(), http.StatusInternalServerError)
		return
	}

	session, idKey, err := auth.Service.Auth.Login(&login)
	if err != nil {
		auth.Log.Error("Failed to Login"+err.Error(), zap.Error(err))
		helper.ResponseError(c, "Failed", err.Error(), http.StatusBadRequest)
		return
	}

	token := session.Token
	IDKEY := idKey

	auth.Log.Info("Saving token to Redis", zap.String("IDKEY", IDKEY), zap.String("token", token))

	err = auth.Cacher.Set(IDKEY, token)
	if err != nil {
		helper.ResponseError(c, "server error", err.Error(), http.StatusInternalServerError)
	}

	helper.ResponseOK(c, session, "successfully login")

}
