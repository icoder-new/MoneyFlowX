package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/pkg/repository"
)

type Authorization interface {
	Attempt(input *dto.LoginRequestBody) (*models.User, error)
	ForgotPass(input *dto.ForgotPasswordRequestBody) (*models.PasswordReset, error)
	ResetPass(input *dto.ResetPasswordRequestBody) (*models.PasswordReset, error)
}

type JWT interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	SendVerificationToken(user *models.User) error
	VerifyUser(token string) (*models.User, *models.Wallet, error)
}

type User interface {
	GetUser(input *dto.UserRequestParams) (*models.User, error)
	CreateUser(input *dto.RegisterRequestBody) (*models.User, error)
}

type Wallet interface {
	GetWalletByUserId(input *dto.WalletRequestBody) (*models.Wallet, error)
	CreateWallet(input *dto.WalletRequestBody) (*models.Wallet, error)
	GetWalletByNumber(number string) (*models.Wallet, error)
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
		Transaction:   NewTransactionService(*repo),
	}
}
