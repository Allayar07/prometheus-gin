package repository

import (
	"context"
	"fmt"
	"gin_prometheus/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Auth struct {
	db *pgxpool.Pool
}

func NewAuth(db *pgxpool.Pool) *Auth {
	return &Auth{db: db}
}

func (r *Auth) Crete(user models.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, phone, address) VALUES ($1, $2, $3)`, "users")
	_, err := r.db.Exec(context.Background(), query, user.Name, user.Phone, user.Address)
	if err != nil {
		return err
	}
	return nil
}
