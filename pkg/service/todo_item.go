package service

import (
	todo "go-todo"
	"go-todo/pkg/repository"
)

type TodoItemService struct {
	repoList repository.TodoList
	repoItem repository.TodoItem
}

func (t *TodoItemService) DeleteItem(userId int, listId int, itemId int) error {
	_, err := t.repoList.GetListById(userId, listId)
	if err != nil {
		return err
	}

	err = t.repoItem.DeleteTodo(itemId, listId)
	if err != nil {
		return err
	}

	return nil

}
func (t *TodoItemService) UpdateItem(userId int, listId int, itemId int, input todo.UpdateTodoInput) error {
	_, err := t.repoList.GetListById(userId, listId)
	if err != nil {
		return err
	}

	err = input.Validate()
	if err != nil {
		return err
	}

	err = t.repoItem.UpdateTodo(itemId, listId, input)
	if err != nil {
		return err
	}

	return nil
}
func (t *TodoItemService) GetItemById(userId int, listId int, itemId int) (*todo.TodoItem, error) {
	_, err := t.repoList.GetListById(userId, listId)
	if err != nil {
		return nil, err
	}

	todoItem, err := t.repoItem.GetById(listId, itemId)

	return &todoItem, err
}

// TODO: decorator

func (t *TodoItemService) GetAllItems(userId int, listId int) ([]todo.TodoItem, error) {
	_, err := t.repoList.GetListById(userId, listId)
	if err != nil {
		return nil, err
	}
	todos, err := t.repoItem.GetAllTodo(listId)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoItemService) CreateItem(userId int, listId int, input todo.CreateTodoInput) (int, error) {
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
