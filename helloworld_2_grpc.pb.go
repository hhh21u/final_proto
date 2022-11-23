// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package final_proto

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	PrintGatech(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintReply, error)
	UpdateMaster(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	ReplyFromWorker(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) PrintGatech(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintReply, error) {
	out := new(PrintReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/PrintGatech", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateMaster(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/UpdateMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ReplyFromWorker(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintReply, error) {
	out := new(PrintReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/ReplyFromWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	PrintGatech(context.Context, *PrintRequest) (*PrintReply, error)
	UpdateMaster(context.Context, *GetRequest) (*GetReply, error)
	ReplyFromWorker(context.Context, *PrintRequest) (*PrintReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) PrintGatech(context.Context, *PrintRequest) (*PrintReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintGatech not implemented")
}
func (UnimplementedGreeterServer) UpdateMaster(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaster not implemented")
}
func (UnimplementedGreeterServer) ReplyFromWorker(context.Context, *PrintRequest) (*PrintReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyFromWorker not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_PrintGatech_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).PrintGatech(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/PrintGatech",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).PrintGatech(ctx, req.(*PrintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/UpdateMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateMaster(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ReplyFromWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ReplyFromWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/ReplyFromWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ReplyFromWorker(ctx, req.(*PrintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintGatech",
			Handler:    _Greeter_PrintGatech_Handler,
		},
		{
			MethodName: "UpdateMaster",
			Handler:    _Greeter_UpdateMaster_Handler,
		},
		{
			MethodName: "ReplyFromWorker",
			Handler:    _Greeter_ReplyFromWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld_2.proto",
}
