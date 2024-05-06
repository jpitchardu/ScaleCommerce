package middleware

import (
	"time"

	"github.com/go-kit/log"
	"github.com/jpitchardu/ScaleCommerce/user/pkg/services"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   services.UserModel
}

func (mv LoggingMiddleware) CreateUser(user *services.User) (output int64, err error) {
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

func (mv LoggingMiddleware) GetUser(id int64) (output *services.User, err error) {
	defer func(begin time.Time) {
		mv.Logger.Log(
			"method", "GetUser",
			"id", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mv.Next.GetUser(id)
	return
}

func (mv LoggingMiddleware) UpdateUser(user *services.User) (output int64, err error) {
	defer func(begin time.Time) {
		mv.Logger.Log(
			"method", "UpdateUser",
			"user", user,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mv.Next.UpdateUser(user)
	return
}

func (mv LoggingMiddleware) DeleteUser(id int64) (err error) {
	defer func(begin time.Time) {
		mv.Logger.Log(
			"method", "DeleteUser",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mv.Next.DeleteUser(id)
	return
}
