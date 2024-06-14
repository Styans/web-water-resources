package users

import (
	"database/sql"
	"test/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, secondname, patronymic, benefits, status, districts, addr) VALUES ($2, $3, $4, $5, $6, $7, $8)",
		user.Name,
		user.Secondname,
		user.Patronymic,
		user.Benefits,
		false,
		user.Districts,
		user.Addr,
	)
	if err != nil {
		switch err.Error() {
		case "UNIQUE constraint failed: users.email":
			return models.ErrDuplicateEmail
		case "UNIQUE constraint failed: users.username":
			return models.ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (s *UserStorage) GetAllUsers() ([]*models.User, error) {
	rows, err := s.db.Query("SELECT id, name, secondname, patronymic, benefits, status, districts, addr FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.Name, &user.Secondname, &user.Patronymic, &user.Benefits, &user.Status, &user.Districts, &user.Addr)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStorage) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.Id,
		&user.Name,
		&user.Secondname,
		&user.Patronymic,
		&user.Benefits,
		&user.Status,
		&user.Districts,
		&user.Addr,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
