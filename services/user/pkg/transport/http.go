package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

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

	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/", createUserHandler.ServeHTTP)
	})

	return r
}

var ErrBadRequest = errors.New("bad request")
