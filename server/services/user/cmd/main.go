package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/log"
	"github.com/jpitchardu/ScaleCommerce/user/pkg/data"
	scaleCommerceMiddleware "github.com/jpitchardu/ScaleCommerce/user/pkg/middleware"
	"github.com/jpitchardu/ScaleCommerce/user/pkg/services"
	"github.com/jpitchardu/ScaleCommerce/user/pkg/transport"
)

var (
	httpAddr    = ":8031"
	httpTimeout = 60 * time.Second
)

func main() {

	fmt.Print("Testing changes")

	logger := log.NewLogfmtLogger(os.Stderr)

	db, err := data.ConnectToDB()

	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}

	service := services.NewUserService(db)
	service = scaleCommerceMiddleware.LoggingMiddleware{
		Logger: logger,
		Next:   service,
	}

	handler := transport.MakeHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(httpTimeout))

	r.Mount("/", handler)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("Listening on ", httpAddr)
		errs <- http.ListenAndServe(httpAddr, r)
	}()

	logger.Log("Exiting: %s", <-errs)

}
