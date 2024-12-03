package auth

import (
	"fmt"
	"voucher_system/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepoInterface interface {
	Login(login *models.Login) (*models.Session, string, error)
}

type authRepo struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewManagementVoucherRepo(db *gorm.DB, log *zap.Logger) AuthRepoInterface {
	return &authRepo{DB: db, Log: log}
}

func (a *authRepo) Login(login *models.Login) (*models.Session, string, error) {

	token := uuid.New().String()

	user := models.User{}
	result := a.DB.Where("email = ? AND password = ?", login.Email, login.Password).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, "", fmt.Errorf("invalid email or password")
		}
		return nil, "", result.Error
	}

	session := models.Session{
		UserID: user.ID,
		Token:  token,
	}

	existingSession := models.Session{}
	err := a.DB.Where("user_id = ?", user.ID).First(&existingSession).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, "", err
	}

	if err == gorm.ErrRecordNotFound {

		if err := a.DB.Create(&session).Error; err != nil {
			return nil, "", err
		}
	} else {

		session.ID = existingSession.ID
		if err := a.DB.Save(&session).Error; err != nil {
			return nil, "", err
		}
	}

	return &session, user.Name, nil
}
