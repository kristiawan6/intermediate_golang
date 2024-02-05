package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(secretKey, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	
	claims["role"] = role

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RefreshToken(secretKey, email string, role string) (refreshToken string, err error) {
	// Refresh Token
	refreshTokenClaims := jwt.New(jwt.SigningMethodHS256).Claims.(jwt.MapClaims)
	refreshTokenClaims["email"] = email
	refreshTokenClaims["role"] = role
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
