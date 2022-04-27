package ports

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

// JSONParseService is an interface for JSON parsing service
type JSONParseService interface {
	Load(path string)
	Watch() <-chan domain.Entry
}

// GRPCClientService is an interface for gRPC client service
type GRPCClientService interface {
	CreateOrUpdatePort(ctx context.Context, in *domain.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*domain.CreateOrUpdatePortResponse, error)
	GetPort(ctx context.Context, in *domain.GetPortRequest, opts ...grpc.CallOption) (*domain.GetPortResponse, error)
}
