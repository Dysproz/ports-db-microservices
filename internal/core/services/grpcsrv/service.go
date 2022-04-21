package grpcsrv

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	"github.com/Dysproz/ports-db-microservices/internal/core/ports"
	"github.com/Dysproz/ports-db-microservices/internal/core/services/portsprotocol"
)

type portsProtocolServer struct {
	portsprotocol.UnimplementedPortServiceServer
	db ports.PortsRepository
}

// NewPortsProtocolServer returns new portsPortocolServer
func NewPortsProtocolServer(portsRepo ports.PortsRepository) *portsProtocolServer {
	return &portsProtocolServer{
		db: portsRepo,
	}
}

// CreateOrUpdatePort method creates or updates port information in database
func (s *portsProtocolServer) CreateOrUpdatePort(ctx context.Context, request *domain.CreateOrUpdatePortRequest) (*domain.CreateOrUpdatePortResponse, error) {
	log.Debug("Got CreateOrUpdatePort request for: ", request.Key, " port")
	err := s.db.InsertOrUpdate(request.Key, *request.Port)
	return &domain.CreateOrUpdatePortResponse{}, err
}

// GetPort method handles requests for port data from database
func (s *portsProtocolServer) GetPort(ctx context.Context, request *domain.GetPortRequest) (*domain.GetPortResponse, error) {
	log.Debug("Got GetPort request for ", request.Key)
	port, err := s.db.Get(request.Key)
	return &domain.GetPortResponse{
		Port: &port,
	}, err
}