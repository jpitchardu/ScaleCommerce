package middleware

import (
	"time"

	"github.com/go-kit/log"
	"github.com/jpitchardu/ScaleCommerce/pkg/services"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   services.UserService
}

func (mv LoggingMiddleware) CreateUser(user *services.UserModel) (output int64, err error) {
	defer func(begin time.Time) {
		mv.Logger.Log(
			"method", "CreateUser",
			"user", user,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mv.Next.CreateUser(user)
	return
}

func (mv LoggingMiddleware) GetUser(id int64) (*services.UserModel, error) {

	return nil, nil
}

func (mv LoggingMiddleware) UpdateUser(user *services.UserModel) (int64, error) {
	return user.ID, nil
}

func (mv LoggingMiddleware) DeleteUser(id int64) error {
	return nil
}
