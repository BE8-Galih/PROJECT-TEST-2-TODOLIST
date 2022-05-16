package main

import (
	projectC "todolist/delivery/controller/project"
	todoC "todolist/delivery/controller/todo"
	userC "todolist/delivery/controller/user"
	"todolist/delivery/routes"
	"todolist/repository/project"
	"todolist/repository/todo"
	"todolist/repository/user"
	"todolist/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {

	DB := utils.InitDB()
	RepoUser := user.NewRepoUser(DB)
	ControlUser := userC.NewControlUser(RepoUser, validator.New())

	RepoTodo := todo.NewRepoTodo(DB)
	ControlTodo := todoC.NewControlTodo(RepoTodo, validator.New())

	RepoProject := project.NewRepoProject(DB)
	ControlProject := projectC.NewControlProject(RepoProject, validator.New())
	e := echo.New()

	routes.Path(e, ControlUser, ControlTodo, ControlProject)

	e.Logger.Fatal(e.Start(":8000"))
}
