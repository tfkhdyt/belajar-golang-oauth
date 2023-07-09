package post

type PostRepository interface {
	FindAllPosts() ([]Post, error)
}
