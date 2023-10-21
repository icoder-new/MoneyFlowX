package repository

import (
	"fr33d0mz/moneyflowx/models"

	"gorm.io/gorm"
)

type SourceOfFundRepository struct {
	db *gorm.DB
}

func NewSourceOfFundRepository(db *gorm.DB) *SourceOfFundRepository {
	return &SourceOfFundRepository{
		db: db,
	}
}

func (s *SourceOfFundRepository) FindById(id string) (*models.SourceOfFund, error) {
	var sourceOfFund *models.SourceOfFund

	err := s.db.Where("id = ?", id).Find(&sourceOfFund).Error
	if err != nil {
		return sourceOfFund, err
	}

	return sourceOfFund, nil
}
