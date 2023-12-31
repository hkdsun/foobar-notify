// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/foobar.proto

package proto

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

const (
	FoobarService_TriggerRescan_FullMethodName = "/foobar.FoobarService/TriggerRescan"
)

// FoobarServiceClient is the client API for FoobarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoobarServiceClient interface {
	TriggerRescan(ctx context.Context, in *RescanRequest, opts ...grpc.CallOption) (*RescanResponse, error)
}

type foobarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoobarServiceClient(cc grpc.ClientConnInterface) FoobarServiceClient {
	return &foobarServiceClient{cc}
}

func (c *foobarServiceClient) TriggerRescan(ctx context.Context, in *RescanRequest, opts ...grpc.CallOption) (*RescanResponse, error) {
	out := new(RescanResponse)
	err := c.cc.Invoke(ctx, FoobarService_TriggerRescan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoobarServiceServer is the server API for FoobarService service.
// All implementations must embed UnimplementedFoobarServiceServer
// for forward compatibility
type FoobarServiceServer interface {
	TriggerRescan(context.Context, *RescanRequest) (*RescanResponse, error)
	mustEmbedUnimplementedFoobarServiceServer()
}

// UnimplementedFoobarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoobarServiceServer struct {
}

func (UnimplementedFoobarServiceServer) TriggerRescan(context.Context, *RescanRequest) (*RescanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerRescan not implemented")
}
func (UnimplementedFoobarServiceServer) mustEmbedUnimplementedFoobarServiceServer() {}

// UnsafeFoobarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoobarServiceServer will
// result in compilation errors.
type UnsafeFoobarServiceServer interface {
	mustEmbedUnimplementedFoobarServiceServer()
}

func RegisterFoobarServiceServer(s grpc.ServiceRegistrar, srv FoobarServiceServer) {
	s.RegisterService(&FoobarService_ServiceDesc, srv)
}

func _FoobarService_TriggerRescan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServiceServer).TriggerRescan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FoobarService_TriggerRescan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServiceServer).TriggerRescan(ctx, req.(*RescanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FoobarService_ServiceDesc is the grpc.ServiceDesc for FoobarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoobarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foobar.FoobarService",
	HandlerType: (*FoobarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerRescan",
			Handler:    _FoobarService_TriggerRescan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/foobar.proto",
}
