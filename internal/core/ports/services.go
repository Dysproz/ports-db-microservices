package ports

import (
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)
type JSONParseService interface {
	Load(path string) error
	Watch() <-chan domain.Entry
}

type GRPCClientService interface {
	CreateOrUpdatePort(key string, port pb.Port) error
	GetPort(key string) (pb.Port, error)
}