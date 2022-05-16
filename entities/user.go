package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Todo     []Todo    `gorm:"foreignKey:UserID;references:id"`
	Projects []Project `gorm:"foreignKey:UserID;references:id"`
}
