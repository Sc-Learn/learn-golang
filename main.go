package main

import (
	"net/http"
	"sclearn/learn-golang/helper"
	"sclearn/learn-golang/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(handler *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:2003",
		Handler: handler,
	}
}

func main()  {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}