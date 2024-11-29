package repository

import (
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	CreateHistory(history *models.History) error
	FindByUserID(userID int) ([]models.History, error)
}

type historyRepository struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewHistoryRepository(db *gorm.DB, log *zap.Logger) HistoryRepository {
	return &historyRepository{DB: db, log: log}
}

func (r *historyRepository) CreateHistory(history *models.History) error {
	return r.DB.Create(history).Error
}

func (r *historyRepository) FindByUserID(userID int) ([]models.History, error) {
	var histories []models.History
	err := r.DB.Where("user_id = ?", userID).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}
