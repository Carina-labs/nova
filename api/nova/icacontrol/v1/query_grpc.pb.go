// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nova/icacontrol/v1/query.proto

package icacontrolv1

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// AllZones returns all the zones registered.
	AllZones(ctx context.Context, in *QueryAllZonesRequest, opts ...grpc.CallOption) (*QueryAllZonesResponse, error)
	Zone(ctx context.Context, in *QueryZoneRequest, opts ...grpc.CallOption) (*QueryZoneResponse, error)
	AutoStakingVersion(ctx context.Context, in *QueryAutoStakingVersion, opts ...grpc.CallOption) (*QueryAutoStakingVersionResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllZones(ctx context.Context, in *QueryAllZonesRequest, opts ...grpc.CallOption) (*QueryAllZonesResponse, error) {
	out := new(QueryAllZonesResponse)
	err := c.cc.Invoke(ctx, "/nova.icacontrol.v1.Query/AllZones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Zone(ctx context.Context, in *QueryZoneRequest, opts ...grpc.CallOption) (*QueryZoneResponse, error) {
	out := new(QueryZoneResponse)
	err := c.cc.Invoke(ctx, "/nova.icacontrol.v1.Query/Zone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AutoStakingVersion(ctx context.Context, in *QueryAutoStakingVersion, opts ...grpc.CallOption) (*QueryAutoStakingVersionResponse, error) {
	out := new(QueryAutoStakingVersionResponse)
	err := c.cc.Invoke(ctx, "/nova.icacontrol.v1.Query/AutoStakingVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// AllZones returns all the zones registered.
	AllZones(context.Context, *QueryAllZonesRequest) (*QueryAllZonesResponse, error)
	Zone(context.Context, *QueryZoneRequest) (*QueryZoneResponse, error)
	AutoStakingVersion(context.Context, *QueryAutoStakingVersion) (*QueryAutoStakingVersionResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) AllZones(context.Context, *QueryAllZonesRequest) (*QueryAllZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllZones not implemented")
}
func (UnimplementedQueryServer) Zone(context.Context, *QueryZoneRequest) (*QueryZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Zone not implemented")
}
func (UnimplementedQueryServer) AutoStakingVersion(context.Context, *QueryAutoStakingVersion) (*QueryAutoStakingVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutoStakingVersion not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_AllZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.icacontrol.v1.Query/AllZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllZones(ctx, req.(*QueryAllZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Zone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Zone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.icacontrol.v1.Query/Zone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Zone(ctx, req.(*QueryZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AutoStakingVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAutoStakingVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AutoStakingVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.icacontrol.v1.Query/AutoStakingVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AutoStakingVersion(ctx, req.(*QueryAutoStakingVersion))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nova.icacontrol.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllZones",
			Handler:    _Query_AllZones_Handler,
		},
		{
			MethodName: "Zone",
			Handler:    _Query_Zone_Handler,
		},
		{
			MethodName: "AutoStakingVersion",
			Handler:    _Query_AutoStakingVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/icacontrol/v1/query.proto",
}
