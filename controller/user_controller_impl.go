package controller

import (
	"net/http"
	"sclearn/learn-golang/helper"
	"sclearn/learn-golang/model/api"
	"sclearn/learn-golang/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userCreateRequest := api.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	apiResponse := api.ApiResponse{
		Code: 201,
		Status: "OK",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)  {
	userUpdateRequest := api.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId, err := strconv.Atoi(param.ByName("id"))
	helper.PanicIfError(err)
	userUpdateRequest.Id = userId

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	apiResponse := api.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userId, err := strconv.Atoi(param.ByName("id"))
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), userId)
	apiResponse := api.ApiResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userId, err := strconv.Atoi(param.ByName("id"))
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), userId)
	apiResponse := api.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)  {
	userResponses := controller.UserService.FindAll(request.Context())
	apiResponse := api.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: userResponses,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}