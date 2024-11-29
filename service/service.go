package service

import "voucher_system/repository"

type Service struct {
	// User UserService

}

func NewService(repo repository.Repository) Service {
	return Service{
		// User: NewUserService(repo.User),

	}
}
