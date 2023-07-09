package user

import (
	"gorm.io/gorm"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/user"
)

type userRepoPostgres struct {
	db *gorm.DB
}

func NewUserRepoPostgres(db *gorm.DB) user.UserRepository {
	return &userRepoPostgres{db}
}

func (u *userRepoPostgres) Register(user *user.User) (*user.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepoPostgres) GetUserByID(id uint) (*user.User, error) {
	var user user.User
	if err := u.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
