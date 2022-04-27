package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/Ankit152/golang-microservices/restful-service-grpc/currency/protos"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	protos.UnimplementedCurrencyServer
}

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	//c := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, &Server{})

	reflection.Register(gs)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}
