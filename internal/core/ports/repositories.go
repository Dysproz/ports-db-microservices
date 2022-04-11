package ports

import "github.com/Dysproz/ports-db-microservices/internal/core/domain"


type PortsRepository interface {
	Get(key string) (domain.Port, error)
	InsertOrUpate(domain.Port) error
}