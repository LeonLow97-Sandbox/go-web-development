package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LeonLow97/router"
	"github.com/LeonLow97/store"
)

func main() {

	// setup database connection
	db := store.New()

	r := router.NewRouter(db)

	s := http.Server{
		Addr:         "127.0.0.1:9000",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Second,
		Handler:      r,
	}

	errs := make(chan error, 2)

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminating API with error: %s", <-errs)

	// Shutdown strategy (graceful shutdown)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
