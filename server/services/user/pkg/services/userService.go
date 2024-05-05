package services

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(user *UserModel) (int64, error)
	GetUser(id int64) (*UserModel, error)
	UpdateUser(user *UserModel) (int64, error)
	DeleteUser(id int64) error
}

type userService struct{ DB *sql.DB }

func NewUserService(db *sql.DB) UserService { return &userService{DB: db} }

func (s *userService) CreateUser(user *UserModel) (int64, error) {

	var id int64

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	err = s.DB.QueryRow("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, 1) RETURNING id", user.Name, user.Email, string(bytes)).Scan(&id)

	return id, err
}

func (s *userService) GetUser(id int64) (*UserModel, error) {

	var user UserModel

	err := s.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}

func (s *userService) UpdateUser(user *UserModel) (int64, error) {

	var id int64

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	err = s.DB.QueryRow("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id", user.Name, user.Email, string(bytes), user.ID).Scan(&id)

	return id, err
}

func (s *userService) DeleteUser(id int64) error {

	_, err := s.DB.Exec("DELETE FROM users WHERE id = $1", id)

	return err
}