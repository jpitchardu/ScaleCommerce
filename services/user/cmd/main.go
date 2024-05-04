package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jpitchardu/ScaleCommerce/pkg/user"
)

var (
	httpAddr    = ":8031"
	httpTimeout = 60 * time.Second
)

func main() {
	service := user.NewUserService()

	handler := user.MakeHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(httpTimeout))

	r.Mount("/user", handler)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Printf("Listening on %s", httpAddr)
		errs <- http.ListenAndServe(httpAddr, r)
	}()

	log.Printf("Exiting: %s", <-errs)

}
