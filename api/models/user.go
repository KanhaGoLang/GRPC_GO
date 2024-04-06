package models

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,validateRole"`
	IsActive bool   `json:"is_active"`
}

var ValidRoles = []string{
	"admin",
	"editor",
	"subscriber",
}
