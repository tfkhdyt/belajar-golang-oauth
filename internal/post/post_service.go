package post

import "github.com/tfkhdyt/belajar-golang-oauth/internal/domain/post"

type PostService struct {
	postRepo post.PostRepository
}

func NewPostService(postRepo post.PostRepository) *PostService {
	return &PostService{postRepo}
}

func (p *PostService) FindAllPosts() ([]post.Post, error) {
	return p.postRepo.FindAllPosts()
}
