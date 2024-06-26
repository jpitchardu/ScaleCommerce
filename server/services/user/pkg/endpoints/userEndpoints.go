package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/jpitchardu/ScaleCommerce/user/pkg/services"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID int64
}

func MakeCreateUserEndpoint(s services.UserModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input, isValid := request.(CreateUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.CreateUser(&services.User{Name: input.Name, Email: input.Email, Password: input.Password})

		return &CreateUserResponse{ID}, err
	}
}

type GetUserRequest struct {
	ID int64
}

type GetUserResponse struct {
	ID    int64
	Name  string
	Email string
}

func MakeGetUserEndpoint(s services.UserModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		input, isValid := request.(GetUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		user, err := s.GetUser(input.ID)

		return &GetUserResponse{ID: user.ID, Name: user.Name, Email: user.Email}, err
	}
}

type UpdateUserRequest struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type UpdateUserResponse struct {
	ID int64
}

func MakeUpdateUserEndpoint(s services.UserModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		input, isValid := request.(UpdateUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.UpdateUser(&services.User{ID: input.ID, Name: input.Name, Email: input.Email, Password: input.Password})

		return &UpdateUserResponse{ID}, err
	}
}

type DeleteUserRequest struct {
	ID int64
}

type DeleteUserResponse struct{}

func MakeDeleteUserEndpoint(s services.UserModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		input, isValid := request.(DeleteUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		err = s.DeleteUser(input.ID)

		return &DeleteUserResponse{}, err
	}
}

var ErrBadRequest = errors.New("bad request")
