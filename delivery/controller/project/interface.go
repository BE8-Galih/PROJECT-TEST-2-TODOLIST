package project

import "github.com/labstack/echo/v4"

type ControlProject interface {
	InsertProject() echo.HandlerFunc
	GetAllProject() echo.HandlerFunc
	GetProjectbyID() echo.HandlerFunc
	UpdateProjectID() echo.HandlerFunc
	DeleteProjectID() echo.HandlerFunc
	AddTodoToProject() echo.HandlerFunc
	GetAllProjectUnCompleteTodo() echo.HandlerFunc
	GetAllProjectCompleteTodo() echo.HandlerFunc
	TodoProjectCompleted() echo.HandlerFunc
	TodoProjectReopen() echo.HandlerFunc
	MoveToHome() echo.HandlerFunc
}
