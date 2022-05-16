package user

import (
	"net/http"
	"strconv"
	middlewares "todolist/delivery/middleware"
	"todolist/delivery/view"
	userV "todolist/delivery/view/user"
	"todolist/entities"
	"todolist/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  user.User
	Valid *validator.Validate
}

func NewControlUser(repo user.User, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (u *UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var InsertData userV.InsertUserRequest

		if err := c.Bind(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := u.Valid.Struct(&InsertData); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		newUser := entities.User{Name: InsertData.Name, Email: InsertData.Email, Password: InsertData.Password, Phone: InsertData.Phone}
		res, err := u.Repo.InsertUser(newUser)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := userV.RespondUser{Name: res.Name, Email: res.Email, Phone: res.Phone, UserID: res.ID}
		log.Info(err)
		return c.JSON(http.StatusCreated, userV.SuccessInsert(response))
	}
}

func (u *UserController) GetUserbyID() echo.HandlerFunc {
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

		res, err := u.Repo.GetUserID(convID)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := userV.RespondUser{Name: res.Name, Email: res.Email, Phone: res.Phone, UserID: res.ID}

		return c.JSON(http.StatusOK, userV.StatusGetIdOk(response))
	}

}

func (u *UserController) UpdateUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update userV.UpdateUserRequest

		if err := c.Bind(&update); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		if UserID != float64(id) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		UpdateUser := entities.User{Email: update.Email, Name: update.Name, Password: update.Password, Phone: update.Phone}

		res, err := u.Repo.UpdateUser(id, UpdateUser)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}
		response := userV.RespondUser{Name: res.Name, Email: res.Email, Phone: res.Phone, UserID: res.ID}

		return c.JSON(http.StatusOK, userV.StatusUpdate(response))
	}

}

func (u *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := userV.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := u.Valid.Struct(&param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		res, err := u.Repo.Login(param.Email, param.Password)

		if err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		response := userV.LoginResponse{}

		if response.Token == "" {
			token, _ := middlewares.CreateToken(float64(res.ID), (res.Name), (res.Email))
			response.Token = token
		}

		return c.JSON(http.StatusOK, userV.LoginOK(response))
	}
}

func (u *UserController) DeleteUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		if UserID != float64(convID) {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		error := u.Repo.DeleteUser(convID)

		if error != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
