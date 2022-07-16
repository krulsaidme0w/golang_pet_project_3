package repository

import (
	"database/sql"

	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Save(user *userservice.User) error {
	query := `
		INSERT INTO library_user(username, email, password) 
		VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password)

	return err
}

func (r *UserRepository) Get(id string) (*userservice.User, error) {
	query := `
		SELECT id, username, email, password
		FROM library_user
		WHERE id = $1`

	row := r.DB.QueryRow(query, id)

	user := &userservice.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *userservice.User) error {
	query := `
		UPDATE library_user
		SET username = $1, email = $2, password = $3,
		WHERE id = $4`

	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.ID)

	return err
}

func (r *UserRepository) Delete(id string) error {
	query := `
		DELETE FROM library_user
		WHERE id = $1`

	_, err := r.DB.Exec(query, id)

	return err
}
