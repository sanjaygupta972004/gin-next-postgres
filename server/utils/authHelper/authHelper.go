package authHelper

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savvy-bit/gin-react-postgres/config"
	"github.com/savvy-bit/gin-react-postgres/models"
)

func SignAccessToken(user *models.User) (string, error) {

	authConfig := config.GetGlobalConfig().AuthToken

	claims := jwt.MapClaims{
		"userID":   user.UserID.String(),
		"email":    user.Email,
		"role":     user.Role,
		"username": user.Username,
		"iss":      "go-gin-postgres",
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	}

	if authConfig.AccessToken == "" {
		return "", errors.New("access token is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signAccessToken, err := token.SignedString([]byte(authConfig.AccessToken))
	if err != nil {
		return "", err
	}

	return signAccessToken, nil
}

func SignRefreshToken(user *models.User) (string, error) {

	authConfig := config.GetGlobalConfig().AuthToken

	claims := jwt.MapClaims{
		"userID":   user.UserID.String(),
		"email":    user.Email,
		"role":     user.Role,
		"username": user.Username,
		"iss":      "go-gin-postgres",
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	if authConfig.RefreshToken == "" {
		return "", errors.New("refresh token is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signRefreshToken, err := token.SignedString([]byte(authConfig.RefreshToken))
	if err != nil {
		return "", err
	}

	return signRefreshToken, nil

}

func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		otp[i] = digits[int(otp[i])%len(digits)]
	}
	return string(otp), nil
}

// will have to implement email verification and password reset logic here in future
