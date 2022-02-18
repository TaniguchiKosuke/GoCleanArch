package infra

import (
	"GoCleanArch/domain/model"
)

type TodoRepository struct {
	SqlHandler
}

func NewTodoRepository(sqlHandler *SqlHandler) *TodoRepository {
	todoRepository := TodoRepository{*sqlHandler}
	return &todoRepository
}

func (tr *TodoRepository) FindAllTodo() ([]*model.Todo, error) {
	var todo []*model.Todo
	if err := tr.Conn.Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (tr *TodoRepository) FindTodo(text string) ([]*model.Todo, error) {
	var todo []*model.Todo
	if err := tr.Conn.Where("task LIKE ?", "%"+text+"%").Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (tr *TodoRepository) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	if err := tr.Conn.Create(todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (tr *TodoRepository) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	if err := tr.Conn.Where("id = ?", todo.ID).Find(todo).Error; err != nil {
		return todo, err
	}

	if err := tr.Conn.Updates(model.Todo{Task: todo.Task, LimitDate: todo.LimitDate, Status: todo.Status}).Error; err != nil {
		return todo, err
	}

	return todo, nil
}