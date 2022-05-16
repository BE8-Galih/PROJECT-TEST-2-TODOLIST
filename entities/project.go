package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint
}
