package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name      string `json:"name"`
	Completed string `json:"completed" gorm:"default:no"`
	UserID    uint
	ProjectID uint `json:"project_id" gorm:"default:0"`
}
