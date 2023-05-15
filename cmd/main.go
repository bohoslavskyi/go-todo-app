package main

import (
	"log"

	"github.com/bohoslavskyi/go-todo-app"
	"github.com/bohoslavskyi/go-todo-app/pkg/handler"
	"github.com/bohoslavskyi/go-todo-app/pkg/repository"
	"github.com/bohoslavskyi/go-todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:         viper.GetString("db.hostname"),
		Port:         viper.GetInt("db.port"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		DatabaseName: viper.GetString("db.name"),
		SSLMode:      viper.GetString("db.sslMode"),
	})
	if err != nil {
		log.Fatal(err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
