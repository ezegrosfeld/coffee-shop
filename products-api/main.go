package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "user-service", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/products", ph)

	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Gracefull shutdown: ", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

	defer cancel()
}