package grpcclient

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	"github.com/Dysproz/ports-db-microservices/internal/core/ports"
)

// GrpcClient is a client handling gRPC requests
type GrpcClient struct {
	client ports.GRPCClientService
}

// NewGrpcClient returns a new GrpcClient
func NewGrpcClient(client ports.GRPCClientService) *GrpcClient {
	return &GrpcClient{client: client}
}

// CreateOrUpdatePort handles grpc requests for creating or upating port entry in portDomainService
func (c *GrpcClient) CreateOrUpdatePort(key string, port domain.Port) error {
	log.Println("CreateOrUpdate ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.client.CreateOrUpdatePort(ctx, &domain.CreateOrUpdatePortRequest{
		Key:  key,
		Port: &port,
	})
	return err
}

// GetPort handles grpc requests for gathering port data from portDomainService
func (c *GrpcClient) GetPort(key string) (domain.Port, error) {
	log.Println("GetPort  ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := c.client.GetPort(ctx, &domain.GetPortRequest{
		Key: key,
	})
	if err != nil {
		return domain.Port{}, err
	}
	return *response.Port, nil
}
