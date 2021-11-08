package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/todd-sudo/todo_app"
)

type AuthPosgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPosgres {
	return &AuthPosgres{db: db}
}

func (r *AuthPosgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		usersTable,
	)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
