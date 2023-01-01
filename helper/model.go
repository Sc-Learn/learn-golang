package helper

import (
	"sclearn/learn-golang/model/api"
	"sclearn/learn-golang/model/domain"
)

func ToUserResponse(user domain.User) api.UserResponse {
	return api.UserResponse{
		Id: user.Id,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []api.UserResponse {
	var userResponse []api.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}

	return userResponse
}