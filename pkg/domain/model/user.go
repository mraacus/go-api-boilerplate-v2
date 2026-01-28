package model

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
	Role string `json:"role" validate:"required,oneof=admin user guest"`
}