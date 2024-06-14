package tariffs

import (
	"test/internal/models"
)

type TariffsService struct {
	repo models.TariffsRepo
}

func NewTariffsService(repo models.TariffsRepo) *TariffsService {
	return &TariffsService{repo}
}

func (s *TariffsService) GetAlltariffs() (*[]models.Tariffs, error) {
	return s.repo.GetAlltariffs()
}

func (s *TariffsService) GetSumByName(name string) (float64, error) {
	return s.repo.GetSumByName(name)
}
