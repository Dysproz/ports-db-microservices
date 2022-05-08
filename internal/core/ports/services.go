package ports

import (
	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

// DataInputService is an interface for JSON parsing service
type DataInputService interface {
	Load(path string)
	Watch() <-chan domain.Entry
}
