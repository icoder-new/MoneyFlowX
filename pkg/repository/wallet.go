package repository

import (
	"fr33d0mz/moneyflowx/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (w *WalletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	err := w.db.Create(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (w *WalletRepository) FindByUserId(id string) (*models.Wallet, error) {
	var wallet *models.Wallet

	err := w.db.Where("user_id = ?", id).Find(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (w *WalletRepository) FindByNumber(number string) (*models.Wallet, error) {
	var wallet *models.Wallet

	err := w.db.Where("number = ?", number).Preload("User").Find(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (w *WalletRepository) Update(wallet *models.Wallet) (*models.Wallet, error) {
	err := w.db.Save(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}
