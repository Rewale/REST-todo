package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "go-todo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string, passwordHash string) (*todo.User, error) {
	user := todo.User{
		Username: username,
		Password: passwordHash,
	}
	query := fmt.Sprintf("SELECT name, id  FROM %s WHERE password_hash=$1 and username=$2", userTable)
	err := r.db.Get(&user, query, passwordHash, username)
	if err != nil {
		err = errors.New("No such user " + username)

	}
	return &user, err
}
