package view

import "net/http"

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Cannot Access Database",
		"status":  false,
	}
}

func DataEmpty() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Is Empty",
		"status":  false,
	}
}

func NotFound() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": "Data Not Found",
		"status":  false,
	}
}

func ConvertID() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Cannot Convert ID",
		"status":  false,
	}
}

func BindData() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnsupportedMediaType,
		"message": "Unsupported Media Type",
		"status":  false,
	}
}

func Validate() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Data Type False or Input Data is Required",
		"status":  false,
	}
}

func StatusDelete() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Deleted",
		"status":  true,
	}
}
