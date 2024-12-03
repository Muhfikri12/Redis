package authservice

import (
	"voucher_system/models"
	"voucher_system/repository"

	"go.uber.org/zap"
)

type AuthServiceInterface interface {
	Login(login *models.Login) (*models.Session, string, error)
}

type authservice struct {
	repo repository.Repository
	log  *zap.Logger
}

func NewManagementVoucherService(repo repository.Repository, log *zap.Logger) AuthServiceInterface {
	return &authservice{repo: repo, log: log}
}

func (as *authservice) Login(login *models.Login) (*models.Session, string, error) {

	session, idKey, err := as.repo.Auth.Login(login)
	if err != nil {
		return nil, "", err
	}

	return session, idKey, nil
}
