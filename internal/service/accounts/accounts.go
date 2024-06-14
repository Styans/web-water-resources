package accounts

import (
	"test/internal/models"
)

type AccountsService struct {
	repo models.AccountRepo
}

func NewAccountsService(repo models.AccountRepo) *AccountsService {
	return &AccountsService{repo}
}

func (s *AccountsService) CreateAccruals(account *models.Accounts) error {
	return s.repo.CreateAccruals(account)
}

func (s *AccountsService) GetAccrualsByUserID(userID int) ([]models.AccrualsDTO, error) {
	return s.repo.GetAccrualsByUserID(userID)
}

func (s *AccountsService) CreatePayment(payment *models.PaymentsDTO, id int) error {
	err := s.repo.DeleteAccrualByID(id)
	if err != nil {
		return err
	}
	err = s.repo.CreatePayment(payment)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsService) GetPaymentsByUserID(userID int) ([]models.Payments, error) {
	return s.repo.GetPaymentsByUserID(userID)
}
