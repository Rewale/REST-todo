package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	todo "go-todo"
	"strings"
)

type TodoPostgresItem struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoPostgresItem {
	return &TodoPostgresItem{db: db}
}

func (r *TodoPostgresItem) CreateTodo(listId int, item todo.CreateTodoInput) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var todoId int
	createTodoItem := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) returning id",
		todoItemsTable)
	row := tx.QueryRow(createTodoItem, item.Title, item.Description)
	if err := row.Scan(&todoId); err != nil {
		logrus.Error(err.Error())
		err = tx.Rollback()
		return 0, err
	}

	createTodoListItem := "INSERT INTO lists_items  (item_id, list_id) values($1, $2)"

	_, err = tx.Exec(createTodoListItem, todoId, listId)
	if err != nil {
		logrus.Error(err.Error())
		err = tx.Rollback()
		return 0, err
	}

	return todoId, tx.Commit()

}
func (r *TodoPostgresItem) GetAllTodo(listId int) ([]todo.TodoItem, error) {
	var todoItems []todo.TodoItem
	query := "SELECT ti.* FROM todo_items ti INNER JOIN lists_items li on ti.id = li.item_id WHERE li.list_id = $1"
	err := r.db.Select(&todoItems, query, listId)

	return todoItems, err
}
func (r *TodoPostgresItem) GetById(listId, todoId int) (todo.TodoItem, error) {
	var todoItem todo.TodoItem
	query := "SELECT ti.* FROM todo_items ti INNER JOIN lists_items li on ti.id = li.item_id " +
		"WHERE li.list_id = $1 and ti.id=$2"
	err := r.db.Get(&todoItem, query, listId, todoId)

	if err != nil {
		err = errors.New(fmt.Sprintf("no such item with id %d", todoId))
	}

	return todoItem, err

}
func (r *TodoPostgresItem) DeleteTodo(todoId int) error {
	query := "DELETE FROM todo_items WHERE id=$1"
	result, err := r.db.Exec(query, todoId)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("nothing to delete")
	}
	return nil
}
func (r *TodoPostgresItem) UpdateTodo(todoId int, input todo.UpdateTodoInput) error {
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

	query := "UPDATE todo_items as tl SET " + setQuery + fmt.Sprintf(" Where id=$%d", argId)

	args = append(args, todoId)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("nothing to update")
	}
	return nil
}
