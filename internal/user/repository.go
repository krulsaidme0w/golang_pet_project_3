package user

import "krulsaidme0w/library/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	Get(userID string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(userID string) error
}
