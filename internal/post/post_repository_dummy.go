package post

import "github.com/tfkhdyt/belajar-golang-oauth/internal/domain/post"

type postRepoDummy struct{}

func NewPostRepositoryDummy() post.PostRepository {
	return &postRepoDummy{}
}

func (p *postRepoDummy) GetAllPosts() ([]post.Post, error) {
	posts := []post.Post{
		{
			ID:    "post-1",
			Title: "Belajar OAuth dengan Golang",
			Body:  "Pada hari ini saya belajar mengimplementasikan OAuth menggunakan bahasa pemrograman Golang",
		},
	}

	return posts, nil
}
