package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
