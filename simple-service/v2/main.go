package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Ankit152/golang-microservice/simple-service/v2/handlers"
)

func main() {
	l := log.New(os.Stdout, "test-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)

	http.ListenAndServe(":8080", sm)

}
