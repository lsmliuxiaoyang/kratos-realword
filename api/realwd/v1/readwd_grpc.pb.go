// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: realwd/v1/readwd.proto

package v1

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
	Readwd_SayHello_FullMethodName = "/readwd.v1.Readwd/SayHello"
	Readwd_Login_FullMethodName    = "/readwd.v1.Readwd/Login"
)

// ReadwdClient is the client API for Readwd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReadwdClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type readwdClient struct {
	cc grpc.ClientConnInterface
}

func NewReadwdClient(cc grpc.ClientConnInterface) ReadwdClient {
	return &readwdClient{cc}
}

func (c *readwdClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Readwd_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readwdClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Readwd_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadwdServer is the server API for Readwd service.
// All implementations must embed UnimplementedReadwdServer
// for forward compatibility
type ReadwdServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedReadwdServer()
}

// UnimplementedReadwdServer must be embedded to have forward compatible implementations.
type UnimplementedReadwdServer struct {
}

func (UnimplementedReadwdServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedReadwdServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedReadwdServer) mustEmbedUnimplementedReadwdServer() {}

// UnsafeReadwdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReadwdServer will
// result in compilation errors.
type UnsafeReadwdServer interface {
	mustEmbedUnimplementedReadwdServer()
}

func RegisterReadwdServer(s grpc.ServiceRegistrar, srv ReadwdServer) {
	s.RegisterService(&Readwd_ServiceDesc, srv)
}

func _Readwd_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadwdServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Readwd_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadwdServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Readwd_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadwdServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Readwd_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadwdServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Readwd_ServiceDesc is the grpc.ServiceDesc for Readwd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Readwd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "readwd.v1.Readwd",
	HandlerType: (*ReadwdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Readwd_SayHello_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Readwd_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realwd/v1/readwd.proto",
}