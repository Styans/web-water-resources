package repository

import (
	"database/sql"
	"test/internal/models"
	"test/internal/repository/accounts"
	"test/internal/repository/tariffs"
	"test/internal/repository/users"
)

type Repository struct {
	UserRepo     models.UserRepo
	AccountsRepo models.AccountRepo
	TariffsRepo  models.TariffsRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo:     users.NewUserStorage(db),
		AccountsRepo: accounts.NewAccountStorage(db),
		TariffsRepo:  tariffs.NewTariffsStorage(db),
	}
}
