// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hasher

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

// HashingServiceClient is the client API for HashingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashingServiceClient interface {
	CheckHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	GetHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	CreateHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
}

type hashingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHashingServiceClient(cc grpc.ClientConnInterface) HashingServiceClient {
	return &hashingServiceClient{cc}
}

func (c *hashingServiceClient) CheckHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, "/hasher.HashingService/CheckHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashingServiceClient) GetHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, "/hasher.HashingService/GetHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashingServiceClient) CreateHash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, "/hasher.HashingService/CreateHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashingServiceServer is the server API for HashingService service.
// All implementations must embed UnimplementedHashingServiceServer
// for forward compatibility
type HashingServiceServer interface {
	CheckHash(context.Context, *HashRequest) (*HashResponse, error)
	GetHash(context.Context, *HashRequest) (*HashResponse, error)
	CreateHash(context.Context, *HashRequest) (*HashResponse, error)
	mustEmbedUnimplementedHashingServiceServer()
}

// UnimplementedHashingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHashingServiceServer struct {
}

func (UnimplementedHashingServiceServer) CheckHash(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHash not implemented")
}
func (UnimplementedHashingServiceServer) GetHash(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHash not implemented")
}
func (UnimplementedHashingServiceServer) CreateHash(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHash not implemented")
}
func (UnimplementedHashingServiceServer) mustEmbedUnimplementedHashingServiceServer() {}

// UnsafeHashingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashingServiceServer will
// result in compilation errors.
type UnsafeHashingServiceServer interface {
	mustEmbedUnimplementedHashingServiceServer()
}

func RegisterHashingServiceServer(s grpc.ServiceRegistrar, srv HashingServiceServer) {
	s.RegisterService(&HashingService_ServiceDesc, srv)
}

func _HashingService_CheckHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashingServiceServer).CheckHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hasher.HashingService/CheckHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashingServiceServer).CheckHash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashingService_GetHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashingServiceServer).GetHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hasher.HashingService/GetHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashingServiceServer).GetHash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashingService_CreateHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashingServiceServer).CreateHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hasher.HashingService/CreateHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashingServiceServer).CreateHash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HashingService_ServiceDesc is the grpc.ServiceDesc for HashingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hasher.HashingService",
	HandlerType: (*HashingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHash",
			Handler:    _HashingService_CheckHash_Handler,
		},
		{
			MethodName: "GetHash",
			Handler:    _HashingService_GetHash_Handler,
		},
		{
			MethodName: "CreateHash",
			Handler:    _HashingService_CreateHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hash_service.proto",
}
