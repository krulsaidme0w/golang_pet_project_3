package userservice

import (
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
)

type UseCase interface {
	Save(userRequest *userservice.UserRequest) error
	Get(id string) (*userservice.User, error)
	Update(user userservice.User, updatedUser userservice.UserRequest) error
	Delete(id string) error
}
