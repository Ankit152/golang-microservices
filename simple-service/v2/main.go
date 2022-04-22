package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ankit152/golang-microservices/simple-service/v2/handlers"
)

func main() {
	l := log.New(os.Stdout, "test-api-", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	s.ListenAndServe()

}
