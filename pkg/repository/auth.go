package repository

import (
	"fr33d0mz/moneyflowx/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a *AuthPostgres) CreateUser(user *models.User) (*models.User, error) {
	err := a.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
