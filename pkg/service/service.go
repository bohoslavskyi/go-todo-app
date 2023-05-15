package service

import (
	"github.com/bohoslavskyi/go-todo-app"
	"github.com/bohoslavskyi/go-todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(use todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repositories),
	}
}
