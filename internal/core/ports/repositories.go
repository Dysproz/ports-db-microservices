package ports

import (
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

type PortsRepository interface {
	Get(key string) (pb.Port, error)
	InsertOrUpate(pb.Port) error
}
