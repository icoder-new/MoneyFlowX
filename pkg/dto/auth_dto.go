package dto

import (
	"github.com/icoder-new/MoneyFlowX/models"
	"time"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type RegisterRequestBody struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponseBody struct {
	ID           string    `json:"uuid"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	WalletNumber string    `json:"wallet_number"`
	Type         string    `json:"type"`
	Token        string    `json:"token,omitempty"`
	CreatedAt    time.Time `json:"created_at" time_format:"2006-01-02"`
}

func FormatLogin(user *models.User, wallet *models.Wallet, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:           user.ID,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Username:     user.Username,
		Email:        user.Email,
		WalletNumber: wallet.Number,
		Type:         user.Type,
		CreatedAt:    user.CreatedAt,
		Token:        token,
	}
}

func FormatForgotPassword(passwordReset *models.PasswordReset) ForgotPasswordResponseBody {
	return ForgotPasswordResponseBody{
		Email: passwordReset.User.Email,
		Token: passwordReset.Token,
	}
}
