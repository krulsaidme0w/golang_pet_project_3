package user

import "krulsaidme0w/library/internal/domain/repository"

type Service struct {
	repository repository.User
}

func NewUserService(repository repository.User) *Service {
	return &Service{
		repository: repository,
	}
}
