package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Manage ManagementVoucherInterface
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Manage: NewManagementVoucherRepo(db, log),
	}
}
