package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	todo "go-todo"
	"strings"
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
func (r *TodoPostgres) GetAllLists(userId int) ([]todo.TodoList, error) {
	var todoLists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.*  FROM %s tl INNER JOIN %s ul"+
		" on tl.id = ul.list_id where ul.user_id=$1", todoListsTable, usersListsTable)
	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoPostgres) GetListById(userId int, listId int) (todo.TodoList, error) {
	var todoList todo.TodoList
	query := fmt.Sprintf("SELECT tl.*  FROM %s tl INNER JOIN %s ul"+
		" on tl.id = ul.list_id where ul.user_id=$1 and tl.id=$2", todoListsTable, usersListsTable)
	err := r.db.Get(&todoList, query, userId, listId)

	if err != nil {
		err = errors.New(fmt.Sprintf("no such list with id %d", listId))
	}

	return todoList, err
}
func (r *TodoPostgres) DeleteListById(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id and ul.user_id=$2 and ul.list_id=$1 ",
		todoListsTable, usersListsTable)
	result, err := r.db.Exec(query, listId, userId)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("nothing to delete")
	}

	return nil
}
func (r *TodoPostgres) UpdateListById(userId int, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := "UPDATE todo_lists as tl SET " + setQuery + " " +
		"FROM users_lists as ul " +
		fmt.Sprintf("WHERE ul.user_id=$%d and ul.list_id=tl.id and tl.id=$%d", argId, argId+1)

	args = append(args, userId, listId)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("nothing to update")
	}
	return nil
}
