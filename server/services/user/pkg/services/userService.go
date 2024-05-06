package services

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel interface {
	CreateUser(user *User) (int64, error)
	GetUser(id int64) (*User, error)
	UpdateUser(user *User) (int64, error)
	DeleteUser(id int64) error
}

type userModel struct{ DB *sql.DB }

func NewUserService(db *sql.DB) UserModel { return &userModel{DB: db} }

func (s *userModel) CreateUser(user *User) (int64, error) {
	var id int64

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	err = s.DB.QueryRow("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, 1) RETURNING id", user.Name, user.Email, string(bytes)).Scan(&id)

	return id, err
}

func (s *userModel) GetUser(id int64) (*User, error) {
	var user User

	err := s.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}

func (s *userModel) UpdateUser(user *User) (int64, error) {
	var id int64

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	err = s.DB.QueryRow("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id", user.Name, user.Email, string(bytes), user.ID).Scan(&id)

	return id, err
}

func (s *userModel) DeleteUser(id int64) error {
	_, err := s.DB.Exec("DELETE FROM users WHERE id = $1", id)

	return err
}
