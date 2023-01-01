package service

import (
	"context"
	"sclearn/learn-golang/model/api"
)

type UserService interface{
	Create(ctx context.Context, request api.UserCreateRequest) api.UserResponse
	Update(ctx context.Context, request api.UserUpdateRequest) api.UserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) api.UserResponse
	FindAll(ctx context.Context) []api.UserResponse
}