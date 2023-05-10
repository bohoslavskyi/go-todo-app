package main

import (
	"log"

	"github.com/bohoslavskyi/go-todo-app"
)

func main() {
	server := new(todo.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
