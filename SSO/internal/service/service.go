package service

import (
	"SSO/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func BuildService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
