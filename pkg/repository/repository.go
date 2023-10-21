package repository

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (*models.User, error)
}

type User interface {
	FindAll() ([]*models.User, error)
	FindById(id string) (*models.User, error)
	FindByUsernameName(username string) (*models.User, error)
	FindByName(name string) ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type Wallet interface {
	CreateWallet(wallet *models.Wallet) (*models.Wallet, error)
	FindByUserId(id string) (*models.Wallet, error)
	FindByNumber(number string) (*models.Wallet, error)
	Update(wallet *models.Wallet) (*models.Wallet, error)
}

type SourceOfFund interface {
	FindById(id string) (*models.SourceOfFund, error)
}

type PasswordReset interface {
	FindByUserId(id string) (*models.PasswordReset, error)
	FindByToken(token string) (*models.PasswordReset, error)
	Save(passwordReset *models.PasswordReset) (*models.PasswordReset, error)
	Delete(passwordReset *models.PasswordReset) (*models.PasswordReset, error)
}

type Transaction interface {
	FindAll(userID string, query *dto.TransactionRequestQuery) ([]*models.Transaction, error)
	Count(userID string) (int64, error)
	Save(transaction *models.Transaction) (*models.Transaction, error)
}

type Repository struct {
	Authorization
	User
	Wallet
	SourceOfFund
	PasswordReset
	Transaction
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
