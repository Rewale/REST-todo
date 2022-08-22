package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	todo "go-todo"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) CreateList(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) returning id",
		todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		logrus.Error(err.Error())
		err = tx.Rollback()
		return 0, err
	}

	createUsersListQuery := "INSERT INTO users_lists (user_id, list_id) values ($1, $2)"

	_, err = tx.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		logrus.Error(err.Error())
		err = tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}
