package todo

import (
	"net/http"
	"strconv"
	middlewares "todolist/delivery/middleware"
	"todolist/delivery/view"
	todoV "todolist/delivery/view/todo"
	"todolist/entities"
	"todolist/repository/todo"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TodoController struct {
	Repo  todo.Todo
	Valid *validator.Validate
}

func NewControlTodo(repo todo.Todo, valid *validator.Validate) *TodoController {
	return &TodoController{
		Repo:  repo,
		Valid: valid,
	}
}

func (t *TodoController) InsertTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var InsertData todoV.InsertTodo

		if err := c.Bind(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := t.Valid.Struct(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		newTodo := entities.Todo{Name: InsertData.Name, UserID: uint(UserID)}
		res, err := t.Repo.InsertTodo(newTodo)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := todoV.RespondTodo{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}
		log.Info(err)
		return c.JSON(http.StatusCreated, todoV.SuccessInsert(response))
	}
}

func (t *TodoController) GetAllUnCompleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.Repo.GetAllUnCompleteTodo(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		var GetAll []todoV.RespondTodo
		for _, v := range result {
			Response := todoV.RespondTodo{Name: v.Name, UserID: v.UserID, Completed: v.Completed, ProjectID: v.ProjectID, TodoID: v.ID}
			GetAll = append(GetAll, Response)
		}
		return c.JSON(http.StatusOK, todoV.StatusGetAllUnComplete(GetAll))
	}
}

func (t *TodoController) GetAllCompleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.Repo.GetAllCompleteTodo(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		var GetAll []todoV.RespondTodo
		for _, v := range result {
			Response := todoV.RespondTodo{Name: v.Name, UserID: v.UserID, Completed: v.Completed, ProjectID: v.ProjectID, TodoID: v.ID}
			GetAll = append(GetAll, Response)
		}
		return c.JSON(http.StatusOK, todoV.StatusGetAllComplete(GetAll))
	}
}

func (t *TodoController) GetTodobyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		res, err := t.Repo.GetTodoID(uint(convID), uint(UserID))

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := todoV.RespondTodo{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}

		return c.JSON(http.StatusOK, todoV.StatusGetIdOk(response))
	}

}

func (t *TodoController) UpdateTodoID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update todoV.InsertTodo

		if err := c.Bind(&update); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := t.Valid.Struct(&update); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateTodo := entities.Todo{Name: update.Name}

		res, err := t.Repo.UpdateTodo(uint(id), uint(UserID), UpdateTodo)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := todoV.RespondTodo{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}

		return c.JSON(http.StatusOK, todoV.StatusUpdate(response))
	}

}

func (t *TodoController) DeleteTodoID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		error := t.Repo.DeleteTodo(uint(convID), uint(UserID))

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}

func (t *TodoController) Completed() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		res, error := t.Repo.Completed(uint(convID), uint(UserID))

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := todoV.RespondTodo{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}

		return c.JSON(http.StatusOK, todoV.StatusCompleted(response))
	}
}

func (t *TodoController) Reopen() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		res, error := t.Repo.Reopen(uint(convID), uint(UserID))

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := todoV.RespondTodo{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}

		return c.JSON(http.StatusOK, todoV.StatusReopen(response))
	}
}
