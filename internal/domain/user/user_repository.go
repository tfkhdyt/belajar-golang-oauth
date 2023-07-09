package user

type UserRepository interface {
	Register(user *User) (*User, error)
	GetUserByID(id uint) (*User, error)
}
