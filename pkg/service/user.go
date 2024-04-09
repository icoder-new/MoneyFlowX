package service

import (
	"github.com/google/uuid"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/pkg/repository"
	"github.com/icoder-new/MoneyFlowX/utils/CustomError"
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
		return &models.User{}, &CustomError.IncorrectCredentialsError{}
	}

	user, err := u.repo.User.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID != "" {
		return user, &CustomError.UserAlreadyExistsError{}
	}

	user, err = u.repo.User.FindByUsernameName(input.Username)
	if err != nil {
		return user, err
	}

	if user.ID != "" {
		return user, &CustomError.UserAlreadyExistsError{}
	}

	user.ID = uuid.New().String()
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

	return u.repo.CreateUser(user)
}
