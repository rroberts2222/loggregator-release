// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/v1/egress.proto

package logcache_v1

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

// EgressClient is the client API for Egress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EgressClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error)
}

type egressClient struct {
	cc grpc.ClientConnInterface
}

func NewEgressClient(cc grpc.ClientConnInterface) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Egress/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *egressClient) Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error) {
	out := new(MetaResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Egress/Meta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EgressServer is the server API for Egress service.
// All implementations must embed UnimplementedEgressServer
// for forward compatibility
type EgressServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Meta(context.Context, *MetaRequest) (*MetaResponse, error)
	mustEmbedUnimplementedEgressServer()
}

// UnimplementedEgressServer must be embedded to have forward compatible implementations.
type UnimplementedEgressServer struct {
}

func (UnimplementedEgressServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedEgressServer) Meta(context.Context, *MetaRequest) (*MetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Meta not implemented")
}
func (UnimplementedEgressServer) mustEmbedUnimplementedEgressServer() {}

// UnsafeEgressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EgressServer will
// result in compilation errors.
type UnsafeEgressServer interface {
	mustEmbedUnimplementedEgressServer()
}

func RegisterEgressServer(s grpc.ServiceRegistrar, srv EgressServer) {
	s.RegisterService(&Egress_ServiceDesc, srv)
}

func _Egress_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Egress_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EgressServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Egress/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EgressServer).Meta(ctx, req.(*MetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Egress_ServiceDesc is the grpc.ServiceDesc for Egress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Egress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Egress_Read_Handler,
		},
		{
			MethodName: "Meta",
			Handler:    _Egress_Meta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/egress.proto",
}
