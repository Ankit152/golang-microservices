package server

import (
	"context"
	"log"

	protos "github.com/Ankit152/golang-microservices/restful-service-grpc/currency/protos"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
}

func NewCurrency() *Currency {
	return &Currency{}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	log.Println("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
