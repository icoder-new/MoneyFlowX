package repository

import (
	"github.com/icoder-new/MoneyFlowX/models"
	"time"

	"gorm.io/gorm"
)

type PasswordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{
		db: db,
	}
}

func (p *PasswordResetRepository) FindByUserId(id string) (*models.PasswordReset, error) {
	var passwordReset *models.PasswordReset

	err := p.db.Where("user_id = ?", id).Find(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (p *PasswordResetRepository) FindByToken(token string) (*models.PasswordReset, error) {
	var passwordReset *models.PasswordReset

	err := p.db.Where("token = ?", token).Where("expired_at >= ?", time.Now()).Preload("User").Find(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (p *PasswordResetRepository) Save(passwordReset *models.PasswordReset) (*models.PasswordReset, error) {
	err := p.db.Save(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (p *PasswordResetRepository) Delete(passwordReset *models.PasswordReset) (*models.PasswordReset, error) {
	err := p.db.Delete(&passwordReset).Error
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}
