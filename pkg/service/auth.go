package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/bohoslavskyi/go-todo-app"
	"github.com/bohoslavskyi/go-todo-app/pkg/repository"
)

const salt = "ams9f8g5haebhpgnp8w4gh9q3gb9n"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
