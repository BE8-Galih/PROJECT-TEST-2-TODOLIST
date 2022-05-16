package todo

import "github.com/labstack/echo/v4"

type ControlTodo interface {
	InsertTodo() echo.HandlerFunc
	GetAllUnCompleteTodo() echo.HandlerFunc
	GetAllCompleteTodo() echo.HandlerFunc
	GetTodobyID() echo.HandlerFunc
	UpdateTodoID() echo.HandlerFunc
	DeleteTodoID() echo.HandlerFunc
	Completed() echo.HandlerFunc
	Reopen() echo.HandlerFunc
}
