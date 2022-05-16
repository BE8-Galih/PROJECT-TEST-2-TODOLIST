package project

type InsertProject struct {
	Name string `json:"name" validate:"required"`
}

type InsertProjectTodo struct {
	TodoID    uint `json:"TodoId" validate:"required"`
	ProjectID uint `json:"ProjectId" validate:"required"`
}
