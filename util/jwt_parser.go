package util

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires  int64
	UserId   int64
	Name     string
	Username string
	Role     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		user_id := int64(claims["user_id"].(float64))
		role := int64(claims["role"].(float64))

		return &TokenMetadata{
			Expires: expires,
			UserId:  user_id,
			Role:    role,
		}, nil
	}

	return nil, err
}

func ExtractTokenMetadataRefresh(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := VerifyTokenRefresh(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user_id := int64(claims["user_id"].(float64))

		return &TokenMetadata{
			UserId: user_id,
		}, nil
	}

	return nil, err
}

func ExtractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func ExtractTokenRefresh(c *fiber.Ctx) string {
	payload := c.Get("refreshToken")

	return payload
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, JwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyTokenRefresh(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractTokenRefresh(c)

	token, err := jwt.Parse(tokenString, JwtKeyRefreshFunc)
	if err != nil {
		return nil, errors.New("Verifikasi Token Failed")
	}

	return token, nil
}

func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func JwtKeyRefreshFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
