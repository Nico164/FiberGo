package helpers

import (
	"time"

	"github.com/Nico164/FiberGo/app/models"
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

func ReadJWT(cookie string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
}

func ExtractToken(token *jwt.Token) (*jwt.StandardClaims, error) {
	return token.Claims.(*jwt.StandardClaims), nil
}

func ExtractUser(c *fiber.Ctx) models.User {
	return c.Locals("user").(models.User)
}
