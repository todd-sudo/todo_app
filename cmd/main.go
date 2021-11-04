package main

import (
	"log"

	todo "github.com/todd-sudo/todo_app"
	"github.com/todd-sudo/todo_app/pkg/handler"
	"github.com/todd-sudo/todo_app/pkg/repository"
	"github.com/todd-sudo/todo_app/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
