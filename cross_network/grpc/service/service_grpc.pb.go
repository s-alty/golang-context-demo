// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NoopClient is the client API for Noop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoopClient interface {
	Noop(ctx context.Context, in *NoopRequest, opts ...grpc.CallOption) (*NoopResponse, error)
}

type noopClient struct {
	cc grpc.ClientConnInterface
}

func NewNoopClient(cc grpc.ClientConnInterface) NoopClient {
	return &noopClient{cc}
}

func (c *noopClient) Noop(ctx context.Context, in *NoopRequest, opts ...grpc.CallOption) (*NoopResponse, error) {
	out := new(NoopResponse)
	err := c.cc.Invoke(ctx, "/Noop/Noop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoopServer is the server API for Noop service.
// All implementations must embed UnimplementedNoopServer
// for forward compatibility
type NoopServer interface {
	Noop(context.Context, *NoopRequest) (*NoopResponse, error)
	mustEmbedUnimplementedNoopServer()
}

// UnimplementedNoopServer must be embedded to have forward compatible implementations.
type UnimplementedNoopServer struct {
}

func (UnimplementedNoopServer) Noop(context.Context, *NoopRequest) (*NoopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Noop not implemented")
}
func (UnimplementedNoopServer) mustEmbedUnimplementedNoopServer() {}

// UnsafeNoopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoopServer will
// result in compilation errors.
type UnsafeNoopServer interface {
	mustEmbedUnimplementedNoopServer()
}

func RegisterNoopServer(s grpc.ServiceRegistrar, srv NoopServer) {
	s.RegisterService(&Noop_ServiceDesc, srv)
}

func _Noop_Noop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoopServer).Noop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Noop/Noop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoopServer).Noop(ctx, req.(*NoopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Noop_ServiceDesc is the grpc.ServiceDesc for Noop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Noop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Noop",
	HandlerType: (*NoopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Noop",
			Handler:    _Noop_Noop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
