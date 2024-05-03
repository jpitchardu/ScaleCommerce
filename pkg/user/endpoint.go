package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

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

		input, isValid := request.(createUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.CreateUser(&UserModel{Name: input.Name, Email: input.Email, Password: input.Password})

		return &createUserResponse{ID}, err
	}
}

type getUserRequest struct {
	ID string
}

type getUserResponse struct {
	User *UserModel
}

func makeGetUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		input, isValid := request.(getUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		user, err := s.GetUser(input.ID)

		return &getUserResponse{User: user}, err
	}
}

type updateUserRequest struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type updateUserResponse struct {
	ID string
}

func makeUpdateUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		input, isValid := request.(updateUserRequest)

		if !isValid {
			return nil, ErrBadRequest
		}

		ID, err := s.UpdateUser(&UserModel{ID: input.ID, Name: input.Name, Email: input.Email, Password: input.Password})

		return &updateUserResponse{ID}, err
	}
}
