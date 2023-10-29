package service

import (
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/utils"
	"fr33d0mz/moneyflowx/utils/CustomError"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) Attempt(input *dto.LoginRequestBody) (*models.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &models.User{}, &CustomError.UserNotFoundError{}
	}

	user, err := a.repo.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, &CustomError.UserNotFoundError{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, &CustomError.IncorrectCredentialsError{}
	}

	return user, nil
}

func (a *AuthService) ForgotPass(input *dto.ForgotPasswordRequestBody) (*models.PasswordReset, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &models.PasswordReset{}, &CustomError.NotValidEmailError{}
	}

	user, err := a.repo.FindByEmail(input.Email)
	if err != nil {
		return &models.PasswordReset{}, err
	}

	if user.ID == "" {
		return &models.PasswordReset{}, &CustomError.UserNotFoundError{}
	}

	passwordReset, err := a.repo.PasswordReset.FindByUserId(user.ID)
	if err != nil {
		return &models.PasswordReset{}, err
	}

	passwordReset.UserID = user.ID
	passwordReset.Token = utils.GenerateString(10)
	passwordReset.ExpiredAt = time.Now().Add(time.Minute * 15)

	passwordReset, err = a.repo.PasswordReset.Save(passwordReset)
	passwordReset.User = *user

	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (a *AuthService) ResetPass(input *dto.ResetPasswordRequestBody) (*models.PasswordReset, error) {
	passwordReset, err := a.repo.PasswordReset.FindByToken(input.Token)
	if err != nil {
		return passwordReset, err
	}

	if passwordReset.User.Email == "" {
		return passwordReset, &CustomError.ResetTokenNotFoundError{}
	}

	if input.Password != input.ConfirmPassword {
		return passwordReset, &CustomError.PasswordNotSameError{}
	}

	user := &passwordReset.User
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return passwordReset, err
	}
	user.Password = string(passwordHash)

	_, err = a.repo.User.Update(user)
	if err != nil {
		return passwordReset, err
	}

	passwordReset, err = a.repo.PasswordReset.Delete(passwordReset)
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}
