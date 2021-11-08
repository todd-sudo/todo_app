package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/todd-sudo/todo_app"
	"github.com/todd-sudo/todo_app/pkg/repository"
)

const salt = "fslkfw9yu928rijoIO"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// generatePasswordHash хэширует пароль пользователя
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
