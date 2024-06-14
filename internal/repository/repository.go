package repository

import (
	"database/sql"
	"test/internal/models"
	"test/internal/repository/users"
)

type Repository struct {
	UserRepo models.UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{UserRepo: users.NewUserStorage(db)}
}
