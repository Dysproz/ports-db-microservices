package ports

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)
type JSONParseService interface {
	Load(path string) error
	Watch() <-chan domain.Entry
}

type GRPCClientService interface {
	CreateOrUpdatePort(ctx context.Context, in *domain.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*domain.CreateOrUpdatePortResponse, error)
	GetPort(ctx context.Context, in *domain.GetPortRequest, opts ...grpc.CallOption) (*domain.GetPortResponse, error)
}
