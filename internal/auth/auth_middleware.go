package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	jwtConfig "github.com/tfkhdyt/belajar-golang-oauth/internal/config/jwt"
)

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(jwtConfig.JwtSecretKey)},
	TokenLookup: "cookie:token",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	},
})
