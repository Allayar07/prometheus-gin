package service

import (
	"gin_prometheus/internal/models"
	"gin_prometheus/internal/repository"
)

type AuthSRV struct {
	repo *repository.Auth
}

func NewAuthSRV(repo *repository.Auth) *AuthSRV {
	return &AuthSRV{repo: repo}
}

func (s *AuthSRV) CreateUser(user models.User) error {
	return s.repo.Crete(user)
}
