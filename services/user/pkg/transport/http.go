package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jpitchardu/ScaleCommerce/pkg/endpoints"
	"github.com/jpitchardu/ScaleCommerce/pkg/services"
)

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}

	return req, nil
}

func encodeCreateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoints.CreateUserResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return nil, ErrBadRequest
	}

	return endpoints.GetUserRequest{ID: id}, nil
}

func encodeGetUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoints.GetUserResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}

	return req, nil
}

func encodeUpdateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoints.UpdateUserResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return nil, ErrBadRequest
	}

	return endpoints.DeleteUserRequest{ID: id}, nil
}

func encodeDeleteUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*endpoints.DeleteUserResponse)
	return json.NewEncoder(w).Encode(res)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func MakeHandler(s services.UserService) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	createUserHandler := httptransport.NewServer(
		endpoints.MakeCreateUserEndpoint(s),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		options...,
	)

	getUserHandler := httptransport.NewServer(
		endpoints.MakeGetUserEndpoint(s),
		decodeGetUserRequest,
		encodeGetUserResponse,
		options...,
	)

	updateUserHandler := httptransport.NewServer(
		endpoints.MakeUpdateUserEndpoint(s),
		decodeUpdateUserRequest,
		encodeUpdateUserResponse,
		options...,
	)

	deleteUserHandler := httptransport.NewServer(
		endpoints.MakeDeleteUserEndpoint(s),
		decodeDeleteUserRequest,
		encodeDeleteUserResponse,
		options...,
	)

	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/", createUserHandler.ServeHTTP)
		r.Get((`/{id:\d+}`), getUserHandler.ServeHTTP)
		r.Put((`/{id:\d+}`), updateUserHandler.ServeHTTP)
		r.Delete((`/{id:\d+}`), deleteUserHandler.ServeHTTP)

	})

	return r
}

var ErrBadRequest = errors.New("bad request")
