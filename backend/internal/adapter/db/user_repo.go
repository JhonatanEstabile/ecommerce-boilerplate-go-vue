package db

import (
	"database/sql"
	"ecommerce/internal/core/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindByEmail(email string) (*user.User, error) {
	var user user.User

	err := r.db.QueryRow(`
		SELECT id, name, email, password FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Create(user user.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`, user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
