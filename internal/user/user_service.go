package user

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (u *UserService) Register(payload *user.RegisterRequest) (*user.RegisterResponse, error) {
	if err := payload.Validate(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Request body is not valid")
	}
	newUser := payload.ToEntity()

	registeredUser, err := u.userRepo.Register(newUser)
	if err != nil {
		return nil, err
	}

	response := &user.RegisterResponse{
		ID:        registeredUser.ID,
		Name:      registeredUser.Name,
		Email:     registeredUser.Email,
		AvatarURL: registeredUser.AvatarURL,
		CreatedAt: registeredUser.CreatedAt,
	}

	return response, nil
}
