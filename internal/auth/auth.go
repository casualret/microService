package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(username string) (string, error) {

	claims := &jwt.RegisteredClaims{Issuer: username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func GenerateRandomKey() string {
//	key := make([]byte, 32) // generate a 256 bit key
//	_, err := rand.Read(key)
//	if err != nil {
//		panic("Failed to generate random key: " + err.Error())
//	}
//
//	return base64.StdEncoding.EncodeToString(key)
//}
