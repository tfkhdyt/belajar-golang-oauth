package auth

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/config/oauth"
)

type AuthHandler struct {
	ctx         *context.Context
	authService *AuthService
}

func NewAuthHandler(ctx *context.Context, authService *AuthService) *AuthHandler {
	return &AuthHandler{ctx, authService}
}

func (a *AuthHandler) GetGitHubLoginURL(c *fiber.Ctx) error {
	url := oauth.GithubOauthConfig.AuthCodeURL(oauth.RandomState)
	if url == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to get github login url")
	}

	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (a *AuthHandler) HandleGitHubCallback(c *fiber.Ctx) error {
	code := c.FormValue("code")
	state := c.FormValue("state")

	jwtToken, err := a.authService.HandleGitHubCallback(code, state)
	if err != nil {
		return c.Redirect("/?error="+err.Error(), fiber.StatusTemporaryRedirect)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    *jwtToken,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
	})

	return c.Redirect("/", fiber.StatusTemporaryRedirect)
}
