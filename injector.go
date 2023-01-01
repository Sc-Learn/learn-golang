//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"github.com/google/wire"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"sclearn/learn-golang/app"	
	"sclearn/learn-golang/middleware"
	"sclearn/learn-golang/repository"	
	"sclearn/learn-golang/service"	
	"sclearn/learn-golang/controller"	
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		userSet, 
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}