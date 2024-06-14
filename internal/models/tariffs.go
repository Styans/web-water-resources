package models

type Tariffs struct {
	Id   int
	Name string
	Sum  int
}

type TariffsRepo interface {
	GetAlltariffs() (*[]Tariffs, error)
	GetSumByName(name string) (float64, error)
}

type TariffsService interface {
	GetAlltariffs() (*[]Tariffs, error)
	GetSumByName(name string) (float64, error)
}
