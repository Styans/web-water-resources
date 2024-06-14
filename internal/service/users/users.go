package users

import (
	"test/internal/models"
)

type UserService struct {
	repo models.UserRepo
}

func NewUserService(repo models.UserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAllUsers()
}
