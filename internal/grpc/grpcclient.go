package grpc

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	"github.com/Dysproz/ports-db-microservices/internal/transport"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a client handling gRPC requests
type Client struct {
	Client     transport.PortServiceClient
	connection *grpc.ClientConn
}

// NewClient returns a new Client
func NewClient(serverAddr string, cancel context.CancelFunc) *Client {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Info("Dialing ", serverAddr, "...")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Error("failed to dial: ", err)
		cancel()
	}
	client := transport.NewPortServiceClient(conn)
	return &Client{Client: client, connection: conn}
}

// CreateOrUpdatePort handles grpc requests for creating or upating port entry in portDomainService
func (c *Client) CreateOrUpdatePort(key string, port domain.Port) error {
	log.Println("CreateOrUpdate ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.Client.CreateOrUpdatePort(ctx, &domain.CreateOrUpdatePortRequest{
		Key:  key,
		Port: &port,
	})
	return err
}

// GetPort handles grpc requests for gathering port data from portDomainService
func (c *Client) GetPort(key string) (domain.Port, error) {
	log.Println("GetPort  ", key, " port")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := c.Client.GetPort(ctx, &domain.GetPortRequest{
		Key: key,
	})
	if err != nil {
		return domain.Port{}, err
	}
	return *response.Port, nil
}

// CloseConnection closes client connection
func (c *Client) CloseConnection() error {
	return c.connection.Close()
}
