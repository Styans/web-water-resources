package models

type ReceptsService interface {
	CreateAccruals(AccrualsDTO) error
	CreateRes(pay PaymentsDTO) error
}
