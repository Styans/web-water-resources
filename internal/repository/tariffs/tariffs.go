package tariffs

import (
	"database/sql"
	"fmt"
	"test/internal/models"
)

type TariffsStorage struct {
	db *sql.DB
}

func NewTariffsStorage(db *sql.DB) *TariffsStorage {
	return &TariffsStorage{db: db}
}

func (s *TariffsStorage) GetAlltariffs() (*[]models.Tariffs, error) {
	rows, err := s.db.Query("SELECT id, name, sum FROM tariffs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tariffs []models.Tariffs
	for rows.Next() {
		tariff := new(models.Tariffs)
		err := rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.Sum,
		)
		if err != nil {
			return nil, err
		}
		tariffs = append(tariffs, *tariff)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &tariffs, nil
}

func (s *TariffsStorage) GetSumByName(name string) (float64, error) {
	var sum float64
	query := "SELECT sum FROM tariffs WHERE name = ?"

	// Выполняем запрос с параметром name
	err := s.db.QueryRow(query, name).Scan(&sum)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если не найдено ни одной строки
			return 0, fmt.Errorf("no tariff found with name %s", name)
		}
		return 0, err
	}

	return sum, nil
}
