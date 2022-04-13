package grpcclient

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

type grpcClient struct {
	client pb.PortServiceClient
}

func NewGrpcClient(client pb.PortServiceClient) *grpcClient {
	return &grpcClient{client: client}
}

// CreateOrUpdatePort handles grpc requests for creating or upating port entry in portDomainService
func (c *grpcClient) CreateOrUpdatePort(key string, port pb.Port) error {
	log.Println("CreateOrUpdate ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.client.CreateOrUpdatePort(ctx, &pb.CreateOrUpdatePortRequest{
		Key:  key,
		Port: &port,
	})
	return err
}

// GetPort handles grpc requests for gathering port data from portDomainService
func (c *grpcClient) GetPort(key string) (pb.Port, error) {
	log.Println("GetPort  ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := c.client.GetPort(ctx, &pb.GetPortRequest{
		Key: key,
	})
	if err != nil {
		return pb.Port{}, err
	}
	return *response.Port, nil
}
