package storage

import (
	"fmt"
	"microService/internal/models"
)

func (p *Postgres) CreateUser(user models.CreateUserReq) error {
	const op = "postgres.CreateUser"
	fmt.Println(user.Password)
	query := `INSERT INTO users (email, hash_password, role) VALUES ($1, $2, $3)`
	_, err := p.database.Query(query, user.Username, user.Password, "admin")
	if err != nil {
		return fmt.Errorf("%s:%d", op, err)
	}

	return nil
}
