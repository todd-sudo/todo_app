package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/todd-sudo/todo_app"
	"github.com/todd-sudo/todo_app/pkg/handler"
	"github.com/todd-sudo/todo_app/pkg/repository"
	"github.com/todd-sudo/todo_app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initilizing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
