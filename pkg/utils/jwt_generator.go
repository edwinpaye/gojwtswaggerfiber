package utils

import (
	"os"
	"strconv"
	"time"

	// "github.com/edwinpaye/gots/pkg/common/config"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken() (string, error) {
	// c, _ := config.LoadConfig()
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	// secret := c.JWTKey

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	// minutesCount, _ := strconv.Atoi(c.JWTKExpire)

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
