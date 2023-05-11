package main

import (
	"log"

	"github.com/bohoslavskyi/go-todo-app"
	"github.com/bohoslavskyi/go-todo-app/pkg/handler"
	"github.com/bohoslavskyi/go-todo-app/pkg/repository"
	"github.com/bohoslavskyi/go-todo-app/pkg/service"
)

func main() {
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
