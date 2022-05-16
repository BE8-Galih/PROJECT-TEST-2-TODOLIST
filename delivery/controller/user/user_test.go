package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	middlewares "todolist/delivery/middleware"
	"todolist/entities"

	"github.com/labstack/echo/v4/middleware"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(3, "Yani", "y@gmail.com")
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"email":    "y",
			"password": "849",
			"phone":    "77979799",
			"status":   "starseller",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set Content to JSON

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := NewControlUser(&mockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, "Success Register", resp.Message)
		assert.True(t, resp.Status)
		assert.Equal(t, 201, resp.Code)
		assert.NotNil(t, resp.Data)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"password": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"phone": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 415, resp.Code)
	})
	t.Run("Error Insert DB", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"email":    "y@gmail.com",
			"password": "849",
			"phone":    "77979799",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 500, resp.Code)
	})
}

func TestGetUserbyID(t *testing.T) {
	t.Run("Success Get User by ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		GetUser := NewControlUser(&mockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.GetUserbyID())(context)
		type Response struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Status  bool        `json:"status"`
			Data    interface{} `json:"data"`
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data ID", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)

	})
	t.Run("Error Konversi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("c")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.GetUserbyID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Get DB", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(userController.GetUserbyID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(userController.GetUserbyID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})

}

func TestUpdateUserID(t *testing.T) {
	t.Run("Success Update Data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email": "y@gmail.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		UserCont := NewControlUser(&mockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(UserCont.UpdateUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		// assert.True(t, result.Status)

	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "Error Bind"

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(userController.UpdateUserID())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 415, resp.Code)
	})
	t.Run("Error Konversi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("c")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.UpdateUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.UpdateUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Not Found Access Token", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(userController.UpdateUserID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

func TestDeleteUserID(t *testing.T) {

	t.Run("Success Delete Address", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		GetUser := NewControlUser(&mockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.DeleteUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Konversi", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.DeleteUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.DeleteUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.NotEqual(t, "data tidak dapat didelete", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Not Found Access Token", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("6")
		GetUser := NewControlUser(&erorrMockUserRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetUser.DeleteUserID())(context)
		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

var token string

func TestLogin(t *testing.T) {
	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "y@gmail.com",
			"password": "yani",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := NewControlUser(&mockUserRepository{}, validator.New())
		controller.Login()(context)

		type ResponseStructure struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.True(t, response.Status)
		assert.NotNil(t, response.Data)
		data := response.Data.(map[string]interface{})
		token = data["Token"].(string)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"password": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		userController.Login()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 406, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"password": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := NewControlUser(&erorrMockUserRepository{}, validator.New())
		userController.Login()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 415, resp.Code)
	})
	t.Run("Error Login", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "y@gmail.com",
			"password": "yani",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")

		controller := NewControlUser(&erorrMockUserRepository{}, validator.New())
		controller.Login()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
}

// Dummy Data

type mockUserRepository struct{}

func (mur *mockUserRepository) InsertUser(newUser entities.User) (entities.User, error) {
	return entities.User{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com"}, nil
}

func (mur *mockUserRepository) GetAllUser() ([]entities.User, error) {
	return []entities.User{{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com"}}, nil
}

func (mur *mockUserRepository) GetUserID(ID int) (entities.User, error) {
	return entities.User{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com"}, nil
}

func (mur *mockUserRepository) UpdateUser(ID int, update entities.User) (entities.User, error) {
	return entities.User{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com"}, nil
}

func (mur *mockUserRepository) DeleteUser(ID int) error {
	return nil
}

func (mur *mockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com"}, nil
}

type erorrMockUserRepository struct{}

func (emur *erorrMockUserRepository) InsertUser(newPegawai entities.User) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa insert data")
}
func (emur *erorrMockUserRepository) GetAllUser() ([]entities.User, error) {
	return nil, errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) DeleteUser(ID int) error {
	return errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) GetUserID(ID int) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}
func (emur *erorrMockUserRepository) UpdateUser(ID int, update entities.User) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}
