package models

type User struct {
	Id         int
	Name       string
	Secondname string
	Patronymic string
	Benefits   string
	Status     bool
	Districts  string
	Addr       string
}

type UserService interface {
	CreateUser(*User) error
	GetAllUsers() ([]*User, error)
}

type UserRepo interface {
	CreateUser(*User) error
	GetAllUsers() ([]*User, error)
}
