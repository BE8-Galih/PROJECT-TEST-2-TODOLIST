package todo

import (
	"net/http"
)

type RespondTodo struct {
	TodoID    uint   `json:"todoId"`
	Name      string `json:"name"`
	Completed string `json:"completed"`
	ProjectID uint   `json:"projectId"`
	UserID    uint   `json:"userId"`
}

func SuccessInsert(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Add Todo",
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

func StatusCompleted(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Completed Todo",
		"status":  true,
		"data":    data,
	}
}

func StatusReopen(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Reopen Todo",
		"status":  true,
		"data":    data,
	}
}

func StatusGetAllComplete(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Todo Complete",
		"status":  true,
		"data":    data,
	}
}

func StatusGetAllUnComplete(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Todo Uncomplete",
		"status":  true,
		"data":    data,
	}
}
