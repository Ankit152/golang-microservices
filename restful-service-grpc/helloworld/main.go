package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/Ankit152/golang-microservices/restful-service-grpc/helloworld/protos"
	server "github.com/Ankit152/golang-microservices/restful-service-grpc/helloworld/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	helloServer := server.NewGreeter()

	protos.RegisterGreeterServer(grpcServer, helloServer)
	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("Unable to create a listerner at 8080! Error: %v\n", err)
	}
	log.Println("Listening and serving at Port 8080!")
	grpcServer.Serve(l)
}
