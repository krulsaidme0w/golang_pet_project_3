package application

import (
	"krulsaidme0w/library/internal/domain/entity"
)

type UserService interface {
	Save(userRequest *entity.UserRequest) error
	Get(userID string) (*entity.User, error)
	Update(user entity.User, updatedUser entity.UserRequest) error
	Delete(userID string) error
}

type AuthService interface {
}
