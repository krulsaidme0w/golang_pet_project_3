package models

import (
	"net/mail"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/security"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/errors"
)

const (
	maxEmailLength    = 64
	maxUsernameLength = 64
)

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u *UserRequest) ToUser() (*User, error) {
	if u.Username == "" {
		return nil, errors.ErrEmptyEmail
	}

	if u.Email == "" {
		return nil, errors.ErrEmptyEmail
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return nil, err
	}

	if len(u.Username) > maxUsernameLength {
		return nil, errors.ErrBadUsernameLength
	}

	if len(u.Email) > maxEmailLength {
		return nil, errors.ErrBadEmailLength
	}

	if u.Password != u.ConfirmPassword {
		return nil, errors.ErrDifferentPasswords
	}

	if err := security.ValidatePassword(u.Password); err != nil {
		return nil, security.ErrBadPassword
	}

	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: security.Hash(u.Password),
	}, nil
}
