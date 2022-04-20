package ports

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol" // TODO: geneate it to normal service and domain
)
type JSONParseService interface {
	Load(path string) error
	Watch() <-chan domain.Entry
}

type GRPCClientService interface {
	CreateOrUpdatePort(ctx context.Context, in *pb.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*pb.CreateOrUpdatePortResponse, error)
	GetPort(ctx context.Context, in *pb.GetPortRequest, opts ...grpc.CallOption) (*pb.GetPortResponse, error)
}