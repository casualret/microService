package service

import (
	"fmt"
	"microService/internal/auth"
	"microService/internal/models"
	"microService/internal/storage"
)

type Authentication interface {
	SignUp(user models.CreateUserReq) error
	SignIn(user models.User) (string, error)
}

type AuthenticationService struct {
	storage *storage.Postgres
}

func NewAuthenticationService(storage *storage.Postgres) *AuthenticationService {
	return &AuthenticationService{storage: storage}
}

func (a *AuthenticationService) SignUp(user models.CreateUserReq) error {
	const op = "service.SignUp"

	var err error
	if user.Password, err = auth.HashPassword(user.Password); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = a.storage.CreateUser(user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *AuthenticationService) SignIn(user models.User) (string, error) {
	const op = "service.SignIn"

	token, err := a.storage.SignIn(user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}
