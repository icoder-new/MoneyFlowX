package service

import (
	"errors"
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/dto"
	"fr33d0mz/moneyflowx/pkg/repository"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUser(input *dto.UserRequestParams) (*models.User, error) {
	user, err := u.repo.User.FindById(input.UserID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserService) CreateUser(input *dto.RegisterRequestBody) (*models.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &models.User{}, errors.New("not valid email")
	}

	user, err := u.repo.User.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	user, err = u.repo.User.FindByUsernameName(input.Username)
	if err != nil {
		return user, err
	}

	user.Firstname = input.Firstname
	user.Lastname = input.Lastname
	user.Email = input.Email
	user.Username = input.Username
	user.IsActive = true
	user.Type = "unidentified" // "identified" after validating

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	_user, err := u.repo.CreateUser(user)
	if err != nil {
		return user, err
	}

	return _user, nil
}
