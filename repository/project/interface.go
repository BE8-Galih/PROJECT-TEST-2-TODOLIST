package project

import "todolist/entities"

type Project interface {
	InsertProject(newProject entities.Project) (entities.Project, error)
	GetAllProject(UserID uint) ([]entities.Project, error)
	GetProjectID(id uint, UserID uint) (entities.Project, error)
	UpdateProject(id uint, UserID uint, update entities.Project) (entities.Project, error)
	DeleteProject(id uint, UserID uint) error
	AddTodoToProject(id uint, TodoID uint, UserID uint) (entities.Todo, error)
	TodoProjectCompleted(id uint, TodoID uint, UserID uint) (entities.Todo, error)
	TodoProjectReopen(id uint, TodoID, UserID uint) (entities.Todo, error)
	GetAllProjectCompleteTodo(id uint, UserID uint) ([]entities.Todo, error)
	GetAllProjectUnCompleteTodo(id uint, UserID uint) ([]entities.Todo, error)
	MoveToHome(id uint, TodoID uint, UserID uint) error
}
