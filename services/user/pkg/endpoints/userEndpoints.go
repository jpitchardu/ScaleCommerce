package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/jpitchardu/ScaleCommerce/pkg/services"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID int64
}

func MakeCreateUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		input, isValid := request.(CreateUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.CreateUser(&services.UserModel{Name: input.Name, Email: input.Email, Password: input.Password})

		return &CreateUserResponse{ID}, err
	}
}

type GetUserRequest struct {
	ID int64
}

type GetUserResponse struct {
	User *services.UserModel
}

func MakeGetUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		input, isValid := request.(GetUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		user, err := s.GetUser(input.ID)

		return &GetUserResponse{User: user}, err
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

func MakeUpdateUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		input, isValid := request.(UpdateUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.UpdateUser(&services.UserModel{ID: input.ID, Name: input.Name, Email: input.Email, Password: input.Password})

		return &UpdateUserResponse{ID}, err
	}
}

var ErrBadRequest = errors.New("bad request")
