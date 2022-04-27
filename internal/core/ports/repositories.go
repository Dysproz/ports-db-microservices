package ports

import (
	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

// Repository is a interface for repository storing ports data
type Repository interface {
	Get(key string) (domain.Port, error)
	InsertOrUpdate(string, domain.Port) error
}
