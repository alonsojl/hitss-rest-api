package token

import (
	"errors"
	"hitss/internal/model"
	"hitss/pkg/helper/cert"
	"hitss/pkg/helper/logger"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

func Generate(user model.User) (string, error) {
	claims := customClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "Hitss",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := jwtToken.SignedString(cert.PrivateKey)
	if err != nil {
		logger.Write(err)
		return "", err
	}
	return token, nil
}

func Validate(token string) (model.User, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cert.PublicKey, nil
	})

	if !jwtToken.Valid {
		return model.User{}, err
	}

	claims, ok := jwtToken.Claims.(*customClaims)
	if !ok {
		return model.User{}, errors.New("data could not be obtained")
	}
	return claims.User, nil
}

func CheckError(err error) (string, int) {
	if v, ok := err.(*jwt.ValidationError); ok && v.Errors == jwt.ValidationErrorExpired {
		return "expired token", http.StatusForbidden
	}
	return "failed token", http.StatusNotAcceptable
}
