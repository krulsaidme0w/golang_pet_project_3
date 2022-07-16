package userservice

import (
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
)

type Repository interface {
	Save(user *userservice.User) error
	Get(id string) (*userservice.User, error)
	Update(user *userservice.User) error
	Delete(id string) error
}
