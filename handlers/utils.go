package handlers

import (
	"os"
	"time"

	"github.com/daniilmikhaylov2005/mangagram/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(hash), err
}

func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func CreateToken(id int) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}
	claims := &models.UserClaims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringKey := os.Getenv("SECRET_KEY")
	byteKey := []byte(stringKey)

	stringToken, err := token.SignedString(byteKey)
	return stringToken, err
}
