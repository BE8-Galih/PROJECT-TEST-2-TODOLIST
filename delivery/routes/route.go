package routes

import (
	"todolist/delivery/controller/project"
	"todolist/delivery/controller/todo"
	"todolist/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, u user.ControllerUser, t todo.ControlTodo, p project.ControlProject) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Login
	e.POST("/login", u.Login())
	// ROUTES USER
	user := e.Group("/user")
	user.POST("", u.InsertUser())
	user.GET("/:id", u.GetUserbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.PUT("/:id", u.UpdateUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.DELETE("/:id", u.DeleteUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// ROUTES TODO
	todo := e.Group("/todo")
	todo.POST("", t.InsertTodo(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.GET("/uncompleted", t.GetAllUnCompleteTodo(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.GET("/completed", t.GetAllCompleteTodo(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.GET("/:id", t.GetTodobyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.PUT("/:id", t.UpdateTodoID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.DELETE("/:id", t.DeleteTodoID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.PUT("/:id/completed", t.Completed(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	todo.PUT("/:id/reopen", t.Reopen(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// ROUTES PROJECT
	project := e.Group("/project")
	project.POST("", p.InsertProject(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.GET("", p.GetAllProject(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.GET("/:id", p.GetProjectbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.PUT("/:id", p.UpdateProjectID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.DELETE("/:id", p.DeleteProjectID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.PUT("/todo", p.AddTodoToProject(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.GET("/:id/todo/uncompleted", p.GetAllProjectUnCompleteTodo(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.GET("/:id/todo/completed", p.GetAllProjectCompleteTodo(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.PUT("/:id/todo/:todo_id", p.TodoProjectCompleted(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.PUT("/:id/todo/:todo_id", p.TodoProjectReopen(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	project.PUT("/:id/todo/:todo_id", p.MoveToHome(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
}
