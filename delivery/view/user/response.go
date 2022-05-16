package user

import (
	"net/http"
	"time"
)

type LoginResponse struct {
	Token string
}

type RespondUser struct {
	UserID    uint      `json:"userId"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
	Gender    string    `json:"jenisKelamin"`
}

func SuccessInsert(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Register",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request Access",
		"status":  false,
		"data":    nil,
	}
}

func LoginOK(data LoginResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Login",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}
