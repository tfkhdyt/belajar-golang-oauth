package oauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	GithubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback/github",
		ClientID:     os.Getenv("OAUTH_GITHUB_ID"),
		ClientSecret: os.Getenv("OAUTH_GITHUB_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
	RandomState = os.Getenv("OAUTH_RANDOM_STATE")
)
