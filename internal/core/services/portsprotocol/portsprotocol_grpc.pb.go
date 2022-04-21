// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: portsprotocol.proto

package portsprotocol

import (
	context "context"
	domain "github.com/Dysproz/ports-db-microservices/internal/core/domain"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortServiceClient is the client API for PortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortServiceClient interface {
	CreateOrUpdatePort(ctx context.Context, in *domain.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*domain.CreateOrUpdatePortResponse, error)
	GetPort(ctx context.Context, in *domain.GetPortRequest, opts ...grpc.CallOption) (*domain.GetPortResponse, error)
}

type portServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortServiceClient(cc grpc.ClientConnInterface) PortServiceClient {
	return &portServiceClient{cc}
}

func (c *portServiceClient) CreateOrUpdatePort(ctx context.Context, in *domain.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*domain.CreateOrUpdatePortResponse, error) {
	out := new(domain.CreateOrUpdatePortResponse)
	err := c.cc.Invoke(ctx, "/portsprotocol.PortService/CreateOrUpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portServiceClient) GetPort(ctx context.Context, in *domain.GetPortRequest, opts ...grpc.CallOption) (*domain.GetPortResponse, error) {
	out := new(domain.GetPortResponse)
	err := c.cc.Invoke(ctx, "/portsprotocol.PortService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortServiceServer is the server API for PortService service.
// All implementations must embed UnimplementedPortServiceServer
// for forward compatibility
type PortServiceServer interface {
	CreateOrUpdatePort(context.Context, *domain.CreateOrUpdatePortRequest) (*domain.CreateOrUpdatePortResponse, error)
	GetPort(context.Context, *domain.GetPortRequest) (*domain.GetPortResponse, error)
	mustEmbedUnimplementedPortServiceServer()
}

// UnimplementedPortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortServiceServer struct {
}

func (UnimplementedPortServiceServer) CreateOrUpdatePort(context.Context, *domain.CreateOrUpdatePortRequest) (*domain.CreateOrUpdatePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdatePort not implemented")
}
func (UnimplementedPortServiceServer) GetPort(context.Context, *domain.GetPortRequest) (*domain.GetPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortServiceServer) mustEmbedUnimplementedPortServiceServer() {}

// UnsafePortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortServiceServer will
// result in compilation errors.
type UnsafePortServiceServer interface {
	mustEmbedUnimplementedPortServiceServer()
}

func RegisterPortServiceServer(s grpc.ServiceRegistrar, srv PortServiceServer) {
	s.RegisterService(&PortService_ServiceDesc, srv)
}

func _PortService_CreateOrUpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.CreateOrUpdatePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).CreateOrUpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsprotocol.PortService/CreateOrUpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortServiceServer).CreateOrUpdatePort(ctx, req.(*domain.CreateOrUpdatePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.GetPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsprotocol.PortService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortServiceServer).GetPort(ctx, req.(*domain.GetPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortService_ServiceDesc is the grpc.ServiceDesc for PortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "portsprotocol.PortService",
	HandlerType: (*PortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdatePort",
			Handler:    _PortService_CreateOrUpdatePort_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _PortService_GetPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portsprotocol.proto",
}