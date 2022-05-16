package todo

import "todolist/entities"

type Todo interface {
	InsertTodo(newTodo entities.Todo) (entities.Todo, error)
	GetTodoID(ID uint, UserID uint) (entities.Todo, error)
	UpdateTodo(ID uint, UserID uint, update entities.Todo) (entities.Todo, error)
	DeleteTodo(ID uint, UserID uint) error
	Completed(id uint, UserID uint) (entities.Todo, error)
	Reopen(id uint, UserID uint) (entities.Todo, error)
	GetAllCompleteTodo(UserID uint) ([]entities.Todo, error)
	GetAllUnCompleteTodo(UserID uint) ([]entities.Todo, error)
}
