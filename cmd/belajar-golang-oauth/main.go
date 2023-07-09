package main

import (
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/auth"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/db"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/index"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/post"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/user"
)

var ctx = context.Background()

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

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
	userRepo := user.NewUserRepoPostgres(db.DB)

	postService := post.NewPostService(postRepo)
	authService := auth.NewAuthService(&ctx, userRepo)

	indexHandler := index.NewIndexHandler()
	postHandler := post.NewPostHandler(postService)
	authHandler := auth.NewAuthHandler(&ctx, authService)

	app.Get("/", indexHandler.Index)

	app.Get("/auth/login/github", authHandler.GetGitHubLoginURL)
	app.Get("/auth/callback/github", authHandler.HandleGitHubCallback)

	app.Get("/posts", auth.JwtMiddleware, postHandler.GetAllPosts)

	log.Fatal(app.Listen(":8080"))
}
