package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/config/oauth"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/user"
)

type AuthService struct {
	ctx        *context.Context
	httpClient *http.Client
}

func NewAuthService(ctx *context.Context, httpClient *http.Client) *AuthService {
	return &AuthService{ctx, httpClient}
}

func (a *AuthService) HandleGitHubCallback(code string, state string) (*string, error) {
	if state != oauth.RandomState {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid state")
	}

	token, err := a.getGitHubToken(code)
	if err != nil {
		return nil, err
	}

	user, errGitHubInfo := a.getGitHubUserInfo(token)
	if errGitHubInfo != nil {
		return nil, errGitHubInfo
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

func (a *AuthService) getGitHubUserInfo(token *oauth2.Token) (*user.User, error) {
	req, errNewRequest := http.NewRequest("GET", "https://api.github.com/user", nil)
	if errNewRequest != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to setup new request")
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	response, errGet := a.httpClient.Do(req)
	if errGet != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to get github user info")
	}

	responseData, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to read response body")
	}

	var user user.User
	if err := json.Unmarshal(responseData, &user); err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Failed to unmarshal response")
	}

	return &user, nil
}
