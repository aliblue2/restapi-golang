package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const superSecretKey = "alireza-golang-azno-space.com"

func GenerateToken(email string, userId int64) (string, error) {

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return tokenString.SignedString([]byte(superSecretKey))

}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(superSecretKey), nil
	})

	if err != nil {
		return -1, err
	}

	validToken := parsedToken.Valid

	if !validToken {
		return -1, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid token")
	}

	userId := int64(claims["userId"].(float64))

	return userId, err

}
