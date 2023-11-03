package repository

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) FindAll(userID string, query *dto.TransactionRequestQuery) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuilder := t.db.Limit(query.Limit).Offset(offset).Order(orderBy)
	err := queryBuilder.Where("user_id = ?", userID).Where("comment ILIKE ?", "%"+query.Search+"%").
		Preload("User").Preload("Wallet.User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (t *TransactionRepository) Count(userID string) (int64, error) {
	var total int64
	db := t.db.Model(&models.Transaction{}).Where("user_id = ?", userID).Count(&total)

	if db.Error != nil {
		return 0, db.Error
	}

	return total, nil
}

func (t *TransactionRepository) Save(transaction *models.Transaction) (*models.Transaction, error) {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
