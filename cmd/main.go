package main

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	todo "go-todo"
	"go-todo/pkg/handler"
	"go-todo/pkg/repository"
	"go-todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while reading config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
