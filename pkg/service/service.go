package service

import "github.com/bohoslavskyi/go-todo-app/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	repositories *repository.Repository
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		repositories: repositories,
	}
}
