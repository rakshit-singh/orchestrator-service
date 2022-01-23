// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: Service1.proto

package __

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

// ProcessRequestClient is the client API for ProcessRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessRequestClient interface {
	GetUserByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error)
}

type processRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessRequestClient(cc grpc.ClientConnInterface) ProcessRequestClient {
	return &processRequestClient{cc}
}

func (c *processRequestClient) GetUserByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ProcessRequest/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessRequestServer is the server API for ProcessRequest service.
// All implementations must embed UnimplementedProcessRequestServer
// for forward compatibility
type ProcessRequestServer interface {
	GetUserByName(context.Context, *UserName) (*User, error)
	mustEmbedUnimplementedProcessRequestServer()
}

// UnimplementedProcessRequestServer must be embedded to have forward compatible implementations.
type UnimplementedProcessRequestServer struct {
}

func (UnimplementedProcessRequestServer) GetUserByName(context.Context, *UserName) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedProcessRequestServer) mustEmbedUnimplementedProcessRequestServer() {}

// UnsafeProcessRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessRequestServer will
// result in compilation errors.
type UnsafeProcessRequestServer interface {
	mustEmbedUnimplementedProcessRequestServer()
}

func RegisterProcessRequestServer(s grpc.ServiceRegistrar, srv ProcessRequestServer) {
	s.RegisterService(&ProcessRequest_ServiceDesc, srv)
}

func _ProcessRequest_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessRequestServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessRequest/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessRequestServer).GetUserByName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessRequest_ServiceDesc is the grpc.ServiceDesc for ProcessRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProcessRequest",
	HandlerType: (*ProcessRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _ProcessRequest_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Service1.proto",
}