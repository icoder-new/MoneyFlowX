package service

import (
	"fmt"
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
	signedToken, err := token.SignedString(utils.AppSettings.AppParams.SecretKey)
	if err != nil {
		return signedToken, err
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
