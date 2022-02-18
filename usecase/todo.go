package usecase

import (
	"GoCleanArch/domain/model"
	"GoCleanArch/domain/repository"
)

type TodoUsecase interface {
    SearchTodo(string) (todo []*model.Todo, err error)
    ViewTodo() (todo []*model.Todo, err error)
    AddTodo(todo *model.Todo) (err error)
    EditTodo(todo *model.Todo) (err error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) *todoUsecase {
	todoUsecase := todoUsecase{todoRepository: todoRepository}
	return &todoUsecase
}

func (usecase *todoUsecase) ViewTodo() ([]*model.Todo, error) {
	todo, err := usecase.todoRepository.FindAllTodo()
	return todo, err
}

func (usecase *todoUsecase) SearchTodo(text string) ([]*model.Todo, error) {
	todo, err := usecase.todoRepository.FindTodo(text)
	return todo, err
}

func (usecase *todoUsecase) AddTodo(todo *model.Todo) error {
	_, err := usecase.todoRepository.CreateTodo(todo)
	return err
}

func (usecase *todoUsecase) EditTodo(todo *model.Todo) error {
	_, err := usecase.todoRepository.UpdateTodo(todo)
	return err
}