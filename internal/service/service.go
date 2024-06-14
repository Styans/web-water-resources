package service

import (
	"log"
	"test/internal/models"
	"test/internal/repository"
	"test/internal/service/users"
)

type Service struct {
	UserService models.UserService
}

func NewService(repo *repository.Repository, log *log.Logger) *Service {
	return &Service{UserService: users.NewUserService(repo.UserRepo)}
}
