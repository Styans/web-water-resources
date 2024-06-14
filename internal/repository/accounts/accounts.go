package accounts

import (
	"database/sql"
	"strings"
	"test/internal/models"
	"time"
)

type AccountsStorage struct {
	db *sql.DB
}

func NewAccountStorage(db *sql.DB) *AccountsStorage {
	return &AccountsStorage{db: db}
}

func (s *AccountsStorage) CreateAccruals(account *models.Accounts) error {
	insertData := `INSERT INTO accruals (date, past, last, user_id, status) VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(
		insertData,
		account.Date,
		account.Past,
		account.Last,
		account.UserId,
		false,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsStorage) DeleteAccrualByID(id int) error {
	deleteData := `DELETE FROM accruals WHERE id = ?`

	_, err := s.db.Exec(deleteData, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsStorage) GetAccrualsByUserID(userID int) ([]models.AccrualsDTO, error) {
	query := `SELECT id, date, past, last, user_id FROM accruals WHERE user_id = ?`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.AccrualsDTO

	for rows.Next() {
		var account models.AccrualsDTO
		var dateStr string
		err := rows.Scan(
			&account.Id,
			&dateStr,
			&account.Past,
			&account.Last,
			&account.UserId,
		)
		if err != nil {
			return nil, err
		}
		dateOnly := strings.SplitN(dateStr, "T", 2)[0]

		Date, err := time.Parse("2006-01-02", dateOnly)
		account.Date = Date.Format("2006-01-02")
		if err != nil {
			return nil, err
		}
		account.Substract = account.Last - account.Past
		account.Sum = int(float64(account.Substract) * 28.22)
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *AccountsStorage) CreatePayment(payment *models.PaymentsDTO) error {
	insertData := `INSERT INTO payments (date, sum, user_id) VALUES (?, ?, ?)`

	_, err := s.db.Exec(
		insertData,
		payment.Date,
		payment.Sum,
		payment.UserId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsStorage) GetPaymentsByUserID(userID int) ([]models.Payments, error) {
	query := `SELECT id, date, sum, user_id FROM payments WHERE user_id = ?`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payments

	for rows.Next() {
		var payment models.Payments
		var dateStr string

		err := rows.Scan(
			&payment.Id,
			&dateStr,
			&payment.Sum,
			&payment.UserId,
		)
		if err != nil {
			return nil, err
		}
		dateOnly := strings.SplitN(dateStr, "T", 2)[0]

		Date, err := time.Parse("2006-01-02", dateOnly)
		payment.Date = Date.Format("2006-01-02")
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}
