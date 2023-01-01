package api

type UserUpdateRequest struct {
	Id int `validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}