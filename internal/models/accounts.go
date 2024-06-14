package models

import "time"

type Accounts struct {
	Date   time.Time
	Past   int
	Last   int
	UserId int
}

type AccrualsDTO struct {
	Id        int
	Date      string
	Past      int
	Last      int
	UserId    int
	Status    bool
	Substract int
	Sum       int
}
type PaymentsDTO struct {
	Id     int
	Date   time.Time
	Sum    int
	UserId int
}

type Payments struct {
	Id     int
	Date   string
	Sum    int
	UserId int
}

type AccountRepo interface {
	CreateAccruals(account *Accounts) error
	DeleteAccrualByID(id int) error
	GetAccrualsByUserID(userID int) ([]AccrualsDTO, error)
	CreatePayment(payment *PaymentsDTO) error
	GetPaymentsByUserID(userID int) ([]Payments, error)
}

type AccountsService interface {
	CreateAccruals(account *Accounts) error

	GetAccrualsByUserID(userID int) ([]AccrualsDTO, error)
	CreatePayment(payment *PaymentsDTO, id int) error
	GetPaymentsByUserID(userID int) ([]Payments, error)
}
