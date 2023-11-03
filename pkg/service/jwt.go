package service

import (
	"errors"
	"fmt"
	"fr33d0mz/moneyflowx/models"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	repo repository.Repository
}

func NewJWTService(repo repository.Repository) *JWTService {
	return &JWTService{
		repo: repo,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func (j *JWTService) GenerateToken(userID string) (string, error) {
	payload := idTokenClaims{}
	payload.ExpiresAt = &jwt.NumericDate{
		Time: time.Now().Add(time.Minute * time.Duration(utils.AppSettings.AppParams.TokenTTL)),
	}
	payload.IssuedAt = &jwt.NumericDate{
		Time: time.Now(),
	}
	payload.Issuer = utils.AppSettings.AppParams.JWTIssuer
	payload.UserID = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(utils.AppSettings.AppParams.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.AppSettings.AppParams.SecretKey), nil
	})

	if err != nil {
		return _token, err
	}

	return _token, nil
}

func (j *JWTService) SendVerificationToken(user *models.User) error {
	token, err := j.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	data := &utils.EmailData{
		URL:       fmt.Sprintf("%v/verify/%v", utils.GenerateAppURL(), token),
		FirstName: user.Firstname,
		Subject:   "Your account verification link",
	}

	return utils.SendMail(user.Email, data)
}

func (j *JWTService) VerifyUser(token string) (*models.User, *models.Wallet, error) {
	_token, err := j.ValidateToken(token)
	if err != nil {
		return &models.User{}, &models.Wallet{}, err
	}

	payload, ok := _token.Claims.(jwt.MapClaims)
	if !ok || !_token.Valid {
		return &models.User{}, &models.Wallet{}, errors.New("parsing claims error")
	}

	userID := payload["user_id"].(string)

	user, err := j.repo.User.FindById(userID)
	if err != nil {
		return &models.User{}, &models.Wallet{}, err
	}

	user.Type = "identified"

	user, err = j.repo.User.Update(user)
	if err != nil {
		return &models.User{}, &models.Wallet{}, err
	}

	wallet, err := j.repo.Wallet.FindByUserId(userID)
	if err != nil {
		return &models.User{}, &models.Wallet{}, err
	}

	wallet.UserType = user.Type

	wallet, err = j.repo.Wallet.Update(wallet)
	if err != nil {
		return &models.User{}, &models.Wallet{}, err
	}

	return user, wallet, nil
}
