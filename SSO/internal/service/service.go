package service

import "SummerHSE/sso/internal/repository"

type Service struct {
	repo *repository.Repository
}

func BuildService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
