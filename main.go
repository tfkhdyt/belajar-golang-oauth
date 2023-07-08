package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubUser struct {
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	ID        uint   `json:"id"`
}

var (
	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback/github",
		ClientID:     os.Getenv("OAUTH_GITHUB_ID"),
		ClientSecret: os.Getenv("OAUTH_GITHUB_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
	randomState  = os.Getenv("OAUTH_RANDOM_STATE")
	ctx          = context.Background()
	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
)

func main() {
	app := fiber.New()

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(jwtSecretKey)},
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`<body>
      <body>
        <a href='/auth/login/github'>
          <button>
            Login with GitHub
          </button>
        </a>
      </body>
    </body>`)
	})

	app.Get("/restricted", jwtMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world!",
		})
	})

	app.Get("/auth/login/github", func(c *fiber.Ctx) error {
		url := githubOauthConfig.AuthCodeURL(randomState)
		return c.Redirect(url, fiber.StatusTemporaryRedirect)
	})

	app.Get("/auth/callback/github", func(c *fiber.Ctx) error {
		state := c.FormValue("state")
		code := c.FormValue("code")

		if state != randomState {
			fmt.Println("state is not valid")
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		token, err := githubOauthConfig.Exchange(ctx, code)
		if err != nil {
			fmt.Printf("could not get token: %v\n", err.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		client := &http.Client{}
		req, errNewRequest := http.NewRequest("GET", "https://api.github.com/user", nil)
		if errNewRequest != nil {
			fmt.Printf("could not create new request: %v\n", errNewRequest.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)

		response, errGet := client.Do(req)
		if errGet != nil {
			fmt.Printf("could not get user data: %v\n", errGet.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		responseData, errRead := io.ReadAll(response.Body)
		if errRead != nil {
			fmt.Printf("could not get parse response data: %v\n", errRead.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		var userInfo GitHubUser
		if err := json.Unmarshal(responseData, &userInfo); err != nil {
			fmt.Printf("could not unmarshal response: %v\n", err.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		jwtToken, errJwt := CreateNewJWT(&userInfo)
		if errJwt != nil {
			fmt.Printf("could not create new jwt: %v\n", errJwt.Error())
			return c.Redirect("/", fiber.StatusTemporaryRedirect)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    *jwtToken,
			Path:     "/",
			Expires:  time.Now().Add(1 * time.Hour),
			HTTPOnly: true,
		})

		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	})

	log.Fatal(app.Listen(":8080"))
}

func CreateNewJWT(claims *GitHubUser) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       "belajar-golang-oauth",
		"sub":       "user",
		"userId":    claims.ID,
		"avatarUrl": claims.AvatarURL,
		"name":      claims.Name,
		"email":     claims.Email,
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
	})
	signedString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}
