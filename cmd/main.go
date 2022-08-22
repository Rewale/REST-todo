package main

import (
	"github.com/labstack/gommon/log"
	todo "go-todo"
	"go-todo/pkg/handler"
	"go-todo/pkg/repository"
	"go-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
