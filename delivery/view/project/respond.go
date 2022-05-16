package project

import (
	"net/http"
)

type RespondProject struct {
	ProjectID uint   `json:"projectId"`
	Name      string `json:"name"`
	UserID    uint   `json:"userId"`
}

type RespondTodoProject struct {
	TodoID    uint   `json:"todoId"`
	Name      string `json:"name"`
	Completed string `json:"completed"`
	ProjectID uint   `json:"projectId"`
	UserID    uint   `json:"userId"`
}

func SuccessInsert(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Add Project",
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

func StatusGetAllProject(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Project",
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

func StatusCompleted(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Completed Project",
		"status":  true,
		"data":    data,
	}
}

func StatusReopen(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Reopen Project",
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

func StatusMove() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Move Project",
		"status":  true,
	}
}
