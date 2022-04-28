package server

import (
	"context"
	"fmt"
	"log"

	protos "github.com/Ankit152/golang-microservices/restful-service-grpc/helloworld/protos"
)

type Greeter struct {
	protos.UnimplementedGreeterServer
}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (c *Greeter) SayHello(ctx context.Context, rr *protos.HelloRequest) (*protos.HelloResponse, error) {
	log.Printf("Recieved: %v\n", rr.GetName())
	return &protos.HelloResponse{Message: fmt.Sprintf("Hello %v", rr.GetName())}, nil
}
