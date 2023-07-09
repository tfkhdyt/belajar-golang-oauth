package user

import (
	"time"

	"github.com/tfkhdyt/belajar-golang-oauth/pkg/validator"
)

type RegisterRequest struct {
	validator.Validator
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	AvatarURL string `json:"avatar_url" validate:"url"`
	ID        uint   `json:"id" validate:"required"`
}

func (r *RegisterRequest) ToEntity() *User {
	var user User

	user.ID = r.ID
	user.Name = r.Name
	user.Email = r.Email
	user.AvatarURL = r.AvatarURL

	return &user
}

type RegisterResponse struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
	ID        uint      `json:"id"`
}
