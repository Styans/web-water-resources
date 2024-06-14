package service

import (
	"log"
	"test/internal/models"
	"test/internal/repository"
	"test/internal/service/accounts"
	"test/internal/service/tariffs"
	"test/internal/service/users"
)

type Service struct {
	UserService     models.UserService
	AccountsService models.AccountsService
	TariffsService  models.TariffsService
}

func NewService(repo *repository.Repository, log *log.Logger) *Service {
	return &Service{
		UserService:     users.NewUserService(repo.UserRepo),
		AccountsService: accounts.NewAccountsService(repo.AccountsRepo),
		TariffsService:  tariffs.NewTariffsService(repo.TariffsRepo),
	}
}
