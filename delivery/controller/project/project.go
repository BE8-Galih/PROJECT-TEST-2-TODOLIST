package project

import (
	"net/http"
	"strconv"
	middlewares "todolist/delivery/middleware"
	"todolist/delivery/view"
	projectV "todolist/delivery/view/project"
	"todolist/entities"
	"todolist/repository/project"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProjectController struct {
	Repo  project.Project
	Valid *validator.Validate
}

func NewControlProject(repo project.Project, valid *validator.Validate) *ProjectController {
	return &ProjectController{
		Repo:  repo,
		Valid: valid,
	}
}

func (p *ProjectController) InsertProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var InsertData projectV.InsertProject

		if err := c.Bind(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := p.Valid.Struct(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		newProject := entities.Project{Name: InsertData.Name, UserID: uint(UserID)}
		res, err := p.Repo.InsertProject(newProject)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := projectV.RespondProject{Name: res.Name, UserID: res.UserID, ProjectID: res.ID}
		log.Info(err)
		return c.JSON(http.StatusCreated, projectV.SuccessInsert(response))
	}
}

func (p *ProjectController) GetAllProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := p.Repo.GetAllProject(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		var GetAll []projectV.RespondProject
		for _, v := range result {
			Response := projectV.RespondProject{Name: v.Name, UserID: v.UserID, ProjectID: v.ID}
			GetAll = append(GetAll, Response)
		}
		return c.JSON(http.StatusOK, projectV.StatusGetAllProject(GetAll))
	}
}

func (p *ProjectController) GetProjectbyID() echo.HandlerFunc {
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

		res, err := p.Repo.GetProjectID(uint(convID), uint(UserID))

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := projectV.RespondProject{Name: res.Name, UserID: res.UserID, ProjectID: res.ID}

		return c.JSON(http.StatusOK, projectV.StatusGetIdOk(response))
	}

}

func (p *ProjectController) UpdateProjectID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update projectV.InsertProject

		if err := c.Bind(&update); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := p.Valid.Struct(&update); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateProject := entities.Project{Name: update.Name}

		res, err := p.Repo.UpdateProject(uint(id), uint(UserID), UpdateProject)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := projectV.RespondProject{Name: res.Name, UserID: res.UserID, ProjectID: res.ID}

		return c.JSON(http.StatusOK, projectV.StatusUpdate(response))
	}

}

func (p *ProjectController) DeleteProjectID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		error := p.Repo.DeleteProject(uint(convID), uint(UserID))

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}

func (p *ProjectController) AddTodoToProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var AddTodo projectV.InsertProjectTodo

		if err := c.Bind(&AddTodo); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := p.Valid.Struct(&AddTodo); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		res, err := p.Repo.AddTodoToProject(AddTodo.ProjectID, AddTodo.TodoID, uint(UserID))

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		Response := projectV.RespondTodoProject{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}
		log.Info(err)
		return c.JSON(http.StatusCreated, projectV.SuccessInsert(Response))
	}
}

func (p *ProjectController) GetAllProjectUnCompleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := p.Repo.GetAllProjectUnCompleteTodo(uint(id), uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		var GetAll []projectV.RespondTodoProject
		for _, v := range result {
			Response := projectV.RespondTodoProject{Name: v.Name, UserID: v.UserID, Completed: v.Completed, ProjectID: v.ProjectID, TodoID: v.ID}
			GetAll = append(GetAll, Response)
		}
		return c.JSON(http.StatusOK, projectV.StatusGetAllUnComplete(GetAll))
	}
}

func (p *ProjectController) GetAllProjectCompleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, err := p.Repo.GetAllProjectCompleteTodo(uint(id), uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		var GetAll []projectV.RespondTodoProject
		for _, v := range result {
			Response := projectV.RespondTodoProject{Name: v.Name, UserID: v.UserID, Completed: v.Completed, ProjectID: v.ProjectID, TodoID: v.ID}
			GetAll = append(GetAll, Response)
		}
		return c.JSON(http.StatusOK, projectV.StatusGetAllComplete(GetAll))
	}
}
func (p *ProjectController) TodoProjectCompleted() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		todo := c.Param("todo_id")
		ProjectID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		TodoID, err := strconv.Atoi(todo)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())

		}
		UserID := middlewares.ExtractTokenUserId(c)

		res, error := p.Repo.TodoProjectCompleted(uint(ProjectID), uint(TodoID), uint(UserID))

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := projectV.RespondTodoProject{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID, TodoID: res.ID}

		return c.JSON(http.StatusOK, projectV.StatusCompleted(response))
	}
}

func (p *ProjectController) TodoProjectReopen() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		todo := c.Param("todo_id")
		ProjectID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		TodoID, err := strconv.Atoi(todo)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		res, errorReopen := p.Repo.TodoProjectReopen(uint(ProjectID), uint(TodoID), uint(UserID))

		if errorReopen != nil {
			log.Warn(errorReopen)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := projectV.RespondTodoProject{Name: res.Name, UserID: res.UserID, Completed: res.Completed, ProjectID: res.ProjectID}

		return c.JSON(http.StatusOK, projectV.StatusReopen(response))
	}
}

func (p *ProjectController) MoveToHome() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		todo := c.Param("todo_id")
		ProjectID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		TodoID, err := strconv.Atoi(todo)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errorMove := p.Repo.MoveToHome(uint(ProjectID), uint(TodoID), uint(UserID))
		if errorMove != nil {
			log.Warn(errorMove)
		}
		return c.JSON(http.StatusOK, projectV.StatusMove())
	}
}
