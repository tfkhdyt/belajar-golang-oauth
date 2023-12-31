package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	jwtConfig "github.com/tfkhdyt/belajar-golang-oauth/internal/config/jwt"
)

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(jwtConfig.JwtSecretKey)},
	TokenLookup: "cookie:token,header:Authorization",
	AuthScheme:  "Bearer",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	},
})
