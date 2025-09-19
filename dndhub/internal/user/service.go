package user

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repository
	jwtSecret []byte
}

func NewService(repo *Repository) *Service {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "changeme" // dev only, set properly in prod
	}
	return &Service{
		repo:      repo,
		jwtSecret: []byte(secret),
	}
}

func (s *Service) Register(ctx context.Context, username, password string) (User, error) {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// store user
	return s.repo.CreateUser(ctx, username, string(hash))
}

func (s *Service) Login(ctx context.Context, username, password string) (string, error) {
	u, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// issue JWT
	claims := jwt.MapClaims{
		"sub": u.ID.String(),
		"usr": u.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (User, error) {
	return s.repo.GetUserByUsername(ctx, id)
}
