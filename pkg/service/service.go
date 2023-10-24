package service

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"

	"github.com/golang-jwt/jwt/v4"
)

type Authorization interface {
	Attempt(input *dto.LoginRequestBody) (*models.User, error)
	ForgotPass(input *dto.ForgotPasswordRequestBody) (*models.PasswordReset, error)
	ResetPass(input *dto.ResetPasswordRequestBody) (*models.PasswordReset, error)
}

type JWT interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type User interface {
	GetUser(input *dto.UserRequestParams) (*models.User, error)
	CreateUser(input *dto.RegisterRequestBody) (*models.User, error)
}

type Wallet interface {
	GetWalletByUserId(input *dto.WalletRequestBody) (*models.Wallet, error)
	CreateWallet(input *dto.WalletRequestBody) (*models.Wallet, error)
}

type Transaction interface {
	GetTransactions(userID string, query *dto.TransactionRequestQuery) ([]*models.Transaction, error)
	TopUp(input *dto.TopUpRequestBody) (*models.Transaction, error)
	Transfer(input *dto.TransferRequestBody) (*models.Transaction, error)
	CountTransaction(userID string) (int64, error)
}

type Service struct {
	Authorization
	JWT
	User
	Wallet
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repo),
		JWT:           NewJWTService(*repo),
		User:          NewUserService(*repo),
		Wallet:        NewWalletService(*repo),
	}
}
