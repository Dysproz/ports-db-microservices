package portsprotocolserver

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/pkg/mongodb"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

// PortsProtocolServer implements gRPC server interface.
type PortsProtocolServer struct {
	pb.UnimplementedPortServiceServer
	MongoDB mongodb.MongoClient
}

// CreateOrUpdatePort method creates or updates port information in database
func (s *PortsProtocolServer) CreateOrUpdatePort(ctx context.Context, request *pb.CreateOrUpdatePortRequest) (*pb.CreateOrUpdatePortResponse, error) {
	log.Debug("Got CreateOrUpdatePort request for: ", request.Key, " port")
	err := s.MongoDB.InsertOrUpdatePort(request.Key, *request.Port)
	return &pb.CreateOrUpdatePortResponse{}, err
}

// GetPort method handles requests for port data from database
func (s *PortsProtocolServer) GetPort(ctx context.Context, request *pb.GetPortRequest) (*pb.GetPortResponse, error) {
	log.Debug("Got GetPort request for ", request.Key)
	port, err := s.MongoDB.GetPort(request.Key)
	return &pb.GetPortResponse{
		Port: &port,
	}, err
}
