package app

import (
	"sclearn/learn-golang/controller"
	"sclearn/learn-golang/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:id", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:id", userController.Update)
	router.DELETE("/api/users/:id", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}