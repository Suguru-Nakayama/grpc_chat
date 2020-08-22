package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	LIFE_TIME = 30 * time.Minute
)

func GenerateToken(userID uint) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat":     now.Unix(),
		"exp":     now.Add(LIFE_TIME).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
