package ports

import "github.com/Dysproz/ports-db-microservices/internal/core/domain"

type JSONParseService interface {
	Load(path string) error
	Watch() <-chan domain.Entry
}
