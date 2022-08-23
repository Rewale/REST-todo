package service

import (
	todo "go-todo"
	"go-todo/pkg/repository"
)

type TodoItemService struct {
	repoList repository.TodoList
	repoItem repository.TodoItem
}

func (t TodoItemService) CreateItem(userId int, listId int, input todo.CreateTodoInput) (int, error) {
	_, err := t.repoList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	createTodo, err := t.repoItem.CreateTodo(listId, input)
	if err != nil {
		return 0, err
	}

	return createTodo, nil
}

func NewTodoItemService(repoList repository.TodoList, repoItem repository.TodoItem) *TodoItemService {
	return &TodoItemService{repoList: repoList, repoItem: repoItem}
}
