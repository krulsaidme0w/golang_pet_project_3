package psql

import (
	"database/sql"

	"krulsaidme0w/library/internal/domain/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Save(user *entity.User) error {
	return nil
}

func (r *UserRepository) Get(userID string) (*entity.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	return nil
}

func (r *UserRepository) Delete(userID string) error {
	return nil
}
