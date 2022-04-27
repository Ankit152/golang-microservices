package main

import (
	"fmt"
	"net"
	"os"

	protos "github.com/Ankit152/golang-microservices/restful-service-grpc/currency/protos"
	"github.com/Ankit152/golang-microservices/restful-service-grpc/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency()
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	log.Info("server listening at %v", 8080)
	// listen for requests
	gs.Serve(l)
}
