package post

type PostRepository interface {
	GetAllPosts() ([]Post, error)
}
