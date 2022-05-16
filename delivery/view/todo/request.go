package todo

type InsertTodo struct {
	Name string `json:"name" validate:"required"`
}
