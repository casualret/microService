package storage

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"microService/internal/auth"
	"microService/internal/models"
)

func (p *Postgres) CreateUser(user models.User) error {
	const op = "postgres.CreateUser"

	//fmt.Println(user.Password)

	query := `INSERT INTO users (email, hash_password, role) VALUES ($1, $2, $3)`
	_, err := p.database.Query(query, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("%s:%d", op, err)
	}

	return nil
}

func (p *Postgres) SignIn(user models.UserLogin) (string, error) {
	const op = "postgres.SignIn"

	query := `SELECT users.hash_password, users.role FROM users WHERE email = $1`

	var password, role string
	err := p.database.QueryRow(query, user.Username).Scan(&password, &role)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := auth.GenerateToken(user.Username, role)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}
