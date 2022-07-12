package application

import "krulsaidme0w/library/internal/domain/repository"

type User struct {
	repository repository.User
}

func NewUserApplication(repository repository.User) *User {
	return &User{
		repository: repository,
	}
}
