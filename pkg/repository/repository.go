package repository

import (
	"github.com/jmoiron/sqlx"
	todo "go-todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (*todo.User, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId, listId int) error
	UpdateListById(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateTodo(listId int, item todo.CreateTodoInput) (int, error)
	GetAllTodo(listId int) ([]todo.TodoItem, error)
	GetById(listId, todoId int) (todo.TodoItem, error)
	DeleteTodo(todoId int) error
	UpdateTodo(todoId int, input todo.UpdateTodoInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
