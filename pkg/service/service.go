package service

import (
	todo "go-todo"
	"go-todo/pkg/repository"
)

//go:generate  mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	CheckToken(jwtToken string) (int, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(id int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId int, listId int) error
	UpdateListById(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userId, listId int, input todo.CreateTodoInput) (int, error)
	GetAllItems(userId int, listId int) ([]todo.TodoItem, error)
	GetItemById(userId int, listId int, itemId int) (*todo.TodoItem, error)
	DeleteItem(userId int, listId int, itemId int) error
	UpdateItem(userId int, listId int, itemId int, input todo.UpdateTodoInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoList, repos.TodoItem),
	}
}
