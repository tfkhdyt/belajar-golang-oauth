package post

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/post"
)

type PostHandler struct {
	postRepo post.PostRepository
}

func NewPostHandler(postRepo post.PostRepository) *PostHandler {
	return &PostHandler{postRepo}
}

func (p *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := p.postRepo.GetAllPosts()
	if err != nil {
		return err
	}

	return c.JSON(posts)
}
