package util

import (
	"backend-c-payment-monitoring/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// generate access token.
func GenerateNewAccessToken(user model.User) (string, error) {
	// Set secret key
	secret := os.Getenv("JWT_SECRET")

	// Set expires minutes = 30 minutes
	minutesCount := 30

	// Create a new claims
	claims := jwt.MapClaims{}
	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["user_id"] = user.ID
	claims["role"] = user.RoleID

	// Create a new JWT token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

// generate refresh token.
func GenerateRefreshToken(user model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	hourCount := 160

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(hourCount)).Unix()
	claims["user_id"] = user.ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
