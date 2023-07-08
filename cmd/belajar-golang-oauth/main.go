package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/auth"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/index"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/post"
)

var ctx = context.Background()

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	postRepo := post.NewPostRepositoryDummy()

	postService := post.NewPostService(postRepo)
	authService := auth.NewAuthService(&ctx, &http.Client{})

	indexHandler := index.NewIndexHandler()
	postHandler := post.NewPostHandler(postService)
	authHandler := auth.NewAuthHandler(&ctx, authService)

	app.Get("/", indexHandler.Index)

	app.Get("/auth/login/github", authHandler.GetGitHubLoginURL)
	app.Get("/auth/callback/github", authHandler.HandleGitHubCallback)

	app.Get("/posts", auth.JwtMiddleware, postHandler.GetAllPosts)

	log.Fatal(app.Listen(":8080"))
}
