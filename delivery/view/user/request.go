package user

type InsertUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required" gorm:"unique"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
