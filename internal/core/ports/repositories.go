package ports

import (
	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

type PortsRepository interface {
	Get(key string) (domain.Port, error)
	InsertOrUpdate(string, domain.Port) error
}
