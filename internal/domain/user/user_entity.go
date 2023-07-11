package user

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	jwtConfig "github.com/tfkhdyt/belajar-golang-oauth/internal/config/jwt"
)

var httpClient = &http.Client{}

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	AvatarURL string `json:"avatar_url"`
}

func (u *User) CreateNewJWT() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "belajar-golang-oauth",
		"userId": u.ID,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	})
	if token == nil {
		return nil, fiber.NewError(
			fiber.StatusUnauthorized,
			"Failed to create new token",
		)
	}

	signedString, err := token.SignedString([]byte(jwtConfig.JwtSecretKey))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return &signedString, nil
}

func (u *User) GetGitHubUserInfo(token *oauth2.Token) error {
	req, errNewRequest := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if errNewRequest != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Failed to setup new request",
		)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	response, errGet := httpClient.Do(req)
	if errGet != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Failed to get github user info",
		)
	}

	responseData, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Failed to read response body",
		)
	}

	var user User
	if err := json.Unmarshal(responseData, &user); err != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Failed to unmarshal response",
		)
	}

	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.AvatarURL = user.AvatarURL

	return nil
}
