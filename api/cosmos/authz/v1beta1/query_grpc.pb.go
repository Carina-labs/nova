// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cosmos/authz/v1beta1/query.proto

package authzv1beta1

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
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error)
	// GranterGrants returns list of `GrantAuthorization`, granted by granter.
	//
	// Since: cosmos-sdk 0.46
	GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error)
	// GranteeGrants returns a list of `GrantAuthorization` by grantee.
	//
	// Since: cosmos-sdk 0.46
	GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error) {
	out := new(QueryGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Grants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error) {
	out := new(QueryGranterGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/GranterGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error) {
	out := new(QueryGranteeGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/GranteeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error)
	// GranterGrants returns list of `GrantAuthorization`, granted by granter.
	//
	// Since: cosmos-sdk 0.46
	GranterGrants(context.Context, *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error)
	// GranteeGrants returns a list of `GrantAuthorization` by grantee.
	//
	// Since: cosmos-sdk 0.46
	GranteeGrants(context.Context, *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grants not implemented")
}
func (UnimplementedQueryServer) GranterGrants(context.Context, *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranterGrants not implemented")
}
func (UnimplementedQueryServer) GranteeGrants(context.Context, *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranteeGrants not implemented")
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

func _Query_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Grants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Grants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Grants(ctx, req.(*QueryGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranterGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranterGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranterGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/GranterGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranterGrants(ctx, req.(*QueryGranterGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranteeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranteeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranteeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/GranteeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranteeGrants(ctx, req.(*QueryGranteeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _Query_Grants_Handler,
		},
		{
			MethodName: "GranterGrants",
			Handler:    _Query_GranterGrants_Handler,
		},
		{
			MethodName: "GranteeGrants",
			Handler:    _Query_GranteeGrants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}
