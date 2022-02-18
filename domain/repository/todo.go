package repository

import "GoCleanArch/domain/model"

type TodoRepository interface {
	FindAllTodo() (todos []*model.Todo, err error)
	FindTodo(word string) (todos []*model.Todo, err error)
	CreateTodo(todo *model.Todo) (*model.Todo, error)
	UpdateTodo(todo *model.Todo) (*model.Todo, error)
}
