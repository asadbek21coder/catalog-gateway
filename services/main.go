package services

import (
	"fmt"

	"github.com/asadbek21coder/catalog/gateway/config"
	"github.com/asadbek21coder/catalog/gateway/genproto/book_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	Service() book_service.ServiceClient
}

type grpcClients struct {
	service book_service.ServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ServiceHost, conf.ServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		service: book_service.NewServiceClient(connService),
	}, nil
}

func (g *grpcClients) Service() book_service.ServiceClient {
	return g.service
}
