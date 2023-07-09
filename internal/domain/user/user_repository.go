package user

type UserRepository interface {
	Register(user *User) (*User, error)
	FindUserByID(id uint) (*User, error)
}
