package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) FindMyUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(float64)

	myUserInfo, errUserInfo := u.userService.FindMyUserInfo(uint(userId))
	if errUserInfo != nil {
		return errUserInfo
	}

	return c.JSON(fiber.Map{
		"data": myUserInfo,
	})
}
