package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, time.Time, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(id)),
		ExpiresAt: jwt.NewNumericDate(expiration),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", time.Time{}, err
	}
	return token, expiration, nil
}

func ParseToken(token string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
