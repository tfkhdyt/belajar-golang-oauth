package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	jwtConfig "github.com/tfkhdyt/belajar-golang-oauth/internal/config/jwt"
)

type User struct {
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	ID        uint   `json:"id"`
}

func (u *User) CreateNewJWT() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       "belajar-golang-oauth",
		"sub":       "user",
		"userId":    u.ID,
		"avatarUrl": u.AvatarURL,
		"name":      u.Name,
		"email":     u.Email,
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
	})
	if token == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to create new token")
	}

	signedString, err := token.SignedString([]byte(jwtConfig.JwtSecretKey))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return &signedString, nil
}
