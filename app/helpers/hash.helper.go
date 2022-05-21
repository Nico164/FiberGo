package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "rahasia"

func GeneratePassWord(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 14)
}

func ComPassWord(hash, pw []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pw)
}

func NewClaim(id string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

func SaveCookie(c *fiber.Ctx, key, value string) fiber.Cookie {
	return fiber.Cookie{
		Name:     key,
		Value:    value,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
}
