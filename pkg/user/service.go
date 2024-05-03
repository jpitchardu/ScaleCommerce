package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type UserService interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type User struct {
	_ID      string
	Name     string
	Email    string
	Password string
}

type userService struct{}

func NewUserService() UserService { return &userService{} }

func (s *userService) CreateUser(user *User) error {

	return nil
}

func (s *userService) GetUser(id string) (*User, error) {

	return nil, nil
}

func (s *userService) UpdateUser(user *User) error {
	return nil
}

func (s *userService) DeleteUser(id string) error {
	return nil
}

var ErrInvalidUser = errors.New("Invalid user")

type createUserRequest struct {
	Name     string
	Email    string
	Password string
}

type createUserResponse struct {
	ID string
}

func makeCreateUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil
	}
}
