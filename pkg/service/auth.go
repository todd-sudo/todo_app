package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	todo "github.com/todd-sudo/todo_app"
	"github.com/todd-sudo/todo_app/pkg/repository"
)

const (
	salt       = "fslkfw9yu928rijoIO"
	signingKey = "OIH&kljo*^((OU)(&^johtg7" // ключ подписи
	tokenTTL   = 12 * time.Hour             // время жизни токена | 12 часов
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken генерирует токен для пользователя
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// получаем пользователя по username and password
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время жизни токена | 12 часов
			IssuedAt:  time.Now().Unix(),               // время, когда токен был сгенерирован
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

// generatePasswordHash хэширует пароль пользователя
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
