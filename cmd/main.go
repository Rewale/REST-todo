package main

import (
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
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

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Post:     "5432",
		Username: "wale1",
		Password: "djeb126",
		DBName:   "todo-go",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("error while connection to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
