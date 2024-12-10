package pkg

import (
	"fmt"
	"time"

	"github.com/bwafi/trendora-backend/internal/entity"
	"github.com/bwafi/trendora-backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GenerateAccessToken(customer *entity.Customers, viper *viper.Viper) (string, error) {
	expiry := viper.GetInt("jwt.expAccessToken")
	secretKey := viper.GetString("jwt.accessToken")

	exp := time.Now().Add(time.Minute * time.Duration(expiry))

	claims := &model.JwtCustomClaims{
		Name: customer.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("app.name"),
			Subject:   customer.ID,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(customer *entity.Customers, viper *viper.Viper) (string, error) {
	expiry := viper.GetInt("jwt.expRefreshToken")
	secretKey := viper.GetString("jwt.refreshToken")

	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &model.JwtCustomClaims{
		Name: customer.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("app.name"),
			Subject:   customer.ID,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string, viper *viper.Viper, log *logrus.Logger, secretKey string) (*model.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			log.Warnf("Failed Parse: %+v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.JwtCustomClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}