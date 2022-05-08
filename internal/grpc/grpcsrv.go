package grpc

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	"github.com/Dysproz/ports-db-microservices/internal/core/ports"
	"github.com/Dysproz/ports-db-microservices/internal/transport"

	"google.golang.org/grpc"
)

// PortsProtocolServer is a server handling ports protocol
type PortsProtocolServer struct {
	transport.UnimplementedPortServiceServer
	server   *grpc.Server
	db       ports.Repository
	listener net.Listener
}

// NewPortsProtocolServer returns new portsPortocolServer
func NewPortsProtocolServer(serverPort int, portsRepo ports.Repository, cancel context.CancelFunc) *PortsProtocolServer {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Error(err)
		cancel()
	}

	newServer := &PortsProtocolServer{
		server:   grpcServer,
		db:       portsRepo,
		listener: lis,
	}
	transport.RegisterPortServiceServer(grpcServer, newServer)
	return newServer
}

// CreateOrUpdatePort method creates or updates port information in database
func (s *PortsProtocolServer) CreateOrUpdatePort(ctx context.Context, request *domain.CreateOrUpdatePortRequest) (*domain.CreateOrUpdatePortResponse, error) {
	log.Debug("Got CreateOrUpdatePort request for: ", request.Key, " port")
	err := s.db.InsertOrUpdate(request.Key, *request.Port)
	return &domain.CreateOrUpdatePortResponse{}, err
}

// GetPort method handles requests for port data from database
func (s *PortsProtocolServer) GetPort(ctx context.Context, request *domain.GetPortRequest) (*domain.GetPortResponse, error) {
	log.Debug("Got GetPort request for ", request.Key)
	port, err := s.db.Get(request.Key)
	return &domain.GetPortResponse{
		Port: &port,
	}, err
}

// GracefulStop gracefully stops grpc server and clsoes listener
func (s *PortsProtocolServer) GracefulStop() error {
	log.Info("Stopping gRPC server...")
	s.server.GracefulStop()
	log.Info("Stopping server listener...")
	return s.listener.Close()
}

// Serve serves the grpc server
func (s *PortsProtocolServer) Serve() error {
	return s.server.Serve(s.listener)
}
