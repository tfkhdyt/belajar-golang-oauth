package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/config/oauth"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/user"
)

type AuthService struct {
	ctx      *context.Context
	userRepo user.UserRepository
}

func NewAuthService(ctx *context.Context, userRepo user.UserRepository) *AuthService {
	return &AuthService{ctx, userRepo}
}

func (a *AuthService) HandleGitHubCallback(code string, state string) (*string, error) {
	if state != oauth.RandomState {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid state")
	}

	token, err := a.getGitHubToken(code)
	if err != nil {
		return nil, err
	}

	var user user.User
	if err := user.GetGitHubUserInfo(token); err != nil {
		return nil, err
	}

	if _, err := a.userRepo.GetUserByID(user.ID); err != nil {
		if _, errRegister := a.userRepo.Register(&user); errRegister != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, errRegister.Error())
		}
	}

	jwtToken, errJwt := user.CreateNewJWT()
	if errJwt != nil {
		return nil, errJwt
	}

	return jwtToken, nil
}

func (a *AuthService) getGitHubToken(code string) (*oauth2.Token, error) {
	token, err := oauth.GithubOauthConfig.Exchange(*a.ctx, code)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to get github token")
	}

	return token, nil
}
